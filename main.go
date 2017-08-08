package main

import (
	"database/sql"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type message struct {
	isActive chan int
}

type user struct {
	id  int
	fio string
}

var (
	mysql *sql.DB
)

func main() {
	// ВСЕГДА ПРОВЕРЯТЬ ВОЗВРАЩАЕМЫЕ ОШИБКИ!!!1

	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal("Can`t read config.yml.")
	}
print(configFile)
	var config MySQLConnectionConfig
	err = yaml.Unmarshal(configFile, &config)
println(config.Base)
	mysql, err = NewMySQL(&config)
	if err != nil {
		log.Fatalf("Mysql error: %s", err)
	}

	//assignments := storage.GetAllAssignments(true)
	//userAssignments := assignments[200132743]
	//println(userAssignments.UserId)
	//
	//for _, assignment := range userAssignments.Items {
	//	println(assignment.ItemName)
	//}

	//permissionItems := GetAllPermissionItems()
	//for key, item := range permissionItems {
	//	println(key, item.Name, item.ItemType)
	//}

	//
	//h := http.Server{}
	//h.Serve()
}

// bulk checker by user and operations

// cli
// --host
// --port
// --dbDsn
// docopt.go (for cli generations) @see git.rn/projects/PORTAL/repos/paged/browse

// for http use standart golang http (/net/http) @see https://cryptic.io/go-http/

// use channels for updating data from db

// handlers for bulk operations

// use Notify for hardware interruptions
// use logger instead of panic, panic as an exclusion

// teeeestsss!!!!!!!!!!!!!!! don`t run? only tests
