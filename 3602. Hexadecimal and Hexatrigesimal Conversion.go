package main

import "fmt"

func concatHex36(n int) string {
	return hexAll(n, "hexa") + hexAll(n, "hexd")
}

func hexAll(n int, typ string) string {
	res := ""

	switch typ {
	case "hexd":
		t := n * n * n
		s := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		fmt.Println(t)
		for t > 0 {
			res = string(s[t%36]) + res
			t /= 36
		}
		fmt.Println(res)
	case "hexa":
		t := n * n
		s := "0123456789ABCDEF"
		for t > 0 {
			res = string(s[t%16]) + res
			t /= 16
		}
	}

	return res
}
