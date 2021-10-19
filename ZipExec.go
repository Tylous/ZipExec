package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Tylous/ZipExec/Loader"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Writefile(outFile, result string) {
	cf, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer cf.Close()
	_, err = cf.Write([]byte(result))
	check(err)
}

type JScriptLoader struct {
	Variables map[string]string
}

type FlagOptions struct {
	outFile   string
	inputFile string
	sandbox   bool
}

func options() *FlagOptions {
	outFile := flag.String("O", "", "Name of output file (e.g. loader.js)")
	inputFile := flag.String("I", "", "Path to the file containing binary to zip.")
	sandbox := flag.Bool("sandbox", false, "Enables sandbox evasion using IsDomainedJoined calls.")
	flag.Parse()
	return &FlagOptions{outFile: *outFile, inputFile: *inputFile, sandbox: *sandbox}
}

func main() {
	fmt.Println(`
__________.__      ___________                     
\____    /|__|_____\_   _____/__  ___ ____   ____  
  /     / |  \____ \|    __)_\  \/  // __ \_/ ___\ 
 /     /_ |  |  |_> >        \>    <\  ___/\  \___ 
/_______ \|__|   __/_______  /__/\_ \\___  >\___  >
        \/   |__|          \/      \/    \/     \/ 
		(@Tyl0us)
		`)
	opt := options()
	if opt.inputFile == "" {
		log.Fatal("Error: Please provide a path to the binary")
	}
	if opt.outFile == "" {
		log.Fatal("Error: Please provide a name for the zip file")
	}

	loadercode := Loader.JScriptLoader_Buff(opt.inputFile, opt.outFile, opt.sandbox)
	Writefile(opt.outFile, loadercode)
	fmt.Println("[+] Loader Compiled")
}
