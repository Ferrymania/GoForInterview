/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、
"5e2" "-123" "3.1416" "-1E-16"都是表示数值。但"12e"、"1a3.14"、"1.2.3"、
"+-5" 及"12e+5.4"都不是
 */
package main


func IsNumeric(str string)bool {
	if str == ""{
		return false
	}
	numeric := scanInteger([]byte(str))

}

func scanInteger(str []byte)bool {
	if str[0] == '+' || str[0]=='-'{
		return scanUnsignedInteger(str[1:])
	}
	return scanUnsignedInteger(str)

}

func scanUnsignedInteger(str []byte)bool{

}