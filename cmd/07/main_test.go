package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseAndList(t *testing.T) {
	t.Parallel()

	const exampleIn = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	const exampleOut = `- / (dir)
	- a (dir)
		- e (dir)
			- i (file, size=584)
		- f (file, size=29116)
		- g (file, size=2557)
		- h.lst (file, size=62596)
	- b.txt (file, size=14848514)
	- c.dat (file, size=8504156)
	- d (dir)
		- j (file, size=4060174)
		- d.log (file, size=8033020)
		- d.ext (file, size=5626152)
		- k (file, size=7214296)
`

	root, err := parseListing(exampleIn)
	if err != nil {
		t.Errorf("parseListing() error = %v, want <nil>", err)
	}

	if out := lsR(root); out != exampleOut {
		t.Errorf("lsR(root) diff: %v", cmp.Diff(exampleOut, out))
	}

	if want, got := 584, cd(root).cd("a").cd("e").size(); got != want {
		t.Errorf("/a/e size = %d, want %d", got, want)
	}

	if want, got := 94853, cd(root).cd("a").size(); got != want {
		t.Errorf("/a size = %d, want %d", got, want)
	}

	if want, got := 24933642, cd(root).cd("d").size(); got != want {
		t.Errorf("/d size = %d, want %d", got, want)
	}

	if want, got := 48381165, root.size(); got != want {
		t.Errorf("/ size = %d, want %d", got, want)
	}

	directories := dirs(root.(dir))

	if want, s := 95437, part1(directories); s != want {
		t.Errorf("part1() = %d, want %d", s, want)
	}

	if want, s := 24933642, part2(root, directories); s != want {
		t.Errorf("part1() = %d, want %d", s, want)
	}
}

type cder struct{ dir }

func (c cder) cd(s string) cder {
	return cder{c.dir.get(s).(dir)}
}

func cd(n node) cder {
	return cder{n.(dir)}
}
