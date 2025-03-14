# Write
Functions to aid in the writing data to a file.

## Table Of Contents
1. [WriteFile](#writefile)
2. [WriteList](#writelist)
3. [ReadObject](#writeobject)

### WriteFile
For the puposes of documentation, this function is called `WriteFile`. In the source code, this function is simply called `Write`. Эта функция записывает фрагмент байтов в файл с заданным параметром пути. Если файл не может быть записан, возвращается `IOError`.

### WriteList
This function takes in a pointer to a filepath &mdash; either relative or absolute &mdash; and a [View](/docs/en-UK/collections/view.md) representing the list of serialized data to be writen to the file. If the file can not be written, an `IOError` is returned. Another exception that can be returned is, if the given data can not be marshalled, an `IOError` is returned.

### WriteObject
This function takes in a pointer to a filepath &mdash; either relative or absolute &mdash; and a pointer to a type representing the serialized data to be written to the file. If the file can not be written, an `IOError` is returned. Another exception that can be returned is, if the given data can not be marshalled, an `IOError` is returned.