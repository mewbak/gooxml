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
	"strconv"

	"baliance.com/gooxml"
)

type CT_DataBinding struct {
	// Unique Identifer
	DataBindingNameAttr *string
	// Binding to External File
	FileBindingAttr *bool
	// Reference to Connection ID
	ConnectionIDAttr *uint32
	// File Binding Name
	FileBindingNameAttr *string
	// XML Data Loading Behavior
	DataBindingLoadModeAttr uint32
	Any                     gooxml.Any
}

func NewCT_DataBinding() *CT_DataBinding {
	ret := &CT_DataBinding{}
	return ret
}

func (m *CT_DataBinding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DataBindingNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "DataBindingName"},
			Value: fmt.Sprintf("%v", *m.DataBindingNameAttr)})
	}
	if m.FileBindingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "FileBinding"},
			Value: fmt.Sprintf("%d", b2i(*m.FileBindingAttr))})
	}
	if m.ConnectionIDAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ConnectionID"},
			Value: fmt.Sprintf("%v", *m.ConnectionIDAttr)})
	}
	if m.FileBindingNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "FileBindingName"},
			Value: fmt.Sprintf("%v", *m.FileBindingNameAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "DataBindingLoadMode"},
		Value: fmt.Sprintf("%v", m.DataBindingLoadModeAttr)})
	e.EncodeToken(start)
	if m.Any != nil {
		m.Any.MarshalXML(e, start)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataBinding) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "DataBindingName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DataBindingNameAttr = &parsed
		}
		if attr.Name.Local == "FileBinding" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FileBindingAttr = &parsed
		}
		if attr.Name.Local == "ConnectionID" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ConnectionIDAttr = &pt
		}
		if attr.Name.Local == "FileBindingName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FileBindingNameAttr = &parsed
		}
		if attr.Name.Local == "DataBindingLoadMode" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.DataBindingLoadModeAttr = uint32(parsed)
		}
	}
lCT_DataBinding:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			default:
				if anyEl, err := gooxml.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lCT_DataBinding
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DataBinding and its children
func (m *CT_DataBinding) Validate() error {
	return m.ValidateWithPath("CT_DataBinding")
}

// ValidateWithPath validates the CT_DataBinding and its children, prefixing error messages with path
func (m *CT_DataBinding) ValidateWithPath(path string) error {
	return nil
}
