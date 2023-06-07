package storage

import (
	"context"
	"fmt"
	"math/rand"
	"push_notification/entity"
	"runtime"
	"testing"
)

func benchmarkMemory_PushAverage(m Storage, b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		id := rand.Intn(1000)
		m.Push(ctx, id, entity.UnreadMessageNotification{Count: i})
	}
	b.StopTimer()
	PrintMemUsage()
}

func benchmarkMemory_PushNewItem(m Storage, b *testing.B) {
	ctx := context.Background()
	counter := 0
	for i := 0; i < b.N; i++ {
		m.Push(ctx, i, entity.UnreadMessageNotification{Count: i})
		counter++
	}
	b.StopTimer()
	b.Log("for ", b.N, "notifications: ")
	PrintMemUsage()
}

func BenchmarkMemoryWithChannel_PushNewItem(b *testing.B) {
	benchmarkMemory_PushNewItem(NewMemoryWithChannel(1000), b)
}

func BenchmarkMemoryWithList_PushNewItem(b *testing.B) {
	benchmarkMemory_PushNewItem(NewMemoryWithList(1000), b)
}
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
