# Version

A semantic versioning module. This module can be used to store versioning objects for feature flags. Or it can used with the flag module to easily create a professional looking version flag.

## Table Of Contents
1. [Usage](#usage)
    1. [Construction](#construction)
    2. [Modify](#modify)
    3. [Compare](#compare)
## Usage
### Construction
There are many ways to construct a version object. The default is to use the `NewVersion` function. All of the constructor functions will return a heap-allocated pointer to a version object. the `NewVersion` function sets the string properties of the new version object to `nil` and each of the number properties to 0. The naming convention of each of the constructor functions is as follows:
- A function with the "convert" keyword will construct the object with the number properties set to the given parametres.
- A function with the "strings" keyword is self-explanatory as it constructs each of the string properties in the version object.
- Any additional properties that are set in the constructor are named explicitly.
### Modify
To set any of the properties in the version object, each property has a specific setter associated with it.
To change any of the properties within the object the following functions are available:
- Release: Increment the object's major value and set its minor and patch values to 0.
- Update: Increment the object's minor value and set the object's patch version to 0 while not affecting the object's major value.
- Patch: Increment the object's patch value while not affecting the object's major or minor values.
- Publish: Set the object's major value to 1. If the object is evaluated to have already been published, an error is raised and no values are changed.
### Compare
Each of the object's number properties have a comparison method. To compare a version object to another a `Compare` method is available. These compare methods determine if the object's numeric properties are greater than or equal to the given compare parametre. An `IsPublic` method is available and is used in the `Publish` method. The `IsPublic` method returns true if  the object's major property is greater than or equal to 1. To determine if the object is equal to 0, an `IsZero` method is available.