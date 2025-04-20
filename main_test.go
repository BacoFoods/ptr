package ptr_test

import (
	"testing"

	. "github.com/BacoFoods/ptr"
)

func TestPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		val := 42
		ptr := Ptr(val)
		if *ptr != val {
			t.Errorf("expected %d, got %d", val, *ptr)
		}
	})

	t.Run("string", func(t *testing.T) {
		val := "test"
		ptr := Ptr(val)
		if *ptr != val {
			t.Errorf("expected %s, got %s", val, *ptr)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type example struct {
			Name string
		}
		val := example{Name: "example"}
		ptr := Ptr(val)
		if ptr.Name != val.Name {
			t.Errorf("expected %s, got %s", val.Name, ptr.Name)
		}
	})
}

func TestVal(t *testing.T) {
	t.Run("nil pointer with default value", func(t *testing.T) {
		var ptr *int
		defaultVal := 10
		val := Val(ptr, defaultVal)
		if val != defaultVal {
			t.Errorf("expected %d, got %d", defaultVal, val)
		}
	})

	t.Run("non-nil pointer without default value", func(t *testing.T) {
		val := 42
		ptr := Ptr(val)
		defaultVal := 10
		result := Val(ptr, defaultVal)
		if result != val {
			t.Errorf("expected %d, got %d", val, result)
		}
	})
}

func TestEqualVals(t *testing.T) {
	t.Run("both nil pointers", func(t *testing.T) {
		var ptr1, ptr2 *int
		if EqualVals(ptr1, ptr2) {
			t.Errorf("nil comparison should return false")
		}
	})

	t.Run("one nil pointer", func(t *testing.T) {
		val := 42
		ptr := Ptr(val)
		if EqualVals(ptr, nil) {
			t.Errorf("expected false, got true")
		}
		if EqualVals(nil, ptr) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("equal values", func(t *testing.T) {
		val1, val2 := 42, 42
		ptr1, ptr2 := Ptr(val1), Ptr(val2)
		if !EqualVals(ptr1, ptr2) {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("different values", func(t *testing.T) {
		val1, val2 := 42, 43
		ptr1, ptr2 := Ptr(val1), Ptr(val2)
		if EqualVals(ptr1, ptr2) {
			t.Errorf("expected false, got true")
		}
	})
}

func TestNilIfZero(t *testing.T) {
	t.Run("non-zero value", func(t *testing.T) {
		val := 42
		ptr := NilIfZero(val)
		if ptr == nil || *ptr != val {
			t.Errorf("expected pointer to %d, got nil or incorrect value", val)
		}
	})

	t.Run("zero value", func(t *testing.T) {
		val := 0
		ptr := NilIfZero(val)
		if ptr != nil {
			t.Errorf("expected nil, got pointer to %d", *ptr)
		}
	})

	t.Run("non-zero string", func(t *testing.T) {
		val := "test"
		ptr := NilIfZero(val)
		if ptr == nil || *ptr != val {
			t.Errorf("expected pointer to %s, got nil or incorrect value", val)
		}
	})

	t.Run("zero string", func(t *testing.T) {
		val := ""
		ptr := NilIfZero(val)
		if ptr != nil {
			t.Errorf("expected nil, got pointer to %s", *ptr)
		}
	})
}
