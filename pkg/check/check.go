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
)

func Exists(fileName string) bool {

	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
}

func Args() {

	var stringCmd string

	var existCmd int

	var lenArgs int

	existCmd = 0

	lenArgs = len(os.Args)

	//
	// == 1 default cmd
	//

	if lenArgs == 1 {

		PrintDefaults()
	}

	//
	// Validate hidden flags
	//

	fmt.Println(len(os.Args))

	Argsx := os.Args

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

	value, ok := arrayParam[4]

	if ok {

		fmt.Println("value: ", value)

	} else {

		fmt.Println("key not found")
	}

	fmt.Println(arrayParam)

	os.Exit(0)

	for x := range os.Args {

		stringCmd = strings.Trim(os.Args[x], " ")
		stringCmd = strings.TrimSpace(stringCmd)
		stringCmd = strings.ToLower(stringCmd)

		//fmt.Println("args: ", stringCmd, " ", x)

		switch stringCmd {

		case "--put":

			fmt.Println(Argsx[x])

			//value, ok := Argsx[x+1]

			// if ok {
			// 	fmt.Println("value: ", value)
			// } else {
			// 	fmt.Println("key not found")
			// }

			// fmt.Println("val: ", val, " ex: ", exists)

			existCmd++

		case "--bucket":

			existCmd++

		case "--crypt":

			existCmd++

		case "--acl":

			existCmd++

		case "--help":

			existCmd++

		case "--h":

			existCmd++

		case "--version":

			existCmd++

		case "--v":

			existCmd++

		case "-put":

			existCmd++

		case "-bucket":

			existCmd++

		case "-crypt":

			existCmd++

		case "-acl":

			existCmd++

		case "-help":

			existCmd++

		case "-h":

			existCmd++

		case "-version":

			existCmd++

		case "-v":

			existCmd++

		}

		//if existCmd ==
		// case "put":

		// 	fmt.Println(len(os.Args))

		// 	if len(os.Args) <= 2 {

		// 		boldRed.Println("\nMissing file as parameter ex: --put file.pdf\n")
		// 		os.Exit(0)
		// 	}

		// 	stringCmd2 = strings.Trim(os.Args[x+1], "-")
		// 	stringCmd2 = strings.TrimSpace(stringCmd2)
		// 	FileUpload = fmt.Sprintf("%s", stringCmd2)

		// 	//
		// 	// if /dir/dir/file
		// 	//

		// case "bucket":

		// 	fmt.Println(len(os.Args))

		// 	//if len(os.Args) <= 2 {

		// 	stringCmd2 = strings.Trim(os.Args[x+1], "-")
		// 	stringCmd2 = strings.TrimSpace(stringCmd2)
		// 	Bucket = fmt.Sprintf("%s", stringCmd2)
		// 	//fmt.Println("Bucket: ", stringCmd2)

		// case "crypt":

		// 	cryptInt += 1

		// case "acl":

		// 	stringCmd2 = strings.Trim(os.Args[x+1], "-")
		// 	stringCmd2 = strings.TrimSpace(stringCmd2)
		// 	stringCmd2 = strings.ToLower(stringCmd2)

		// 	stringAclTmp = fmt.Sprintf("%s", stringCmd2)

		// 	if stringAclTmp == "read" {

		// 		stringAcl = BucketOwnerRead

		// 	} else if stringAclTmp == "write" {

		// 		stringAcl = PublicReadWrite

		// 	} else if stringAclTmp == "all" {

		// 		stringAcl = BucketOwnerFull

		// 	} else {

		// 		boldYellow.Println("Acl does not exist! Try red | write | all")
		// 		os.Exit(0)

		// 	}

		// case "version":

		// 	boldYellow.Println("v.1.0")
		// 	os.Exit(0)

		// case "v":

		// 	boldYellow.Println("v.1.0")
		// 	os.Exit(0)

		// case "help":

		// 	flag.PrintDefaults()
		// 	os.Exit(0)

		// case "h":

		// 	flag.PrintDefaults()
		// 	os.Exit(0)

		// default:
		// 	//flag.PrintDefaults()
		// 	//os.Exit(0)() {
		// }
	}

	fmt.Println("Len: ", lenArgs)
	fmt.Println("Exist: ", existCmd)
}

func PrintDefaults() {

	var help string

	help = `	
	Use: 
	s3goamz [OPTION]...
 	or: s3goamz --put file.pdf --bucket name-bucket [options]
 	or: s3goamz --put file.pdf --bucket name-bucket --acl read [options]
 	or: s3goamz --put file.pdf --bucket name-bucket --acl read --crypt

 	Put and bucket arguments are required.
 	-put, --put <file>
`
	fmt.Println(help)
}
