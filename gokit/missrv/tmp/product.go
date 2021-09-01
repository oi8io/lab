package tmp

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"time"
)

type Product struct {
	ID    int
	Title string
	Price float64
}

func RecProduct() Product {
	return Product{
		ID:    999,
		Title: "Golang 入门到精通",
		Price: 99.99,
	}
}

func GetProduct() (Product, error) {
	id := rand.Intn(10)
	if id < 6 {
		time.Sleep(2 * time.Second)
	}
	return Product{
		ID:    id,
		Title: "Golang 入门到精通",
		Price: 66.66,
	}, nil
}

type UserClaim struct {
	Uname string
	jwt.StandardClaims
}

func CreateToken() {
	sec := []byte("2020-07-20T14:28:56.963585Z")
	signing, err := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{Uname: "Jorden"}).SignedString(sec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signing)
	//getToken, err := jwt.Parse(signing, func(token *jwt.Token) (interface{}, error) {
	//	return sec, nil
	//})
	//if getToken.Valid {
	//	fmt.Println(getToken.Claims)
	//}
	u := &UserClaim{}
	getToken, err := jwt.ParseWithClaims(signing, u, func(token *jwt.Token) (interface{}, error) {
		return sec, nil
	})
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*UserClaim).Uname)
	}

}
func CreateTokenRS() {
	sec := []byte("2020-07-20T14:28:56.963585Z")
	signing, err := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{Uname: "Jorden"}).SignedString(sec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signing)
	//getToken, err := jwt.Parse(signing, func(token *jwt.Token) (interface{}, error) {
	//	return sec, nil
	//})
	//if getToken.Valid {
	//	fmt.Println(getToken.Claims)
	//}
	u := &UserClaim{}
	getToken, err := jwt.ParseWithClaims(signing, u, func(token *jwt.Token) (interface{}, error) {
		return sec, nil
	})
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*UserClaim).Uname)
	}

}
