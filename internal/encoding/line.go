package encoding

//text
//binary
//json
//ndjson
//influx

import (
	"io"
)

func EncodeLine(w io.Writer, data interface{}) error {
	return nil
}

func DecodeLine(r io.Reader) (interface{}, error) {

	return nil, nil
}
