package staff

import (
	"bytes"
	"encoding/xml"

	"golang.org/x/net/html/charset"
)

func Win2UTF(data []byte) *xml.Decoder {
	reader := bytes.NewReader(data)
	xmlDecoder := xml.NewDecoder(reader)
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	return xmlDecoder
}
