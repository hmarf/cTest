# ctest
## Overview
Give color to the output according to the test result.
#### *Normal test code output*
<img src="https://github.com/hmarf/cTest/blob/master/image/normal.png" width="700px">

#### *Test code output when using ctest*
<img src="https://github.com/hmarf/cTest/blob/master/image/no_option.png" width="700px">

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
| command | explain | output |
|:---|:---|:---|
| go test -v ./... \| ctest | Output test results in color | <img src="https://github.com/hmarf/cTest/blob/master/image/no_option.png" width="700px">|
| go test -v ./... \| ctest -r | Output test results in color <br> If the first characters of the output is `=== RUN:`, the line is not printed | https://github.com/hmarf/cTest/blob/master/image/option_r.png |
| go test -v ./... \| ctest -p -r | Output test results in color <br> If the first characters of the output is `=== RUN:` and `--- PASS:`, the line is not printed | https://github.com/hmarf/cTest/blob/master/image/option_p_r.png |