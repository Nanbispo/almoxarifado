package entity

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"nome" binding:"required,min=2"`
	Description string `json:"descricao,omitempty"`
	Department  string `json:"departamento" binding:"required"`
}

