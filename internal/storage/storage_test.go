package storage

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"

	if pathKey.PathName != expectedPathName {
		t.Errorf("expected: [%s], but got: [%s]\n", expectedPathName, pathKey.PathName)
	}
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("expected: [%s], but got: [%s]\n", expectedOriginalKey, pathKey.Filename)
	}
}

func TestStore(t *testing.T) {
	s := newStore()

	key := "momsspecial"

	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewBuffer(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	if !reflect.DeepEqual(b, data) {
		t.Errorf("want %s, got %s\n", data, b)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	s := newStore()
	key := "momsspecial"
	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}
