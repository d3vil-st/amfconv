package amf

import (
	"encoding/binary"
	"math"
	"time"
)

const (
	amf3MaxInt = 268435455  // (2^28)-1
	amf3MinInt = -268435456 // -(2^28)
)

func EncodeAMF3(v interface{}) []byte {
	switch v.(type) {
	case float64:
		return encodeDouble3(v.(float64))
	case int:
		return encodeInteger3(v.(int))
	case uint:
		return encodeInteger3(v.(int))
	case bool:
		return encodeBoolean3(v.(bool))
	case string:
		return encodeString3(v.(string))
	case nil:
		return encodeNull3()
	// case map[string]interface{}:
	// 	return encodeObject3(v.(map[string]interface{}))
	case time.Time:
		return encodeDate3(v.(time.Time))
		// case []interface{}:
		// 	return encodeStrictArr3(v.([]interface{}))
	}
	return nil
}

func encodeU29(v uint) []byte {
	msg := make([]byte, 0, 4)
	v &= 0x1fffffff
	if v <= 0x7f {
		msg = append(msg, byte(v))
	} else if v <= 0x3fff {
		msg = append(msg, byte((v>>7)|0x80))
		msg = append(msg, byte(v&0x7f))
	} else if v <= 0x1fffff {
		msg = append(msg, byte((v>>14)|0x80))
		msg = append(msg, byte((v>>7)|0x80))
		msg = append(msg, byte(v&0x7f))
	} else {
		msg = append(msg, byte((v>>22)|0x80))
		msg = append(msg, byte((v>>14)|0x80))
		msg = append(msg, byte((v>>7)|0x80))
		msg = append(msg, byte(v&0x7f))
	}
	return msg
}

func encodeInteger3(v int) []byte {
	if v >= amf3MinInt && v <= amf3MaxInt {
		msg := make([]byte, 0, 1+4) // 1 header + up to 4 U29
		msg = append(msg, amf3Integer)
		msg = append(msg, encodeU29(uint(v))...)
		return msg
	} else {
		return encodeDouble3(float64(v))
	}
}

func encodeDouble3(v float64) []byte {
	msg := make([]byte, 1+8) // 1 header + 8 float64
	msg[0] = amf3Double
	binary.BigEndian.PutUint64(msg[1:], uint64(math.Float64bits(v)))
	return msg
}

func encodeBoolean3(v bool) []byte {
	if v {
		return []byte{amf3True}
	} else {
		return []byte{amf3False}
	}
	return []byte{amf3False}
}

func encodeNull3() []byte {
	return []byte{amf3Null}
}

func encodeString3(v string) []byte {
	var strlen = len(v)
	if strlen > amf3MaxInt {
		strlen = amf3MaxInt
	}
	var msg []byte
	msg = append(msg, amf3String)
	msg = append(msg, encodeU29(uint((strlen<<1)|1))...)
	msg = append(msg, []byte(v)[:strlen]...)
	return msg
}

func encodeDate3(v time.Time) []byte {
	msg := make([]byte, 0, 1+1+8) // 1 header + 1 U29 + 8 float64
	msg = append(msg, amf3Date)
	msg = append(msg, encodeU29(1)...)
	msg = append(msg, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}...)
	binary.BigEndian.PutUint64(msg[2:], uint64(v.UnixNano()/1000000))
	return msg
}
