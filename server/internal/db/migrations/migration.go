package migrations

import (
	"github.com/369guang/tg-im/core/database"
	"github.com/369guang/tg-im/server/internal/db/model"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&model.User{}, &model.UserInfo{},
		&model.Friend{}, &model.Conversation{},
		&model.Group{}, &model.GroupMember{},
		&model.Message{},
	)

	if err != nil {
		panic(err)
	}

}
