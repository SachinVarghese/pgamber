// Code generated by ent, DO NOT EDIT.

package individual

const (
	// Label holds the string label denoting the individual type in the database.
	Label = "individual"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// Table holds the table name of the individual in the database.
	Table = "individuals"
)

// Columns holds all SQL columns for individual fields.
var Columns = []string{
	FieldID,
	FieldAge,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
)
