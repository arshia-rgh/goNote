package main

import (
	"fmt"
	"goNote/note"
	"goNote/utils"
	"os"
)

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
		os.Exit(1)
	}
}

func main() {

	for {

		fmt.Println("welcome to the GoNote....")
		fmt.Println("Select what do you want to do? ")
		fmt.Println("1 - Create new note")
		fmt.Println("2 - Check exists notes")
		fmt.Println("3 - Exit")

		option, err := utils.GetUserData("Please enter your choice : ")
		checkError(err, "Failed to get user input ")

		if option == "1" {
			title, err := utils.GetUserData("Please enter the title of your note: ")
			checkError(err, "Failed to get title")

			content, err := utils.GetUserData("Please enter the content of your note: ")
			checkError(err, "Failed to get the content")

			newNote, err := note.New(title, content)
			checkError(err, "Failed to create new note")

			fmt.Printf("New note added with %v title and following content:\n %v\n", newNote.Title, newNote.Content)

			fileName, err := utils.GetUserData("Please enter your file name to save your notes (must suffix with .json)")
			checkError(err, "Failed to get the file name")

			err = utils.WriteToFile(fileName, *newNote)
			checkError(err, "Failed to write to the file")

			fmt.Printf("Your note saved to the %v file successfully\n", fileName)

			continue

		} else if option == "2" {

		} else if option == "3" {
			fmt.Println("Goodbye...")
			break

		} else {
			fmt.Println("Invalid option please try again")

		}
	}
}
