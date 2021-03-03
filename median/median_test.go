package median

import (
	"testing"
)

func Test_medianFunc(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{values: []float64{3.0 , 4.0, 5.0}}, want: 4.0},
		{name: "1", args: args{values: []float64{3.0 , 3.5, 4.0, 5.0}}, want: 4.0},
		{name: "1", args: args{values: []float64{3.0 , 3.5, 4.0, 5.0}}, want: 4.0},
		{name: "1", args: args{values: []float64{1.0 , 1.5, 1.0, 5.0}}, want: 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianFunc(tt.args.values); got != tt.want {
				t.Errorf("median() = %v, want %v", got, tt.want)
			}
		})
	}
}
