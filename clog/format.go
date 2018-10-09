package clog

type Formater interface {
	Write(m interface{}) (err error)
}
