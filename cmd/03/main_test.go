package main

import (
	"io"
	"strings"
	"testing"
)

func Test_task1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"task1-example",
			args{strings.NewReader(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)},
			157,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := task1(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("task1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("task1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task2(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"task2-example",
			args{strings.NewReader(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)},
			70,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := task2(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("task2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("task2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMatching(t *testing.T) {
	type args struct {
		contents string
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr bool
	}{
		{"1", args{"vJrwpWtwJgWrhcsFMMfFFhFp"}, 'p', false},
		{"2", args{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}, 'L', false},
		{"3", args{"PmmdzqPrVvPwwTWBwg"}, 'P', false},
		{"4", args{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"}, 'v', false},
		{"5", args{"ttgJtRGJQctTZtZT"}, 't', false},
		{"6", args{"CrZsJsPPZsGzwwsLwLmpwMDw"}, 's', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findMatching(tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("findMatching() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_matching(t *testing.T) {
	type args struct {
		sets []string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"1", args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}}, 'r'},
		{"2", args{[]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}}, 'Z'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sets []map[rune]int
			for _, s := range tt.args.sets {
				sets = append(sets, set(s))
			}
			got := matching(sets...)
			if got != tt.want {
				t.Errorf("findMatching() = %c, want %c", got, tt.want)
			}
		})
	}
}

func Test_score(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"p", args{'p'}, 16},
		{"L", args{'L'}, 38},
		{"P", args{'P'}, 42},
		{"v", args{'v'}, 22},
		{"t", args{'t'}, 20},
		{"s", args{'s'}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.args.r); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}
