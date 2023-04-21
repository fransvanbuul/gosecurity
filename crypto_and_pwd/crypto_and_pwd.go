package main

import (
    "crypto/aes"
    "crypto/rand"
    "fmt"
    "io"
)

func main() {
    // [Vulnerability] Hardcoded password
    password := []byte("Alice and Bob........")
    encrypted := make([]byte, len(password))
    plaintext2 := make([]byte, len(password))
    key := make([]byte, 16)
    _, err := io.ReadFull(rand.Reader, key)
    if err != nil { panic(err.Error()) }

    // [Vulnerability] Weak Encryption: Insufficient Key Size
    _, err = aes.NewCipher([]byte("Not sixteen b"))
    if err != nil { panic(err.Error()) }

    // [Vulnerability] Key Management: Hardcoded Encryption Key
    _, err = aes.NewCipher([]byte("Sixteen Byte Key"))
    if err != nil { panic(err.Error()) }

    // [Vulnerability] Key Management: Null Encryption Key
    _, err = aes.NewCipher(nil)
    if err != nil { panic(err.Error()) }

    // [Vulnerability] Key Management: Empty Encryption Key
    _, err = aes.NewCipher([]byte(""))
    if err != nil { panic(err.Error()) }

    // [Safe] Weak Encryption: Hardcoded Key Size (128 bits, sufficient)
    aes_blk, _ := aes.NewCipher(key)

    aes_blk.Encrypt(encrypted, password)
    aes_blk.Decrypt(plaintext2, encrypted)

    // [Vulnerability] Privacy Violation (the password is being sent to the console)
    fmt.Printf("%s\n", plaintext2)
}
