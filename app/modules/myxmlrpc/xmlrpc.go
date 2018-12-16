package myxmlrpc

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type MethodRequest struct {
	Name string
	Params []interface{}
}

type MethodResponse struct {
	Name string
	Params []interface{}
}

func (r *MethodRequest) AddParameter(p interface{}) {
	r.Params = append(r.Params, p)
}

func (r *MethodRequest) GetParameter(i int) interface{} {
	return r.Params[i]
}

func (r *MethodResponse) AddParameter(p interface{}) {
	r.Params = append(r.Params, p)
}

func (r *MethodResponse) GetParameter(i int) interface{} {
	return r.Params[i]
}

/**
编码
 */
func EncodeRequest(name string, v interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)

	buf.WriteString(`<?xml version="1.0"?><methodCall>`)
	buf.WriteString("<methodName>" + xmlEscape(name) + "</methodName>")
	buf.WriteString("<params>")

	buf.WriteString("<param><value>")

	buf.WriteString(toXml(v, true))
	buf.WriteString("</value></param>")
	buf.WriteString("</params></methodCall>")

	return buf
}

/**
编码
 */
func EncodeResponse(v interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)

	buf.WriteString(`<?xml version="1.0"?><methodResponse>`)
	buf.WriteString("<params>")
	buf.WriteString("<param><value>")

	if v != nil {
		buf.WriteString(toXml(v, true))
	}

	buf.WriteString("</value></param>")
	buf.WriteString("</params></methodResponse>")

	return buf
}

/**
编码
 */
func EncodeFault(v interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)

	buf.WriteString(`<?xml version="1.0"?><methodResponse>`)
	buf.WriteString("<fault><value>")

	//field := st.Field(i).Name
	buf.WriteString(toXml(v, true))
	buf.WriteString("</value></fault>")
	buf.WriteString("</methodResponse>")

	return buf
}

/**
 解码
 */
func DecodeRequest(r io.Reader) (v *MethodRequest, e error)  {
	//req := MethodRequest{Params: make([]interface{}, 0)}
	p := xml.NewDecoder(r)
	token, e := startElement(p) // methodResponse

	if token.Name.Local != "methodCall" {
		return nil, errors.New("invalid request: missing key word: methodCall")
	}

	token, e = startElement(p) // params
	if token.Name.Local != "methodName" {
		return nil, errors.New("invalid request: missing methodName")
	}

	//parse the method name
	req := MethodRequest{}
	p.DecodeElement(&req.Name, &token)

	token, e = startElement(p) // params
	if token.Name.Local != "params" {
		return nil, errors.New("invalid response: wrong xml tag:" + token.Name.Local + ", expect params")
	}

	//godump.Dump(token.Name.Local)

	for token, e = startElement2(p); e != io.EOF; token, e = startElement2(p){

		if token.Name.Local == "boolean" {
			continue
		}

		if token.Name.Local != "param" {
			return nil, errors.New("invalid response: wrong xml tag:" + token.Name.Local + ", expect param")
		}

		token, e = startElement(p) // param
		if token.Name.Local != "value" {
			return nil, errors.New("invalid request: missing value")
		}

		_, v2, e2 := next(p)
		if e2 != nil {
			return nil, errors.New("parse request parameter failed: " + e2.Error())
		}

		req.AddParameter(v2)
	} // value

	return &req, nil
}

func DecodeResponse(r io.Reader) (v *MethodResponse, e error)  {
	p := xml.NewDecoder(r)
	token, e := startElement(p) // methodResponse

	if token.Name.Local != "methodResponse" {
		return nil, errors.New("invalid request: missing key word: methodResponse")
	}

	rsp := MethodResponse{}

	token, e = startElement(p) // params
	if token.Name.Local == "params" {
		for token, e = startElement(p); e != io.EOF; token, e = startElement(p){
			if token.Name.Local != "param" {
				return nil, errors.New("invalid response: wrong xml tag:" + token.Name.Local + ", expect param")
			}

			token, e = startElement(p) // param
			if token.Name.Local != "value" {
				return nil, errors.New("invalid request: missing value")
			}

			_, v2, e2 := next(p)
			if e2 != nil {
				return nil, errors.New("parse request parameter failed: " + e2.Error())
			}

			rsp.AddParameter(v2)
		} // value

		rsp.Name = "Response"
		return &rsp, nil
	} else if token.Name.Local == "fault" {
		token, e = startElement(p) // value
		if token.Name.Local != "value" {
			return nil, errors.New("invalid response: missing value under fault element")
		}

		_, v2, _ := next(p)
		rsp.AddParameter(v2)
		rsp.Name = "Fault"
		return v, nil
	} else {
		return nil, errors.New("invalid response: missing params or fault")
	}
}

