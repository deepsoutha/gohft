package src


func Strategy(priceSlice []float64,period int) int {


	ma1 := MA(priceSlice, period/2)
	ma2 := MA(priceSlice, period)
	if ma1[len(ma1)-1] > ma2[len(ma2)-1] && ma1[len(ma1)-2] < ma2[len(ma2)-2] {
		return 1
	} else if ma1[len(ma1)-1] < ma2[len(ma2)-1] && ma1[len(ma1)-2] > ma2[len(ma2)-2] {
		return -1
	} else {
		return 0
	}
}

func MA(data []float64, windowSize int) []float64 {
	ma := make([]float64, len(data)-windowSize+1)
	sum := 0.0
	for i := 0; i < windowSize; i++ {
		sum += data[i]
	}
	ma[0] = sum / float64(windowSize)
	for i := windowSize; i < len(data); i++ {
		sum += data[i] - data[i-windowSize]
		ma[i-windowSize+1] = sum / float64(windowSize)
	}
	return ma
}
