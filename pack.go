package pack

import (
	"context"
	"io"
)

func Process(ctx context.Context, in io.Reader, out io.Writer) error{
	for {
		_, err := io.Copy(out, in)
		if err != nil {
			return err
		}
	}
}

func Add(arg1 int, arg2 int) int {
	return arg1 + arg2
}