package usecase

import (
	"aozorarodoku-service/aozorarodokuweb"
	"aozorarodoku-service/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var dbx *sqlx.DB

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {

		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	dbx = sqlx.NewDb(db, "postgres")
	/*
		// Open an atlas driver.
		driver, err := postgres.Open(db)
		if err != nil {
			log.Fatalf("failed opening atlas driver: %s", err)
		}

			// Inspect the created table.
			sch, err := driver.InspectSchema(context.Background(), "public", &schema.InspectOptions{
				Tables: []string{"example"},
			})
			if err != nil {
				log.Fatalf("failed inspecting schema: %s", err)
			}
			tbl, ok := sch.Table("example")
			fmt.Println(tbl)
			fmt.Println(ok)
	*/
	output, err := exec.Command("atlas", "schema", "apply", "--to", "file://"+rootDir()+`/db/schema.hcl`, "--auto-approve", "-u", databaseUrl).Output()
	if err != nil {
		log.Fatalf("failed setup docker test schema: %s, %s", err, output)
	}

	//Run tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

const ContentsFixture = `作品ルビ,作品,作家ルビ,作家,読み手ルビ,読み手,mp3,新着,時間
ああき,ア、秋,だざいおさむ,太宰 治,なかむらあきよ,中村 昭代,rd260.mp3,,6分52秒
ああしんど,ああしんど,いけだしょうえん,池田 焦園,いけどみか,池戸 美香,rd158.mp3,,2分40秒
ああとうきょうはくいだおれ,ああ東京は食い倒れ,ふるかわろっぱ,古川 緑波,のむらようじ,野村 洋二,rd164.mp3,2020-11-01,9分25秒
`

func TestRegisterContentsToDB(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", aozorarodokuweb.AppMetaJsonURL,
		httpmock.NewStringResponder(200, `{
  "version":  "1.7.4",
  "url": "https://aozoraroudoku.jp/app/dummy.csv"
}`),
	)
	httpmock.RegisterResponder("GET", "https://aozoraroudoku.jp/app/dummy.csv",
		httpmock.NewStringResponder(200, ContentsFixture),
	)
	err := RegisterContentsToDB(context.Background(), dbx)
	assert.NoError(t, err)

	contents, err := database.FindContents(context.Background(), dbx)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(contents))
	assert.Equal(t, []database.Content{
		{TitleRuby: "ああき", Title: "ア、秋", AuthorRuby: "だざいおさむ", Author: "太宰 治", SpeakerRuby: "なかむらあきよ", Speaker: "中村 昭代", FileName: "rd260.mp3", NewArrivalDate: "", Time: "6分52秒"},
		{TitleRuby: "ああしんど", Title: "ああしんど", AuthorRuby: "いけだしょうえん", Author: "池田 焦園", SpeakerRuby: "いけどみか", Speaker: "池戸 美香", FileName: "rd158.mp3", NewArrivalDate: "", Time: "2分40秒"},
		{TitleRuby: "ああとうきょうはくいだおれ", Title: "ああ東京は食い倒れ", AuthorRuby: "ふるかわろっぱ", Author: "古川 緑波", SpeakerRuby: "のむらようじ", Speaker: "野村 洋二", FileName: "rd164.mp3", NewArrivalDate: "2020-11-01", Time: "9分25秒"},
	}, contents)
}
