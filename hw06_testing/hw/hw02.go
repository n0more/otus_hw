package hw

import (
	"fmt"

	"github.com/n0more/otus_hw/hw06_testing/printer"
	"github.com/n0more/otus_hw/hw06_testing/reader"
	"github.com/n0more/otus_hw/hw06_testing/types"
)

func Runhw02(path string) string {
	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	return printer.PrintStaff(staff)
}
