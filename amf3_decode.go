package amf

func DecodeAMF3(v []byte) interface{} {
	switch v[0] {
	// case amf0Number:
	// 	return decodeNumber(v)
	// case amf0Boolean:
	// 	return decodeBoolean(v)
	// case amf0String, amf0StringExt:
	// 	return decodeString(v)
	// case amf0Object:
	// 	return decodeObject(v)
	// case amf0Null:
	// 	return nil
	// case amf0Array:
	// 	return decodeECMAArray(v)
	// case amf0StrictArr:
	// 	return decodeStrictArr(v)
	// case amf0Date:
	// 	return decodeDate(v)
	}
	return nil
}
