# AMF encoding/decoding library for Go

The Adobe Integrated Runtime and Adobe Flash Player use AMF to communicate between an application and a remote server. AMF encodes remote procedure calls (RPC) into a compact binary representation that can be transferred over HTTP/HTTPS or the RTMP/RTMPS protocol. Objects and data values are serialized into this binary format, which increases performance, allowing applications to load data up to 10 times faster than with text-based formats such as XML or SOAP.

## AMF0

 - [x] `int` / Number
 - [x] `float64` / Number
 - [x] `bool` / Boolean
 - [x] `string` / String
 - [x] `map[string]interface{}` / Object
 - [x] `nil` / Null
 - [x] `[]interface{}` / Array
 - [x] `time.Time` / Date

## AMF3

 - [x] `int`, `uint` / Number
 - [x] `float64` / Number
 - [x] `bool` / Boolean
 - [x] `string` / String
 - [x] `nil` / Null
 - [x] `time.Time` / Date

Work in progress
