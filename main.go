package main

import (
	"fmt"

	repo "github.com/jdscdev/fs-developer-2505-porto-test/internal/models"
)

func main() {
	repositories, _ := repo.GetRepositoriesFromCSVFile("commits.csv")

	fmt.Println(repositories)
}
