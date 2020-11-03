package hooks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Eldius/go-version-manager/config"
	"github.com/google/uuid"
)

func AddHook(hook string) {
	fmt.Println("arg:", hook)
	if i, err := os.Stat(hook); err != nil {
		addCommand(hook)
	} else {
		if i.IsDir() {
			fmt.Println("Arg isn't a file")
			os.Exit(1)
		}
		if !isExecutable(i.Mode()) {
			fmt.Println("Script file isn't an executable file")
			os.Exit(1)
		}
		addScriptFile(hook)
	}
}

func ListHooks() []string {
	files, err := ioutil.ReadDir(config.GetHooksDir())
	var result []string
	if err != nil {
		fmt.Println("Failed to list hooks")
		log.Fatal(err)
	}

	for _, f := range files {
		result = append(result, f.Name())
	}

	return result
}

func ExecuteHook(hook string) {
	fmt.Println("Executing hook", hook)
	cmd := exec.Command(filepath.Join(config.GetHooksDir(), hook))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to execute hook script:\n%s\n", err.Error())
	}
	fmt.Println("---")
}

func addCommand(command string) {
	fmt.Println("Adding a command to the hooks list")
	panic("Not implemented yet.")
}

func addScriptFile(script string) {
	fmt.Println("Adding a script to the hooks list")
	input, err := ioutil.ReadFile(script)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = os.MkdirAll(config.GetHooksDir(), os.ModePerm)

	destinationFile := filepath.Join(config.GetHooksDir(), uuid.New().String())
	err = ioutil.WriteFile(destinationFile, input, 0700)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}

func isExecutable(mode os.FileMode) bool {
	return mode&0100 != 0 ||
		mode&0010 != 0 ||
		mode&0001 != 0
}
