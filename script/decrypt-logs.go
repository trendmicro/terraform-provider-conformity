package main

import (
	"flag"
)

func main() {

	privateKeyPath := flag.String("key-path", "", "Path to the private key")
	logFilePath := flag.String("log-file-path", "", "Path to the log file")

	flag.Parse()

	print(*logFilePath)
	print(*privateKeyPath)
}
