package string

import (
	"strings"
	"fmt"
)

//FindKeyString finds string value when given a specified parameter-list(json,key string)
func FindKeyString(s string, key string) (string, error) {
	if !strings.Contains(s, ",") {
		if strings.Contains(s, "\""+key+"\"") {
			t := strings.Split(s, "\"")
			for i, t2 := range t {
				if strings.Contains(t2, ":") {
					if t[i-1] == key {
						return t[i+1], nil
					}
				}
			}
		}
	}
	line := strings.Split(s, ",")
	for _, v := range line {
		if strings.Contains(v, "\""+key+"\":") {
			t := strings.Split(v, "\"")
			for i, t2 := range t {
				if strings.Contains(t2, ":") {
					if t[i-1] == key {
						return t[i+1], nil
					}
				}
			}
		}
	}
	return "", fmt.Errorf("FindKeyString: can't find %s from %s", key, s)
}
