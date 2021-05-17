/*
HackerRank: https://www.hackerrank.com/challenges/drawing-book/problem

A teacher asks the class to open their books to a page number. A student can either start turning pages from the front of the book or from the back of the book. They always turn pages one at a time. When they open the book, page  is always on the right side:

image

When they flip page , they see pages  and . Each page except the last page will always be printed on both sides. The last page may only be printed on the front, given the length of the book. If the book is  pages long, and a student wants to turn to page , what is the minimum number of pages to turn? They can start at the beginning or the end of the book.

Given  and , find and print the minimum number of pages that must be turned in order to arrive at page .

Example



Untitled Diagram(4).png

Using the diagram above, if the student wants to get to page , they open the book to page , flip  page and they are on the correct page. If they open the book to the last page, page , they turn  page and are at the correct page. Return .

Function Description

Complete the pageCount function in the editor below.

pageCount has the following parameter(s):

int n: the number of pages in the book
int p: the page number to turn to
Returns

int: the minimum number of pages to turn
Input Format

The first line contains an integer , the number of pages in the book.
The second line contains an integer, , the page to turn to.

Constraints

Sample Input 0

6
2
Sample Output 0

1
Explanation 0

If the student starts turning from page , they only need to turn  page:

Untitled Diagram(6).png

If a student starts turning from page , they need to turn  pages:

Untitled Diagram(3).png

Return the minimum value, .

Sample Input 1

5
4
Sample Output 1

0
Explanation 1

If the student starts turning from page , they need to turn  pages:

Untitled Diagram(4).png

If they start turning from page , they do not need to turn any pages:

Untitled Diagram(5).png

Return the minimum value, 0.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the number of pages in a book: ")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Enter the page number you want to open: ")
	scanner.Scan()
	p, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println(err)
		return
	}
	pc, err := pageCount(n, p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("You need to turn %d pages to reach to the desired page \n", pc)
}

func pageCount(n int, p int) (int, error) {
	var r int
	if p < 1 || p > n {
		return -1, fmt.Errorf("Desired page %d is out of range from the book having %d pages", p, n)
	}
	//Diff between last page and desired page
	diffb := n - p
	// Diff between desired page and first page
	difff := p - 1
	f := func(num int) int {
		return num/2 + num%2
	}
	if difff < diffb {
		fmt.Println("Start turning the page from the front")
		r = f(difff)
	} else {
		fmt.Println("Start turning the page from the back")
		if n%2 == 0 {
			r = f(diffb)
		} else {
			r = diffb / 2
		}
	}
	return r, nil
}
