package stringreverse

import "testing"

func TestRevstr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			s:    "ABCD",
			want: "DCBA",
		},
		{
			name: "Test1",
			s:    "АБЦ",
			want: "ЦБА",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Revstr(tt.s); got != tt.want {
				t.Errorf("Revstr() = %v, want %v", got, tt.want)
			}
		})
	}
}
