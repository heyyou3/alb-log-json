package alb_log

import "testing"

func TestOutputALBLogStdoutAdapter_Save(t *testing.T) {
	type args struct {
		albLog []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Wrote stdout", args{albLog: []string{"wrote stdout test"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OutputALBLogStdoutAdapter{}
			if got := o.Save(tt.args.albLog); got != tt.want {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
