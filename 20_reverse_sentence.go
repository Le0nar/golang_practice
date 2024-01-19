package main

import "fmt"

func main()  {
 sentence := "snow dog sun"

 resersedSentence := reverseSentence(sentence)

 fmt.Printf("resersedSentence: %v\n", resersedSentence)
}

func reverseSentence(sentence string) string {
 var reversedSentence string
 var word string

 for index, v := range sentence {
  chararcter := string(v)
  isLastIndex := index == len(sentence) - 1

  if isLastIndex {
   word = word + chararcter + " "

   reversedSentence = word + reversedSentence
  }
 
  if chararcter == " " {
   reversedSentence = word + chararcter + reversedSentence
   word = ""
  } else {
   word = word + chararcter
  }
 }

 return reversedSentence
}