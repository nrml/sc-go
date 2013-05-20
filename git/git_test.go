package git

import (
	"fmt"
	"os"
	"testing"
)

func Test_0_Clone(t *testing.T) {
	hm := os.Getenv("HOME")
	path := hm + "/sc-test/github.com/nrml/sc-go/"
	os.RemoveAll(path)
	repo := Repo{"https://github.com", "nrml/sc-go", path}

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

	repo := Repo{"https://github.com", "nrml/sc-go", path}

	out, err := repo.Checkout([]string{"-f"})

	if err != nil {
		fmt.Println("OUTPUT: " + out)
		fmt.Println("FAIL: -----" + err.Error())
		t.Error(err)
		return
	}

	t.Log(out)
}
