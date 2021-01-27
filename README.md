# ONEIROS

## What is this

A new way of logging in

## Installation

``` sh
go get -u github.com/thecsw/oneiros
```

## Usage

It will return a 16 digit number represented as a string.

``` go
oneiros.Generate()
```

It takes about ~1700ns to generate one. Tested on MacBook Pro M1 (ARM mode).
