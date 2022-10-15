package _func

import "bufio"

func ReadFull(r bufio.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err != nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
