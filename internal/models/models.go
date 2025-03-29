package models

import "time"

// In models/user.go
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"unique" json:"username"`
	Password     string    `gorm:"-" json:"password"` // Only for input, not stored
	PasswordHash string    `json:"-"`                 // Only stored in DB
	CreatedAt    time.Time `json:"created_at"`
}

type Puzzle struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Date     time.Time `gorm:"unique" json:"date"`
	Subjects string    `json:"subjects"` // Comma-separated, e.g., "Math,Logic,Riddles"
	Content  string    `json:"content"`
	Solution string    `json:"-"`
}

type UserPuzzle struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	UserID   int       `gorm:"index" json:"user_id"`
	PuzzleID int       `gorm:"index" json:"puzzle_id"`
	SolvedAt time.Time `json:"solved_at"`
}
