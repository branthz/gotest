package main

import (
	"fmt"
	"package/tools"
    "encoding/hex"
    "encoding/base64"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"crypto"
	"encoding/pem"
)

var(
    aesiv="ca3911ea34bf9326f89ea2b68c0d6729"
    aeskey="74cab937aa5bbbb9702465b463635f11"
)

func testBase64() {
    data:="broadlink+broadlink"
    kk,_:=hex.DecodeString(aeskey)
    iv,_:=hex.DecodeString(aesiv)
    cipher,_:=tools.AesEncryptPkcs7([]byte(data),kk,iv)
    fmt.Printf("cipher:%v\n",cipher)
    bb:=base64.StdEncoding.EncodeToString(cipher)
    fmt.Printf("base64:%s\n",bb)
    bbp,_:=base64.StdEncoding.DecodeString(bb)
    fmt.Printf("base64 decode:%v\n",bbp)
    datap,_:=tools.AesDecryptPkcs7(bbp,kk,iv)
    fmt.Printf("aesdecode:%s\n",string(datap))
	//xxx := tools.Sha1Cal([]byte("broadlink"))
	//fmt.Printf("%s\n%d\n", xxx, len(xxx))
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)

func Sign(data []byte) (sig []byte, err error) {
	block, _ := pem.Decode(privateKey)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	
    hashFunc := crypto.MD5
    h := hashFunc.New()
    h.Write(data)
    digest := h.Sum(nil)
    return rsa.SignPKCS1v15(rand.Reader,priv, hashFunc, digest)
}

func checkSign(sign []byte,data []byte) error {
	block, _ := pem.Decode(publicKey)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
  	if err != nil {
    		return err
  	}
	pub := pubInterface.(*rsa.PublicKey)
	hashFunc := crypto.MD5
    	h := hashFunc.New()
    	h.Write(data)
    	digest := h.Sum(nil)

	return rsa.VerifyPKCS1v15(pub,hashFunc,digest,sign)
}


func rsaEncrypt(data []byte,pubkey []byte)([]byte,error){
	block, _ := pem.Decode(pubkey)
	if block==nil{
        	return nil,fmt.Errorf("pem decode publick key,get block nil")
    	}
			
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
        if err != nil {
                return nil,err
        }
        pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}


func main(){
	var data=[]byte{1,2,3,4,5,6,7,8,9,0,11,12,13,14,15,16,17}	
	var data = []byte{11,2,3,4,5,6,7,8,9,0,11,12,13,14,15,16,17}
	code,err:=rsaEncrypt(data,publicKey)	
	if err!=nil{
		fmt.Println(err)
		return	
	}
	fmt.Printf("sign code:%v,len:%d\n",code,len(code))	
	//err=checkSign(code,data)
	//fmt.Printf("checkSign return:%v\n",err)
}
