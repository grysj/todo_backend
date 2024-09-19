package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreateUserTxParams struct {
	Username string `json:"username"`
}

type CreateUserTxResult struct {
	User User `json:"user"`
}

func (store *Store) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		var check bool
		check, err = q.CheckIfUsernameExists(ctx, arg.Username)
		if err != nil {
			return err
		}
		if check {
			return fmt.Errorf("User with a given username already exists")
		}
		result.User, err = q.CreateUser(ctx, arg.Username)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

type AddPermissionTxParams struct {
	FromUser       int32 `json:"from_user"`
	ToUser         int32 `json:"to_user"`
	ListID         int32 `json:"list"`
	PermissionType int32 `json:"permission_type"`
}

type AddPermissionTxResult struct {
	Permission Permission `json:"permission"`
	ToUser     User       `json:"to_user"`
}

func (store *Store) AddPermissionTx(ctx context.Context, arg AddPermissionTxParams) (AddPermissionTxResult, error) {
	var result AddPermissionTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		check, err := q.CheckUserPermission(ctx, CheckUserPermissionParams{
			FromUser: arg.FromUser,
			ListID:   arg.FromUser,
		})
		if err != nil {
			return err
		}
		if check >= 3 {
			return fmt.Errorf("User not permitted to add permission")
		}
		result.Permission, err = q.CreatePermission(ctx, CreatePermissionParams{
			FromUser: arg.FromUser,
			ToUser:   arg.ToUser,
			ListID:   arg.ListID,
			PermType: arg.PermissionType,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

type Change int

const (
	EditContent Change = iota
	EditCheck
	EditPosition
	AddPoint
	DeletePoint
)

type ChangeParams struct {
	ChangeType Change `json:"change"`
	PointID    int32  `json:"point_id,omitempty"`
	Bool       bool   `json:"bool,omitempty"`
	Text       string `json:"text,omitempty"`
	Position   int32  `json:"position,omitempty"`
}

type ModifyContentTxParams struct {
	UserID  int32          `json:"user_id"`
	ListID  int32          `json:"list_id"`
	Changes []ChangeParams `json:"changes"`
}
type ModifyContentTxResult struct {
	ListPoints []ListPoint `json:"list_points"`
}

func (store *Store) ModifyContentTx(ctx context.Context, arg ModifyContentTxParams) (ModifyContentTxResult, error) {
	var result ModifyContentTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		check, err := q.CheckUserPermission(ctx, CheckUserPermissionParams{
			FromUser: arg.UserID,
			ListID:   arg.ListID,
		})
		if err != nil {
			return err
		}
		if check >= 3 {
			return fmt.Errorf("User not permitted to add point to this list")
		}

		for _, change := range arg.Changes {
			switch {
			case change.ChangeType == EditContent:
				err = q.ChangePointContent(ctx, ChangePointContentParams{
					PointID: change.PointID,
					NewText: change.Text,
				})
			case change.ChangeType == EditCheck:
				err = q.ChangePointCheck(ctx, ChangePointCheckParams{
					PointID: change.PointID,
					Checked: change.Bool,
				})
			case change.ChangeType == EditPosition:
				err = q.ChangePointPosition(ctx, ChangePointPositionParams{
					PointID: change.PointID,
					NewPos:  change.Position,
				})
			case change.ChangeType == AddPoint:
				_, err = q.CreatePoint(ctx, CreatePointParams{
					ListID:   arg.ListID,
					Content:  change.Text,
					Position: change.Position,
					AddedBy:  arg.UserID,
				})
			case change.ChangeType == DeletePoint:
				err = q.DeletePoint(ctx, change.PointID)
			default:
				err = fmt.Errorf("Invalid type of change")
			}
			if err != nil {
				return err
			}
		}
		result.ListPoints, err = q.GetPointsByListID(ctx, arg.ListID)
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
