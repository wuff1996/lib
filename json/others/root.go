package others

import (
	"fmt"
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
				key := i + 3
				if  err := plusRule(fs[key], fs[i+3]); err != NormalErr{
					log.
				}
				t += string(fs[key])
				for err := plusRule(fs[key], fs[i+3]); err == NormalErr; i++ {
					t += string(fs[i+4])
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

func PlusRule() func(rune, rune) error {
	count := 0
	return func(s, n rune) error {
		var head rune
		var tail rune
		if s == []rune("[")[0] {
			head, tail = []rune("[")[0], []rune("]")[0]
		} else if s == []rune("{" )[0] {
			head, tail = []rune("{" )[0], []rune("}" )[0]
		} else {
			return fmt.Errorf("set key: only { or [ is permmited,what you put is %s", string(s))
		}
		if n == head {
			count++
		}
		if n == tail {
			count--
		}
		if count != 0 {
			return fmt.Errorf("EOF")
		}
		return nil
	}
}
