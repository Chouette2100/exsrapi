package exsrapi

import (
	"fmt"
	"testing"
)

func TestFuncNameOfThisFunction(t *testing.T) {
	tests := []struct {
		name         string
		wantFuncname string
	}{
		// TODO: Add test cases.
		{"test1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("FuncNameOfThisFunction() = %v, want %v\n", FuncNameOfThisFunction(), tt.wantFuncname)
			if gotFuncname := FuncNameOfThisFunction(); gotFuncname != tt.wantFuncname {
				t.Errorf("FuncNameOfThisFunction() = %v, want %v", gotFuncname, tt.wantFuncname)
			}
		})
	}
}
