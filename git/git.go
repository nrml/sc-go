package git

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type Repo struct {
	Host string
	Name string
	Path string
}
type PrivateRepo struct {
	User  string
	Token string
	_     Repo
}

func (repo *Repo) Clone() (string, error) {
	err := repo.verifyPath(repo.Path)

	if err != nil {
		return "failed to create path", err
	}
	log.Printf("%s %s %s %s\n", "git", "clone", repo.Host+"/"+repo.Name+".git", repo.Path)
	cmd := exec.Command("git", "clone", repo.Host+"/"+repo.Name+".git", repo.Path)

	out, err := cmd.Output()

	return string(out), err
}
func (repo *Repo) Pull() {

}
func (repo *Repo) Checkout(flags []string) (string, error) {
	err := repo.verifyPath(repo.Path)

	if err != nil {
		return "failed to create path", err
	}

	log.Printf("%s %s %s\n", "git", "checkout", strings.Join(flags, " "))

	err = os.Chdir(repo.Path)

	if err != nil {
		return "cannnot change dir to path", err
	}

	cmd := exec.Command("git", "checkout", strings.Join(flags, " "))

	out, err := cmd.Output()

	return string(out), err
}
func (repo *Repo) Push() {

}
func (repo *Repo) verifyPath(path string) error {
	err := os.MkdirAll(path, os.ModeDir|os.ModePerm|os.ModeAppend)
	return err
}
