package main

import (
	"fmt"
	"sort"

	repo "github.com/jdscdev/fs-developer-2505-porto-test/internal/models"
)

func main() {
	repositories, err := repo.GetRepositoriesFromCSVFile("commits.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Sort activity score for each repository
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].ActivityScore > repositories[j].ActivityScore
	})

	fmt.Println(" ********** Top 10 repositories by activity score ********** \n",
		"-------------------------------------------------------------------------------- \n",
		" Repo   | #Commits | Files Changed | Lines Added | Lines Deleted | ActivityScore \n",
		"--------------------------------------------------------------------------------")

	for i := 0; i < 10; i++ {
		fmt.Printf("%s | %d       | %d          | %d      | %d       | %.2f\n",
			repositories[i].RepoName,
			len(repositories[i].Commits),
			repositories[i].FilesChanged,
			repositories[i].LinesAdded,
			repositories[i].LinesDeleted,
			repositories[i].ActivityScore)
	}
}
