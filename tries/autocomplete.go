package autocomplete

type characterNode struct {
	children map[rune]*characterNode
	endWord  bool
	value    *rune
}

func newCharacter(value rune) *characterNode {
	return &characterNode{
		value:    &value,
		children: make(map[rune]*characterNode),
	}
}

// Autocomplete represents a builder to build a dictionary of words and to suggest words off a word
// This is built using a trie internally
type Autocomplete struct {
	root *characterNode
}

// Initializes a new autocomplete to so you can start adding words to the autocomplete dictionary
func New() *Autocomplete {
	return &Autocomplete{
		root: &characterNode{
			children: make(map[rune]*characterNode),
		},
	}
}

// Adds a new word to the autocomplete dictionary.
// Returns back an autocomplete instance for easy method chaining
func (a *Autocomplete) AdddWord(word string) *Autocomplete {
	node := a.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = newCharacter(ch)
		}
		node = node.children[ch]
	}
	node.endWord = true

	return a
}

// Takes in a word/prefix and gets the predictions/suggestions to complete based off this word.
func (a *Autocomplete) Predict(word string) []string {
	foundWords := make([]string, 0, 10)

	subTree := a.getSubTree(word)
	if subTree == nil {
		return foundWords
	}
	a.suggest(subTree, word, &foundWords)

	return foundWords
}

func (a *Autocomplete) suggest(node *characterNode, prefix string, foundWords *[]string) {
	if node.endWord {
		*foundWords = append(*foundWords, prefix)
	}

	for ch, currNode := range node.children {
		a.suggest(currNode, prefix+string(ch), foundWords)
	}
}

func (a *Autocomplete) getSubTree(word string) *characterNode {
	firstCharacter := rune(word[0])

	node := a.root.children[firstCharacter]
	if node == nil {
		return nil
	}

	for _, value := range word[1:] {
		node = node.children[rune(value)]
	}

	return node
}
