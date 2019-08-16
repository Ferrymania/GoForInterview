/*
cal sum of prime numbers between 1 - 100
*/
package main

import "fmt"

func isPrime(num int) bool {
	for i:=2;i*i <=num;i++{
		if num%i == 0{
			return false
		}
	}
	return true
}

func main(){
	sum :=0
	for i:=2;i<=100;i++ {
		if isPrime(i){
			sum +=i
			fmt.Println("prime",i)
		}
	}

	fmt.Println("sum is ",sum)
}
