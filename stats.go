package main

import (
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type RepoStats struct {
	TotalCommits      int
	CommitsByType     map[string]int
	NightCommits      int
	WeekendCommits    int
	ConventionalCount int
	AvgMessageLength  float64
	Timestamps        []time.Time
	TopAuthor         string
}

func CollectStats(repoPath string) (*RepoStats, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	head, err := repo.Head()
	if err != nil {
		return nil, err
	}

	iter, err := repo.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	stats := &RepoStats{
		CommitsByType: make(map[string]int),
		Timestamps:    []time.Time{},
	}
	authorCount := make(map[string]int)
	totalMsgLen := 0

	err = iter.ForEach(func(c *object.Commit) error {
		stats.TotalCommits++
		ts := c.Author.When
		stats.Timestamps = append(stats.Timestamps, ts)

		authorCount[c.Author.Email]++
		msg := c.Message
		totalMsgLen += len(msg)

		commitType, isConv := classifyMessage(msg)
		stats.CommitsByType[commitType]++
		if isConv {
			stats.ConventionalCount++
		}

		hour := ts.Hour()
		if hour >= 22 || hour < 5 {
			stats.NightCommits++
		}

		wd := ts.Weekday()
		if wd == time.Saturday || wd == time.Sunday {
			stats.WeekendCommits++
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if stats.TotalCommits == 0 {
		return nil, fmt.Errorf("no commits found in repository")
	}

	stats.AvgMessageLength = float64(totalMsgLen) / float64(stats.TotalCommits)

	maxAuthor := ""
	maxCount := 0
	for email, cnt := range authorCount {
		if cnt > maxCount {
			maxCount = cnt
			maxAuthor = email
		}
	}
	stats.TopAuthor = maxAuthor

	sort.Slice(stats.Timestamps, func(i, j int) bool {
		return stats.Timestamps[i].Before(stats.Timestamps[j])
	})

	return stats, nil
}

func classifyMessage(msg string) (string, bool) {
	convRe := regexp.MustCompile(`^(?i)(feat|fix|docs|style|refactor|test|chore)(\([^\)]+\))?!?:`)
	matches := convRe.FindStringSubmatch(msg)
	if len(matches) > 1 {
		return strings.ToLower(matches[1]), true
	}

	lower := strings.ToLower(msg)
	switch {
	case strings.Contains(lower, "fix"):
		return "fix", false
	case strings.Contains(lower, "feat"):
		return "feat", false
	case strings.Contains(lower, "doc"):
		return "docs", false
	case strings.Contains(lower, "style"):
		return "style", false
	case strings.Contains(lower, "refactor"):
		return "refactor", false
	case strings.Contains(lower, "test"):
		return "test", false
	case strings.Contains(lower, "chore"):
		return "chore", false
	default:
		return "other", false
	}
}