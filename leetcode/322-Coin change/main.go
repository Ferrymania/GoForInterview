/*
You are given coins of different denominations and a total amount of money
amount. Write a function to compute the fewest number of coins that you need to
make up that amount. If that amount of money cannot be made up by any
combination of the coins, return -1.
Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Input: coins = [2], amount = 3
Output: -1

 */
package main

import "golang.org/x/text/currency"

/*
当总价等于11的时候，设f(n)为凑成n元最少的币数
取一张1元的时候：cost = f(10) + 1
取一张2的时候 cost = f(9) + 1
取一张5的时候 cost = f(6) + 1
所以 f(n)只与f(n-1) f(n-2) f(n-5) 有关系
f(n) = min{f(n-1),f(n-2),f(n-5)}+1
 */

func coinChange(coins []int, amount int) int {
	if amount < 0{
		return 0
	}
	return coinChangeDP(coins,amount)
}

func coinChangeDP(coins []int,amount int) int {
	if amount <0 {
		return -1
	}
	if amount == 0 {
		return  0
	}
	count := make([]int,amount)
}

func main(){

}
