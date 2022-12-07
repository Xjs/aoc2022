package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type node interface {
	name() string
	size() int
	isDir() bool
}

type dir interface {
	node
	add(n node)
	get(name string) node
	entries() []node
}

type folder struct {
	// name
	n string
	// contents
	c []node
	// contents by name
	byName map[string]int

	sizeCache int
}

func (f *folder) name() string {
	return f.n
}

func (f *folder) size() int {
	if f.sizeCache != 0 {
		return f.sizeCache
	}

	var sum int
	for _, node := range f.c {
		sum += node.size()
	}
	return sum
}

func (f *folder) isDir() bool {
	return true
}

func (f *folder) add(n node) {
	f.sizeCache = 0

	idx, ok := f.byName[n.name()]
	if ok {
		f.c[idx] = n
		return
	}

	if f.byName == nil {
		f.byName = make(map[string]int)
	}

	idx = len(f.c)
	f.c = append(f.c, n)
	f.byName[n.name()] = idx
}

func (f *folder) get(n string) node {
	idx, ok := f.byName[n]
	if !ok {
		return nil
	}
	return f.c[idx]
}

func (f *folder) entries() []node {
	return f.c
}

type file struct {
	// name
	n string
	// size
	s int
}

func (f file) name() string { return f.n }
func (f file) size() int    { return f.s }
func (f file) isDir() bool  { return false }

func parseListing(lst string) (node, error) {
	var currentFolder []string
	root := new(folder)

	lines := strings.Split(lst, "\n")
	var isLs bool

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if line[0] == '$' {
			if isLs {
				isLs = false
			}

			if len(line) < 3 {
				return nil, fmt.Errorf("invalid line %q", line)
			}

			fs := strings.Fields(line[2:])
			switch fs[0] {
			case "cd":
				if len(fs) != 2 {
					return nil, fmt.Errorf("cd expects 1 argument, got %q", line)
				}
				if fs[1] == "/" {
					currentFolder = nil
				} else if fs[1] == ".." {
					if len(currentFolder) < 1 {
						return nil, fmt.Errorf("cannot cd out of root")
					}
					currentFolder = currentFolder[:len(currentFolder)-1]
				} else {
					currentFolder = append(currentFolder, fs[1])
				}
			case "ls":
				if len(fs) != 1 {
					return nil, fmt.Errorf("ls expects no arguments, got %q", line)
				}
				isLs = true
				continue
			default:
				return nil, fmt.Errorf("invalid command line %q", line)
			}
		}

		if isLs {
			var cur dir = root
			for _, leaf := range currentFolder {
				next := cur.get(leaf)
				if next == nil {
					return nil, fmt.Errorf("error in cd to %v (%q does not exist)", currentFolder, leaf)
				}
				c, ok := next.(dir)
				if !ok {
					return nil, fmt.Errorf("error in cd to %v (%q is not a folder)", currentFolder, leaf)
				}
				cur = c
			}

			fs := strings.Fields(line)
			if len(fs) != 2 {
				return nil, fmt.Errorf("invalid ls output: %q", line)
			}
			if fs[0] == "dir" {
				cur.add(&folder{n: fs[1]})
			} else {
				s, err := strconv.Atoi(fs[0])
				if err != nil {
					return nil, fmt.Errorf("could not parse size %q of %q: %w", fs[0], fs[1], err)
				}
				cur.add(file{n: fs[1], s: s})
			}
		}
	}

	return root, nil
}

func lsR(n node) string {
	return ls(n, "")
}

func ls(n node, prefix string) string {
	var b strings.Builder

	if n.isDir() {
		f := n.(dir)
		name := f.name()
		if name == "" {
			name = "/"
		}
		fmt.Fprintf(&b, "%s- %s (dir)\n", prefix, name)
		for _, c := range f.entries() {
			b.WriteString(ls(c, "\t"+prefix))
		}
	} else {
		f := n.(file)
		fmt.Fprintf(&b, "%s- %s (file, size=%d)\n", prefix, f.n, f.s)
	}

	return b.String()
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	root, err := parseListing(string(input))
	if err != nil {
		log.Fatal(err)
	}

	directories := dirs(root.(dir))

	fmt.Printf("Part 1: %d\n", part1(directories))
	fmt.Printf("Part 2: %d\n", part2(root, directories))
}

func part1(directories []dir) int {
	var sum int
	for _, d := range directories {
		s := d.size()
		if s <= 100000 {
			sum += s
		}
	}

	return sum
}

func dirs(root dir) []dir {
	var directories []dir
	var getD func(d dir)
	getD = func(d dir) {
		for _, n := range d.entries() {
			if n.isDir() {
				directories = append(directories, n.(dir))
				getD(n.(dir))
			}
		}
	}
	getD(root)
	return directories
}

func part2(root node, directories []dir) int {
	const (
		total       = 70000000
		needAtLeast = 30000000
	)

	rootSize := root.size()
	unused := total - rootSize
	needDelete := needAtLeast - unused

	if needDelete <= 0 {
		return 0
	}

	var smallest int = rootSize
	for _, d := range directories {
		s := d.size()
		if s > needDelete && s < smallest {
			smallest = s
		}
	}

	return smallest
}
