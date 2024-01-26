package file

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func GetLastLine(f *os.File) (string, error) {
	line := ""
	var cursor int64 = 0
	stat, _ := f.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		f.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		f.Read(char)

		if char[0] == 0 { // stop if file is empty
			return "", errors.New("file is empty")
		}

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the beginning
			break
		}
	}

	return line, nil
}
