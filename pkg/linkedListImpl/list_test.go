package linkedListImpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_GetListLength(t *testing.T) {
	tests := []struct {
		name      string
		want      int
		dataToAdd []string
	}{
		{
			name:      "positive case - 1",
			dataToAdd: []string{"some_data", "another_one_test_data"},
			want:      2,
		},
		{
			name:      "positive case - 2",
			dataToAdd: []string{},
			want:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testList := InitLinkedList[string]()
			for _, v := range tt.dataToAdd {
				testList.AddItemToBack(v)
			}
			if got := testList.GetListLength(); got != tt.want {
				t.Errorf("GetListLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_GetHead(t *testing.T) {
	tests := []struct {
		name       string
		want       string
		valueToAdd []string
		nilHead    bool
	}{
		{
			name:       "positive case - 1",
			want:       "test_value",
			valueToAdd: []string{"test_value"},
		},
		{
			name:    "positive case - 2 - empty list",
			want:    "",
			nilHead: true,
		},
		{
			name:       "positive case - 3 - couple elements in list",
			want:       "another_test_value",
			valueToAdd: []string{"test_value", "another_test_value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for _, v := range tt.valueToAdd {
				l.AddItemToFront(v)
			}
			got := l.GetHead()
			if tt.nilHead {
				assert.Nil(t, got)
			} else {
				assert.Equal(t, tt.want, got.Value, "GetHead")
			}
		})
	}
}

func TestLinkedList_GetTail(t *testing.T) {
	tests := []struct {
		name       string
		want       string
		valueToAdd []string
		nilTail    bool
	}{
		{
			name:       "positive case - 1",
			want:       "test_value",
			valueToAdd: []string{"test_value"},
		},
		{
			name:    "positive case - 2 - empty list",
			want:    "",
			nilTail: true,
		},
		{
			name:       "positive case - 3 - couple elements",
			want:       "another_test_value",
			valueToAdd: []string{"test_value", "another_test_value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for _, v := range tt.valueToAdd {
				l.AddItemToBack(v)
			}
			got := l.GetTail()
			if tt.nilTail {
				assert.Nil(t, got)
			} else {
				assert.Equal(t, tt.want, got.Value, "GetTail()")
			}
		})
	}
}
