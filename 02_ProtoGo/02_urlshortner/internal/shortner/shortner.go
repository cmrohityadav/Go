package shortner

import (
	"crypto/rand"
	"encoding/hex"
)


func GenerateString()(string,error){
	byte:=make([]byte,4);

	_,err:=rand.Read(byte);
	if err!=nil{
		return "",err;
	}

	return hex.EncodeToString(byte),nil
}