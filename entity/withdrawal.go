package entity

import "time"

type withdrawalStatus string

const (
	pendingStatus   withdrawalStatus = "PENDENTE"
	completedStatus withdrawalStatus = "CONCLUIDO"
	lateStatus      withdrawalStatus = "ATRASADO"
)

type Withdrawal struct {
	ID                  uint             `json:"id" gorm:"primaryKey"`

	ProductID           uint             `json:"produto_id" binding:"required"`
	Product             Product          `json:"produto,omitempty" gorm:"foreignKey:ProductID"`

	RepresentativeID    uint             `json:"retirante_id" binding:"required"`
	Representative      Person           `json:"retirante,omitempty" gorm:"foreignKey:RepresentativeID"`

	AuthorizingLeaderID uint             `json:"lider_autorizador_id" binding:"required"`
	AuthorizingLeader   Person           `json:"lider_autorizador,omitempty" gorm:"foreignKey:AuthorizingLeaderID"`

	DateRetrieved       time.Time        `json:"data_retirada"`
	ExpectedReturnDate  time.Time        `json:"data_devolucao_prevista" binding:"required"`
	ActualReturnDate    *time.Time       `json:"data_devolucao_efetiva,omitempty"`
	Status              withdrawalStatus `json:"status" gorm:"default:'PENDENTE'"`
}