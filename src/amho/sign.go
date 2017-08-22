package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func sign() {

	// 서명과 인증

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048) // 개인키와 공캐기 만듦
	publicKey := &privateKey.PublicKey

	message := "나니?"
	hash := md5.New()
	hash.Write([]byte(message)) // 해시값을 만들고 문자열을 추가했음
	digest := hash.Sum(nil)     // 해시값 추출

	var h1 crypto.Hash
	signature, _ := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		h1,
		digest, // 개인 키로 서명
	)

	var h2 crypto.Hash
	err := rsa.VerifyPKCS1v15(
		publicKey,
		h2,
		digest,
		signature, // 공개키로 서명검증
	)

	if err != nil {
		fmt.Println("실패")
	}

	// 메시지, 메시기 해시값, 서명, 공개키는 공개
	// 개인키만 숨겨져 있음
	// 메시지 해시값으로 메시지 변조 여부를 알 수 있음
	// 서명 값으로 메시지가 올바른 사람에게서 왔는지 알 수 있음

	// 예 ) 공인인증서

}
