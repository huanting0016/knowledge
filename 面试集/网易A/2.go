package main

func findKthBit(n, k int) byte {
	pre := "a"
	for i := 2; i <= n; i++ {
		pre += string(rune('a'+i-1)) + rein(pre)
	}
	return pre[k-1]
}

func rein(s string) string {
	bs := make([]byte, len(s))
	for i, b := range s {
		if b >= 'n' {
			b = 'm' - (b - 'n')
		} else {
			b = 'n' + ('m' - b)
		}
		bs[len(s)-i-1] = byte(b)
	}
	return string(bs)
}
