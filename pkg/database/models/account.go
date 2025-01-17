package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const AccountPasswdLen = 8

type Account struct {
	ID    uint `gorm:"primarykey;autoIncrement;"`
	Email string

	Type      AcountType
	FirstName string
	LastName  string
	Passwd    string // Considered as plaintext, but can be encrypted by frontend
}

type Auth struct {
	Email           string `gorm:"primarykey;"` // Not a refer key !!!
	AuthCode        string
	AuthCodeExpires time.Time
}

type AcountType string // Type of account

const (
	PatientType AcountType = "patient"
	DoctorType  AcountType = "doctor"
	AdminType   AcountType = "admin"
)

type Doctor struct {
	ID           uint `gorm:"primarykey;autoIncrement;"`
	DepartmentID uint

	AccountID uint
	CaseID    uint
	ChatID    uint
	// Cases     []cases.Case `gorm:"foreignkey:ID"`
	// Chats     []chat.Chat  `gorm:"foreignkey:ID"`
}

type Patient struct {
	ID uint `gorm:"primarykey;autoIncrement;"`

	AccountID uint
	CaseID    uint
	ChatID    uint
	// Account Account      `gorm:"foreignkey:ID"`
	// Cases   []cases.Case `gorm:"foreignkey:ID"`
	// Chats   []chat.Chat  `gorm:"foreignkey:ID"`
}

/**
 * @brief private method for hashing password
 */
func (u *Account) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Passwd), bcrypt.DefaultCost)
	u.Passwd = string(bytes)
}

/**
 * @brief private method for generateing token
 */
func (u *Account) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().Add(7 * 24 * time.Hour),
	})

	jwtKey := []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
