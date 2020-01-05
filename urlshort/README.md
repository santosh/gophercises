# urlshort

URL shortner with data backstore supporting YAML, JSON, BoltDB.

Attention to the copy pasters. The examples are of course not production ready and needs changes. This was an example how we can interact with various data backstores in golang. There was a lose design document therefore implementation might not as you expected. Notable irritation may include but not limited to: 

- Commnad line argument does not makes sense. It will look for all the yaml, json and db file regardless of what you're actually running.
- Design at the moment is not unit-testable.

I'm trying not to fix them and move on to next exercises. Take them as TODO.
