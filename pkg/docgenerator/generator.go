// Copyright (c) TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package docgenerator

import (
	"reflect"

	"github.com/topfreegames/pitaya/v3/pkg/component"
)

type docs struct {
	Handlers docMap `json:"handlers"`
	Remotes  docMap `json:"remotes"`
}

type docMap map[string]*doc

type doc struct {
	Input  interface{}   `json:"input"`
	Output []interface{} `json:"output"`
}

// HandlersDocs returns a map from route to input and output
func HandlersDocs(serverType string, services map[string]*component.Service, getPtrNames bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemotesDocs returns a map from route to input and output
func RemotesDocs(serverType string, services map[string]*component.Service, getPtrNames bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d docMap) toMap() (map[string]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func docForMethod(method reflect.Method, getPtrNames bool) *doc {
	_ = "STUB: not implemented"
	return nil
}

func parseStruct(typ reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func docForType(typ reflect.Type, isOutput bool, getPtrNames bool) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func validName(field reflect.StructField) bool { _ = "STUB: not implemented"; return false }

func firstLetterToLower(name string, isOutput bool) string { _ = "STUB: not implemented"; return "" }

func getName(field reflect.StructField, isOutput bool) (name string, valid bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseType(typ reflect.Type, isOutput bool, getPtrNames bool) interface{} {
	_ = "STUB: not implemented"
	return nil
}
