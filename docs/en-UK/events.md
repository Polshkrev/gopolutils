# Events
Utility defining a standardization of the event pattern.

## Subscribe
Subscribing to an event defines a passed in event &mdash; or callback &mdash; called when the event gets posted. To subscribe to an event, the `EventType` enum will need to be extended. Extended the `EventType` can be defined as below:
```go
package main

import (
	"fmt"

	"github.com/Polshkrev/gopolutils/events"
)

const (
	applicationStart events.EventType = "applicationStart"
	applicationEnd   events.EventType = "applicationEnd"
)

func registerApplicationStart() {
	events.Subscribe(applicationStart, func() { fmt.Println("Application has started.") })
}

func registerApplicationEnd() {
	events.Subscribe(applicationEnd, func() { fmt.Println("Application has ended.") })
}

func init() {
	registerApplicationStart()
	registerApplicationEnd()
}
```
## Posting
Posting an event defines triggering each function subscribed to at the given `EventType`. Such as in the example below:
```go
func main() {
    events.Post(applicationStart)
	/* Application Code */
	events.Post(applicationEnd)
}
```