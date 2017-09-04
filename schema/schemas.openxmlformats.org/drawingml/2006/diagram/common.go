// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml"
)

func ParseUnionST_ModelId(s string) (ST_ModelId, error) {
	// TODO: implement
	return ST_ModelId{}, nil
}
func ParseUnionST_LayoutShapeType(s string) (ST_LayoutShapeType, error) {
	// TODO: implement
	return ST_LayoutShapeType{}, nil
}
func ParseUnionST_ParameterVal(s string) (ST_ParameterVal, error) {
	// TODO: implement
	return ST_ParameterVal{}, nil
}
func ParseUnionST_FunctionArgument(s string) (ST_FunctionArgument, error) {
	// TODO: implement
	return ST_FunctionArgument{}, nil
}
func ParseUnionST_FunctionValue(s string) (ST_FunctionValue, error) {
	// TODO: implement
	return ST_FunctionValue{}, nil
}
func ParseUnionST_PrSetCustVal(s string) (ST_PrSetCustVal, error) {
	// TODO: implement
	return ST_PrSetCustVal{}, nil
}
func ParseSliceST_AxisTypes(s string) (ST_AxisTypes, error) {
	// TODO: implement
	return ST_AxisTypes{}, nil
}
func ParseSliceST_ElementTypes(s string) (ST_ElementTypes, error) {
	// TODO: implement
	return ST_ElementTypes{}, nil
}
func ParseSliceST_Booleans(s string) (ST_Booleans, error) {
	// TODO: implement
	return ST_Booleans{}, nil
}
func ParseSliceST_Ints(s string) (ST_Ints, error) {
	// TODO: implement
	return ST_Ints{}, nil
}
func ParseSliceST_UnsignedInts(s string) (ST_UnsignedInts, error) {
	// TODO: implement
	return ST_UnsignedInts{}, nil
}

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

type ST_ClrAppMethod byte

const (
	ST_ClrAppMethodUnset  ST_ClrAppMethod = 0
	ST_ClrAppMethodSpan   ST_ClrAppMethod = 1
	ST_ClrAppMethodCycle  ST_ClrAppMethod = 2
	ST_ClrAppMethodRepeat ST_ClrAppMethod = 3
)

func (e ST_ClrAppMethod) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ClrAppMethodUnset:
		attr.Value = ""
	case ST_ClrAppMethodSpan:
		attr.Value = "span"
	case ST_ClrAppMethodCycle:
		attr.Value = "cycle"
	case ST_ClrAppMethodRepeat:
		attr.Value = "repeat"
	}
	return attr, nil
}

func (e *ST_ClrAppMethod) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "span":
		*e = 1
	case "cycle":
		*e = 2
	case "repeat":
		*e = 3
	}
	return nil
}

func (m ST_ClrAppMethod) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ClrAppMethod) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "span":
			*m = 1
		case "cycle":
			*m = 2
		case "repeat":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ClrAppMethod) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "span"
	case 2:
		return "cycle"
	case 3:
		return "repeat"
	}
	return ""
}

func (m ST_ClrAppMethod) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ClrAppMethod) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HueDir byte

const (
	ST_HueDirUnset ST_HueDir = 0
	ST_HueDirCw    ST_HueDir = 1
	ST_HueDirCcw   ST_HueDir = 2
)

func (e ST_HueDir) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HueDirUnset:
		attr.Value = ""
	case ST_HueDirCw:
		attr.Value = "cw"
	case ST_HueDirCcw:
		attr.Value = "ccw"
	}
	return attr, nil
}

func (e *ST_HueDir) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "cw":
		*e = 1
	case "ccw":
		*e = 2
	}
	return nil
}

func (m ST_HueDir) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HueDir) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "cw":
			*m = 1
		case "ccw":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HueDir) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "cw"
	case 2:
		return "ccw"
	}
	return ""
}

func (m ST_HueDir) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HueDir) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PtType byte

const (
	ST_PtTypeUnset    ST_PtType = 0
	ST_PtTypeNode     ST_PtType = 1
	ST_PtTypeAsst     ST_PtType = 2
	ST_PtTypeDoc      ST_PtType = 3
	ST_PtTypePres     ST_PtType = 4
	ST_PtTypeParTrans ST_PtType = 5
	ST_PtTypeSibTrans ST_PtType = 6
)

func (e ST_PtType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PtTypeUnset:
		attr.Value = ""
	case ST_PtTypeNode:
		attr.Value = "node"
	case ST_PtTypeAsst:
		attr.Value = "asst"
	case ST_PtTypeDoc:
		attr.Value = "doc"
	case ST_PtTypePres:
		attr.Value = "pres"
	case ST_PtTypeParTrans:
		attr.Value = "parTrans"
	case ST_PtTypeSibTrans:
		attr.Value = "sibTrans"
	}
	return attr, nil
}

func (e *ST_PtType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "node":
		*e = 1
	case "asst":
		*e = 2
	case "doc":
		*e = 3
	case "pres":
		*e = 4
	case "parTrans":
		*e = 5
	case "sibTrans":
		*e = 6
	}
	return nil
}

func (m ST_PtType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PtType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "node":
			*m = 1
		case "asst":
			*m = 2
		case "doc":
			*m = 3
		case "pres":
			*m = 4
		case "parTrans":
			*m = 5
		case "sibTrans":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PtType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "node"
	case 2:
		return "asst"
	case 3:
		return "doc"
	case 4:
		return "pres"
	case 5:
		return "parTrans"
	case 6:
		return "sibTrans"
	}
	return ""
}

func (m ST_PtType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PtType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CxnType byte

const (
	ST_CxnTypeUnset               ST_CxnType = 0
	ST_CxnTypeParOf               ST_CxnType = 1
	ST_CxnTypePresOf              ST_CxnType = 2
	ST_CxnTypePresParOf           ST_CxnType = 3
	ST_CxnTypeUnknownRelationship ST_CxnType = 4
)

func (e ST_CxnType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CxnTypeUnset:
		attr.Value = ""
	case ST_CxnTypeParOf:
		attr.Value = "parOf"
	case ST_CxnTypePresOf:
		attr.Value = "presOf"
	case ST_CxnTypePresParOf:
		attr.Value = "presParOf"
	case ST_CxnTypeUnknownRelationship:
		attr.Value = "unknownRelationship"
	}
	return attr, nil
}

func (e *ST_CxnType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "parOf":
		*e = 1
	case "presOf":
		*e = 2
	case "presParOf":
		*e = 3
	case "unknownRelationship":
		*e = 4
	}
	return nil
}

func (m ST_CxnType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CxnType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "parOf":
			*m = 1
		case "presOf":
			*m = 2
		case "presParOf":
			*m = 3
		case "unknownRelationship":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_CxnType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "parOf"
	case 2:
		return "presOf"
	case 3:
		return "presParOf"
	case 4:
		return "unknownRelationship"
	}
	return ""
}

func (m ST_CxnType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CxnType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Direction byte

const (
	ST_DirectionUnset ST_Direction = 0
	ST_DirectionNorm  ST_Direction = 1
	ST_DirectionRev   ST_Direction = 2
)

func (e ST_Direction) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DirectionUnset:
		attr.Value = ""
	case ST_DirectionNorm:
		attr.Value = "norm"
	case ST_DirectionRev:
		attr.Value = "rev"
	}
	return attr, nil
}

func (e *ST_Direction) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "norm":
		*e = 1
	case "rev":
		*e = 2
	}
	return nil
}

func (m ST_Direction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Direction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "norm":
			*m = 1
		case "rev":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Direction) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "norm"
	case 2:
		return "rev"
	}
	return ""
}

func (m ST_Direction) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Direction) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HierBranchStyle byte

const (
	ST_HierBranchStyleUnset ST_HierBranchStyle = 0
	ST_HierBranchStyleL     ST_HierBranchStyle = 1
	ST_HierBranchStyleR     ST_HierBranchStyle = 2
	ST_HierBranchStyleHang  ST_HierBranchStyle = 3
	ST_HierBranchStyleStd   ST_HierBranchStyle = 4
	ST_HierBranchStyleInit  ST_HierBranchStyle = 5
)

func (e ST_HierBranchStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HierBranchStyleUnset:
		attr.Value = ""
	case ST_HierBranchStyleL:
		attr.Value = "l"
	case ST_HierBranchStyleR:
		attr.Value = "r"
	case ST_HierBranchStyleHang:
		attr.Value = "hang"
	case ST_HierBranchStyleStd:
		attr.Value = "std"
	case ST_HierBranchStyleInit:
		attr.Value = "init"
	}
	return attr, nil
}

func (e *ST_HierBranchStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "l":
		*e = 1
	case "r":
		*e = 2
	case "hang":
		*e = 3
	case "std":
		*e = 4
	case "init":
		*e = 5
	}
	return nil
}

func (m ST_HierBranchStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HierBranchStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "l":
			*m = 1
		case "r":
			*m = 2
		case "hang":
			*m = 3
		case "std":
			*m = 4
		case "init":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HierBranchStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "l"
	case 2:
		return "r"
	case 3:
		return "hang"
	case 4:
		return "std"
	case 5:
		return "init"
	}
	return ""
}

func (m ST_HierBranchStyle) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HierBranchStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AnimOneStr byte

const (
	ST_AnimOneStrUnset  ST_AnimOneStr = 0
	ST_AnimOneStrNone   ST_AnimOneStr = 1
	ST_AnimOneStrOne    ST_AnimOneStr = 2
	ST_AnimOneStrBranch ST_AnimOneStr = 3
)

func (e ST_AnimOneStr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AnimOneStrUnset:
		attr.Value = ""
	case ST_AnimOneStrNone:
		attr.Value = "none"
	case ST_AnimOneStrOne:
		attr.Value = "one"
	case ST_AnimOneStrBranch:
		attr.Value = "branch"
	}
	return attr, nil
}

func (e *ST_AnimOneStr) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "one":
		*e = 2
	case "branch":
		*e = 3
	}
	return nil
}

func (m ST_AnimOneStr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AnimOneStr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "one":
			*m = 2
		case "branch":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_AnimOneStr) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "one"
	case 3:
		return "branch"
	}
	return ""
}

func (m ST_AnimOneStr) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AnimOneStr) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AnimLvlStr byte

const (
	ST_AnimLvlStrUnset ST_AnimLvlStr = 0
	ST_AnimLvlStrNone  ST_AnimLvlStr = 1
	ST_AnimLvlStrLvl   ST_AnimLvlStr = 2
	ST_AnimLvlStrCtr   ST_AnimLvlStr = 3
)

func (e ST_AnimLvlStr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AnimLvlStrUnset:
		attr.Value = ""
	case ST_AnimLvlStrNone:
		attr.Value = "none"
	case ST_AnimLvlStrLvl:
		attr.Value = "lvl"
	case ST_AnimLvlStrCtr:
		attr.Value = "ctr"
	}
	return attr, nil
}

func (e *ST_AnimLvlStr) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "lvl":
		*e = 2
	case "ctr":
		*e = 3
	}
	return nil
}

func (m ST_AnimLvlStr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AnimLvlStr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "lvl":
			*m = 2
		case "ctr":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_AnimLvlStr) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "lvl"
	case 3:
		return "ctr"
	}
	return ""
}

func (m ST_AnimLvlStr) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AnimLvlStr) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ResizeHandlesStr byte

const (
	ST_ResizeHandlesStrUnset ST_ResizeHandlesStr = 0
	ST_ResizeHandlesStrExact ST_ResizeHandlesStr = 1
	ST_ResizeHandlesStrRel   ST_ResizeHandlesStr = 2
)

func (e ST_ResizeHandlesStr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ResizeHandlesStrUnset:
		attr.Value = ""
	case ST_ResizeHandlesStrExact:
		attr.Value = "exact"
	case ST_ResizeHandlesStrRel:
		attr.Value = "rel"
	}
	return attr, nil
}

func (e *ST_ResizeHandlesStr) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "exact":
		*e = 1
	case "rel":
		*e = 2
	}
	return nil
}

func (m ST_ResizeHandlesStr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ResizeHandlesStr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "exact":
			*m = 1
		case "rel":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ResizeHandlesStr) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "exact"
	case 2:
		return "rel"
	}
	return ""
}

func (m ST_ResizeHandlesStr) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ResizeHandlesStr) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AlgorithmType byte

const (
	ST_AlgorithmTypeUnset     ST_AlgorithmType = 0
	ST_AlgorithmTypeComposite ST_AlgorithmType = 1
	ST_AlgorithmTypeConn      ST_AlgorithmType = 2
	ST_AlgorithmTypeCycle     ST_AlgorithmType = 3
	ST_AlgorithmTypeHierChild ST_AlgorithmType = 4
	ST_AlgorithmTypeHierRoot  ST_AlgorithmType = 5
	ST_AlgorithmTypePyra      ST_AlgorithmType = 6
	ST_AlgorithmTypeLin       ST_AlgorithmType = 7
	ST_AlgorithmTypeSp        ST_AlgorithmType = 8
	ST_AlgorithmTypeTx        ST_AlgorithmType = 9
	ST_AlgorithmTypeSnake     ST_AlgorithmType = 10
)

func (e ST_AlgorithmType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AlgorithmTypeUnset:
		attr.Value = ""
	case ST_AlgorithmTypeComposite:
		attr.Value = "composite"
	case ST_AlgorithmTypeConn:
		attr.Value = "conn"
	case ST_AlgorithmTypeCycle:
		attr.Value = "cycle"
	case ST_AlgorithmTypeHierChild:
		attr.Value = "hierChild"
	case ST_AlgorithmTypeHierRoot:
		attr.Value = "hierRoot"
	case ST_AlgorithmTypePyra:
		attr.Value = "pyra"
	case ST_AlgorithmTypeLin:
		attr.Value = "lin"
	case ST_AlgorithmTypeSp:
		attr.Value = "sp"
	case ST_AlgorithmTypeTx:
		attr.Value = "tx"
	case ST_AlgorithmTypeSnake:
		attr.Value = "snake"
	}
	return attr, nil
}

