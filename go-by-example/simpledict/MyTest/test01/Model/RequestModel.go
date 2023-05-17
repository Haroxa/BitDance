package Model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"` // 字段名必须正确
	UserID    string `json:"user_id"`
}

func ToIOReader(word string) io.Reader {
	//fmt.Printf("Your word is [%s]!\n", word)
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal("ToIOReader : ", err)
	}
	var data = bytes.NewReader(buf)
	//fmt.Printf("ToIOReader : data[%s] !\n", data)
	return data
}

const (
	YouDaoRow string = "le=en&t=9&client=web&sign=4ad1707568da35b30d1d4ce1700f35ef&keyfrom=webdict"
)

func ToYouDaoReader(word string) io.Reader {
	//fmt.Printf("Your word is [%s]!\n", word)
	request := fmt.Sprintf("q=%s&%s", word, YouDaoRow)
	buf := []byte(request)
	var data = bytes.NewReader(buf)
	//fmt.Printf("ToIOReader : data[%s] !\n", data)
	return data
}
