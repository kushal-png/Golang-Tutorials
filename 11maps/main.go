package main
import "fmt"

func main(){

   fmt.Println("Maps")
   languages:= make(map[int]string)
   languages[1]="c++"
   languages[2]="python"
   languages[3]="ruby"
   languages[4]="java"
   
   fmt.Println(languages)

   delete(languages,2)
   fmt.Println(languages)
   

   // loops
   for key, value := range(languages){
	 fmt.Println(key,":", value)
   }
}