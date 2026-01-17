package main

import (
	"context"
	"database/sql"
	"fmt"
	"migration/filestorage"
	"migration/sql/repo_sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("sqlite3", "./videos.db") // change this
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	queries := repo_sqlite.New(conn)
	AddAllModulesFromJson(ctx, queries)
	AddAllVideosFromJson(ctx, queries)
}

func AddAllModulesFromJson(ctx context.Context, queries *repo_sqlite.Queries) {
	queries.AddCourse(ctx, "Golang course")
	modulesJson := filestorage.GetJsonConfig("./config/mod.json")
	for _, m := range modulesJson {
		mod := repo_sqlite.AddModuleParams{
			Name:     m.Name,
			Folder:   m.Folder,
			CourseID: 1,
		}
		err := queries.AddModule(ctx, mod)
		if err!= nil {
            panic(err)
        }
	}
}

func AddAllVideosFromJson(ctx context.Context, queries *repo_sqlite.Queries) {
	modulesJson := filestorage.GetJsonConfig("./config/mod.json")
	modulosDB, err := queries.GetAllModules(ctx, 1)
	if err != nil {
		panic(err)
	}

	for _, m := range modulosDB {
		mj := modulesJson[m.Folder]
		for video := range mj.Videos {

			video := repo_sqlite.AddVideoParams{
				Name:     mj.Videos[video],
				VideoKey: "---",
				ModuleID: m.ID,
			}
			fmt.Println(video)
			err := queries.AddVideo(ctx, video)
			if err != nil {
				panic(err)
			}
		}
	}
}
