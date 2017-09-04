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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_CustomWorkbookView struct {
	// Custom View Name
	NameAttr string
	// Custom View GUID
	GuidAttr string
	// Auto Update
	AutoUpdateAttr *bool
	// Merge Interval
	MergeIntervalAttr *uint32
	// Changes Saved Win
	ChangesSavedWinAttr *bool
	// Only Synch
	OnlySyncAttr *bool
	// Personal View
	PersonalViewAttr *bool
	// Include Print Settings
	IncludePrintSettingsAttr *bool
	// Include Hidden Rows & Columns
	IncludeHiddenRowColAttr *bool
	// Maximized
	MaximizedAttr *bool
	// Minimized
	MinimizedAttr *bool
	// Show Horizontal Scroll
	ShowHorizontalScrollAttr *bool
	// Show Vertical Scroll
	ShowVerticalScrollAttr *bool
	// Show Sheet Tabs
	ShowSheetTabsAttr *bool
	// Top Left Corner (X Coordinate)
	XWindowAttr *int32
	// Top Left Corner (Y Coordinate)
	YWindowAttr *int32
	// Window Width
	WindowWidthAttr uint32
	// Window Height
	WindowHeightAttr uint32
	// Sheet Tab Ratio
	TabRatioAttr *uint32
	// Active Sheet in Book View
	ActiveSheetIdAttr uint32
	// Show Formula Bar
	ShowFormulaBarAttr *bool
	// Show Status Bar
	ShowStatusbarAttr *bool
	// Show Comments
	ShowCommentsAttr ST_Comments
	// Show Objects
	ShowObjectsAttr ST_Objects
	ExtLst          *CT_ExtensionList
}

func NewCT_CustomWorkbookView() *CT_CustomWorkbookView {
	ret := &CT_CustomWorkbookView{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_CustomWorkbookView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	if m.AutoUpdateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoUpdate"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoUpdateAttr))})
	}
	if m.MergeIntervalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mergeInterval"},
			Value: fmt.Sprintf("%v", *m.MergeIntervalAttr)})
	}
	if m.ChangesSavedWinAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "changesSavedWin"},
			Value: fmt.Sprintf("%d", b2i(*m.ChangesSavedWinAttr))})
	}
	if m.OnlySyncAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "onlySync"},
			Value: fmt.Sprintf("%d", b2i(*m.OnlySyncAttr))})
	}
	if m.PersonalViewAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "personalView"},
			Value: fmt.Sprintf("%d", b2i(*m.PersonalViewAttr))})
	}
	if m.IncludePrintSettingsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "includePrintSettings"},
			Value: fmt.Sprintf("%d", b2i(*m.IncludePrintSettingsAttr))})
	}
	if m.IncludeHiddenRowColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "includeHiddenRowCol"},
			Value: fmt.Sprintf("%d", b2i(*m.IncludeHiddenRowColAttr))})
	}
	if m.MaximizedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maximized"},
			Value: fmt.Sprintf("%d", b2i(*m.MaximizedAttr))})
	}
	if m.MinimizedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minimized"},
			Value: fmt.Sprintf("%d", b2i(*m.MinimizedAttr))})
	}
	if m.ShowHorizontalScrollAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showHorizontalScroll"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowHorizontalScrollAttr))})
	}
	if m.ShowVerticalScrollAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showVerticalScroll"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowVerticalScrollAttr))})
	}
	if m.ShowSheetTabsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showSheetTabs"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowSheetTabsAttr))})
	}
	if m.XWindowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xWindow"},
			Value: fmt.Sprintf("%v", *m.XWindowAttr)})
	}
	if m.YWindowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "yWindow"},
			Value: fmt.Sprintf("%v", *m.YWindowAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "windowWidth"},
		Value: fmt.Sprintf("%v", m.WindowWidthAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "windowHeight"},
		Value: fmt.Sprintf("%v", m.WindowHeightAttr)})
	if m.TabRatioAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tabRatio"},
			Value: fmt.Sprintf("%v", *m.TabRatioAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "activeSheetId"},
		Value: fmt.Sprintf("%v", m.ActiveSheetIdAttr)})
	if m.ShowFormulaBarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showFormulaBar"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowFormulaBarAttr))})
	}
	if m.ShowStatusbarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showStatusbar"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowStatusbarAttr))})
	}
	if m.ShowCommentsAttr != ST_CommentsUnset {
		attr, err := m.ShowCommentsAttr.MarshalXMLAttr(xml.Name{Local: "showComments"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowObjectsAttr != ST_ObjectsUnset {
		attr, err := m.ShowObjectsAttr.MarshalXMLAttr(xml.Name{Local: "showObjects"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomWorkbookView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
		}
		if attr.Name.Local == "autoUpdate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoUpdateAttr = &parsed
		}
		if attr.Name.Local == "mergeInterval" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MergeIntervalAttr = &pt
		}
		if attr.Name.Local == "changesSavedWin" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ChangesSavedWinAttr = &parsed
		}
		if attr.Name.Local == "onlySync" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OnlySyncAttr = &parsed
		}
		if attr.Name.Local == "personalView" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PersonalViewAttr = &parsed
		}
		if attr.Name.Local == "includePrintSettings" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IncludePrintSettingsAttr = &parsed
		}
		if attr.Name.Local == "includeHiddenRowCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IncludeHiddenRowColAttr = &parsed
		}
		if attr.Name.Local == "maximized" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MaximizedAttr = &parsed
		}
		if attr.Name.Local == "minimized" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MinimizedAttr = &parsed
		}
		if attr.Name.Local == "showHorizontalScroll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowHorizontalScrollAttr = &parsed
		}
		if attr.Name.Local == "showVerticalScroll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowVerticalScrollAttr = &parsed
		}
		if attr.Name.Local == "showSheetTabs" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowSheetTabsAttr = &parsed
		}
		if attr.Name.Local == "xWindow" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.XWindowAttr = &pt
		}
		if attr.Name.Local == "yWindow" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.YWindowAttr = &pt
		}
		if attr.Name.Local == "windowWidth" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.WindowWidthAttr = uint32(parsed)
		}
		if attr.Name.Local == "windowHeight" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.WindowHeightAttr = uint32(parsed)
		}
		if attr.Name.Local == "tabRatio" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TabRatioAttr = &pt
		}
		if attr.Name.Local == "activeSheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ActiveSheetIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "showFormulaBar" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowFormulaBarAttr = &parsed
		}
		if attr.Name.Local == "showStatusbar" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowStatusbarAttr = &parsed
		}
		if attr.Name.Local == "showComments" {
			m.ShowCommentsAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "showObjects" {
			m.ShowObjectsAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_CustomWorkbookView:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_CustomWorkbookView %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomWorkbookView
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomWorkbookView and its children
func (m *CT_CustomWorkbookView) Validate() error {
	return m.ValidateWithPath("CT_CustomWorkbookView")
}

// ValidateWithPath validates the CT_CustomWorkbookView and its children, prefixing error messages with path
func (m *CT_CustomWorkbookView) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if err := m.ShowCommentsAttr.ValidateWithPath(path + "/ShowCommentsAttr"); err != nil {
		return err
	}
	if err := m.ShowObjectsAttr.ValidateWithPath(path + "/ShowObjectsAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
