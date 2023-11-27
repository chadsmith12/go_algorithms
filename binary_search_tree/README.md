# Binary Search Tree

This is basic implementation of Binary Search Tree in Go. This implementation is done using generics and can accept any type that you would like to store in a Binary Search Tree. Currently is supports the following methods:

* Insert(v) - Inserts the item v into the tree
* Find(v) - Attempts to find the item v in the tree. Will return true, if it finds the value, false if it can't find it.
* Delete(v) - Attempts to delete the item v from the tree.
* InorderSlice() - Returns the tree as a slice using in order traversal. This would essentially give you a sorted slice from the BST.

## Usage

Getting started is pretty easy. You just need to initialize a new BST using the `NewBST` and just pass in the `Compparator`. The `Comparator` is needed because this BST is designed to work off of any type, so the tree needs a way to be able to compare the different nodes. This `Comparator` function is the same that the standard Go functions would use. Just need a method that takes in the two values and returns and integer. Example below shows creating a BST for integers.

```go
func intComparator(a, b) int {
    return a - b
}

tree := bst.NewBST[int](intComparator)
```

This will give you a basic empty tree that you can start inserting items into.