func (e *ST_AlgorithmType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "composite":
		*e = 1
	case "conn":
		*e = 2
	case "cycle":
		*e = 3
	case "hierChild":
		*e = 4
	case "hierRoot":
		*e = 5
	case "pyra":
		*e = 6
	case "lin":
		*e = 7
	case "sp":
		*e = 8
	case "tx":
		*e = 9
	case "snake":
		*e = 10
	}
	return nil
}

func (m ST_AlgorithmType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AlgorithmType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "composite":
			*m = 1
		case "conn":
			*m = 2
		case "cycle":
			*m = 3
		case "hierChild":
			*m = 4
		case "hierRoot":
			*m = 5
		case "pyra":
			*m = 6
		case "lin":
			*m = 7
		case "sp":
			*m = 8
		case "tx":
			*m = 9
		case "snake":
			*m = 10
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_AlgorithmType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "composite"
	case 2:
		return "conn"
	case 3:
		return "cycle"
	case 4:
		return "hierChild"
	case 5:
		return "hierRoot"
	case 6:
		return "pyra"
	case 7:
		return "lin"
	case 8:
		return "sp"
	case 9:
		return "tx"
	case 10:
		return "snake"
	}
	return ""
}

func (m ST_AlgorithmType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AlgorithmType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AxisType byte

const (
	ST_AxisTypeUnset       ST_AxisType = 0
	ST_AxisTypeSelf        ST_AxisType = 1
	ST_AxisTypeCh          ST_AxisType = 2
	ST_AxisTypeDes         ST_AxisType = 3
	ST_AxisTypeDesOrSelf   ST_AxisType = 4
	ST_AxisTypePar         ST_AxisType = 5
	ST_AxisTypeAncst       ST_AxisType = 6
	ST_AxisTypeAncstOrSelf ST_AxisType = 7
	ST_AxisTypeFollowSib   ST_AxisType = 8
	ST_AxisTypePrecedSib   ST_AxisType = 9
	ST_AxisTypeFollow      ST_AxisType = 10
	ST_AxisTypePreced      ST_AxisType = 11
	ST_AxisTypeRoot        ST_AxisType = 12
	ST_AxisTypeNone        ST_AxisType = 13
)

func (e ST_AxisType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AxisTypeUnset:
		attr.Value = ""
	case ST_AxisTypeSelf:
		attr.Value = "self"
	case ST_AxisTypeCh:
		attr.Value = "ch"
	case ST_AxisTypeDes:
		attr.Value = "des"
	case ST_AxisTypeDesOrSelf:
		attr.Value = "desOrSelf"
	case ST_AxisTypePar:
		attr.Value = "par"
	case ST_AxisTypeAncst:
		attr.Value = "ancst"
	case ST_AxisTypeAncstOrSelf:
		attr.Value = "ancstOrSelf"
	case ST_AxisTypeFollowSib:
		attr.Value = "followSib"
	case ST_AxisTypePrecedSib:
		attr.Value = "precedSib"
	case ST_AxisTypeFollow:
		attr.Value = "follow"
	case ST_AxisTypePreced:
		attr.Value = "preced"
	case ST_AxisTypeRoot:
		attr.Value = "root"
	case ST_AxisTypeNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_AxisType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "self":
		*e = 1
	case "ch":
		*e = 2
	case "des":
		*e = 3
	case "desOrSelf":
		*e = 4
	case "par":
		*e = 5
	case "ancst":
		*e = 6
	case "ancstOrSelf":
		*e = 7
	case "followSib":
		*e = 8
	case "precedSib":
		*e = 9
	case "follow":
		*e = 10
	case "preced":
		*e = 11
	case "root":
		*e = 12
	case "none":
		*e = 13
	}
	return nil
}

func (m ST_AxisType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AxisType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "self":
			*m = 1
		case "ch":
			*m = 2
		case "des":
			*m = 3
		case "desOrSelf":
			*m = 4
		case "par":
			*m = 5
		case "ancst":
			*m = 6
		case "ancstOrSelf":
			*m = 7
		case "followSib":
			*m = 8
		case "precedSib":
			*m = 9
		case "follow":
			*m = 10
		case "preced":
			*m = 11
		case "root":
			*m = 12
		case "none":
			*m = 13
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_AxisType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "self"
	case 2:
		return "ch"
	case 3:
		return "des"
	case 4:
		return "desOrSelf"
	case 5:
		return "par"
	case 6:
		return "ancst"
	case 7:
		return "ancstOrSelf"
	case 8:
		return "followSib"
	case 9:
		return "precedSib"
	case 10:
		return "follow"
	case 11:
		return "preced"
	case 12:
		return "root"
	case 13:
		return "none"
	}
	return ""
}

func (m ST_AxisType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AxisType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BoolOperator byte

const (
	ST_BoolOperatorUnset ST_BoolOperator = 0
	ST_BoolOperatorNone  ST_BoolOperator = 1
	ST_BoolOperatorEqu   ST_BoolOperator = 2
	ST_BoolOperatorGte   ST_BoolOperator = 3
	ST_BoolOperatorLte   ST_BoolOperator = 4
)

func (e ST_BoolOperator) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BoolOperatorUnset:
		attr.Value = ""
	case ST_BoolOperatorNone:
		attr.Value = "none"
	case ST_BoolOperatorEqu:
		attr.Value = "equ"
	case ST_BoolOperatorGte:
		attr.Value = "gte"
	case ST_BoolOperatorLte:
		attr.Value = "lte"
	}
	return attr, nil
}

func (e *ST_BoolOperator) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "equ":
		*e = 2
	case "gte":
		*e = 3
	case "lte":
		*e = 4
	}
	return nil
}

func (m ST_BoolOperator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BoolOperator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "equ":
			*m = 2
		case "gte":
			*m = 3
		case "lte":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_BoolOperator) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "equ"
	case 3:
		return "gte"
	case 4:
		return "lte"
	}
	return ""
}

func (m ST_BoolOperator) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BoolOperator) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ChildOrderType byte

const (
	ST_ChildOrderTypeUnset ST_ChildOrderType = 0
	ST_ChildOrderTypeB     ST_ChildOrderType = 1
	ST_ChildOrderTypeT     ST_ChildOrderType = 2
)

func (e ST_ChildOrderType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ChildOrderTypeUnset:
		attr.Value = ""
	case ST_ChildOrderTypeB:
		attr.Value = "b"
	case ST_ChildOrderTypeT:
		attr.Value = "t"
	}
	return attr, nil
}

func (e *ST_ChildOrderType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "b":
		*e = 1
	case "t":
		*e = 2
	}
	return nil
}

func (m ST_ChildOrderType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ChildOrderType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "b":
			*m = 1
		case "t":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ChildOrderType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "b"
	case 2:
		return "t"
	}
	return ""
}

func (m ST_ChildOrderType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ChildOrderType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ConstraintType byte

const (
	ST_ConstraintTypeUnset         ST_ConstraintType = 0
	ST_ConstraintTypeNone          ST_ConstraintType = 1
	ST_ConstraintTypeAlignOff      ST_ConstraintType = 2
	ST_ConstraintTypeBegMarg       ST_ConstraintType = 3
	ST_ConstraintTypeBendDist      ST_ConstraintType = 4
	ST_ConstraintTypeBegPad        ST_ConstraintType = 5
	ST_ConstraintTypeB             ST_ConstraintType = 6
	ST_ConstraintTypeBMarg         ST_ConstraintType = 7
	ST_ConstraintTypeBOff          ST_ConstraintType = 8
	ST_ConstraintTypeCtrX          ST_ConstraintType = 9
	ST_ConstraintTypeCtrXOff       ST_ConstraintType = 10
	ST_ConstraintTypeCtrY          ST_ConstraintType = 11
	ST_ConstraintTypeCtrYOff       ST_ConstraintType = 12
	ST_ConstraintTypeConnDist      ST_ConstraintType = 13
	ST_ConstraintTypeDiam          ST_ConstraintType = 14
	ST_ConstraintTypeEndMarg       ST_ConstraintType = 15
	ST_ConstraintTypeEndPad        ST_ConstraintType = 16
	ST_ConstraintTypeH             ST_ConstraintType = 17
	ST_ConstraintTypeHArH          ST_ConstraintType = 18
	ST_ConstraintTypeHOff          ST_ConstraintType = 19
	ST_ConstraintTypeL             ST_ConstraintType = 20
	ST_ConstraintTypeLMarg         ST_ConstraintType = 21
	ST_ConstraintTypeLOff          ST_ConstraintType = 22
	ST_ConstraintTypeR             ST_ConstraintType = 23
	ST_ConstraintTypeRMarg         ST_ConstraintType = 24
	ST_ConstraintTypeROff          ST_ConstraintType = 25
	ST_ConstraintTypePrimFontSz    ST_ConstraintType = 26
	ST_ConstraintTypePyraAcctRatio ST_ConstraintType = 27
	ST_ConstraintTypeSecFontSz     ST_ConstraintType = 28
	ST_ConstraintTypeSibSp         ST_ConstraintType = 29
	ST_ConstraintTypeSecSibSp      ST_ConstraintType = 30
	ST_ConstraintTypeSp            ST_ConstraintType = 31
	ST_ConstraintTypeStemThick     ST_ConstraintType = 32
	ST_ConstraintTypeT             ST_ConstraintType = 33
	ST_ConstraintTypeTMarg         ST_ConstraintType = 34
	ST_ConstraintTypeTOff          ST_ConstraintType = 35
	ST_ConstraintTypeUserA         ST_ConstraintType = 36
	ST_ConstraintTypeUserB         ST_ConstraintType = 37
	ST_ConstraintTypeUserC         ST_ConstraintType = 38
	ST_ConstraintTypeUserD         ST_ConstraintType = 39
	ST_ConstraintTypeUserE         ST_ConstraintType = 40
	ST_ConstraintTypeUserF         ST_ConstraintType = 41
	ST_ConstraintTypeUserG         ST_ConstraintType = 42
	ST_ConstraintTypeUserH         ST_ConstraintType = 43
	ST_ConstraintTypeUserI         ST_ConstraintType = 44
	ST_ConstraintTypeUserJ         ST_ConstraintType = 45
	ST_ConstraintTypeUserK         ST_ConstraintType = 46
	ST_ConstraintTypeUserL         ST_ConstraintType = 47
	ST_ConstraintTypeUserM         ST_ConstraintType = 48
	ST_ConstraintTypeUserN         ST_ConstraintType = 49
	ST_ConstraintTypeUserO         ST_ConstraintType = 50
	ST_ConstraintTypeUserP         ST_ConstraintType = 51
	ST_ConstraintTypeUserQ         ST_ConstraintType = 52
	ST_ConstraintTypeUserR         ST_ConstraintType = 53
	ST_ConstraintTypeUserS         ST_ConstraintType = 54
	ST_ConstraintTypeUserT         ST_ConstraintType = 55
	ST_ConstraintTypeUserU         ST_ConstraintType = 56
	ST_ConstraintTypeUserV         ST_ConstraintType = 57
	ST_ConstraintTypeUserW         ST_ConstraintType = 58
	ST_ConstraintTypeUserX         ST_ConstraintType = 59
	ST_ConstraintTypeUserY         ST_ConstraintType = 60
	ST_ConstraintTypeUserZ         ST_ConstraintType = 61
	ST_ConstraintTypeW             ST_ConstraintType = 62
	ST_ConstraintTypeWArH          ST_ConstraintType = 63
	ST_ConstraintTypeWOff          ST_ConstraintType = 64
)

func (e ST_ConstraintType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ConstraintTypeUnset:
		attr.Value = ""
	case ST_ConstraintTypeNone:
		attr.Value = "none"
	case ST_ConstraintTypeAlignOff:
		attr.Value = "alignOff"
	case ST_ConstraintTypeBegMarg:
		attr.Value = "begMarg"
	case ST_ConstraintTypeBendDist:
		attr.Value = "bendDist"
	case ST_ConstraintTypeBegPad:
		attr.Value = "begPad"
	case ST_ConstraintTypeB:
		attr.Value = "b"
	case ST_ConstraintTypeBMarg:
		attr.Value = "bMarg"
	case ST_ConstraintTypeBOff:
		attr.Value = "bOff"
	case ST_ConstraintTypeCtrX:
		attr.Value = "ctrX"
	case ST_ConstraintTypeCtrXOff:
		attr.Value = "ctrXOff"
	case ST_ConstraintTypeCtrY:
		attr.Value = "ctrY"
	case ST_ConstraintTypeCtrYOff:
		attr.Value = "ctrYOff"
	case ST_ConstraintTypeConnDist:
		attr.Value = "connDist"
	case ST_ConstraintTypeDiam:
		attr.Value = "diam"
	case ST_ConstraintTypeEndMarg:
		attr.Value = "endMarg"
	case ST_ConstraintTypeEndPad:
		attr.Value = "endPad"
	case ST_ConstraintTypeH:
		attr.Value = "h"
	case ST_ConstraintTypeHArH:
		attr.Value = "hArH"
	case ST_ConstraintTypeHOff:
		attr.Value = "hOff"
	case ST_ConstraintTypeL:
		attr.Value = "l"
	case ST_ConstraintTypeLMarg:
		attr.Value = "lMarg"
	case ST_ConstraintTypeLOff:
		attr.Value = "lOff"
	case ST_ConstraintTypeR:
		attr.Value = "r"
	case ST_ConstraintTypeRMarg:
		attr.Value = "rMarg"
	case ST_ConstraintTypeROff:
		attr.Value = "rOff"
	case ST_ConstraintTypePrimFontSz:
		attr.Value = "primFontSz"
	case ST_ConstraintTypePyraAcctRatio:
		attr.Value = "pyraAcctRatio"
	case ST_ConstraintTypeSecFontSz:
		attr.Value = "secFontSz"
	case ST_ConstraintTypeSibSp:
		attr.Value = "sibSp"
	case ST_ConstraintTypeSecSibSp:
		attr.Value = "secSibSp"
	case ST_ConstraintTypeSp:
		attr.Value = "sp"
	case ST_ConstraintTypeStemThick:
		attr.Value = "stemThick"
	case ST_ConstraintTypeT:
		attr.Value = "t"
	case ST_ConstraintTypeTMarg:
		attr.Value = "tMarg"
	case ST_ConstraintTypeTOff:
		attr.Value = "tOff"
	case ST_ConstraintTypeUserA:
		attr.Value = "userA"
	case ST_ConstraintTypeUserB:
		attr.Value = "userB"
	case ST_ConstraintTypeUserC:
		attr.Value = "userC"
	case ST_ConstraintTypeUserD:
		attr.Value = "userD"
	case ST_ConstraintTypeUserE:
		attr.Value = "userE"
	case ST_ConstraintTypeUserF:
		attr.Value = "userF"
	case ST_ConstraintTypeUserG:
		attr.Value = "userG"
	case ST_ConstraintTypeUserH:
		attr.Value = "userH"
	case ST_ConstraintTypeUserI:
		attr.Value = "userI"
	case ST_ConstraintTypeUserJ:
		attr.Value = "userJ"
	case ST_ConstraintTypeUserK:
		attr.Value = "userK"
	case ST_ConstraintTypeUserL:
		attr.Value = "userL"
	case ST_ConstraintTypeUserM:
		attr.Value = "userM"
	case ST_ConstraintTypeUserN:
		attr.Value = "userN"
	case ST_ConstraintTypeUserO:
		attr.Value = "userO"
	case ST_ConstraintTypeUserP:
		attr.Value = "userP"
	case ST_ConstraintTypeUserQ:
		attr.Value = "userQ"
	case ST_ConstraintTypeUserR:
		attr.Value = "userR"
	case ST_ConstraintTypeUserS:
		attr.Value = "userS"
	case ST_ConstraintTypeUserT:
		attr.Value = "userT"
	case ST_ConstraintTypeUserU:
		attr.Value = "userU"
	case ST_ConstraintTypeUserV:
		attr.Value = "userV"
	case ST_ConstraintTypeUserW:
		attr.Value = "userW"
	case ST_ConstraintTypeUserX:
		attr.Value = "userX"
	case ST_ConstraintTypeUserY:
		attr.Value = "userY"
	case ST_ConstraintTypeUserZ:
		attr.Value = "userZ"
	case ST_ConstraintTypeW:
		attr.Value = "w"
	case ST_ConstraintTypeWArH:
		attr.Value = "wArH"
	case ST_ConstraintTypeWOff:
		attr.Value = "wOff"
	}
	return attr, nil
}

