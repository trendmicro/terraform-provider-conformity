package cloudconformity

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
    "log"
    "encoding/base64"
)

// GenerateKeyPair generates a new key pair
func CheckError(e error) {
    if e != nil {
        log.Println(e.Error)
    }
}
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	CheckError(err)
	return privkey, &privkey.PublicKey
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	CheckError(err)
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		CheckError(err)
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	CheckError(err)
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		log.Println("not ok")
	}
	return key
}

func EncryptWithPublicKey(secretMessage string, key *rsa.PublicKey) string {
    label := []byte("OAEP Encrypted")
    rng := rand.Reader
    ciphertext, err := rsa.EncryptOAEP(sha512.New(), rng, key, []byte(secretMessage), label)
    CheckError(err)
    return base64.StdEncoding.EncodeToString(ciphertext)
}
