package generator

import (
	"encoding/base64"
)

const m_UUID_LENGTH = 64

func UuidGenerator() string {
	bytes := make([]byte, m_UUID_LENGTH)
	return base64.URLEncoding.EncodeToString(bytes)
}