func (e *ST_ConstraintType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "alignOff":
		*e = 2
	case "begMarg":
		*e = 3
	case "bendDist":
		*e = 4
	case "begPad":
		*e = 5
	case "b":
		*e = 6
	case "bMarg":
		*e = 7
	case "bOff":
		*e = 8
	case "ctrX":
		*e = 9
	case "ctrXOff":
		*e = 10
	case "ctrY":
		*e = 11
	case "ctrYOff":
		*e = 12
	case "connDist":
		*e = 13
	case "diam":
		*e = 14
	case "endMarg":
		*e = 15
	case "endPad":
		*e = 16
	case "h":
		*e = 17
	case "hArH":
		*e = 18
	case "hOff":
		*e = 19
	case "l":
		*e = 20
	case "lMarg":
		*e = 21
	case "lOff":
		*e = 22
	case "r":
		*e = 23
	case "rMarg":
		*e = 24
	case "rOff":
		*e = 25
	case "primFontSz":
		*e = 26
	case "pyraAcctRatio":
		*e = 27
	case "secFontSz":
		*e = 28
	case "sibSp":
		*e = 29
	case "secSibSp":
		*e = 30
	case "sp":
		*e = 31
	case "stemThick":
		*e = 32
	case "t":
		*e = 33
	case "tMarg":
		*e = 34
	case "tOff":
		*e = 35
	case "userA":
		*e = 36
	case "userB":
		*e = 37
	case "userC":
		*e = 38
	case "userD":
		*e = 39
	case "userE":
		*e = 40
	case "userF":
		*e = 41
	case "userG":
		*e = 42
	case "userH":
		*e = 43
	case "userI":
		*e = 44
	case "userJ":
		*e = 45
	case "userK":
		*e = 46
	case "userL":
		*e = 47
	case "userM":
		*e = 48
	case "userN":
		*e = 49
	case "userO":
		*e = 50
	case "userP":
		*e = 51
	case "userQ":
		*e = 52
	case "userR":
		*e = 53
	case "userS":
		*e = 54
	case "userT":
		*e = 55
	case "userU":
		*e = 56
	case "userV":
		*e = 57
	case "userW":
		*e = 58
	case "userX":
		*e = 59
	case "userY":
		*e = 60
	case "userZ":
		*e = 61
	case "w":
		*e = 62
	case "wArH":
		*e = 63
	case "wOff":
		*e = 64
	}
	return nil
}

func (m ST_ConstraintType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ConstraintType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "alignOff":
			*m = 2
		case "begMarg":
			*m = 3
		case "bendDist":
			*m = 4
		case "begPad":
			*m = 5
		case "b":
			*m = 6
		case "bMarg":
			*m = 7
		case "bOff":
			*m = 8
		case "ctrX":
			*m = 9
		case "ctrXOff":
			*m = 10
		case "ctrY":
			*m = 11
		case "ctrYOff":
			*m = 12
		case "connDist":
			*m = 13
		case "diam":
			*m = 14
		case "endMarg":
			*m = 15
		case "endPad":
			*m = 16
		case "h":
			*m = 17
		case "hArH":
			*m = 18
		case "hOff":
			*m = 19
		case "l":
			*m = 20
		case "lMarg":
			*m = 21
		case "lOff":
			*m = 22
		case "r":
			*m = 23
		case "rMarg":
			*m = 24
		case "rOff":
			*m = 25
		case "primFontSz":
			*m = 26
		case "pyraAcctRatio":
			*m = 27
		case "secFontSz":
			*m = 28
		case "sibSp":
			*m = 29
		case "secSibSp":
			*m = 30
		case "sp":
			*m = 31
		case "stemThick":
			*m = 32
		case "t":
			*m = 33
		case "tMarg":
			*m = 34
		case "tOff":
			*m = 35
		case "userA":
			*m = 36
		case "userB":
			*m = 37
		case "userC":
			*m = 38
		case "userD":
			*m = 39
		case "userE":
			*m = 40
		case "userF":
			*m = 41
		case "userG":
			*m = 42
		case "userH":
			*m = 43
		case "userI":
			*m = 44
		case "userJ":
			*m = 45
		case "userK":
			*m = 46
		case "userL":
			*m = 47
		case "userM":
			*m = 48
		case "userN":
			*m = 49
		case "userO":
			*m = 50
		case "userP":
			*m = 51
		case "userQ":
			*m = 52
		case "userR":
			*m = 53
		case "userS":
			*m = 54
		case "userT":
			*m = 55
		case "userU":
			*m = 56
		case "userV":
			*m = 57
		case "userW":
			*m = 58
		case "userX":
			*m = 59
		case "userY":
			*m = 60
		case "userZ":
			*m = 61
		case "w":
			*m = 62
		case "wArH":
			*m = 63
		case "wOff":
			*m = 64
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ConstraintType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "alignOff"
	case 3:
		return "begMarg"
	case 4:
		return "bendDist"
	case 5:
		return "begPad"
	case 6:
		return "b"
	case 7:
		return "bMarg"
	case 8:
		return "bOff"
	case 9:
		return "ctrX"
	case 10:
		return "ctrXOff"
	case 11:
		return "ctrY"
	case 12:
		return "ctrYOff"
	case 13:
		return "connDist"
	case 14:
		return "diam"
	case 15:
		return "endMarg"
	case 16:
		return "endPad"
	case 17:
		return "h"
	case 18:
		return "hArH"
	case 19:
		return "hOff"
	case 20:
		return "l"
	case 21:
		return "lMarg"
	case 22:
		return "lOff"
	case 23:
		return "r"
	case 24:
		return "rMarg"
	case 25:
		return "rOff"
	case 26:
		return "primFontSz"
	case 27:
		return "pyraAcctRatio"
	case 28:
		return "secFontSz"
	case 29:
		return "sibSp"
	case 30:
		return "secSibSp"
	case 31:
		return "sp"
	case 32:
		return "stemThick"
	case 33:
		return "t"
	case 34:
		return "tMarg"
	case 35:
		return "tOff"
	case 36:
		return "userA"
	case 37:
		return "userB"
	case 38:
		return "userC"
	case 39:
		return "userD"
	case 40:
		return "userE"
	case 41:
		return "userF"
	case 42:
		return "userG"
	case 43:
		return "userH"
	case 44:
		return "userI"
	case 45:
		return "userJ"
	case 46:
		return "userK"
	case 47:
		return "userL"
	case 48:
		return "userM"
	case 49:
		return "userN"
	case 50:
		return "userO"
	case 51:
		return "userP"
	case 52:
		return "userQ"
	case 53:
		return "userR"
	case 54:
		return "userS"
	case 55:
		return "userT"
	case 56:
		return "userU"
	case 57:
		return "userV"
	case 58:
		return "userW"
	case 59:
		return "userX"
	case 60:
		return "userY"
	case 61:
		return "userZ"
	case 62:
		return "w"
	case 63:
		return "wArH"
	case 64:
		return "wOff"
	}
	return ""
}

func (m ST_ConstraintType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ConstraintType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ConstraintRelationship byte

const (
	ST_ConstraintRelationshipUnset ST_ConstraintRelationship = 0
	ST_ConstraintRelationshipSelf  ST_ConstraintRelationship = 1
	ST_ConstraintRelationshipCh    ST_ConstraintRelationship = 2
	ST_ConstraintRelationshipDes   ST_ConstraintRelationship = 3
)

func (e ST_ConstraintRelationship) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ConstraintRelationshipUnset:
		attr.Value = ""
	case ST_ConstraintRelationshipSelf:
		attr.Value = "self"
	case ST_ConstraintRelationshipCh:
		attr.Value = "ch"
	case ST_ConstraintRelationshipDes:
		attr.Value = "des"
	}
	return attr, nil
}

func (e *ST_ConstraintRelationship) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "self":
		*e = 1
	case "ch":
		*e = 2
	case "des":
		*e = 3
	}
	return nil
}

func (m ST_ConstraintRelationship) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ConstraintRelationship) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "self":
			*m = 1
		case "ch":
			*m = 2
		case "des":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ConstraintRelationship) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "self"
	case 2:
		return "ch"
	case 3:
		return "des"
	}
	return ""
}

func (m ST_ConstraintRelationship) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ConstraintRelationship) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ElementType byte

const (
	ST_ElementTypeUnset    ST_ElementType = 0
	ST_ElementTypeAll      ST_ElementType = 1
	ST_ElementTypeDoc      ST_ElementType = 2
	ST_ElementTypeNode     ST_ElementType = 3
	ST_ElementTypeNorm     ST_ElementType = 4
	ST_ElementTypeNonNorm  ST_ElementType = 5
	ST_ElementTypeAsst     ST_ElementType = 6
	ST_ElementTypeNonAsst  ST_ElementType = 7
	ST_ElementTypeParTrans ST_ElementType = 8
	ST_ElementTypePres     ST_ElementType = 9
	ST_ElementTypeSibTrans ST_ElementType = 10
)

func (e ST_ElementType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ElementTypeUnset:
		attr.Value = ""
	case ST_ElementTypeAll:
		attr.Value = "all"
	case ST_ElementTypeDoc:
		attr.Value = "doc"
	case ST_ElementTypeNode:
		attr.Value = "node"
	case ST_ElementTypeNorm:
		attr.Value = "norm"
	case ST_ElementTypeNonNorm:
		attr.Value = "nonNorm"
	case ST_ElementTypeAsst:
		attr.Value = "asst"
	case ST_ElementTypeNonAsst:
		attr.Value = "nonAsst"
	case ST_ElementTypeParTrans:
		attr.Value = "parTrans"
	case ST_ElementTypePres:
		attr.Value = "pres"
	case ST_ElementTypeSibTrans:
		attr.Value = "sibTrans"
	}
	return attr, nil
}

func (e *ST_ElementType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "all":
		*e = 1
	case "doc":
		*e = 2
	case "node":
		*e = 3
	case "norm":
		*e = 4
	case "nonNorm":
		*e = 5
	case "asst":
		*e = 6
	case "nonAsst":
		*e = 7
	case "parTrans":
		*e = 8
	case "pres":
		*e = 9
	case "sibTrans":
		*e = 10
	}
	return nil
}

func (m ST_ElementType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ElementType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "all":
			*m = 1
		case "doc":
			*m = 2
		case "node":
			*m = 3
		case "norm":
			*m = 4
		case "nonNorm":
			*m = 5
		case "asst":
			*m = 6
		case "nonAsst":
			*m = 7
		case "parTrans":
			*m = 8
		case "pres":
			*m = 9
		case "sibTrans":
			*m = 10
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ElementType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "all"
	case 2:
		return "doc"
	case 3:
		return "node"
	case 4:
		return "norm"
	case 5:
		return "nonNorm"
	case 6:
		return "asst"
	case 7:
		return "nonAsst"
	case 8:
		return "parTrans"
	case 9:
		return "pres"
	case 10:
		return "sibTrans"
	}
	return ""
}

