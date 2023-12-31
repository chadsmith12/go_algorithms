## Heap DataStructure

This is sometimes called a Priority Queue or a heap.

This is a Binary Tree in that every child and grand child is smaller (MaxHeap), or larger (MinHeap) than the current node

* Whenever a node is added, we must adjust the tree
* Whenever a node is deleted, we must adjust the tree
* There is not travering the tree

### Example

We start with 50. With a minheap the top value must be the smallest. So if we insert into the tree then they must be larger. Example tree down below

       50
      /  \
     71   100
    /  \
   101  80

It's important to note that heaps are ordered but they maintain a weak order.

### Operations

* You can peek and get the smallest (or largest when using a maxheap) element easily in O(1) operation.
* Adding into the tree -
  - We add 3 the tree above. This does not match the minheap conditions cause this is very small.
  - bubble up all the way up. Check the parent. If smaller than the parent, then swap
  - Keep repeating.
* Deleting/Popping from the tree
  - We now delete 3 from the tree. We have a hole in tree
  - Take the very last node in the tree and put it in the top
  - Now we need to heapify down, and put it back in the correct spot
  - Take the minimum of the two children
  - if it's > then we need to swap
  - Bubble down now

### Storing the tree

We have always used node based ways on how to store tree's but using this approach with a heap would be quite slow to find the node we need to work with. We could instead use an array or arraylist so it can be stored.

`[50, 71, 100, 101, 80, 200, 101, 200]`

Parent-Child Relationships:
Index 2 has a child of index 5 and 6.
Index 2 = 100, Index 5 = 200, and Index 6 = 101

Formula: 2i+1 = Left
         2i+2 = Right
This solves how to go and bubble down. Now we need to solve the generic case for bubbling up.
We can use integer division to our advantage here and get a generic formula of: (i-1)/2
Example: (14-1) / 2 = (13/2) = 6
Example: (13-1) / 2 (12/2) = 6
From these examples they both end up at the correct parent.

heapifyDown(0)
heapifyUp(length)

### Updating

If you need to update a priority queue, though usually not an operation that priority queues usually do, then you will need to find a way to store a index->value relaionship, or a value->relationship. Something like a hashmap can work here.
