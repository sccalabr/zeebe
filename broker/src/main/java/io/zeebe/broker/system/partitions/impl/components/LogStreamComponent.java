/*
 * Copyright Camunda Services GmbH and/or licensed to Camunda Services GmbH under
 * one or more contributor license agreements. See the NOTICE file distributed
 * with this work for additional information regarding copyright ownership.
 * Licensed under the Zeebe Community License 1.0. You may not use this file
 * except in compliance with the Zeebe Community License 1.0.
 */
package io.zeebe.broker.system.partitions.impl.components;

import io.zeebe.broker.system.partitions.Component;
import io.zeebe.broker.system.partitions.ZeebePartitionState;
import io.zeebe.logstreams.log.LogStream;
import io.zeebe.util.sched.future.ActorFuture;

public class LogStreamComponent implements Component<LogStream> {

  @Override
  public ActorFuture<LogStream> open(final ZeebePartitionState state) {
    return LogStream.builder()
        .withLogStorage(state.getAtomixLogStorage())
        .withLogName("logstream-" + state.getRaftPartition().name())
        .withNodeId(state.getNodeId())
        .withPartitionId(state.getRaftPartition().id().id())
        .withMaxFragmentSize(state.getMaxFragmentSize())
        .withActorScheduler(state.getScheduler())
        .buildAsync();
  }

  @Override
  public ActorFuture<Void> close(final ZeebePartitionState state) {
    state.getComponentHealthMonitor().removeComponent(state.getLogStream().getLogName());
    return state.getLogStream().closeAsync();
  }

  @Override
  public void onOpen(final ZeebePartitionState state, final LogStream logStream) {
    state.setLogStream(logStream);

    if (state.getDeferredCommitPosition() > 0) {
      state.getLogStream().setCommitPosition(state.getDeferredCommitPosition());
      state.setDeferredCommitPosition(-1);
    }
    state.getComponentHealthMonitor().registerComponent(logStream.getLogName(), logStream);
  }

  @Override
  public String getName() {
    return "logstream";
  }
}
