package main

import (
	"fmt"
	"goNote/note"
	"goNote/utils"
	"os"
)

func main() {
	fileName, err := utils.GetUserData("Please enter your file name to save and read your notes (must suffix with .json)")
	checkError(err, "Failed to get the file name")
	utils.ClearTerminal()

	for {

		fmt.Println("welcome to the GoNote....")
		fmt.Println("Select what do you want to do? ")
		fmt.Println("1 - Create new note")
		fmt.Println("2 - Check exists notes")
		fmt.Println("3 - Exit")

		option, err := utils.GetUserData("Please enter your choice : ")
		utils.ClearTerminal()
		checkError(err, "Failed to get user input ")

		if option == "1" {
			title, err := utils.GetUserData("Please enter the title of your note: ")
			checkError(err, "Failed to get title")

			content, err := utils.GetUserData("Please enter the content of your note: ")
			checkError(err, "Failed to get the content")
			utils.ClearTerminal()

			newNote, err := note.New(title, content)
			checkError(err, "Failed to create new note")

			fmt.Printf("New note added with %v title and following content:\n %v\n", newNote.Title, newNote.Content)

			err = utils.WriteToFile(fileName, *newNote)
			checkError(err, "Failed to write to the file")

			fmt.Printf("Your note saved to the %v file successfully\n", fileName)

			fmt.Println("Select any key to redirect...")
			fmt.Scanln()

			utils.ClearTerminal()

			continue

		} else if option == "2" {
			notes, err := utils.ReadFromFile(fileName)
			checkError(err, "Failed to read from file")

			fmt.Println(notes)

			fmt.Println("Select any key to redirect...")
			fmt.Scanln()
			utils.ClearTerminal()

		} else if option == "3" {
			fmt.Println("Goodbye...")
			break

		} else {
			fmt.Println("Invalid option please try again")

		}
	}
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
		os.Exit(1)
	}
}