var xmlSpecial = map[byte]string{
	'<':  "&lt;",
	'>':  "&gt;",
	'"':  "&quot;",
	'\'': "&apos;",
	'&':  "&amp;",
}

func xmlEscape(s string) string {
	var b bytes.Buffer
	for i := 0; i < len(s); i++ {
		c := s[i]
		if s, ok := xmlSpecial[c]; ok {
			b.WriteString(s)
		} else {
			b.WriteByte(c)
		}
	}
	return b.String()
}

func toXml(v interface{}, typ bool) (s string) {
	if v == nil {
		return "<nil/>"
	}
	r := reflect.ValueOf(v)
	t := r.Type()
	k := t.Kind()

	if b, ok := v.([]byte); ok {
		return "<base64>" + base64.StdEncoding.EncodeToString(b) + "</base64>"
	}

	switch k {
	case reflect.Invalid:
		panic("unsupported type")
	case reflect.Bool:
		return fmt.Sprintf("<boolean>%v</boolean>", v)
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if typ {
			return fmt.Sprintf("<int>%v</int>", v)
		}
		return fmt.Sprintf("%v", v)
	case reflect.Uintptr:
		panic("unsupported type")
	case reflect.Float32, reflect.Float64:
		if typ {
			return fmt.Sprintf("<double>%v</double>", v)
		}
		return fmt.Sprintf("%v", v)
	case reflect.Complex64, reflect.Complex128:
		panic("unsupported type")
	case reflect.Array:
		s = "<array><data>"
		for n := 0; n < r.Len(); n++ {
			s += "<value>"
			s += toXml(r.Index(n).Interface(), typ)
			s += "</value>"
		}
		s += "</data></array>"
		return s
	case reflect.Chan:
		panic("unsupported type")
	case reflect.Func:
		panic("unsupported type")
	case reflect.Interface:
		return toXml(r.Elem(), typ)
	case reflect.Map:
		s = "<struct>"
		for _, key := range r.MapKeys() {
			s += "<member>"
			s += "<name>" + xmlEscape(key.Interface().(string)) + "</name>"
			s += "<value>" + toXml(r.MapIndex(key).Interface(), typ) + "</value>"
			s += "</member>"
		}
		s += "</struct>"
		return s
	case reflect.Ptr:
		panic("unsupported type")
	case reflect.Slice:
		s = "<array><data>"
		for n := 0; n < r.Len(); n++ {
			s += "<value>"
			s += toXml(r.Index(n).Interface(), typ)
			s += "</value>"
		}
		s += "</data></array>"
		return s
	case reflect.String:
		if typ {
			return fmt.Sprintf("<string>%v</string>", xmlEscape(v.(string)))
		}
		return xmlEscape(v.(string))
	case reflect.Struct:
		if t.Name() == "Time" {
			return fmt.Sprintf("<dateTime.iso8601>%v</dateTime.iso8601>", v.(time.Time).String())
		}

		s = "<struct>"
		for n := 0; n < r.NumField(); n++ {
			rv := r.FieldByIndex([]int{n}).Interface()
			if rv != nil {
				s += "<member>"

				if t.Field(n).Tag.Get("xml") != "" {
					s += "<name>" + t.Field(n).Tag.Get("xml") + "</name>"
				} else {
					s += "<name>" + t.Field(n).Name + "</name>"
				}
				s += "<value>" + toXml(rv, true) + "</value>"
				s += "</member>"
			}
		}
		s += "</struct>"
		return s
	case reflect.UnsafePointer:
		return toXml(r.Elem(), typ)
	}
	return
}

type Array []interface{}
type Struct map[string]interface{}

