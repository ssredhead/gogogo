package main

import "fmt"

func main() {
  var publisher, writer, artist, title string
  var year, pageNumber int8
  var grade float32

  title = "Mr. GoToSleep"
  writer = "Tracy Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "for grade", grade, "with", pageNumber, "pages")

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  // publisher is the same
  year = 2013
  pageNumber = 160
  grade = 9.0

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "of grade", grade, "with", pageNumber, "pages")

}
