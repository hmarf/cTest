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
   ctest [global options] [go test options]

AUTHOR:
   hmarf

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -r          Do not output '=== RUN: ~'
   -p          Do not output '--- PASS: ~'
   -f          Do not output '--- FAIL: ~'
   --help, -h  show help
```

## How to install
```
$ brew tap hmarf/tap
$ brew install hmarf/tap/ctest
```

## How to use & Example
| command | explain | output |
|:---|:---|:---|
| go test -v ./... \| ctest | Output test results in color | <img src="https://github.com/hmarf/cTest/blob/master/image/no_option.png" width="800px">|
| go test -v ./... \| ctest -r | Output test results in color <br> If the first characters of the output is `=== RUN:`, the line is not printed | <img src="https://github.com/hmarf/cTest/blob/master/image/option_r.png" width="800px">|
| go test -v ./... \| ctest -p -r | Output test results in color <br> If the first characters of the output is `=== RUN:` and `--- PASS:`, the line is not printed | <img src="https://github.com/hmarf/cTest/blob/master/image/option_p_r.png" width="800px">|
