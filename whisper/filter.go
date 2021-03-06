package whisper

import "crypto/ecdsa"

type Filter struct {
	To     *ecdsa.PublicKey
	From   *ecdsa.PublicKey
	Topics [][]byte
	Fn     func(*Message)
}
