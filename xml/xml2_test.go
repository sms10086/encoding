package xml

import (
	"testing"
)

type Form struct {
	XMLName Name `xml:"s:form"`
	Something string `xml:"Something"`
}

type Request struct {
	XMLName Name `xml:"http://example.org/ s:request"`
	Form Form
}

func TestMarshalWithPrefix(t* testing.T) {
	
	expected := `<s:request xmlns="http://example.org/"><s:form><Something>test</Something></s:form></s:request>`
	r := new(Request)
	r.Form.Something = "test"
	s, err := Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(s) != expected {
		t.Fatalf("marshal fail")
	}
	err = Unmarshal([]byte(expected), r)
	if err != nil {
		t.Fatal(err)
	}
}