package piscine

func ft_profit(prices []int) int {
	min := prices[0]
	pose_min := 0
	max := prices[0]
	var dif int
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			if max-min > dif {
				dif = max - min
			}
			min = prices[i]
			pose_min = i
			max = prices[i]
		} else if prices[i] > max && i > pose_min {
			max = prices[i]
		}
		if max-min > dif {
			dif = max - min
		}
	}
	return dif
}
