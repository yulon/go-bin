package bin

import "io"

type Writer struct{
	io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w}
}

func (w *Writer) Byte(data interface{}) (int, error) {
	return w.Write(Byte(data))
}

func (w *Writer) Word(data interface{}) (int, error) {
	return w.Write(Word(data))
}

func (w *Writer) Dword(data interface{}) (int, error) {
	return w.Write(Dword(data))
}

func (w *Writer) Qword(data interface{}) (int, error) {
	return w.Write(Qword(data))
}

func (w *Writer) WordB(data interface{}) (int, error) {
	return w.Write(WordB(data))
}

func (w *Writer) DwordB(data interface{}) (int, error) {
	return w.Write(DwordB(data))
}

func (w *Writer) QwordB(data interface{}) (int, error) {
	return w.Write(QwordB(data))
}

func (w *Writer) Cstr(text string) (int, error) {
	return w.Write(Cstr(text))
}

func (w *Writer) Zeros(size int64) (int, error) {
	return w.Write(Zeros(size))
}
