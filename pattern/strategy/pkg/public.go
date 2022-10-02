package pkg

import "fmt"

type PublicStrategy struct {
}

func (rs *PublicStrategy) Route(start, end int) {
	avgSpeed := 35

	total := end - start
	totalTime := total * 40

	fmt.Printf("Public S:[%d] E:[%d] Avg Speed: [%d] Total: [%d] Total time: [%d]\n", start, end, avgSpeed, total, totalTime)
}