package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/santosh/gophercises/phnorm/db"
)

func init() {
	must(godotenv.Load())
}

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	dbname   = "gophercises_phone"
)

func main() {
	password := os.Getenv("PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, username, password)
	must(db.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(db.Migrate("postgres", psqlInfo))

	db, err := db.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)

	phones, err := db.AllPhones()
	must(err)

	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

// func normalize(phone string) string {
// 	var buf bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}

// 	return buf.String()
// }
