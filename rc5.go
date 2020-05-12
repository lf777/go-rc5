package rc5

import "crypto/cipher"

//cip, err := rc5.NewCipher16(key, NUM_RC5_ROUNDS)
//func NewCipher16(key []byte, rounds uint) (cipher.Block, error) {

func NewRc5Cipher16(key []byte, rounds uint) (cipher.Block, error) {
	return NewCipher16(key, rounds)
}
