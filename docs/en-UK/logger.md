# Logger
A logger.

## Table Of Contents
1. [Usage](#usage)
    1. [Construction](#construction)
    2. [Logging](#logging)
    3. [Destruction](#destruction)

## Usage
### Construction
To construct a new logger object, use the `NewLogger` function. This function will return a heap allocated logger object. The parametres of the function is the name and default logging level enum value. Just as in my [wrapper](https://github.com/Polshkrev/Utilities/tree/main/docs/en-UK/globals/log) around the python logging package, the level parametre passed to this function sets the minimum level that the logger will output.
To setup the logger correctly, you will need to either add `stdout` or a file. To do so, you will either need to call the `AddConsole` method or `AddFile` method respectively. To use both `stdout` and a file, you need to call the `FullSetup` method. Each of these methods can fail. If either `stdout` or a file can not be added, an `IOError` will be returned.
To modify the logger after it has been constructed, a public getter and setter for the name and level properties are defined.
### Logging
To log a message, you need to call the `Log` method. This method takes a log message without a newline, and a level with which the message will be logged.
### Destruction
If a file is bound to the logger, you will need to close the file. The `Close` is typicalled called deffered to close any file bound to the logger. This method will not affect the standard output if it is bound to the logger.