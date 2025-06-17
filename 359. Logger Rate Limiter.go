package main

type Logger struct {
	m map[string]int
}

func ConstructorIX() Logger {
	m := make(map[string]int)
	return Logger{m}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	val, ok := this.m[message]
	if !ok {
		this.m[message] = timestamp + 10
		return true
	}
	if val > timestamp {
		return false
	}
	this.m[message] = timestamp + 10
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
