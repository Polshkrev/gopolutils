/*
Events provides an event-like push-pull system.

As an example taken from the [github.com/Polshkrev/parol] package:

	// Register the application end event.
	func registerApplicationEnd(manager *parol.Parol, database table.Table[password.Password]) {
		events.Subscribe(settings.ApplicationEnd, func(any) {
			var except *gopolutils.Exception = database.InsertMany(parol.ObjectToView(manager.Passwords()))
			if except != nil {
				panic(except)
			}
		})
	}

To call an event:

	// Paint the gui.
	func (gui *GUI) Paint() {
		...
		(*gui.window).SetCloseIntercept(func() {
			events.Post(settings.ApplicationEnd, nil)
		})
	}

Or with data:

	// Default callback for when a given card is deleted from its given parent.
	// Returns a callback triggered when a given card is deleted from its given parent.
	func removeCallback(card *Card) Callback {
		return func() {
			events.Post(settings.CardDeleted, card)
		}
	}
*/
package events
