package tests

import (
	"TAREA1/estructuras"
	"testing"
)

func TestDictionaryPutGet(t *testing.T) {

	dict := estructuras.NewDictionary()
	dict.Put("age", 21)

	value, ok := dict.Get("age")
	if !ok {
		t.Error("Expected key 'age' to exist")
	}

	if value != 21 {
		t.Errorf("Expected 22 but got %d", value)
	}

}

func TestDictionaryUpdateValue(t *testing.T) {

	dict := estructuras.NewDictionary()

	dict.Put("year", 2025)
	dict.Put("year", 2026)

	value, _ := dict.Get("year")

	if value != 2026 {
		t.Errorf("Expected updated value 2025 but got %d", value)
	}

}

func TestDictionaryRemove(t *testing.T) {

	dict := estructuras.NewDictionary()

	dict.Put("a", 1)
	dict.Put("b", 2)

	dict.Remove("a")

	if dict.Contains("a") {
		t.Error("Expected key 'a' to be removed")
	}

}

func TestDictionaryContains(t *testing.T) {

	dict := estructuras.NewDictionary()

	dict.Put("name", 1)

	if !dict.Contains("name") {
		t.Error("Expected dictionary to contain key 'name'")
	}

	if dict.Contains("age") {
		t.Error("Expected dictionary to not contain key 'age'")
	}

}

func TestDictionarySize(t *testing.T) {

	dict := estructuras.NewDictionary()

	dict.Put("a", 1)
	dict.Put("b", 2)
	dict.Put("c", 3)

	if dict.Size() != 3 {
		t.Errorf("Expected size 3 but got %d", dict.Size())
	}

}
