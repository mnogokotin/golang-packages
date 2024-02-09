package file

import (
	"fmt"
	"io"
	"os"
)

func NewEmptyFileError(op, fileName string, err error) *EmptyFileError {
	return &EmptyFileError{op: op, fileName: fileName, err: err}
}

type EmptyFileError struct {
	op       string
	fileName string
	err      error
}

func (err EmptyFileError) Error() string {
	return fmt.Errorf("%s: file %s is empty: %v", err.op, err.fileName, err.err).Error()
}

func GetLastLine(f *os.File) (string, error) {
	const op = "utils.file.GetLastLine"

	line := ""
	var cursor int64 = 0
	stat, _ := f.Stat()
	fileName := stat.Name()
	fileSize := stat.Size()
	for {
		cursor -= 1
		_, err := f.Seek(cursor, io.SeekEnd)
		if err != nil { // stop if file is empty
			return "", NewEmptyFileError(op, fileName, err)
		}

		char := make([]byte, 1)
		_, err = f.Read(char)
		if err != nil {
			return "", fmt.Errorf("%s: error read file %s: %w", op, fileName, err)
		}

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -fileSize { // stop if we are at the beginning
			break
		}
	}

	return line, nil
}
