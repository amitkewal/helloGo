package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	"fmt"
    // "html/template"
    // "log"
    "net/http"
    //"strings"
)

func main() {
	beego.Run()
}

func submit(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
         r.ParseForm()
        // logic part of log in
        fmt.Println("Name:", r.Form["fullName"])
        fmt.Println("Email:", r.Form["email"])
        fmt.Println("Country:", r.Form["country"])
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

