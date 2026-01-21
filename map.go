package main

import (
	"fmt"
)

func Map() {
	currnecyCode := make(map[string]string)

	currnecyCode["USD"] = "US Dollar"
	currnecyCode["EUR"] = "Euro"
	currnecyCode["INR"] = "Indian Rupee"

	fmt.Println("\n", currnecyCode)

	newMap := map[string]string{
		"HE": "Homo Erectus",
		"HS": "Homo Sapiens",
		"HN": "Homo Neanderthals",
	}

	fmt.Println(newMap)
}
