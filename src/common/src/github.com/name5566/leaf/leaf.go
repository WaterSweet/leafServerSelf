package leaf

import (
	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/console"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"os"
	"os/signal"
)

func Run(mods ...module.Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Leaf %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)


	//d, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil{
	//	fmt.Println("Connect to redis error", err)
	//	return
	//}
	//fmt.Println("Connect to redis successful!")
	//defer d.Close()
	//
	//
	//_, err = d.Do("SET", "myKey", "superWang")
	//if err != nil{
	//	fmt.Println("redis set failed", err)
	//}
	//
	//userName, err := redis.String(d.Do("GET", "myKey"))
	//if err != nil{
	//	fmt.Println("redis get failed:", err)
	//}else {
	//	fmt.Println("Get userName", userName)
	//}


	sig := <-c
	log.Release("Leaf closing down (signal: %v)", sig)
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
