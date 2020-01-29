package main

import (
    "fmt"
    "gopkg.in/src-d/go-git.v4"
	"os/exec"
	"time"
	"os"
	"path/filepath"
	"bytes"
	"io"
)

var runCmd *exec.Cmd

func main() {
	args := os.Args
	if len(args) != 2 {
		prog := args[0]
		prog = filepath.Base(prog)
		fmt.Printf("Usage: %s <repo-path>\n", prog)
		return
	}
	path := args[1]

	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	valid1 := validEntryPoint(path)
	valid2 := validRepo(path)
	if !valid1 || !valid2 {
		fmt.Println("Failed to locate entry point.")
		fmt.Println("Make sure the repo-path provided is a valid git repo and there is a main.go file inside.")
		return
	}

	fmt.Println("Starting...")
	go start(path)

	for {
		pulled, err := pull(path)
		if err != nil {
			fmt.Println(err)
		}
		if pulled {
			fmt.Println("Update received. Restarting process...")
			if runCmd != nil {
				err = runCmd.Process.Kill()
				if err != nil {
					fmt.Println("Failed to kill process")
					return
				}
				runCmd = nil

			}
			go start(path)
		}
		time.Sleep(1 * time.Second)
	}

}

func pull(path string) (bool, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return false, err
	}

	w, err := r.Worktree()
	if err != nil {
		return false, err
	}

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})

	if err != nil && err == git.NoErrAlreadyUpToDate {
		return false, nil
	} else if err != nil && err != git.ErrUnstagedChanges {
		return false, err
	} else if err != nil {
		return false, err
	}

	ref, err := r.Head()
	if err != nil {
		return false, err
	}

	_, err = r.CommitObject(ref.Hash())
	if err != nil {
		return false, err
	}

	return true, nil
}

func start(path string) {
	buildCmd := exec.Command("cmd", "/c", "cd", path, "&&", "go", "build", "main.go")

	err := buildCmd.Run()
	if err != nil {
		fmt.Println("dafuq")
		fmt.Println(err)
		return
	}

	runCmd = exec.Command(path + "/main.exe")

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	runCmd.Stdout = mw
	runCmd.Stderr = mw


	err = runCmd.Start()
	if err != nil {
	    fmt.Println(err)
	}
	fmt.Print("\033[H\033[2J")
	fmt.Print(stdBuffer.String())
}

func validEntryPoint(path string) bool {
	info, err := os.Stat(path + "/main.go")
    if os.IsNotExist(err) {
        return false
    }
	if info.IsDir() {
		return false
	}
    return true
}

func validRepo(path string) bool {
	r, err := git.PlainOpen(path)
	if err != nil {
		return false
	}
	w, err := r.Worktree()
	if err != nil {
		return false
	}
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err == git.NoErrAlreadyUpToDate {
		return true
	} else if err != nil {
		return false
	}
	return true
}
