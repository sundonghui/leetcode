package cancompletecircuit

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumGas, sumCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumCost += cost[j]
			sumGas += gas[j]
			if sumCost > sumGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
