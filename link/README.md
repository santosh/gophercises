# HTML Link Parser

For each extracted link this program returns a data structure `Link` which includes both the `href`, as well as the text inside the link. Any HTML inside of the link is stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Links will be nested in different HTML elements, and it is very possible that you will have to deal with HTML similar to code below.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

In situations like these, we want to get output that looks roughly like:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```

## Usage

    go run main.go <filename>

Response would be something like this:

    [{/other-page A link to another page} {/second-page A link to second page}]
