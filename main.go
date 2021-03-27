package main
import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
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
func CheckHash(hash,new_hash_string,password *string){
	if *hash == *new_hash_string{
		fmt.Printf("Found password:%s\nTime: %s\n",*password,GetTime())
		os.Exit(1)
	}
}
func StartCrack(hash,wordlist_path,encryption_type *string){
	
	file, err := os.Open(*wordlist_path)
	defer file.Close()
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		//if hash == GetMD5Hash(scanner.Text()){
		var password string =scanner.Text()
		if *encryption_type == "md5" {
			CheckHash(hash,GetMD5Hash(scanner.Text()),&password)
		} else if *encryption_type == "sha1" {
			fmt.Printf("Using sha1")
			CheckHash(hash,GetSHA1Hash(scanner.Text()),&password) 
		} else if *encryption_type == "sha256" {
			CheckHash(hash,GetSHA256Hash(scanner.Text()),&password)
		} else if *encryption_type == "sha512" {
			CheckHash(hash,GetSHA512Hash(scanner.Text()),&password)
		}
		if err := scanner.Err(); err != nil {
        		fmt.Printf("Error: %v", err)
        		printExit(*wordlist_path)
   		}
	}
	printExit(*wordlist_path)
}
func printExit(worldlist_path string){
	fmt.Printf("Didn't found any password in %s\nTime:%s",worldlist_path,GetTime())
	os.Exit(1)	
}
func GetTime() time.Duration {
	var t = time.Now()
	return t.Sub(start_time)
}
func GetMD5Hash(text string) *string {
    var output string = fmt.Sprintf("%x",md5.Sum([]byte(text)))
    return &output
}
func GetSHA1Hash(text string) *string {
    var output string = fmt.Sprintf("%x",sha1.Sum([]byte(text)))
    return &output
}
func GetSHA256Hash(text string) *string {
    var output string = fmt.Sprintf("%x",sha256.Sum256([]byte(text)))
    return &output
}
func GetSHA512Hash(text string) *string {
    var output string = fmt.Sprintf("%x",sha512.Sum512([]byte(text)))
    return &output
}
func areArgumentsCorrect(args [3]string) bool{
	for _, argument := range args{
		if (argument != ""){
			continue
		} else{
			return false
		}
	}
	return true
}
func main(){
	fmt.Printf(banner+"\nStarting cracking\n") 
	var hash = flag.String("hash", "", "Hash to crack")
	var worldlist_path = flag.String("wordlist","","Wordlist to use")
	var encryption_type =flag.String("type","","Type of hash: md5, sha1, sha256, sha512")
	flag.Parse()
	arguments := [3] string {*hash,*worldlist_path,*encryption_type}
	
	if (areArgumentsCorrect(arguments)){
		StartCrack(hash,worldlist_path,encryption_type)	

	} else {
		fmt.Printf("Not all arguments were set\n")
		os.Exit(1)
	}
}
