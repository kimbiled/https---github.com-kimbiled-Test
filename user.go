package main

import
(
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/solite"
)

type User struct
{
	gorm.Model

	
}

func AllUsers(w http.ResponseWriter, r *http.Request)
{
	fmt.Fprintf(w,"All users list")
}
func UpdateUser(w http.ResponseWriter, r *http.Request)
{
	fmt.Fprintf(w,"Update users list")
}

func CreateUser(w http.ResponseWriter, r *http.Request)
{
	fmt.Fprintf(w,"Created users list")
}
func DeleteUser(w http.ResponseWriter, r *http.Request)
{
	fmt.Fprintf(w,"Delete users list")
}

