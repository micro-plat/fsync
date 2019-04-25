package main

import "github.com/micro-plat/hydra/hydra"
import _ "github.com/go-sql-driver/mysql"

type fsync struct {
	*hydra.MicroApp
}

func main() {
	app := &fsync{
		hydra.NewApp(
			hydra.WithPlatName("fsync"),
			hydra.WithSystemName("fsync"),
			hydra.WithServerTypes("api"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}
