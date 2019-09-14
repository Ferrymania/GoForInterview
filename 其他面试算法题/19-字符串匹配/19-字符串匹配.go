/*
给定主字符串s与模式字符串p，判断p是否是s的子串，如果是，则找出p在s中第一次出现的下标
 */
package _9_字符串匹配

import "fmt"

/*
直接计算法 O(n(m-n))
 */

func directMatch(str,pattern string)int{
	sLen := len(str)
	pLen := len(pattern)
	if sLen <pLen{
		return -1
	}
	i,j:=0,0
	for i<sLen && j<pLen {
		if str[i] == pattern [j]{  //如果字符相同则继续比较下一字符
			i++
			j++
		}else {        //如果不相同，则退回到最初比较首元素下个元素
			i = i-j+1
			j =0
			if i>sLen-pLen{  //剩余字符串长度小于模板则直接退出
				return -1
			}

		}
	}

	if j >= pLen {
		return i-pLen
	}
	return -1
}

/*
KMP算法
 */
// str -> text pattern -> template  q prime number
func kmp(str string,pattern string,q int)int{
	sLen := len(str)
	pLen := len(pattern)
	if sLen < pLen {
		return -1
	}
	d :=256 //the number of characters in the input alphabet
	sHash :=0  //hash value for text
	pHash :=0  //hash value for txt

	h :=1
	// The value of h would be "pow(d, M-1)%q"
	for i:=0;i<pLen-1;i++{
		h = (h*d)%q
	}
	// Calculate the hash value of pattern and first window of text
	for i:=0;i<pLen;i++{
		pHash = (d*pHash+int(pattern[i]))%q
		sHash = (d*sHash+int(str[i]))%q
	}
	// Slide the pattern over text one by one
	var i,j int
	for i= 0;i<=sLen-pLen;i++{
		// Check the hash values of current window of text
		// and pattern. If the hash values match then only
		// check for characters on by one
		if pHash == sHash {
			//Check for characters one by one
			for j=0;j<pLen;j++{
				if str[i+j] != pattern[j]{
					break
				}
			}
			if j == pLen {
				fmt.Println("kmp found match at",i)
				return i
			}
		}
		if i< sLen-pLen {
			sHash = (d*(sHash-int(str[i]))*h) + int(str[i+pLen])%q
			if sHash < 0{
				sHash += q
			}
		}
	}
	return -1
}

func main(){
	str := "sadfhlkjasflkjad"
	pattern := "lkja"
	fmt.Println(directMatch(str,pattern))

	fmt.Println(kmp(str,pattern,101))
}