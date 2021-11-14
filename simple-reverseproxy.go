package main

import (
    "net/http"
    "net/url"
    "reverse/rbpl"
    "flag"
    "fmt"
    "os"
)

func main() {
    flag.Parse()
    if len(flag.Arg(0)) == 0 {
	fmt.Println("command :PORT http(s)://hogehoge.com(:PORT)")
	os.Exit(0)
    } else if len(flag.Arg(1)) == 0 {
	fmt.Println("command :PORT http(s)://hogehoge.com(:PORT)")
	os.Exit(0)
    } else {
	fmt.Println(flag.Arg(0), " => ", flag.Arg(1))
        http.ListenAndServe(flag.Arg(0), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            path, err := url.Parse(flag.Arg(1))
            if err != nil {
                panic(err)
                return
            }
            proxy := reverseproxy.NewReverseProxy(path)
            proxy.ServeHTTP(w, r)
        }))
    }
}
