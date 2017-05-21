Go Logger Interface
===================

Golang is gaining a lot of traction at work so here is my basic program using interfaces.

My experience is that Golangs interface comes across as a strongly typed duck typing! Quite cool actually.

Usage
-----
* Download source code of this project from github
* From command line type "go build"
* From command line type "interface_logger"

You will see some log messages!

The Code
--------
### logger.go
Contains the interface which consists of four methods
* Debug
* Info
* Warn
* Error

Each interface method accepts a string parameter.
Also contains a "factory" that will return a logger adhering to the interface.

### console_logger.go
An implimentation of a console logger

### file_logger.go
An implimentation of a file logger

### main.go
Uses the factory to create a logger adhering to the logger interface.
Also creates a concrete type of ConsoleLogger

Contributing
------------
Please see [CONTRIBUTING.md][contributor] and read the [CODE_OF_CONDUCT.md][conduct]

License and Authors
-------------------
Please see [LICENSE][licence]
Authors: Chris Sullivan

[contributor]: https://github.com/chrisgit/go-interface_logger/blob/master/CONTRIBUTING.md
[conduct]: https://github.com/chrisgit/go-interface_logger/blob/master/CODE_OF_CONDUCT.md
[licence]: https://github.com/chrisgit/go-interface_logger/blob/master/LICENSE