# go-cdg

### Basics

* go is static typed language

|type  |example    |
|:---: |:---:|
|bool  | true false|
|string| "test"    |
|int   | 0 -1 999  |
|float64| 10.00 0.0009 -100.009|

* Переменные можно инициализировать вне функции, но нельзя присвоить переменную.
* Переменные сначала необходимо инициализировать с помощью оператора ': =' или синтаксиса 'var variableName type'.
* Files in the same package do not have to be imported into each other. Example two separate files with package name main.

### Arrays and Slices

* Array - fixed length list of things
* Slice An array that can grow or shrink (сокращаться)
* every element must be of same type

|Arrays                 |Slices             |
|:---:                  |:---:              |
|Primitive data sructure|Can grow and shrink|
|Can't resize           |Used 99% of the time for lists of elements|
|Rarely used directly||


|Value Types|Reference Types|
|:---:      |:---:          |
|int        |slices         |
|float      |maps           |
|string     |channels       |
|bool       |pointers       |
|structs    |functions      |