package amf

import (
	"bytes"
	"encoding/binary"
	"math"
	"sort"
	"time"
)

func EncodeAMF0(v interface{}) []byte {
	switch v.(type) {
	case float64:
		return encodeNumber(v.(float64))
	case int:
		return encodeNumber(float64(v.(int)))
	case bool:
		return encodeBoolean(v.(bool))
	case string:
		return encodeString(v.(string))
	case nil:
		return encodeNull()
	case map[string]interface{}:
		return encodeObject(v.(map[string]interface{}))
	case Amf0ECMAArray:
		return encodeECMAArray(v.(Amf0ECMAArray))
	case time.Time:
		return encodeDate(v.(time.Time))
	case []interface{}:
		return encodeStrictArr(v.([]interface{}))
	}
	return nil
}

func encodeNumber(v float64) []byte {
	msg := make([]byte, 1+8) // 1 header + 8 float64
	msg[0] = amf0Number
	binary.BigEndian.PutUint64(msg[1:], uint64(math.Float64bits(v)))
	return msg
}

func encodeBoolean(v bool) []byte {
	msg := make([]byte, 1+1) // 1 header + 1 boolean
	msg[0] = amf0Boolean
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
		msg[0] = amf0String
		binary.BigEndian.PutUint16(msg[1:], uint16(len(v)))
		copy(msg[3:], v)
	} else {
		msg = make([]byte, 1+4+len(v)) // 1 header + 4 length + length of string
		msg[0] = amf0StringExt
		binary.BigEndian.PutUint32(msg[1:], uint32(len(v)))
		copy(msg[5:], v)
	}
	return msg
}

func encodeObject(v map[string]interface{}) []byte {
	buf := new(bytes.Buffer)
	var keys []string
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := v[key]
		buf.Write(EncodeAMF0(key))
		switch value.(type) {
		case int:
			buf.Write(EncodeAMF0(value.(int)))
		case float64:
			buf.Write(EncodeAMF0(value.(float64)))
		case string:
			buf.Write(EncodeAMF0(value.(string)))
		case bool:
			buf.Write(EncodeAMF0(value.(bool)))
		case time.Time:
			buf.Write(EncodeAMF0(value.(time.Time)))
		case nil:
			buf.Write(EncodeAMF0(nil))
		}
	}
	buf.Write(encodeObjectEnd())
	msg := make([]byte, 1+buf.Len()) // 1 header + length
	msg[0] = amf0Object
	copy(msg[1:], buf.Bytes())
	return msg
}

func encodeNull() []byte {
	return []byte{amf0Null}
}

func encodeECMAArray(v Amf0ECMAArray) []byte {
	msg_body := encodeObject(v)
	summary_length := len(msg_body) - 4
	msg := make([]byte, 1+4+summary_length) // 1 header + 4 length + sum length
	msg[0] = amf0Array
	binary.BigEndian.PutUint32(msg[1:], uint32(len(v)))
	copy(msg[5:], msg_body[1:summary_length+1])
	return msg
}

func encodeObjectEnd() []byte {
	return []byte{0x00, 0x00, amf0ObjectEnd}
}

func encodeDate(v time.Time) []byte {
	msg := make([]byte, 1+8+2) // 1 header + 8 float64 + 2 timezone
	msg[0], msg[9], msg[10] = amf0Date, 0x0, 0x0
	binary.BigEndian.PutUint64(msg[1:], uint64(v.UnixNano()/1000000))
	return msg
}

func encodeStrictArr(v []interface{}) []byte {
	buf := new(bytes.Buffer)
	StringExt := false
	for _, k := range v {
		switch k.(type) {
		case string:
			if len(k.(string)) > 0xffff {
				StringExt = true
			}
		}
	}
	for _, k := range v {
		if StringExt {
			msg := make([]byte, 1+4+len(k.(string))) // 1 header + 4 length + length of string
			msg[0] = amf0StringExt
			binary.BigEndian.PutUint32(msg[1:], uint32(len(k.(string))))
			copy(msg[5:], k.(string))
			buf.Write(msg)
			continue
		}
		buf.Write(EncodeAMF0(k))
	}
	msg := make([]byte, 1+8+buf.Len()) // 1 header + 8 array count + length
	msg[0] = amf0StrictArr
	binary.BigEndian.PutUint32(msg[1:], uint32(len(v)))
	copy(msg[9:], buf.Bytes())
	return msg
}
