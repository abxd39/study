package encrypt_test

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/abxd39/myproject/encrypt"
)

/*
参数说明
key 的长度为 16 24 32
iv 长度为 必须为 16

*/

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
}

func ExampleAes_AesDecrypt() {
	aes := &encrypt.Aesu{}
	var buffer bytes.Buffer
	origData := `{\"token\":\"hgd5\",\"domain\":\"abc.com\"}`
	buffer.WriteString(origData)
	key := "github.com/abxd39/myproject/encr"
	iv := "github.com/abxd3"
	result, err := aes.AesEncrypt(buffer.Bytes(), key, iv)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(hex.EncodeToString(result))
	result, err = aes.AesDecrypt(result, key, iv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}

func TestAes_AesDecrypt(t *testing.T) {
	//http://tool.chacuo.net/cryptaes
	//http://blog.studygolang.com/2013/01/go加密解密之aes/
	//https://yq.aliyun.com/articles/711747?spm=a2c4e.11155472.0.0.6ed03abemsl9wt
/*
 iv=mgv3b132zX21K91F 加密的内容={"token":"mgv","domain":"http://campus-card-api.dev-ymj.snsshop.net"}
[T] 2019/09/05 11:05:05.105284 student-info.go:416: base64=e9RQLDhaPvDEc8hoU10LfaSS93vo2Pu4eDy6ggdGbfZheal2We3JcaQfkPxykFl+z0KbezBYVfxSk8BGf+MAG204o9eYwnNHShNUAd+84qo=
[T] 2019/09/05 11:05:05.105293 encrypt.go:86: 80 11
[T] 2019/09/05 11:05:05.105298 student-info.go:424: 解密结果为 {"token":"mgv","domain":"http://campus-card-api.dev-ymj.snsshop.net"}
{"appid":"eafgjswetr344njivlke","iv":"mgv3b132zX21K91F","str":""}
*/
	log.Println("test begin")
	aes := &encrypt.Aesu{}
	var buffer bytes.Buffer
	origData := "e393ff88-d14c-4827-8a1b-ad33624a824d"
	buffer.WriteString(origData)
	key := "dsfhet7346fgefsx"
	iv := "BJSH4K16Ji9d5KZT"
	t.Log(len(iv), iv)
	result, err := aes.AesEncrypt(buffer.Bytes(), key, iv)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println(base64.StdEncoding.EncodeToString(result))

	buf, err := base64.StdEncoding.DecodeString("fJANNe8ljHlQIRziAaZktdb0O+v/Hkdd61rsSqKNOJIP9i/p2yxjhRMIcY/guXJF")
	if err != nil {
		log.Println(err)
		return
	}

	result, err = aes.AesDecrypt(buf, key, iv)

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(result))

	//gocrypto.SetAesKey("dsfhet7346fgefsx")
	//enb, err := gocrypto.AesCBCEncrypt([]byte(origData))
	//if err != nil {
	//	t.Fatal(err)
	//	return
	//}
	//t.Log(string(enb))
	//res, err := gocrypto.AesCBCDecrypt(enb)
	//if err != nil {
	//	t.Fatal(err)
	//	return
	//}
	//t.Log(base64.StdEncoding.EncodeToString(res))

}
