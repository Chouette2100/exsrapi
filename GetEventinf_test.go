// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi

import "testing"

func TestGetEventinf(t *testing.T) {
	type args struct {
		eventid   string
		eventinfo *Event_Inf
	}

	var eventinfo Event_Inf

	tests := []struct {
		name       string
		args       args
		wanterr error
	}{
		{
		name: "test",
		args: args{
			eventid: "bestofhawaiianwedding2023_3d",
			eventinfo: &eventinfo,
		},
		wanterr: nil,		//	実行経過を確認するためPASSさせない
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if goterr := GetEventinf(tt.args.eventid, tt.args.eventinfo); goterr != tt.wanterr {
				t.Errorf(" eventinf = %v, GetEventinf() = %v, want %v", eventinfo, goterr, tt.wanterr)
			}
		})
	}
}
