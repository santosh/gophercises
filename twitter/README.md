# twitter retweet contest

In this exercise, given a tweet ID, we pick n number of retweeters of a tweet.

Things learned are OAuth2 authentication with _client credential_ grant type.

## Usage 

1. We need to get the credential of the app. This app is the entity on behalf of which we'll make a request to Twitter's API. `consumer_key` and `consumer_secret`. You'll get this after registering your application with Twitter at <https://apps.twitter.com>

2. Run the program with `go run main.go -tweet <tweet ID> -winners <# of winners to choose>`.

Below is full list of flags supported.

```
$ go run main.go --help
Usage of /tmp/go-build651276016/b001/exe/main:
  -key string
        credential files (default ".keys.json")
  -tweet string
        The ID of the tweet you want retweeters of. (default "991053593250758658")
  -users string
        file where users have retweeted the tweet are stored. (default "users.csv")
  -winners int
        The number of winners to pick for the contest.
```

A sample run would result something like below:

```
$ go run main.go -tweet 1293305328696688641 -winners 2
The winners are:
        anhpham64
        jgmc3012
```
