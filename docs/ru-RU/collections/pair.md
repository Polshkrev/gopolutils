# Пара
Реализация кортежа типа pair де-факто. Вдохновлено [std::pair](https://en.cppreference.com/w/cpp/utility/pair) в c++.

## Оглавление
1. [Доступить](#доступить)
2. [Утилиты](#утилиты)

### Доступить
Эта структура предоставляет несколько методов доступа. Для каждого из двух определенных свойств предоставляется один метод доступа. Эти методы возвращают пойнтеры на данные, хранящиеся в каждом соответствующем свойстве.

### Утилиты
Как и в оригинальной реализации c++, эта структура имеет метод `Swap`. Этот метод принимает заданное с теми же дженериками типа, что и оригинал, и меняет местами каждое из значений; первое с первым, второе со вторым.