func (m ST_ElementType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ElementType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ParameterId byte

const (
	ST_ParameterIdUnset            ST_ParameterId = 0
	ST_ParameterIdHorzAlign        ST_ParameterId = 1
	ST_ParameterIdVertAlign        ST_ParameterId = 2
	ST_ParameterIdChDir            ST_ParameterId = 3
	ST_ParameterIdChAlign          ST_ParameterId = 4
	ST_ParameterIdSecChAlign       ST_ParameterId = 5
	ST_ParameterIdLinDir           ST_ParameterId = 6
	ST_ParameterIdSecLinDir        ST_ParameterId = 7
	ST_ParameterIdStElem           ST_ParameterId = 8
	ST_ParameterIdBendPt           ST_ParameterId = 9
	ST_ParameterIdConnRout         ST_ParameterId = 10
	ST_ParameterIdBegSty           ST_ParameterId = 11
	ST_ParameterIdEndSty           ST_ParameterId = 12
	ST_ParameterIdDim              ST_ParameterId = 13
	ST_ParameterIdRotPath          ST_ParameterId = 14
	ST_ParameterIdCtrShpMap        ST_ParameterId = 15
	ST_ParameterIdNodeHorzAlign    ST_ParameterId = 16
	ST_ParameterIdNodeVertAlign    ST_ParameterId = 17
	ST_ParameterIdFallback         ST_ParameterId = 18
	ST_ParameterIdTxDir            ST_ParameterId = 19
	ST_ParameterIdPyraAcctPos      ST_ParameterId = 20
	ST_ParameterIdPyraAcctTxMar    ST_ParameterId = 21
	ST_ParameterIdTxBlDir          ST_ParameterId = 22
	ST_ParameterIdTxAnchorHorz     ST_ParameterId = 23
	ST_ParameterIdTxAnchorVert     ST_ParameterId = 24
	ST_ParameterIdTxAnchorHorzCh   ST_ParameterId = 25
	ST_ParameterIdTxAnchorVertCh   ST_ParameterId = 26
	ST_ParameterIdParTxLTRAlign    ST_ParameterId = 27
	ST_ParameterIdParTxRTLAlign    ST_ParameterId = 28
	ST_ParameterIdShpTxLTRAlignCh  ST_ParameterId = 29
	ST_ParameterIdShpTxRTLAlignCh  ST_ParameterId = 30
	ST_ParameterIdAutoTxRot        ST_ParameterId = 31
	ST_ParameterIdGrDir            ST_ParameterId = 32
	ST_ParameterIdFlowDir          ST_ParameterId = 33
	ST_ParameterIdContDir          ST_ParameterId = 34
	ST_ParameterIdBkpt             ST_ParameterId = 35
	ST_ParameterIdOff              ST_ParameterId = 36
	ST_ParameterIdHierAlign        ST_ParameterId = 37
	ST_ParameterIdBkPtFixedVal     ST_ParameterId = 38
	ST_ParameterIdStBulletLvl      ST_ParameterId = 39
	ST_ParameterIdStAng            ST_ParameterId = 40
	ST_ParameterIdSpanAng          ST_ParameterId = 41
	ST_ParameterIdAr               ST_ParameterId = 42
	ST_ParameterIdLnSpPar          ST_ParameterId = 43
	ST_ParameterIdLnSpAfParP       ST_ParameterId = 44
	ST_ParameterIdLnSpCh           ST_ParameterId = 45
	ST_ParameterIdLnSpAfChP        ST_ParameterId = 46
	ST_ParameterIdRtShortDist      ST_ParameterId = 47
	ST_ParameterIdAlignTx          ST_ParameterId = 48
	ST_ParameterIdPyraLvlNode      ST_ParameterId = 49
	ST_ParameterIdPyraAcctBkgdNode ST_ParameterId = 50
	ST_ParameterIdPyraAcctTxNode   ST_ParameterId = 51
	ST_ParameterIdSrcNode          ST_ParameterId = 52
	ST_ParameterIdDstNode          ST_ParameterId = 53
	ST_ParameterIdBegPts           ST_ParameterId = 54
	ST_ParameterIdEndPts           ST_ParameterId = 55
)

func (e ST_ParameterId) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ParameterIdUnset:
		attr.Value = ""
	case ST_ParameterIdHorzAlign:
		attr.Value = "horzAlign"
	case ST_ParameterIdVertAlign:
		attr.Value = "vertAlign"
	case ST_ParameterIdChDir:
		attr.Value = "chDir"
	case ST_ParameterIdChAlign:
		attr.Value = "chAlign"
	case ST_ParameterIdSecChAlign:
		attr.Value = "secChAlign"
	case ST_ParameterIdLinDir:
		attr.Value = "linDir"
	case ST_ParameterIdSecLinDir:
		attr.Value = "secLinDir"
	case ST_ParameterIdStElem:
		attr.Value = "stElem"
	case ST_ParameterIdBendPt:
		attr.Value = "bendPt"
	case ST_ParameterIdConnRout:
		attr.Value = "connRout"
	case ST_ParameterIdBegSty:
		attr.Value = "begSty"
	case ST_ParameterIdEndSty:
		attr.Value = "endSty"
	case ST_ParameterIdDim:
		attr.Value = "dim"
	case ST_ParameterIdRotPath:
		attr.Value = "rotPath"
	case ST_ParameterIdCtrShpMap:
		attr.Value = "ctrShpMap"
	case ST_ParameterIdNodeHorzAlign:
		attr.Value = "nodeHorzAlign"
	case ST_ParameterIdNodeVertAlign:
		attr.Value = "nodeVertAlign"
	case ST_ParameterIdFallback:
		attr.Value = "fallback"
	case ST_ParameterIdTxDir:
		attr.Value = "txDir"
	case ST_ParameterIdPyraAcctPos:
		attr.Value = "pyraAcctPos"
	case ST_ParameterIdPyraAcctTxMar:
		attr.Value = "pyraAcctTxMar"
	case ST_ParameterIdTxBlDir:
		attr.Value = "txBlDir"
	case ST_ParameterIdTxAnchorHorz:
		attr.Value = "txAnchorHorz"
	case ST_ParameterIdTxAnchorVert:
		attr.Value = "txAnchorVert"
	case ST_ParameterIdTxAnchorHorzCh:
		attr.Value = "txAnchorHorzCh"
	case ST_ParameterIdTxAnchorVertCh:
		attr.Value = "txAnchorVertCh"
	case ST_ParameterIdParTxLTRAlign:
		attr.Value = "parTxLTRAlign"
	case ST_ParameterIdParTxRTLAlign:
		attr.Value = "parTxRTLAlign"
	case ST_ParameterIdShpTxLTRAlignCh:
		attr.Value = "shpTxLTRAlignCh"
	case ST_ParameterIdShpTxRTLAlignCh:
		attr.Value = "shpTxRTLAlignCh"
	case ST_ParameterIdAutoTxRot:
		attr.Value = "autoTxRot"
	case ST_ParameterIdGrDir:
		attr.Value = "grDir"
	case ST_ParameterIdFlowDir:
		attr.Value = "flowDir"
	case ST_ParameterIdContDir:
		attr.Value = "contDir"
	case ST_ParameterIdBkpt:
		attr.Value = "bkpt"
	case ST_ParameterIdOff:
		attr.Value = "off"
	case ST_ParameterIdHierAlign:
		attr.Value = "hierAlign"
	case ST_ParameterIdBkPtFixedVal:
		attr.Value = "bkPtFixedVal"
	case ST_ParameterIdStBulletLvl:
		attr.Value = "stBulletLvl"
	case ST_ParameterIdStAng:
		attr.Value = "stAng"
	case ST_ParameterIdSpanAng:
		attr.Value = "spanAng"
	case ST_ParameterIdAr:
		attr.Value = "ar"
	case ST_ParameterIdLnSpPar:
		attr.Value = "lnSpPar"
	case ST_ParameterIdLnSpAfParP:
		attr.Value = "lnSpAfParP"
	case ST_ParameterIdLnSpCh:
		attr.Value = "lnSpCh"
	case ST_ParameterIdLnSpAfChP:
		attr.Value = "lnSpAfChP"
	case ST_ParameterIdRtShortDist:
		attr.Value = "rtShortDist"
	case ST_ParameterIdAlignTx:
		attr.Value = "alignTx"
	case ST_ParameterIdPyraLvlNode:
		attr.Value = "pyraLvlNode"
	case ST_ParameterIdPyraAcctBkgdNode:
		attr.Value = "pyraAcctBkgdNode"
	case ST_ParameterIdPyraAcctTxNode:
		attr.Value = "pyraAcctTxNode"
	case ST_ParameterIdSrcNode:
		attr.Value = "srcNode"
	case ST_ParameterIdDstNode:
		attr.Value = "dstNode"
	case ST_ParameterIdBegPts:
		attr.Value = "begPts"
	case ST_ParameterIdEndPts:
		attr.Value = "endPts"
	}
	return attr, nil
}

func (e *ST_ParameterId) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "horzAlign":
		*e = 1
	case "vertAlign":
		*e = 2
	case "chDir":
		*e = 3
	case "chAlign":
		*e = 4
	case "secChAlign":
		*e = 5
	case "linDir":
		*e = 6
	case "secLinDir":
		*e = 7
	case "stElem":
		*e = 8
	case "bendPt":
		*e = 9
	case "connRout":
		*e = 10
	case "begSty":
		*e = 11
	case "endSty":
		*e = 12
	case "dim":
		*e = 13
	case "rotPath":
		*e = 14
	case "ctrShpMap":
		*e = 15
	case "nodeHorzAlign":
		*e = 16
	case "nodeVertAlign":
		*e = 17
	case "fallback":
		*e = 18
	case "txDir":
		*e = 19
	case "pyraAcctPos":
		*e = 20
	case "pyraAcctTxMar":
		*e = 21
	case "txBlDir":
		*e = 22
	case "txAnchorHorz":
		*e = 23
	case "txAnchorVert":
		*e = 24
	case "txAnchorHorzCh":
		*e = 25
	case "txAnchorVertCh":
		*e = 26
	case "parTxLTRAlign":
		*e = 27
	case "parTxRTLAlign":
		*e = 28
	case "shpTxLTRAlignCh":
		*e = 29
	case "shpTxRTLAlignCh":
		*e = 30
	case "autoTxRot":
		*e = 31
	case "grDir":
		*e = 32
	case "flowDir":
		*e = 33
	case "contDir":
		*e = 34
	case "bkpt":
		*e = 35
	case "off":
		*e = 36
	case "hierAlign":
		*e = 37
	case "bkPtFixedVal":
		*e = 38
	case "stBulletLvl":
		*e = 39
	case "stAng":
		*e = 40
	case "spanAng":
		*e = 41
	case "ar":
		*e = 42
	case "lnSpPar":
		*e = 43
	case "lnSpAfParP":
		*e = 44
	case "lnSpCh":
		*e = 45
	case "lnSpAfChP":
		*e = 46
	case "rtShortDist":
		*e = 47
	case "alignTx":
		*e = 48
	case "pyraLvlNode":
		*e = 49
	case "pyraAcctBkgdNode":
		*e = 50
	case "pyraAcctTxNode":
		*e = 51
	case "srcNode":
		*e = 52
	case "dstNode":
		*e = 53
	case "begPts":
		*e = 54
	case "endPts":
		*e = 55
	}
	return nil
}

