package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the CLI! Type 'exit' to quit.")
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}
		executeCommand(input)
	}
}

func executeCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}
	command := args[0]
	switch strings.ToLower(command) {
	case "cls":
		clearScreen()
	case "dir":
		listDirectory()
	case "ls":
		listDirectory()
	case "cd":
		changeDirectory(args[1:])
	case "pwd":
		printWorkingDirectory()
	case "mkdir":
		makeDirectory(args[1:])
	case "rmdir":
		removeDirectory(args[1:])
	case "help":
		displayHelp()
	case "rm":
		removeFileOrDirectory(args[1:])
	case "cp":
		copyFileOrDirectory(args[1:])
	case "mv":
		moveFileOrDirectory(args[1:])
	case "touch":
		createEmptyFile(args[1:])
	case "cat":
		displayFileContents(args[1:])
	case "more", "less":
		displayFileContentsByPage(args[1:])
	case "head":
		displayFirstFewLines(args[1:])
	case "tail":
		displayLastFewLines(args[1:])
	case "find":
		findFilesOrDirectories(args[1:])
	case "grep":
		searchTextWithinFiles(args[1:])
	case "chmod":
		changeFilePermissions(args[1:])
	case "chown":
		changeFileOwner(args[1:])
	case "chgrp":
		changeFileGroup(args[1:])
	case "df":
		displayDiskSpaceUsage(args[1:])
	case "du":
		displayDiskUsage(args[1:])
	// Add more cases for other commands
	default:
		fmt.Println("Unknown command:", command)
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if isWindows() {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	runCommand(cmd)
}

func listDirectory() {
	var cmd *exec.Cmd
	if isWindows() {
		cmd = exec.Command("cmd", "/c", "dir")
	} else {
		cmd = exec.Command("ls", "-l")
	}
	runCommand(cmd)
}

func changeDirectory(args []string) {
	if len(args) == 0 {
		fmt.Println("cd: missing argument")
		return
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println("cd:", err)
	}
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd:", err)
		return
	}
	fmt.Println(dir)
}

func makeDirectory(args []string) {
	if len(args) == 0 {
		fmt.Println("mkdir: missing argument")
		return
	}
	err := os.Mkdir(args[0], 0755)
	if err != nil {
		fmt.Println("mkdir:", err)
	}
}

func removeDirectory(args []string) {
	if len(args) == 0 {
		fmt.Println("rmdir: missing argument")
		return
	}
	err := os.Remove(args[0])
	if err != nil {
		fmt.Println("rmdir:", err)
	}
}

func removeFileOrDirectory(args []string) {
	if len(args) == 0 {
		fmt.Println("rm: missing argument")
		return
	}
	err := os.RemoveAll(args[0])
	if err != nil {
		fmt.Println("rm:", err)
	}
}

func copyFileOrDirectory(args []string) {
	if len(args) < 2 {
		fmt.Println("cp: missing argument")
		return
	}
	cmd := exec.Command("cp", "-r", args[0], args[1])
	runCommand(cmd)
}

func moveFileOrDirectory(args []string) {
	if len(args) < 2 {
		fmt.Println("mv: missing argument")
		return
	}
	err := os.Rename(args[0], args[1])
	if err != nil {
		fmt.Println("mv:", err)
	}
}

func createEmptyFile(args []string) {
	if len(args) == 0 {
		fmt.Println("touch: missing argument")
		return
	}
	file, err := os.Create(args[0])
	if err != nil {
		fmt.Println("touch:", err)
		return
	}
	file.Close()
}

func displayFileContents(args []string) {
	if len(args) == 0 {
		fmt.Println("cat: missing argument")
		return
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("cat:", err)
		return
	}
	fmt.Println(string(data))
}

func displayFileContentsByPage(args []string) {
	if len(args) == 0 {
		fmt.Println("more/less: missing argument")
		return
	}
	cmd := exec.Command("less", args[0])
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("more/less:", err)
	}
}

