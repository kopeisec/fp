package fp

import (
	"errors"
	"reflect"
	"testing"
)

func TestAccumulateCombineErrors(t *testing.T) {
	type args struct {
		err1 error
		err2 error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				err1: nil,
				err2: nil,
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				err1: nil,
				err2: errors.New("1"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AccumulateCombineErrors(tt.args.err1, tt.args.err2); (err != nil) != tt.wantErr {
				t.Errorf("AccumulateCombineErrors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccumulateStringJoin(t *testing.T) {
	type args struct {
		sep string
		s1  string
		s2  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				sep: "-",
				s1:  "1",
				s2:  "2",
			},
			want: "1-2",
		},
		{
			name: "",
			args: args{
				sep: "-",
				s1:  "1",
				s2:  "",
			},
			want: "1",
		},
		{
			name: "",
			args: args{
				sep: "-",
				s1:  "",
				s2:  "2",
			},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccumulateStringJoin(tt.args.sep)(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccumulateStringJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	type args struct {
		b1 bool
		b2 bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				b1: false,
				b2: false,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				b1: true,
				b2: false,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				b1: false,
				b2: true,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				b1: true,
				b2: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Or(tt.args.b1, tt.args.b2); got != tt.want {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	type args struct {
		b1 bool
		b2 bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				b1: false,
				b2: false,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				b1: false,
				b2: true,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				b1: true,
				b2: false,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				b1: true,
				b2: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := And(tt.args.b1, tt.args.b2); got != tt.want {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}
