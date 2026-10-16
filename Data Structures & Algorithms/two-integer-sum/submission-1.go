func twoSum(nums []int, target int) []int {
	historico := make(map[int]int)
	for i, num := range nums {
		  diferenca :=  target - num
		if indiceAntigo, encontrado := historico[diferenca]; encontrado {
			return []int{indiceAntigo, i}
		}
		historico[num] = i
	}
	return []int{}
}
