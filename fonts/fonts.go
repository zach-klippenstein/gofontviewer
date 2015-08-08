package fonts

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"

	"compress/gzip"
)

//go:generate ./encode_ttf.sh "Roboto-Regular.ttf" RobotoRegular

type Font string

func LoadFont(font Font) ([]byte, error) {
	fontDataEncoded := bytes.NewBuffer([]byte(font))
	fontDataCompressed := base64.NewDecoder(base64.StdEncoding, fontDataEncoded)
	fontDataTtf, err := gzip.NewReader(fontDataCompressed)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(fontDataTtf)
}
