/*
HackerRank: https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem

There is a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. The player can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus  or . The player must avoid the thunderheads. Determine the minimum number of jumps it will take to jump from the starting postion to the last cloud. It is always possible to win the game.

For each game, you will get an array of clouds numbered  if they are safe or  if they must be avoided.

Example

Index the array from . The number on each cloud is its index in the list so the player must avoid the clouds at indices  and . They could follow these two paths:  or . The first path takes  jumps while the second takes . Return .

Function Description

Complete the jumpingOnClouds function in the editor below.

jumpingOnClouds has the following parameter(s):

int c[n]: an array of binary integers
Returns

int: the minimum number of jumps required
Input Format

The first line contains an integer , the total number of clouds. The second line contains  space-separated binary integers describing clouds  where .

Constraints

Output Format

Print the minimum number of jumps needed to win the game.

Sample Input 0

7
0 0 1 0 0 1 0
Sample Output 0

4
Explanation 0:
The player must avoid  and . The game can be won with a minimum of  jumps:

jump(2).png

Sample Input 1

6
0 0 0 0 1 0
Sample Output 1

3
Explanation 1:
The only thundercloud to avoid is . The game can be won in  jumps:

jump(5).png
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the sequence of thunderheads (represented as 1) and cumulus (represented as 0) clouds seperated by space: ")
	scanner.Scan()
	cs := scanner.Text()
	ias := strings.Split(cs, " ")
	ia, err := sliceAtoi(ias)
	if err != nil {
		fmt.Println(err)
		return
	}
	steps, err := jumpingOnClouds(ia)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The number of steps required to complete the game is %d", steps)
}

func sliceAtoi(c []string) ([]int, error) {
	var r []int
	for _, v := range c {
		if v == "0" || v == "1" {
			i, _ := strconv.Atoi(v)
			r = append(r, i)
		} else {
			return r, fmt.Errorf("Invalid input provided. Only sequence of 0 and 1 are allowed but received %v", c)
		}
	}
	return r, nil
}

// This function expects input array containing values only 0 and 1
// If the array contain elements other than 0 and 1 and error is returned
// Similarly if there is no way to reach to destination as well and error is returned
func jumpingOnClouds(c []int) (int, error) {

	if c[0] == 1 || c[len(c)-1] == 1 {
		return -1, errors.New("There is no way to reach the final state safely. Please correct you input next time")
	}
	// Stores selected path
	var sp []string
	// starting index to consider for each iteration
	si := 0
	for i, _ := range c {
		if si == len(c)-1 {
			break
		}
		if i < si {
			continue
		}

		if si == len(c)-2 {
			sp = append(sp, fmt.Sprintf("[%d -> %d]", si, len(c)-1))
			break
		} else {
			if c[i+1] == 1 && c[i+2] == 1 {
				//This problem in unsolvable
				return -1, errors.New("There is no way to reach the final state safely. Please correct you input next time")
			}
			if c[i+2] == 0 {
				sp = append(sp, fmt.Sprintf("[%d -> %d]", si, i+2))
				si = i + 2
			} else {
				sp = append(sp, fmt.Sprintf("[%d -> %d]", si, i+1))
				si = i + 1
			}
		}
	}
	fmt.Println(sp)
	return len(sp), nil

}
