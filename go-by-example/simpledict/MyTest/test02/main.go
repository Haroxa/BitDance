package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"../test02/Model"
)

func main() {
	client := &http.Client{}
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"go"}`)
	//fmt.Printf("ToIOReader : data[%s] !\n", data)
	word := ScanWord()
	data := ToIOReader(word)
	//fmt.Printf("ToIOReader : data[%s] !\n", data)

	req := ToHttpRequest(data)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func ScanWord() string {
	fmt.Printf("Please input a word to query: ")
	var word string
	fmt.Scanf("%s\n", &word)
	fmt.Printf("Your input is [%s]!\n", word)
	return word
}

func ToIOReader(word string) *bytes.Reader {
	//fmt.Printf("Your word is [%s]!\n", word)
	request := Model.DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	fmt.Printf("ToIOReader : data[%s] !\n", data)
	return data
}

func ToHttpRequest(data io.Reader) *http.Request {
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "664ab23505777488c3ab7f71c7c0b958")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.0.0")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	return req
}
