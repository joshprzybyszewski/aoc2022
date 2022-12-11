package seven

func isDir(line string) bool {
	// this assumes that there cannot be a file named "dir".
	return len(line) >= 4 && line[:4] == `dir `
}
