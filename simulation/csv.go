package simulation

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/iotaledger/hive.go/autopeering/peer"
)

func initCSV(records [][]string, filename string) error {
	createDirIfNotExist("data")
	f, err := os.Create("data/result_" + filename + ".csv")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err = w.WriteAll(records); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	return err
}

func WriteCSV(records [][]string, filename string, header ...[]string) error {
	var err error
	if header != nil {
		err = initCSV(header, filename) // requires format into [][]string
	} else {
		err = initCSV([][]string{}, filename) // requires format into [][]string
	}
	if err != nil {
		log.Fatalln("error initializing csv:", err)
	}
	f, err := os.OpenFile("data/result_"+filename+".csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	return err
}

func WriteAdjlist(nodeMap map[peer.ID]Node, filename string) error {
	const separator = ' '

	f, err := os.Create("data/result_" + filename + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for id, node := range nodeMap {
		_, _ = w.WriteString(id.String())
		out := node.GetOutgoingNeighbors()
		for _, n := range out {
			_, _ = w.WriteRune(separator)
			_, _ = w.WriteString(n.ID().String())
		}
		_, _ = w.WriteRune('\n')
	}
	return w.Flush()
}
