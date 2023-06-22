package events

import "errors"

type Subscriber *func(event Event) (Event, error)

var (
	ErrNoSubscribers = errors.New("no subscribers for event type")
)

var (
	subscribers = make(map[EventType][]Subscriber)
)

// Subscribe adds a subscriber to the list of subscribers for the event type
func Subscribe(eventType EventType, subscriber Subscriber) {
	subscribers[eventType] = append(subscribers[eventType], subscriber)
}

// Unsubscribe removes a subscriber from the list of subscribers for the event type
func Unsubscribe(eventType EventType, subscriber Subscriber) {
	for i, s := range subscribers[eventType] {
		if s == subscriber {
			subscribers[eventType] = append(subscribers[eventType][:i], subscribers[eventType][i+1:]...)
			break
		}
	}
}

// Broadcast sends an event to all subscribers of the event type
func Broadcast(event Event) (results []Event, err error) {
	h := event.Header()
	if _, ok := subscribers[h.Type]; !ok {
		return nil, ErrNoSubscribers
	}

	for _, subscriber := range subscribers[h.Type] {
		res, err := (*subscriber)(event)
		if err != nil {
			return nil, err
		}

		results = append(results, res)
	}

	return results, nil
}
