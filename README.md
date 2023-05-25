# encoding
clone from golang encoding/xml

support prefix

marshal/unmarshal example:
```
type Form struct {
	XMLName xml.Name `xml:"s:form"`
	Something string `xml:"Something"`
}

type Request struct {
	XMLName xml.Name `xml:"http://example.org/ s:request"`
	Form Form
}

func main() {
  r := new(Request)
	r.Form.Something = "test"
	s, _:= xml.Marshal(r) 
  println(s)
}


type Body struct {
  XMLName xml.Name `xml:"s:Body"`
  Value interface{} // set value before marshal/unmarshal
}

type Envelope struct {
  XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ s:Envelope"`
  Body Body
}

```