package bitbucket

import "github.com/daeMOn63/bitclient"

func LoadAllProjects(client *bitclient.BitClient) (projects []bitclient.Project, err error) {
	var page = 1

	for {
		next, err := client.GetProjects(
			bitclient.PagedRequest{
				Limit: 1000,
				Start: uint(len(projects)),
			},
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, next.Values...)

		if next.IsLastPage {
			break
		}

		page++
	}

	return projects, nil
}

func LoadAllRepositoriesForProject(client *bitclient.BitClient, projectKey string) (repos []bitclient.Repository, err error) {
	var page = 1

	for {
		next, err := client.GetRepositories(
			projectKey,
			bitclient.PagedRequest{
				Limit: 1000,
				Start: uint(len(repos)),
			},
		)
		if err != nil {
			return nil, err
		}

		repos = append(repos, next.Values...)

		if next.IsLastPage {
			break
		}

		page++
	}

	return repos, nil
}
