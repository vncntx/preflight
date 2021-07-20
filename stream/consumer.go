package stream

import "os"

// Consumer is a function that consumes a file
type Consumer func(w *os.File)
