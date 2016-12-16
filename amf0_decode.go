package amf

import (
	"encoding/binary"
	"math"
	"time"
)

func DecodeAMF0(v []byte) interface{} {
	switch v[0] {
	case amf0Number:
		return decodeNumber(v)
	case amf0Boolean:
		return decodeBoolean(v)
	case amf0String, amf0StringExt:
		return decodeString(v)
	case amf0Object:
		return decodeObject(v)
	case amf0Null:
		return nil
	case amf0Array:
		return decodeECMAArray(v)
	case amf0StrictArr:
		return decodeStrictArr(v)
	case amf0Date:
		return decodeDate(v)
	}
	return nil
}

func decodeNumber(v []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(v[1:]))
}

func decodeBoolean(v []byte) bool {
	if v[1] == 0x1 {
		return true
	} else {
		return false
	}
}

func decodeString(v []byte) string {
	if v[0] == amf0String {
		return string(v[3:])
	} else {
		return string(v[5:])
	}
}

func decodeECMAArray(v []byte) Amf0ECMAArray {
	data := make([]byte, len(v)-4)
	data[0] = amf0Object
	copy(data[1:], v[5:])
	return Amf0ECMAArray(DecodeAMF0(data).(map[string]interface{}))
}

func decodeStrictArr(v []byte) interface{} {
	elem_len := uint(len(v)-9) / uint(binary.BigEndian.Uint32(v[1:9]))
	var arr []interface{}
	if v[9] == amf0String {
		for position := uint(10); position < uint(len(v))-1; {
			elem_len = uint(binary.BigEndian.Uint16(v[position : position+2]))
			arr = append(arr, DecodeAMF0(v[position-1:position+elem_len+2]))
			position += 3 + elem_len
		}
		return arr
	}
	if v[9] == amf0StringExt {
		for position := uint(10); position < uint(len(v))-1; {
			elem_len = uint(binary.BigEndian.Uint32(v[position : position+4]))
			arr = append(arr, DecodeAMF0(v[position-1:position+elem_len+4]))
			position += 5 + elem_len
		}
		return arr
	}
	for position := uint(9); position < uint(len(v))-1; position += elem_len {
		arr = append(arr, DecodeAMF0(v[position:position+elem_len]))
	}
	return arr
}

func decodeDate(v []byte) time.Time {
	return time.Unix(0, int64(binary.BigEndian.Uint64(v[1:9])*1000000))
}

func decodeObject(v []byte) map[string]interface{} {
	msg := make(map[string]interface{})
	for position := 1; position < len(v)-1; {
		if v[position] == byte(0x00) &&
			v[position+1] == byte(0x00) &&
			v[position+2] == byte(0x09) {
			break
		}
		elem_len := int(binary.BigEndian.Uint16(v[position+1 : position+3]))
		key := DecodeAMF0(v[position : position+3+elem_len])
		position += 3 + elem_len
		switch v[position] {
		case amf0Number:
			msg[key.(string)] = DecodeAMF0(v[position : position+9])
			position += 9
		case amf0Boolean:
			msg[key.(string)] = DecodeAMF0(v[position : position+2])
			position += 2
		case amf0String:
			elem_len := int(binary.BigEndian.Uint16(v[position+1 : position+3]))
			msg[key.(string)] = DecodeAMF0(v[position : position+3+elem_len])
			position += 3 + elem_len
		case amf0Null:
			msg[key.(string)] = nil
			position += 1
		case amf0Date:
			msg[key.(string)] = DecodeAMF0(v[position : position+11])
			position += 11
		case amf0StringExt:
			elem_len := int(binary.BigEndian.Uint32(v[position+1 : position+5]))
			msg[key.(string)] = DecodeAMF0(v[position : position+5+elem_len])
			position += 5 + elem_len
		}
	}
	return msg
}
