package dicesprobability

func dicesProbability(n int) []float64 {
	db := make([][]float64, n)
	for i := range db {
		db[i] = make([]float64, 5*i+6)
	}
	for i := 0; i < 6; i++ {
		db[0][i] = 1.0 / 6.0
	}
	for i := 0; i < len(db)-1; i++ {
		for j := range db[i] {
			for k := 0; k < 6; k++ {
				db[i+1][j+k] += db[i][j] / 6.0
			}
		}
	}
	return db[n-1]
}
