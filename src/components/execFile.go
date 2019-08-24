package components

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

// ExecFile executes the file passed to it.
func ExecFile(name string, t int) error {
	proc := exec.Command(name)
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
	io.WriteString(stdin, "echoecho ")
	scanner.Scan()
	fmt.Println(scanner.Text())
	proc.Wait()
	fmt.Println(time.Since(startTime))
	return nil
}
