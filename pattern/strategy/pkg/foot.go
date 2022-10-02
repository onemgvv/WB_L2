package pkg

import "fmt"

type FootStrategy struct {
}

func (rs *FootStrategy) Route(start, end int) {
	avgSpeed := 4
	total := end - start
	totalTime := total * 80

	fmt.Printf("Foot S:[%d] E:[%d] Avg Speed: [%d] Total: [%d] Total time: [%d]\n", start, end, avgSpeed, total, totalTime)
}