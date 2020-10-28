# golang-unsplash
a simple api wrapper for search photos by word

# Usage
```golang
package main

import (
	"fmt"
	"github.com/alonsopf/golang-unsplash"
)

func main() {
	unsplash.UNSPLASH_ACCESS_KEY = "YOUR_KEY"
	Photos, err, total, totalPages := unsplash.SearchPhotosByWord("medical",1,9)
    if err != nil {
    	fmt.Println(err)
    	return
    }
    fmt.Println(len(*Photos), total, totalPages)
    if len(*Photos) > 0 {
    	fmt.Println((*Photos)[0].Url)	
    }   
}
```
