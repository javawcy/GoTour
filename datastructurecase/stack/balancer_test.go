package stack

import "testing"

func TestIsBalance(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"balancer1", args{"(hello)"}, true},
		{"balancer2", args{"((hello)"}, false},
		{"balancer3", args{"([{<hello>}])"}, true},
		{"balancer4", args{"([{<hello}>)"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalance(tt.args.str); got != tt.want {
				t.Errorf("IsBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
