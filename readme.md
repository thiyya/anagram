## Anagram matching
It is a program that takes as argument the path to a file containing one word per line, groups the words that are anagrams to each other, and writes to the standard output each of these groups.
The groups should be separated by new lines, and the words inside each group by commas.

* It is assumed that the words in the file are ordered by size. In this way, words are first read according to their size and kept in separate lists.
* A separate goroutine is created for each list.
* Words are saved in a safe map, normalized versions of the words are saved as keys and themselves as values
* Each record in the resulting safe map actually contains anagram words. 
* If there is only one word in it, it is understood that the related word doesn't have any anagram pair.

## How to run:

```bash
go install
anagram data/example.txt
```

Output:
```text
abc,bac,cba
unf,fun
```
