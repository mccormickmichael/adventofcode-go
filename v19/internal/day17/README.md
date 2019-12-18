# Day 17

## Set and Forget

### Part 1 - Locating Intersecttions

Part 1 is straightforward. Run an Intcode program to generate scaffolding
represented by an array of characters and locate the intersections.

First cut, generate the scaffolding map as [][]byte to see its extent and
whether it is rectangular. If it's not regularly rectangular, fill in
the missing bits until it is.

Looking for intersections: Scan then interior (i.e. one away from the edge
on all sides). If a cell is '#' and its neighbors up, down, left, and right
are also '#', then it's an intersection.

Finally, output a map of the scaffolding to see what it looks like.

### Part 2 - Navigation Subroutines

This was not at all what I expected for Part 2. It turned out that generating
the instructions and parceling them into subroutines was best done by hand.
I expected some trickery, splitting turns and distances across subroutines,
but no such shenanigans.

Generating the input for the program was then just a matter of concatenation
and letting it run.