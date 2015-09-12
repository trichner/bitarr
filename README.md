# BitArray
This `Go` package `bitarr` provides a bitarray, can be helpful to efficiently store an array of bools


## Usage


```Go

// create a new bitarray, initialised with all zeros
arr := bitarr.NewArray(65)

arr.Set(33, 1)
arr.Set(23, 1)
arr.Set(18, 1)

arr.Set(3, 0)

// returns 0
arr.Get(12)

// returns 1
arr.Get(33)

// returns 0
arr.Get(17)

// returns 1
arr.Get(18)

// returns 65
arr.Len()

```