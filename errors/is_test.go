package errors

import (
	"errors"
	"testing"
)

type MyError1 struct {
}

// Error implements error.
func (*MyError1) Error() string {
	return "myerror1"
}

func NewMyError1() error {
	return &MyError1{}
}

type MyError2 struct {
}

// Error implements error.
func (*MyError2) Error() string {
	return "myerror2"
}

func NewMyError2() error {
	return &MyError2{}
}

func TestIs(t *testing.T) {
	t.Parallel()
	error1 := errors.New("error1")
	error2 := errors.New("error2")
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true/compare the errors generated by errors.New",
			args: args{
				err:    error1,
				target: error1,
			},
			want: true,
		},
		{
			name: "false/compare the errors generated by errors.New",
			args: args{
				err:    error1,
				target: error2,
			},
			want: false,
		},
		{
			name: "true/compare the errors generated by NewMyError",
			args: args{
				err:    NewMyError1(),
				target: NewMyError1(),
			},
			want: true,
		},
		{
			name: "false/compare the errors generated by NewMyError",
			args: args{
				err:    NewMyError1(),
				target: NewMyError2(),
			},
			want: false,
		},
		{
			name: "false/compare the errors generated by NewMyError",
			args: args{
				err:    NewMyError1(),
				target: errors.New("test"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}