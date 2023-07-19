nostrkey
========
key pair utility for [Nostr](https://github.com/nostr-protocol/nostr).

### Environment
* Ubuntu 23.04 and later
* Go Language 1.20.5 and later

### Features
* Generating a key pair

### Requirements
* [nbd-wtf / go-nostr](https://github.com/nbd-wtf/go-nostr)

### Install nostrkey:
```bash
go install github.com/mitsugu/nostrkey@latest
```

### Usage
``` bash
nostrkey     : Print a hex type private key to standard output.  
nostrkey -n  : Reads a hex type private key from standard input, converts it to nsec format, and outputs it to standard output.  
nostrkey -p  : Reads a hex type private key from standard input, generates a hex type public key, and outputs it to standard output.  
nostrkey -np : Reads a hex type public key from standard input, generates an npub type public key, and outputs it to standard output.  
```

### Usage Example
``` bash
nostrkey | tee hsec.txt
cat hsec.txt | nostrkey -n | tee nsec.txt
cat hsec.txt | nostrkey -p | tee hpub.txt
cat hpub.txt | nostrkey -np | tee npub.txt
```

