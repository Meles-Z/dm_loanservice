package uuid

import (
	"crypto/rand"
	"io"
	"math"
	"math/bits"

	"github.com/google/uuid"
)

var defaultCharset = []byte("_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// UniqueIDString generates a secure randomized id
func UniqueIDString(size int) (string, error) {
	buf := make([]byte, size)

	_, err := UniqueID(defaultCharset, buf, 0, size)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// UniquesID generates a secure randomized id
func UniqueID(charset, outb []byte, offset, size int) (int, error) {
	if len(charset) == 0 {
		charset = defaultCharset
	}

	ncharset := len(charset)
	size = min(len(outb)-offset, size)

	mask := (2 << uint(31-bits.LeadingZeros32(uint32(ncharset-1)))) - 1
	step := int(math.Ceil(1.6 * float64(mask) * float64(size) / float64(ncharset)))

	var (
		idx int
		err error

		iout = offset
		nout int

		istep1 int
		step1  int
	)

	for {
		istep1 = iout
		step1 = min(size-nout, step)

		_, err = io.ReadAtLeast(rand.Reader, outb[iout:], step1)
		if err != nil {
			return nout, err
		}

		for i := 0; i < step1; i++ {
			idx = int(outb[istep1+i]) & mask

			if idx < ncharset {
				outb[iout] = charset[idx]
				iout++
				nout++

				if nout >= size {
					return nout, nil
				}
			}
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func UUID() string {
	return uuid.New().String()
}
