/*
HackerRank: https://www.hackerrank.com/challenges/electronics-shop/problem

A person wants to determine the most expensive computer keyboard and USB drive that can be purchased with a give budget. Given price lists for keyboards and USB drives and a budget, find the cost to buy them. If it is not possible to buy both items, return .

Example



The person can buy a , or a . Choose the latter as the more expensive option and return .

Function Description

Complete the getMoneySpent function in the editor below.

getMoneySpent has the following parameter(s):

int keyboards[n]: the keyboard prices
int drives[m]: the drive prices
int b: the budget
Returns

int: the maximum that can be spent, or  if it is not possible to buy both items
Input Format

The first line contains three space-separated integers , , and , the budget, the number of keyboard models and the number of USB drive models.
The second line contains  space-separated integers , the prices of each keyboard model.
The third line contains  space-separated integers , the prices of the USB drives.

Constraints

The price of each item is in the inclusive range .
Sample Input 0

10 2 3
3 1
5 2 8
Sample Output 0

9
Explanation 0

Buy the  keyboard and the  USB drive for a total cost of .

Sample Input 1

5 1 1
4
5
Sample Output 1

-1
Explanation 1

There is no way to buy one keyboard and one USB drive because 4+5 > 5, so return -1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMoneySpent(keyboards []int, drives []int, b int) int {
	res := -1
	fmt.Println("keyboards", keyboards)
	for _, k := range keyboards {
		if k >= b {
			continue
		}
		for _, d := range drives {
			n := k + d
			if n > b {
				continue
			}
			if res < n {
				res = n
			}
		}
	}
	return res
}

func stringSliceToInt(sa []string) ([]int, error) {
	r := make([]int, 0, len(sa))
	for _, v := range sa {
		i, err := strconv.Atoi(v)
		if err != nil || i < 1 {
			return r, fmt.Errorf("Only integer value should be provided and should be greater than 0 \n")
		}
		r = append(r, i)
	}
	return r, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the list of number of keyboards separated by space: ")
	scanner.Scan()
	ks := strings.Fields(scanner.Text())
	keyboards, err := stringSliceToInt(ks)
	if err != nil {
		fmt.Printf("Keyboards input error. %s", err)
		return
	}

	fmt.Printf("Enter the list of number of drives separated by space: ")
	scanner.Scan()
	ds := strings.Fields(scanner.Text())
	drives, err := stringSliceToInt(ds)
	if err != nil {
		fmt.Printf("Drives input error. %s", err)
		return
	}

	fmt.Printf("Enter the allocated budget: ")
	scanner.Scan()
	b, err := strconv.Atoi(scanner.Text())
	if err != nil || b < 1 {
		fmt.Printf("Budget input error. Budget should be an intger and should be greater than 0")
		return
	}
	n := getMoneySpent(keyboards, drives, b)
	fmt.Printf("Maximum amount that can be spent is %d \n", n)
}
