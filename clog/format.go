package clog

import (
	"io"
)

type Formater interface {
	io.ReadWriter
}
