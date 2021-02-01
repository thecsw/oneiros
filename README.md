# ONEIROS

## What is this

A new way of logging in

## Installation

``` sh
go get -u github.com/thecsw/oneiros
```

## Usage

It will return a 16 digit number as a `*math/big.Int`

``` go
oneiros.Generate()
```

It takes about ~380ns to generate one. Tested on MacBook Pro M1 (ARM mode).
