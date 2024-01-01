# Tries

The easiest way to visualise a trie is to think of auto-complete. 

## English Language Tree

You will start with a root. This doesn't really have a value. Lets say we were adding the word "cat."

```
          *
         /
        c
       /
      a
     /
    t
     \
      *
```

Notice that 't' gets marked, and it is marked with an extra child, an asterics child. This marks that this parth you're on is a word. Another approach is to have a "boolean" flag that marks it as `isWord`.

Now we want to add the word "cats". We see we have a c link, a link, and a t link, so we just need to add a s link and mark that node as a potential word.

```
          *
         /
        c
       /
      a
     /
    t
   /
  s *(isWord)
```

Next we add the world "cattle"

```
            *
           /
          c
         /
        a
       /
      t
     / \
    s   t
         \
          l 
           \
            e *(isWord)
```

If someone starts typing the letter 'c' then we start from the c branch and see if have anything there. This would be represented with index 2. Example of how to find this out.

```go
// create a zero value of a
const zero rune = 'a'

func index(str rune) int {
    return int(str - zero)
}

fmt.Println(index('c')) // this will give you index 2
```

If we want to get the list of words we will need to a depth first search or even a BFS. If you want your words to come out sorted, then using a preorder dfs will keep the words sorted like that.

### Insertion

Simple psuedo code for insertion:

```js
insertion(str) {
    curr = head
    for c in str {
        if curr[c] {
            curr = curr[c]
        } else {
            node = createNode()
            current[c] = node
            curr = node
        }
    }
}

```

We walk to the end and insert if we need to insert.
