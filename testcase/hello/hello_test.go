package hello

import "testing"

func TestSayHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"java"}, "Hello java"},
		{"case2", args{"golang"}, "Hello golang"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SayHello(tt.args.name); got != tt.want {
				t.Errorf("SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
