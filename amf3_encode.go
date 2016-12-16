package amf

import (
	"encoding/binary"
	"math"
)

func EncodeAMF3(v interface{}) []byte {
	switch v.(type) {
	case float64:
		return encodeDouble3(v.(float64))
	/*case int:
	  return encodeInteger(float64(v.(int)))*/
	case bool:
		return encodeBoolean3(v.(bool))
		/*case string:
		  return encodeString3(v.(string))*/
	case nil:
		return encodeNull3()
		/*     case Amf0Reference:
		return encodeReference3(v.(Amf0Reference))
		case map[string]interface{}:
		return encodeObject3(v.(map[string]interface{}))
		case Amf0ECMAArray:
		return encodeECMAArray3(v.(Amf0ECMAArray))
		case time.Time:
		return encodeDate3(v.(time.Time))
		case []interface{}:
		return encodeStrictArr3(v.([]interface{}))
		case Amf0Xml:
		return encodeXml3(v.(Amf0Xml)) */
	}
	return nil
}

func encodeDouble3(v float64) []byte {
	msg := make([]byte, 1+8) // 1 header + 8 float64
	msg[0] = byte(amf3Double)
	binary.BigEndian.PutUint64(msg[1:], uint64(math.Float64bits(v)))
	return msg
}

func encodeBoolean3(v bool) []byte {
	if v {
		return []byte{byte(amf3True)}
	} else {
		return []byte{byte(amf3False)}
	}
	return []byte{byte(amf3False)}
}

func encodeNull3() []byte {
	return []byte{byte(amf3Null)}
}

func encodeString3(v string) []byte {
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

/* func encodeUndefined3() []byte {
    return []byte{amf3Undefined}
}*/
