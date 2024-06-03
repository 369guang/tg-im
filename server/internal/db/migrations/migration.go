package migrations

import (
	"github.com/space-water-bear/tg-im/internal/database"
	"github.com/space-water-bear/tg-im/server/internal/db/model"
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
