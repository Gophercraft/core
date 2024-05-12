package connection

type dummy_cipher struct {
}

func (d dummy_cipher) Decrypt(data, tag []byte) error {
	return nil
}

func (d dummy_cipher) Encrypt(data, tag []byte) error {
	return nil
}

func NewDummyCipher() (connection_cipher Cipher) {
	return dummy_cipher{}
}
