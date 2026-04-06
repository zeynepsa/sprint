package main

func main() {
	print(ToUpperCase("Hello! How's your day going?"))
}

func ToUpperCase(s string) string {
	str:=""

	for _,v :=range s{
		if v>='a' && v<='z'{
			str +=string(v-32)
		}else{
			str +=string(v)		
		}
	}
	return str
}
