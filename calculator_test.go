package calculator

import "testing"

func TestCalculateInterest(t *testing.T) {
	type args struct {
		principal       float64
		rateOfInterest  float64
		timeInYears     int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Case 1",
			args: args{
				principal:       1000,
				rateOfInterest:  5,
				timeInYears:     2,
			},
			want: 100,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateInterest(tt.args.principal, tt.args.rateOfInterest, tt.args.timeInYears); got != tt.want {
				t.Errorf("CalculateInterest() = %v, want %v", got, tt.want)
			}
		})
	}
}