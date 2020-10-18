# Adding Flags

Adding a flag to **yctf** is easy. While the project is written using the [gin-gonic/gin](https://github.com/gin-gonic/gin) framework, you do not need extensive knowledge using the framework to add a flag. That said, some knowledge is beneficial for you to be able to create unique flags in your handlers. 

Each flag will require five additions to the code base.

## Adding Metadatda

Each flag has an entry in the [flags.json](../flags.json) file in the following format.

```json
[
    ...
    {
       "ID": 0,
       "Flag": "yctf-5f39410913126cea50a3e6fcdb7312628a869da376985fdb7ef7d0a41254a0e3", 
       "Template": "flag0.html"
    }
    ...
]
```

- `ID` is the flag identifier. This must be unique.
- `Flag` is a sha-2 (with a digest of 256) string with the `yctf-` prefix that is to be found by the participant. This must be unique. The sha-2 string is the hash of `yctf-ID`.
- `Template` refers to a Go HTML Template by filename. These are stored in the [templates](../templates/)directory. This should closely mirror the format of existing templates (e.g. `flag{ID}.html`).

You will need to add an entry to [flags.json](../flags.json) for each flag that you want to introduced to `yctf`.a

## Adding Your Handler

Each flag has a handler function that handles HTTP requests to the flag URL (i.e. http://localhost:8080/flag{ID}). Each handler function is in a separate file as a part of the `main` package (e.g. [handlers.flag0.go](../handlers.flag0.go))

The following is flag0, but with additional comments inline to help you understsand how things are connected. 

```go
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func getFlag99(c *gin.Context) {
    // Get the flag details. These are pulled from the flags.json file.
    f, err := getFlag(99)
    // If we were unable to get the flag from the json file, we return an error
	if err != nil {
		log.Fatal(err)
    }
    
    // If we were successful in getting the flag from flags.json, we set our page title
    // for when we render.
	title := fmt.Sprintf("flag%v", f.ID)

    // Finally, we render the page to the user. The template used to render the page is
    // pulled from flags.json and is read from the filesystem. The payload contains
    // the flag string and is also pulled from flags.json.
	render(c, gin.H{
		"payload": f.Flag,
		"title":   title}, f.Template)
}
```

You may use other approaches to rendering so long as you handle rendering of your flag's clue (which is always HTML). 


## Adding Your Template

The template is the next piece you need. It should live in the [templates/](../templates/) directory and should take on the same name you defined for your flag in [flags.json](../flags.json). 

Templates are written using Golang's [html/template](https://golang.org/pkg/html/template/) package and are generally expected to container the **Header** and **Footer** elements `yctf` provides.

```go
<!--Embed the header.html template at this location-->
{{ template "header.html" .}}

    // your content goes here

<!--Embed the footer.html template at this location-->
{{ template "footer.html" .}}
```

## Adding Your Route

Once you have the core components of your flag in place, you need to tell the router how to handle requests for your new flag. This is done in the [routes.go](../routes.go) file in function `initializeRoutes()`. Add a new entry mapping your flag's function to your URI `yctf` will automatically handle requests to your flag for you. Here's an example for a sample "flag99".

```go
package main

func initializeRoutes() {

    //... all other routes ...

	// Handle flag99
    router.GET("/flag99", getFlag99)

    //... all other routes ...
}
```

## Adding Your Tests

This section is coming soon.