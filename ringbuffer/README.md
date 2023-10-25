# ArrayBuffer or RingBuffer

Think of them like a ArrayList. Except you have an index based head and index based tail.

```
[0............n]
    h   t
```

Everything to the left of the head is null, and everything to the right of the tail is null. Between the head and tail would be your values.


## Removing from the front

If you would like to remove from the front you could then just +1 to the head. That would make the new head 1 index higher, and that is your new head.

## Add to the tail

If you would like to add to the tail you can then just +1 to the tail. This would make the new tail 1 index higher, and you have your new tail.

---

## Head Past Length

It is possible that your head and tail pass each other. You can use the `%` operator to cycle around the array list. `list.tail % len(list)`

## Resizing

If you run into the situation that your tail has poassed your head then it is time for a resize. Example, lets say you went to push something onto the list, and the tail moved to being at the same place as the head or past it, then you should start at the head and write everything to the length and then push it. This would put the head at 0 and the tail at the length, and you can just move the tail one higher now as the capacity increased.


