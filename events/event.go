package events

import (
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

// Lookup name of an event.
type EventType gopolutils.StringEnum

// Type alias for an event to be called.
type Event func(any)

// Type alias for a mapping of event types to a collection of events.
type EventManager collections.Mapping[EventType, collections.Collection[Event]]

var (
	events EventManager = NewEventManager() // Default event manager.
)

// Construct a new event manager.
// Returns a new event manager.
func NewEventManager() EventManager {
	return collections.NewMap[EventType, collections.Collection[Event]]()
}

// Subscribe to a given event type with a given event.
func Subscribe(eventType EventType, event Event) {
	if !events.HasKey(eventType) {
		events.Insert(eventType, collections.NewArray[Event]())
	}
	var subscribedEvents *collections.Collection[Event] = gopolutils.Must(events.At(eventType))
	(*subscribedEvents).Append(event)
}

// Trigger each event stored at the given event type.
func Post(eventType EventType, data any) {
	if !events.HasKey(eventType) {
		return
	}
	var subscribedEvents *collections.Collection[Event] = gopolutils.Must(events.At(eventType))
	var i int
	for i = range (*subscribedEvents).Collect() {
		var subscribedEvent Event = (*subscribedEvents).Collect()[i]
		subscribedEvent(data)
	}
}

// Obtain the current [EventManager].
// Returns a mapping of [EventType]s to a [collections.View] of [Event]s.
func Manager() EventManager {
	return events
}

// Obtain all the [Event]s stored at the given [EventType].
// Returns all the [Event]s stored at the given [EventType].
func Events(eventType EventType) collections.View[Event] {
	var i int
	var result collections.View[Event]
	for i = range events.Collect() {
		var bucket collections.Pair[EventType, collections.Collection[Event]] = events.Collect()[i]
		if (*(*bucket).First()) != eventType {
			continue
		}
		result = (*(*bucket).Second())
	}
	return result
}

// Change the default [EventManager].
func SetEventManager(manager EventManager) {
	events = manager
}
