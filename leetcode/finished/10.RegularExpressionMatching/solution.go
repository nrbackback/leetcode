package solution

func isMatch(s string, p string) bool {
	if len(s) < 1 || len(s) > 20 {
		return false
	}
	if len(p) < 1 || len(p) > 30 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	for i := 0; i < len(p); i++ {
		v := p[i]
		if v != '.' && v != '*' {
			if v < 'a' || v > 'z' {
				return false
			}
		}
	}
	beginS, endS, beginP, endP := 0, 0, 0, 0
	for {
		if beginS >= len(s) || beginP >= len(p) {
			break
		}
		pv := p[beginP]
		sv := s[beginS]
		if pv == '*' {
			// 不让这种情况出现
			return false
		}
		if pv != '.' {
			// 字母

			// 判断下一位是否是*
			if beginP == len(p)-1 {
				// 最后一位了
				if pv != sv {
					return false
				} else {
					break
				}
			} else {
				// 不是最后一位
				nextPv := p[beginP+1]
				
			}

		}
		if pv == '.' {

		}
	}
	// TODO
	return false
}
