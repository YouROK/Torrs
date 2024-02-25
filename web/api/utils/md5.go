package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func HashStr(str string) string {
	return HashBytes([]byte(str))
}

func HashBytes(buf []byte) string {
	md5 := md5.New()
	_, err := md5.Write(buf)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(md5.Sum(nil))
}

func HashReader(reader io.Reader) string {
	md5 := md5.New()
	io.Copy(md5, reader)
	return hex.EncodeToString(md5.Sum(nil))
}
