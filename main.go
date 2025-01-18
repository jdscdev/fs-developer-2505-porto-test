package main

import (
	"fmt"

	repo "github.com/jdscdev/fs-developer-2505-porto-test/internal/models"
)

func main() {
	repositories, err := repo.GetRepositoriesFromCSVFile("commits.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// repo.GetRepositoriesFromCSVFile("commits.csv")

	// fmt.Println(repositories)

	for _, repository := range repositories {
		fmt.Println("Repo:", repository.RepoName, "Commits:", len(repository.Commits), "ActivityStore:", repository.ActivityStore)
	}
}