func displayFirstFewLines(args []string) {
	if len(args) == 0 {
		fmt.Println("head: missing argument")
		return
	}
	cmd := exec.Command("head", args[0])
	runCommand(cmd)
}

func displayLastFewLines(args []string) {
	if len(args) == 0 {
		fmt.Println("tail: missing argument")
		return
	}
	cmd := exec.Command("tail", args[0])
	runCommand(cmd)
}

func findFilesOrDirectories(args []string) {
	if len(args) == 0 {
		fmt.Println("find: missing argument")
		return
	}
	cmd := exec.Command("find", args...)
	runCommand(cmd)
}

func searchTextWithinFiles(args []string) {
	if len(args) < 2 {
		fmt.Println("grep: missing argument")
		return
	}
	cmd := exec.Command("grep", args...)
	runCommand(cmd)
}

func changeFilePermissions(args []string) {
	if len(args) < 2 {
		fmt.Println("chmod: missing argument")
		return
	}
	cmd := exec.Command("chmod", args...)
	runCommand(cmd)
}

func changeFileOwner(args []string) {
	if len(args) < 2 {
		fmt.Println("chown: missing argument")
		return
	}
	cmd := exec.Command("chown", args...)
	runCommand(cmd)
}

func changeFileGroup(args []string) {
	if len(args) < 2 {
		fmt.Println("chgrp: missing argument")
		return
	}
	cmd := exec.Command("chgrp", args...)
	runCommand(cmd)
}

func displayDiskSpaceUsage(args []string) {
	cmd := exec.Command("df", args...)
	runCommand(cmd)
}

func displayDiskUsage(args []string) {
	cmd := exec.Command("du", args...)
	runCommand(cmd)
}

func displayHelp() {
	fmt.Println("Available commands:")
	fmt.Printf("%s:\t%s\n", "cat", "Display the contents of a file")
	fmt.Printf("%s:\t%s\n", "cd", "Change the current directory")
	fmt.Printf("%s:\t%s\n", "cls", "Clear the screen")
	fmt.Printf("%s:\t%s\n", "cp", "Copy files or directories")
	fmt.Printf("%s:\t%s\n", "dir", "List the contents of the current directory")
	fmt.Printf("%s:\t%s\n", "head", "Display the first few lines of a file")
	fmt.Printf("%s:\t%s\n", "help", "Display this help message")
	fmt.Printf("%s:\t%s\n", "less", "Display the contents of a file one page at a time")
	fmt.Printf("%s:\t%s\n", "ls", "List the contents of the current directory")
	fmt.Printf("%s:\t%s\n", "mkdir", "Create a new directory")
	fmt.Printf("%s:\t%s\n", "more", "Display the contents of a file one page at a time")
	fmt.Printf("%s:\t%s\n", "mv", "Move or rename files or directories")
	fmt.Printf("%s:\t%s\n", "pwd", "Print the current directory path")
	fmt.Printf("%s:\t%s\n", "rm", "Remove a file or directory")
	fmt.Printf("%s:\t%s\n", "rmdir", "Remove an empty directory")
	fmt.Printf("%s:\t%s\n", "tail", "Display the last few lines of a file")
	fmt.Printf("%s:\t%s\n", "touch", "Create an empty file")
	fmt.Printf("%s:\t%s\n", "find", "Searches for files or directories based on specified criteria")
	fmt.Printf("%s:\t%s\n", "grep", "Searches for text within files")
	fmt.Printf("%s:\t%s\n", "chmod", "Modifies file permissions")
	fmt.Printf("%s:\t%s\n", "chown", "Changes the owner of a file or directory")
	fmt.Printf("%s:\t%s\n", "chgrp", "Changes the group of a file or directory")
	fmt.Printf("%s:\t%s\n", "df", "Displays disk space usage")
	fmt.Printf("%s:\t%s\n", "du", "Displays the disk usage of files and directories")
	// Add more commands here
}

func runCommand(cmd *exec.Cmd) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println(string(output))
}

func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
