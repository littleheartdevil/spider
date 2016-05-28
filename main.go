package main

import (
    "fmt" 
    "net/http"
    "io/ioutil"
)

func main() {
    sayHi()
    txt, err := fetch("https://github.com")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(txt)
}

func sayHi() {
    fmt.Println("Spider say hi to everybody")
}

func fetch(urlAddr string) (string, error) {
    res, requestErr :=  http.Get(urlAddr)
    if requestErr == nil {
        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        return string(body), err 
    }
    return "", requestErr
}