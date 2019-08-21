package components

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/agzuniverse/go-testcases/src/utils"
)

// ExecFile executes the file passed to it.
func ExecFile(name string, time int) error {
	proc := exec.Command(name)
	stdin, err := proc.StdinPipe()
	if err != nil {
		utils.HandleErr(err)
	}
	stdout, err2 := proc.StdoutPipe()
	if err != nil {
		utils.HandleErr(err2)
	}
	defer stdout.Close()
	if err := proc.Start(); err != nil {
		utils.HandleErr(err)
	}
	io.WriteString(stdin, "echo echo")
	buf := make([]byte, 100)
	stdout.Read(buf)
	fmt.Println(buf)
	proc.Wait()
	return nil
}
