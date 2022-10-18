//// Fetchall fetches URLs in parallel and reports their times and sizes.
//package main
//
//import (
//	"fmt"
//	"io"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//	"time"
//)
//
//func main() {
//	start := time.Now()
//	ch := make(chan string)
//	//用make函数创建了一个传递string类型参数的channel
//	for _, url := range os.Args[1:] {
//		go fetch(url, ch) // start a goroutine
//	}
//	for range os.Args[1:] {
//		fmt.Println(<-ch) // receive from channel ch
//	}
//	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
//}
//
//func fetch(url string, ch chan<- string) {
//	start := time.Now()
//	resp, err := http.Get(url)
//	if err != nil {
//		ch <- fmt.Sprint(err) // send to channel ch
//		return
//	}
//	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
//	resp.Body.Close() // don't leak resources
//	if err != nil {
//		ch <- fmt.Sprintf("while reading %s: %v", url, err)
//		return
//	}
//	secs := time.Since(start).Seconds()
//	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
//}
//

//// Server1 is a minimal "echo" server.
//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//func main() {
//	http.HandleFunc("/", handler) // each request calls handler
//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
//}
//
//// handler echoes the Path component of the request URL r.
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//}
//package main
//
//func add1(r rune) rune {
//	return r + 1
//}
//
//fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
//fmt.Println(strings.Map(add1, "VMS")) // "WNT"
//fmt.Println(strings.Map(add1, "Admix")) // "Benjy"

package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name

}
func main() {
	fmt.Println(Hello("world"))

}
