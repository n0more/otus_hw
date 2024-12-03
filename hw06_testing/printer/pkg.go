package printer

import (
	"fmt"
	"strings"

	"github.com/n0more/otus_hw/hw06_testing/types"
)

func PrintStaff(staff []types.Employee) string {
	output := []string{}
	for i := 0; i < len(staff); i++ {
		str := fmt.Sprintf(
			"User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		output = append(output, str)
	}
	return strings.Join(output, "\n")
}
