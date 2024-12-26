package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: Adoption{}},
		{Model: Pet{}},
		{Model: Admin{}},
		{Model: PetImage{}},
	}
}
