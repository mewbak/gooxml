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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_RevisionComment struct {
	// Sheet Id
	SheetIdAttr uint32
	// Cell
	CellAttr string
	// GUID
	GuidAttr string
	// User Action
	ActionAttr ST_RevisionAction
	// Always Show Comment
	AlwaysShowAttr *bool
	// Old Comment
	OldAttr *bool
	// Comment In Hidden Row
	HiddenRowAttr *bool
	// Hidden Column
	HiddenColumnAttr *bool
	// Author
	AuthorAttr string
	// Original Comment Length
	OldLengthAttr *uint32
	// New Comment Length
	NewLengthAttr *uint32
}

func NewCT_RevisionComment() *CT_RevisionComment {
	ret := &CT_RevisionComment{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_RevisionComment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cell"},
		Value: fmt.Sprintf("%v", m.CellAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	if m.ActionAttr != ST_RevisionActionUnset {
		attr, err := m.ActionAttr.MarshalXMLAttr(xml.Name{Local: "action"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AlwaysShowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "alwaysShow"},
			Value: fmt.Sprintf("%d", b2i(*m.AlwaysShowAttr))})
	}
	if m.OldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "old"},
			Value: fmt.Sprintf("%d", b2i(*m.OldAttr))})
	}
	if m.HiddenRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hiddenRow"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenRowAttr))})
	}
	if m.HiddenColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hiddenColumn"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenColumnAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.OldLengthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldLength"},
			Value: fmt.Sprintf("%v", *m.OldLengthAttr)})
	}
	if m.NewLengthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "newLength"},
			Value: fmt.Sprintf("%v", *m.NewLengthAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionComment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "cell" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CellAttr = parsed
		}
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
		}
		if attr.Name.Local == "action" {
			m.ActionAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "alwaysShow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AlwaysShowAttr = &parsed
		}
		if attr.Name.Local == "old" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OldAttr = &parsed
		}
		if attr.Name.Local == "hiddenRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenRowAttr = &parsed
		}
		if attr.Name.Local == "hiddenColumn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenColumnAttr = &parsed
		}
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
		}
		if attr.Name.Local == "oldLength" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.OldLengthAttr = &pt
		}
		if attr.Name.Local == "newLength" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NewLengthAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RevisionComment: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RevisionComment and its children
func (m *CT_RevisionComment) Validate() error {
	return m.ValidateWithPath("CT_RevisionComment")
}

// ValidateWithPath validates the CT_RevisionComment and its children, prefixing error messages with path
func (m *CT_RevisionComment) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if err := m.ActionAttr.ValidateWithPath(path + "/ActionAttr"); err != nil {
		return err
	}
	return nil
}
