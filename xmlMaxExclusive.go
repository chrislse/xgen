// Copyright 2020 The xgen Authors. All rights reserved. Use of this source
// code is governed by a BSD-style license that can be found in the LICENSE
// file.
//
// Package xgen written in pure Go providing a set of functions that allow you
// to parse XSD (XML schema files). This library needs Go version 1.10 or
// later.

package xgen

import "encoding/xml"

// EndMaxExclusive handles parsing event on the maxExclusive end elements.
// MaxExclusive specifies the upper bounds for numeric values (the value must
// be less than this value).
func (opt *Options) EndMaxExclusive(ele xml.EndElement, protoTree []interface{}) (err error) {
	if opt.SimpleType.Len() > 0 && opt.Element.Len() > 0 {
		if opt.Element.Peek().(*Element).Type, err = opt.GetValueType(opt.SimpleType.Pop().(*SimpleType).Base, opt.ProtoTree); err != nil {
			return
		}
		opt.CurrentEle = ""
	}
	return
}