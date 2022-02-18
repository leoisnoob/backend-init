package cmd

import (
	"example-backend/config"
	"example-backend/server"
	"example-backend/store"

	"github.com/urfave/cli/v2"
)

var configPath string

func GetApp() *cli.App {
	app := &cli.App{
		Name: "example-backend",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "conf",
				Value:       "config.yml",
				Destination: &configPath,
			},
		},
		Before: cli.BeforeFunc(func(c *cli.Context) error {
			config.InitConfig(configPath)
			server.InitLogger()
			return nil
		}),
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "serve http api",
				Action: server.Serve,
				Before: cli.BeforeFunc(func(c *cli.Context) error {
					store.InitMysql(config.Cfg.MysqlDNS)
					cfg2 := config.Cfg.Minio
					store.InitMinio(cfg2.URL, cfg2.AccessKey, cfg2.SecretKey, cfg2.BucketName)
					return nil
				}),
			},
			{
				Name:   "migrate",
				Usage:  "migrate db table",
				Action: Migrate,
			},
		},
	}
	return app
}
