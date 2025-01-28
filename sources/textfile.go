package sources

import (
	"errors"
	"io"
	"math/rand"
	"os"
	"path/filepath"
)

// TextFile implements DataSource for reading and outputting text file contents
type TextFile struct {
	fpath  string
	filter string
	files  []string
}

// Configure sets up the TextFile source with the provided filepath and optional filter
func (t *TextFile) Configure(config map[string]string) error {
	fpath, exists := config["filepath"]
	if !exists {
		return errors.New("filepath configuration is required")
	}

	t.fpath = fpath
	t.filter = config["filter"]

	// Check if path is a directory
	info, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		// If no filter is provided, default to all files
		if t.filter == "" {
			t.filter = "*"
		}

		// Enumerate files matching the glob pattern
		matches, err := filepath.Glob(filepath.Join(fpath, t.filter))
		if err != nil {
			return err
		}

		if len(matches) == 0 {
			return errors.New("no files found matching the filter")
		}

		t.files = matches
	} else {
		// If it's a single file, just add it to files
		t.files = []string{fpath}
	}

	return nil
}

// Print reads and writes the text file contents to the provided writer
func (t *TextFile) Print(writer io.Writer) {
	// Select a file (random if multiple, first if single)
	var selectedFile string
	if len(t.files) > 1 {
		selectedFile = t.files[rand.Intn(len(t.files))]
	} else if len(t.files) == 1 {
		selectedFile = t.files[0]
	} else {
		writer.Write([]byte("No files available"))
		return
	}

	file, err := os.Open(selectedFile)
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
