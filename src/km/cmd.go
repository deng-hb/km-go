package km

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func KillProcess(port string) {

	cmd := exec.Command(`netstat`, `-ano`)
	output, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		cmd.Run()
	}()
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, `LISTENING`) && strings.Contains(line, port) {
			parts := strings.Fields(line)
			pid := parts[len(parts)-1]
			fmt.Printf("kill pid %s\n", pid)

			killCmd := exec.Command(`taskkill`, `/F`, `/PID`, pid)
			killCmd.Run()
		}
	}
}
