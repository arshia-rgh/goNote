package main

import (
	"fmt"
	"goNote/note"
	"goNote/utils"
)

func main() {

	for {

		fmt.Println("welcome to the GoNote....")
		fmt.Println("Select what do you want to do? ")
		fmt.Println("1 - Create new note")
		fmt.Println("2 - Check exists notes")
		fmt.Println("3 - Exit")

		option, err := utils.GetUserData("Please enter your choice : ")

		if err != nil {
			fmt.Println(err)
			return
		}
		if option == "1" {
			var title string
			var content string

			title, err = utils.GetUserData("Please enter the title of your note: ")

			if err != nil {
				fmt.Println(err)
				return
			}

			content, err = utils.GetUserData("Please enter the content of your note: ")

			if err != nil {
				fmt.Println(err)
				return
			}
			var newNote *note.Note

			newNote, err = note.New(title, content)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("New note added with %v title and following content:\n %v", newNote.Title, newNote.Content)

			var fileName string
			fileName, err = utils.GetUserData("Please enter your file name to save your notes (must suffix with .json)")

			if err != nil {
				fmt.Println(err)
				return
			}

			err = utils.WriteToFile(fileName, *newNote)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Your note saved to the %v file successfully\n", fileName)

			continue

		} else if option == "2" {

		} else if option == "3" {

		} else {

		}
	}
}
