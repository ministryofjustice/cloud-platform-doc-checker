package main

import (
	"flag"
	"log"
	"os"

	"github.com/ministryofjustice/cloud-platform-doc-checker/internal/check"

	ghaction "github.com/sethvargo/go-githubactions"
)

// Require a GitHub team name at runtime.
var (
	team = flag.String("team", "", "team and orgOwner are the GitHub team and organisation that we're using to validate the user.")
)

// Everything else has default values that most teams will use.
var (
	prOwner  = flag.String("owner", os.Getenv("PR_OWNER"), "contains the value of an environment variable that is set in the GH action container")
	orgOwner = flag.String("org", "ministryofjustice", "who owns the repository")
	fileName = flag.String("file", "changes", "the file created by a GitHub action, it contains the output of a git diff")
	token    = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "Personal access token from GitHub.")
)

func main() {
	flag.Parse()

	// The user must either specify or set the PR_OWNER and token environment variable.
	if *team == "" || *token == "" || *prOwner == "" {
		log.Fatalln("you must have the GITHUB_OAUTH_TOKEN, PR_OWNER and team name defined.")
	}

	// Parse the pull request body file and return false if the PR has more than a review change.
	prRelevant, err := check.ParsePR(*fileName)
	if err != nil {
		log.Println("Unable to parse the PR", err)
	}

	// If the PR is relevant and the user is allowed to be auto approved, we'll return a success.
	userAllowed, err := check.GitHubTeam(*team, *orgOwner, *token, *prOwner)
	if err != nil {
		log.Println("Unable to check if the user is valid.", err)
	}

	// We don't want a hard fail so we set the output to false and log.
	if prRelevant && userAllowed {
		log.Println("Success: The changes in this PR are only review dates and the user is valid.")
		ghaction.SetOutput("review_pr", "true")
	} else {
		log.Println("Fail: Either the PR contains more than review date changes or the user isn't a member of the webops team.")
		ghaction.SetOutput("review_pr", "false")
	}
}