func (m ST_ParameterId) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ParameterId) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "horzAlign":
			*m = 1
		case "vertAlign":
			*m = 2
		case "chDir":
			*m = 3
		case "chAlign":
			*m = 4
		case "secChAlign":
			*m = 5
		case "linDir":
			*m = 6
		case "secLinDir":
			*m = 7
		case "stElem":
			*m = 8
		case "bendPt":
			*m = 9
		case "connRout":
			*m = 10
		case "begSty":
			*m = 11
		case "endSty":
			*m = 12
		case "dim":
			*m = 13
		case "rotPath":
			*m = 14
		case "ctrShpMap":
			*m = 15
		case "nodeHorzAlign":
			*m = 16
		case "nodeVertAlign":
			*m = 17
		case "fallback":
			*m = 18
		case "txDir":
			*m = 19
		case "pyraAcctPos":
			*m = 20
		case "pyraAcctTxMar":
			*m = 21
		case "txBlDir":
			*m = 22
		case "txAnchorHorz":
			*m = 23
		case "txAnchorVert":
			*m = 24
		case "txAnchorHorzCh":
			*m = 25
		case "txAnchorVertCh":
			*m = 26
		case "parTxLTRAlign":
			*m = 27
		case "parTxRTLAlign":
			*m = 28
		case "shpTxLTRAlignCh":
			*m = 29
		case "shpTxRTLAlignCh":
			*m = 30
		case "autoTxRot":
			*m = 31
		case "grDir":
			*m = 32
		case "flowDir":
			*m = 33
		case "contDir":
			*m = 34
		case "bkpt":
			*m = 35
		case "off":
			*m = 36
		case "hierAlign":
			*m = 37
		case "bkPtFixedVal":
			*m = 38
		case "stBulletLvl":
			*m = 39
		case "stAng":
			*m = 40
		case "spanAng":
			*m = 41
		case "ar":
			*m = 42
		case "lnSpPar":
			*m = 43
		case "lnSpAfParP":
			*m = 44
		case "lnSpCh":
			*m = 45
		case "lnSpAfChP":
			*m = 46
		case "rtShortDist":
			*m = 47
		case "alignTx":
			*m = 48
		case "pyraLvlNode":
			*m = 49
		case "pyraAcctBkgdNode":
			*m = 50
		case "pyraAcctTxNode":
			*m = 51
		case "srcNode":
			*m = 52
		case "dstNode":
			*m = 53
		case "begPts":
			*m = 54
		case "endPts":
			*m = 55
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ParameterId) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "horzAlign"
	case 2:
		return "vertAlign"
	case 3:
		return "chDir"
	case 4:
		return "chAlign"
	case 5:
		return "secChAlign"
	case 6:
		return "linDir"
	case 7:
		return "secLinDir"
	case 8:
		return "stElem"
	case 9:
		return "bendPt"
	case 10:
		return "connRout"
	case 11:
		return "begSty"
	case 12:
		return "endSty"
	case 13:
		return "dim"
	case 14:
		return "rotPath"
	case 15:
		return "ctrShpMap"
	case 16:
		return "nodeHorzAlign"
	case 17:
		return "nodeVertAlign"
	case 18:
		return "fallback"
	case 19:
		return "txDir"
	case 20:
		return "pyraAcctPos"
	case 21:
		return "pyraAcctTxMar"
	case 22:
		return "txBlDir"
	case 23:
		return "txAnchorHorz"
	case 24:
		return "txAnchorVert"
	case 25:
		return "txAnchorHorzCh"
	case 26:
		return "txAnchorVertCh"
	case 27:
		return "parTxLTRAlign"
	case 28:
		return "parTxRTLAlign"
	case 29:
		return "shpTxLTRAlignCh"
	case 30:
		return "shpTxRTLAlignCh"
	case 31:
		return "autoTxRot"
	case 32:
		return "grDir"
	case 33:
		return "flowDir"
	case 34:
		return "contDir"
	case 35:
		return "bkpt"
	case 36:
		return "off"
	case 37:
		return "hierAlign"
	case 38:
		return "bkPtFixedVal"
	case 39:
		return "stBulletLvl"
	case 40:
		return "stAng"
	case 41:
		return "spanAng"
	case 42:
		return "ar"
	case 43:
		return "lnSpPar"
	case 44:
		return "lnSpAfParP"
	case 45:
		return "lnSpCh"
	case 46:
		return "lnSpAfChP"
	case 47:
		return "rtShortDist"
	case 48:
		return "alignTx"
	case 49:
		return "pyraLvlNode"
	case 50:
		return "pyraAcctBkgdNode"
	case 51:
		return "pyraAcctTxNode"
	case 52:
		return "srcNode"
	case 53:
		return "dstNode"
	case 54:
		return "begPts"
	case 55:
		return "endPts"
	}
	return ""
}

func (m ST_ParameterId) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ParameterId) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FunctionType byte

const (
	ST_FunctionTypeUnset    ST_FunctionType = 0
	ST_FunctionTypeCnt      ST_FunctionType = 1
	ST_FunctionTypePos      ST_FunctionType = 2
	ST_FunctionTypeRevPos   ST_FunctionType = 3
	ST_FunctionTypePosEven  ST_FunctionType = 4
	ST_FunctionTypePosOdd   ST_FunctionType = 5
	ST_FunctionTypeVar      ST_FunctionType = 6
	ST_FunctionTypeDepth    ST_FunctionType = 7
	ST_FunctionTypeMaxDepth ST_FunctionType = 8
)

func (e ST_FunctionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FunctionTypeUnset:
		attr.Value = ""
	case ST_FunctionTypeCnt:
		attr.Value = "cnt"
	case ST_FunctionTypePos:
		attr.Value = "pos"
	case ST_FunctionTypeRevPos:
		attr.Value = "revPos"
	case ST_FunctionTypePosEven:
		attr.Value = "posEven"
	case ST_FunctionTypePosOdd:
		attr.Value = "posOdd"
	case ST_FunctionTypeVar:
		attr.Value = "var"
	case ST_FunctionTypeDepth:
		attr.Value = "depth"
	case ST_FunctionTypeMaxDepth:
		attr.Value = "maxDepth"
	}
	return attr, nil
}

func (e *ST_FunctionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "cnt":
		*e = 1
	case "pos":
		*e = 2
	case "revPos":
		*e = 3
	case "posEven":
		*e = 4
	case "posOdd":
		*e = 5
	case "var":
		*e = 6
	case "depth":
		*e = 7
	case "maxDepth":
		*e = 8
	}
	return nil
}

func (m ST_FunctionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FunctionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "cnt":
			*m = 1
		case "pos":
			*m = 2
		case "revPos":
			*m = 3
		case "posEven":
			*m = 4
		case "posOdd":
			*m = 5
		case "var":
			*m = 6
		case "depth":
			*m = 7
		case "maxDepth":
			*m = 8
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FunctionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "cnt"
	case 2:
		return "pos"
	case 3:
		return "revPos"
	case 4:
		return "posEven"
	case 5:
		return "posOdd"
	case 6:
		return "var"
	case 7:
		return "depth"
	case 8:
		return "maxDepth"
	}
	return ""
}

func (m ST_FunctionType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FunctionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FunctionOperator byte

const (
	ST_FunctionOperatorUnset ST_FunctionOperator = 0
	ST_FunctionOperatorEqu   ST_FunctionOperator = 1
	ST_FunctionOperatorNeq   ST_FunctionOperator = 2
	ST_FunctionOperatorGt    ST_FunctionOperator = 3
	ST_FunctionOperatorLt    ST_FunctionOperator = 4
	ST_FunctionOperatorGte   ST_FunctionOperator = 5
	ST_FunctionOperatorLte   ST_FunctionOperator = 6
)

func (e ST_FunctionOperator) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FunctionOperatorUnset:
		attr.Value = ""
	case ST_FunctionOperatorEqu:
		attr.Value = "equ"
	case ST_FunctionOperatorNeq:
		attr.Value = "neq"
	case ST_FunctionOperatorGt:
		attr.Value = "gt"
	case ST_FunctionOperatorLt:
		attr.Value = "lt"
	case ST_FunctionOperatorGte:
		attr.Value = "gte"
	case ST_FunctionOperatorLte:
		attr.Value = "lte"
	}
	return attr, nil
}

func (e *ST_FunctionOperator) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "equ":
		*e = 1
	case "neq":
		*e = 2
	case "gt":
		*e = 3
	case "lt":
		*e = 4
	case "gte":
		*e = 5
	case "lte":
		*e = 6
	}
	return nil
}

func (m ST_FunctionOperator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FunctionOperator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "equ":
			*m = 1
		case "neq":
			*m = 2
		case "gt":
			*m = 3
		case "lt":
			*m = 4
		case "gte":
			*m = 5
		case "lte":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FunctionOperator) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "equ"
	case 2:
		return "neq"
	case 3:
		return "gt"
	case 4:
		return "lt"
	case 5:
		return "gte"
	case 6:
		return "lte"
	}
	return ""
}

func (m ST_FunctionOperator) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FunctionOperator) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DiagramHorizontalAlignment byte

const (
	ST_DiagramHorizontalAlignmentUnset ST_DiagramHorizontalAlignment = 0
	ST_DiagramHorizontalAlignmentL     ST_DiagramHorizontalAlignment = 1
	ST_DiagramHorizontalAlignmentCtr   ST_DiagramHorizontalAlignment = 2
	ST_DiagramHorizontalAlignmentR     ST_DiagramHorizontalAlignment = 3
	ST_DiagramHorizontalAlignmentNone  ST_DiagramHorizontalAlignment = 4
)

func (e ST_DiagramHorizontalAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DiagramHorizontalAlignmentUnset:
		attr.Value = ""
	case ST_DiagramHorizontalAlignmentL:
		attr.Value = "l"
	case ST_DiagramHorizontalAlignmentCtr:
		attr.Value = "ctr"
	case ST_DiagramHorizontalAlignmentR:
		attr.Value = "r"
	case ST_DiagramHorizontalAlignmentNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_DiagramHorizontalAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "l":
		*e = 1
	case "ctr":
		*e = 2
	case "r":
		*e = 3
	case "none":
		*e = 4
	}
	return nil
}

func (m ST_DiagramHorizontalAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DiagramHorizontalAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "l":
			*m = 1
		case "ctr":
			*m = 2
		case "r":
			*m = 3
		case "none":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DiagramHorizontalAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "l"
	case 2:
		return "ctr"
	case 3:
		return "r"
	case 4:
		return "none"
	}
	return ""
}

func (m ST_DiagramHorizontalAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DiagramHorizontalAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VerticalAlignment byte

const (
	ST_VerticalAlignmentUnset ST_VerticalAlignment = 0
	ST_VerticalAlignmentT     ST_VerticalAlignment = 1
	ST_VerticalAlignmentMid   ST_VerticalAlignment = 2
	ST_VerticalAlignmentB     ST_VerticalAlignment = 3
	ST_VerticalAlignmentNone  ST_VerticalAlignment = 4
)

func (e ST_VerticalAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VerticalAlignmentUnset:
		attr.Value = ""
	case ST_VerticalAlignmentT:
		attr.Value = "t"
	case ST_VerticalAlignmentMid:
		attr.Value = "mid"
	case ST_VerticalAlignmentB:
		attr.Value = "b"
	case ST_VerticalAlignmentNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_VerticalAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "mid":
		*e = 2
	case "b":
		*e = 3
	case "none":
		*e = 4
	}
	return nil
}

func (m ST_VerticalAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VerticalAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "t":
			*m = 1
		case "mid":
			*m = 2
		case "b":
			*m = 3
		case "none":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_VerticalAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "mid"
	case 3:
		return "b"
	case 4:
		return "none"
	}
	return ""
}

func (m ST_VerticalAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VerticalAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ChildDirection byte

const (
	ST_ChildDirectionUnset ST_ChildDirection = 0
	ST_ChildDirectionHorz  ST_ChildDirection = 1
	ST_ChildDirectionVert  ST_ChildDirection = 2
)

func (e ST_ChildDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ChildDirectionUnset:
		attr.Value = ""
	case ST_ChildDirectionHorz:
		attr.Value = "horz"
	case ST_ChildDirectionVert:
		attr.Value = "vert"
	}
	return attr, nil
}

func (e *ST_ChildDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "horz":
		*e = 1
	case "vert":
		*e = 2
	}
	return nil
}

func (m ST_ChildDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ChildDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "horz":
			*m = 1
		case "vert":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ChildDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "horz"
	case 2:
		return "vert"
	}
	return ""
}

func (m ST_ChildDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ChildDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ChildAlignment byte

const (
	ST_ChildAlignmentUnset ST_ChildAlignment = 0
	ST_ChildAlignmentT     ST_ChildAlignment = 1
	ST_ChildAlignmentB     ST_ChildAlignment = 2
	ST_ChildAlignmentL     ST_ChildAlignment = 3
	ST_ChildAlignmentR     ST_ChildAlignment = 4
)

func (e ST_ChildAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ChildAlignmentUnset:
		attr.Value = ""
	case ST_ChildAlignmentT:
		attr.Value = "t"
	case ST_ChildAlignmentB:
		attr.Value = "b"
	case ST_ChildAlignmentL:
		attr.Value = "l"
	case ST_ChildAlignmentR:
		attr.Value = "r"
	}
	return attr, nil
}

func (e *ST_ChildAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "b":
		*e = 2
	case "l":
		*e = 3
	case "r":
		*e = 4
	}
	return nil
}

func (m ST_ChildAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ChildAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "t":
			*m = 1
		case "b":
			*m = 2
		case "l":
			*m = 3
		case "r":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ChildAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "b"
	case 3:
		return "l"
	case 4:
		return "r"
	}
	return ""
}

func (m ST_ChildAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ChildAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SecondaryChildAlignment byte

const (
	ST_SecondaryChildAlignmentUnset ST_SecondaryChildAlignment = 0
	ST_SecondaryChildAlignmentNone  ST_SecondaryChildAlignment = 1
	ST_SecondaryChildAlignmentT     ST_SecondaryChildAlignment = 2
	ST_SecondaryChildAlignmentB     ST_SecondaryChildAlignment = 3
	ST_SecondaryChildAlignmentL     ST_SecondaryChildAlignment = 4
	ST_SecondaryChildAlignmentR     ST_SecondaryChildAlignment = 5
)

func (e ST_SecondaryChildAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SecondaryChildAlignmentUnset:
		attr.Value = ""
	case ST_SecondaryChildAlignmentNone:
		attr.Value = "none"
	case ST_SecondaryChildAlignmentT:
		attr.Value = "t"
	case ST_SecondaryChildAlignmentB:
		attr.Value = "b"
	case ST_SecondaryChildAlignmentL:
		attr.Value = "l"
	case ST_SecondaryChildAlignmentR:
		attr.Value = "r"
	}
	return attr, nil
}

func (e *ST_SecondaryChildAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "t":
		*e = 2
	case "b":
		*e = 3
	case "l":
		*e = 4
	case "r":
		*e = 5
	}
	return nil
}

func (m ST_SecondaryChildAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SecondaryChildAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "t":
			*m = 2
		case "b":
			*m = 3
		case "l":
			*m = 4
		case "r":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_SecondaryChildAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "t"
	case 3:
		return "b"
	case 4:
		return "l"
	case 5:
		return "r"
	}
	return ""
}

func (m ST_SecondaryChildAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SecondaryChildAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LinearDirection byte

const (
	ST_LinearDirectionUnset ST_LinearDirection = 0
	ST_LinearDirectionFromL ST_LinearDirection = 1
	ST_LinearDirectionFromR ST_LinearDirection = 2
	ST_LinearDirectionFromT ST_LinearDirection = 3
	ST_LinearDirectionFromB ST_LinearDirection = 4
)

func (e ST_LinearDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LinearDirectionUnset:
		attr.Value = ""
	case ST_LinearDirectionFromL:
		attr.Value = "fromL"
	case ST_LinearDirectionFromR:
		attr.Value = "fromR"
	case ST_LinearDirectionFromT:
		attr.Value = "fromT"
	case ST_LinearDirectionFromB:
		attr.Value = "fromB"
	}
	return attr, nil
}

