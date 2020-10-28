# golang-unsplash
a simple golang api wrapper for search photos by word, using the api of https://unsplash.com/

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
# Example result

```console
foo@bar:~$ 
9 1070 119
https://images.unsplash.com/photo-1560582861-45078880e48e?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max&ixid=eyJhcHBfaWQiOjE3MDg1MX0
```
