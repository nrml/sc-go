package git

import (
	"fmt"
	"os"
	"testing"
)

func Test_Clone(t *testing.T) {
	hm := os.Getenv("HOME")
	repo := Repo{"https://github.com", "nrml/sc-go", hm + "/sc-test/github.com/"}

	out, err := repo.Clone()

	if err != nil {
		fmt.Println("OUTPUT: " + out)
		fmt.Println("FAIL: -----" + err.Error())
		t.Error(err)
		return
	}

	t.Log(out)
}
