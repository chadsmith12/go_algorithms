# Bubble Sort

Lets say we have the following array: 

```
[1, 3, 7, 4, 2]
```

Quick Sort will start in the `0th` position. It will look at the item next to it and see if it is greater than it. If it is then they should swap positions. If not, then we should continue on.

Example Steps:

* End Index = `len(array)` = 5
* 1 < 3 
* 3 < 7
* 7 > 4 - Swap 7 and 4 
* 7 > 2 - Swap 7 and 2 

After one interation of the loop the largest item is always at the end:

```
[1, 3, 4, 2, 7]
```

Because the largest element is now the last element, in the next iteration we can leave that one off and not have to check it.

Example Steps:

* End Index = `len(array - 1)` = 4
* 1 < 3 
* 3 < 4 
* 4 > 2 - Swap 4 and 2 

After this next iteration our array will be:

```
[1, 3, 2, 4, 7]
```

Following what we just learned we don't need to include 4 in our iteration next time.

Example Steps:

* End Index = `len(array - 2)` = 3
* 1 < 3 
* 3 > 2 - Swap 2 and 2 

After this next iteration our array will be: 

```
[1, 2, 3, 4, 7]
```

Our array is now sorted, but we will keep running until we have one position left.

Example Steps:

* End Index = `len(array - 2)` = 2
* 1 < 2 

We now have a sorted array with a final result of: `[1, 2, 3, 4, 7]`

---

### Generalization

First time we will do `N` checks. The next time through we don't have to check the end so we will do `N-1` amount of checks. This gives us the following:

* `N` Checks
* `N-1` Checks
* `N-2` Checks
* ...
* `N-N+1` Checks

`(N + 1)N/2` = Sum of numbers over an arrange.

`O(N^2 + N) = O(N^2)`
