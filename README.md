# gophercises

Go exercises from gophercises.com offers learning by doing.

Following are the exercises covered.

## Quiz Game

This application reads from a CSV file (default: problems.csv). Each row is have a question in the first column and answer in the second. There is a default 30 seconds timer after which quiz will end. Results shows how many questions were asked and how many of them were correct.

## URL Shortner

A BoldDB backed URL Shortner. Fallbacks to YAML and JSON seed files.

## Choose Your Own Adventure

Web simulation of the famous book Choose Your Own Adventure. The story is fed via a JSON input.

## Link Parser

Fetches `href` and `text` of an `<a>` tag to a custom `Link` data structure.

## Sitemap Builder

Built on top of Link Parser is Sitemap Builder. It crawls and finds all the links n devel deep and generate sitemap with all the links. Uses BFS algorithm.

## Hacker Rank Problem

Solves Camel Problem and Cipler Problem from HackerRank. The idea was to use services like HackerRank, LeetCode and Project Euler to work on algorithms and data structures.

## Task Manager

Task manager is separately uploaded to <https://gitlab.com/sntshk/task>

## Phone Number Normalizer

This lesson is designed in a way to teach interaction with SQL databases in Go.
Uses Postgres with <https://github.com/lib/pq>.

## Deck of Cards

The `deck` package, which will be used in future card game exercises. Implements Suit, Rank and Card types.

## Blackjack

Using the `deck` module, implements a simple version of the Blackjack game.

## File Renaming Tool

A tool used to rename files with a common pattern. Eg we might want to take many files with names like “Dog (1 of 100).jpg”, “Dog (2 of 100).jpg”, … and rename them to “Dog_001.jpg”, “Dog_002.jpg”, …

## Quite Hacker News

Take <https://github.com/tomspeak/quiet-hacker-news>, and make it fast by adding caching and concurrency.

## Recover Middleware

HTTP middleware that recovers from panic in an application.

## Recover Middleware w/ Source Code

Expand upon the recover middleware and add links to source code along with syntax highlighting of the source code in order to make a useful development tool.

## Twitter Highlight Contest CLI

Create a CLI to help run a Twitter contest where users retweet a tweet for entry, and after some time you pick one or more of the users who retweeted as the winner.

