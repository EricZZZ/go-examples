package main

import (
	"strings"
	"testing"
)

type FakeStorage struct{}

func (f *FakeStorage) Upload(path string, data []byte) error {
	return nil
}

func TestSavePhoto(t *testing.T) {
	service := &PhotoService{storage: &FakeStorage{}}

	path, err := service.SavePhoto("user-123", []byte("photo data"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(path, "user-123") {
		t.Errorf("expected path to contain user ID, got %s", path)
	}
}
