package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

    cmd := exec.Command("sp", "current")

    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()

    if err != nil {
        fmt.Errorf("error exec command: %s\n", err)
    }

    if strings.HasPrefix(out.String(), "Error") {
        fmt.Errorf("spotify not running: %s\n", out.String())
        os.Exit(0)
    }

    lines := strings.Split(out.String(), "\n")
    var title string
    var artist string

    for _, line := range(lines) {
        if strings.HasPrefix(line, "Title") {
            title = strings.TrimSpace(strings.TrimPrefix(line, "Title"))
        } else if strings.HasPrefix(line, "Artist") {
            artist = strings.TrimSpace(strings.TrimPrefix(line, "Artist"))
        }
    }
    fmt.Printf("%s - %s", artist, title)

}
