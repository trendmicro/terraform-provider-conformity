package main

import (
	"flag"
	"fmt"
    "io/ioutil"
    "os"
    "log"
    "crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
    "encoding/base64"
    "strings"
    "regexp"
)

// GenerateKeyPair generates a new key pair
func CheckError(e error) {
    if e != nil {
        fmt.Println(e.Error)
    }
}

func CreateFile(filename string, data string) {

    file, err := os.Create(filename)

    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }

    // Defer is used for purposes of cleanup like
    // closing a running file after the file has
    // been written and main //function has
    // completed execution
    defer file.Close()

    // len variable captures the length
    // of the string written to the file.
    len, err := file.WriteString(data)

    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
    fmt.Printf("\nFile Name: %s", file.Name())
    fmt.Printf("\nLength of data: %s", len)
}

func read_file(file_path string) string{

    fmt.Println("\n\nReading a file in Go lang ")
    fmt.Println("\nFile Name: ", file_path)

    // The ioutil package contains inbuilt
    // methods like ReadFile that reads the
    // filename and returns the contents.
    data, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println("failed reading data from file: %s", err)
    }
    return string(data)

}

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

func DecryptWithPrivateKey(cipherText string, privKey *rsa.PrivateKey) string {
    ct, _ := base64.StdEncoding.DecodeString(cipherText)
    label := []byte("OAEP Encrypted")
    rng := rand.Reader
    plaintext, err := rsa.DecryptOAEP(sha512.New(), rng, privKey, ct, label)
    CheckError(err)
    return string(plaintext)
}

func replace_encrypted(log_data string, priv_key string) string{
    //fmt.Printf(*privateKeyPath)
	privkey := BytesToPrivateKey([]byte(priv_key))
    r := regexp.MustCompile(`-----\s*(.*?)\s*-----`)
    matches := r.FindAllStringSubmatch(log_data, -1)
    var dycryMsg string
    for _, v := range matches {
//         fmt.Println(v[1])
        dycryMsg = DecryptWithPrivateKey(v[1], privkey)
        log_data = strings.Replace(log_data, v[1], dycryMsg, -1)
    }

    return log_data
}

func main() {

	privateKeyPath := flag.String("key-path", "", "Path to the private key")
	logFilePath := flag.String("log-file-path", "", "Path to the log file")
	flag.Parse()

	priv_key := read_file(*privateKeyPath)

	flag.Parse()

	log_data := read_file(*logFilePath)

    dycryMsg := replace_encrypted(log_data, priv_key)

    split_file_name := strings.Split(*logFilePath, ".")
    output_filename := split_file_name[0] + "_decrypted." + split_file_name[1]
    CreateFile(output_filename, dycryMsg)
}
