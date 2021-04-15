# ctest
## Overview
Give color to the output according to the test result. \
#### *Normal test code output*
$ go test -v ./...

#### *Test code output when using ctest*
$ go test -v ./... | ctest

## Usage
```
NAME:
   cTest - Give color to the output according to the test result.

USAGE:
   ctest [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   hmarf

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --run, -r      Do not output '=== RUN: ~'
   --pass, -p     Do not output '--- PASS: ~'
   --fail, -f     Do not output '--- FAIL: ~'
   --help, -h     show help
   --version, -v  print the version
```

## How to use & Example
| command | output|
|:---|:---|
| go test -v ./... \| ctest |string|
| go test -v ./... \| ctest |string|
| go test -v ./... \| ctest |string|
| go test -v ./... \| ctest |string|