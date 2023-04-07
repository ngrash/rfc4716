// Package rfc4716 provides a function to encode a byte slice into a RFC4716 formatted string.
package rfc4716

import (
	"bytes"
	"encoding/base64"
)

// Encode encodes a byte slice into a RFC4716 formatted string.
// Note that this function assumes that the input byte slice contains a valid public key in a format
// that can be encoded as base64. If the input is not a valid public key, the output will not be a
// valid RFC4716-formatted public key.
func Encode(bs []byte) string {
	encoded := base64.StdEncoding.EncodeToString(bs)
	var buf bytes.Buffer
	buf.WriteString("---- BEGIN SSH2 PUBLIC KEY ----\n")
	for i := 0; i < len(encoded); i += 70 {
		end := i + 70
		if end > len(encoded) {
			end = len(encoded)
		}
		buf.WriteString(encoded[i:end] + "\n")
	}
	buf.WriteString("---- END SSH2 PUBLIC KEY ----\n")
	return buf.String()
}
