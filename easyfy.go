package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"
)

var dirPath string = "shortcuts"

func getYesOrNo(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt + " (yes/no): ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "yes" || response == "y" {
			return true
		} else if response == "no" || response == "n" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
		}
	}
}

func checkDirCreated() bool{
	fmt.Print("Checking if dircetory is created :) \n\n")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Println("Directory Doesnot exist , creating Directory :) \n\n")
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return false
		}else{
			fmt.Println("Directory created sucessfully :) \n\n")
			return true
		}
	}else{
		fmt.Println("Directory is alredy created :) \n\n")
		return true
	}
	
}

func main() {
	exePath, err := os.Executable()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	exeDir := filepath.Dir(exePath) + string(os.PathSeparator)
	dirPath = exeDir + dirPath

	if checkDirCreated() {
		var shortcut_name string
		var shortcut_commands string
		cwd,_:=os.Getwd()
		directory := filepath.VolumeName(cwd)
		if len(os.Args) < 2 {
			fmt.Print("Enter Shortcut Name: ")
			fmt.Scanln(&shortcut_name)
			fmt.Print("Enter Commands to execute: ")
			reader := bufio.NewReader(os.Stdin)
			shortcutcommands, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			shortcut_commands = strings.TrimSpace(shortcutcommands)
		}else{
			shortcut_commands = strings.Join(os.Args[1:], " ")
			fmt.Println("Commands to execute:", shortcut_commands)
			fmt.Print("Enter Shortcut Name: ")
			fmt.Scanln(&shortcut_name)
		}

		if(shortcut_name != "" && shortcut_commands != ""){
			if getYesOrNo("Is the command location specific ?") {
				fmt.Println(cwd)
				fmt.Println(directory)
				shortcut_commands = directory + " && cd " + cwd + " && " + shortcut_commands
				fmt.Println("\nfinal command : " ,shortcut_commands)
			}

			err := ioutil.WriteFile(dirPath+"/"+shortcut_name+".bat", []byte(shortcut_commands), 0644)
			if err != nil {
				panic(err)
			}
		}
	}else{
		fmt.Println("sorry something went wrong")
	}
	
}
