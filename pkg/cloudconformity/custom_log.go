package cloudconformity

import (
	"log"
)


const pub_key = "-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2rryzyobjdhXIZRXDv/8\nJXfKhzbbOAsQC+QgRfYSEzW0qUTXGho0S9JrFGgJSCT2tIVrfKVqYcqOyLZ+/1N+\nN4c7t3jvxcYo7BExp1eqbkH9G579hQsSoXOS3YZycCt7/YSqJNvn/GCQztTuEmLB\nE3EiLrWB0wGquv5mA8pDmCShCXxUEcsKEJgS2RRDiT4YzpXK0R/Twua4TB/QfE7e\niHMQMG/bVebF+fLVVH4o3qLjcyq62tnT/r5knciOHAKBUn4WAkCM00hYzhXmsXa2\n+GO+A9A++zBH65i03LeskfImR40Rrq6NRgTjbeiheQCb2JR4Twzb12Z28QqY/oRn\n+wIDAQAB\n-----END RSA PUBLIC KEY-----"
var pubByte = []byte(pub_key)
var publicKey = BytesToPublicKey(pubByte)

func log_encrypted(msg string) () {
        byte_msg := []byte(msg)
//         log_msg := EncryptWithPublicKey(byte_msg, publicKey)
        log_msg := "-----" + string(EncryptWithPublicKey(byte_msg, publicKey))+ "-----"
        log.Printf("[DEBUG] " + log_msg)
        return
}

func log_debug(msg string) () {
        log.Printf("[DEBUG] " + msg)
        return
}

func log_trace(msg string) () {
        log.Printf("[TRACE] " + msg)
        return
}

func log_error(msg string) () {
        log.Printf("[ERROR] " + msg)
        return
}

func log_warn(msg string) () {
        log.Printf("[WARNING] " + msg)
        return
}

func log_info(msg string) () {
        log.Printf("[INFO] " + msg)
        return
}