// package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
    "log"
)

// func main() {
//     // privkey, PublicKey := GenerateKeyPair(2048)
//     // fmt.Println(privkey)
//     // fmt.Println(PublicKey)
//     // privBytes :=  PrivateKeyToBytes(privkey)
//     // pubBytes :=  PublicKeyToBytes(PublicKey)
//     privBytes := []byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA2rryzyobjdhXIZRXDv/8JXfKhzbbOAsQC+QgRfYSEzW0qUTX\nGho0S9JrFGgJSCT2tIVrfKVqYcqOyLZ+/1N+N4c7t3jvxcYo7BExp1eqbkH9G579\nhQsSoXOS3YZycCt7/YSqJNvn/GCQztTuEmLBE3EiLrWB0wGquv5mA8pDmCShCXxU\nEcsKEJgS2RRDiT4YzpXK0R/Twua4TB/QfE7eiHMQMG/bVebF+fLVVH4o3qLjcyq6\n2tnT/r5knciOHAKBUn4WAkCM00hYzhXmsXa2+GO+A9A++zBH65i03LeskfImR40R\nrq6NRgTjbeiheQCb2JR4Twzb12Z28QqY/oRn+wIDAQABAoIBAALOjVkdODdMxGl4\n5tkZbdnpPJ8ZlByW/8C3T7a2HqtCcCwP5xa9qVgjvh4H676SQtw0LhnuYXwZxlVL\nCjwqjR8XTCvhkGogTdwhqFp2ZIh/rkjRdH1lk/qgag0PsZ5A4JlzP1+ztrllX9ZX\noyp3O/UM6Zxh1eWoStGVfCemS9HC+H74itVRiArThdSEDzUUMNcptL9sr8OfkXxA\nKMnQOlbve/WzfWrhIeHMz6Z6z8FX/jp7eNYCODF3Y/NTH8GvTQkAvrR13fyHgBMz\noxAyzTmqwK5K75/50KAwpbAsbSrMUXQiX+dZeJvjsntPxSoy20ko/AUHhZk7hRY2\nz9pWZUECgYEA8aOIFutPT38Om16ClDoTyU+0tHYTYpboQYsI3TwShBoX2eUN3Yaq\n7R0Ep7lg51v9bXk1gYQzDwFZWudgVkiH3TLB9J/fArPK4X3ggto3SK2OguxeqVhB\nBwxkWKRqrb+Wv2zT8k8rzVNITqh9LTSD2LgWDZ0IK3LedM4uDyeKqJsCgYEA57re\nr6Wcfn4wE+2qbOh2IWxy7AEeWB3j8rMyB6dsGO95rxWr7KaUzezaQU2n/ZT8pjsw\nyhYaE1DtSDN+HC8V+FmjQjJqiRbRmc0Zm+aCRQ5DnzmnLxSlZ9VhcszQFn8btKSX\noIXAQEuHMsB5AdKiqik9OHcauBPUUMCSPpuKxCECgYEA7pFGC1rHMvV+tmbZBP9S\nCa9n+cOZ3/yd0hgy6DonDcW1RquexNfwaan8rpuX0NRBoZPJ/9VFk8sBLX7C3m09\ntmYmmB4/T6uy4m4k+wv3CQpRaXF1BDzd9teFOv8ZU/GUI+qOVu1TkaRn/0DaVYdD\npPQa0dX3+u2uNCRb1Rp1C4ECgYEAjMQssCB3XzPCeuid5YiU8hrR+OF5EGgf53fJ\nhXLDrKYUkjIk/R34ONuPfanxyY4up8A/FBO3BVLcwUZebjqAKxwwm27K3roY20gH\nLXgqXE6c72VzVJtDGz848ibOpUvThbmSTjXRonz/BOp814mfvKROhzV66qVJDUDd\n1eBgvGECgYEAnRo8It+Z/15SIPUYyxc0OSEKIOpOnSY5HodIR9C5WN1WK3Kvc/KM\nJg2kLiIX4nx/1zR11Y/t7PP0peKJSfHjckQns29pQofemOGG3fy3JPlf9SVSGPqg\nt6HmlnXA3Xz14wylkh3OX6JItV4pPn8ZfhPTCDLO0mSvqVt+E0myw0E=\n-----END RSA PRIVATE KEY-----")
//     pubBytes := []byte("-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2rryzyobjdhXIZRXDv/8\nJXfKhzbbOAsQC+QgRfYSEzW0qUTXGho0S9JrFGgJSCT2tIVrfKVqYcqOyLZ+/1N+\nN4c7t3jvxcYo7BExp1eqbkH9G579hQsSoXOS3YZycCt7/YSqJNvn/GCQztTuEmLB\nE3EiLrWB0wGquv5mA8pDmCShCXxUEcsKEJgS2RRDiT4YzpXK0R/Twua4TB/QfE7e\niHMQMG/bVebF+fLVVH4o3qLjcyq62tnT/r5knciOHAKBUn4WAkCM00hYzhXmsXa2\n+GO+A9A++zBH65i03LeskfImR40Rrq6NRgTjbeiheQCb2JR4Twzb12Z28QqY/oRn\n+wIDAQAB\n-----END RSA PUBLIC KEY-----")
//     fmt.Println(string(privBytes))
//     fmt.Println(string(pubBytes))
//     privkey := BytesToPrivateKey(privBytes)
//     PublicKey := BytesToPublicKey(pubBytes)
//     fmt.Println(privkey)
//     fmt.Println(PublicKey)
//     msg := []byte("my test")
//     encryMsg := EncryptWithPublicKey(msg, PublicKey)
//     fmt.Println(encryMsg)
//     dycryMsg := DecryptWithPrivateKey(encryMsg, privkey)
//     fmt.Println(string(dycryMsg))
// }

// GenerateKeyPair generates a new key pair
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Println(err)
	}
	return privkey, &privkey.PublicKey
}

// PrivateKeyToBytes private key to bytes
func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		log.Println(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Println(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		log.Println(err)
	}
	return key
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
		if err != nil {
			log.Println(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		log.Println(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		log.Println("not ok")
	}
	return key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Println(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}
	return plaintext
}
