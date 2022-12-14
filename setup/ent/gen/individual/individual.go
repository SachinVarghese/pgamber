// Code generated by ent, DO NOT EDIT.

package individual

const (
	// Label holds the string label denoting the individual type in the database.
	Label = "individual"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldWorkclass holds the string denoting the workclass field in the database.
	FieldWorkclass = "workclass"
	// FieldEducation holds the string denoting the education field in the database.
	FieldEducation = "education"
	// FieldMaritalStatus holds the string denoting the marital_status field in the database.
	FieldMaritalStatus = "marital_status"
	// FieldOccupation holds the string denoting the occupation field in the database.
	FieldOccupation = "occupation"
	// FieldRelationship holds the string denoting the relationship field in the database.
	FieldRelationship = "relationship"
	// FieldRace holds the string denoting the race field in the database.
	FieldRace = "race"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldCapitalGain holds the string denoting the capital_gain field in the database.
	FieldCapitalGain = "capital_gain"
	// FieldCapitalLoss holds the string denoting the capital_loss field in the database.
	FieldCapitalLoss = "capital_loss"
	// FieldHoursPerWeek holds the string denoting the hours_per_week field in the database.
	FieldHoursPerWeek = "hours_per_week"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// EdgeBracket holds the string denoting the bracket edge name in mutations.
	EdgeBracket = "bracket"
	// Table holds the table name of the individual in the database.
	Table = "individuals"
	// BracketTable is the table that holds the bracket relation/edge.
	BracketTable = "income_brackets"
	// BracketInverseTable is the table name for the IncomeBracket entity.
	// It exists in this package in order to avoid circular dependency with the "incomebracket" package.
	BracketInverseTable = "income_brackets"
	// BracketColumn is the table column denoting the bracket relation/edge.
	BracketColumn = "individual_bracket"
)

// Columns holds all SQL columns for individual fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldWorkclass,
	FieldEducation,
	FieldMaritalStatus,
	FieldOccupation,
	FieldRelationship,
	FieldRace,
	FieldSex,
	FieldCapitalGain,
	FieldCapitalLoss,
	FieldHoursPerWeek,
	FieldCountry,
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
	AgeValidator func(float64) error
	// WorkclassValidator is a validator for the "workclass" field. It is called by the builders before save.
	WorkclassValidator func(int) error
	// EducationValidator is a validator for the "education" field. It is called by the builders before save.
	EducationValidator func(int) error
	// MaritalStatusValidator is a validator for the "marital_status" field. It is called by the builders before save.
	MaritalStatusValidator func(int) error
	// OccupationValidator is a validator for the "occupation" field. It is called by the builders before save.
	OccupationValidator func(int) error
	// RelationshipValidator is a validator for the "relationship" field. It is called by the builders before save.
	RelationshipValidator func(int) error
	// RaceValidator is a validator for the "race" field. It is called by the builders before save.
	RaceValidator func(int) error
	// SexValidator is a validator for the "sex" field. It is called by the builders before save.
	SexValidator func(int) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(int) error
)
