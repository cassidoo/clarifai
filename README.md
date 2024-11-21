![image](https://github.com/user-attachments/assets/eed8dfe5-bfb6-4dfc-86b4-51e1f9a887f4)

# Clarifai Golang Library

Library for our v1 [API](https://clarifai.com/developer/guide-v1/).

[![GoDoc](https://godoc.org/github.com/Clarifai/clarifai-go?status.svg)](https://godoc.org/github.com/Clarifai/clarifai-go)

## Disclaimer

This API client only supports Clarifai v1 API. Stay tuned for the v2 support.

## Usage
`go get github.com/clarifai/clarifai-go`


```go
package main

import (
	"fmt"

	"github.com/clarifai/clarifai-go"
)

func main() {
	client := clarifai.NewClient("<client_id>", "<client_secret>")
	// Get the current status of things
	info, err := client.Info()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", info)
	}
	// Let's get some context about these images
	urls := []string{"https://samples.clarifai.com/metro-north.jpg", "https://samples.clarifai.com/puppy.jpeg"}
	// Give it to Clarifai to run their magic
	tag_data, err := client.Tag(clarifai.TagRequest{URLs: urls})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", tag_data) // See what we got!
	}

	// Tag a base64 image
	data = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII="
	tag_data, err = client.TagEncoded(clarifai.TagEncodedRequest{
		EncodedData: data
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", tag_data)
	}

	feedback_data, err := client.Feedback(clarifai.FeedbackForm{
		URLs:    urls,
		AddTags: []string{"cat", "animal"},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", feedback_data)
	}
}
```

## Testing
Run `go test`

##Special Thanks

Thanks so much to [Sam Couch](https://github.com/samuelcouch) for his amazing contribution to the Clarifai client libraries by starting this one and graciously handing it off to us. Follow him on Twitter [@SamuelCouch](http://twitter.com/SamuelCouch).
