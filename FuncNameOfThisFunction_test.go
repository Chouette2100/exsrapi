// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi

import (
	"fmt"
	"testing"
)

func TestFuncNameOfThisFunction(t *testing.T) {
	tests := []struct {
		name         string
		level int
		wantFuncname string
	}{
		// TODO: Add test cases.
		{"test1", 1, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("FuncNameOfThisFunction() = %v, want %v\n", FuncNameOfThisFunction(tt.level), tt.wantFuncname)
			if gotFuncname := FuncNameOfThisFunction(tt.level); gotFuncname != tt.wantFuncname {
				t.Errorf("FuncNameOfThisFunction() = %v, want %v", gotFuncname, tt.wantFuncname)
			}
		})
	}
}
