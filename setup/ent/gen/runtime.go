// Code generated by ent, DO NOT EDIT.

package gen

import (
	"github.com/SachinVarghese/pgamber/setup/ent/gen/individual"
	"github.com/SachinVarghese/pgamber/setup/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	incomebracketFields := schema.IncomeBracket{}.Fields()
	_ = incomebracketFields
	individualFields := schema.Individual{}.Fields()
	_ = individualFields
	// individualDescAge is the schema descriptor for age field.
	individualDescAge := individualFields[0].Descriptor()
	// individual.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	individual.AgeValidator = individualDescAge.Validators[0].(func(float64) error)
	// individualDescWorkclass is the schema descriptor for workclass field.
	individualDescWorkclass := individualFields[1].Descriptor()
	// individual.WorkclassValidator is a validator for the "workclass" field. It is called by the builders before save.
	individual.WorkclassValidator = individualDescWorkclass.Validators[0].(func(int) error)
	// individualDescEducation is the schema descriptor for education field.
	individualDescEducation := individualFields[2].Descriptor()
	// individual.EducationValidator is a validator for the "education" field. It is called by the builders before save.
	individual.EducationValidator = individualDescEducation.Validators[0].(func(int) error)
	// individualDescMaritalStatus is the schema descriptor for marital_status field.
	individualDescMaritalStatus := individualFields[3].Descriptor()
	// individual.MaritalStatusValidator is a validator for the "marital_status" field. It is called by the builders before save.
	individual.MaritalStatusValidator = individualDescMaritalStatus.Validators[0].(func(int) error)
	// individualDescOccupation is the schema descriptor for occupation field.
	individualDescOccupation := individualFields[4].Descriptor()
	// individual.OccupationValidator is a validator for the "occupation" field. It is called by the builders before save.
	individual.OccupationValidator = individualDescOccupation.Validators[0].(func(int) error)
	// individualDescRelationship is the schema descriptor for relationship field.
	individualDescRelationship := individualFields[5].Descriptor()
	// individual.RelationshipValidator is a validator for the "relationship" field. It is called by the builders before save.
	individual.RelationshipValidator = individualDescRelationship.Validators[0].(func(int) error)
	// individualDescRace is the schema descriptor for race field.
	individualDescRace := individualFields[6].Descriptor()
	// individual.RaceValidator is a validator for the "race" field. It is called by the builders before save.
	individual.RaceValidator = individualDescRace.Validators[0].(func(int) error)
	// individualDescSex is the schema descriptor for sex field.
	individualDescSex := individualFields[7].Descriptor()
	// individual.SexValidator is a validator for the "sex" field. It is called by the builders before save.
	individual.SexValidator = individualDescSex.Validators[0].(func(int) error)
	// individualDescCountry is the schema descriptor for country field.
	individualDescCountry := individualFields[11].Descriptor()
	// individual.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	individual.CountryValidator = individualDescCountry.Validators[0].(func(int) error)
}
