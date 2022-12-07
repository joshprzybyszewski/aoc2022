package seven

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	rootDirName = `/`
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	db := newDirBuilder()
	db.processLines(lines)

	dirs := db.allDirsWithMaxSize(100000)

	total := 0
	for _, d := range dirs {
		total += d.size()
	}

	return fmt.Sprintf("%d", total), nil
}

func isCDCmd(line string) bool {
	return len(line) >= 4 && line[:4] == `$ cd`
}

func isLSCmd(line string) bool {
	return len(line) >= 4 && line[:4] == `$ ls`
}

func isDir(line string) bool {
	return len(line) >= 4 && line[:4] == `dir `
}

type dirBuilder struct {
	root *dir

	curDir *dir
}

func newDirBuilder() *dirBuilder {
	root := newRootDir()
	return &dirBuilder{
		root:   root,
		curDir: root,
	}
}

func (db *dirBuilder) allDirs() []*dir {
	children := db.root.getAllChildren()

	output := make([]*dir, 0, 1+len(children))
	output = append(output, db.root)
	output = append(output, children...)

	return output
}

func (db *dirBuilder) allDirsWithMaxSize(
	maxSize int,
) []*dir {

	var output []*dir
	if db.root.size() < maxSize {
		output = append(output, db.root)
	}

	children := db.root.getAllChildren()
	for _, c := range children {
		if c.size() < maxSize {
			output = append(output, c)
		}
	}

	return output
}

func (db *dirBuilder) processLines(
	lines []string,
) error {

	for _, line := range lines {
		db.processLine(line)
	}

	return nil
}

func (db *dirBuilder) processLine(line string) error {
	if isDir(line) {
		d, err := newDir(
			db.curDir,
			line,
		)
		if err != nil {
			return err
		}
		db.curDir.addChild(d)
		return nil
	}
	if isCDCmd(line) {
		// change the curDir in the builder
		return db.handleCD(line)
	}
	if isLSCmd(line) {
		// do nothing, we're gonna parse the output soon
		return nil
	}

	f, err := newFile(line)
	if err != nil {
		return err
	}
	db.curDir.addFile(f)

	return nil
}

func (db *dirBuilder) handleCD(line string) error {
	parts := strings.Split(line, ` `)
	if len(parts) != 3 {
		return fmt.Errorf("line should have three parts: %q", line)
	}
	newDir := parts[2]
	if newDir == db.root.name {
		db.curDir = db.root
		return nil
	}
	if newDir == `..` {
		db.curDir = db.curDir.parent
		return nil
	}

	child := db.curDir.getChild(newDir)
	if child == nil {
		return fmt.Errorf("does not have a child with name: %q", newDir)
	}
	db.curDir = child
	return nil
}

type dir struct {
	name  string
	files []*file

	parent   *dir
	children []*dir
}

func newRootDir() *dir {
	return &dir{
		name: rootDirName,
	}
}

func newDir(
	parent *dir,
	line string,
) (*dir, error) {
	if parent == nil {
		return nil, fmt.Errorf("requires a parent")
	}

	parts := strings.Split(line, ` `)
	if len(parts) != 2 {
		return nil, fmt.Errorf("line should have two parts: %q", line)
	}

	if parts[0] != `dir` {
		return nil, fmt.Errorf("first part should be 'dir': %q", line)
	}

	name := parts[1]
	if name == rootDirName {
		return nil, fmt.Errorf("only the root can be named %q: %q", rootDirName, line)
	}

	return &dir{
		parent: parent,
		name:   name,
	}, nil
}

func (d *dir) addChild(c *dir) {
	d.children = append(d.children, c)
}

func (d *dir) getChild(name string) *dir {
	for _, c := range d.children {
		if c.name == name {
			return c
		}
	}
	return nil
}

func (d *dir) addFile(f *file) {
	d.files = append(d.files, f)
}

func (d *dir) size() int {
	// TODO memoize
	total := 0
	for _, f := range d.files {
		total += f.size
	}
	for _, d := range d.children {
		total += d.size()
	}
	return total
}

func (d *dir) getAllChildren() []*dir {
	output := make([]*dir, 0, len(d.children))
	for _, c := range d.children {
		output = append(output, c)
		output = append(output, c.getAllChildren()...)
	}
	return output
}

type file struct {
	name string
	size int
}

func newFile(line string) (*file, error) {
	parts := strings.Split(line, ` `)
	if len(parts) != 2 {
		return nil, fmt.Errorf("line should have two parts: %q", line)
	}

	size, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	return &file{
		name: parts[1],
		size: size,
	}, nil
}
