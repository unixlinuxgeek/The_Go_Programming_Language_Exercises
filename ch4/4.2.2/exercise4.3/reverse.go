// Перепишите функцию reverse так, чтобы вместо среза
// она использовала указатель на массив.

package reverse

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
