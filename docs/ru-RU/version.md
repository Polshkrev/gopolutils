# Версия

Возникнув в семантическом версионировании, этот глобальный объект может использоваться для хранения версий. Альтернативно, его можно использовать с библиотекой флагов, чтоб легко создать профессионально выглядящий флаг версии.

## Оглавление
1. [Использование](#использование)
    1. [Конструкция](#конструкция)
    2. [Модификация](#модификация)
    3. [Сравнивать](#сравнивать)
## Использование
### Конструкция
Существует много методов построения версии. По умолчанию используется функция `NewVersion`. Все функции конструктора вернут пойнтер на версии, выделенный в куче. Функция `NewVersion`, устанавливающая строковые свойства нового версии в `nil`, а каждое из числовых свойств в ноль. Соглашение об именовании каждой из функций конструктора следующее:
- Функции словом "convert" создает версию со свойствами чисел, заданными для указанных параметров.
- Функции словом "strings" говорят сами за себя, поскольку они создают каждое из свойств строки версии.
- Любые другие свойства, заданные в конструкторе, именуются эксплицитно.
### Модификация
Чтоб установить любого из свойств версии с каждым свойством связан определенный сеттер.
Чтоб изменить любого свойства версии, существуют следующие функции:
- `Release`: Увеличьте основное свойство версии и установите ее свойства `minor` и `patch` на ноль.
- `Update`: Увеличьте свойство версии `minor` и установите свойство версии `patch` на ноль, не влияя на свойство версии `major`.
- `Patch`: Увеличьте значение свойства `patch` версии, не затрагивая при этом `major` или `minor` свойства версии.
- `Publish`: Установите значение `major` версии на единицу. Если версия оценивается как уже опубликованная, возникает ошибка и никакие свойства не изменяются.
### Сравнивать
Каждое из числовых свойств версии имеет метод сравнения. Чтобы сравнить с другой одной версии доступен метод `Compare`. Эти методы сравнения определяют, больше ли числовые свойства версии или равны заданному параметру сравнения. Метод `IsPublic` доступен и используется в методе `Publish`. Метод `IsPublic` возвращает значение `true`, если свойство версии «major» больше или равно единице. Чтоб определить того, равна ли версия нулю, доступен метод `IsZero`.