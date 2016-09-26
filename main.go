package main

import (
	"fmt"
	"os"

	"github.com/midnightfreddie/goleveldb/leveldb"
	"github.com/midnightfreddie/goleveldb/leveldb/opt"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "MCPE Tool"
	app.Version = "0.0.0"
	app.Usage = "A utility to access Minecraft Portable Edition .mcworld exported world files."

	app.Commands = []cli.Command{
		{
			Name:    "keys",
			Aliases: []string{"k"},
			Usage:   "Lists all keys in the database. Be sure to include the path to the db, e.g. 'McpeTool keys db'",
			Action: func(c *cli.Context) error {
				o := &opt.Options{
					ReadOnly: true,
				}
				db, err := leveldb.OpenFile(c.Args().First(), o)
				if err != nil {
					panic("error")
				}
				defer db.Close()

				iter := db.NewIterator(nil, nil)
				for iter.Next() {
					key := iter.Key()
					switch {
					case len(key) == 9:
						switch key[8] {
						case 0x30, 0x31, 0x32, 0x76:
							fmt.Println(key)
						default:
							fmt.Println(string(key[:]))
						}
					case len(key) == 13:
						switch key[12] {
						case 0x30, 0x31, 0x32, 0x76:
							fmt.Println(key)
						default:
							fmt.Println(string(key[:]))
						}
					default:
						fmt.Println(string(key[:]))
					}
				}
				iter.Release()
				err = iter.Error()
				if err != nil {
					panic(err.Error())
				}
				return nil
			},
		},
		{
			Name:    "develop",
			Aliases: []string{"dev"},
			Usage:   "Random thing the dev is working on",
			Action: func(c *cli.Context) error {
				db, err := leveldb.OpenFile(c.Args().First(), nil)
				if err != nil {
					panic("error")
				}
				defer db.Close()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
