package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	threadPoolSize = 64
	numKeys        = 16384
	keySize        = 16    // 16 bytes for the key
	valueSize      = 16384 // 16368 bytes for the value to make total 16KiB
)

var prefix = ""
var value = ""

func main() {
	prefix = fmt.Sprintf("%016x", time.Now().UnixNano())

	numKeysFlag := flag.Int("num-keys", numKeys, "Number of key-value pairs to create")
	threadPoolSizeFlag := flag.Int("thread-pool-size", threadPoolSize, "Number of workers in the pool")
	valueSizeFlag := flag.Int("value-size", valueSize, "Size of the value in bytes")

	flag.Parse()

	numKeys = *numKeysFlag
	threadPoolSize = *threadPoolSizeFlag
	valueSize = *valueSizeFlag

	valueBytes := make([]byte, valueSize)

	for j := range valueBytes {
		valueBytes[j] = 'a'
	}
	value = string(valueBytes)

	// Connect to etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://control1:2379"}, // Replace with your etcd endpoints
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	// Create a worker pool
	var wg sync.WaitGroup
	workChan := make(chan int, numKeys)

	// Start workers
	for i := 0; i < threadPoolSize; i++ {
		wg.Add(1)
		go worker(cli, workChan, &wg)
	}

	// Distribute work
	startTime := time.Now()
	for i := 0; i < numKeys; i++ {
		workChan <- i
	}
	close(workChan)

	// Wait for all workers to finish
	wg.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Created %d key-value pairs in %v\n", numKeys, elapsedTime)
}

func worker(cli *clientv3.Client, workChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range workChan {
		key := fmt.Sprintf("key-%s-%016d", prefix, i)

		_, err := cli.Put(context.Background(), key, value)
		if err != nil {
			log.Printf("Failed to put key %s: %v", key, err)
		}
	}
}
