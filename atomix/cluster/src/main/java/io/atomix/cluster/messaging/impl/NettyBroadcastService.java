/*
 * Copyright 2018-present Open Networking Foundation
 * Copyright © 2020 camunda services GmbH (info@camunda.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.atomix.cluster.messaging.impl;

import static com.google.common.base.Preconditions.checkNotNull;
import static io.atomix.utils.concurrent.Threads.namedThreads;

import com.google.common.collect.Maps;
import com.google.common.collect.Sets;
import io.atomix.cluster.messaging.BroadcastService;
import io.atomix.cluster.messaging.ManagedBroadcastService;
import io.atomix.utils.net.Address;
import io.atomix.utils.serializer.Namespace;
import io.atomix.utils.serializer.Namespaces;
import io.atomix.utils.serializer.Serializer;
import io.netty.bootstrap.Bootstrap;
import io.netty.buffer.ByteBuf;
import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.DatagramChannel;
import io.netty.channel.socket.DatagramPacket;
import io.netty.channel.socket.InternetProtocolFamily;
import io.netty.channel.socket.nio.NioDatagramChannel;
import java.io.UncheckedIOException;
import java.net.InetSocketAddress;
import java.net.NetworkInterface;
import java.net.SocketException;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.Consumer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/** Netty broadcast service. */
public class NettyBroadcastService implements ManagedBroadcastService {

  private static final Logger LOG = LoggerFactory.getLogger(NettyMessagingService.class);
  private static final Serializer SERIALIZER =
      Serializer.using(
          Namespace.builder()
              .register(Namespaces.BASIC)
              .nextId(Namespaces.BEGIN_USER_CUSTOM_ID)
              .register(Message.class)
              .setCompatible(true)
              .build());

  private final boolean enabled;
  private final InetSocketAddress localAddress;
  private final InetSocketAddress groupAddress;
  private final NetworkInterface iface;
  private EventLoopGroup group;
  private Channel serverChannel;
  private DatagramChannel clientChannel;
  private final Map<String, Set<Consumer<byte[]>>> listeners = Maps.newConcurrentMap();
  private final AtomicBoolean started = new AtomicBoolean();

  public NettyBroadcastService(
      final String host, final Address groupAddress, final boolean enabled) {
    this.enabled = enabled;
    localAddress = new InetSocketAddress(host, groupAddress.port());
    this.groupAddress = new InetSocketAddress(groupAddress.host(), groupAddress.port());
    try {
      iface = NetworkInterface.getByInetAddress(localAddress.getAddress());
    } catch (final SocketException e) {
      throw new UncheckedIOException(e);
    }
  }

  /**
   * Returns a new broadcast service builder.
   *
   * @return a new broadcast service builder
   */
  public static Builder builder() {
    return new Builder();
  }

  @Override
  public void broadcast(final String subject, final byte[] payload) {
    if (enabled) {
      final Message message = new Message(subject, payload);
      final byte[] bytes = SERIALIZER.encode(message);
      final ByteBuf buf = serverChannel.alloc().buffer(4 + bytes.length);
      buf.writeInt(bytes.length).writeBytes(bytes);
      serverChannel.writeAndFlush(new DatagramPacket(buf, groupAddress));
    }
  }

  @Override
  public synchronized void addListener(final String subject, final Consumer<byte[]> listener) {
    listeners.computeIfAbsent(subject, s -> Sets.newCopyOnWriteArraySet()).add(listener);
  }

  @Override
  public synchronized void removeListener(final String subject, final Consumer<byte[]> listener) {
    final Set<Consumer<byte[]>> listeners = this.listeners.get(subject);
    if (listeners != null) {
      listeners.remove(listener);
      if (listeners.isEmpty()) {
        this.listeners.remove(subject);
      }
    }
  }

  private CompletableFuture<Void> bootstrapServer() {
    final Bootstrap serverBootstrap =
        new Bootstrap()
            .group(group)
            .channelFactory(() -> new NioDatagramChannel(InternetProtocolFamily.IPv4))
            .handler(
                new SimpleChannelInboundHandler<>() {
                  @Override
                  public void channelRead0(final ChannelHandlerContext ctx, final Object msg)
                      throws Exception {
                    // Nothing will be sent.
                  }
                })
            .option(ChannelOption.IP_MULTICAST_IF, iface)
            .option(ChannelOption.SO_REUSEADDR, true);

    final CompletableFuture<Void> future = new CompletableFuture<>();
    final ChannelFuture result = serverBootstrap.bind(localAddress);
    result.addListener(
        f -> {
          if (f.isSuccess()) {
            serverChannel = result.channel();
            LOG.debug("Bound broadcasting service to {}", localAddress);
            future.complete(null);
          } else {
            LOG.warn("Failed to bind broadcasting service to {}", localAddress, f.cause());
            future.completeExceptionally(f.cause());
          }
        });
    return future;
  }

