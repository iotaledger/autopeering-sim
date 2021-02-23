package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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
		neighborhoods[peer.ID()] = mgrMap[peer.ID()].GetNeighbors()

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

	record := []string{
		fmt.Sprintf("%v", "Avg"),
		fmt.Sprintf("%v", float64(avgResult.outbound)/float64(N)),
		fmt.Sprintf("%v", float64(avgResult.accepted)/float64(N)),
		fmt.Sprintf("%v", float64(avgResult.rejected)/float64(N)),
		fmt.Sprintf("%v", float64(avgResult.incoming)/float64(N)),
		fmt.Sprintf("%v", float64(avgResult.dropped)/float64(N)),
	}

	output = append(output, record)

	return output
}

/// Distance to strings //////////////////////////////////////////////////////////////////

func distanceToString() (output, outputHisto [][]string) {
	distList := make([]float64, len(allPeers))
	distListHisto := make(map[float64]int)
	var totalDist, numDist uint64
	for i, peer := range allPeers {
		dist := mgrMap[peer.ID()].GetNeighborsDistance()
		minDist := uint32(4294967295)
		for _, d := range dist {
			if d < minDist {
				minDist = d
			}
		}
		numDist++
		totalDist += uint64(minDist)
		distList[i] = float64(minDist)
		distkey, _ := strconv.ParseFloat(strconv.FormatFloat(float64(minDist), 'e', 1, 64), 64)
		distListHisto[distkey]++
	}
	fmt.Println("avgDistance: ", float64(totalDist)/float64(numDist))
	// data
	for i, dist := range distList {
		record := []string{
			fmt.Sprintf("%v", i),
			fmt.Sprintf("%v", dist),
		}
		output = append(output, record)
	}
	// PDF
	for key, num := range distListHisto {
		record := []string{
			fmt.Sprintf("%v", key),
			fmt.Sprintf("%v", float64(num)/float64(numDist)),
		}
		outputHisto = append(outputHisto, record)
	}
	return output, outputHisto
}
func distanceInboundToString() (output, outputHisto [][]string) {
	distList := make([]float64, len(allPeers))
	distListHisto := make(map[float64]int)
	var totalDist, numDist uint64
	for i, peer := range allPeers {
		dist := mgrMap[peer.ID()].GetNeighborsDistanceInbound()
		minDist := uint32(4294967295)
		for _, d := range dist {
			if d < minDist {
				minDist = d
			}
		}
		numDist++
		totalDist += uint64(minDist)
		distList[i] = float64(minDist)
		distkey, _ := strconv.ParseFloat(strconv.FormatFloat(float64(minDist), 'e', 1, 64), 64)
		distListHisto[distkey]++
	}
	fmt.Println("avgDistance: ", float64(totalDist)/float64(numDist))
	// data
	for i, dist := range distList {
		record := []string{
			fmt.Sprintf("%v", i),
			fmt.Sprintf("%v", dist),
		}
		output = append(output, record)
	}
	// PDF
	for key, num := range distListHisto {
		record := []string{
			fmt.Sprintf("%v", key),
			fmt.Sprintf("%v", float64(num)/float64(numDist)),
		}
		outputHisto = append(outputHisto, record)
	}
	return output, outputHisto
}
func distanceOutboundToString() (output, outputHisto [][]string) {
	distList := make([]float64, len(allPeers))
	distListHisto := make(map[float64]int)
	var totalDist, numDist uint64
	for i, peer := range allPeers {
		dist := mgrMap[peer.ID()].GetNeighborsDistanceOutbound()
		minDist := uint32(4294967295)
		for _, d := range dist {
			if d < minDist {
				minDist = d
			}
		}
		numDist++
		totalDist += uint64(minDist)
		distList[i] = float64(minDist)
		distkey, _ := strconv.ParseFloat(strconv.FormatFloat(float64(minDist), 'e', 1, 64), 64)
		distListHisto[distkey]++
	}
	fmt.Println("avgDistance: ", float64(totalDist)/float64(numDist))
	// data
	for i, dist := range distList {
		record := []string{
			fmt.Sprintf("%v", i),
			fmt.Sprintf("%v", dist),
		}
		output = append(output, record)
	}
	// PDF
	for key, num := range distListHisto {
		record := []string{
			fmt.Sprintf("%v", key),
			fmt.Sprintf("%v", float64(num)/float64(numDist)),
		}
		outputHisto = append(outputHisto, record)
	}
	return output, outputHisto
}

/// Distance to strings END //////////////////////////////////////////////////////////////////
