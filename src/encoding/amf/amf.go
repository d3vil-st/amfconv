package amf

type AMFVersion uint8

const (
	AMF0 AMFVersion = 0x1
	AMF3 AMFVersion = 0x3
)

type amf0 byte
type Amf0Reference uint16
type Amf0ECMAArray map[string]interface{}
type Amf0Xml string
type amf3 byte

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
	amf3Undefined = amf3(0x00) // done
	amf3Null      = amf3(0x01) // done
	amf3False     = amf3(0x02) // done
	amf3True      = amf3(0x03) // done
	amf3Integer   = amf3(0x04)
	amf3Double    = amf3(0x05) // done
	amf3String    = amf3(0x06)
	amf3XmlDoc    = amf3(0x07)
	amf3Date      = amf3(0x08)
	amf3Array     = amf3(0x09)
	amf3Object    = amf3(0x0a)
	amf3Xml       = amf3(0x0b)
	amf3ByteArray = amf3(0x0c)
)
