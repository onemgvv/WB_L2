package pkg

import "fmt"

type PrivateStrategy struct {
}

func (rs *PrivateStrategy) Route(start, end int) {
	avgSpeed := 30
	trafficJam := 2
	total := end - start
	totalTime := total * 40 * trafficJam

	fmt.Printf("Private S:[%d] E:[%d] Avg Speed: [%d] Traffic jam: [%d] Total: [%d] Total time: [%d]\n", start, end, avgSpeed, trafficJam, total, totalTime)
}