package repo

import (
	"fmt"

	"github.com/jdscdev/fs-developer-2505-porto-test/internal/utils"
)

type Commit struct {
	Timestamp int
	Username  string
	Files     int
	Additions int
	Deletions int
}

type Repository struct {
	RepoName      string
	Commits       []Commit
	FilesChanged  int
	LinesAdded    int
	LinesDeleted  int
	ActivityScore float64
}

const (
	TIMESTAMP_POS = iota
	USERNAME_POS
	REPO_NAME_POS
	FILES_POS
	ADDITIONS_POS
	DELETIONS_POS
)

const (
	WEIGHT_COMMITS       = 0.4
	WEIGHT_FILES_CHANGED = 0.2
	WEIGHT_LINES_ADDED   = 0.2
	WEIGHT_LINES_DELETED = 0.2
)

// FindOrCreateRepository finds or creates a repository struct
func FindOrCreateRepository(repositories []Repository, repoName string) (*Repository, bool) {
	for i, repository := range repositories {
		if repository.RepoName == repoName {
			return &repositories[i], true
		}
	}
	return &Repository{}, false
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

			var commit Commit = MapsNewCommitFromCommitArray(commitArray)

			if !repoFound {
				MapsNewRepoFromCommitArray(repository, commitArray, commit)
				repositories = append(repositories, *repository)
			} else {
				repository.Commits = append(repository.Commits, commit)
			}

			SetActivityScoreFromCommit(repository, commit)
		}
	}

	return repositories, error
}

// MapsNewRepoFromCommitArray creates a new repository mapped from a commit array
func MapsNewRepoFromCommitArray(repository *Repository, commitArray []string, commit Commit) {
	repository.RepoName = commitArray[REPO_NAME_POS]
	repository.Commits = []Commit{commit}
	repository.FilesChanged = commit.Files
	repository.LinesAdded = commit.Additions
	repository.LinesDeleted = commit.Deletions
}

// MapsNewCommitFromCommitArray maps a commit from a commit array
func MapsNewCommitFromCommitArray(commitArray []string) Commit {
	return Commit{
		Timestamp: utils.ConvertToInt(commitArray[TIMESTAMP_POS]),
		Username:  commitArray[USERNAME_POS],
		Files:     utils.ConvertToInt(commitArray[FILES_POS]),
		Additions: utils.ConvertToInt(commitArray[ADDITIONS_POS]),
		Deletions: utils.ConvertToInt(commitArray[DELETIONS_POS]),
	}
}

// SetActivityScoreFromCommit sets the activity score from a commit
func SetActivityScoreFromCommit(repository *Repository, commit Commit) {
	repository.FilesChanged += commit.Files
	repository.LinesAdded += commit.Additions
	repository.LinesDeleted += commit.Deletions
	repository.ActivityScore = float64(
		(float64(len(repository.Commits)) * WEIGHT_COMMITS) +
			(float64(repository.FilesChanged) * WEIGHT_FILES_CHANGED) +
			(float64(repository.LinesAdded) * WEIGHT_LINES_ADDED) +
			(float64(repository.LinesDeleted) * WEIGHT_LINES_DELETED))
}
