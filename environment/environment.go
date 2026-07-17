package environment

import (
	"os"
	"strings"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/fayl"
	"github.com/joho/godotenv"
)

const (
	Seperator string = "=" // Seperator between the system variables.
)

type Variable collections.Pair[string, string] // Representation of a single system variable.

func check(key string, resultChannel chan<- bool) {
	defer close(resultChannel)
	var ok bool
	_, ok = os.LookupEnv(key)
	resultChannel <- ok
}

// Check if the given key is stored in the system's enviornment.
// Returns true if the given key is set within the system's enviornment.
func Check(key string) bool {
	var resultChannel chan bool = make(chan bool, 1)
	go check(key, resultChannel)
	var result bool = <-resultChannel
	return result
}

// Concurrently set a given value at a given key within the system's enviornment.
func set(key, value string, errorChannel chan<- *gopolutils.Exception) {
	defer close(errorChannel)
	var upperValue string = strings.ToUpper(key)
	if Check(upperValue) {
		errorChannel <- nil
		return
	}
	var setError error = os.Setenv(upperValue, value)
	if setError != nil {
		errorChannel <- gopolutils.NewNamedException(gopolutils.KeyError, "%s", setError)
		return
	}
	errorChannel <- nil
}

// Set the given value at the given key within the system enviornment.
// If the given key already exists within the system enviornment, the function returns with a nil exception value.
func Set(key, value string) *gopolutils.Exception {
	var errorChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go set(key, value, errorChannel)
	var except *gopolutils.Exception = <-errorChannel
	return except
}

// Concurrently obtain the given key from the given enviornment.
func get(key string, resultChannel chan<- string, errorChannel chan<- *gopolutils.Exception) {
	defer close(resultChannel)
	defer close(errorChannel)
	if !Check(key) {
		errorChannel <- gopolutils.NewNamedException(gopolutils.IOError, "Can not find key \"%s\".", key)
		resultChannel <- ""
	}
	resultChannel <- os.Getenv(key)
	errorChannel <- nil
}

// Obtain the value set at the given key within the system variables.
// Returns the value set at the given key within the system variables.
func Get(key string) string {
	var resultChannel chan string = make(chan string, 1)
	var errorChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go get(key, resultChannel, errorChannel)
	var except *gopolutils.Exception = <-errorChannel
	if except != nil {
		panic(except)
	}
	var result string = <-resultChannel
	return result
}

// Load a the system enviornment from a given file.
// Returns an [Enviornment] from the given [fayl.Path].
// If the given file does not exist, the enviornment is loaded from the system's internal enviornment.
// If the .env file can not be read, this function `panics` with an [OSError].
// If the key is already in the mapping, instead of just quietly not inserting into the mapping, this function `panics` with a [KeyError].
func From(file *fayl.Path) collections.View[Variable] {
	if !file.Exists() {
		return Load()
	}
	var result collections.Collection[Variable] = collections.NewArray[Variable]()
	var raw map[string]string
	var readError error
	raw, readError = godotenv.Read(file.String())
	if readError != nil {
		panic(gopolutils.NewNamedException(gopolutils.OSError, "%s\n", readError.Error()))
	}
	var key, value string
	for key, value = range raw {
		result.Append(Variable(*collections.NewPair(key, value)))
	}
	return result
}

// Concurrently split the given raw enviornment variable based on the default seperator.
func split(variable string, resultChannel chan<- Variable, errorChannel chan<- *gopolutils.Exception) {
	defer close(resultChannel)
	defer close(errorChannel)
	var split []string = strings.Split(variable, Seperator)
	if len(split) > 2 {
		errorChannel <- gopolutils.NewNamedException(gopolutils.ValueError, "Can not seperate the variable \"%s\".", variable)
		return
	}
	resultChannel <- Variable(*collections.NewPair(split[0], split[1]))
	errorChannel <- nil
}

// Split a given string representation of a given enviornment variable.
// Returns the given string representation as a [Variable].
// If the variable can not be split, a [ValueError] is returned.
func Split(variable string) (Variable, *gopolutils.Exception) {
	var resultChannel chan Variable = make(chan Variable, 1)
	var errorChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go split(variable, resultChannel, errorChannel)
	var except *gopolutils.Exception = <-errorChannel
	var result Variable = <-resultChannel
	return result, except
}

// Concurrently load the raw system enviornment.
func loadResult(resultChannel chan<- []string) {
	defer close(resultChannel)
	resultChannel <- os.Environ()
}

// Load the raw system enviornment.
// Returns a slice of raw enviornment variables.
func loadRaw() []string {
	var resultChannel chan []string = make(chan []string, 1)
	go loadResult(resultChannel)
	var result []string = <-resultChannel
	return result
}

// Concurrently load the system enviornment as an [Enviornment].
func assign(resultChannel chan<- collections.Collection[Variable], errorChannel chan<- *gopolutils.Exception) {
	defer close(resultChannel)
	defer close(errorChannel)
	var raw []string = loadRaw()
	var result collections.Collection[Variable] = collections.NewArray[Variable]()
	var i int
	for i = range raw {
		var variable collections.Pair[string, string] = collections.Pair[string, string](gopolutils.Must(Split(raw[i])))
		result.Append(Variable(variable))
	}
	resultChannel <- result
	errorChannel <- nil
}

// Load the system enviornment variables as an [Enviornment].
// Returns the system enviornment variables as an [Enviornment].
// If the enviornment can not be parsed or loaded, a [KeyError] is returns.
func load() (collections.View[Variable], *gopolutils.Exception) {
	var resultChannel chan collections.Collection[Variable] = make(chan collections.Collection[Variable], 1)
	var errorChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception)
	go assign(resultChannel, errorChannel)
	var except *gopolutils.Exception = <-errorChannel
	if except != nil {
		return nil, except
	}
	var result collections.Collection[Variable] = <-resultChannel
	return result, nil
}

// Load the system enviornment from the internal enviornment.
// Returns an [Enviornment] based on the system's internal enviornment.
func Load() collections.View[Variable] {
	return gopolutils.Must(load())
}
