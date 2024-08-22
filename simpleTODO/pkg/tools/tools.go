package tools

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
)

//other tools
func GenerateUUID() string {
	bytes := make([]byte, 33)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}
	return base64.RawStdEncoding.EncodeToString(bytes)

}

func HashThis(originalText string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(originalText))
	hashed_byte := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hashed_byte)
	return hashedPassword
}

func ValidateFormInputs(tobevalidated string, valuetobevalidated string) bool {
	validationFlag := true

	if tobevalidated == "username" {
		if valuetobevalidated == "" {
			fmt.Println("Empty username")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 5 || len(valuetobevalidated) > 30 {
				fmt.Println("invalid length for username")
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(valuetobevalidated)) {
					fmt.Println("username must only contain alpha numeric characters")
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "password" {
		if tobevalidated == "" {
			fmt.Println("Empty Password")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 5 {
				fmt.Println("invalid length for username")
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`[a-zA-Z]`).MatchString(valuetobevalidated) && regexp.MustCompile(`\d`).MatchString(valuetobevalidated) && regexp.MustCompile(`[\W_]`).MatchString(valuetobevalidated)) {
					fmt.Println("Password must be alpha numeric and at least a symbol character")
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "firstname" {
		if valuetobevalidated == "" {
			fmt.Println("Empty First name")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				fmt.Println("invalid length for First name")
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(valuetobevalidated)) {
					fmt.Println("firstname must only contain alpha numeric characters")
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "lastname" {
		if valuetobevalidated == "" {
			fmt.Println("Empty Last name")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) < 3 || len(valuetobevalidated) > 20 {
				fmt.Println("invalid length for First name")
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(valuetobevalidated)) {
					fmt.Println("firstname must only contain alpha numeric characters")
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "email" {
		if valuetobevalidated == "" {
			fmt.Println("Empty email")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) > 40 {
				fmt.Println("invalid length for email")
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(valuetobevalidated)) {
					fmt.Println("invalid email")
					validationFlag = false
					return validationFlag
				}
			}
		}
	} else if tobevalidated == "phonenumber" {
		if valuetobevalidated == "" {
			fmt.Println("Empty phone number")
			validationFlag = false
			return validationFlag
		} else {
			if len(valuetobevalidated) != 10 {
				fmt.Println("Invalid length for phone number")
				validationFlag = false
				return validationFlag
			} else {
				if !(regexp.MustCompile(`(^\d+$)`).MatchString(valuetobevalidated)) {
					fmt.Println("Phone number must be only numeric")
					validationFlag = false
					return validationFlag
				}
			}
		}
	}

	return validationFlag
}
