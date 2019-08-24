package components

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// ExecFile executes the file passed to it.
func ExecFile(name string, t int, min int, max int, n int) error {
	args := strings.Split(name, " ")
	var proc *exec.Cmd
	if len(args) == 2 {
		proc = exec.Command("cmd", "/C", name+" < ../temp.txt")
	} else if len(args) == 1 {
		proc = exec.Command(args[0], " < ../temp.txt")
	} else {
		return errors.New("Incorrect number of arguments given to process-fail")
	}
	stderr, err := proc.StderrPipe()
	if err != nil {
		return err
	}
	errScanner := bufio.NewScanner(bufio.NewReader(stderr))
	input, err := GenMultiRandom(n, min, max)
	if err != nil {
		return err
	}
	f, err := os.Create("../temp.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(strconv.Itoa(n) + "\n")
	for _, v := range input {
		f.WriteString(strconv.Itoa(v) + " ")
	}
	f.WriteString("\n")

	startTime := time.Now()
	if err := proc.Start(); err != nil {
		return err
	}
	for errScanner.Scan() {
		fmt.Println(errScanner.Text())
	}
	proc.Wait()

	fmt.Println(time.Since(startTime))
	return nil
}
