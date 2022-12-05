package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const sample = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Test_parseStacks(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		want    []stack
		wantErr bool
	}{
		{
			"example",
			args{bufio.NewScanner(strings.NewReader(sample))},
			[]stack{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStacks(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStacks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("parseStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task1(t *testing.T) {
	crates, err := task(strings.NewReader(sample), false)
	if err != nil {
		t.Fatalf("task1() error = %v, want <nil>", err)
	}
	if !cmp.Equal(crates, []crate{'C', 'M', 'Z'}) {
		t.Errorf("task1() = %v, want CMZ", crates)
	}
}

func Test_task2(t *testing.T) {
	crates, err := task(strings.NewReader(sample), true)
	if err != nil {
		t.Fatalf("task2() error = %v, want <nil>", err)
	}
	if !cmp.Equal(crates, []crate{'M', 'C', 'D'}) {
		t.Errorf("task2() = %v, want MCD", crates)
	}
}
