package types

import (
	"encoding/xml"
)

// ContentNamespace is the namespace for RSS format's content module.
type ContentNamespace string

func (ns *ContentNamespace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:content"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:content"}, Value: "http://purl.org/rss/1.0/modules/content/"}, nil
}
