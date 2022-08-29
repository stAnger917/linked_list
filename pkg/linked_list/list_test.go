package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testString1 = "first_string"
var testString2 = "second_string"
var testString3 = "third_string"

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
		name         string
		expectedLen  int
		expectedData []string
		listData     func() (*LinkedList[string], *Item[string])
	}{
		{
			name:        "positive case - 1 - one item in list",
			expectedLen: 0,
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				itemToRemove := l.AddItemToBack("some_data")
				return l, itemToRemove
			},
			expectedData: []string{},
		},
		{
			name:        "positive case - 2 - three items in list, deleted from middle",
			expectedLen: 3,
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				l.AddItemToBack(testString2)
				itemToRemove := l.AddItemToBack("some_data")
				l.AddItemToBack(testString3)
				itemToRemove.next = l.tail
				return l, itemToRemove
			},
			expectedData: []string{testString1, testString2, testString3},
		},
		{
			name:        "positive case - 3 - delete first elem",
			expectedLen: 3,
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				itemToRemove := l.AddItemToFront("some_data")
				l.AddItemToBack(testString1)
				itemToRemove.next = l.tail
				l.AddItemToBack(testString2)
				l.AddItemToBack(testString3)
				return l, itemToRemove
			},
			expectedData: []string{testString1, testString2, testString3},
		},
		{
			name:        "positive case - 3 - delete last item",
			expectedLen: 3,
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				l.AddItemToBack(testString2)
				l.AddItemToBack(testString3)
				itemToRemove := l.AddItemToBack("some_data")
				return l, itemToRemove
			},
			expectedData: []string{testString1, testString2, testString3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, item := tt.listData()
			l.RemoveItem(item)
			assert.Equal(t, tt.expectedLen, l.len)
			assert.Equal(t, tt.expectedData, l.ListToSlice())
			assert.Equal(t, item.isDeleted, true)
		})
	}
}

func TestLinkedList_InsertBeforeElem(t *testing.T) {
	type args struct {
		element string
	}
	tests := []struct {
		name         string
		args         args
		expectedData []string
		listData     func() (*LinkedList[string], *Item[string])
	}{
		{
			name: "positive case - 1 - mark == head",
			args: args{
				element: testString1,
			},
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				mark := l.AddItemToFront("some_data")
				l.AddItemToBack(testString1)
				mark.next = l.tail
				l.AddItemToBack(testString2)
				return l, mark
			},
			expectedData: []string{"test_data_to_add", "some_data", "first_string", "second_string"},
		},
		{
			name: "positive case - 2 - mark == tail",
			args: args{
				element: testString1,
			},
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				l.AddItemToBack(testString2)
				mark := l.AddItemToBack("some_data")
				return l, mark
			},
			expectedData: []string{"first_string", "second_string", "test_data_to_add", "some_data"},
		},
		{
			name: "positive case - 3 - mark == tail, 2 items in list",
			args: args{
				element: testString1,
			},
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				mark := l.AddItemToBack("some_data")
				return l, mark
			},
			expectedData: []string{"first_string", "test_data_to_add", "some_data"},
		},
		{
			name: "positive case - 4 - 4 items in list, mark in middle",
			args: args{
				element: testString1,
			},
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				mark := l.AddItemToBack("some_data")
				l.AddItemToBack(testString2)
				mark.next = l.tail
				l.AddItemToBack(testString3)
				return l, mark
			},
			expectedData: []string{"first_string", "test_data_to_add", "some_data", "second_string", "third_string"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, mark := tt.listData()
			_ = l.InsertBeforeElem("test_data_to_add", mark)
			assert.Equal(t, tt.expectedData, l.ListToSlice())
		})
	}
}

func TestLinkedList_InsertAfterElem(t *testing.T) {
	tests := []struct {
		name         string
		expectedData []string
		listData     func() (*LinkedList[string], *Item[string])
	}{
		{
			name: "positive case - 1 - mark == head",
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				mark := l.AddItemToFront("some_data")
				l.AddItemToBack(testString1)
				mark.next = l.tail
				return l, mark
			},
			expectedData: []string{"some_data", "test_data_to_add", "first_string"},
		},
		{
			name: "positive case - 2 - mark == head, three elements",
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				mark := l.AddItemToFront("some_data")
				l.AddItemToBack(testString1)
				mark.next = l.tail
				l.AddItemToBack(testString2)
				return l, mark
			},
			expectedData: []string{"some_data", "test_data_to_add", "first_string", "second_string"},
		},
		{
			name: "positive case - 3 - mark == tail",
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				l.AddItemToBack(testString2)
				mark := l.AddItemToBack("some_data")
				return l, mark
			},
			expectedData: []string{"first_string", "second_string", "some_data", "test_data_to_add"},
		},
		{
			name: "positive case - 4 - mark == tail, 2 items in list",
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				mark := l.AddItemToBack("some_data")
				return l, mark
			},
			expectedData: []string{"first_string", "some_data", "test_data_to_add"},
		},
		{
			name: "positive case - 5 - 4 items in list, mark in middle",
			listData: func() (*LinkedList[string], *Item[string]) {
				l := InitLinkedList[string]()
				l.AddItemToBack(testString1)
				mark := l.AddItemToBack("some_data")
				l.AddItemToBack(testString2)
				mark.next = l.tail
				l.AddItemToBack(testString3)
				return l, mark
			},
			expectedData: []string{"first_string", "some_data", "test_data_to_add", "second_string", "third_string"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, mark := tt.listData()
			l.InsertAfterElem("test_data_to_add", mark)
			assert.Equal(t, tt.expectedData, l.ListToSlice())
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

func TestLinkedList_RemoveBackItem(t *testing.T) {
	tests := []struct {
		name         string
		finallyLen   int
		expectedData []string
		listData     func() *LinkedList[string]
		wantErr      bool
	}{
		{
			name:         "positive case - 1",
			finallyLen:   2,
			expectedData: []string{testString1, testString2},
			listData: func() *LinkedList[string] {
				l := InitLinkedList[string]()
				l.AddItemToFront(testString1)
				l.AddItemToBack(testString2)
				l.AddItemToBack(testString3)
				return l
			},
		},
		{
			name:         "positive case - 2 - single element in list",
			finallyLen:   0,
			expectedData: []string{},
			listData: func() *LinkedList[string] {
				l := InitLinkedList[string]()
				l.AddItemToFront(testString1)
				return l
			},
		},
		{
			name:         "positive case - 3 - empty list",
			finallyLen:   0,
			expectedData: []string{},
			listData: func() *LinkedList[string] {
				l := InitLinkedList[string]()
				return l
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.listData()
			err := l.RemoveBackItem()
			if tt.wantErr {
				assert.Errorf(t, err, ErrNoItem.Error())
			}
			assert.Equal(t, tt.finallyLen, l.len)
			assert.Equal(t, tt.expectedData, l.ListToSlice())
		})
	}
}
