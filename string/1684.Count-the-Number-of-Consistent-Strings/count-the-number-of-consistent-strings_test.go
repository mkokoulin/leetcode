package main

import "testing"

func Test_countConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{
				allowed: "ab",
				words:   []string{"ad", "bd", "aaab", "baa", "badab"},
			},
			want: 2,
		},
		{
			name: "Test #2",
			args: args{
				allowed: "abc",
				words:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsistentStrings(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
