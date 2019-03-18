import activitylogs

from event_manager.events import tensorboard

activitylogs.subscribe(tensorboard.TensorboardStartedTriggeredEvent)
activitylogs.subscribe(tensorboard.TensorboardSoppedTriggeredEvent)
activitylogs.subscribe(tensorboard.TensorboardViewedEvent)
activitylogs.subscribe(tensorboard.TensorboardStatusesViewedEvent)
activitylogs.subscribe(tensorboard.TensorboardUpdatedEvent)
activitylogs.subscribe(tensorboard.TensorboardBookmarkedEvent)
activitylogs.subscribe(tensorboard.TensorboardUnBookmarkedEvent)
activitylogs.subscribe(tensorboard.TensorboardDeletedTriggeredEvent)
activitylogs.subscribe(tensorboard.TensorboardArchivedEvent)
activitylogs.subscribe(tensorboard.TensorboardRestoredEvent)
