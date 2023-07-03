package main

import "testing"

func Test_validateParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive - 1",
			args: args{
				"((({{{[[[]]]}}})))",
			},
			want: true,
		},
		{
			name: "Positive - 2",
			args: args{
				"{{{[][][]}}}",
			},
			want: true,
		},
		{
			name: "Positive - 3",
			args: args{
				"()(){{{}}}[][]",
			},
			want: true,
		},
		{
			name: "Positive - 4",
			args: args{
				"()()()()",
			},
			want: true,
		},
		{
			name: "Negative - 1",
			args: args{
				"(((())})",
			},
			want: false,
		},
		{
			name: "Negative - 2",
			args: args{
				"(((())})",
			},
			want: false,
		},
		{
			name: "Negative - 3",
			args: args{
				"((()()()))(()",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateParentheses(tt.args.s); got != tt.want {
				t.Errorf("validateParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
