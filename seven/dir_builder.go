package seven

import (
	"fmt"
	"strings"
)

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
