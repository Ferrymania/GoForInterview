package main

import "fmt"

func isLow(ch byte)bool{
	return 'a'<=ch && ch <='z'
}

func min(a,b int) int{
	if a<b{
		return a
	}else{
		return b
	}
}
func getButtonFrequency(s string)int{
	if s == ""{
		return 0
	}
	sByte := []byte(s)
	sLength := len(s)
	lows :=make([]int,sLength)
	ups :=make([]int,sLength)
	last := sByte[sLength-1]
	if isLow(last){
		lows[sLength-1] = 1
		ups[sLength-1] = 2
	}else {
		lows[sLength-1] = 2
		ups[sLength-1] = 1
	}

	for i:=sLength-2;i>-1;i--{
		ch :=sByte[i]
		if isLow(ch){
			lows[i] = 1 + lows[i+1]
			ups[i] = 2 + min(lows[i+1],ups[i+1])
		}else {
			ups[i] = 1 + ups[i+1]
			lows[i] = 2+min(lows[i+1],ups[i+1])
		}
	}
	return lows[0]
}
func main(){
	fmt.Println(getButtonFrequency("AaAAAA"))
}
