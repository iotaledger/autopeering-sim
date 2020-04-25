package simulation

import (
	"fmt"
	"os"
	"sort"

	"github.com/iotaledger/autopeering-sim/simulation/config"
	"github.com/iotaledger/hive.go/identity"
)

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func LinksToString(linkList []Link) (output [][]string) {
	links := linkSurvival(linkList)
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

func ConvergenceToString() (output [][]string) {
	c := RecordConv.convergence
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

func MessagesToString(nodeMap map[identity.ID]Node, status *StatusMap) (output [][]string) {
	avgResult := StatusSum{}

	//fmt.Printf("\nID\tOUT\tACC\tREJ\tIN\tDROP\n")
	for id := range nodeMap {
		summary := status.GetSummary(id)
		record := []string{
			fmt.Sprintf("%v", id),
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

	N := config.NumberNodes()
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
