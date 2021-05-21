/*
HackerRank: https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem

A left rotation operation on an array shifts each of the array's elements  unit to the left. For example, if  left rotations are performed on array , then the array would become . Note that the lowest index item moves to the highest index in a rotation. This is called a circular array.

Given an array  of  integers and a number, , perform  left rotations on the array. Return the updated array to be printed as a single line of space-separated integers.

Function Description

Complete the function rotLeft in the editor below.

rotLeft has the following parameter(s):

int a[n]: the array to rotate
int d: the number of rotations
Returns

int a'[n]: the rotated array
Input Format

The first line contains two space-separated integers  and , the size of  and the number of left rotations.
The second line contains  space-separated integers, each an .

Constraints

Sample Input

5 4
1 2 3 4 5
Sample Output

5 1 2 3 4
Explanation

When we perform  left rotations, the array undergoes the following sequence of changes:


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

func rotLeft(a []string, d int) []string {
	//  arctual rotation required
	rr := d % len(a)
	res := append(a[rr:], a[0:rr]...)
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the array elements separted by space:")
	scanner.Scan()
	a := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	fmt.Printf("Enter number of left rotation you would like to perform on the array: ")
	scanner.Scan()

	d, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Number of rotation provided is not integer")
		return
	}
	ra := rotLeft(a, d)
	fmt.Printf("Final result is %+v \n", ra)
}
