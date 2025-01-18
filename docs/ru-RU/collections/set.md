# Множество
Реализация множества структуры данных. Под капотом эта структура использует интерфейс [сопоставления](/docs/en-UK/collections/mapping.md).

## Оглавление
1. [Добавить](#добавить)
2. [Удалить](#удалить)
3. [Доступить](#доступить)
4. [Утилиты](#утилиты)
5. [Реализовать](#реализовать)

### Добавить
Чтобы добавить к множеству есть два метода можно использовать; есть &mdash; конечно &mdash; метод `Append`. Этот метод добавляет один уникальный элемент к набору. Есть тоже метод `Extend`. Этот метод добавляет несколько уникальных элементов к множеству в форме `View`.

### Удалить
Чтоб удалить элемента из множества доступен метод `Remove`. Этот метод принимает индекс, по которому в множестве находится удаляемый элемент. Если множество пуст или заданный индекс больше размера множества, возвращается ошибка `IndexOutOfRange`. В редком случае, когда данные по указанному индексу не могут быть найдены, возвращается `IndexError`. Тоже доступен метод `Discard`. Этот метод представляет собой метод удаления, который принимает элемент для удаления и &mdash; вместо того, чтобы возвращать ошибку или просто аварийно завершать работу &mdash; этот метод просто не изменяет множество и просто возвращает значение, если заданный элемент не найден в множестве или если множество пуст. В **редком критическом случае** программа **вылетит**, если данный элемент **не может быть удален** из множества.

### Доступить
Доступ к множеству не допускается. Не определен метод, позволяющий получить доступ к элементу внутри множества. Определен метод `At`, но он всегда возвращает `NotImplementedError`.

### Утилиты
В этой структуре определены два других метода: `Difference` и `Intersection`. Как следует из названия, метод `Difference` принимает множество операндов и возвращает пойнтер на новое множество, содержащий все элементы в множестве операндов, которые не содержатся в исходном множестве. Метод `Intersection` является полной противоположностью метода `Difference`. Метод `Intersection` принимает множество операндов и возвращает пойнтер на новое множество, содержащий все элементы, которые содержатся в обоих множествах.

### Реализовать
Эта структура реализует интерфейсы `Collection`, `View` и `Sized`.