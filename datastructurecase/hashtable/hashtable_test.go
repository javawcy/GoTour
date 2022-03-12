package hashtable

import "testing"

func TestFindFirstNoRepeatSubStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"repeatStr1", args{"a green apple"}, "g", false},
		{"repeatStr2", args{"a"}, "a", false},
		{"repeatStr3", args{""}, "", true},
		{"repeatStr4", args{"abab"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFirstNoRepeatSubStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindFirstNoRepeatSubStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindFirstNoRepeatSubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
