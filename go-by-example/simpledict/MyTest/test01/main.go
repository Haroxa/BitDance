package main

import "test01/Query"

func main() {

	word := Query.ScanWord()

	go Query.Query(word, Query.CaiYun)
	Query.Query(word, Query.YouDao)

}
