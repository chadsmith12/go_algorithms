# Sorting

## Bubble Sort

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

---

## QuickSort

QuickSort uses a Divide and Conqure algorithm.

We will pick some element out of an array, we will call it `p`. In this case, we will pick the last element of the array. Anything that is smaller or equal to the pivot will be on one side of the array, and everything that is larger than the pivot on the other side of it. 

You will split the array, not including the pivot, and break it down into two arrays and repeat the operation on the two arrays.

### Example:

We have an array that is 32 elements large, 0-31. We pick a pivot in the middle so `p` is 16. We break this down so we have two sub arrays. This will be: 0-15 and 17-31. Note we don't include the pivot.

The pivot on the first half will then be 8. Going down this half if we keep breaking it down you will then have 1-7 and 9-15.

Breaking it down further, you have 1-7, a pivot of 4, and 9-15 with a pivot of 12.

Breaking it down furthe, you have 1-3, a pivot of 2. this breaks down to only one element. At this point when you are at the bottom part, all of the pieces are now sorted.

## Running Time

You scan each input at least once. Every break down you go down each input. So you're down `N` operations until `N=1`, which will give us the equation: `n/2^k = 1`. This ends up equaling `logn`. So therefor the running for this would be `O(nlogn)` amount of times.  You will do `logn` operations `n` times.

Though there is a catch that if your array is reverse sorted and you ended up picking the pivot of the smallest element, it could end up being `O(n^2)`. There is a strategy to help with this and a simple strategy is just pick the middle element. So there is a chance to pick the worst condition, though the chances quite low.


