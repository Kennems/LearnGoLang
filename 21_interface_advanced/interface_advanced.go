package main

type stringer interface {
	String() string
}

type userName string

func (u userName) String() string {
	return string(u)
}

func assertString(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok
}

func convertInt64(v int) int64 {
	return int64(v)
}

func compareInterfaces(a, b any) bool {
	return a == b
}
