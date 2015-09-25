package main
import (
	"fmt"
	"os"
	"log"
)
func main(){
file,err:= os.Open("C:/users/muthusubbu459/GoWorkspace/src/maximum/maximum.go")
 if err != nil {
		log.Fatal(err)
	}
	data := make([]byte,500)
	count,err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf( "read %d bytes: %q\n",count,data[:count])
}
