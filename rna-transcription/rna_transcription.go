// Package strand provides a suite of functions about the DNA calculations
package strand

// ToRNA convert a dna sequence into and rna sequence
func ToRNA(dna string) string {
	conversion := map[uint8]string{'G': "C", 'C': "G", 'T': "A", 'A': "U"}
	res := ""
	for i := range dna {
		res += conversion[dna[i]]
	}
	return res
}
