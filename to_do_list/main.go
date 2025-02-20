package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func show_list(todo_list []string){
	fmt.Println("To Do List:")
	fmt.Println("----------------")
	for i := 0; i < len(todo_list); i++ {
		fmt.Println(i+1, todo_list[i])
	}
	fmt.Println("----------------")
}


func main() {
	var todo_list []string

	file, err := os.OpenFile("todo_list.txt", os.O_RDWR|os.O_CREATE, 0755)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		todo_list = append(todo_list, scanner.Text())
	}
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for{
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		
		commands := strings.Split(input, " ")
		command := strings.TrimSpace(commands[0])
	
		item := strings.TrimSpace(strings.Join(commands[1:], " "))
	
		if command  == "add" {
			todo_list = append(todo_list, item)
			show_list(todo_list)
		}else if command == "remove" {
			if item == "all"{
				todo_list = []string{}
				show_list(todo_list)
				continue
			}
			index, _ := strconv.Atoi(item)
			fmt.Println(index)
			if index > 0 && index <= len(todo_list) {
				todo_list = append(todo_list[:index-1], todo_list[index:]...)
				show_list(todo_list)
			}else{
				fmt.Println("Invalid index")
			}
		}else if command == "save"{
			file, err := os.OpenFile("todo_list.txt", os.O_WRONLY|os.O_TRUNC, 0755)
			if err != nil {
                fmt.Println(err)
                return
            }
            defer file.Close()
			for _, item := range todo_list {
				file.WriteString(item + "\n")
			}
			break
		}else if command == "show" {
			show_list(todo_list)
		}
	}
}