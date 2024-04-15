package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {

	store := NewStore(StoreConfig{
		TransformPathFunc: NoTransformPathFunc,
	})

	data := bytes.NewReader([]byte("some data"))
	if err := store.writeStream("myfile", data); err != nil {
		t.Error(err)
	}

}

func TestTransformPathFunc(t *testing.T) {
	key := "myfile"
	path := PathTransformation(key)
	expectedPath := "b3580/ab45c/b088b/a47ff/070aa/81c2d/ae1be/56ca2 b3580ab45cb088ba47ff070aa81c2dae1be56ca2"

	if path.Key != expectedPath {
		t.Errorf("Expected path to be %s, got %s", expectedPath, path)
	}
}
