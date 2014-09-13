This is not ready for production.  This is the result of staying up until 
2am attempting to learn Go for a command-line program I'm working on. 
This was actually my first Go package and I've learned a bit since this was
born.  Use at your own risk.  It does depend on a few other packages I've 
thrown together. 


## Usage:

```
package main

import (
  "fmt"
  "github.com/tlorens/pullbarz"
  "github.com/tlorens/ansiout"
)


func main() {
	ansiout.ClearScr()
	ansiout.Wait(5)

	ansiout.PrintFile("header.ans")

	var options = map[int]string {
		1: "Reply",
		2: "Delete",
		3: "Forward",
		4: "Quit",
		5: "eXit",
	}

	pullbarz.Init(options)
	ch := pullbarz.LightBar(0, 0)

	fmt.Print(ch)
}

```


### Here's what it does
![BBS Menu](https://dl.dropboxusercontent.com/u/2934311/stupid-bbs.png)

### DOS ANSI color codes from TheDraw
![Color Codes](https://dl.dropboxusercontent.com/u/2934311/doscolors.png)
