package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/go-redis/redis/v7"
    "encoding/json"
    //"github.com/garyburd/redigo/redis"

)
var count = 0
type Author struct {
   Count int `json:"count"`
   }
func main() {

         client := redis.NewClient(&redis.Options{
                 Addr: "172.17.0.1:6379",
                Password: "",
                DB: 0,
    })
     http.HandleFunc("/hits", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("counter: %d\n", count)
        count = count +1
        json, err := json.Marshal(Author{Count: count})
        //json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
    if err != nil {
        fmt.Println(err)
    }
        //err = client.Set("id1234", json, 0).Err()
        err = client.Set("counter", json, 0).Err()
     if err != nil {
        fmt.Println(err)
    }

        //fmt.Fprintf("%+v\n", count)
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request){
       val, err := client.Get("counter").Result()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Fprintf(w, val)
    //fmt.Sprintf("t1: %s\n", string(val))
            //fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
