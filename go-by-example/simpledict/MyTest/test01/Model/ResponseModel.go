package Model

import (
	"encoding/json"
	"fmt"
	"log"
)

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func ToResponse(bodyText []byte) {
	var response DictResponse
	err := json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Fatal("ToResponse : ", err)
	}
	//fmt.Printf("%#v\n", response)
	fmt.Printf("\n")
	fmt.Println("[", response.Dictionary.Entry, "] UK:", response.Dictionary.Prons.En, "US:", response.Dictionary.Prons.EnUs)
	for _, item := range response.Dictionary.Explanations {
		fmt.Println(item)
	}
	fmt.Printf("\n")
}

func ToYouDaoResponse(bodyText []byte) {
	var response YouDaoResponse
	err := json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Fatal("ToYouDaoResponse : ", err)
	}
	//fmt.Printf("%#v\n", response)
	fmt.Printf("\n")
	fmt.Println("[", response.Ec.Word.ReturnPhrase, "] Prons : ", response.Ec.Word.Usphone)
	for _, item := range response.Ec.Word.Trs {
		fmt.Println(item)
	}
	for i, item := range response.BlngSentsPart.TrsClassify {
		if i == 0 {
			continue
		}
		fmt.Printf("%v ", item)
	}
	fmt.Printf("\n")
}
