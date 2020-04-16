package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	allEntries := make([]string, len(rhyme))
	for i := 0; i < len(allEntries); i++ {
		if i < len(rhyme) - 1 {
			allEntries[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		} else {
			allEntries[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		}
	}
	return allEntries
}
