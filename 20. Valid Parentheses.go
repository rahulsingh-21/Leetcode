package main

func isValid(s string) bool {
	st := []byte{}
	//st = append(st, s[0])
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st = append(st, s[i])
			continue
		}
		if len(st) == 0 {
			return false
		}
		temp := s[i]
		top := st[len(st)-1]

		if top == '(' && temp != ')' {
			return false
		} else if top == '[' && temp != ']' {
			return false
		} else if top == '{' && temp != '}' {
			return false
		}
		st = st[:len(st)-1]
	}
	return len(st) == 0
}
