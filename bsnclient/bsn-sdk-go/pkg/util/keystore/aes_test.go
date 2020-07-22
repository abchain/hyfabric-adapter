/**
 * @Author: Gao Chenxi
 * @Description: 
 * @File:  aes_test
 * @Version: 1.0.0
 * @Date: 2020/4/22 14:54
 */

package keystore

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAESCBCPKCS7Decrypt(t *testing.T) {

	data :=[]byte("abc")
	key :=[]byte("123456")

	//CBC模式，秘钥不足16位 PKCS7填充秘钥
	key = Pkcs7PaddingKey(key)
	//加密
	cr ,err :=AESCBCPKCS7Encrypt(key,data)
	if err !=nil{
		t.Fatal(err)
	}

	//转hex输出
	fmt.Println("Encrypt：",hex.EncodeToString(cr))

	//解密
	data,err = AESCBCPKCS7Decrypt(key,cr)
	if err !=nil{
		t.Fatal(err)
	}

	fmt.Println("Decrypt：",string(data))


}