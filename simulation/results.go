package main

import (
	"fmt"
	"os"
	"sort"
)

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func linksToString(links map[int64]int) (output [][]string) {
	keys := make([]int, len(links))
	i := 0
	total := 0
	for k, v := range links {
		keys[i] = int(k)
		i++
		total += v
	}
	sort.Ints(keys)
	for _, key := range keys {
		record := []string{
			fmt.Sprintf("%v", key),
			fmt.Sprintf("%v", float64(links[int64(key)])/float64(total)),
		}
		output = append(output, record)
	}
	return output
}

func convergenceToString(c []Convergence) (output [][]string) {
	for _, line := range c {
		record := []string{
			fmt.Sprintf("%v", line.timestamp.Seconds()),
			fmt.Sprintf("%v", line.counter),
			fmt.Sprintf("%v", line.avgNeighbors),
		}
		output = append(output, record)
	}
	return output
}

func messagesToString(status *StatusMap) (output [][]string) {
	avgResult := StatusSum{}

	//fmt.Printf("\nID\tOUT\tACC\tREJ\tIN\tDROP\n")
	for _, peer := range allPeers {
		neighborhoods[peer.ID()] = protocolMap[peer.ID()].GetNeighbors()

		summary := status.GetSummary(idMap[peer.ID()])

		record := []string{
			fmt.Sprintf("%v", idMap[peer.ID()]),
			fmt.Sprintf("%v", summary.outbound),
			fmt.Sprintf("%v", summary.accepted),
			fmt.Sprintf("%v", summary.rejected),
			fmt.Sprintf("%v", summary.incoming),
			fmt.Sprintf("%v", summary.dropped),
		}

		output = append(output, record)

		avgResult.outbound += summary.outbound
		avgResult.accepted += summary.accepted
		avgResult.rejected += summary.rejected
		avgResult.incoming += summary.incoming
		avgResult.dropped += summary.dropped

	}

	// record := []string{
	// 	fmt.Sprintf("%v", "Avg"),
	// 	fmt.Sprintf("%v", float64(avgResult.outbound)/float64(N)),
	// 	fmt.Sprintf("%v", float64(avgResult.accepted)/float64(N)),
	// 	fmt.Sprintf("%v", float64(avgResult.rejected)/float64(N)),
	// 	fmt.Sprintf("%v", float64(avgResult.incoming)/float64(N)),
	// 	fmt.Sprintf("%v", float64(avgResult.dropped)/float64(N)),
	// }

	// output = append(output, record)

	return output
}

func countMsgToString(cm []int) (output [][]string, avg float64) {
	total := 0
	for _, msg := range cm {
		record := []string{
			fmt.Sprintf("%v", msg),
		}
		output = append(output, record)
		total += msg
	}
	return output, float64(total) / float64(N)
}

func msgInPerTToString() (output [][]string, avg float64) {
	total := 0
	for i := range allPeers {
		record := []string{}
		for k := 0; k < len(msgInPerTList[uint16(i)]); k++ {
			record = append(record, fmt.Sprintf("%v", msgInPerTList[uint16(i)][k]))
			total += msgInPerTList[uint16(i)][k]
		}
		output = append(output, record)
	}
	return output, float64(total) / float64(N*len(msgInPerTList[0]))
}

func msgOutPerTToString() (output [][]string, avg float64) {
	total := 0
	for i := range allPeers {
		record := []string{}
		for k := 0; k < len(msgOutPerTList[uint16(i)]); k++ {
			record = append(record, fmt.Sprintf("%v", msgOutPerTList[uint16(i)][k]))
			total += msgOutPerTList[uint16(i)][k]
		}
		output = append(output, record)
	}
	return output, float64(total) / float64(N*len(msgOutPerTList[0]))
}
