package others

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

var NormalErr = fmt.Errorf("EOF")

/*FindKey finds the specific in string s and return it's value and error.
warning:!!it can only find struct value {...} and array value [...],if you
figure out some improvements and bugs,fork it and commit ,i will appreciate it :),
or you contact me by 1027309561@qq.com */
func FindKey(s, key string) (string, error) {
	rs := []rune(s)
	plusRule := PlusRule()
	var fs []rune
	t := ""
	index := 0
	for _, v2 := range rs {
		if string(v2) == " " || string(v2) == "\n" || string(v2) == "\t" {
			continue
		}
		fs = append(fs, v2)
	}
	for i, v := range fs {
		if v == int32(key[index]) {
			if index == len(key)-1 {
				if fs[i-len(key)] != []rune("\"")[0] || fs[i+2] != []rune(":")[0] {
					index = 0
					continue
				}
				key := i + 3
				t += string(fs[key])
				for !plusRule(fs[key], fs[i+3]) {
					t += string(fs[i+4])
					i++
				}
				return t, nil
			}
			index++
			continue
		}
		index = 0
	}
	return "", fmt.Errorf(fmt.Sprintf("no %s in %s", key, s))
}

func PlusRule() func(rune, rune) bool {
	count := 0
	return func(s, n rune) bool {
		var head rune
		var tail rune
		if s == []rune("[")[0] {
			head, tail = []rune("[")[0], []rune("]")[0]
		} else if s == []rune("{" )[0] {
			head, tail = []rune("{" )[0], []rune("}" )[0]
		} else {
			log.Error(fmt.Errorf("findKey: set key: only { or [ is permmited,what i get is %s", string(s)))
			count = 0
		}
		if n == head {
			count++
		}
		if n == tail {
			count--
		}
		if count != 0 {
			return false
		}
		return true
	}
}
