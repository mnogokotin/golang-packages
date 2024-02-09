package file

import (
	"fmt"
	"io"
	"os"
)

type EmptyFileError struct{}

func (err EmptyFileError) Error() string {
	return fmt.Sprint("file is empty")
}

func GetLastLine(f *os.File) (string, error) {
	line := ""
	var cursor int64 = 0
	stat, _ := f.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		_, err := f.Seek(cursor, io.SeekEnd)
		if err != nil {
			return "", err
		}

		char := make([]byte, 1)
		_, err = f.Read(char)
		if err != nil {
			return "", err
		}

		if char[0] == 0 { // stop if file is empty
			var err EmptyFileError
			return "", err
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
