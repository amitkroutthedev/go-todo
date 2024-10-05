package main

import "fmt"

type TODOItems struct{
	
}

func main() {
	var choices int;
	allTodoList:=[]string{}
	fmt.Println("ToDO application in GO")
    
    for{
		fmt.Println("Enter a Choices")
		fmt.Println("1. Add a task")
		fmt.Println("2. Delete a task")
		fmt.Println("3. Edit a task")
		fmt.Println("4. View all task")
		fmt.Println("5. Exit form the application")
   
		fmt.Scan(&choices)
   
		fmt.Println("User selected",choices)

		switch choices{
		case 1:
			var usertask string
			fmt.Println("Enter the task")
			fmt.Scan(&usertask)
			allTodoList = append(allTodoList,usertask)
		case 2:
			var taskNumber int
			fmt.Println("Enter the task number to delete:")
			fmt.Scan(&taskNumber)
			if taskNumber > 0 && taskNumber <= len(allTodoList) {
				allTodoList = append(allTodoList[:taskNumber], allTodoList[taskNumber+1:]...)
			} else {
				fmt.Println("Invalid task number")
			}
		case 3:
			 var taskNumber int
			 var newTask string
			 fmt.Println("Enter the task to edit:")
             fmt.Scan(&taskNumber)
			 if taskNumber >= 0 && taskNumber <= len(allTodoList) {
				fmt.Println("Enter the new task:")
				fmt.Scan(&newTask)
				allTodoList[taskNumber] = newTask
			 }else{
				fmt.Println("Invalid task number")
			 }
			case 4:
			fmt.Println("All tasks:")
			for i, task:= range allTodoList{
				fmt.Println(i,". ",task)
			}
		case 5:
			fmt.Println("Exiting the application")
			return
		default:
			fmt.Println("Enter the correct choice 1 to 5")
	}
	}
}
