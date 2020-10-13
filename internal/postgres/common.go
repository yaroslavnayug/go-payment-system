package postgres

import (
	"fmt"
	"strings"
)

// creates sequence ($1, $2, $3, ...) base on columns length
func getSubstitutionVerbsForColumns(columns []string) string {
	var verbs []string
	for i := 1; i <= len(columns); i++ {
		verbs = append(verbs, fmt.Sprintf("$%d", i))
	}
	return strings.Join(verbs, ", ")
}
