package main

import (
	"encoding/hex"
	"fmt"
	"github.com/speps/amfconv/src/encoding/amf"
	"time"
)

func main() {
	assoc := map[string]interface{}{"1": 1,
		"2": 1.797693134862315708145274237317043567981e+308,
		"3": "three",
		"4": nil,
		"5": true}

	fmt.Print(hex.Dump(amf.EncodeAMF0(1.3)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(1)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(-1)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(true)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(false)))
	fmt.Print(hex.Dump(amf.EncodeAMF0("foo")))
	fmt.Print(hex.Dump(amf.EncodeAMF0("")))
	fmt.Print(hex.Dump(amf.EncodeAMF0(nil)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(amf.Amf0Reference(5))))
	fmt.Print(hex.Dump(amf.EncodeAMF0(assoc)))
	fmt.Print(hex.Dump(amf.EncodeAMF0(amf.Amf0ECMAArray(assoc))))
	fmt.Print(hex.Dump(amf.EncodeAMF0(time.Now())))
	fmt.Print(hex.Dump(amf.EncodeAMF0([]interface{}{"one", "two", "three"})))
	fmt.Print(hex.Dump(amf.EncodeAMF0([]interface{}{1, 2, 1})))
	fmt.Print(hex.Dump(amf.EncodeAMF0([]interface{}{true, false, true})))
	fmt.Print(hex.Dump(amf.EncodeAMF0(amf.Amf0Xml("one two three"))))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(1.3)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(1)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(-1)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(true)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(false)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0("foo")))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0("")))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(assoc)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(nil)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(amf.Amf0Reference(5))))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(assoc)))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(amf.Amf0ECMAArray(assoc))))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(time.Now())))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0([]interface{}{"one", "two", "three"})))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0([]interface{}{1, 2})))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0([]interface{}{true, false, true})))
	fmt.Println(amf.DecodeAMF0(amf.EncodeAMF0(amf.Amf0Xml("one two three"))))
	fmt.Print(hex.Dump(amf.EncodeAMF3(1.3)))
	fmt.Print(hex.Dump(amf.EncodeAMF3(true)))
	fmt.Print(hex.Dump(amf.EncodeAMF3(false)))
	fmt.Print(hex.Dump(amf.EncodeAMF3(nil)))
}
