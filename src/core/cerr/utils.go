package cerr

// Writer is the error writer.
type Writer interface {
	Write(s string) int
}

func WriteJoinE(w Writer, sep string, s ...E) int {
	if len(s) == 0 {
		return w.Write("")
	}

	n := s[0].WriteErr(w)
	for _, e := range s[1:] {
		n += w.Write(sep) + e.WriteErr(w)
	}

	return n
}

func WriteJoinS(w Writer, sep string, segments ...string) int {
	if len(segments) == 0 {
		return w.Write("")
	}

	n := w.Write(segments[0])
	for _, s := range segments[1:] {
		n += w.Write(sep) + w.Write(s)
	}

	return n
}
