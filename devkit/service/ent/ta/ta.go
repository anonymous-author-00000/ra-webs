// Code generated by ent, DO NOT EDIT.

package ta

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the ta type in the database.
	Label = "ta"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEvidence holds the string denoting the evidence field in the database.
	FieldEvidence = "evidence"
	// FieldRepository holds the string denoting the repository field in the database.
	FieldRepository = "repository"
	// FieldCommitID holds the string denoting the commit_id field in the database.
	FieldCommitID = "commit_id"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// Table holds the table name of the ta in the database.
	Table = "tas"
)

// Columns holds all SQL columns for ta fields.
var Columns = []string{
	FieldID,
	FieldEvidence,
	FieldRepository,
	FieldCommitID,
	FieldPublicKey,
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

// OrderOption defines the ordering options for the TA queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEvidence orders the results by the evidence field.
func ByEvidence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEvidence, opts...).ToFunc()
}

// ByRepository orders the results by the repository field.
func ByRepository(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRepository, opts...).ToFunc()
}

// ByCommitID orders the results by the commit_id field.
func ByCommitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommitID, opts...).ToFunc()
}
