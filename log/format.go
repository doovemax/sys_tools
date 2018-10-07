package log

import (
	"io"
)

type Formater interface {
	io.ReadWriter
}
