package main

import (
	"os"
	"bufio"
	"time"
	"fmt"
	"log"
	"errors"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

/*,
main
	nostrkey
	nostrakey -n < hsec
	nostrakey -p < hsec
	nostrakey -np < hsec
*/


func main() {
	cn := make(chan string, 1)
		go func() {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		cn <- sc.Text()
	}()

	switch (len(os.Args)) {
	case 1:
		sk := nostr.GeneratePrivateKey()
		fmt.Println(sk)
	case 2:
		timer := time.NewTimer(time.Second * 5)
		defer timer.Stop()
		switch (os.Args[1]) {
		case "-n":
			select {
			case text := <-cn:
				nsec, err := nip19.EncodePrivateKey(text)
				if err!=nil {
					log.Fatal(err)
				}
				fmt.Println(nsec)
			case <-timer.C:
				log.Fatal(errors.New("Time out input from standerd input"))
			}
		case "-p":
			select {
			case text := <-cn:
				pk, err := nostr.GetPublicKey(text)
				if err!=nil{
					log.Fatal(err)
				}
				fmt.Println(pk)
			case <-timer.C:
				log.Fatal(errors.New("Time out input from standerd input"))
			}
		case "-np":
			select {
			case text := <-cn:
				npub, err := nip19.EncodePublicKey(text)
				if err!=nil{
					log.Fatal(err)
				}
				fmt.Println(npub)
			case <-timer.C:
				log.Fatal(errors.New("Time out input from standerd input"))
			}
		case "-h":
			dispHelp()
		case "--help":
			dispHelp()
		default:
			fmt.Println("Invalid option.")
			fmt.Println()
			dispHelp()
		}
	}
}

//

func dispHelp() {
	fmt.Println("nostrkey")
	fmt.Println("    print the key for Nostr to standard output.")
	fmt.Println("")
	fmt.Println("Usage :")
	fmt.Println("        nostrkey : Print a hex type private key to standard output.")
	fmt.Println("        nostrkey -n : Reads a hex type private key from standard input,")
	fmt.Println("                      converts it to nsec format,")
	fmt.Println("                      and outputs it to standard output.")
	fmt.Println("        nostrkey -p : Reads a hex type private key from standard input,")
	fmt.Println("                      generates a hex type public key,")
	fmt.Println("                      and outputs it to standard output.")
	fmt.Println("        nostrkey -np : Reads a hex type public key from standard input,")
	fmt.Println("                       generates an npub type public key,")
	fmt.Println("                       and outputs it to standard output.")
}

//

