package components

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

// ExecFile executes the file passed to it.
func ExecFile(name string, t int) error {
	args := strings.Split(name, " ")
	var proc *exec.Cmd
	if len(args) == 2 {
		proc = exec.Command(args[0], args[1])
	} else if len(args) == 1 {
		proc = exec.Command(args[0])
	} else {
		return errors.New("Incorrect number of arguments given to process-fail")
	}
	stdin, err := proc.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err2 := proc.StdoutPipe()
	if err != nil {
		return err2
	}
	startTime := time.Now()
	if err3 := proc.Start(); err3 != nil {
		return err3
	}
	reader := bufio.NewReader(stdout)
	scanner := bufio.NewScanner(reader)
	io.WriteString(stdin, "echoecho\n")
	scanner.Scan()
	fmt.Println(scanner.Text())
	proc.Wait()

	fmt.Println(time.Since(startTime))
	return nil
}
