package io.zeebe.broker.system.partitions.impl.components;

import io.zeebe.broker.system.partitions.Component;
import io.zeebe.broker.system.partitions.ZeebePartitionState;
import io.zeebe.util.sched.future.ActorFuture;
import io.zeebe.util.sched.future.CompletableActorFuture;

public class FollowerPostStorageComponent implements Component<Void> {

  @Override
  public ActorFuture<Void> open(final ZeebePartitionState state) {
    state.getSnapshotController().consumeReplicatedSnapshots();
    return CompletableActorFuture.completed(null);
  }

  @Override
  public ActorFuture<Void> close(final ZeebePartitionState state) {
    return CompletableActorFuture.completed(null);
  }

  @Override
  public void onOpen(final ZeebePartitionState state, final Void aVoid) {}

  @Override
  public String getName() {
    return "ConsumeReplicatedSnapshots";
  }
}
