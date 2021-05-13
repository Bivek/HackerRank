/*
HackerRank URL: https://www.hackerrank.com/challenges/sock-merchant/problem

There is a large pile of socks that must be paired by color. Given an array of strings representing the color of each sock, determine how many pairs of socks with matching colors there are.

Example


There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .

Function Description

Complete the sockMerchant function in the editor below.

sockMerchant has the following parameter(s):

int n: the number of socks in the pile
int ar[n]: the colors of each sock
Returns

int: the number of pairs
Input Format

The first line contains an integer , the number of socks represented in .
The second line contains  space-separated integers, , the colors of the socks in the pile.

Constraints

 where
Sample Input

STDIN                       Function
-----                       --------
9                           n = 9
10 20 20 10 10 30 50 10 20  ar = ['10', '20', '20', '10', '10', '30', '50', '10', '20']

Sample Output
3

Explanation

10 -> 10
10 -> 10
20 -> 20
50
30
20

There are three pairs of socks.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER and err
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY ar
 */

func sockMerchant(n int, ar []string) (int, error) {
	found := 0
	if n != len(ar) {
		return -1, errors.New("Values provided and the count do not match")
	}
	m := make(map[string]int)
	for _, v := range ar {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	for _, v := range m {
		found += v / 2
	}
	return found, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the number of socks in the pile: ")
	scanner.Scan()
	socks := scanner.Text()
	n, err := strconv.Atoi(socks)
	if err != nil {
		fmt.Println("Encounter an error ", err)
		return
	}
	fmt.Printf("Enter the colour of each sock by space between the multiple socks: ")
	scanner.Scan()
	sc := scanner.Text()
	ar := strings.Split(sc, " ")
	result, err := sockMerchant(n, ar)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
