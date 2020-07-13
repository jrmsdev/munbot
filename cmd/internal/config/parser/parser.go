// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/munbot/master/config"
)

type marshalFunc func(interface{}) ([]byte, error)
type unmarshalFunc func([]byte, interface{}) error

var jsonMarshal marshalFunc = json.Marshal
var jsonUnmarshal unmarshalFunc = json.Unmarshal

func Parse(c *config.Munbot, filter string) (map[string]string, error) {
	m := make(map[string]interface{})
	if blob, err := jsonMarshal(c); err != nil {
		return nil, err
	} else {
		if err := jsonUnmarshal(blob, &m); err != nil {
			return nil, err
		}
	}
	dst := make(map[string]string)
	return walk(dst, "", m, filter), nil
}

func walk(dst map[string]string, prefix string, m map[string]interface{}, filter string) map[string]string {
	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			p := fmt.Sprintf("%s%s.", prefix, k)
			walk(dst, p, v.(map[string]interface{}), filter)
		default:
			opt := prefix + k
			if filter == "" || strings.HasPrefix(opt, filter) {
				dst[opt] = fmt.Sprint(v)
			}
		}
	}
	return dst
}
