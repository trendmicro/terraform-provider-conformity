import (
	"log"
)


const pubByte = []byte("-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2rryzyobjdhXIZRXDv/8\nJXfKhzbbOAsQC+QgRfYSEzW0qUTXGho0S9JrFGgJSCT2tIVrfKVqYcqOyLZ+/1N+\nN4c7t3jvxcYo7BExp1eqbkH9G579hQsSoXOS3YZycCt7/YSqJNvn/GCQztTuEmLB\nE3EiLrWB0wGquv5mA8pDmCShCXxUEcsKEJgS2RRDiT4YzpXK0R/Twua4TB/QfE7e\niHMQMG/bVebF+fLVVH4o3qLjcyq62tnT/r5knciOHAKBUn4WAkCM00hYzhXmsXa2\n+GO+A9A++zBH65i03LeskfImR40Rrq6NRgTjbeiheQCb2JR4Twzb12Z28QqY/oRn\n+wIDAQAB\n-----END RSA PUBLIC KEY-----")
const publicKey =: BytesToPublicKey(pubByte)

func custom_log_print(msg string, is_encrypt bool, mode string) () {
    if is_encrypt {
        msg = "-----"EncryptWithPublicKey(payload, publicKey)+ "-----"
        log.Printf(mode + " " + msg)
    }
    else{
        log.Printf(mode + " " + msg)
    }

}