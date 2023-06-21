/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"reflect"
	"testing"

	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

func TestGetEventidOfEventBox(t *testing.T) {
	type args struct {
		eventid string
	}
	tests := []struct {
		name         string
		args         args
		wantNamelist []string
		wantErr      bool
	}{
		{
			name: "test1",
			args: args{
				eventid: "bestofhawaiianwedding2023_hk3",
			},
			wantNamelist: []string{
				"bestofhawaiianwedding2023_3",
				"bestofhawaiianwedding2023_3b",
				"bestofhawaiianwedding2023_3c",
				"bestofhawaiianwedding2023_3d",
				"bestofhawaiianwedding2023_3_2a",
				"bestofhawaiianwedding2023_3_2b",
				"bestofhawaiianwedding2023_3fin"},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNamelist, err := GetEventidOfEventBox(tt.args.eventid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventidOfEventBox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNamelist, tt.wantNamelist) {
				t.Errorf("GetEventidOfEventBox() = %v, want %v", gotNamelist, tt.wantNamelist)
			}
		})
	}
}
