package repo

import (
	"fmt"

	"github.com/jdscdev/fs-developer-2505-porto-test/internal/utils"
)

type Commit struct {
	Timestamp int    `json:"timestamp"`
	Username  string `json:"username"`
	Files     string `json:"files"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
}

type Repository struct {
	RepoName        string   `json:"repository"`
	Commits         []Commit `json:"commits"`
	StargazersCount int      `json:"stargazers_count"`
	ForksCount      int      `json:"forks_count"`
	OpenIssuesCount int      `json:"open_issues_count"`
}

const TIMESTAMP_POS = 0
const USERNAME_POS = 1
const REPO_NAME_POS = 2
const FILES_POS = 3
const ADDITIONS_POS = 4
const DELETIONS_POS = 5

// FindOrCreateRepository finds or creates a repository struct
func FindOrCreateRepository(repositories []Repository, repoName string) (Repository, bool) {
	for _, repository := range repositories {
		if repository.RepoName == repoName {
			return repository, true
		}
	}
	return Repository{}, false
}

// GetRepositoriesFromCSVFile reads a CSV file and returns an array of repositories
func GetRepositoriesFromCSVFile(csvFile string) ([]Repository, error) {
	commits, error := utils.ReadCSVFile(csvFile)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}

	var repositories []Repository

	for ind, commitArray := range commits {
		if ind > 0 {
			repository, repoFound := FindOrCreateRepository(repositories, commitArray[REPO_NAME_POS])

			if !repoFound {
				MapsNewRepoFromCommitArray(&repository, commitArray)
			} else {
				repository.Commits = append(repository.Commits, NewCommitMappedFromCommitArray(commitArray))
			}

			repositories = append(repositories, repository)
		}
	}

	return repositories, error
}

// MapsNewRepoFromCommitArray creates a new repository mapped from a commit array
func MapsNewRepoFromCommitArray(repository *Repository, commitArray []string) {
	var commit Commit = NewCommitMappedFromCommitArray(commitArray)

	repository.RepoName = commitArray[REPO_NAME_POS]
	repository.Commits = []Commit{commit}
	repository.StargazersCount = 0
	repository.ForksCount = 0
	repository.OpenIssuesCount = 0
}

// MapCommitFromCommitArray maps a commit from a commit array
func NewCommitMappedFromCommitArray(commitArray []string) Commit {
	return Commit{
		Timestamp: utils.ConvertToInt(commitArray[TIMESTAMP_POS]),
		Username:  commitArray[USERNAME_POS],
		Files:     commitArray[FILES_POS],
		Additions: utils.ConvertToInt(commitArray[ADDITIONS_POS]),
		Deletions: utils.ConvertToInt(commitArray[DELETIONS_POS]),
	}
}
