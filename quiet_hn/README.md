# quiet_hn

Quiet HackerNews is a command-line tool to browse HackerNews. It is quiet in the sense that it only shows only items that are links to other sites (`itemType == 'story'`). In addition to that, it also keeps away the comments. Topics which are _discussion_ are also excluded. 

This exercise is based on [tomspeak's quiet-hacker-news](https://github.com/tomspeak/quiet-hacker-news), and on top of that is adds concurrency and caching. 

The tool only shows 30 stories (overridable) and in same order as the original. See source code to learn how different gorountines are synced in same order where not all 30 item are story. The tool _has_ to fetch 30 stories, no matter what number of stories comes in first cycle.

## Usage 

```
$ go run main.go --help
Usage of /tmp/go-build089732964/b001/exe/main:
  -num_stories int
        the number of top stories to display (default 30)
  -port int
        the port to start the web server on (default 3000)
```

Simply run go run main.go and head over to <http://127.0.0.1:3000>


## Screenshot

![Quiet HackerNews Custom Landing Page](./screenshot.png "Custom Quiet HackerNews")
