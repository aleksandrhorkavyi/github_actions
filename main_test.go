package main

import "testing"

func Test_testPrint(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{want: "TEST..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testPrint(); got != tt.want {
				t.Errorf("testPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
