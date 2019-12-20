package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main(){
	token := "test"
	//username := "ibc_chengxiang"
	usersecret := "ea6b2efbdd4255a9f1b3bbc6399b58f4"
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("%d%s%s%d", timestamp, token, usersecret, timestamp)
	fmt.Println(key)
	key="1553826571testea6b2efbdd4255a9f1b3bbc6399b58f41553826571"
	fmt.Println(key)
	secret := md5.Sum([]byte(key))
	signature := hex.EncodeToString(secret[:])
	fmt.Println(timestamp)
	fmt.Println(signature)
}
