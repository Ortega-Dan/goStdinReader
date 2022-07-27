# golang-stdin


Usage for STDIN reading:

```golang
import "github.com/Ortega-Dan/golang-stdin/stdin"

func myFunction(){

  text := stdin.ReadLine()
  
  // ...
  // do your thing with the text string
  // ...
}
```
(`stdin.ReadLineNoTrim()` also exists if empty space around string wants to be preserved)
___
Usage for file reading:

```golang
import "github.com/Ortega-Dan/golang-stdin/stdin"

func readFile() {

	// starting reader
	reada, err := stdin.NewLineByLineFileReader("test.txt")
	if err != nil {
		fmt.Println("ERROR OPENING FILE !!!")
	}

	// variables to work with
	var readingErr error
	var str string

	// line reading loop
	for readingErr == nil {
		str, readingErr = reada()

		// fmt.Println(str) // do your stuff
	}

	// identify if different error than end of file
	if readingErr != io.EOF {
		fmt.Println("REAL READING ERROR !!!")
	}

	fmt.Println("**** Went well ******")
}
```