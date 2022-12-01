package processor

import (
	"anagram/entity"
	"anagram/utils"
	"log"
	"sync"
)

// Processor Initializes a safe map and creates workers
func Processor(contentList []entity.Content) []string {
	var waitGroup sync.WaitGroup
	safeMap := New()
	wgCounter := 0
	for _, content := range contentList {
		waitGroup.Add(1)
		go process(content.Words, &waitGroup, safeMap)
		wgCounter = wgCounter + 1
	}
	log.Println("Counter of working group: ", wgCounter)
	waitGroup.Wait()
	result := make([]string, 0)
	for key, _ := range safeMap.safeMap {
		res := safeMap.get(key)
		if len(res) > 0 {
			result = append(result, res)
		}
	}
	safeMap.flush()

	return result
}

func process(words []string, waitGroup *sync.WaitGroup, safeMap *SafeMap) error {
	defer waitGroup.Done()
	for _, word := range words {
		key := utils.SortStringByCharacter(word)
		safeMap.add(key, word)
	}

	return nil
}
