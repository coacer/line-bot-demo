// Code generated by yo. DO NOT EDIT.
// Package golang contains the types.
package golang

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// User represents a row from 'User'.
type User struct {
	ID               int64     `spanner:"Id" json:"Id"`                             // Id
	LineBotChannelID int64     `spanner:"LineBotChannelId" json:"LineBotChannelId"` // LineBotChannelId
	LineUID          string    `spanner:"LineUID" json:"LineUID"`                   // LineUID
	CreatedAt        time.Time `spanner:"CreatedAt" json:"CreatedAt"`               // CreatedAt
	UpdatedAt        time.Time `spanner:"UpdatedAt" json:"UpdatedAt"`               // UpdatedAt
}

func UserPrimaryKeys() []string {
	return []string{
		"Id",
	}
}

func UserColumns() []string {
	return []string{
		"Id",
		"LineBotChannelId",
		"LineUID",
		"CreatedAt",
		"UpdatedAt",
	}
}

func UserWritableColumns() []string {
	return []string{
		"Id",
		"LineBotChannelId",
		"LineUID",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (u *User) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "Id":
			ret = append(ret, &u.ID)
		case "LineBotChannelId":
			ret = append(ret, &u.LineBotChannelID)
		case "LineUID":
			ret = append(ret, &u.LineUID)
		case "CreatedAt":
			ret = append(ret, &u.CreatedAt)
		case "UpdatedAt":
			ret = append(ret, &u.UpdatedAt)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (u *User) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "Id":
			ret = append(ret, u.ID)
		case "LineBotChannelId":
			ret = append(ret, u.LineBotChannelID)
		case "LineUID":
			ret = append(ret, u.LineUID)
		case "CreatedAt":
			ret = append(ret, u.CreatedAt)
		case "UpdatedAt":
			ret = append(ret, u.UpdatedAt)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newUser_Decoder returns a decoder which reads a row from *spanner.Row
// into User. The decoder is not goroutine-safe. Don't use it concurrently.
func newUser_Decoder(cols []string) func(*spanner.Row) (*User, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*User, error) {
		var u User
		ptrs, err := u.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &u, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (u *User) Insert(ctx context.Context) *spanner.Mutation {
	values, _ := u.columnsToValues(UserWritableColumns())
	return spanner.Insert("User", UserWritableColumns(), values)
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (u *User) Update(ctx context.Context) *spanner.Mutation {
	values, _ := u.columnsToValues(UserWritableColumns())
	return spanner.Update("User", UserWritableColumns(), values)
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (u *User) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	values, _ := u.columnsToValues(UserWritableColumns())
	return spanner.InsertOrUpdate("User", UserWritableColumns(), values)
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (u *User) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, UserPrimaryKeys()...)

	values, err := u.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "User.UpdateColumns", "User", err)
	}

	return spanner.Update("User", colsWithPKeys, values), nil
}

// FindUser gets a User by primary key
func FindUser(ctx context.Context, db YORODB, id int64) (*User, error) {
	key := spanner.Key{id}
	row, err := db.ReadRow(ctx, "User", key, UserColumns())
	if err != nil {
		return nil, newError("FindUser", "User", err)
	}

	decoder := newUser_Decoder(UserColumns())
	u, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindUser", "User", err)
	}

	return u, nil
}

// ReadUser retrieves multiples rows from User by KeySet as a slice.
func ReadUser(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*User, error) {
	var res []*User

	decoder := newUser_Decoder(UserColumns())

	rows := db.Read(ctx, "User", keys, UserColumns())
	err := rows.Do(func(row *spanner.Row) error {
		u, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, u)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadUser", "User", err)
	}

	return res, nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := u.columnsToValues(UserPrimaryKeys())
	return spanner.Delete("User", spanner.Key(values))
}
