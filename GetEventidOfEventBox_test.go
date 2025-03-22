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
				eventid: "tsutsuzyuku2",
			},
			wantNamelist: []string{
				"sr_tsutsuzyuku_geinin_2",
				"sr_tsutsuzyuku_2",
				"sr_tsutsuzyuku_talent_2",
				"sr_tsutsuzyuku_assistantmc_2",
			},
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
