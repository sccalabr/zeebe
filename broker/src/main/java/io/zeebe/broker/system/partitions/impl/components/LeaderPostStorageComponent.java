package io.zeebe.broker.system.partitions.impl.components;

import io.zeebe.broker.system.partitions.Component;
import io.zeebe.broker.system.partitions.ZeebePartitionState;
import io.zeebe.util.sched.future.ActorFuture;
import io.zeebe.util.sched.future.CompletableActorFuture;

public class LeaderPostStorageComponent implements Component<Void> {

  @Override
  public ActorFuture<Void> open(final ZeebePartitionState state) {
    state
        .getSnapshotStoreSupplier()
        .getPersistedSnapshotStore(state.getRaftPartition().name())
        .addSnapshotListener(state.getSnapshotController());

    return CompletableActorFuture.completed(null);
  }

  @Override
  public ActorFuture<Void> close(final ZeebePartitionState state) {
    state
        .getSnapshotStoreSupplier()
        .getPersistedSnapshotStore(state.getRaftPartition().name())
        .removeSnapshotListener(state.getSnapshotController());
    return CompletableActorFuture.completed(null);
  }

  @Override
  public void onOpen(final ZeebePartitionState state, final Void aVoid) {}

  @Override
  public String getName() {
    return "RegisterSnapshotListener";
  }
}
