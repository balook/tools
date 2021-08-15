package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/twmb/murmur3"
)

func main() {

	stat, _ := os.Stdin.Stat()

	// if (stat.Mode() & os.ModeCharDevice) == 0 {
	// 	fmt.Println("data is being piped to stdin")
	// } else {
	// 	fmt.Println("stdin is from a terminal")
	// }

	if len(os.Args) == 2 {
		data := os.Args[1]
		request(data)
	} else if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			request(scanner.Text())
		}
	} else {
		os.Exit(0)
	}
}

func request(input string) {
	response, err := http.Get(input)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
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

	if err != nil {
		log.Fatal("error")
	}

	result := murmur3.Sum32([]byte(str))

	fmt.Println(input, "->", result)

}
