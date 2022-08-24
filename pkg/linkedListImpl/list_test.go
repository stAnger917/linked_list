package linkedListImpl

import "testing"

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
			name:            "positive case - 1",
			needPreTestFunc: false,
			want:            0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testList := InitLinkedList()
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
