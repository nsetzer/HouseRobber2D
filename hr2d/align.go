package hr2d

type Alignment []int
type Alignments []Alignment

// Return True if the alignment contains no adjacent true values.
func (a *Alignment) Valid() bool {
	var last int
	for _, v := range *a {
		if last == 1 && v == 1 {
			return false
		}
		last = v
	}
	return true
}

// compute dot product of two vectors of equal length.
func (a *Alignment) Dot(vector []int) int {
	var total int
	for i, v := range *a {
		total += vector[i] * v
	}
	return total
}

// Generate all the set of all valid alignments for C columns
//
// permute all combinations of a C-bit binary number and return
// only the combinations which are Valid.
func generate_alignments(C int) (o Alignments) {
	// C-bit number has N=2^C possible combinations
	N := C * C
	ali := make(Alignment, C)
	for i := 0; i < N; i++ {
		if ali.Valid() {
			o = append(o, append(Alignment{}, ali...))
		}
		// add one and carry
		for j := 0; j < C; j++ {
			ali[j] = 1 - ali[j]
			if ali[j] == 1 {
				break
			}
		}
	}
	return
}

// generate the list of indices for each alignment i which are valid
// when situated on top of each other.
func generate_alignment_pairs(alignments Alignments) [][]int {
	pairs := make([][]int, len(alignments))
	for i, a := range alignments {
		for j, b := range alignments {
			if a.Dot(b) == 0 {
				pairs[i] = append(pairs[i], j)
			}
		}
	}
	return pairs
}
