# fs-developer-2505-porto-test

## Implementation Plan in Go
1. Read and Parse the CSV File:
    1.1 Use the encoding/csv package to load the data from "commits.csv" file.

2. Aggregate Data by Repository:
    2.1 Use an array of Repository structs to group data by repository name and accumulate metrics (commits, files changed, additions, deletions).

3. Compute Scores:
    3.1 Use the predefined formula to calculate the activity score for each repository.

4. Rank Repositories:
    4.1 Store the results and sort it by activity score in descending order.

5. Output Top 10 Repositories:
    5.1 Print the top 10 repositories and their activity scores.

## How to Run the Code
1. Ensure the commits.csv file is in the same directory as main.go.

2. Run the code using the following bash command:
`go run main.go`

3. The program will output the top 10 most active repositories along with their activity scores.

## Documentation
### Algorithm
1. Each repository is scored based on:
    1.1 Weights:
        - Commits (wc = 0.4)
        - Files Changed (wf = 0.2)
        - Additions (wa = 0.2)
        - Deletions (wd = 0.2)
    1.2 Activity Score formula for each repository:
        - Number of Commits * wc + Sum of Files * wf + Sum of Additions * wa + Sum of Deletions * wd

2. Activity Scores are calculated and stored in an array of Repository structs for each repository.

3. Repositories are ranked by their scores, and the top 10 are printed.

4. The program will print the top 10 repositories in descending order of their scores:
