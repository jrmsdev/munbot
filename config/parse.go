// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/munbot/master/log"
)

func Parse(c *Munbot, filter string) (map[string]string, error) {
	log.Debugf("parse filter='%s'", filter)
	m := make(map[string]interface{})
	if blob, err := Bytes(c); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(blob, &m); err != nil {
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
				dst[opt] = fmt.Sprintf("%v", v)
			}
		}
	}
	return dst
}

func ParseJSON(c *Munbot, filter string) (map[string]interface{}, error) {
	log.Debugf("parse json filter='%s'", filter)
	m := make(map[string]interface{})
	if blob, err := Bytes(c); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(blob, &m); err != nil {
			return nil, err
		}
	}
	if filter == "" {
		return m, nil
	}
	dst := make(map[string]interface{})
	return walkJSON(dst, "", m, filter), nil
}

func walkJSON(dst map[string]interface{}, prefix string, m map[string]interface{}, filter string) map[string]interface{} {
	for k, v := range m {
		opt := prefix + k
		if filter == "" || strings.HasPrefix(opt, filter) || strings.HasPrefix(filter, opt) {
			switch v.(type) {
				case map[string]interface{}:
					d := make(map[string]interface{})
					dst[k] = d
					p := fmt.Sprintf("%s%s.", prefix, k)
					walkJSON(d, p, v.(map[string]interface{}), filter)
				default:
					dst[k] = v
			}
		}
	}
	return dst
}
