**Disclaimer:** *I'm new to golang and just starting to learn it, so feel free to give feedback :)*

# Basics of decibinary.go
Define a deci-binary/pseudo-binary number as any base-10 number whose digits are composed of only `1` or `0`. For example, `10, 11, 110, 101` are all deci-binary/pseudo-binary numbers. 

Given an input integer `n`, find the minimal combination of deci-binary/pseudo-binary numbers that sum to that input. For example, `n = 23` would return `[11, 11, 1]`. 

Working as of go 1.12.5

# Usage
Verify everything works as expected with `go test`.

Build using `go build`.

Input integers can be specified as command line arguments:

`./decibinary 123 0 99 -1`

Which would print out:
```sh
$ ./decibinary 123 0 99 -1
  Input:  123
  	 3 steps: [111 11 1]
  Input:  0
  	 1 steps: [0]
  Input:  99
  	 9 steps: [11 11 11 11 11 11 11 11 11]
  Input:  -1
  	 0 steps: []
```

## Examples
* `123` would return `[111 11 1]`
* `91` would return `[11 10 10 10 10 10 10 10 10]`
* `0` would return `[0]`