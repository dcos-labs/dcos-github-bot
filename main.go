package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	// TODO(joerg84): get via env variables
	org := "joerg84"
	repository := "universe"
	communityPrLabel := "community-pr1"

	client := github.NewClient(nil)

	prs, _, err := client.PullRequests.List(context.Background(), org, repository, nil)
	if err != nil {
		fmt.Printf("PullRequests.List returned error: %v", err)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total number of PRs: %d\n", len(prs))

	for _, pr := range prs {
		fmt.Printf("Handling PR number: %d\n", pr.GetNumber())

		// Check if PR needs Label & comment
		// TODO(joerg84): check whether PR id from member in mesosphere org
		hasCommunityPRLabel := false
		for _, label := range pr.Labels {
			fmt.Println("%v", label)
			if *label.Name == communityPrLabel {
				hasCommunityPRLabel = true
				break
			}
		}

		if !hasCommunityPRLabel {
			fmt.Printf("PR needs label.")

			// Add comment & label
		}

	}

	// organizations, err := FetchOrganizations(username)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }

}
