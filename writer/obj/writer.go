package obj

import "os"

type writer struct {
	lines
	fp *os.File
}
