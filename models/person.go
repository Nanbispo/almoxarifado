type Person struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"nome" binding:"required"`
	Department  string `json:"departamento" binding:"required"`
}