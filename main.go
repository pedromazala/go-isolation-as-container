package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)
// Docker synas and our syntax
// docker run <container> command args
// go run main.go run command args

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

func run() {
	// Create a "self" process passing "child" argument and our command + parameters
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// Vinculate standard input, output and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// isolate UTS, PID and File system
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	// execute our command
	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as pid %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}