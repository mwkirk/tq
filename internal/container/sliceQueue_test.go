package container

import "testing"

func TestSliceQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		expected int64
	}{
		{"Enqueue 1 item", []int{1}, 1},
		{"Enqueue 3 items", []int{1, 2, 3}, 3},
		{"Enqueue 5 items", []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			length, err := queue.Length()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if length != 0 {
				t.Errorf("Expected length to be 0, got %d", length)
			}

			for _, v := range tt.enqueue {
				err := queue.Enqueue(v)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			length, err = queue.Length()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if length != tt.expected {
				t.Errorf("Expected length to be %d, got %d", tt.expected, length)
			}
		})
	}
}

func TestSliceQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		dequeue  []int
		expected []int
	}{
		{"Dequeue 1 item", []int{1}, []int{1}, []int{}},
		{"Dequeue 3 items", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"Dequeue 5 items", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			for _, item := range tt.enqueue {
				err := queue.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			for _, expected := range tt.dequeue {
				item, err := queue.Dequeue()
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if item != expected {
					t.Errorf("Expected item to be %d, got %d", expected, item)
				}
			}
		})
	}
}

func TestSliceQueue_Front(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		expected int
	}{
		{"Front 1 item", []int{1}, 1},
		{"Front 3 items", []int{1, 2, 3}, 1},
		{"Front 5 items", []int{1, 2, 3, 4, 5}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			for _, item := range tt.enqueue {
				err := queue.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			item, err := queue.Front()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if item != tt.expected {
				t.Errorf("Expected front item to be %d, got %d", tt.expected, item)
			}
		})
	}
}

func TestSliceQueue_Back(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		expected int
	}{
		{"Back 1 item", []int{1}, 1},
		{"Back 3 items", []int{1, 2, 3}, 3},
		{"Back 5 items", []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			for _, item := range tt.enqueue {
				err := queue.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			item, err := queue.Back()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if item != tt.expected {
				t.Errorf("Expected back item to be %d, got %d", tt.expected, item)
			}
		})
	}
}

func TestSliceQueue_Length(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		expected int64
	}{
		{"Length 1 item", []int{1}, 1},
		{"Length 3 items", []int{1, 2, 3}, 3},
		{"Length 5 items", []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			for _, item := range tt.enqueue {
				err := queue.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			length, err := queue.Length()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if length != tt.expected {
				t.Errorf("Expected length to be %d, got %d", tt.expected, length)
			}
		})
	}
}

func TestSliceQueue_Filter(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		filter   func(int) bool
		expected []int
	}{
		{"Filter odd numbers", []int{1, 2, 3, 4, 5}, func(v int) bool {
			return v%2 == 1
		}, []int{1, 3, 5}},
		{"Filter even numbers", []int{1, 2, 3, 4, 5}, func(v int) bool {
			return v%2 == 0
		}, []int{2, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			for _, item := range tt.enqueue {
				err := queue.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			filtered := queue.Filter(tt.filter)
			for i, item := range filtered {
				if item != tt.expected[i] {
					t.Errorf("Expected item to be %d, got %d", tt.expected[i], item)
				}
			}
		})
	}
}

func TestSliceQueue_Equal(t *testing.T) {
	tests := []struct {
		name     string
		enqueue1 []int
		enqueue2 []int
		equal    bool
	}{
		{"Equal 1 item", []int{1}, []int{1}, true},
		{"Equal 3 items", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"Equal 5 items", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, true},
		{"Not equal different length", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, false},
		{"Not equal different values", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 6}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue1 := NewSliceQueue[int]()
			for _, item := range tt.enqueue1 {
				err := queue1.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			queue2 := NewSliceQueue[int]()
			for _, item := range tt.enqueue2 {
				err := queue2.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if queue1.Equal(queue2) != tt.equal {
				t.Errorf("Expected queues to be equal: %v", tt.equal)
			}
		})
	}
}

func TestSliceQueue_FindFirst(t *testing.T) {
	tests := []struct {
		name     string
		enqueue  []int
		find     func(int) bool
		expected int
		found    bool
	}{
		{"FindFirst odd number", []int{1, 2, 3, 4, 5}, func(v int) bool {
			return v%2 == 1
		}, 1, true},
		{"FindFirst even number", []int{1, 2, 3, 4, 5}, func(v int) bool {
			return v%2 == 0
		}, 2, true},
		{"FindFirst missing", []int{1, 2, 3, 4, 5}, func(v int) bool {
			return v == 6
		}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			for _, item := range tt.enqueue {
				err := queue.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			item, found := queue.FindFirst(tt.find)
			if found != tt.found {
				t.Errorf("Expected found to be %t, got %t", tt.found, found)
			}

			if found && item != tt.expected {
				t.Errorf("Expected item to be %d, got %d", tt.expected, item)
			}
		})
	}
}

func TestSliceQueue_FindFirstAndDelete(t *testing.T) {
	tests := []struct {
		name          string
		enqueue       []int
		pred          func(int) bool
		findExpected  int
		queueExpected []int
	}{
		{name: "FindFirstAndDelete first value 3", enqueue: []int{1, 2, 3, 4, 5}, pred: func(v int) bool {
			return v == 3
		}, findExpected: 3, queueExpected: []int{1, 2, 4, 5}},
		{name: "FindFirstAndDelete first odd value", enqueue: []int{1, 2, 3, 4, 5}, pred: func(v int) bool {
			return v%2 == 1
		}, findExpected: 1, queueExpected: []int{2, 3, 4, 5}},
		{name: "FindFirstAndDelete first even value", enqueue: []int{1, 2, 3, 4, 5}, pred: func(v int) bool {
			return v%2 == 0
		}, findExpected: 2, queueExpected: []int{1, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewSliceQueue[int]()
			for _, item := range tt.enqueue {
				err := queue.Enqueue(item)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			v, ok := queue.FindFirstAndDelete(tt.pred)
			if !ok {
				t.Errorf("Expected to find item")
			}

			if v != tt.findExpected {
				t.Errorf("Expected to find %d, got %d", tt.findExpected, v)
			}

			if !sliceEqual(queue.s, tt.queueExpected) {
				t.Errorf("Expected queue to be %v, got %v", tt.queueExpected, queue.s)
			}
		})
	}
}
