package amf

import (
	"bytes"
	"encoding/binary"
	"math"
	"sort"
)

type AMFVersion uint8

const (
	AMF0 AMFVersion = 0x1
	AMF3 AMFVersion = 0x3
)

type amf0 uint8
type Amf0Reference uint16
type Amf0Date float64
type Amf0ECMAArray map[string]interface{}
type Amf0Xml string
type amf3 uint8

const (
	amf0Number    = amf0(0x00) // done
	amf0Boolean   = amf0(0x01) // done
	amf0String    = amf0(0x02) // done
	amf0Object    = amf0(0x03) // done
	amf0Null      = amf0(0x05) // done
	amf0Undefined = amf0(0x06) // done
	amf0Reference = amf0(0x07) // done
	amf0Array     = amf0(0x08) // done
	amf0ObjectEnd = amf0(0x09) // done
	amf0StrictArr = amf0(0x0a) // done
	amf0Date      = amf0(0x0b) // done
	amf0StringExt = amf0(0x0c) // done
	amf0Xml       = amf0(0x0f) // done
	amf0Instance  = amf0(0x10)
)

const (
	amf3Undefined = amf3(0x00)
	amf3Null      = amf3(0x01)
	amf3False     = amf3(0x02)
	amf3True      = amf3(0x03)
	amf3Integer   = amf3(0x04)
	amf3Double    = amf3(0x05)
	amf3String    = amf3(0x06)
	amf3XmlDoc    = amf3(0x07)
	amf3Date      = amf3(0x08)
	amf3Array     = amf3(0x09)
	amf3Object    = amf3(0x0a)
	amf3Xml       = amf3(0x0b)
	amf3ByteArray = amf3(0x0c)
)

func EncodeAMF0(v interface{}) []byte {
	var msg []byte
	switch v.(type) {
	case float64:
		msg = encodeNumber(v.(float64))
	case int:
		msg = encodeNumber(float64(v.(int)))
	case bool:
		msg = encodeBoolean(v.(bool))
	case string:
		msg = encodeString(v.(string))
	case nil:
		msg = encodeNull()
	case Amf0Reference:
		msg = encodeReference(v.(Amf0Reference))
	case map[string]interface{}:
		msg = encodeObject(v.(map[string]interface{}))
	case Amf0ECMAArray:
		msg = encodeECMAArray(v.(Amf0ECMAArray))
	case Amf0Date:
		msg = encodeDate(v.(Amf0Date))
	case []interface{}:
		msg = encodeStrictArr(v.([]interface{}))
	case Amf0Xml:
		msg = encodeXml(v.(Amf0Xml))
	}
	return msg
}

func encodeNumber(v float64) []byte {
	msg := make([]byte, 1+8) // 1 header + 8 float64
	msg[0] = byte(amf0Number)
	binary.BigEndian.PutUint64(msg[1:], uint64(math.Float64bits(v)))
	return msg
}

func encodeBoolean(v bool) []byte {
	msg := make([]byte, 1+1) // 1 header + 1 boolean
	msg[0] = byte(amf0Boolean)
	if v {
		msg[1] = 0x1
	} else {
		msg[1] = 0x0
	}
	return msg
}

func encodeString(v string) []byte {
	var msg []byte
	if len(v) < 0xffff {
		msg = make([]byte, 1+2+len(v)) // 1 header + 2 length + length of string
		msg[0] = byte(amf0String)
		binary.BigEndian.PutUint16(msg[1:], uint16(len(v)))
		copy(msg[3:], v)
	} else {
		msg = make([]byte, 1+4+len(v)) // 1 header + 4 length + length of string
		msg[0] = byte(amf0StringExt)
		binary.BigEndian.PutUint32(msg[1:], uint32(len(v)))
		copy(msg[5:], v)
	}
	return msg
}

func encodeObject(v map[string]interface{}) []byte {
	var keys []string
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	buf := new(bytes.Buffer)
	for key := range keys {
		buf.Write(EncodeAMF0(keys[key]))
		switch v[keys[key]].(type) {
		case int:
			buf.Write(EncodeAMF0(v[keys[key]].(int)))
		case float64:
			buf.Write(EncodeAMF0(v[keys[key]].(float64)))
		case string:
			buf.Write(EncodeAMF0(v[keys[key]].(string)))
		case bool:
			buf.Write(EncodeAMF0(v[keys[key]].(bool)))
		case Amf0Date:
			buf.Write(EncodeAMF0(v[keys[key]].(Amf0Date)))
		case Amf0Reference:
			buf.Write(EncodeAMF0(v[keys[key]].(Amf0Reference)))
		case nil:
			buf.Write(EncodeAMF0(nil))
		}
	}
	buf.Write(encodeObjectEnd())
	msg := make([]byte, 1+buf.Len()) // 1 header + length
	msg[0] = byte(amf0Object)
	copy(msg[1:], buf.Bytes())
	return msg
}

func encodeNull() []byte {
	msg := make([]byte, 1)
	msg[0] = byte(amf0Null)
	return msg
}

func encodeReference(v Amf0Reference) []byte {
	msg := make([]byte, 1+2) // 1 header + 2 uint16
	msg[0] = byte(amf0Reference)
	binary.BigEndian.PutUint16(msg[1:], uint16(v))
	return msg
}

/*func encodeUndefined() []byte {
    msg := make([]byte, 1)
    msg[0] = byte(amf0Undefined)
    return msg
}*/

func encodeECMAArray(v Amf0ECMAArray) []byte {
	msg_body := encodeObject(v)
	summary_length := len(msg_body) - 4
	msg := make([]byte, 1+4+summary_length) // 1 header + 4 length + sum length
	msg[0] = byte(amf0Array)
	binary.BigEndian.PutUint32(msg[1:], uint32(len(v)))
	copy(msg[5:], msg_body[1:summary_length+1])
	return msg
}

func encodeObjectEnd() []byte {
	msg := make([]byte, 3)
	msg[0], msg[1], msg[2] = 0x0, 0x0, byte(amf0ObjectEnd)
	return msg
}

func encodeDate(v Amf0Date) []byte {
	msg := make([]byte, 1+8+2) // 1 header + 8 float64 + 2 timezone
	msg[0], msg[9], msg[10] = byte(amf0Date), 0x0, 0x0
	binary.BigEndian.PutUint64(msg[1:], uint64(math.Float64bits(float64(v))))
	return msg
}

func encodeStrictArr(v []interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, k := range v {
		buf.Write(EncodeAMF0(k))
	}
	msg := make([]byte, 1+8+buf.Len()) // 1 header + 8 array count + length
	msg[0] = byte(amf0StrictArr)
	binary.BigEndian.PutUint32(msg[1:], uint32(buf.Len()))
	copy(msg[9:], buf.Bytes())
	return msg
}

func encodeXml(v Amf0Xml) []byte {
	msg := make([]byte, 1+4+len(v)) // 1 header + 4 string length + string
	msg[0] = byte(amf0Xml)
	binary.BigEndian.PutUint32(msg[1:], uint32(len(v)))
	copy(msg[5:], v)
	return msg
}
