
package main

import (
    "fmt"
    "runtime"
  "app/test"
)

func main() {
    // Run garbage collection to get accurate memory statistics
    runtime.GC()

    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)

    fmt.Println("Initial Memory Usage:")
    fmt.Println("Alloc:", memStats.Alloc)
    fmt.Println("Total Alloc:", memStats.TotalAlloc)
    fmt.Println("HeapAlloc:", memStats.HeapAlloc)
    fmt.Println("HeapSys:", memStats.HeapSys)

    // Your code here
    test.Mytest(100)
    // Run garbage collection again to get updated memory statistics
    runtime.GC()
    runtime.ReadMemStats(&memStats)

    fmt.Println("Final Memory Usage:")
    fmt.Println("Alloc:", memStats.Alloc)
    fmt.Println("Total Alloc:", memStats.TotalAlloc)
    fmt.Println("HeapAlloc:", memStats.HeapAlloc)
    fmt.Println("HeapSys:", memStats.HeapSys)
}
