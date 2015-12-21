//go test
//go test -bench=.*

package hr2d

import "testing"
import "math"
import "math/rand"

func Test2Align(t *testing.T) {

	table := [][]int{
		[]int{5, 0},
		[]int{10,4} }

	expected := 10
	cost, alignments, err := SolveDP2(2, 2, table)

	if err != nil {
		t.Fatal(err)
	}
	if cost != expected {
		for _, ali := range alignments {
			t.Logf("%s",ali)
		}
		t.Fatalf("Found Score of %d. Expected %d.",cost,expected)
	}
}

func Test2DP(t *testing.T) {

	table := [][]int{
		[]int{5, 0},
		[]int{10,4} }

	expected := 10
	cost := SolveDP(2, 2, table)

	if cost != expected {
		t.Fatalf("Found Score of %d. Expected %d.",cost,expected)
	}
}


func Test2Flow(t *testing.T) {

	table := [][]int{
		[]int{5, 0},
		[]int{10,4} }

	expected := 10
	cost := SolveFlow(2, 2, table)

	if cost != expected {
		t.Fatalf("Found Score of %d. Expected %d.",cost,expected)
	}
}

func Test3Align(t *testing.T) {

	table := [][]int{
		[]int{10, 20, 10},
		[]int{20, 30, 20},
		[]int{10, 20, 10}}

	expected := 80
	cost, alignments, err := SolveDP2(3, 3, table)

	if err != nil {
		t.Fatal(err)
	}
	if cost != expected {
		for _, ali := range alignments {
			t.Logf("%s",ali)
		}
		t.Fatalf("Found Score of %d. Expected %d.",cost,expected)
	}
}

func Test3DP(t *testing.T) {

	table := [][]int{
		[]int{10, 20, 10},
		[]int{20, 30, 20},
		[]int{10, 20, 10}}

	expected := 80
	cost := SolveDP(3, 3, table)

	if cost != expected {
		t.Fatalf("Found Score of %d. Expected %d.",cost,expected)
	}
}


func Test3Flow(t *testing.T) {

	table := [][]int{
		[]int{10, 20, 10},
		[]int{20, 30, 20},
		[]int{10, 20, 10}}

	expected := 80
	cost := SolveFlow(3, 3, table)

	if cost != expected {
		t.Fatalf("Found Score of %d. Expected %d.",cost,expected)
	}
}

func createBenchmarkTable(R,C int) [][]int {
	table := make([][]int,R)

	rand.Seed(4)

	for i:=0;i<R;i++ {
		table[i] = make([]int,C)
		for j:=0;j<C;j++ {
			table[i][j] = rand.Intn(10)
		}
	}

	return table
}

// benchmark runtime given a retangle of height N
// by contraining the number of columns very little work is needed
// to be done per row.
func BenchmarkDPRectangle(b *testing.B) {

	C:=6
	table := createBenchmarkTable(b.N,C)

	b.ResetTimer()

	SolveDP(b.N, C, table)

}


// benchmark runtime given a square whose area is equal to N
// setting each side to sqrt(N) contrains the size of the problem
// but will result in n^3 runtime in the worst case.
func BenchmarkDPSquare(b *testing.B) {

	R := int(math.Sqrt( float64(b.N)))
	table := createBenchmarkTable(R,R)

	b.ResetTimer()

	SolveDP(R, R, table)

}

func BenchmarkFlowRectangle(b *testing.B) {

	C:=6
	table := createBenchmarkTable(b.N,C)

	b.ResetTimer()

	SolveFlow(b.N, C, table)

}

func BenchmarkFlowSquare(b *testing.B) {

	R := int(math.Sqrt( float64(b.N)))
	table := createBenchmarkTable(R,R)

	b.ResetTimer()

	SolveFlow(R, R, table)

}
