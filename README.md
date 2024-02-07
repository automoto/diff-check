# diff-check
Commandline app for calculating the difference between numbers. We currently support two kinds of calculations:
- Percent Difference: Percent difference is typically used to compare two values and determine how much greater or less one value is than the other
- Percent Change: Percent change is used to express how much a number has increased or decreased relative to another number. It's a way to describe the change in a value over time as a percentage.

This mostly uses the go-difference-check library to do its calculations which is available here https://github.com/sydneyshile/go-difference-check

## Install
The easiest way to install diff-check is to use the provided installation script or with go if you already have it installed.
#### Using the install script:
```shell
wget https://raw.githubusercontent.com/automoto/diff-check/master/scripts/get-diff-check.sh
# you're welcome to examine the script, it just grabs the latest release from github for your OS and installs it
sh get-diff-check.sh
```

#### Install Using Go

If you already have an updated version of go lang, installing via go is easy:
`go get -u github.com/automoto/diff-check/`

If you prefer to install it yourself, you can get the latest release binary directly from github https://github.com/automoto/diff-check/releases and extract it to your `/usr/local/bin`.

## Usage

You can list out the available commands for diff-check by using the `-h` option
```
diff-check -h
```

Which outputs:

```
Usage of diff-check:
  -1 float
        first number to do calculations on
  -2 float
        second number to do calculations on
  -decimal-places uint
        How many decimal places would you like to round your answer to? (default 2)
  -dp uint
        How many decimal places would you like to round your answer to? (default 2)
  -first float
        first number to do calculations on
  -second float
        second number to do calculations on
  -t string
        What type of calculation do you want to do, Percent Difference or Percent Change? valid options are 'change' or 'diff' (default "diff")
  -type string
        What type of calculation do you want to do? valid options are 'change' or 'diff' (default "diff")

```

A simple calculation could be done like so:
```
diff-check -1 22 -2 200
```

_Note: it defaults to percent difference, to check for percent change you can use a command like so:_
```
diff-check -1 22 -2 200 -t change
```
