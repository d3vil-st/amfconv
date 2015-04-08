package main

import (
	"./src/encoding/amf"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Print(hex.Dump(amf.EncodeAMF0(1.0)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(1)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(-1)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(true)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(false)))
	fmt.Print(hex.Dump(amf.EncodeAMF0("foo")))
	fmt.Print(hex.Dump(amf.EncodeAMF0("")))
	fmt.Print(hex.Dump(amf.EncodeAMF0(nil)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(amf.Amf0Reference(5))))
	assoc := map[string]interface{}{"1": 1,
		"2": 1.797693134862315708145274237317043567981e+308,
		"3": "three",
		"4": nil,
		"5": true}
	//assoc := map[string]interface{}{"1": "2", "3": "4", "5": "6", "7": "8"}
	fmt.Print(hex.Dump(amf.EncodeAMF0(assoc)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(amf.Amf0ECMAArray(assoc))))
	fmt.Print(hex.Dump(amf.EncodeAMF0(amf.Amf0Date(353464561))))
	fmt.Print(hex.Dump(amf.EncodeAMF0([]interface{}{"one", "two", "three"})))
	fmt.Print(hex.Dump(amf.EncodeAMF0([]interface{}{1, 2, 1})))
	fmt.Print(hex.Dump(amf.EncodeAMF0([]interface{}{true, false, true})))
	fmt.Print(hex.Dump(amf.EncodeAMF0(amf.Amf0Xml("one two three"))))
}
