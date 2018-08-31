package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
  "reflect"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello golang server"))
}

func test(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("test golang server"))
}

func main() {
    //将func (w http.ResponseWriter, r *http.Request)类型转为http.HandlerFunc类型，为其增加了ServeHTTP方法，支持了http.Handler接口 
    th := http.HandlerFunc(timeHandler)
    fmt.Println(reflect.TypeOf(th))

    http.Handle("/time", th)
    http.Handle("/", http.HandlerFunc(home))
    
    //http路由可以通过http.Handle指定，也可以通过http.HandlFunc指定
    http.HandleFunc("/test",test)

    log.Println("Listening...")
    // And pass nil as the handler to ListenAndServe.
    http.ListenAndServe(":3000", nil)
}
