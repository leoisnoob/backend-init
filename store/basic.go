package store

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GlobalStore Store
var zapLog *zap.Logger

func InjectLog(l *zap.Logger) {
	zapLog = l
}

type Store struct {
	MysqlCli *gorm.DB
	MinioCli *minio.Client
}

func InitMysql(dns string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // Disable color
		},
	)
	log.Println("dns ", dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	GlobalStore.MysqlCli = db

}

func InitMinio(url, accessKey, secretKey, bucketName string) {
	minioClient, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	found, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		panic(err)
	}
	if !found {
		err = fmt.Errorf("bucket %s not exist", bucketName)
		panic(err)
	}
	GlobalStore.MinioCli = minioClient
}
