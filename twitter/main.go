package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/santosh/gophercises/twitter/twitter"
)

var (
	keyFile    string
	usersFile  string
	tweetID    string
	numWinners int
)

func init() {
	flag.StringVar(&keyFile, "key", ".keys.json", "credential files")
	flag.StringVar(&usersFile, "users", "users.csv", "file where users have retweeted the tweet are stored.")
	flag.StringVar(&tweetID, "tweet", "991053593250758658", "The ID of the tweet you want retweeters of.")
	flag.IntVar(&numWinners, "winners", 0, "The number of winners to pick for the contest.")
	flag.Parse()
}

func main() {
	key, secret, err := keys(keyFile)
	if err != nil {
		panic(err)
	}
	client, err := twitter.New(key, secret)
	if err != nil {
		panic(err)
	}
	newUsernames, err := client.Retweeters(tweetID)
	if err != nil {
		panic(err)
	}
	existingUsername := existing(usersFile)
	allUsernames := merge(newUsernames, existingUsername)
	err = writeUsers(usersFile, allUsernames)
	if err != nil {
		panic(err)
	}
	if numWinners == 0 {
		return
	}
	existingUsername = existing(usersFile)
	winners := pickWinners(existingUsername, numWinners)
	fmt.Println("The winners are:")
	for _, username := range winners {
		fmt.Printf("\t%s\n", username)
	}
}

func keys(keyFile string) (key, secret string, err error) {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}
	f, err := os.Open(keyFile)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	return keys.Key, keys.Secret, nil
}

func existing(usersFile string) []string {
	f, err := os.Open(usersFile)
	if err != nil {
		return []string{}
	}

	defer f.Close()
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	users := make([]string, 0, len(lines))
	for _, line := range lines {
		users = append(users, line[0])
	}
	return users
}

func merge(a, b []string) []string {
	uniq := make(map[string]struct{}, 0)
	for _, user := range a {
		uniq[user] = struct{}{}
	}
	for _, user := range b {
		uniq[user] = struct{}{}
	}
	ret := make([]string, 0, len(uniq))
	for user := range uniq {
		ret = append(ret, user)
	}
	return ret
}

func writeUsers(usersFile string, users []string) error {
	f, err := os.OpenFile(usersFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for _, username := range users {
		if err := w.Write([]string{username}); err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func pickWinners(users []string, numWinners int) []string {
	existingUsername := existing(usersFile)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(existingUsername))
	winners := perm[:numWinners]
	var ret []string
	for _, idx := range winners {
		ret = append(ret, users[idx])
	}
	return ret
}
