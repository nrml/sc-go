package git

import (
	"fmt"
	"log"
	"os/exec"
)

type PrivateRepo struct {
	Token string
	Repo  Repo
}

func (repo *PrivateRepo) Clone() (string, error) {
	err := repo.Repo.verifyPath(repo.Repo.Path)

	if err != nil {
		return "failed to create path", err
	}
	remote := fmt.Sprintf("%s://%s:x-oauth-basic@%s/%s.git", repo.Repo.Protocol, repo.Token, repo.Repo.Host, repo.Repo.Name)
	log.Printf("%s %s %s %s\n", "git", "clone", remote, repo.Repo.Path)

	cmd := exec.Command("git", "clone", remote, repo.Repo.Path)

	out, err := cmd.Output()

	return string(out), err
}