func (e *ST_LinearDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "fromL":
		*e = 1
	case "fromR":
		*e = 2
	case "fromT":
		*e = 3
	case "fromB":
		*e = 4
	}
	return nil
}

func (m ST_LinearDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_LinearDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "fromL":
			*m = 1
		case "fromR":
			*m = 2
		case "fromT":
			*m = 3
		case "fromB":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_LinearDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "fromL"
	case 2:
		return "fromR"
	case 3:
		return "fromT"
	case 4:
		return "fromB"
	}
	return ""
}

func (m ST_LinearDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LinearDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SecondaryLinearDirection byte

const (
	ST_SecondaryLinearDirectionUnset ST_SecondaryLinearDirection = 0
	ST_SecondaryLinearDirectionNone  ST_SecondaryLinearDirection = 1
	ST_SecondaryLinearDirectionFromL ST_SecondaryLinearDirection = 2
	ST_SecondaryLinearDirectionFromR ST_SecondaryLinearDirection = 3
	ST_SecondaryLinearDirectionFromT ST_SecondaryLinearDirection = 4
	ST_SecondaryLinearDirectionFromB ST_SecondaryLinearDirection = 5
)

func (e ST_SecondaryLinearDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SecondaryLinearDirectionUnset:
		attr.Value = ""
	case ST_SecondaryLinearDirectionNone:
		attr.Value = "none"
	case ST_SecondaryLinearDirectionFromL:
		attr.Value = "fromL"
	case ST_SecondaryLinearDirectionFromR:
		attr.Value = "fromR"
	case ST_SecondaryLinearDirectionFromT:
		attr.Value = "fromT"
	case ST_SecondaryLinearDirectionFromB:
		attr.Value = "fromB"
	}
	return attr, nil
}

func (e *ST_SecondaryLinearDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "fromL":
		*e = 2
	case "fromR":
		*e = 3
	case "fromT":
		*e = 4
	case "fromB":
		*e = 5
	}
	return nil
}

func (m ST_SecondaryLinearDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SecondaryLinearDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "fromL":
			*m = 2
		case "fromR":
			*m = 3
		case "fromT":
			*m = 4
		case "fromB":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_SecondaryLinearDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "fromL"
	case 3:
		return "fromR"
	case 4:
		return "fromT"
	case 5:
		return "fromB"
	}
	return ""
}

func (m ST_SecondaryLinearDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SecondaryLinearDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StartingElement byte

const (
	ST_StartingElementUnset ST_StartingElement = 0
	ST_StartingElementNode  ST_StartingElement = 1
	ST_StartingElementTrans ST_StartingElement = 2
)

func (e ST_StartingElement) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StartingElementUnset:
		attr.Value = ""
	case ST_StartingElementNode:
		attr.Value = "node"
	case ST_StartingElementTrans:
		attr.Value = "trans"
	}
	return attr, nil
}

func (e *ST_StartingElement) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "node":
		*e = 1
	case "trans":
		*e = 2
	}
	return nil
}

func (m ST_StartingElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StartingElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "node":
			*m = 1
		case "trans":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StartingElement) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "node"
	case 2:
		return "trans"
	}
	return ""
}

func (m ST_StartingElement) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StartingElement) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RotationPath byte

const (
	ST_RotationPathUnset     ST_RotationPath = 0
	ST_RotationPathNone      ST_RotationPath = 1
	ST_RotationPathAlongPath ST_RotationPath = 2
)

func (e ST_RotationPath) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RotationPathUnset:
		attr.Value = ""
	case ST_RotationPathNone:
		attr.Value = "none"
	case ST_RotationPathAlongPath:
		attr.Value = "alongPath"
	}
	return attr, nil
}

func (e *ST_RotationPath) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "alongPath":
		*e = 2
	}
	return nil
}

func (m ST_RotationPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_RotationPath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "alongPath":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_RotationPath) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "alongPath"
	}
	return ""
}

func (m ST_RotationPath) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_RotationPath) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CenterShapeMapping byte

const (
	ST_CenterShapeMappingUnset ST_CenterShapeMapping = 0
	ST_CenterShapeMappingNone  ST_CenterShapeMapping = 1
	ST_CenterShapeMappingFNode ST_CenterShapeMapping = 2
)

func (e ST_CenterShapeMapping) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CenterShapeMappingUnset:
		attr.Value = ""
	case ST_CenterShapeMappingNone:
		attr.Value = "none"
	case ST_CenterShapeMappingFNode:
		attr.Value = "fNode"
	}
	return attr, nil
}

func (e *ST_CenterShapeMapping) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "fNode":
		*e = 2
	}
	return nil
}

func (m ST_CenterShapeMapping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CenterShapeMapping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "fNode":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_CenterShapeMapping) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "fNode"
	}
	return ""
}

func (m ST_CenterShapeMapping) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CenterShapeMapping) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BendPoint byte

const (
	ST_BendPointUnset ST_BendPoint = 0
	ST_BendPointBeg   ST_BendPoint = 1
	ST_BendPointDef   ST_BendPoint = 2
	ST_BendPointEnd   ST_BendPoint = 3
)

func (e ST_BendPoint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BendPointUnset:
		attr.Value = ""
	case ST_BendPointBeg:
		attr.Value = "beg"
	case ST_BendPointDef:
		attr.Value = "def"
	case ST_BendPointEnd:
		attr.Value = "end"
	}
	return attr, nil
}

func (e *ST_BendPoint) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "beg":
		*e = 1
	case "def":
		*e = 2
	case "end":
		*e = 3
	}
	return nil
}

func (m ST_BendPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BendPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "beg":
			*m = 1
		case "def":
			*m = 2
		case "end":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_BendPoint) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "beg"
	case 2:
		return "def"
	case 3:
		return "end"
	}
	return ""
}

func (m ST_BendPoint) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BendPoint) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ConnectorRouting byte

const (
	ST_ConnectorRoutingUnset     ST_ConnectorRouting = 0
	ST_ConnectorRoutingStra      ST_ConnectorRouting = 1
	ST_ConnectorRoutingBend      ST_ConnectorRouting = 2
	ST_ConnectorRoutingCurve     ST_ConnectorRouting = 3
	ST_ConnectorRoutingLongCurve ST_ConnectorRouting = 4
)

func (e ST_ConnectorRouting) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ConnectorRoutingUnset:
		attr.Value = ""
	case ST_ConnectorRoutingStra:
		attr.Value = "stra"
	case ST_ConnectorRoutingBend:
		attr.Value = "bend"
	case ST_ConnectorRoutingCurve:
		attr.Value = "curve"
	case ST_ConnectorRoutingLongCurve:
		attr.Value = "longCurve"
	}
	return attr, nil
}

func (e *ST_ConnectorRouting) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "stra":
		*e = 1
	case "bend":
		*e = 2
	case "curve":
		*e = 3
	case "longCurve":
		*e = 4
	}
	return nil
}

func (m ST_ConnectorRouting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ConnectorRouting) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "stra":
			*m = 1
		case "bend":
			*m = 2
		case "curve":
			*m = 3
		case "longCurve":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ConnectorRouting) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "stra"
	case 2:
		return "bend"
	case 3:
		return "curve"
	case 4:
		return "longCurve"
	}
	return ""
}

func (m ST_ConnectorRouting) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ConnectorRouting) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ArrowheadStyle byte

const (
	ST_ArrowheadStyleUnset ST_ArrowheadStyle = 0
	ST_ArrowheadStyleAuto  ST_ArrowheadStyle = 1
	ST_ArrowheadStyleArr   ST_ArrowheadStyle = 2
	ST_ArrowheadStyleNoArr ST_ArrowheadStyle = 3
)

func (e ST_ArrowheadStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ArrowheadStyleUnset:
		attr.Value = ""
	case ST_ArrowheadStyleAuto:
		attr.Value = "auto"
	case ST_ArrowheadStyleArr:
		attr.Value = "arr"
	case ST_ArrowheadStyleNoArr:
		attr.Value = "noArr"
	}
	return attr, nil
}

func (e *ST_ArrowheadStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "arr":
		*e = 2
	case "noArr":
		*e = 3
	}
	return nil
}

func (m ST_ArrowheadStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ArrowheadStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "arr":
			*m = 2
		case "noArr":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ArrowheadStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "arr"
	case 3:
		return "noArr"
	}
	return ""
}

func (m ST_ArrowheadStyle) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ArrowheadStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ConnectorDimension byte

const (
	ST_ConnectorDimensionUnset ST_ConnectorDimension = 0
	ST_ConnectorDimension1D    ST_ConnectorDimension = 1
	ST_ConnectorDimension2D    ST_ConnectorDimension = 2
	ST_ConnectorDimensionCust  ST_ConnectorDimension = 3
)

func (e ST_ConnectorDimension) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ConnectorDimensionUnset:
		attr.Value = ""
	case ST_ConnectorDimension1D:
		attr.Value = "1D"
	case ST_ConnectorDimension2D:
		attr.Value = "2D"
	case ST_ConnectorDimensionCust:
		attr.Value = "cust"
	}
	return attr, nil
}

func (e *ST_ConnectorDimension) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "1D":
		*e = 1
	case "2D":
		*e = 2
	case "cust":
		*e = 3
	}
	return nil
}

func (m ST_ConnectorDimension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ConnectorDimension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "1D":
			*m = 1
		case "2D":
			*m = 2
		case "cust":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ConnectorDimension) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "1D"
	case 2:
		return "2D"
	case 3:
		return "cust"
	}
	return ""
}

func (m ST_ConnectorDimension) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ConnectorDimension) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ConnectorPoint byte

const (
	ST_ConnectorPointUnset  ST_ConnectorPoint = 0
	ST_ConnectorPointAuto   ST_ConnectorPoint = 1
	ST_ConnectorPointBCtr   ST_ConnectorPoint = 2
	ST_ConnectorPointCtr    ST_ConnectorPoint = 3
	ST_ConnectorPointMidL   ST_ConnectorPoint = 4
	ST_ConnectorPointMidR   ST_ConnectorPoint = 5
	ST_ConnectorPointTCtr   ST_ConnectorPoint = 6
	ST_ConnectorPointBL     ST_ConnectorPoint = 7
	ST_ConnectorPointBR     ST_ConnectorPoint = 8
	ST_ConnectorPointTL     ST_ConnectorPoint = 9
	ST_ConnectorPointTR     ST_ConnectorPoint = 10
	ST_ConnectorPointRadial ST_ConnectorPoint = 11
)

func (e ST_ConnectorPoint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ConnectorPointUnset:
		attr.Value = ""
	case ST_ConnectorPointAuto:
		attr.Value = "auto"
	case ST_ConnectorPointBCtr:
		attr.Value = "bCtr"
	case ST_ConnectorPointCtr:
		attr.Value = "ctr"
	case ST_ConnectorPointMidL:
		attr.Value = "midL"
	case ST_ConnectorPointMidR:
		attr.Value = "midR"
	case ST_ConnectorPointTCtr:
		attr.Value = "tCtr"
	case ST_ConnectorPointBL:
		attr.Value = "bL"
	case ST_ConnectorPointBR:
		attr.Value = "bR"
	case ST_ConnectorPointTL:
		attr.Value = "tL"
	case ST_ConnectorPointTR:
		attr.Value = "tR"
	case ST_ConnectorPointRadial:
		attr.Value = "radial"
	}
	return attr, nil
}

func (e *ST_ConnectorPoint) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "bCtr":
		*e = 2
	case "ctr":
		*e = 3
	case "midL":
		*e = 4
	case "midR":
		*e = 5
	case "tCtr":
		*e = 6
	case "bL":
		*e = 7
	case "bR":
		*e = 8
	case "tL":
		*e = 9
	case "tR":
		*e = 10
	case "radial":
		*e = 11
	}
	return nil
}

func (m ST_ConnectorPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ConnectorPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "bCtr":
			*m = 2
		case "ctr":
			*m = 3
		case "midL":
			*m = 4
		case "midR":
			*m = 5
		case "tCtr":
			*m = 6
		case "bL":
			*m = 7
		case "bR":
			*m = 8
		case "tL":
			*m = 9
		case "tR":
			*m = 10
		case "radial":
			*m = 11
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ConnectorPoint) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "bCtr"
	case 3:
		return "ctr"
	case 4:
		return "midL"
	case 5:
		return "midR"
	case 6:
		return "tCtr"
	case 7:
		return "bL"
	case 8:
		return "bR"
	case 9:
		return "tL"
	case 10:
		return "tR"
	case 11:
		return "radial"
	}
	return ""
}

func (m ST_ConnectorPoint) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ConnectorPoint) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_NodeHorizontalAlignment byte

const (
	ST_NodeHorizontalAlignmentUnset ST_NodeHorizontalAlignment = 0
	ST_NodeHorizontalAlignmentL     ST_NodeHorizontalAlignment = 1
	ST_NodeHorizontalAlignmentCtr   ST_NodeHorizontalAlignment = 2
	ST_NodeHorizontalAlignmentR     ST_NodeHorizontalAlignment = 3
)

func (e ST_NodeHorizontalAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_NodeHorizontalAlignmentUnset:
		attr.Value = ""
	case ST_NodeHorizontalAlignmentL:
		attr.Value = "l"
	case ST_NodeHorizontalAlignmentCtr:
		attr.Value = "ctr"
	case ST_NodeHorizontalAlignmentR:
		attr.Value = "r"
	}
	return attr, nil
}

func (e *ST_NodeHorizontalAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "l":
		*e = 1
	case "ctr":
		*e = 2
	case "r":
		*e = 3
	}
	return nil
}

