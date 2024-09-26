package fp

import "testing"

func TestConditionHasError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				err: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConditionHasError(tt.args.err); got != tt.want {
				t.Errorf("ConditionHasError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConditionShouldNotEmpty(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConditionShouldNotEmpty(tt.args.s); got != tt.want {
				t.Errorf("ConditionShouldNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
