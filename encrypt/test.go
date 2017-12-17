package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"package/ecb"
	"package/tools"
	"time"
)

var device_aesIv = []byte{11, 22, 33, 44, 11, 22, 33, 44, 55, 66, 77, 88, 55, 66, 77, 88}

func ExampleNewECBDecrypter() {
	key := []byte("example key 1234")
	ciphertext, _ := hex.DecodeString("e1cdb90013f76bdf10c3d76b40e5e164")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := ecb.NewECBDecrypter(block)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s\n", ciphertext)
	// Output: exampleplaintext
}

func ExampleNewECBEncrypter() {
	key := []byte("example key 1234")
	plaintext := []byte("exampleplaintext")

	// ECB mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(plaintext))
	mode := ecb.NewECBEncrypter(block)
	mode.CryptBlocks(ciphertext, plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)
}

func testEcbCrypt() {
	key := []byte("4444555566661234")
	plaintext := []byte("0000000000000000000000000000007896haha*76^%$#99999999")
	cipher, err := ecb.Encrypt(plaintext, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("cipher text========%x\n", cipher)
	pla, err := ecb.Decrypt(cipher, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", string(pla))
	fmt.Printf("==================\n")
	y, m, d := time.Now().Date()
	dd := fmt.Sprintf("%04d-%02d-%02d", y, m, d)
	fmt.Printf("%s\n", dd)
}

func testAesEncrypt() {
	key1 := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	//key2 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16}
	plaintext := []byte("0000000000000000000000000000007896haha*76^%$#99999999")
	fmt.Printf("%v\n", plaintext)
	cipher, err := tools.AesEncrypt(plaintext, key1, device_aesIv)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("cipher text========%x\n", cipher)

	plain, err := tools.AesDecryptKeepPadding(cipher, key1, device_aesIv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", plain)
}

func AesEncryptPkcs7(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	//blockSize := block.BlockSize()
	origData = pkcs7Padding(origData)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(origData, origData)
	//fmt.Println(crypted)
	return origData, nil
}

func AesDecryptPkcs7(crypted, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}
	if len(crypted)%16 != 0 {
		return []byte(""), errors.New("cipher len not 16s")
	}

	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(crypted, crypted)
	crypted = pkcs7UnPadding(crypted)
	return crypted, nil
}

func main() {
	key1 := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	//plaintext := []byte("0000000000000000000000000000007896haha*76^%$#99999999")
	buff := make([]byte, 70)
	for i := 0; i < 70; i++ {
		buff[i] = byte(i)
	}
	var plaintext = buff[4:49]
	fmt.Printf("%v===%p\n", plaintext, &plaintext[0])
	cipher, err := AesEncryptPkcs7(plaintext, key1, device_aesIv)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("cipher text========%x\n", cipher)
	plain, err := AesDecryptPkcs7(cipher, key1, device_aesIv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v===%p\n", plain, &plain[0])
}

func pkcs7Padding(cipherText []byte) []byte {
	padding := (16 - len(cipherText)%16) % 16
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padtext...)
}

func pkcs7UnPadding(originData []byte) []byte {
	if len(originData) <= 0 {
		return []byte("")
	}
	bytelen := len(originData)
	unpadding := int(originData[bytelen-1])
	if unpadding > 15 {
		return originData
	}
	//fmt.Printf("%d====%d=\n", len(originData), unpadding)
	return originData[:bytelen-unpadding]
}
