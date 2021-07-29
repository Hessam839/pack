package pack

import "io"

func Process(in io.Reader, out io.Writer) error{
	for {
		_, err := io.Copy(out, in)
		if err != nil {
			return err
		}
	}
}