func maxProfit(prices []int) int {
	res := 0
	minPrice := prices[0]
	for _, price := range prices {
		minPrice = min(minPrice, price)
		res = max(res, price-minPrice)
	}

	return res
}
