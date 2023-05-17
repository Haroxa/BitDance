package Query

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"test01/Model"
	"test01/UrlApi"
)

var QueryWay = []string{
	"CaiYun", "YouDao",
}

func Select() int {
	fmt.Printf("Please select way to query: \n")
	for k, v := range QueryWay {
		fmt.Printf("%d-%s\n", k+1, v)
	}
	var option int
	fmt.Scanf("%d", &option)
	if option < 1 && option > len(QueryWay) {
		log.Fatal("Unknown option ??? ")
	}
	return option
}

func ScanWord_1() string {
	fmt.Printf("Please input a word to query: ")
	var word string
	fmt.Scanf("%s\n", &word)
	fmt.Printf("Your input is [%s]!\n", word)
	return word
}

func ScanWord() string {
	fmt.Printf("Please input a word/words to query: ")
	var word string
	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	word = strings.TrimSuffix(word, "\r\n")
	//fmt.Scanf("%s\n", &word)
	fmt.Printf("Your input is [%s]!\n", word)
	return word
}

const (
	CaiYun = iota + 1
	YouDao
)

func Query(word string, option int) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"go"}`)
	//fmt.Printf("ToIOReader : data[%s] !\n", data)
	//word := ScanWord()
	//option := Select()

	//fmt.Printf("ToIOReader : data[%s] !\n", data)
	var data io.Reader
	var req *http.Request
	switch option {
	case CaiYun:
		data = Model.ToIOReader(word)
		req = UrlApi.ToHttpRequest(data)
	case YouDao:
		data = Model.ToYouDaoReader(word)
		req = UrlApi.ToYoudaoHttpRequest(data)
	default:
		log.Fatal("Unknown option ??? ")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	switch option {
	case CaiYun:
		Model.ToResponse(bodyText)
	case YouDao:
		Model.ToYouDaoResponse(bodyText)
	default:
		log.Fatal("Unknown case ??? ")
	}
}
