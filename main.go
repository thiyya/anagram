package main

import (
	"anagram/processor"
	"anagram/utils"
	"fmt"
	"log"
)

func main() {
	filePath, err := utils.HandleFlag()
	if err != nil {
		log.Fatalln(err)

		return
	}

	content, err := utils.ReadFileContent(filePath)
	if err != nil {
		log.Fatalln(err)

		return
	}

	result := processor.Processor(content)
	for _, s := range result {
		fmt.Println(s)
	}
}
