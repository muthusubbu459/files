package main
import(
	"fmt"
	"io/ioutil"
	"os"
	"log"
//	"io"
//	"go/token"
	//"crypto/x509"
	"io"
)
func write(f[] byte){
	a,err:= os.Create("C:/users/muthusubbu459/GoWorkspace/src/maximum/new.txt")
	if err!=nil{
		log.Fatal(err)
	}
	a.Write(f)
	// a.Close()
	fmt.Println("the data is written into newfile")
	l:=make([]byte,5000)
	k,err:=a.Read(l)
	if err!=nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}


	fmt.Println("the new file is written")
	fmt.Printf( "read %d bytes: %q\n",k,l[:k])
	//fmt.Println(string(k),string(l))
	m,err:=ioutil.ReadFile("C:/users/muthusubbu459/GoWorkspace/src/maximum/defer.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(m))

}
func main() {
	f, err := os.Open("C:/users/muthusubbu459/GoWorkspace/src/maximum/defer.txt")
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println("the data in file is")
	//	fmt.Println(string(f))
	//	write(f)
	l := make([]byte, 5000)
	k, err := f.Read(l)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q", k, l[:k])
	write(l)
}







