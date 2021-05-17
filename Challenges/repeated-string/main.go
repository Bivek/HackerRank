/*

HackerRank: https://www.hackerrank.com/challenges/repeated-string/problem

There is a string, , of lowercase English letters that is repeated infinitely many times. Given an integer, , find and print the number of letter a's in the first  letters of the infinite string.

Example


The substring we consider is , the first  characters of the infinite string. There are  occurrences of a in the substring.

Function Description

Complete the repeatedString function in the editor below.

repeatedString has the following parameter(s):

s: a string to repeat
n: the number of characters to consider
Returns

int: the frequency of a in the substring
Input Format

The first line contains a single string, .
The second line contains an integer, .

Constraints

For  of the test cases, .
Sample Input

Sample Input 0

aba
10
Sample Output 0

7
Explanation 0
The first  letters of the infinite string are abaabaabaa. Because there are  a's, we return .

Sample Input 1

a
1000000000000
Sample Output 1

1000000000000
Explanation 1
Because all of the first  letters of the infinite string are a, we return .
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter string: ")
	scanner.Scan()
	s := strings.TrimSpace(scanner.Text())
	fmt.Printf("Enter the value of n: ")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	rs := repeatedString(s, n)

	fmt.Printf("'a' is repated %d times", rs)
}

func repeatedString(s string, n int) int {
	// Write your code here
	var result int
	sz := len(s)
	t := n / sz
	r := n % sz
	ss := s[0:r]
	f := func(s string) int {
		var res int
		for _, v := range s {
			if string(v) == "a" {
				res += 1
			}
		}
		return res
	}
	result = f(s)*t + f(ss)
	return result
}
