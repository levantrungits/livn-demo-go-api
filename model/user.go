// Package model by TRUNGLV
package model

// User struct is table users in data base.
// Save user object to db.
// Response user json for response.
type User struct {
	ID     int    // ok ID
	Name   string // ok Name
	Gender string // BUG(1): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
	Email  string // BUG(2): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
}
