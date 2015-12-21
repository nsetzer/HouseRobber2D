# House Robber 2D

http://www.programcreek.com/2014/03/leetcode-house-robber-java/

This project solves the House Robber problem in 2 Dimensions. In a city
there are R rows and C columns of houses aligned on a grid. We know
the total amount of valuables in each house. The goal is to rob houses to
maximize the amount of valuables. However robbing two adjacent houses
will raise an alarm.

This is solved in two ways. A generalized solution is given using
Dynamic Programming. Restricting the problem to positive integers greater
than or equal to zero allows for the problem to be represented as a
maximum flow / min cut problem.


Benchmark and Analysis

The DP solution has a runtime of `O(N*φ^N)` ( Actually, `O(R*φ^C)` ),
while the flow algorithm is `O(N^2)`.
The benchmarks below are for finding the optimum value.

| Solver Function | Grid Shape (Row x Col) | runtime (ms) |
| --------------- | ---------------------- | ------------ |
| DP-Rectangle    |  3,000,000 x 6         |        1,611 |
| Flow-Rectangle  | 20,000,000 x 6         |        2,040 |
| DP-Square       | 447 x 447              |       25,704 |
| Flow-Square     | 10,000 x 10,000        |        1,639 |

Results were generated using the go benchmark tool. The grid sizes are not the
same, so this is not exactly comparing apples to apples. However the results
are pretty clear. The optimized Flow solution allows for processing
significantly larger problems.