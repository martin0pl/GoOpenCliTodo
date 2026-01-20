package main

import (
	"GoOpenCliTodo/tache"
	"fmt"
	"slices"
)

func main() {
	var taches []tache.Tache

	tache_test := *tache.NewTache("Tache test", "10/10/2010")
	taches = append(taches, tache_test)

	var execution bool = true

	for execution {
		var message string

		fmt.Print("\nEntrer une commande (exit pour quitter) : ")
		fmt.Scan(&message)

		if message == "exit" {
			execution = false
		} else if message == "list" {

			if len(taches) == 0 {
				fmt.Println("\nPas de tâches")
			} else {
				fmt.Println("")

				for i, tache := range taches {
					fmt.Print(i, " - ")
					tache.Afficher()
				}
			}

		} else if message == "newtask" {
			var name, echeance string

			fmt.Scanln(&name, &echeance)

			taches = append(taches, *tache.NewTache(name, echeance))
		} else if message == "done" {
			var index int

			fmt.Scanln(&index)

			taches[index].Realisee = true
		} else if message == "notdone" {
			var index int

			fmt.Scanln(&index)

			taches[index].Realisee = false
		} else if message == "delete" {
			var index int

			fmt.Scanln(&index)

			taches = slices.Delete(taches, index, index+1)
		} else if message == "cleandone" {

			for i := len(taches) - 1; i >= 0; i-- {
				if taches[i].Realisee == true {
					taches = slices.Delete(taches, i, i+1)
				}
			}

		}

	}
}
