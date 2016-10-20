package main

import (
    "log"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("https://baidu.com")
    if err != nil {
        log.Fatalln(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(len(body))
}
