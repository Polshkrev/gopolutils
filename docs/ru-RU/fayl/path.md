# Адрес
Реализация [Python Pathlib](https://docs.python.org/3/library/pathlib.html)-а на go.

## Оглавление
1. [Использование](#использование)
    1. [Конструкция](#конструкция)
        1. [NewPath](#newpath)
        2. [PathFrom](#pathfrom)
        3. [PathFromParts](#pathfromparts)
    2. [Модификация](#модификация)
## Использование
### Конструкция
Чтобы конструировать нового адреса можно использовать три функции: `NewPath`, `PathFrom` и `PathFromParts`. Некоторые методы адреса тоже вернут новый адрес, например `Absolute`, `Parent` и `Root`.
#### NewPath
Используя `NewPath` создаст адрес из текущего рабочего каталога. Если текущий рабочий каталог не может быть получен, `OSError` выводится в стандартный поток ошибок и программа завершает работу.
#### PathFrom
Используя `PathFrom` создаст новый адрес из заданной строки адреса. Этот конструктор не терпит неудачу.
#### PathFromParts
Используя `PathFromParts` создаст адрес из заданного имени папки, имени файла и суффикса файла. Например:
```go
var path *gopolutils.Path = gopolutils.PathFromParts(".", "test", "txt")
// The path string should equal: "{currentDirectory}/test.txt"
```
Этот конструктор использует разделитель адресов, зависящий от операционной системы.

### Модификация
Существует два простых способа изменения адреса, такие как `Append` и `AppendAs`. Используя `Append` добавит адрес. Используя `AppendAs` добавит строку.