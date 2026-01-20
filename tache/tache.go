package tache

import "fmt"

type Tache struct {
	Name     string
	Echeance string
	Realisee bool
}

// Création d'une nouvelle tâche
func NewTache(name string, echeance string) *Tache {
	return &Tache{name, echeance, false}
}

// Affichage d'une tâche
func (t *Tache) Afficher() {

	var cEcheance string
	if t.Realisee {
		cEcheance = "x"
	} else {
		cEcheance = "_"
	}

	fmt.Println("["+cEcheance+"] Nom = ", t.Name, ", Echeance = ", t.Echeance)
}
