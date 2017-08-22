package main

// 1. 해시 알고리즘

// 2. 대칭키 알고리즘

// 3. 공개키 알고리즘

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
)

func main() {

	s := "Hello World?"
	h1 := sha512.Sum512([]byte(s)) // 문자열 sha512 해시값

	sha := sha512.New() // sha512 해시값 생성
	sha.Write([]byte("Hello, "))
	sha.Write([]byte("world?"))
	h2 := sha.Sum(nil) // sha 해시값 추출

	fmt.Println(h1, h2)

	key := "AES KEY"

	block, err := aes.NewCipher([]byte(key)) // AES 대칭키 암호화블록
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s)) // 평문을 알고리즘으로 암호화 했다.

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext) // 암호화된 데이터를 평문으로 복호화

	fmt.Println(plaintext)

	// 키와 암호화할 데이터의 크기가 일정해야 함.
	// 긴 데이터는 잘라서 암호화해야하는데 그럴 시 보안에 취약(ECB방식)

	//그래서 나온건 CBC방식 crypto/rand 와 crypto/cipher crypto/aes 이용

	// rsa 공개키 방식 : 좀 더 안전하다

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey // rsa로 생성시 개인 키, 공개 키가 생성되고 개인 키 안에 공개 키가 들어 있음.
	s2 := `
	아이 졸령
	`
	ciphertext2, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s2),
	) // 암호화

	plaintext2, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext2, // 개인 키를 사용해 다시 복호화했음
	)

	fmt.Println(plaintext2)
}
