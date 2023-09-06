// Package cerr defines core error types.
package cerr

type E interface {
	WriteErr(w Writer) int
}
