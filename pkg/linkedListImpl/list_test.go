package linkedListImpl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_GetListLength(t *testing.T) {
	tests := []struct {
		name            string
		want            int
		needPreTestFunc bool
	}{
		{
			name:            "positive case - 1",
			needPreTestFunc: true,
			want:            2,
		},
		{
			name:            "positive case - 2",
			needPreTestFunc: false,
			want:            0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testList := InitLinkedList[string]()
			if tt.needPreTestFunc {
				testList.AddItemToFront("SomeTestData")
				testList.AddItemToFront("AnotherTestData")
			}
			if got := testList.GetListLength(); got != tt.want {
				t.Errorf("GetListLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_GetHead(t *testing.T) {
	tests := []struct {
		name               string
		howManyElementsAdd int
		want               *Item[any]
		valueToAdd         []string
		wantErr            bool
		error              error
	}{
		{
			name: "positive case - 1",
			want: &Item[any]{
				Value: "test_value",
				Next:  nil,
				Prev:  nil,
			},
			howManyElementsAdd: 1,
			valueToAdd:         []string{"test_value"},
		},
		{
			name:               "positive case - 2 - empty list",
			want:               nil,
			howManyElementsAdd: 0,
			wantErr:            true,
			error:              ErrNilHead,
		},
		{
			name: "positive case - 3 - couple elements in list",
			want: &Item[any]{
				Value: "another_test_value",
			},
			howManyElementsAdd: 2,
			valueToAdd:         []string{"test_value", "another_test_value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for i := 0; i < tt.howManyElementsAdd; i++ {
				l.AddItemToFront(fmt.Sprintf("%v", tt.valueToAdd[i]))
			}
			got, err := l.GetHead()
			if tt.wantErr {
				assert.Equal(t, tt.error, err, "GetHead")
			} else {
				assert.Equal(t, tt.want.Value, got.Value, "GetHead")
			}
		})
	}
}

func TestLinkedList_GetTail(t *testing.T) {
	tests := []struct {
		name               string
		want               *Item[any]
		wantErr            bool
		error              error
		howManyElementsAdd int
		valueToAdd         []string
	}{
		{
			name: "positive case - 1",
			want: &Item[any]{
				Value: "test_value",
				Next:  nil,
				Prev:  nil,
			},
			howManyElementsAdd: 1,
			valueToAdd:         []string{"test_value"},
		},
		{
			name:               "positive case - 2 - empty list",
			want:               nil,
			howManyElementsAdd: 0,
			wantErr:            true,
			error:              ErrNilTail,
		},
		{
			name: "positive case - 3 - couple elements",
			want: &Item[any]{
				Value: "another_test_value",
			},
			howManyElementsAdd: 2,
			valueToAdd:         []string{"test_value", "another_test_value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for i := 0; i < tt.howManyElementsAdd; i++ {
				l.AddItemToBack(fmt.Sprintf("%v", tt.valueToAdd[i]))
			}
			listSlice := l.ListToSlice()
			got, err := l.GetTail()
			if tt.wantErr {
				assert.Equal(t, tt.error, err, "GetTail")
			} else {
				assert.Equal(t, tt.want.Value, got.Value, "GetTail()")
				if len(listSlice) > 0 {
					assert.Equalf(t, listSlice[len(listSlice)-1], got, "GetTail()")
				}
			}
		})
	}
}
