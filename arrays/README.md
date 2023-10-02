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

---

### Searching Algorithms

* Linear Search - This is the most basic searching algorithm that searches the array for a certain value and tells you if the items is inside of the array. This algorithm will linearly search this array until it finds the item and return that it found it. If it can not find the item then it will return back that the item was not found in the array. Because we are searching the array linearly we have to do that n times, so therefore the algorithm is `O(n)`.
* Binary Search - When you have your data source already sorted then you can start to make some assumptions on how to search your data. This is where Binary Search comes in. The basic idea of a binary search is that you keep track of a low point, high point, calculate a midpoint, and see if the value you're searching for is at that midpoint. If it is, then we are done. If the value is > the value we are searching for then we know we can half our search space and search from the midpoint + 1 to high. Find the Midpoint there, and repeat. We can use the same technique if the value is < the value at the midpoint. By doing this we are essentially halving our search space each time, so therefore the algorithm is `O(logn)`
