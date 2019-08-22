package components

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/agzuniverse/go-testcases/src/utils"
)

// ExecFile executes the file passed to it.
func ExecFile(name string, t int) error {
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
	startTime := time.Now()
	if err := proc.Start(); err != nil {
		utils.HandleErr(err)
	}
	reader := bufio.NewReader(stdout)
	scanner := bufio.NewScanner(reader)
	io.WriteString(stdin, "echoecho ")
	scanner.Scan()
	fmt.Println(scanner.Text())
	proc.Wait()
	fmt.Println(time.Since(startTime))
	return nil
}
