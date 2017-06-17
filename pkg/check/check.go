/***********
*
*
* project s3 Upload
*
* @package     main
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*
 */

package check

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

//
//
//
func Exists(fileName string) bool {

	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
}

//
// FileUpload, Bucket, stringAcl
//
func GenArgs() (string, string, string, int, string) {

	//
	//
	//
	mapCommand := map[int]string{0: "put", 1: "get", 2: "bucket", 3: "acl", 4: "crypt", 5: "decrypt", 6: "version", 7: "help", 8: "h", 9: "v"}

	//
	//
	// white := color.New(color.FgWhite)
	// boldWhite := white.Add(color.Bold)

	//
	//
	//
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	//
	//
	//
	yellow := color.New(color.FgYellow)
	boldYellow := yellow.Add(color.Bold)

	//
	//
	//
	var stringCmd string

	//
	//
	//
	var FileUpload string

	//
	//
	//
	var FileGetPut string

	//
	//
	//
	var Bucket string

	//
	//
	//
	var existCmd int

	//
	//
	//
	var lenArgs int

	//
	//
	//
	var cryptInt int

	//
	//
	//
	cryptInt = 0

	//
	//
	//
	existCmd = 0

	//
	//
	//
	stringAcl := ""

	//
	//
	//
	lenArgs = len(os.Args)

	//
	// == 1 default cmd
	//
	if lenArgs == 1 {

		PrintDefaults()
		os.Exit(0)
	}

	//
	//
	//
	arrayParam := map[int]string{}

	//
	//
	//
	var cmdFor string

	//
	//
	//
	exitInGetPut := 0

	//
	//
	//
	for j, val := range os.Args {

		//
		//
		//
		arrayParam[j] = val

		i := strings.Index(val, "-")

		if i == 0 {

			// fmt.Println(":::", val)

			stringCmd = strings.Trim(val, "-")
			stringCmd = strings.TrimSpace(stringCmd)
			stringCmd = strings.ToLower(stringCmd)

			exitInMap := 0

			if stringCmd == "get" {

				exitInGetPut++

			} else if stringCmd == "put" {

				exitInGetPut++
			}

			for _, val2 := range mapCommand {

				if val2 == stringCmd {

					exitInMap = 1
					break

				} else {

					cmdFor = val
				}
			}

			if exitInMap == 0 {

				boldRed.Println("\nCommand [" + cmdFor + "] does not exist!")
				PrintDefaults()
				os.Exit(0)
			}

			//
			// Exists in the command map
			//
		}
	}

	if exitInGetPut >= 2 {

		boldRed.Println("\nCommand get and put not allowed!")
		PrintDefaults()
		os.Exit(0)
	}

	for x := range arrayParam {

		stringCmd = strings.Trim(arrayParam[x], " ")
		stringCmd = strings.TrimSpace(stringCmd)
		stringCmd = strings.ToLower(stringCmd)

		switch stringCmd {

		case "--put":

			FileUpload = validPut(x, arrayParam)
			FileGetPut = "put"
			existCmd++

		case "-put":

			FileUpload = validPut(x, arrayParam)
			FileGetPut = "put"
			existCmd++

		case "--get":

			FileGetPut = "get"
			FileUpload = validGet(x, arrayParam)
			existCmd++

		case "-get":

			FileGetPut = "get"
			FileUpload = validGet(x, arrayParam)
			existCmd++

		case "--bucket":

			Bucket = validBucket(x, arrayParam)
			existCmd++

		case "-bucket":

			Bucket = validBucket(x, arrayParam)
			existCmd++

		case "--crypt":

			_, ok := arrayParam[x+1]

			if ok {

				boldRed.Println("\nThere is no value for this parameter\n")
				os.Exit(0)

			}

			cryptInt += 1
			existCmd++

		case "-crypt":

			_, ok := arrayParam[x+1]

			if ok {

				boldRed.Println("\nThere is no value for this parameter\n")
				os.Exit(0)

			}

			cryptInt += 1
			existCmd++

		case "--acl":

			stringAcl = validAcl(x, arrayParam)
			existCmd++

		case "-acl":

			stringAcl = validAcl(x, arrayParam)
			existCmd++

		case "--decrypt":

			_, ok := arrayParam[x+1]

			if ok {

				boldRed.Println("\nThere is no value for this parameter\n")
				os.Exit(0)

			}

			cryptInt += 1
			existCmd++

		case "-decrypt":

			_, ok := arrayParam[x+1]

			if ok {

				boldRed.Println("\nThere is no value for this parameter\n")
				os.Exit(0)

			}

			cryptInt += 1
			existCmd++

		case "--help":

			PrintDefaults()
			os.Exit(0)

		case "--h":

			PrintDefaults()
			os.Exit(0)

		case "-help":

			PrintDefaults()
			os.Exit(0)

		case "-h":

			PrintDefaults()
			os.Exit(0)

		case "--version":

			boldYellow.Println("v.1.0")
			os.Exit(0)

		case "--v":

			boldYellow.Println("v.1.0")
			os.Exit(0)

		case "-version":

			boldYellow.Println("v.1.0")
			os.Exit(0)

		case "-v":

			boldYellow.Println("v.1.0")
			os.Exit(0)
		}
	}

	//
	//
	//
	if existCmd == 0 || existCmd == 1 {

		PrintDefaults()
		os.Exit(0)
	}

	return FileUpload, Bucket, stringAcl, cryptInt, FileGetPut
}

