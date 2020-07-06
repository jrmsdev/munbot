// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"fmt"

	"github.com/jrmsdev/munbot/log"
)

func Parse(c *Munbot) (map[string]string, error) {
	log.Debug("parse...")
	m := make(map[string]interface{})
	if blob, err := Bytes(c); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(blob, &m); err != nil {
			return nil, err
		}
	}
	dst := make(map[string]string)
	return walk(dst, "", m), nil
}

func walk(dst map[string]string, prefix string, m map[string]interface{}) map[string]string {
	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			p := fmt.Sprintf("%s%s.", prefix, k)
			walk(dst, p, v.(map[string]interface{}))
		default:
			dst[prefix + k] = fmt.Sprintf("%v", v)
		}
	}
	return dst
}
