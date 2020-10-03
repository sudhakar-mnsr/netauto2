package dbutils
import "log"
import "database/sql"
func Initialize(dbDriver *sql.DB) {
   statement, driverError := dbDriver.Prepare(train)
   if driverError != nil {
      log.Println(driverError)
   }

