# gopolutils
 Несколько утилит golang.

## Оглавление
1. [Установить](#установить)
2. [Диаграммы](/docs/ru-RU/diagrams.md)
3. Глобалы
    1. [Логгер](/docs/ru-RU/logger.md)
    2. [Версия](/docs/ru-RU/version.md)
    3. [Ошибка](/docs/ru-RU/exception.md)
4. [Скопления](/docs/ru-RU/collections/collections.md)
    1. Интерфейсы
        1. [Скопление](/docs/ru-RU/collections/collection.md)
        2. [Сопоставление](/docs/ru-RU/collections/mapping.md)
        3. [Размерный](/docs/ru-RU/collections/sized.md)
        4. [Толкование](/docs/ru-RU/collections/view.md)
        5. [Обёртка](/docs/ru-RU/collections/wrapper.md)
        6. Конкурентный
            1. [Скопление](/docs/ru-RU/collections/safe/collection.md)
            2. [Сопоставление](/docs/ru-RU/collections/safe/mapping.md)
    3. Структуры
        1. [Список](/docs/ru-RU/collections/array.md)
        2. [Таблица](/docs/ru-RU/collections/map.md)
        3. [Множество](/docs/ru-RU/collections/set.md)
        4. [Очередь](/docs/ru-RU/collections/queue.md)
        5. [Стек](/docs/ru-RU/collections/stack.md)
        6. [Пара](/docs/ru-RU/collections/pair.md)
        7. Конкурентный
            1. [Список](/docs/ru-RU/collections/safe/array.md)
            2. [Таблица](/docs/ru-RU/collections/safe/map.md)
            3. [Очередь](/docs/ru-RU/collections/safe/queue.md)
            4. [Стек](/docs/ru-RU/collections/safe/stack.md)
5. Файл
    1. [Читать](/docs/ru-RU/fayl/read.md)
    2. [Записывать](/docs/ru-RU/fayl/write.md)
    3. [Адрес](/docs/ru-RU/fayl/path.md)
    4. [Запись](/docs/ru-RU/fayl/entry.md)
    5. [Сжимать](/docs/ru-RU/fayl/compress.md)


### Установить
При правильно настроенном наборе инструментов Go:
```console
go get -u github.com/Polshkrev/gopolutils
```