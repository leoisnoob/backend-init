package cmd

import (
	"crypto/md5"
	"fmt"
	"log"
	"reflect"

	"example-backend/config"
	"example-backend/model"
	"example-backend/store"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var models = []interface{}{
	&model.User{},
}

func Migrate(c *cli.Context) error {

	store.InitMysql(config.Cfg.MysqlDNS)
	db := store.GlobalStore.MysqlCli
	// for _, v := range models {
	// 	key := reflect.TypeOf(v).Elem().Name()
	// 	log.Printf("drop table %s\n", key)
	// 	if err := db.Migrator().DropTable(v); err != nil {
	// 		log.Fatalf("drop table %s failed, error : %s\n", key, err)
	// 	}
	// }

	for _, v := range models {
		key := reflect.TypeOf(v).Elem().Name()
		log.Printf("migrate table %s\n", key)
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB,CHARSET=utf8mb4,COLLATE=utf8mb4_general_ci").AutoMigrate(v); err != nil {
			log.Fatalf("migrate %s failed, error %s\n", key, err)
		}
	}
	log.Println("migrate success ...........")
	// mockData(db)
	return nil
}

func mockData(db *gorm.DB) {
	// mockUser(db)
	// mockPullTasks(db)
	// mockDatesets(db)
	// mockLabelTasks(db)
}

func mockUser(db *gorm.DB) {
	users := []model.User{
		{
			UserID: "admin",
			IsRoot: 1,
			Pass:   md5Hash("admin"),
		},
		{
			UserID: "lissel",
			Pass:   md5Hash("123456"),
		},
		{
			UserID: "jack",
			Pass:   md5Hash("123456"),
		},
	}
	for i := range users {
		if err := db.Create(&users[i]).Error; err != nil {
			log.Printf("create user failed, error %s\n", err)
		}
	}
}

func md5Hash(v string) string {
	if v == "" {
		return ""
	}
	h := md5.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
