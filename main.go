package main

import (
	"os/exec"
	"log"
	"bufio"
	"strings"
	"fmt"
	"os"
	"strconv"
)


type device struct {
	Name string
	Serial string
}

func main() {

	cmd := exec.Command("adb", "devices")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	var deviceList []device
	r  := bufio.NewReader(stdout)
	for buf, _, err := r.ReadLine(); err == nil; buf, _, err = r.ReadLine() {
		l := string(buf)
		if strings.Index(l, "List of") == 0 {
			continue
		}
		if strings.TrimSpace(l) == "" {
			continue
		}
		sep := strings.Split(l, "\t")
		
		deviceList = append(deviceList, device{ sep[1], sep[0] })
	}

	serial := askDevice(deviceList)
	fmt.Printf(serial)

}

func askDevice(deviceList[] device) (serial string) {

	showDevices(deviceList)

	for {
		fmt.Fprintf(os.Stderr, "which device?: ")
		in := bufio.NewReader(os.Stdin)
		buf, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		num, err := strconv.Atoi(string(buf))
		if err == nil && num >= 1 && num <= len(deviceList) {
			serial = deviceList[num - 1].Serial
			return
		}
		fmt.Fprintf(os.Stderr, "error.\n")
	}

	return ""
}


func showDevices(deviceList []device) {

	for i, device := range(deviceList) {
		fmt.Fprintf(os.Stderr, "%d) %s %s\n", i + 1, device.Name, device.Serial)
	}
}
