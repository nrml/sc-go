package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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

//const gitPath = "/usr/local/git/bin/git"

func (repo *Repo) Clone() (string, error) {
	err := repo.verifyPath(repo.Path)

	if err != nil {
		return "failed to create path", err
	}

	fmt.Println("trying to clone")

	fmt.Println("cloning with : git clone " + repo.Host + "/" + repo.Name + ".git " + repo.Path)

	cmd := exec.Command("git", "clone", repo.Host+"/"+repo.Name+".git", repo.Path)

	out, err := cmd.Output()

	if err != nil {

		fmt.Println("error exec git:  " + err.Error())
	}

	return string(out), err
}
func (repo *Repo) Pull() {

}
func (repo *Repo) Checkout() {

}
func (repo *Repo) Push() {

}
func (repo *Repo) verifyPath(path string) error {
	err := os.MkdirAll(path, os.ModeDir|os.ModePerm|os.ModeAppend)
	if err != nil {
		return errors.New("failed to create paths: " + path)
	}
	return err

	//return err
}
