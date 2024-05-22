package writer

import (
	"bytes"
	"context"
	"io"
)

type IOWriteCloser struct {
	io.WriteCloser
	context    context.Context
	wof_writer Writer
	key        string
}

func NewIOWriteCloser(ctx context.Context, wof_writer Writer, key string) io.Writer {

	wr := &IOWriteCloser{
		context:    ctx,
		wof_writer: wof_writer,
		key:        key,
	}

	return wr
}

func (wr *IOWriteCloser) Write(p []byte) (int, error) {

	br := bytes.NewReader(p)
	n, err := wr.wof_writer.Write(wr.context, wr.key, br)

	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func (wr *IOWriteCloser) Close() error {
	return wr.wof_writer.Close(wr.context)
}
