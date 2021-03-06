package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_0_Clone(t *testing.T) {
	hm := os.Getenv("HOME")
	path := hm + "/sc-test/github.com/nrml/sc-go/"
	os.RemoveAll(path)
	repo := Repo{"https", "github.com", "nrml/sc-go", path}

	out, err := repo.Clone()

	if err != nil {
		fmt.Println("OUTPUT: " + out)
		fmt.Println("FAIL: -----" + err.Error())
		t.Error(err)
		return
	}

	t.Log(out)
}

func Test_1_CheckoutForce(t *testing.T) {
	hm := os.Getenv("HOME")
	path := hm + "/sc-test/github.com/nrml/sc-go/"

	repo := Repo{"https", "github.com", "nrml/sc-go", path}

	out, err := repo.Checkout([]string{"-f"})

	if err != nil {
		fmt.Println("OUTPUT: " + out)
		fmt.Println("FAIL: -----" + err.Error())
		t.Error(err)
		return
	}

	t.Log(out)
}

func Test_2_ClonePrivate(t *testing.T) {
	hm := os.Getenv("HOME")
	path := hm + "/sc-test/github.com/brentmn/notation.io"
	os.RemoveAll(path)

	b, err := ioutil.ReadFile(hm + "/tokens/github.com.tkn")

	token := string(b)

	repo := Repo{"https", "github.com", "brentmn/notation.io", path}
	prvt := PrivateRepo{token, repo}

	out, err := prvt.Clone()

	if err != nil {
		fmt.Println("OUTPUT: " + out)
		fmt.Println("FAIL: -----" + err.Error())
		t.Error(err)
		return
	}

	t.Log(out)
}
