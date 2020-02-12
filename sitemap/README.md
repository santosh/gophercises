# Sitemap Builder

Crawl a website and generate sitemap with all the links. Sitemap should follow the sitemap protocol.

Considerations:

- Only internal links. Links with different domain will be filtered out.
- Only links to other page (used to jump to other section of same page. e.g.`#links`).
- Refrain going into visited links to prevent cyclic condition.
- Parse concurrently for efficiency.


## TODO

- Parse concurrently for efficiency.
