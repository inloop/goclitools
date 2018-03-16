package goclitools

import (
	"bytes"
	"os"
	"strings"
)

func SecureString(input string, secrets []string) string {
	if len(secrets) == 0 {
		return input
	}

	r := make([]string, 2*len(secrets))
	for i, e := range secrets {
		r[i*2] = e
		r[i*2+1] = "*******"
	}

	return strings.NewReplacer(r...).Replace(input)
}

func SecureByteArray(input []byte, secrets []string) []byte {
	if len(secrets) == 0 {
		return input
	}

	data := input
	for _, e := range secrets {
		data = bytes.Replace(data, []byte(e), []byte("*******"), -1)
	}
	return data
}

func SecureStd(out *os.File, secrets []string) *os.File {
	if len(secrets) == 0 {
		return out
	}

	readFile, writeFile, err := os.Pipe()
	if err != nil {
		return out
	}

	go func() {
		defer readFile.Close()
		// MEMO: secrets may be split to multiple buffers and not detected
		//  possible solution https://golang.org/pkg/bufio/#example_Scanner_lines
		var data [250]byte
		for {
			n, err := readFile.Read(data[:])
			if err != nil {
				break
			}
			out.Write(SecureByteArray(data[:n], secrets))
		}
	}()

	return writeFile
}
