# Trees

In a sense a tree is just a linked list with multiple paths. Examples of a tree:

* File System
* DOM
* Abstract Syntax Tree

The basic definition of a tree would be like:

```
node<T>
    value: T
    children: T[]

```

## Terminology:

* Height - Height of a tree is described as the height from the root to the most grandchild node from that root. With this, not all trees have to be balanced.
* Binary Tree - A tree with at max two children. A lot of times when using Binary Tree you'll see `left` and `right`
* Binary Search Tree - A Tree in which has a specific ordering and at most 2 children.
* Leaves - A node with no more children.
* Balanced - A tree is perfectly balanced when any node's left and right children have the same height.
* Branching Factor - The amount of children a tree has.

## Traversals:

There are different ways to traverse a tree. They really just depend on when you visit the node.

* Pre Order - First Visit node, and then do the recursion. The root will be at the beginning
* In Order - You first recurse all the way down the left side, Visit the node, and then you will recurse all the way down the right side. Root will be in the middle.
* Post Order - You first recurse both the left and right side, then visit the node. Root will be at the end.

## Running Time

The input is the size of the tree. There are N Nodes. The traversal would be `O(N)`.