func (m ST_NodeHorizontalAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_NodeHorizontalAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "l":
			*m = 1
		case "ctr":
			*m = 2
		case "r":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_NodeHorizontalAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "l"
	case 2:
		return "ctr"
	case 3:
		return "r"
	}
	return ""
}

func (m ST_NodeHorizontalAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_NodeHorizontalAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_NodeVerticalAlignment byte

const (
	ST_NodeVerticalAlignmentUnset ST_NodeVerticalAlignment = 0
	ST_NodeVerticalAlignmentT     ST_NodeVerticalAlignment = 1
	ST_NodeVerticalAlignmentMid   ST_NodeVerticalAlignment = 2
	ST_NodeVerticalAlignmentB     ST_NodeVerticalAlignment = 3
)

func (e ST_NodeVerticalAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_NodeVerticalAlignmentUnset:
		attr.Value = ""
	case ST_NodeVerticalAlignmentT:
		attr.Value = "t"
	case ST_NodeVerticalAlignmentMid:
		attr.Value = "mid"
	case ST_NodeVerticalAlignmentB:
		attr.Value = "b"
	}
	return attr, nil
}

func (e *ST_NodeVerticalAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "mid":
		*e = 2
	case "b":
		*e = 3
	}
	return nil
}

func (m ST_NodeVerticalAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_NodeVerticalAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "t":
			*m = 1
		case "mid":
			*m = 2
		case "b":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_NodeVerticalAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "mid"
	case 3:
		return "b"
	}
	return ""
}

func (m ST_NodeVerticalAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_NodeVerticalAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FallbackDimension byte

const (
	ST_FallbackDimensionUnset ST_FallbackDimension = 0
	ST_FallbackDimension1D    ST_FallbackDimension = 1
	ST_FallbackDimension2D    ST_FallbackDimension = 2
)

func (e ST_FallbackDimension) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FallbackDimensionUnset:
		attr.Value = ""
	case ST_FallbackDimension1D:
		attr.Value = "1D"
	case ST_FallbackDimension2D:
		attr.Value = "2D"
	}
	return attr, nil
}

func (e *ST_FallbackDimension) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "1D":
		*e = 1
	case "2D":
		*e = 2
	}
	return nil
}

func (m ST_FallbackDimension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FallbackDimension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "1D":
			*m = 1
		case "2D":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FallbackDimension) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "1D"
	case 2:
		return "2D"
	}
	return ""
}

func (m ST_FallbackDimension) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FallbackDimension) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextDirection byte

const (
	ST_TextDirectionUnset ST_TextDirection = 0
	ST_TextDirectionFromT ST_TextDirection = 1
	ST_TextDirectionFromB ST_TextDirection = 2
)

func (e ST_TextDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextDirectionUnset:
		attr.Value = ""
	case ST_TextDirectionFromT:
		attr.Value = "fromT"
	case ST_TextDirectionFromB:
		attr.Value = "fromB"
	}
	return attr, nil
}

func (e *ST_TextDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "fromT":
		*e = 1
	case "fromB":
		*e = 2
	}
	return nil
}

func (m ST_TextDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "fromT":
			*m = 1
		case "fromB":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "fromT"
	case 2:
		return "fromB"
	}
	return ""
}

func (m ST_TextDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PyramidAccentPosition byte

const (
	ST_PyramidAccentPositionUnset ST_PyramidAccentPosition = 0
	ST_PyramidAccentPositionBef   ST_PyramidAccentPosition = 1
	ST_PyramidAccentPositionAft   ST_PyramidAccentPosition = 2
)

func (e ST_PyramidAccentPosition) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PyramidAccentPositionUnset:
		attr.Value = ""
	case ST_PyramidAccentPositionBef:
		attr.Value = "bef"
	case ST_PyramidAccentPositionAft:
		attr.Value = "aft"
	}
	return attr, nil
}

func (e *ST_PyramidAccentPosition) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bef":
		*e = 1
	case "aft":
		*e = 2
	}
	return nil
}

func (m ST_PyramidAccentPosition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PyramidAccentPosition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "bef":
			*m = 1
		case "aft":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PyramidAccentPosition) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bef"
	case 2:
		return "aft"
	}
	return ""
}

func (m ST_PyramidAccentPosition) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PyramidAccentPosition) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PyramidAccentTextMargin byte

const (
	ST_PyramidAccentTextMarginUnset ST_PyramidAccentTextMargin = 0
	ST_PyramidAccentTextMarginStep  ST_PyramidAccentTextMargin = 1
	ST_PyramidAccentTextMarginStack ST_PyramidAccentTextMargin = 2
)

func (e ST_PyramidAccentTextMargin) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PyramidAccentTextMarginUnset:
		attr.Value = ""
	case ST_PyramidAccentTextMarginStep:
		attr.Value = "step"
	case ST_PyramidAccentTextMarginStack:
		attr.Value = "stack"
	}
	return attr, nil
}

func (e *ST_PyramidAccentTextMargin) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "step":
		*e = 1
	case "stack":
		*e = 2
	}
	return nil
}

func (m ST_PyramidAccentTextMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PyramidAccentTextMargin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "step":
			*m = 1
		case "stack":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PyramidAccentTextMargin) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "step"
	case 2:
		return "stack"
	}
	return ""
}

func (m ST_PyramidAccentTextMargin) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PyramidAccentTextMargin) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextBlockDirection byte

const (
	ST_TextBlockDirectionUnset ST_TextBlockDirection = 0
	ST_TextBlockDirectionHorz  ST_TextBlockDirection = 1
	ST_TextBlockDirectionVert  ST_TextBlockDirection = 2
)

func (e ST_TextBlockDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextBlockDirectionUnset:
		attr.Value = ""
	case ST_TextBlockDirectionHorz:
		attr.Value = "horz"
	case ST_TextBlockDirectionVert:
		attr.Value = "vert"
	}
	return attr, nil
}

func (e *ST_TextBlockDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "horz":
		*e = 1
	case "vert":
		*e = 2
	}
	return nil
}

func (m ST_TextBlockDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextBlockDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "horz":
			*m = 1
		case "vert":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextBlockDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "horz"
	case 2:
		return "vert"
	}
	return ""
}

func (m ST_TextBlockDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextBlockDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextAnchorHorizontal byte

const (
	ST_TextAnchorHorizontalUnset ST_TextAnchorHorizontal = 0
	ST_TextAnchorHorizontalNone  ST_TextAnchorHorizontal = 1
	ST_TextAnchorHorizontalCtr   ST_TextAnchorHorizontal = 2
)

func (e ST_TextAnchorHorizontal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextAnchorHorizontalUnset:
		attr.Value = ""
	case ST_TextAnchorHorizontalNone:
		attr.Value = "none"
	case ST_TextAnchorHorizontalCtr:
		attr.Value = "ctr"
	}
	return attr, nil
}

func (e *ST_TextAnchorHorizontal) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "ctr":
		*e = 2
	}
	return nil
}

func (m ST_TextAnchorHorizontal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextAnchorHorizontal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "ctr":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextAnchorHorizontal) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "ctr"
	}
	return ""
}

func (m ST_TextAnchorHorizontal) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextAnchorHorizontal) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextAnchorVertical byte

const (
	ST_TextAnchorVerticalUnset ST_TextAnchorVertical = 0
	ST_TextAnchorVerticalT     ST_TextAnchorVertical = 1
	ST_TextAnchorVerticalMid   ST_TextAnchorVertical = 2
	ST_TextAnchorVerticalB     ST_TextAnchorVertical = 3
)

func (e ST_TextAnchorVertical) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextAnchorVerticalUnset:
		attr.Value = ""
	case ST_TextAnchorVerticalT:
		attr.Value = "t"
	case ST_TextAnchorVerticalMid:
		attr.Value = "mid"
	case ST_TextAnchorVerticalB:
		attr.Value = "b"
	}
	return attr, nil
}

func (e *ST_TextAnchorVertical) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "mid":
		*e = 2
	case "b":
		*e = 3
	}
	return nil
}

func (m ST_TextAnchorVertical) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextAnchorVertical) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "t":
			*m = 1
		case "mid":
			*m = 2
		case "b":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextAnchorVertical) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "mid"
	case 3:
		return "b"
	}
	return ""
}

func (m ST_TextAnchorVertical) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextAnchorVertical) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DiagramTextAlignment byte

const (
	ST_DiagramTextAlignmentUnset ST_DiagramTextAlignment = 0
	ST_DiagramTextAlignmentL     ST_DiagramTextAlignment = 1
	ST_DiagramTextAlignmentCtr   ST_DiagramTextAlignment = 2
	ST_DiagramTextAlignmentR     ST_DiagramTextAlignment = 3
)

func (e ST_DiagramTextAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DiagramTextAlignmentUnset:
		attr.Value = ""
	case ST_DiagramTextAlignmentL:
		attr.Value = "l"
	case ST_DiagramTextAlignmentCtr:
		attr.Value = "ctr"
	case ST_DiagramTextAlignmentR:
		attr.Value = "r"
	}
	return attr, nil
}

func (e *ST_DiagramTextAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "l":
		*e = 1
	case "ctr":
		*e = 2
	case "r":
		*e = 3
	}
	return nil
}

func (m ST_DiagramTextAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DiagramTextAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "l":
			*m = 1
		case "ctr":
			*m = 2
		case "r":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DiagramTextAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "l"
	case 2:
		return "ctr"
	case 3:
		return "r"
	}
	return ""
}

func (m ST_DiagramTextAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DiagramTextAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AutoTextRotation byte

const (
	ST_AutoTextRotationUnset ST_AutoTextRotation = 0
	ST_AutoTextRotationNone  ST_AutoTextRotation = 1
	ST_AutoTextRotationUpr   ST_AutoTextRotation = 2
	ST_AutoTextRotationGrav  ST_AutoTextRotation = 3
)

func (e ST_AutoTextRotation) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AutoTextRotationUnset:
		attr.Value = ""
	case ST_AutoTextRotationNone:
		attr.Value = "none"
	case ST_AutoTextRotationUpr:
		attr.Value = "upr"
	case ST_AutoTextRotationGrav:
		attr.Value = "grav"
	}
	return attr, nil
}

func (e *ST_AutoTextRotation) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "upr":
		*e = 2
	case "grav":
		*e = 3
	}
	return nil
}

func (m ST_AutoTextRotation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AutoTextRotation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "upr":
			*m = 2
		case "grav":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_AutoTextRotation) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "upr"
	case 3:
		return "grav"
	}
	return ""
}

func (m ST_AutoTextRotation) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AutoTextRotation) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_GrowDirection byte

const (
	ST_GrowDirectionUnset ST_GrowDirection = 0
	ST_GrowDirectionTL    ST_GrowDirection = 1
	ST_GrowDirectionTR    ST_GrowDirection = 2
	ST_GrowDirectionBL    ST_GrowDirection = 3
	ST_GrowDirectionBR    ST_GrowDirection = 4
)

func (e ST_GrowDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_GrowDirectionUnset:
		attr.Value = ""
	case ST_GrowDirectionTL:
		attr.Value = "tL"
	case ST_GrowDirectionTR:
		attr.Value = "tR"
	case ST_GrowDirectionBL:
		attr.Value = "bL"
	case ST_GrowDirectionBR:
		attr.Value = "bR"
	}
	return attr, nil
}

func (e *ST_GrowDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "tL":
		*e = 1
	case "tR":
		*e = 2
	case "bL":
		*e = 3
	case "bR":
		*e = 4
	}
	return nil
}

func (m ST_GrowDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_GrowDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "tL":
			*m = 1
		case "tR":
			*m = 2
		case "bL":
			*m = 3
		case "bR":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_GrowDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "tL"
	case 2:
		return "tR"
	case 3:
		return "bL"
	case 4:
		return "bR"
	}
	return ""
}

func (m ST_GrowDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_GrowDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FlowDirection byte

const (
	ST_FlowDirectionUnset ST_FlowDirection = 0
	ST_FlowDirectionRow   ST_FlowDirection = 1
	ST_FlowDirectionCol   ST_FlowDirection = 2
)

func (e ST_FlowDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FlowDirectionUnset:
		attr.Value = ""
	case ST_FlowDirectionRow:
		attr.Value = "row"
	case ST_FlowDirectionCol:
		attr.Value = "col"
	}
	return attr, nil
}

func (e *ST_FlowDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "row":
		*e = 1
	case "col":
		*e = 2
	}
	return nil
}

func (m ST_FlowDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FlowDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "row":
			*m = 1
		case "col":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FlowDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "row"
	case 2:
		return "col"
	}
	return ""
}

func (m ST_FlowDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FlowDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ContinueDirection byte

const (
	ST_ContinueDirectionUnset   ST_ContinueDirection = 0
	ST_ContinueDirectionRevDir  ST_ContinueDirection = 1
	ST_ContinueDirectionSameDir ST_ContinueDirection = 2
)

func (e ST_ContinueDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ContinueDirectionUnset:
		attr.Value = ""
	case ST_ContinueDirectionRevDir:
		attr.Value = "revDir"
	case ST_ContinueDirectionSameDir:
		attr.Value = "sameDir"
	}
	return attr, nil
}

func (e *ST_ContinueDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "revDir":
		*e = 1
	case "sameDir":
		*e = 2
	}
	return nil
}

func (m ST_ContinueDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ContinueDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "revDir":
			*m = 1
		case "sameDir":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ContinueDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "revDir"
	case 2:
		return "sameDir"
	}
	return ""
}

func (m ST_ContinueDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ContinueDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Breakpoint byte

const (
	ST_BreakpointUnset  ST_Breakpoint = 0
	ST_BreakpointEndCnv ST_Breakpoint = 1
	ST_BreakpointBal    ST_Breakpoint = 2
	ST_BreakpointFixed  ST_Breakpoint = 3
)

func (e ST_Breakpoint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BreakpointUnset:
		attr.Value = ""
	case ST_BreakpointEndCnv:
		attr.Value = "endCnv"
	case ST_BreakpointBal:
		attr.Value = "bal"
	case ST_BreakpointFixed:
		attr.Value = "fixed"
	}
	return attr, nil
}

