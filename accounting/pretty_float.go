package accounting

import (
	"encoding/xml"
	"fmt"
)

type Float float64

func (f Float) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(fmt.Sprintf("%.2f", f), start)
}
