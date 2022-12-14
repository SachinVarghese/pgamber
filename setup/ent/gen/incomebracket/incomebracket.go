// Code generated by ent, DO NOT EDIT.

package incomebracket

import (
	"fmt"
)

const (
	// Label holds the string label denoting the incomebracket type in the database.
	Label = "income_bracket"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldClass holds the string denoting the class field in the database.
	FieldClass = "class"
	// EdgePerson holds the string denoting the person edge name in mutations.
	EdgePerson = "person"
	// Table holds the table name of the incomebracket in the database.
	Table = "income_brackets"
	// PersonTable is the table that holds the person relation/edge.
	PersonTable = "income_brackets"
	// PersonInverseTable is the table name for the Individual entity.
	// It exists in this package in order to avoid circular dependency with the "individual" package.
	PersonInverseTable = "individuals"
	// PersonColumn is the table column denoting the person relation/edge.
	PersonColumn = "individual_bracket"
)

// Columns holds all SQL columns for incomebracket fields.
var Columns = []string{
	FieldID,
	FieldClass,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "income_brackets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"individual_bracket",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Class defines the type for the "class" enum field.
type Class string

// ClassLte50K is the default value of the Class enum.
const DefaultClass = ClassLte50K

// Class values.
const (
	ClassLte50K Class = "lte_50K"
	ClassGt50K  Class = "gt_50K"
)

func (c Class) String() string {
	return string(c)
}

// ClassValidator is a validator for the "class" field enum values. It is called by the builders before save.
func ClassValidator(c Class) error {
	switch c {
	case ClassLte50K, ClassGt50K:
		return nil
	default:
		return fmt.Errorf("incomebracket: invalid enum value for class field: %q", c)
	}
}
