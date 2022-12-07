package seven

import (
	"fmt"
	"strings"
)

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
