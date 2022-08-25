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
				assert.Equal(t, tt.want, got.value, "GetHead")
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
				assert.Equal(t, tt.want, got.value, "GetTail()")
			}
		})
	}
}

func TestLinkedList_ListToSlice(t *testing.T) {
	tests := []struct {
		name       string
		want       []string
		valueToAdd []string
	}{
		{
			name:       "positive case - 1",
			want:       []string{"test_data", "another_test_data"},
			valueToAdd: []string{"test_data", "another_test_data"},
		},
		{
			name:       "positive case - 2 - empty list",
			want:       []string{},
			valueToAdd: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for _, v := range tt.valueToAdd {
				l.AddItemToBack(v)
			}
			assert.Equalf(t, tt.want, l.ListToSlice(), "ListToSlice()")
		})
	}
}

func TestLinkedList_AddItemToFront_Logic(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "positive case - 1",
			args: []string{"test_data"},
			want: "test_data",
		},
		{
			name: "positive case - 2 - couple objects",
			args: []string{"test_data", "another_test_data"},
			want: "another_test_data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for _, v := range tt.args {
				l.AddItemToFront(v)
			}
			assert.Equalf(t, tt.want, l.GetHead().value, "AddItemToFront")
			assert.Equal(t, len(tt.args), l.len)
		})
	}
}

func TestLinkedList_AddItemToFront_Value(t *testing.T) {
	tests := []struct {
		name            string
		arg             string
		want            string
		needPretestData bool
		nextElemNotNil  bool
	}{
		{
			name: "positive case - 1",
			arg:  "test_data",
			want: "test_data",
		},
		{
			name:            "positive case - 2",
			arg:             "test_data",
			want:            "test_data",
			nextElemNotNil:  true,
			needPretestData: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			if tt.needPretestData {
				l.AddItemToBack(tt.arg)
			}
			elem := l.AddItemToFront(tt.arg)
			assert.Equal(t, tt.want, elem.value)
			assert.Nil(t, elem.prev)
			if tt.nextElemNotNil {
				assert.NotEmpty(t, elem.next)
			}
		})
	}
}

func TestLinkedList_AddItemToBack_Logic(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "positive case - 1",
			args: []string{"test_data"},
			want: "test_data",
		},
		{
			name: "positive case - 2 - couple objects",
			args: []string{"test_data", "another_test_data"},
			want: "another_test_data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for _, v := range tt.args {
				l.AddItemToBack(v)
			}
			assert.Equalf(t, tt.want, l.GetTail().value, "AddItemToFront")
			assert.Equal(t, len(tt.args), l.len)
		})
	}
}

func TestLinkedList_AddItemToBack_Value(t *testing.T) {
	tests := []struct {
		name            string
		arg             string
		want            string
		needPretestData bool
		prevElemNotNil  bool
	}{
		{
			name: "positive case - 1",
			arg:  "test_data",
			want: "test_data",
		},
		{
			name:            "positive case - 2",
			arg:             "test_data",
			want:            "test_data",
			prevElemNotNil:  true,
			needPretestData: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			if tt.needPretestData {
				l.AddItemToFront(tt.arg)
			}
			elem := l.AddItemToBack(tt.arg)
			assert.Equal(t, tt.want, elem.value)
			assert.Nil(t, elem.next)
			if tt.prevElemNotNil {
				assert.NotEmpty(t, elem.prev)
			}
		})
	}
}

func TestLinkedList_RemoveItem(t *testing.T) {
	tests := []struct {
		name        string
		expectedLen int
		dataToAdd   []string
	}{
		{
			name:        "positive case - 1",
			expectedLen: 0,
			dataToAdd:   []string{},
		},
		{
			name:        "positive case - 2",
			expectedLen: 1,
			dataToAdd:   []string{"some_data"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for _, v := range tt.dataToAdd {
				l.AddItemToFront(v)
			}
			elem := l.AddItemToBack("test_data")
			l.RemoveItem(elem)
			assert.Equal(t, tt.expectedLen, l.len)
		})
	}
}

func TestLinkedList_RemoveFrontItem(t *testing.T) {
	tests := []struct {
		name          string
		itemsToAdd    []string
		wantErr       bool
		finallyLen    int
		itemValueLeft string
	}{
		{
			name:       "positive case - 1",
			itemsToAdd: []string{"test_data"},
			wantErr:    false,
			finallyLen: 0,
		},
		{
			name:          "positive case - 1",
			itemsToAdd:    []string{"test_data", "another_test_data"},
			wantErr:       false,
			finallyLen:    1,
			itemValueLeft: "test_data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := InitLinkedList[string]()
			for _, v := range tt.itemsToAdd {
				l.AddItemToFront(v)
			}
			err := l.RemoveFrontItem()
			item := l.GetHead()
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.finallyLen, l.len)
			if tt.itemValueLeft != "" {
				assert.Equal(t, tt.itemValueLeft, item.value)
			}
		})
	}
}
