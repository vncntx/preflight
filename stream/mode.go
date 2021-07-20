package stream

import "io/fs"

func isAppend(mode fs.FileMode) bool {
	return mode&fs.ModeAppend == fs.ModeAppend
}

func isPipe(mode fs.FileMode) bool {
	return mode&fs.ModeNamedPipe == fs.ModeNamedPipe
}
