package db
import (

   "log"
   "gorm.io/driver/postgres"
   "gorm.io/gorm"
   "github.com/pramishkarkee/gql-test-struct/graph/model"
)
func Connect() *gorm.DB {
	// dbURL := os.Getenv("DB_URL")
	databasename := "postgres"
	// database := "postgres"
	databasepassword := "root1234"
	dbURL := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
        log.Fatalln(err)
    }
//    db := pg.Connect(opt)
//    if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
//        panic("PostgreSQL is down")
//    }
   db.AutoMigrate(&model.Movie{})

    return db
}