package autocomplete

import (
	"testing"
)

func assertRootChildren(a *Autocomplete, t *testing.T, expected int) {
	if len(a.root.children) != expected {
		t.Errorf("len() of roots children is %d, expected %d\n", len(a.root.children), expected)
	}
}

func assertCharacterChildren(a *Autocomplete, t *testing.T, value rune, expected int) {
	valueChildren := a.root.children[value]
	if valueChildren == nil {
		t.Errorf("children for %c was nil, expected %d\n", value, expected)
	}
	childrenLen := len(valueChildren.children)
	if childrenLen != expected {
		t.Errorf("children for %c had %d children, expected %d\n", value, childrenLen, expected)
	}
}

func assertNumberPredictions(t *testing.T, values []string, expected int) {
	actual := len(values)

	if actual != expected {
		t.Errorf("number predictions was %d, expected %d\n", actual, expected)
	}
}

func TestAddAutocomplete(t *testing.T) {
	builder := New()

	builder.AdddWord("a")
	assertRootChildren(builder, t, 1)
	builder.AdddWord("cats").AdddWord("cattle")
	assertRootChildren(builder, t, 2)
	assertCharacterChildren(builder, t, 'c', 1)
	builder.AdddWord("brother").AdddWord("brothers")
	assertRootChildren(builder, t, 3)
	assertCharacterChildren(builder, t, 'b', 1)
	builder.AdddWord("crane")
	assertCharacterChildren(builder, t, 'c', 2)
}

func TestPredictAutocomplete(t *testing.T) {
	autoComplete := New()

	autoComplete.AdddWord("cats").AdddWord("cattle")
	assertCharacterChildren(autoComplete, t, 'c', 1)

	suggestedWords := autoComplete.Predict("ca")
	assertNumberPredictions(t, suggestedWords, 2)

	autoComplete.AdddWord("apple").AdddWord("app").AdddWord("application")
	suggestedWords = autoComplete.Predict("app")
	assertNumberPredictions(t, suggestedWords, 3)
}
