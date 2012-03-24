package shattr

import "testing"
import "github.com/sdegutis/assert"

import "bytes"
import "fmt"

func TestShellAttributes(t *testing.T) {
  var buffer bytes.Buffer

  output := NewWriter(&buffer, Underline, Blue)
  output.Write([]byte("testing"))
  assert.Equals(t, buffer.String(), fmt.Sprintf("%c[0;4;34mtesting%c[0m", 27, 27))
}
