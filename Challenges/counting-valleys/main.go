/*
HackerRank: https://www.hackerrank.com/challenges/counting-valleys/problem

An avid hiker keeps meticulous records of their hikes. During the last hike that took exactly  steps, for every step it was noted if it was an uphill, , or a downhill,  step. Hikes always start and end at sea level, and each step up or down represents a  unit change in altitude. We define the following terms:

A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.
A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.
Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.

Example



The hiker first enters a valley  units deep. Then they climb out and up onto a mountain  units high. Finally, the hiker returns to sea level and ends the hike.

Function Description

Complete the countingValleys function in the editor below.

countingValleys has the following parameter(s):

int steps: the number of steps on the hike
string path: a string describing the path
Returns

int: the number of valleys traversed
Input Format

The first line contains an integer , the number of steps in the hike.
The second line contains a single string , of  characters that describe the path.

Constraints

Sample Input

8
UDDDUDUU
Sample Output

1
Explanation

If we represent _ as sea level, a step up as /, and a step down as \, the hike can be drawn as:

_/\      _
   \    /
    \/\/
The hiker enters and leaves one valley.

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
	fmt.Printf("Enter the number of steps taken by hiker:")
	scanner.Scan()
	s, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Errorf("Steps should be an integer. Error: %v \n", err)
		return
	}
	fmt.Printf("Enter the path taken by the hiker (path should cosist of U & D characters only and should match with the steps provided and should have end at the sea level:")
	scanner.Scan()
	p := strings.TrimSpace(scanner.Text())

	vc, err := countingValleys(s, p)
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}
	fmt.Printf("Number of valley crossed is %d \n", vc)
}

func countingValleys(steps int, path string) (int, error) {
	if len(path) != steps {
		return -1, fmt.Errorf("The steps count provided and the length of the path did not match \n")
	}
	var vc, seaLevel, prevSeaLevel int
	for i := 0; i < steps; i++ {
		prevSeaLevel = seaLevel
		if path[i] == 'D' {
			seaLevel -= 1
		} else if path[i] == 'U' {
			seaLevel += 1
		} else {
			return -1, fmt.Errorf("Path should contain characters 'U' and 'D' only. Provided %s \n", path)
		}

		if prevSeaLevel < 0 && seaLevel == 0 {
			vc += 1
		}
	}
	if seaLevel != 0 {
		return -1, fmt.Errorf("The input path is wrong and the hiker could not made to sea level at the end. Provided path: %s \n", path)
	}
	return vc, nil
}
