package exsrapi

import "testing"

func TestMakePeriod(t *testing.T) {
	type args struct {
		started_at int64
		ended_at   int64
	}
	tests := []struct {
		name       string
		args       args
		wantPeriod string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{1690448400, 1691672399},
			wantPeriod: "Jul 27, 2023 6:00 PM - Aug 10, 2023 9:59 PM",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPeriod, err := MakePeriod(tt.args.started_at, tt.args.ended_at)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakePeriod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPeriod != tt.wantPeriod {
				t.Errorf("MakePeriod() = %v, want %v", gotPeriod, tt.wantPeriod)
			}
		})
	}
}
