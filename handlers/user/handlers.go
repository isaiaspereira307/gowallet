package handlers

import (
	"github.com/isaiaspereira307/gowallet/internal/db"
)

var queries *db.Queries

func InitializeUserHandlers(q *db.Queries) {
	queries = q
}
