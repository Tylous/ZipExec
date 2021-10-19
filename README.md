# ZipExec
ZipExec is a Proof-of-Concept (POC) tool to wrap binary-based tools into a password-protected zip file. This zip file is then base64 encoded into a string that is rebuilt on disk. This encoded string is then loaded into a JScript file that when executed, would rebuild the password-protected zip file on disk and execute it. This is done programmatically by using COM objects to access the GUI-based functions in Windows via the generated JScript loader, executing the loader inside the password-protected zip without having to unzip it first. By password protecting the zip file, it protects the binary from EDRs and disk-based or anti-malware scanning mechanisms.



<p align="center"><img src="Screenshots/demo.gif" border="2px solid #555" />


## Installation

The first step as always is to clone the repo. Before you compile ZipExec you'll need to install the dependencies. To install them, run following commands:
```
go get github.com/yeka/zip
```


Then build it

```
go build ZipExec.go
``` 

## Help
```
./ZipExec -h

__________.__      ___________                     
\____    /|__|_____\_   _____/__  ___ ____   ____  
  /     / |  \____ \|    __)_\  \/  // __ \_/ ___\ 
 /     /_ |  |  |_> >        \>    <\  ___/\  \___ 
/_______ \|__|   __/_______  /__/\_ \\___  >\___  >
        \/   |__|          \/      \/    \/     \/ 
                (@Tyl0us)

Usage of ./ZipExec:
  -I string
        Path to the file containing binary to zip.
  -O string
        Name of output file (e.g. loader.js)
  -sandbox
        Enables sandbox evasion using IsDomainedJoined calls.
```



