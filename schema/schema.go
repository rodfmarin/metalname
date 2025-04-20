// Package schema defines schemas for names, etc.
package schema

type Name struct {
	FirstName string
	LastName  string
}

func (n Name) String() string {
	return n.FirstName + " " + n.LastName
} 
