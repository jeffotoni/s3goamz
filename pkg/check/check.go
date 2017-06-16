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
func GenArgs() (string, string, string, int) {

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
	}

	//
	//
	//
	arrayParam := map[int]string{}

	//
	//
	//
	for j, val := range os.Args {

		//
		//
		//
		arrayParam[j] = val
	}

	for x := range arrayParam {

		stringCmd = strings.Trim(arrayParam[x], " ")
		stringCmd = strings.TrimSpace(stringCmd)
		stringCmd = strings.ToLower(stringCmd)

		switch stringCmd {

		case "--put":

			FileUpload = validPut(x, arrayParam)
			existCmd++

		case "-put":

			FileUpload = validPut(x, arrayParam)
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
	if cryptInt == 0 {

		PrintDefaults()
		os.Exit(0)
	}

	return FileUpload, Bucket, stringAcl, cryptInt
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

		if i > -1 {

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
func validBucket(x int, arrayParam map[int]string) string {

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	var stringCmd2 string
	var Bucket string

	value, ok := arrayParam[x+1]

	if ok {

		i := strings.Index(value, "-")

		if i > -1 {

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

   Put and bucket arguments are required.
   -put,     --put      <file>    The file and its respective path
   -bucket,  --bucket   <name>    Bucket name s3
   -acl,     --acl      <options> read, write, all
   -crypt,   --crypt    has no parameter
   -help,    --help     -h
   -version, --version, -v

`
	fmt.Println(help)
}
