package task

import (
    "fmt"
    "runtime"
)

func GarbageExecute() {
	 // Allocate memory
	 var a []int
	 for i := 0; i < 1000000; i++ {
		 a = append(a, i)
	 }
 
	 // Print memory stats before garbage collection
	 var memStatsBefore runtime.MemStats
	 runtime.ReadMemStats(&memStatsBefore)
	 fmt.Printf("Allocated memory before GC: %d bytes\n", memStatsBefore.Alloc)
 
	 // Force garbage collection
	 runtime.GC()
 
	 // Print memory stats after garbage collection
	 var memStatsAfter runtime.MemStats
	 runtime.ReadMemStats(&memStatsAfter)
	 fmt.Printf("Allocated memory after GC: %d bytes\n", memStatsAfter.Alloc)


	 
    // // Print initial memory stats
    // var memStats runtime.MemStats
    // runtime.ReadMemStats(&memStats)
    // fmt.Printf("Initial Alloc = %v MiB\n", bToMb(memStats.Alloc))
    // fmt.Printf("Total Alloc = %v MiB\n", bToMb(memStats.TotalAlloc))
    // fmt.Printf("Sys = %v MiB\n", bToMb(memStats.Sys))
    // fmt.Printf("NumGC = %v\n", memStats.NumGC)

    // // Allocate memory
    // makeAllocations()

    // // Print memory stats after allocations
    // runtime.ReadMemStats(&memStats)
    // fmt.Printf("After Allocations Alloc = %v MiB\n", bToMb(memStats.Alloc))
    // fmt.Printf("Total Alloc = %v MiB\n", bToMb(memStats.TotalAlloc))
    // fmt.Printf("Sys = %v MiB\n", bToMb(memStats.Sys))
    // fmt.Printf("NumGC = %v\n", memStats.NumGC)

    // // Force garbage collection
    // runtime.GC()

    // // Print memory stats after garbage collection
    // runtime.ReadMemStats(&memStats)
    // fmt.Printf("After GC Alloc = %v MiB\n", bToMb(memStats.Alloc))
    // fmt.Printf("Total Alloc = %v MiB\n", bToMb(memStats.TotalAlloc))
    // fmt.Printf("Sys = %v MiB\n", bToMb(memStats.Sys))
    // fmt.Printf("NumGC = %v\n", memStats.NumGC)
}

func makeAllocations() {
    for i := 0; i < 1000000; i++ {
        _ = make([]byte, 1024) // Allocate 1 KiB
    }
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
