# go-concurrent-word-counter

# 🧠 Go Parallel Word Counter

A high-performance word counting utility written in Go, using goroutines, channels, mutexes, and worker pools. It reads multiple text files concurrently and computes word frequencies — then prints the **most frequent word** found across all files.

---

## 🚀 Features

- ✅ Concurrent file processing using goroutines  
- ✅ Safe shared access with `sync.Mutex`  
- ✅ Channels to distribute file processing jobs  
- ✅ Aggregates and counts words from all files  
- ✅ Finds the **most frequent word** across all inputs

---
