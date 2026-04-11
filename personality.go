package main

import (
	"fmt"
	"math"
	"sort"
)

type Profile struct {
	Name        string
	Description string
	Traits      []string
}

func GenerateProfile(stats *RepoStats) *Profile {
	total := float64(stats.TotalCommits)
	docsRatio := float64(stats.CommitsByType["docs"]) / total
	featRatio := float64(stats.CommitsByType["feat"]) / total
	fixRatio := float64(stats.CommitsByType["fix"]) / total
	nightRatio := float64(stats.NightCommits) / total
	weekendRatio := float64(stats.WeekendCommits) / total
	convRatio := float64(stats.ConventionalCount) / total

	consistent := false
	if stats.TotalCommits > 5 {
		intervals := []float64{}
		for i := 1; i < len(stats.Timestamps); i++ {
			diff := stats.Timestamps[i].Sub(stats.Timestamps[i-1]).Hours()
			intervals = append(intervals, diff)
		}
		if len(intervals) > 0 {
			mean := 0.0
			for _, v := range intervals {
				mean += v
			}
			mean /= float64(len(intervals))
			variance := 0.0
			for _, v := range intervals {
				variance += math.Pow(v-mean, 2)
			}
			stdDev := math.Sqrt(variance / float64(len(intervals)))
			if stdDev < 24 {
				consistent = true
			}
		}
	}

	switch {
	case docsRatio > 0.25:
		return &Profile{
			Name:        "Documentation Prophet",
			Description: "You speak in READMEs and comment every function. The code is merely a vessel for your divine prose.",
			Traits:      []string{"Verbose", "Organized", "Example-driven"},
		}
	case nightRatio > 0.35:
		return &Profile{
			Name:        "The Midnight Committer",
			Description: "The sun is your enemy. You refactor entire modules under moonlight and drink coffee like water.",
			Traits:      []string{"Nocturnal", "Caffeinated", "Bold"},
		}
	case featRatio > 0.3 && fixRatio < 0.15:
		return &Profile{
			Name:        "Rogue Feature Implementer",
			Description: "Why fix bugs when you can ship features? Tests are for the weak, deadlines are for the strong.",
			Traits:      []string{"Creative", "Rebellious", "Fast-paced"},
		}
	case consistent && convRatio > 0.5:
		return &Profile{
			Name:        "Guardian of the Main Branch",
			Description: "Your commit messages are pristine, your rebases surgical. The main branch has never been safer.",
			Traits:      []string{"Disciplined", "Methodical", "Protective"},
		}
	case weekendRatio > 0.25:
		return &Profile{
			Name:        "The Weekend Warrior",
			Description: "Monday is just a suggestion. You ship on Saturdays and squash bugs on Sundays.",
			Traits:      []string{"Dedicated", "Unstoppable", "Sleep-deprived"},
		}
	default:
		return &Profile{
			Name:        "The Agile Devotee",
			Description: "You follow the rituals: daily standups, sprint planning, and just enough commits to look busy.",
			Traits:      []string{"Balanced", "Predictable", "Team player"},
		}
	}
}

func PrintProfile(p *Profile, stats *RepoStats) {
	fmt.Println("======================================")
	fmt.Printf("  Commitment Issues - Maintainer Profile\n")
	fmt.Println("======================================")
	fmt.Printf("  %s\n", p.Name)
	fmt.Println("--------------------------------------")
	fmt.Printf("  %s\n", p.Description)
	fmt.Println("--------------------------------------")
	fmt.Println("  Key Stats:")
	fmt.Printf("    Total commits: %d\n", stats.TotalCommits)
	fmt.Printf("    Top author:   %s\n", stats.TopAuthor)
	fmt.Printf("    Commit types: ")
	types := make([]string, 0, len(stats.CommitsByType))
	for t := range stats.CommitsByType {
		types = append(types, t)
	}
	sort.Strings(types)
	for i, t := range types {
		fmt.Printf("%s:%d", t, stats.CommitsByType[t])
		if i < len(types)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
	fmt.Printf("    Night commits:  %.1f%%\n", float64(stats.NightCommits)/float64(stats.TotalCommits)*100)
	fmt.Printf("    Weekend ratio:  %.1f%%\n", float64(stats.WeekendCommits)/float64(stats.TotalCommits)*100)
	fmt.Printf("    Conventional:   %.1f%%\n", float64(stats.ConventionalCount)/float64(stats.TotalCommits)*100)
	fmt.Printf("    Avg msg length: %.0f chars\n", stats.AvgMessageLength)
	fmt.Println("--------------------------------------")
	fmt.Println("  Personality Traits:")
	for _, trait := range p.Traits {
		fmt.Printf("    - %s\n", trait)
	}
	fmt.Println("======================================")
}