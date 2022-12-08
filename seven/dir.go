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

	knownSize int
}

func newRootDir() *dir {
	return &dir{
		name:      rootDirName,
		knownSize: -1,
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
		return nil, fmt.Errorf("dir line should have two parts: %q", line)
	}

	if parts[0] != `dir` {
		return nil, fmt.Errorf("first part should be 'dir': %q", line)
	}

	name := parts[1]
	if name == rootDirName {
		return nil, fmt.Errorf("only the root can be named %q: %q", rootDirName, line)
	}

	return &dir{
		parent:    parent,
		name:      name,
		knownSize: -1,
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

func (d *dir) complete() {
	d.knownSize = d.size()
	for _, c := range d.children {
		c.complete()
	}
}

func (d *dir) size() int {
	if d.knownSize >= 0 {
		return d.knownSize
	}

	total := 0
	for _, f := range d.files {
		total += f.size
	}
	for _, d := range d.children {
		total += d.size()
	}
	return total
}

func (d *dir) getChildrenWithMaxSize(
	maxSize int,
) []*dir {

	output := make([]*dir, 0, len(d.children))
	for _, c := range d.children {
		if c.size() <= maxSize {
			output = append(output, c)
		}
		output = append(output, c.getChildrenWithMaxSize(maxSize)...)
	}
	return output
}

func (d *dir) getMinSizeGreaterThan(
	floor int,
) int {
	if d.size() < floor {
		return -1
	}

	min := d.size()
	var otherMin int
	for _, c := range d.children {
		otherMin = c.getMinSizeGreaterThan(floor)
		if otherMin != -1 && otherMin < min {
			min = otherMin
		}
	}
	return min
}
