package main

import (
	"fmt"
	"os"
	//"sort"
)

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func knownPeersToString(c []Convergence) (output [][]string) {
	for _, line := range c {
		record := []string{
			fmt.Sprintf("%v", line.timestamp.Seconds()),
			fmt.Sprintf("%v", float64(line.counter)/float64(N)*100.),
			fmt.Sprintf("%v", line.avgKnown),
		}
		output = append(output, record)
	}
	return output
}

func sentMsgPdfToString(m map[uint16][]byte) (output [][]string) {
	keys := make(map[int]int)
	for _, v := range m {
		keys[len(v)]++
	}
	for l, key := range keys {
		record := []string{
			fmt.Sprintf("%v", l),
			fmt.Sprintf("%v", float64(key)/float64(N)),
		}
		output = append(output, record)
	}
	return output
}

func sentMsgToString(m map[uint16][]byte) (output [][]string) {
	for pID, msgList := range m {
		record := []string{
			fmt.Sprintf("%v", pID),
			fmt.Sprintf("%v", len(msgList)),
		}
		output = append(output, record)
	}
	return output
}
