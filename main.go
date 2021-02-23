package main
import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"os"
	"bufio"
	"flag"
	"time"
)
var banner string = `
                  ▄▄  ▄▄▄▄▄▄▄▄                                ▄▄        ▄▄
       ██        ██   ▀▀▀▀▀███                        ██       █▄        █▄
      ██        ██        ██▀    ▄████▄    ██▄████  ███████     █▄        █▄
     ██        ██       ▄██▀    ██▄▄▄▄██   ██▀        ██         █▄        █▄
    ▄█▀       ▄█▀      ▄██      ██▀▀▀▀▀▀   ██         ██          █▄        █
   ▄█▀       ▄█▀      ███▄▄▄▄▄  ▀██▄▄▄▄█   ██         ██▄▄▄        █▄        █▄
  ▄█▀       ▄█▀       ▀▀▀▀▀▀▀▀    ▀▀▀▀▀    ▀▀          ▀▀▀▀         █▄        █▄
                            Hash Cracker but on Golang
`
var start_time = time.Now()
func CheckError(err error){
	if err != nil{
		fmt.Println("Error while reading a wordlist")
		os.Exit(1)
	}
}
func CheckHash(hash,md5_string,password *string){
	if *hash == *md5_string{
		fmt.Printf("Found password:%s\nTime: %s",*password,GetTime())
		os.Exit(1)
	}
}
func StartCrack(hash,wordlist_path string){
	file, err := os.Open(wordlist_path)
	defer file.Close()
	CheckError(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		//if hash == GetMD5Hash(scanner.Text()){
		var password string = scanner.Text()
		var password_hash string = GetMD5Hash(scanner.Text()) 
		CheckHash(&hash,&password_hash,&password)
		if err := scanner.Err(); err != nil {
        	fmt.Printf("Error: %v", err)
        	printExit(wordlist_path)
        	
    	}
	}
	printExit(wordlist_path)
}
func printExit(worldlist_path string){
	fmt.Printf("Didn't found any password in %s\nTime:%s",worldlist_path,GetTime())
	os.Exit(1)	
}
func GetTime() time.Duration {
	var t = time.Now()
	return t.Sub(start_time)
}
func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
func main(){
	fmt.Printf(banner+"\nStarting cracking\n") 
	var hash = flag.String("hash", "", "Hash to crack")
	var worldlist_path = flag.String("wordlist","","Wordlist to use")
	flag.Parse()
	if (*hash != "") && (*worldlist_path != ""){
		StartCrack(*hash,*worldlist_path)
	}	
}