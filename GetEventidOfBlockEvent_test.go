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
)

func TestGetEventidOfBlockEvent(t *testing.T) {
	type args struct {
		eventid string
	}
	tests := []struct {
		name          string
		args          args
		wantBlocklist []Block
		wantErr       bool
	}{
		{
			name: "test1",
			args: args{eventid: "circle2023_2nd_b"},
			wantBlocklist: []Block{
				{
					Label:    "A枠",
					Block_id: 8101,
				},
				{
					Label:    "B枠",
					Block_id: 8102,
				},
				{
					Label:    "C枠",
					Block_id: 8103,
				},
				{
					Label:    "D枠",
					Block_id: 8104,
				},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBlocklist, err := GetEventidOfBlockEvent(tt.args.eventid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventidOfBlockEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBlocklist, tt.wantBlocklist) {
				t.Errorf("GetEventidOfBlockEvent() = %v, want %v", gotBlocklist, tt.wantBlocklist)
			}
		})
	}
}
