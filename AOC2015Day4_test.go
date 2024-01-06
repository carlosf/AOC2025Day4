package main

import "testing"

func Test_calcMD5(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcMDFive(tt.args.s); got != tt.want {
				t.Errorf("CalcMDFive() = %v, want %v", got, tt.want)
			}
		})
	}
}
