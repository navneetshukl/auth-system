package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string)(string,error){
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),10)
	if err!=nil{
		return "",err
	}

	return string(hashedPassword),nil
}

func ComparePassword(hashed,original string)(error){
	err:=bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(original))
	if err!=nil{
		return err
	}
	return nil
}