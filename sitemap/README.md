# Sitemap Builder

Sitemap Builder makes use of link package from the last exercise and uses that to crawl a website and generate sitemap with all the links. Sitemap should follow the sitemap protocol.

Considerations:

- Only internal links. Links with different domain will be filtered out.
- Only links to other page (used to jump to other section of same page. e.g.`#links`).
- Refrain going into visited links to prevent cyclic condition.
- Parse concurrently for efficiency.

## Usage 

```
usage: sitemap [-h|--help] [-s|--site "<value>"] [-d|--depth <integer>]

               Generates sitemaps of given website.

Arguments:

  -h  --help   Print help information
  -s  --site   Site to be crawled. Default: http://127.0.0.1:6060/
  -d  --depth  Max depth to go starting from --site. Default: 2
```

**Output**:

```
$ go run main.go -s http://0.0.0.0:6060/ex1.html

<?xml version="1.0" encoding="UTF-8"?>

<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">

   <url>
      <loc>http://0.0.0.0:6060/ex1.html</loc>
   </url>

   <url>
      <loc>http://0.0.0.0:6060/other-page</loc>
   </url>

   <url>
      <loc>http://0.0.0.0:6060/second-page</loc>
   </url>

</urlset>
```

`ex1.html` is reused from the last lesson.

