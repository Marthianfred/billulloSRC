package Models

import "github.com/jinzhu/gorm"


type UserModel struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Edad        uint   `json:"edad"`
	Telefono    string `json:"telefono" gorm:"size:20"`
	Direccion   string `json:"direccion"`
	Email       string `json:"email"`
	Descripcion string `json:"descripcion" gorm:"type:TEXT"`
}