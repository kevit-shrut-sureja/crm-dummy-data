package main

import (
	"fmt"
	"math/rand/v2"
)

func probSingle(prob float32, data any) any {
	if rand.Float32() > 1-prob {
		fmt.Println("------------------")
		return nil
	}

	return data
}

func probArray[T any](prob float32, data []T, selectOne bool) []T {
	if rand.Float32() > 1-prob {
		fmt.Println("------------------")
		return []T{}
	}

	return getRandomSubset(data, selectOne)
}

func getRandomSubset[T any](arr []T, selectOne bool) []T {
	n := len(arr)
	if n == 0 {
		return []T{}
	}

	// Shuffle the slice in place.
	rand.Shuffle(n, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	if selectOne {
		return []T{arr[0]}
	}

	// Choose a random subset size from 0 to n.
	subsetSize := rand.IntN(n) + 1
	return arr[:subsetSize]
}

func safePtr[T any](v any) *T {
	if v == nil {
		return nil
	}
	x, ok := v.(T)
	if !ok {
		return nil
	}
	return &x
}

func safeValue[T any](v any) T {
	var empty T
	if v == nil {
		return empty
	}
	x, ok := v.(T)
	if !ok {
		return empty
	}
	return x
}
