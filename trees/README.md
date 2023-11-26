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

When doing these traversals we are doing a Depth First Search.

## Running Time

The input is the size of the tree. There are N Nodes. The traversal would be `O(N)`.

## Breadth-First Seach

A Breadth-First Search works the opposite of a Depth First Search. When you use a DFS and using recurrsion you will notice that you are implicitly using a stack. When you start using a BFS the opposite happens and now you are going to be using a Queue.

It's important to note that Depth First Search will preserve shape, while a BFS will not preserve shape. So if you were to compare that two binary trees are in-fact the same in contents and shape, then you would want to use a DFS.
