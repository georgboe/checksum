package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/clearsign"
	"golang.org/x/crypto/openpgp/packet"
	"golang.org/x/crypto/ssh/terminal"
)

func getChecksum(file string) []byte {
	hash := sha256.New()
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(hash, f)
	return hash.Sum(nil)
}

func checksumFiles(files []string) string {
	var output []string
	for _, arg := range files {

		checksum := getChecksum(arg)
		output = append(output, fmt.Sprintf("SHA256 (%s) = %s", arg, hex.EncodeToString(checksum)))
	}

	hashes := strings.Join(output, "\n")
	return hashes
}

func decryptPrivateKey(privateKey *packet.PrivateKey) {
	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Println()

	if err != nil {
		log.Fatalln("Could not read password from terminal.")
	}

	err = privateKey.Decrypt(password)

	if err != nil {
		log.Fatalln("Could not decrypt password from keyring.")
	}
}

func ensureFileExists(file string) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		log.Fatalln("File does not exist:", file)
	}
}

func getSigningKey() *openpgp.Entity {
	usr, err := user.Current()

	if err != nil {
		log.Fatalln("Could not get current user.")
	}

	keyringFile := usr.HomeDir + "/.gnupg/secring.gpg"

	ensureFileExists(keyringFile)
	keyring, err := os.Open(keyringFile)

	if err != nil {
		log.Fatalln("Could not read keyring file.")
	}
	defer keyring.Close()

	signers, err := openpgp.ReadKeyRing(keyring)

	if err != nil {
		log.Fatalln("Could not read GPG keyring.")
	}

	if len(signers) == 0 {
		log.Fatalln("No keys found in GPG keyring.")
	}

	signer := signers[0]
	return signer
}

func main() {
	log.SetFlags(0)
	program, files := os.Args[0], os.Args[1:]

	if len(files) == 0 {
		log.Fatal("usage: ", program, " [FILE]...")
	}

	hashes := checksumFiles(files)
	signer := getSigningKey()

	if signer.PrivateKey.Encrypted {
		decryptPrivateKey(signer.PrivateKey)
	}

	buf := bytes.NewBuffer(nil)
	w, err := clearsign.Encode(buf, signer.PrivateKey, nil)

	if err != nil {
		log.Fatalln("Could not create signature.")
	}

	w.Write([]byte(hashes))
	w.Close()

	fmt.Println(buf.String())
	ioutil.WriteFile("CHECKSUM", buf.Bytes(), 0644)
}
