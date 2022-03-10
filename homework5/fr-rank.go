package main

import "fmt"

type Rank interface {
	Register(p *Person)
	GetRank(p *Person)
}

type FatRateRank struct {
	clients map[float64]*Person
	frRank  []float64
}

func (fr *FatRateRank) Register(p *Person) {
	fr.clients[p.FatRate] = p
}

func (fr *FatRateRank) GetRank(p *Person) {
	for _, item := range fr.clients {
		fr.frRank = append(fr.frRank, item.FatRate)
	}

	// bubbleSort(&fr.frRank)
	quickSort(&fr.frRank, 0, len(fr.frRank)-1)

	for i, frItem := range fr.frRank {
		if p.FatRate == frItem {
			fmt.Printf("%s的体脂排名是: %d", p.Name, i+1)
		}
	}
}

func bubbleSort(arr *[]float64) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}

}

func quickSort(arr *[]float64, start, end int) {
	pivotIdx := (start + end) / 2
	pivotVal := (*arr)[pivotIdx]
	l, r := start, end
	for l <= r {
		for (*arr)[l] < pivotVal {
			l++
		}
		for (*arr)[r] > pivotVal {
			r--
		}
		if l >= r {
			break
		}
		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(arr, start, r)
	}
	if l < end {
		quickSort(arr, l, end)
	}

}
