/*
============= Usage ===========
urls.txt => list of favicon.txt urls
cat urls.txt | go run favicon.

Run :-
	go get github.com/spaolacci/murmur3

credits @sw33tlie
my twitter => @prob0x01
==============================
*/

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/twmb/murmur3"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		data := sc.Text()
		fmt.Println(data, " -> ", Request(data))
	}
}

func Request(data string) uint32 {
	response, _ := http.Get(data)
	body, _ := ioutil.ReadAll(response.Body)
	str := base64.StdEncoding.EncodeToString(body)

	final := ""
	fix := 76
	s := make([]string, 0)

	for i := 0; i*fix+fix < len(str); i++ {
		it := str[i*fix : i*fix+fix]
		s = append(s, it)
	}

	findlen := len(s) * fix
	last := str[findlen:] + "\n"

	for _, s := range s {
		final = final + s + "\n"
	}

	str = final + last

	return murmur3.Sum32([]byte(str))
}
