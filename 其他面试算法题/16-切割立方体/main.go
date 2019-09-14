/*
有一个x*y*z的立方体，要在这个立方体上砍k刀，每一刀可以看作是用一个平行于立方体某一面的平面切割立方体，且必须在坐标为整数的位置切割，如在x=0.5处用平面切割是非法的。
问在切割k刀之后，最多可以把立方体切割成多少块。
 */
package main

import "fmt"

func maxParts(x,y,z,k int)int{
	pointX := x-1
	pointY := y-1
	pointZ := z-1
	if k == 0 {
		return 1
	}
	if pointX <1 && pointY<1&& pointZ<1&& k>0 {
		return 1
	}
	parts :=2
	for i:=1;pointX+pointX+pointY >=3 && k>1; {
		parts = parts+(i*i+i)/2+1
		fmt.Println("parts ",parts)
		i++
		k--
	}
	for i:=2;pointX+pointY+pointZ <=2&& pointX+pointY+pointZ>0 && k>1;{
		parts += i
		k--
		i++
	}
	return parts
}


func main() {
	//var x,y,z,k int
	//fmt.Scanln(&x,&y,&z,&k)
	fmt.Println(maxParts(2,2,2,3))
}