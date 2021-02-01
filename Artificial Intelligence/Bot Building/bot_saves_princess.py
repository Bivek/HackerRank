'''
Link to HackerRank: https://www.hackerrank.com/challenges/saveprincess/problem

Princess Peach is trapped in one of the four corners of a square grid. You are in the center of the grid and can move one step at a time in any of the four directions. Can you rescue the princess?

Input format

The first line contains an odd integer N (3 <= N < 100) denoting the size of the grid. This is followed by an NxN grid. Each cell is denoted by '-' (ascii value: 45). The bot position is denoted by 'm' and the princess position is denoted by 'p'.

Grid is indexed using Matrix Convention

Output format

Print out the moves you will take to rescue the princess in one go. The moves must be separated by '\n', a newline. The valid moves are LEFT or RIGHT or UP or DOWN.

Sample input

3
---
-m-
p--
Sample output

DOWN
LEFT
Task

Complete the function displayPathtoPrincess which takes in two parameters - the integer N and the character array grid. The grid will be formatted exactly as you see it in the input, so for the sample input the princess is at grid[2][0]. The function shall output moves (LEFT, RIGHT, UP or DOWN) on consecutive lines to rescue/reach the princess. The goal is to reach the princess in as few moves as possible.

The above sample input is just to help you understand the format. The princess ('p') can be in any one of the four corners.

Scoring
Your score is calculated as follows : (NxN - number of moves made to rescue the princess)/10, where N is the size of the grid (3x3 in the sample testcase).
'''


#!/usr/bin/python

def displayPathtoPrincess(n,grid):
    #print all the moves here
    bot_coordinate = find_coordinate_for(grid, 'm')
    prince_coordinate = find_coordinate_for(grid, 'p')
    if bot_coordinate[0] is None:
        return ""
    if prince_coordinate[0] is None:
        return ""
    row_diff_wrt_bot = bot_coordinate[0] - prince_coordinate[0]
    col_diff_wrt_bot = bot_coordinate[1] - prince_coordinate[1]
    if row_diff_wrt_bot > 0:
        vertical_movement = "UP"
    elif row_diff_wrt_bot < 0:
        vertical_movement = "DOWN"
    if col_diff_wrt_bot > 0:
        horizontal_movement = "LEFT"
    elif col_diff_wrt_bot < 0:
        horizontal_movement = "RIGHT"
    movement = []
    for i in range(0, abs(row_diff_wrt_bot)):
        movement.append(vertical_movement)
        
    for i in range(0, abs(col_diff_wrt_bot)):
        movement.append(horizontal_movement)
        
    print("\n".join(movement))
    return("\n".join(movement))
    
def find_coordinate_for(grid, marker):
    posx = None
    posy = None
    for x, row in enumerate(grid):
        for y, col in enumerate(row):
            if grid[x][y] == marker:
                posx = x + 1
                posy = y + 1
    return(posx,posy)
    

m = int(input())
grid = [] 
for i in range(0, m): 
    grid.append(input().strip())
                
displayPathtoPrincess(m,grid)