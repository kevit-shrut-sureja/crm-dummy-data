package main

import (
	"math/rand/v2"
	"time"
)

func probSingle(prob float32, data any) any {
	if prob != 0 && rand.Float32() > 1-prob {
		return nil
	}

	return data
}

func probArray[T any](prob float32, data []T, selectOne bool) []T {
	if prob != 0 && rand.Float32() > 1-prob {
		return []T{}
	}

	return getRandomSubset(data, selectOne)
}

func getRandomSubset[T any](arr []T, selectOne bool) []T {
	n := len(arr)
	if n == 0 {
		return []T{}
	}

	// Create a copy to avoid modifying the original slice.
	copiedArr := append([]T{}, arr...)

	// Shuffle the copied slice.
	rand.Shuffle(n, func(i, j int) { copiedArr[i], copiedArr[j] = copiedArr[j], copiedArr[i] })

	if selectOne {
		return []T{copiedArr[0]}
	}

	// Choose a random subset size from 1 to n.
	subsetSize := rand.IntN(n) + 1
	return copiedArr[:subsetSize]
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

func ptr[T any](s T) *T {
	return &s
}

func randomTimePicker() string {
	start := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, time.February, 31, 23, 59, 59, 0, time.UTC)

	// Generate a random duration between start and end
	randomDuration := time.Duration(rand.Int64N(int64(end.Sub(start))))
	randomTime := start.Add(randomDuration)

	// Format the time in the required format
	timeFormat := "2006-01-02T15:04:05.000Z"

	return randomTime.Format(timeFormat)
}
