package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

//
type BoundedString struct {
	a     string
	limit int
}

/*
 Writer is the interface that wraps the basic Write method.

Write writes len(p) bytes from p to the underlying data stream. It returns the number
of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the
write to stop early. Write must return a non-nil error if it returns n < len(p).
 Write must not modify the slice data, even temporarily.

Implementations must not retain p.

type Writer interface {
        Write(p []byte) (n int, err error)
}
*/

// We now implement Writer
func (sr *BoundedString) Write(p []byte) (n int, err error) {
	sr.a = string(p[0:sr.limit])
	if len(p) > sr.limit {
		return sr.limit, errors.New("exceeded limit")
	}
	return len(p), nil
}

/*
 Reader is the interface that wraps the basic Read method.

Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any
error encountered. Even if Read returns n < len(p), it may use all of p as scratch space during the
call. If some data is available but not len(p) bytes, Read conventionally returns what is available
instead of waiting for more.

When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it
returns the number of bytes read. It may return the (non-nil) error from the same call or return the
error (and n == 0) from a subsequent call. An instance of this general case is that a Reader returning
a non-zero number of bytes at the end of the input stream may return either err == EOF or err == nil.
The next Read should return 0, EOF.

Callers should always process the n > 0 bytes returned before considering the error err. Doing so
correctly handles I/O errors that happen after reading some bytes and also both of the allowed EOF
behaviors.

Implementations of Read are discouraged from returning a zero byte count with a nil error, except
when len(p) == 0. Callers should treat a return of 0 and nil as indicating that nothing happened;
in particular it does not indicate EOF.

Implementations must not retain p.

type Reader interface {
        Read(p []byte) (n int, err error)
}
*/

func (sr *BoundedString) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p) && i < sr.limit; i++ {
		p[i] = sr.a[i]
	}
	if len(p) > sr.limit {
		return sr.limit, nil
	}
	return len(p), nil
}

func main() {
	msr := BoundedString{limit: 3}
	n, err := msr.Write([]byte{'N', 'i', 'k', 'o', 'l', 'a', 's'})
	fmt.Println(n, err, msr.a)

	buf := make([]byte, 5)
	fmt.Println(msr.Read(buf))

	buf = make([]byte, 2)
	fmt.Println(msr.Read(buf))

	buf = make([]byte, 40)
	msr.Read(buf)
	os.Stdout.Write(buf)
}