  private CompletableFuture<Void> bootstrapClient() {
    final Bootstrap clientBootstrap =
        new Bootstrap()
            .group(group)
            .channelFactory(() -> new NioDatagramChannel(InternetProtocolFamily.IPv4))
            .handler(
                new SimpleChannelInboundHandler<DatagramPacket>() {
                  @Override
                  protected void channelRead0(
                      final ChannelHandlerContext context, final DatagramPacket packet)
                      throws Exception {
                    final byte[] payload = new byte[packet.content().readInt()];
                    packet.content().readBytes(payload);
                    final Message message = SERIALIZER.decode(payload);
                    final Set<Consumer<byte[]>> listeners =
                        NettyBroadcastService.this.listeners.get(message.subject());
                    if (listeners != null) {
                      for (final Consumer<byte[]> listener : listeners) {
                        listener.accept(message.payload());
                      }
                    }
                  }
                })
            .option(ChannelOption.IP_MULTICAST_IF, iface)
            .option(ChannelOption.SO_REUSEADDR, true)
            .localAddress(localAddress.getPort());

    final CompletableFuture<Void> future = new CompletableFuture<>();
    clientBootstrap
        .bind()
        .addListener(
            (ChannelFutureListener)
                f -> {
                  if (f.isSuccess()) {
                    clientChannel = (DatagramChannel) f.channel();
                    LOG.info(
                        "{} joining multicast group {} on port {}",
                        localAddress.getHostName(),
                        groupAddress.getHostName(),
                        groupAddress.getPort());
                    clientChannel
                        .joinGroup(groupAddress, iface)
                        .addListener(
                            f2 -> {
                              if (f2.isSuccess()) {
                                LOG.info(
                                    "{} successfully joined multicast group {} on port {}",
                                    localAddress.getHostName(),
                                    groupAddress.getHostName(),
                                    groupAddress.getPort());
                                future.complete(null);
                              } else {
                                LOG.info(
                                    "{} failed to join group {} on port {}",
                                    localAddress.getHostName(),
                                    groupAddress.getHostName(),
                                    groupAddress.getPort());
                                future.completeExceptionally(f2.cause());
                              }
                            });
                  } else {
                    future.completeExceptionally(f.cause());
                  }
                });
    return future;
  }

  @Override
  public CompletableFuture<BroadcastService> start() {
    if (!enabled) {
      return CompletableFuture.completedFuture(this);
    }
    group = new NioEventLoopGroup(0, namedThreads("netty-broadcast-event-nio-client-%d", LOG));
    return bootstrapServer()
        .thenCompose(v -> bootstrapClient())
        .thenRun(() -> started.set(true))
        .thenApply(v -> this);
  }

  @Override
  public boolean isRunning() {
    return started.get();
  }

  @Override
  public CompletableFuture<Void> stop() {
    if (!enabled) {
      return CompletableFuture.completedFuture(null);
    }
    if (clientChannel != null) {
      final CompletableFuture<Void> future = new CompletableFuture<>();
      clientChannel
          .leaveGroup(groupAddress, iface)
          .addListener(
              f -> {
                started.set(false);
                group.shutdownGracefully();
                future.complete(null);
              });
      return future;
    }
    started.set(false);
    return CompletableFuture.completedFuture(null);
  }

  /** Netty broadcast service builder. */
  public static class Builder implements BroadcastService.Builder {
    private String host;
    private Address groupAddress;
    private boolean enabled = true;

    /**
     * Sets the host interface against which to bind. The port will be taken from the {@code
     * groupAddress}. Defaults to 0.0.0.0.
     *
     * @param host the host interface to bind to
     * @return this builder
     */
    public Builder withHost(final String host) {
      this.host = checkNotNull(host);
      return this;
    }

    /**
     * Sets the group address.
     *
     * @param address the group address
     * @return the broadcast service builder
     */
    public Builder withGroupAddress(final Address address) {
      groupAddress = checkNotNull(address);
      return this;
    }

    /**
     * Sets whether the service is enabled.
     *
     * @param enabled whether the service is enabled
     * @return the broadcast service builder
     */
    public Builder withEnabled(final boolean enabled) {
      this.enabled = enabled;
      return this;
    }

    @Override
    public ManagedBroadcastService build() {
      return new NettyBroadcastService(host, groupAddress, enabled);
    }
  }

  /** Internal broadcast service message. */
  static class Message {
    private final String subject;
    private final byte[] payload;

    Message() {
      this(null, null);
    }

    Message(final String subject, final byte[] payload) {
      this.subject = subject;
      this.payload = payload;
    }

    String subject() {
      return subject;
    }

    byte[] payload() {
      return payload;
    }
  }
}
