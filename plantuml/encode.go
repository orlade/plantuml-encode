package plantuml

import (
	"bytes"
	"compress/zlib"
	"io"
)

// DeflateAndEncode compresses some PlantUML source text into the URL-encoded format accepted by the
// PlantUML server.
func DeflateAndEncode(text string) (string, error) {
	return DeflateAndEncodeBytes([]byte(text))
}

// The functions below ported from https://github.com/dougn/python-plantuml/blob/master/plantuml.py

// DeflateAndEncode compresses some PlantUML source text into the URL-encoded format accepted by the
// PlantUML server.
func DeflateAndEncodeBytes(text []byte) (string, error) {
	var buf bytes.Buffer
	zw, err := zlib.NewWriterLevel(&buf, zlib.BestCompression)
	if err != nil {
		return "", err
	}
	if _, err := zw.Write(text); err != nil {
		return "", err
	}
	if err := zw.Close(); err != nil {
		return "", err
	}
	return encode(buf.Bytes()), nil
}

func encode(data []byte) string {
	var buf bytes.Buffer
	i := 0
	for wholeTripleBytes := len(data) / 3 * 3; i < wholeTripleBytes; i += 3 {
		encode3bytes(&buf, data[i], data[i+1], data[i+2])
	}
	switch len(data) - i {
	case 1:
		encode3bytes(&buf, data[i], 0, 0)
	case 2:
		encode3bytes(&buf, data[i], data[i+1], 0)
	}
	return buf.String()
}

// 3 bytes takes 24 bits. This splits 24 bits into 4 bytes of which lower 6-bit takes account.
func encode3bytes(w io.ByteWriter, b1, b2, b3 byte) {
	if err := w.WriteByte(encode6bit(0x3F & (b1 >> 2))); err != nil {
		panic(err)
	}
	if err := w.WriteByte(encode6bit(0x3F & (((b1 & 0x3) << 4) | (b2 >> 4)))); err != nil {
		panic(err)
	}
	if err := w.WriteByte(encode6bit(0x3F & (((b2 & 0xF) << 2) | (b3 >> 6)))); err != nil {
		panic(err)
	}
	if err := w.WriteByte(encode6bit(0x3F & b3)); err != nil {
		panic(err)
	}
}

func encode6bit(b byte) byte {
	// 6 bit makes value 0 to 63. The func maps 0-63 to characters
	// '0'-'9', 'A'-'Z', 'a'-'z', '-', '_'. '?' should never be reached.
	if b > 63 {
		panic("unexpected character!")
	}
	return "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"[b]
}
