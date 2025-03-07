# Ошибка
Обёртке ведения типа `error` в golang, основанная на [ошибки Python](https://docs.python.org/3/tutorial/errors.html).

## Оглавление
1. [Использование](#использование)
	1. [Надо](#надо)

## Использование
В качестве возвращаемого значения может использовать модуль ошибки. В пакете коллекций используется множество примеров, например метод `Collection[Type].At(index uint64) (*Type, *gopolutils.Exception)`. Например, метод списка `At`:
```go
func (array Array[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	if index > array.size {
		return nil, gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	}
	return &array.items[index], nil
}
```
### Надо
В этом модуле определена функция `Must`. Эта функция принимает кортеж результата параметра типа и `Exception`. Если переданное исключение не равно nil, функция паникует.