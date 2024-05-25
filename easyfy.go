package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

var dirPath string = "./shortcuts"

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
	if checkDirCreated() {
		if len(os.Args) < 2 {
			fmt.Print("Enter Shortcut Name: ")
			var shortcut_name string
			fmt.Scanln(&shortcut_name)

			fmt.Print("Enter Commands to execute: ")
			reader := bufio.NewReader(os.Stdin)
			shortcut_commands, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			shortcut_commands = strings.TrimSpace(shortcut_commands)
			err = ioutil.WriteFile("shortcuts/"+shortcut_name+".bat", []byte(shortcut_commands), 0755)
			if err != nil {
				panic(err)
			}
		}else{
			shortcutCommands := strings.Join(os.Args[1:], " ")
			fmt.Println("Commands to execute:", shortcutCommands)
			fmt.Print("Enter Shortcut Name: ")
			var shortcut_name string
			fmt.Scanln(&shortcut_name)
			err := ioutil.WriteFile("shortcuts/"+shortcut_name+".bat", []byte(shortcutCommands), 0755)
			if err != nil {
				panic(err)
			}
		}
	}else{
		fmt.Println("sorry something went wrong")
	}
	
}
