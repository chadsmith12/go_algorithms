# Arrays

Arrays are a continious part of memory. They are also a fixed amount in memory, which can be confusing or different to a list because a list is dynamic. In Go arrays are a first class data structure that is a value. In Go the follow is an array of numbers that is a length of 8:

```go
array := [8]int
```

## Operations on an array

Arrays have a few specific operations that they can do, though it's important to remember you cannot grow an array.

* Getting at a specific index
* Setting at a specific index
* Deleting at a specific index

Deleting is not really a full deletion and is more just a "set." You set the value to something that signifies that it is deleted to you.

All these operations are a `O(1)`.
