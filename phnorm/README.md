# Phone Number Noarmalizer

This exercise deals with interacting with PostgreSQL using <https://github.com/lib/pq> as a driver.

## Getting Started

I have used Docker to run an instance of Postgres on my machine:

    docker run -d --name postgres -e POSTGRES_PASSWORD=Pass2020! -v postgres-data:/var/lib/postgresql/data -p 5432:5432 postgres:11-alpine

## Usage

For sake of demonstration, the table is dropped on every run. 

Seed data looks something like this:

```go
data := []string{
    "1234567890",
    "123 456 7891",
    "(123) 456 7892",
    "(123) 456-7893",
    "123-456-7894",
    "123-456-7890",
    "1234567892",
    "(123)456-7892",
}
```

```
$ go run main.go
Working on... {ID:1 Number:1234567890}
No changes required
Working on... {ID:2 Number:123 456 7891}
Updating or removing... 1234567891
Working on... {ID:3 Number:(123) 456 7892}
Updating or removing... 1234567892
Working on... {ID:4 Number:(123) 456-7893}
Updating or removing... 1234567893
Working on... {ID:5 Number:123-456-7894}
Updating or removing... 1234567894
Working on... {ID:6 Number:123-456-7890}
Updating or removing... 1234567890
Working on... {ID:7 Number:1234567892}
No changes required
Working on... {ID:8 Number:(123)456-7892}
Updating or removing... 1234567892
```

After normalization the table looks like this:

```sql
gophercises_phone=# SELECT * from phone_numbers;
 id |   value    
----+------------
  1 | 1234567890
  7 | 1234567892
  2 | 1234567891
  4 | 1234567893
  5 | 1234567894
(5 rows)
```
