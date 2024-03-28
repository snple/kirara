package edge

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/snple/kirara/consts"
	"github.com/snple/kirara/edge/model"
	"github.com/snple/kirara/util"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

func Seed(db *bun.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = seed(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func seed(db bun.Tx) error {
	var err error
	ctx := context.Background()

	{
		seedDevice := func() error {
			device := model.Device{
				ID:      util.RandomID(),
				Name:    consts.DEFAULT_DEVICE,
				Status:  consts.ON,
				Created: time.Now(),
				Updated: time.Now(),
			}

			_, err = db.NewInsert().Model(&device).Exec(ctx)
			if err != nil {
				return err
			}

			return nil
		}

		err = db.NewSelect().Model(&model.Device{}).Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				if err = seedDevice(); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	{
		seedUser := func() error {
			randPass := util.RandPass(10)

			passwd, err := bcrypt.GenerateFromPassword([]byte(randPass), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			user := model.User{
				ID:      util.RandomID(),
				Name:    "root",
				Desc:    "root",
				Pass:    string(passwd),
				Status:  consts.ON,
				Created: time.Now(),
				Updated: time.Now(),
			}

			_, err = db.NewInsert().Model(&user).Exec(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("seed: The initial password for the root user is: %v\n", randPass)

			return nil
		}

		err = db.NewSelect().Model(&model.User{}).Where("name = ?", "root").Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				if err = seedUser(); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
