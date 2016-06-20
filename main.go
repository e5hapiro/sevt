package main

import (
	"fmt"
	"github.com/CiscoZeus/go-zeusclient"
	"io"
	"net/http"
	"os"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>This is a house of learned doctors.</h1>")
	io.WriteString(w, "<a href=\"https://www.youtube.com/watch?v=hh1oaumUoyc\">https://www.youtube.com/watch?v=hh1oaumUoyc</a>")
//	fmt.Fprintf(w, "<h2>Hi there, I love %s!</h2>", r.URL.Path[1:])
	fmt.Fprintf(w, "<h1>Its Summer, time for cold %s!</h1>", r.URL.Path[1:])
}

func main() {
	// print env variables
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("USER:", os.Getenv("USER"))
	fmt.Println("HOME:", os.Getenv("HOME"))
	fmt.Println("HOSTNAME:", os.Getenv("HOSTNAME"))

	// send env variables to zeus
	z := &zeus.Zeus{ApiServ: "http://api.ciscozeus.io", Token: os.Getenv("ZEUS_TOKEN")}
	logs := zeus.LogList{
		Name: "syslog",
		Logs: []zeus.Log{
			zeus.Log{
				"OS":       runtime.GOOS,
				"USER":     os.Getenv("USER"),
				"HOME":     os.Getenv("HOME"),
				"HOSTNAME": os.Getenv("HOSTNAME")},
		},
	}
	suc, err := z.PostLogs(logs)
	if err != nil {
		fmt.Println("Zeus error response:", err.Error())
	} else {
		fmt.Println("Zeus success response:", suc)
	}

	fmt.Printf("Starting http server...\n")
	http.HandleFunc("/", handler)
	err2 := http.ListenAndServe(":8080", nil)
	fmt.Println(err2.Error())
}