//
//
//
func validAcl(x int, arrayParam map[int]string) string {

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	var stringCmd2 string
	var stringAclTmp string
	var stringAcl string

	stringCmd2 = strings.Trim(arrayParam[x+1], "-")
	stringCmd2 = strings.TrimSpace(stringCmd2)
	stringCmd2 = strings.ToLower(stringCmd2)

	stringAclTmp = fmt.Sprintf("%s", stringCmd2)

	if stringAclTmp == "read" {

		stringAcl = "read"

	} else if stringAclTmp == "write" {

		stringAcl = "write"

	} else if stringAclTmp == "all" {

		stringAcl = "all"

	} else {

		boldRed.Println("\nAcl does not exist! Try red | write | all\n")
		os.Exit(0)
	}

	return stringAcl
}

//
//
//
func validPut(x int, arrayParam map[int]string) string {

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	var stringCmd2 string
	var FileUpload string

	value, ok := arrayParam[x+1]

	if ok {

		i := strings.Index(value, "-")

		if i == 0 {

			// exist
			boldRed.Println("\nMissing file as parameter ex: --put <file>\n")
			os.Exit(0)
		}

		stringCmd2 = strings.Trim(value, "-")
		stringCmd2 = strings.TrimSpace(stringCmd2)
		FileUpload = fmt.Sprintf("%s", stringCmd2)

	} else {

		boldRed.Println("\nMissing file as parameter ex: --put <file>\n")
		os.Exit(0)
	}

	return FileUpload
}

//
//
//
func validGet(x int, arrayParam map[int]string) string {

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	var stringCmd2 string
	var FileGet string

	value, ok := arrayParam[x+1]

	if ok {

		i := strings.Index(value, "-")

		if i == 0 {

			// exist
			boldRed.Println("\nMissing file as parameter ex: --get <file>\n")
			os.Exit(0)
		}

		stringCmd2 = strings.Trim(value, "-")
		stringCmd2 = strings.TrimSpace(stringCmd2)
		FileGet = fmt.Sprintf("%s", stringCmd2)

	} else {

		boldRed.Println("\nMissing file as parameter ex: --get <file>\n")
		os.Exit(0)
	}

	return FileGet
}

//
//
//
func validBucket(x int, arrayParam map[int]string) string {

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	var stringCmd2 string
	var Bucket string

	value, ok := arrayParam[x+1]

	if ok {

		i := strings.Index(value, "-")

		if i == 0 {

			// exist
			boldRed.Println("\nMissing file as parameter ex: --bucket <file>\n")
			os.Exit(0)
		}

		stringCmd2 = strings.Trim(value, "-")
		stringCmd2 = strings.TrimSpace(stringCmd2)
		Bucket = fmt.Sprintf("%s", stringCmd2)

	} else {

		boldRed.Println("\nMissing file as parameter ex: --bucket <name>\n")
		os.Exit(0)
	}

	return Bucket
}

//
//
//
func PrintDefaults() {

	var help string

	help = `	
  Use: 
   s3goamz [OPTION]...
   or: s3goamz --put file.pdf --bucket name-bucket [options]
   or: s3goamz --put file.pdf --bucket name-bucket --acl read [options]
   or: s3goamz --put file.pdf --bucket name-bucket --acl read --crypt
   or: s3goamz --get file.pdf --bucket name-bucket

   Put and bucket arguments are required.
   -put,     --put      <file>    The file and its respective path
   -get,     --get      <file>    The file and its respective path
   -bucket,  --bucket   <name>    Bucket name s3
   -acl,     --acl      <options> read, write, all
   -crypt,   --crypt    has no parameter
   -help,    --help     -h
   -version, --version, -v

`
	fmt.Println(help)
}
