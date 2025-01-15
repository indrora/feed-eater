package sources

import (
	"errors"
	"io"
	"os"
)

// TextFile implements DataSource for reading and outputting text file contents
type TextFile struct {
	filepath string
}

// Configure sets up the TextFile source with the provided filepath
func (t *TextFile) Configure(config map[string]string) error {
	filepath, exists := config["filepath"]
	if !exists {
		return errors.New("filepath configuration is required")
	}

	t.filepath = filepath
	return nil
}

// Print reads and writes the text file contents to the provided writer
func (t *TextFile) Print(writer io.Writer) {
	file, err := os.Open(t.filepath)
	if err != nil {
		writer.Write([]byte("Error reading file: " + err.Error()))
		return
	}
	defer file.Close()

	_, err = io.Copy(writer, file)
	if err != nil {
		writer.Write([]byte("Error copying file contents: " + err.Error()))
	}
}
