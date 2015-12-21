/*

The 2D House Robber Problem

Given a City laid out on a grid, there are R rows and C columns of houses.
Each house is known to contain some amount of valuables v_ij.
Your task is to rob as many houses as possible to maximize the amount of loot.
However there is a security system in place and if you rob two adjacent houses
an alarm will go off.

Find the optimum set of non-adjacent houses to rob.

Example:
            Houses:  alignment:
          10 20 10     0  1  0
          20 30 20 =>  1  0  1
          10 20 10     0  1  0
        This alignment results in the maximum value of 80.

Dynamic Programming

Dynamic Programming requires comparing all possible alignments to
determine which alignment produces the optimal score. I first compute
the set of valid alignments in a single row then generate the pairs of
alignments that can be stacked vertically. Next for each row the optimal
alignment given the solution to all previous rows is computed. Once the
final row is processed the overall maximum value can be found. The solutions
can then be read back in reverse to determine which alignment produced this
result

Maximum Flow

This problem is an example of Maximum Bipartite Matching and can be
represented a Maximum Flow problem. Partition the input table into
two sets U and V such that no element in either set is adjacent to another
within the same set. Create a source vertex S and a sink vertex T. For
every node u in U with weight w connect an edge from s to u with weight w.
For every node v in V with weight w connect an edge from v to t with
weight w. Finally compute the maximum flow through the network. The
maximum flow equals the minimum cut and is isomorphic to the selection
of vertices's. Thus the maximum value is equal to the total value minus the
maximum flow. The code here is loosely based on the Dinic's Algorithm for Max
Flow. The Dynamic Trees are embedded inside the matrix structure.


*/
package hr2d

// Use dynamic programming to determine maximum value
func SolveDP(R, C int, table [][]int) (int) {
	alignments := generate_alignments(C)
	pairs := generate_alignment_pairs(alignments)

	N := len(alignments)

	// predecessors array each index p[i][j] is the index of the
	// alignment from the previous row.

	// current and previous best scores for the current row i
	// given the set of all possible valid alignments.
	dc := make([]int, N)
	dp := make([]int, N)

	for j, ali := range alignments {
		dp[j] = ali.Dot(table[0])
	}

	row_scores := make([]int, N)

	// process each row
	for i := 1; i < R; i++ {

		// determine alignment scores for the current row
		for j, ali := range alignments {
			row_scores[j] = ali.Dot(table[i])
			dc[j] = 0 //reset this rows current best score
		}

		// determine best score for each alignment given previous answers
		for a := 0; a < N; a++ {
			for _, b := range pairs[a] {
				c := dp[a] + row_scores[b]
				if c > dc[b] {
					dc[b] = c
				}
			}
		}

		// copy current best scores to the previous scores table
		for j := 0; j < N; j++ {
			dp[j] = dc[j]
		}
	}

	var maximum_value int
	for _, v := range dp {
		if v > maximum_value {
			maximum_value = v
		}
	}

	return maximum_value
}

// Use dynamic programming to determine maximum value and alignment
func SolveDP2(R, C int, table [][]int) (int, Alignments, error) {
	alignments := generate_alignments(C)
	pairs := generate_alignment_pairs(alignments)

	N := len(alignments)

	// predecessors array each index p[i][j] is the index of the
	// alignment from the previous row.
	p := make([][]int, R)

	//p[0] = make([]int, N) p[0] is not interesting

	// current and previous best scores for the current row i
	// given the set of all possible valid alignments.
	dc := make([]int, N)
	dp := make([]int, N)

	for j, ali := range alignments {
		dp[j] = ali.Dot(table[0])
	}

	row_scores := make([]int, N)

	// process each row
	for i := 1; i < R; i++ {

		p[i] = make([]int, N)

		// determine alignment scores for the current row
		for j, ali := range alignments {
			row_scores[j] = ali.Dot(table[i])
			dc[j] = 0 //reset this rows current best score
		}

		// determine best score for each alignment given previous answers
		for a := 0; a < N; a++ {
			for _, b := range pairs[a] {
				c := dp[a] + row_scores[b]
				if c > dc[b] {
					dc[b] = c
					p[i][b] = a
				}
			}
		}

		// copy current best scores to the previous scores table
		for j := 0; j < N; j++ {
			dp[j] = dc[j]
		}
	}

	final := make(Alignments, R)

	// determine which alignment generated the highest score in the final row
	var a, c int // best alignment index; best cost
	for j, v := range dp {
		if v > c {
			c = v
			a = j
		}
	}

	// walk the parent array in reverse
	// select the optimal alignment at each row
	for i := R - 1; i > 0; i-- {
		final[i] = alignments[a]
		a = p[i][a]
	}
	final[0] = alignments[a]

	return c, final, nil
}

// Use Maximum Flow to determine maximum value
func SolveFlow(R, C int, table [][]int) (int) {

	var total_value int;
	var max_flow int;

	// O(N^2)
	// determine total value of all houses
	for i:=0; i<R; i++ {
		for j:=0; j<C; j++ {
			total_value += table[i][j];
		}
	}

	// O(N * 4*(N/2)) => O(2*N^2) => O(N^2)
	// determine the maximum flow, by finding the min cut
	// of all paths from the source to the sink.
	// This loop iterates over all vertices's in U, then determines
	// the min cut for all adjacent vertices's in V.
	for i:=0; i<R; i++ {
		for j:=i%2; j<C; j+=2 {
			if i>0 {
				max_flow += update_path(table,i,j,i-1,j);
			}
			if j>0 {
				max_flow += update_path(table,i,j,i,j-1);
			}
			if j+1 < C {
				max_flow += update_path(table,i,j,i,j+1);
			}
			if i+1 < R {
				max_flow += update_path(table,i,j,i+1,j);
			}
		}
	}

	maximum_value := total_value - max_flow;

	return maximum_value
}

// compute blocking flow through s->u->v->t
func update_path(table [][]int, i, j, x, y int) int {

	u := table[i][j]
	v := table[x][y]

	// determine the blocking flow
	f := u
	if v < u {
		f = v
	}

	// ensure that this path has not been cut already
	if f > 0 {
		table[i][j] -= f
		table[x][y] -= f
	}

	return f;
}