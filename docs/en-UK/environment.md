# Evironment
Module for standardizing setting, getting, and loading environment variables.

## Set
To set an environment variable, the function `Set` is defined. This function takes in a key and a value to set within the system's environment. If the value can not be set, a `KeyError` is returned.

## Get
To obtain an environment variable, the function `Get` is defined. This function takes a key and returns the variable, if any, stored at the given key within the system environment. If the key is not stored within the system environment, the function panics with an `IOError`.

## Load
To load the environment from a file, the functions `From` or `Load` is defined. `From` takes a `Path` and returns a collection of `Variable`s. If the path exists, the file is loaded into the environment, else the result is loaded from the current system environment. `Load` simply loads the current system environment.