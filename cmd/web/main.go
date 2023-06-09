package main

/*
Dependencies:
	github.com/jackc/pgconn
	github.com/jackc/pgx/v4
	github.com/alexedwards/scs/v2
	github.com/alexedwards/scs/redisstore
	github.com/go-chi/chi/v5
*/
import (
	"database/sql"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// connect to the database
	initDB()
	// create sessions
	initSession()
	// create loggers

	// create channels

	// create waitgroup

	// set up the application config

	// set up mail

	// listen for web connections
}

// initDB connects to Postgres and returns a pool of connections
func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return conn
}

// connectToDB tries to connect to postgres, and backs off until a connection
// is made, or we have not connected after 10 tries
func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Print("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++

		continue
	}
}

// openDB opens a connection to Postgres, using a DSN read
// from the environment variable DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// initSession sets up a session, using Redis for session store
func initSession() *scs.SessionManager {
	// set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

// initRedis returns a pool of connections to Redis using the
// environment variable REDIS
func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}
