// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_CellSmartTag struct {
	// Smart Tag Type Index
	TypeAttr uint32
	// Deleted
	DeletedAttr *bool
	// XML Based
	XmlBasedAttr *bool
	// Smart Tag Properties
	CellSmartTagPr []*CT_CellSmartTagPr
}

func NewCT_CellSmartTag() *CT_CellSmartTag {
	ret := &CT_CellSmartTag{}
	return ret
}

func (m *CT_CellSmartTag) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
		Value: fmt.Sprintf("%v", m.TypeAttr)})
	if m.DeletedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "deleted"},
			Value: fmt.Sprintf("%d", b2i(*m.DeletedAttr))})
	}
	if m.XmlBasedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlBased"},
			Value: fmt.Sprintf("%d", b2i(*m.XmlBasedAttr))})
	}
	e.EncodeToken(start)
	if m.CellSmartTagPr != nil {
		secellSmartTagPr := xml.StartElement{Name: xml.Name{Local: "x:cellSmartTagPr"}}
		e.EncodeElement(m.CellSmartTagPr, secellSmartTagPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellSmartTag) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.TypeAttr = uint32(parsed)
		}
		if attr.Name.Local == "deleted" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DeletedAttr = &parsed
		}
		if attr.Name.Local == "xmlBased" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.XmlBasedAttr = &parsed
		}
	}
lCT_CellSmartTag:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cellSmartTagPr":
				tmp := NewCT_CellSmartTagPr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CellSmartTagPr = append(m.CellSmartTagPr, tmp)
			default:
				log.Printf("skipping unsupported element on CT_CellSmartTag %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellSmartTag
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CellSmartTag and its children
func (m *CT_CellSmartTag) Validate() error {
	return m.ValidateWithPath("CT_CellSmartTag")
}

// ValidateWithPath validates the CT_CellSmartTag and its children, prefixing error messages with path
func (m *CT_CellSmartTag) ValidateWithPath(path string) error {
	for i, v := range m.CellSmartTagPr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CellSmartTagPr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
