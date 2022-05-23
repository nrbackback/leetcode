package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	r := strings.TrimSpace(s)
	if len(r) == 0 {
		return 0
	}
	var passitive bool
	var beginIndex int
	if r[0] == '-' {
		passitive = true
		beginIndex = 1
	}
	if r[0] == '+' {
		beginIndex = 1
	}
	var endIndex int
	endIndex = len(r)
	for i := beginIndex; i < len(r); i++ {
		if r[i] < '0' || r[i] > '9' {
			endIndex = i
			break
		}
	}

	v := r[beginIndex:endIndex]
	var parsed int64
	if v != "" {
		var err error
		parsed, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			// 这个转换失败是给越界处理用的
			fmt.Printf("parse %s error %v\n", v, err)

			parsed = 2 << 30
		}
	} else {
		// v为""的时候特殊处理
		parsed = 0
	}

	if passitive {
		parsed = -parsed
	}
	if parsed < -2<<30 {
		parsed = -2 << 30
	} else if parsed > 2<<30-1 {
		parsed = 2<<30 - 1
	}

	return int(parsed)
}
