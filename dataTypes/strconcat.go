package main

func main() {
	print(StrConcat("r", "ProgrammerHumour", "/"))
}

func StrConcat(s1, s2, delim string) string {
	str := s1 + delim + s2
	return str
}
