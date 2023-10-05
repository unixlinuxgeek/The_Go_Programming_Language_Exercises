package main

// nonempty возвращает срез, содержащий только непустые строки.
// Содержимое базового массива при работе функции изменится.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] //
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