func next(p *xml.Decoder) (xml.Name, interface{}, error) {
	se, e := startElement(p)
	if e != nil {
		return xml.Name{}, nil, e
	}

	var nv interface{}
	switch se.Name.Local {
	case "string":
		var s string
		if e = p.DecodeElement(&s, &se); e != nil {
			return xml.Name{}, nil, e
		}
		return xml.Name{}, s, nil
	case "boolean":
		var s string
		if e = p.DecodeElement(&s, &se); e != nil {
			return xml.Name{}, nil, e
		}
		s = strings.TrimSpace(s)
		var b bool
		switch s {
		case "true", "1":
			b = true
		case "false", "0":
			b = false
		default:
			e = errors.New("invalid boolean value")
		}
		return xml.Name{}, b, e
	case "int", "i1", "i2", "i4", "i8":
		var s string
		var i int
		if e = p.DecodeElement(&s, &se); e != nil {
			return xml.Name{}, nil, e
		}
		i, e = strconv.Atoi(strings.TrimSpace(s))
		return xml.Name{}, i, e
	case "double":
		var s string
		var f float64
		if e = p.DecodeElement(&s, &se); e != nil {
			return xml.Name{}, nil, e
		}
		f, e = strconv.ParseFloat(strings.TrimSpace(s), 64)
		return xml.Name{}, f, e
	case "dateTime.iso8601":
		var s string
		if e = p.DecodeElement(&s, &se); e != nil {
			return xml.Name{}, nil, e
		}
		t, e := time.Parse("20060102T15:04:05", s)
		if e != nil {
			t, e = time.Parse("2006-01-02T15:04:05-07:00", s)
			if e != nil {
				t, e = time.Parse("2006-01-02T15:04:05", s)
			}
		}
		return xml.Name{}, t, e
	case "base64":
		var s string
		if e = p.DecodeElement(&s, &se); e != nil {
			return xml.Name{}, nil, e
		}
		if b, e := base64.StdEncoding.DecodeString(s); e != nil {
			return xml.Name{}, nil, e
		} else {
			return xml.Name{}, b, nil
		}
	case "member":
		startElement(p)
		return next(p)
	case "value":
		startElement(p)
		return next(p)
	case "name":
		startElement(p)
		return next(p)
	case "struct":
		st := Struct{}

		se, e = startElement(p)
		for e == nil && se.Name.Local == "member" {
			// name
			se, e = startElement(p)
			if se.Name.Local != "name" {
				return xml.Name{}, nil, errors.New("invalid response")
			}
			if e != nil {
				break
			}
			var name string
			if e = p.DecodeElement(&name, &se); e != nil {
				return xml.Name{}, nil, e
			}
			se, e = startElement(p)
			if e != nil {
				break
			}

			// value
			_, value, e := next(p)
			if se.Name.Local != "value" {
				return xml.Name{}, nil, errors.New("invalid response")
			}

			if e != nil {
				break
			}
			st[name] = value

			se, e = startElementInTag(p, "struct")
			if e != nil {
				break
			}
		}
		return xml.Name{}, st, nil
	case "array":
		var ar Array
		startElement(p) // data
		se, e = startElementInTag(p, "array") // top of value
		if se.Name.Local != "value" {
			return xml.Name{}, ar, nil
		}
		for {
			_, value, e := next(p)
			if e != nil {
				break
			}
			ar = append(ar, value)
		}
		return xml.Name{}, ar, nil
	case "nil":
		return xml.Name{}, nil, nil
	}

	if e = p.DecodeElement(nv, &se); e != nil {
		return xml.Name{}, nil, e
	}
	return se.Name, nv, e
}

func startElement2(p *xml.Decoder) (xml.StartElement, error) {
	for {
		t, e := p.Token()
		if e != nil {
			return xml.StartElement{}, e
		}
		switch t := t.(type) {
		case xml.StartElement:
			//godump.Dump(t);godump.Dump("=====\n")
			return t, nil
		}
	}
	panic("unreachable")
}

func startElement(p *xml.Decoder) (xml.StartElement, error) {
	for {
		t, e := p.Token()
		if e != nil {
			return xml.StartElement{}, e
		}
		switch t := t.(type) {
		case xml.StartElement:
			//godump.Dump(t);godump.Dump("=====\n")
			return t, nil
		}
	}
	panic("unreachable")
}


func startElementInTag(p *xml.Decoder, tag string) (xml.StartElement, error) {
	for {
		t, e := p.Token()
		if e != nil {
			return xml.StartElement{}, e
		}
		switch t := t.(type) {
		case xml.StartElement:
			return t, nil
		case xml.EndElement:
			if tag != "" && t.Name.Local == tag {
				return xml.StartElement{}, nil
			}
		}
	}
	panic("unreachable")
}
