# События
Утилита, определяющая стандартизацию шаблона событий.

## Подписка
Подписка на событие определяет переданное событие &mdash; или обратный вызов &mdash; вызываемый при отправке события. Для подписки на событие необходимо расширить перечисление `EventType`. Расширенное перечисление `EventType` можно определить следующим образом:
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
	events.Subscribe(applicationStart, func(any) { fmt.Println("Application has started.") })
}

func registerApplicationEnd() {
	events.Subscribe(applicationEnd, func(any) { fmt.Println("Application has ended.") })
}

func init() {
	registerApplicationStart()
	registerApplicationEnd()
}
```
## Отправка события
Отправка события определяет запуск каждой функции, на которую оформлена подписка, для заданного `EventType`. Например, как в приведенном ниже примере:
```go
func main() {
    events.Post(applicationStart, nil)
	/* Application Code */
	events.Post(applicationEnd, nil)
}
```