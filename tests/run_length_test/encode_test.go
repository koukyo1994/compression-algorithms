package runlength_test

import (
	"io/ioutil"
	"testing"

	runlength "github.com/koukyo1994/compression-algorithms/runlength"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{name: "abc", filePath: "../../testdata/run_length/abc.txt", expected: "a8b6c2\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer, err := ioutil.ReadFile(tt.filePath)
			if err != nil {
				t.Fatal(err)
			}
			lines := string(buffer)

			actual := runlength.Encode(lines)
			if actual != tt.expected {
				assert.Equal(t, tt.expected, actual, "Encode: %s", tt.name)
			}
		})
	}
}
