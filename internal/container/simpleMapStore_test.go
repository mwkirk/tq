package container

import (
	"errors"
	"testing"
)

func TestSimpleMapStore_Exists(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	if exists, _ := s.Exists("key"); exists {
		t.Error("Expected key to not exist")
	}

	s.Put("key", 1)
	exists, err := s.Exists("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !exists {
		t.Error("Expected key to exist")
	}
}

func TestSimpleMapStore_Get(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	if _, err := s.Get("key"); !errors.Is(err, ErrorNotFound) {
		t.Errorf("Expected ErrorNotFound, got %v", err)
	}

	s.Put("key", 1)
	v, err := s.Get("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if v != 1 {
		t.Errorf("Expected 1, got %v", v)
	}
}

func TestSimpleMapStore_GetAndDelete(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	if _, err := s.GetAndDelete("key"); !errors.Is(err, ErrorNotFound) {
		t.Errorf("Expected ErrorNotFound, got %v", err)
	}

	s.Put("key", 1)
	v, err := s.GetAndDelete("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if v != 1 {
		t.Errorf("Expected 1, got %v", v)
	}
	if exists, _ := s.Exists("key"); exists {
		t.Error("Expected key to not exist")
	}
}

func TestSimpleMapStore_Put(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	s.Put("key", 1)
	if exists, _ := s.Exists("key"); !exists {
		t.Error("Expected key to exist")
	}
}

func TestSimpleMapStore_Delete(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	s.Put("key", 1)
	if exists, _ := s.Exists("key"); !exists {
		t.Error("Expected key to exist")
	}

	s.Delete("key")
	if exists, _ := s.Exists("key"); exists {
		t.Error("Expected key to not exist")
	}
}

func TestSimpleMapStore_Update(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	s.Put("key", 1)
	if exists, _ := s.Exists("key"); !exists {
		t.Error("Expected key to exist")
	}

	err := s.Update("key", func(v int) int {
		return v + 1
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	v, err := s.Get("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if v != 2 {
		t.Errorf("Expected 2, got %v", v)
	}
}

func TestSimpleMapStore_Update_KeyNotFound(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	if err := s.Update("key", func(v int) int {
		return v + 1
	}); !errors.Is(err, ErrorNotFound) {
		t.Errorf("Expected ErrorNotFound, got %v", err)
	}
}

func TestSimpleMapStore_Update_Panic(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	s.Put("key", 1)
	if exists, _ := s.Exists("key"); !exists {
		t.Error("Expected key to exist")
	}

	err := s.Update("key", func(v int) int {
		panic("test")
	})
	if !errors.Is(err, ErrorOperationPanicked) {
		t.Error("Expected operation panicked error")
	}
}

func TestSimpleMapStore_Filter(t *testing.T) {
	s := NewSimpleMapStore[string, int]()
	s.Put("key1", 1)
	s.Put("key2", 2)
	s.Put("key3", 3)

	filtered := s.Filter(func(v int) bool {
		return v%2 == 0
	})
	if len(filtered) != 1 {
		t.Errorf("Expected 1, got %v", len(filtered))
	}
	if filtered[0] != 2 {
		t.Errorf("Expected 2, got %v", filtered[0])
	}
}
