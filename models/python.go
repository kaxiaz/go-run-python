package model

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type PythonModel struct{}

// RunProgram :
func (pm PythonModel) RunProgram(code string, in string) (string, int64, error) {
	cmd := exec.Command("python3", "-c", code)
	cmdIn := exec.Command("echo", in)
	cmd.Stdin, _ = cmdIn.StdoutPipe()
	cmd.Stdout = os.Stdout

	var buf bytes.Buffer
	cmd.Stdout = &buf

	// start time counter
	start := time.Now().UnixNano()
	// run program
	cmd.Start()
	// insert input
	err := cmdIn.Run()
	if err != nil {
		log.Println(err)
		return "IE", 2000, errors.New("input error")
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(3 * time.Second)

	select {
	case <-timeout:
		cmd.Process.Kill()
		return "TLE", 2000, errors.New("command timed out")
	case err := <-done:
		if err != nil {
			log.Println(err)
			return "RE", 2000, errors.New("runtime error")
		}
		return buf.String(), (time.Now().UnixNano() - start) / int64(time.Millisecond), nil
	}
}

// EvaluateProgram :
func (pm PythonModel) EvaluateProgram(code string, cases []string) (string, error) {
	if len(cases)%2 != 0 {
		return "System Error", errors.New("test cases are invalid")
	}

	for i := 0; i < len(cases); i += 2 {
		// run program
		out, _, err := pm.RunProgram(code, cases[i])
		// check error output
		if err != nil {
			return out, err
		}
		// trim \n and white space
		runOutput := strings.ReplaceAll(string(out), `\n`, "\n")
		caseOutput := strings.ReplaceAll(string(cases[i+1]), `\n`, "\n")
		runOutput = strings.TrimSpace(string(runOutput))
		caseOutput = strings.TrimSpace(string(caseOutput))
		// compare output
		if runOutput != caseOutput {
			msg := fmt.Sprintf("Wrong output on test case #%d\n", ((i + 1) / 2))
			log.Print(msg)
			return "WA", errors.New("WA")
		}
	}

	return "", nil
}

// GetTestCases :
func (pm PythonModel) GetTestCases(problemID uint64) ([]string, error) {
	var cases []string

	// db := dbhandler.OpenConnection()
	// defer dbhandler.CloseConnection(db)

	// stmt, err := db.Prepare(`...`)
	// if err != nil {
	// 	return cases, err
	// }
	// defer stmt.Close()

	// rows, err := stmt.Query(problemID)
	// if err != nil {
	// 	return cases, err
	// }

	// for rows.Next() {
	// 	var in, out string
	// 	if err := rows.Scan(&in, &out); err != nil {
	// 		return cases, err
	// 	}
	// 	cases = append(cases, in)
	// 	cases = append(cases, out)
	// }

	// just test data
	cases = append(cases, "hello") // input
	cases = append(cases, "hello") // output
	return cases, nil
}
