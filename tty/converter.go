package tty

import (
	"io"
)

// TTYConverter handles conversion of modern ASCII text to Teletype-compatible format
type TTYConverter struct {
	writer       io.Writer
	replacements map[rune]string
}

// NewTTYConverter creates a new TTY converter with default mappings
func NewTTYConverter(w io.Writer) *TTYConverter {
	t := &TTYConverter{
		writer:       w,
		replacements: make(map[rune]string),
	}
	t.initDefaultMappings()
	return t
}

func (t *TTYConverter) initDefaultMappings() {
	t.replacements = map[rune]string{
		'@':  "AT",
		'\\': "/",
		'[':  "(",
		']':  ")",
		'{':  "(",
		'}':  ")",
		'|':  "!",
		'~':  "-",
		'`':  "'",
		'©':  "(C)",
		'®':  "(R)",
		'™':  "(TM)",
		'•':  "*",
		'…':  "...",
		'"':  "\"",
		'\'': "'",
		'—':  "--",
		'–':  "-",
		'¡':  "!",
		'¿':  "?",
		'☎':  "\a", // Bell character
	}
}

// Write implements io.Writer interface
func (t *TTYConverter) Write(p []byte) (n int, err error) {
	input := string(p)
	var output []byte

	for _, r := range input {
		if replacement, exists := t.replacements[r]; exists {
			output = append(output, []byte(replacement)...)
		} else if r <= 127 && r >= 32 {
			output = append(output, byte(r))
		} else {
			output = append(output, ' ')
		}
	}

	return t.writer.Write(output)
}

// AddReplacement allows adding custom character mappings
func (t *TTYConverter) AddReplacement(from rune, to string) {
	t.replacements[from] = to
}
