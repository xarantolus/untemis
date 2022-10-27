package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/xarantolus/untemis/bitbucket"
	"github.com/xarantolus/untemis/config"

	"github.com/daeMOn63/bitclient"
)

func main() {
	var (
		flagConfigFile = flag.String("cfg", "config.yml", "Path to config file")
		destination    = flag.String("dst", "repos", "Destination folder")
	)
	flag.Parse()

	settings, err := config.Parse(*flagConfigFile)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	client := bitclient.NewBitClient("https://"+settings.Server, settings.Username, settings.Password)

	projects, err := bitbucket.LoadAllProjects(client)
	if err != nil {
		log.Fatalf("Error loading projects: %v", err)
	}

	var errors []string

	errLog := func(err string) {
		log.Println(err)
		errors = append(errors, err)
	}

	for _, project := range projects {
		repos, err := bitbucket.LoadAllRepositoriesForProject(client, project.Key)
		if err != nil {
			errLog(fmt.Sprintf("Error loading repositories for project %s: %v", project.Key, err))
			continue
		}

		log.Printf("Project %s has %d repositories", project.Key, len(repos))

		for _, repo := range repos {
			log.Printf("Cloning %s/%s", project.Key, repo.Slug)

			dest := path.Join(*destination, repo.Slug)

			cmd := exec.Command("git", "clone", "https://"+settings.Server+"/scm/"+project.Key+"/"+repo.Slug+".git", dest)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				errLog(fmt.Sprintf("Error cloning repository %s/%s: %v", project.Key, repo.Slug, err))
			}
		}
	}

	if len(errors) > 0 {
		log.Printf("There were %d errors while processing:", len(errors))
		for _, err := range errors {
			log.Println("    " + err)
		}
	} else {
		log.Println("Done, no errors!")
	}
}