func (e *ST_Breakpoint) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "endCnv":
		*e = 1
	case "bal":
		*e = 2
	case "fixed":
		*e = 3
	}
	return nil
}

func (m ST_Breakpoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Breakpoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "endCnv":
			*m = 1
		case "bal":
			*m = 2
		case "fixed":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Breakpoint) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "endCnv"
	case 2:
		return "bal"
	case 3:
		return "fixed"
	}
	return ""
}

func (m ST_Breakpoint) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Breakpoint) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Offset byte

const (
	ST_OffsetUnset ST_Offset = 0
	ST_OffsetCtr   ST_Offset = 1
	ST_OffsetOff   ST_Offset = 2
)

func (e ST_Offset) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OffsetUnset:
		attr.Value = ""
	case ST_OffsetCtr:
		attr.Value = "ctr"
	case ST_OffsetOff:
		attr.Value = "off"
	}
	return attr, nil
}

func (e *ST_Offset) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "ctr":
		*e = 1
	case "off":
		*e = 2
	}
	return nil
}

func (m ST_Offset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Offset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "ctr":
			*m = 1
		case "off":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Offset) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "ctr"
	case 2:
		return "off"
	}
	return ""
}

func (m ST_Offset) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Offset) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HierarchyAlignment byte

const (
	ST_HierarchyAlignmentUnset   ST_HierarchyAlignment = 0
	ST_HierarchyAlignmentTL      ST_HierarchyAlignment = 1
	ST_HierarchyAlignmentTR      ST_HierarchyAlignment = 2
	ST_HierarchyAlignmentTCtrCh  ST_HierarchyAlignment = 3
	ST_HierarchyAlignmentTCtrDes ST_HierarchyAlignment = 4
	ST_HierarchyAlignmentBL      ST_HierarchyAlignment = 5
	ST_HierarchyAlignmentBR      ST_HierarchyAlignment = 6
	ST_HierarchyAlignmentBCtrCh  ST_HierarchyAlignment = 7
	ST_HierarchyAlignmentBCtrDes ST_HierarchyAlignment = 8
	ST_HierarchyAlignmentLT      ST_HierarchyAlignment = 9
	ST_HierarchyAlignmentLB      ST_HierarchyAlignment = 10
	ST_HierarchyAlignmentLCtrCh  ST_HierarchyAlignment = 11
	ST_HierarchyAlignmentLCtrDes ST_HierarchyAlignment = 12
	ST_HierarchyAlignmentRT      ST_HierarchyAlignment = 13
	ST_HierarchyAlignmentRB      ST_HierarchyAlignment = 14
	ST_HierarchyAlignmentRCtrCh  ST_HierarchyAlignment = 15
	ST_HierarchyAlignmentRCtrDes ST_HierarchyAlignment = 16
)

func (e ST_HierarchyAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HierarchyAlignmentUnset:
		attr.Value = ""
	case ST_HierarchyAlignmentTL:
		attr.Value = "tL"
	case ST_HierarchyAlignmentTR:
		attr.Value = "tR"
	case ST_HierarchyAlignmentTCtrCh:
		attr.Value = "tCtrCh"
	case ST_HierarchyAlignmentTCtrDes:
		attr.Value = "tCtrDes"
	case ST_HierarchyAlignmentBL:
		attr.Value = "bL"
	case ST_HierarchyAlignmentBR:
		attr.Value = "bR"
	case ST_HierarchyAlignmentBCtrCh:
		attr.Value = "bCtrCh"
	case ST_HierarchyAlignmentBCtrDes:
		attr.Value = "bCtrDes"
	case ST_HierarchyAlignmentLT:
		attr.Value = "lT"
	case ST_HierarchyAlignmentLB:
		attr.Value = "lB"
	case ST_HierarchyAlignmentLCtrCh:
		attr.Value = "lCtrCh"
	case ST_HierarchyAlignmentLCtrDes:
		attr.Value = "lCtrDes"
	case ST_HierarchyAlignmentRT:
		attr.Value = "rT"
	case ST_HierarchyAlignmentRB:
		attr.Value = "rB"
	case ST_HierarchyAlignmentRCtrCh:
		attr.Value = "rCtrCh"
	case ST_HierarchyAlignmentRCtrDes:
		attr.Value = "rCtrDes"
	}
	return attr, nil
}

func (e *ST_HierarchyAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "tL":
		*e = 1
	case "tR":
		*e = 2
	case "tCtrCh":
		*e = 3
	case "tCtrDes":
		*e = 4
	case "bL":
		*e = 5
	case "bR":
		*e = 6
	case "bCtrCh":
		*e = 7
	case "bCtrDes":
		*e = 8
	case "lT":
		*e = 9
	case "lB":
		*e = 10
	case "lCtrCh":
		*e = 11
	case "lCtrDes":
		*e = 12
	case "rT":
		*e = 13
	case "rB":
		*e = 14
	case "rCtrCh":
		*e = 15
	case "rCtrDes":
		*e = 16
	}
	return nil
}

func (m ST_HierarchyAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HierarchyAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "tL":
			*m = 1
		case "tR":
			*m = 2
		case "tCtrCh":
			*m = 3
		case "tCtrDes":
			*m = 4
		case "bL":
			*m = 5
		case "bR":
			*m = 6
		case "bCtrCh":
			*m = 7
		case "bCtrDes":
			*m = 8
		case "lT":
			*m = 9
		case "lB":
			*m = 10
		case "lCtrCh":
			*m = 11
		case "lCtrDes":
			*m = 12
		case "rT":
			*m = 13
		case "rB":
			*m = 14
		case "rCtrCh":
			*m = 15
		case "rCtrDes":
			*m = 16
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HierarchyAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "tL"
	case 2:
		return "tR"
	case 3:
		return "tCtrCh"
	case 4:
		return "tCtrDes"
	case 5:
		return "bL"
	case 6:
		return "bR"
	case 7:
		return "bCtrCh"
	case 8:
		return "bCtrDes"
	case 9:
		return "lT"
	case 10:
		return "lB"
	case 11:
		return "lCtrCh"
	case 12:
		return "lCtrDes"
	case 13:
		return "rT"
	case 14:
		return "rB"
	case 15:
		return "rCtrCh"
	case 16:
		return "rCtrDes"
	}
	return ""
}

func (m ST_HierarchyAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HierarchyAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VariableType byte

const (
	ST_VariableTypeUnset         ST_VariableType = 0
	ST_VariableTypeNone          ST_VariableType = 1
	ST_VariableTypeOrgChart      ST_VariableType = 2
	ST_VariableTypeChMax         ST_VariableType = 3
	ST_VariableTypeChPref        ST_VariableType = 4
	ST_VariableTypeBulEnabled    ST_VariableType = 5
	ST_VariableTypeDir           ST_VariableType = 6
	ST_VariableTypeHierBranch    ST_VariableType = 7
	ST_VariableTypeAnimOne       ST_VariableType = 8
	ST_VariableTypeAnimLvl       ST_VariableType = 9
	ST_VariableTypeResizeHandles ST_VariableType = 10
)

func (e ST_VariableType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VariableTypeUnset:
		attr.Value = ""
	case ST_VariableTypeNone:
		attr.Value = "none"
	case ST_VariableTypeOrgChart:
		attr.Value = "orgChart"
	case ST_VariableTypeChMax:
		attr.Value = "chMax"
	case ST_VariableTypeChPref:
		attr.Value = "chPref"
	case ST_VariableTypeBulEnabled:
		attr.Value = "bulEnabled"
	case ST_VariableTypeDir:
		attr.Value = "dir"
	case ST_VariableTypeHierBranch:
		attr.Value = "hierBranch"
	case ST_VariableTypeAnimOne:
		attr.Value = "animOne"
	case ST_VariableTypeAnimLvl:
		attr.Value = "animLvl"
	case ST_VariableTypeResizeHandles:
		attr.Value = "resizeHandles"
	}
	return attr, nil
}

func (e *ST_VariableType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "orgChart":
		*e = 2
	case "chMax":
		*e = 3
	case "chPref":
		*e = 4
	case "bulEnabled":
		*e = 5
	case "dir":
		*e = 6
	case "hierBranch":
		*e = 7
	case "animOne":
		*e = 8
	case "animLvl":
		*e = 9
	case "resizeHandles":
		*e = 10
	}
	return nil
}

func (m ST_VariableType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VariableType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "orgChart":
			*m = 2
		case "chMax":
			*m = 3
		case "chPref":
			*m = 4
		case "bulEnabled":
			*m = 5
		case "dir":
			*m = 6
		case "hierBranch":
			*m = 7
		case "animOne":
			*m = 8
		case "animLvl":
			*m = 9
		case "resizeHandles":
			*m = 10
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_VariableType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "orgChart"
	case 3:
		return "chMax"
	case 4:
		return "chPref"
	case 5:
		return "bulEnabled"
	case 6:
		return "dir"
	case 7:
		return "hierBranch"
	case 8:
		return "animOne"
	case 9:
		return "animLvl"
	case 10:
		return "resizeHandles"
	}
	return ""
}

func (m ST_VariableType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VariableType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_OutputShapeType byte

const (
	ST_OutputShapeTypeUnset ST_OutputShapeType = 0
	ST_OutputShapeTypeNone  ST_OutputShapeType = 1
	ST_OutputShapeTypeConn  ST_OutputShapeType = 2
)

func (e ST_OutputShapeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OutputShapeTypeUnset:
		attr.Value = ""
	case ST_OutputShapeTypeNone:
		attr.Value = "none"
	case ST_OutputShapeTypeConn:
		attr.Value = "conn"
	}
	return attr, nil
}

func (e *ST_OutputShapeType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "conn":
		*e = 2
	}
	return nil
}

func (m ST_OutputShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_OutputShapeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "conn":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_OutputShapeType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "conn"
	}
	return ""
}

func (m ST_OutputShapeType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_OutputShapeType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_CTName", NewCT_CTName)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_CTDescription", NewCT_CTDescription)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_CTCategory", NewCT_CTCategory)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_CTCategories", NewCT_CTCategories)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Colors", NewCT_Colors)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_CTStyleLabel", NewCT_CTStyleLabel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ColorTransform", NewCT_ColorTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ColorTransformHeader", NewCT_ColorTransformHeader)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ColorTransformHeaderLst", NewCT_ColorTransformHeaderLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Pt", NewCT_Pt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_PtList", NewCT_PtList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Cxn", NewCT_Cxn)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_CxnList", NewCT_CxnList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_DataModel", NewCT_DataModel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Constraint", NewCT_Constraint)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Constraints", NewCT_Constraints)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_NumericRule", NewCT_NumericRule)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Rules", NewCT_Rules)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_PresentationOf", NewCT_PresentationOf)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Adj", NewCT_Adj)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_AdjLst", NewCT_AdjLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Shape", NewCT_Shape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Parameter", NewCT_Parameter)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Algorithm", NewCT_Algorithm)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_LayoutNode", NewCT_LayoutNode)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ForEach", NewCT_ForEach)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_When", NewCT_When)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Otherwise", NewCT_Otherwise)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Choose", NewCT_Choose)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_SampleData", NewCT_SampleData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Category", NewCT_Category)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Categories", NewCT_Categories)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Name", NewCT_Name)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Description", NewCT_Description)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_DiagramDefinition", NewCT_DiagramDefinition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_DiagramDefinitionHeader", NewCT_DiagramDefinitionHeader)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_DiagramDefinitionHeaderLst", NewCT_DiagramDefinitionHeaderLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_RelIds", NewCT_RelIds)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ElemPropSet", NewCT_ElemPropSet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_OrgChart", NewCT_OrgChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ChildMax", NewCT_ChildMax)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ChildPref", NewCT_ChildPref)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_BulletEnabled", NewCT_BulletEnabled)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_Direction", NewCT_Direction)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_HierBranchStyle", NewCT_HierBranchStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_AnimOne", NewCT_AnimOne)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_AnimLvl", NewCT_AnimLvl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_ResizeHandles", NewCT_ResizeHandles)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_LayoutVariablePropertySet", NewCT_LayoutVariablePropertySet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_SDName", NewCT_SDName)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_SDDescription", NewCT_SDDescription)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_SDCategory", NewCT_SDCategory)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_SDCategories", NewCT_SDCategories)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_TextProps", NewCT_TextProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_StyleLabel", NewCT_StyleLabel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_StyleDefinition", NewCT_StyleDefinition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_StyleDefinitionHeader", NewCT_StyleDefinitionHeader)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "CT_StyleDefinitionHeaderLst", NewCT_StyleDefinitionHeaderLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "colorsDef", NewColorsDef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "colorsDefHdr", NewColorsDefHdr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "colorsDefHdrLst", NewColorsDefHdrLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "dataModel", NewDataModel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "layoutDef", NewLayoutDef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "layoutDefHdr", NewLayoutDefHdr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "layoutDefHdrLst", NewLayoutDefHdrLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "relIds", NewRelIds)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "styleDef", NewStyleDef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "styleDefHdr", NewStyleDefHdr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "styleDefHdrLst", NewStyleDefHdrLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "AG_IteratorAttributes", NewAG_IteratorAttributes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "AG_ConstraintAttributes", NewAG_ConstraintAttributes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/diagram", "AG_ConstraintRefAttributes", NewAG_ConstraintRefAttributes)
}

type ST_AxisTypes []ST_AxisType
type ST_ElementTypes []ST_ElementType
type ST_Ints []int32
type ST_UnsignedInts []uint32
type ST_Booleans []bool
