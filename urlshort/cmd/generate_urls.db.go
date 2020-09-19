package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// Change this to 'main' while generating bolt.db files. This also means you'll heave
// Change back to something else e.g. 'mmain' when running the main.
func main() {
	var redirects = []byte("redirects")
	db, err := bolt.Open("urls.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pathsToUrls := map[string]string{
		"/so": "https://stackoverflow.com/u/939986/",
		"/li": "https://www.linkedin.com/in/sntshk",
	}

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(redirects)
		if err != nil {
			return err
		}

		for key, value := range pathsToUrls {
			err = bucket.Put([]byte(key), []byte(value))
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(redirects)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", redirects)
		}

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
