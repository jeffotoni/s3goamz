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

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/fatih/color"
	cry "github.com/jeffotoni/gocry/pkg"
	auth "github.com/jeffotoni/s3goamz/pkg/auth"
	check "github.com/jeffotoni/s3goamz/pkg/check"
	erro "github.com/jeffotoni/s3goamz/pkg/erro"
	"launchpad.net/goamz/s3"
)

const (
	Private           = s3.ACL("private")
	PublicRead        = s3.ACL("public-read")
	PublicReadWrite   = s3.ACL("public-read-write")
	AuthenticatedRead = s3.ACL("authenticated-read")
	BucketOwnerRead   = s3.ACL("bucket-owner-read")
	BucketOwnerFull   = s3.ACL("bucket-owner-full-control")
)

const (
	fileChunk = 5 * (1 << 20) // 5MB
)

const (
	keyDefault = "DKYPENJXW43SMOJCU6F5TMFVOUANMJNL"
)

func main() {

	//
	//
	//
	stringAcl := BucketOwnerRead

	//
	//
	//
	stringAclTmp := "read"

	//
	//
	//
	white := color.New(color.FgWhite)
	boldWhite := white.Add(color.Bold)

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
	var FileUpload string

	//
	//
	//
	var Bucket string

	//
	//
	//
	flag.String("put", "", "Ex: file.pdf")

	//
	//
	//
	flag.String("bucket", "", "Ex: name-bucket")

	//
	//
	//
	flag.String("crypt", "", "empty value")

	//
	//
	//
	flag.String("acl", "read", "Ex: read|write|all")

	//
	//
	//
	sizeArgs := len(os.Args)

	//
	// Validate flags
	// and Validate hidden flags
	// --v
	// -v
	// --version
	// -version
	//
	// -help
	// --help
	// -h
	// --h
	//
	if sizeArgs <= 1 {

		flag.PrintDefaults()
		os.Exit(0)
	}

	//
	//
	//
	var stringCmd string

	//
	//
	//
	var stringCmd2 string

	var cryptInt int

	cryptInt = 0
	//
	// Validate hidden flags
	//
	for x := range os.Args {

		stringCmd = strings.Trim(os.Args[x], "-")
		stringCmd = strings.TrimSpace(stringCmd)
		stringCmd = strings.ToLower(stringCmd)

		//fmt.Println("args: ", sizeArgs, " ", x)

		switch stringCmd {

		case "put":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)
			FileUpload = fmt.Sprintf("%s", stringCmd2)
			//fmt.Println("put: ", stringCmd2)

			//
			// if /dir/dir/file
			//

		case "bucket":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)
			Bucket = fmt.Sprintf("%s", stringCmd2)
			//fmt.Println("Bucket: ", stringCmd2)

		case "crypt":

			cryptInt += 1

		case "acl":

			stringCmd2 = strings.Trim(os.Args[x+1], "-")
			stringCmd2 = strings.TrimSpace(stringCmd2)
			stringCmd2 = strings.ToLower(stringCmd2)

			stringAclTmp = fmt.Sprintf("%s", stringCmd2)

			if stringAclTmp == "read" {

				stringAcl = BucketOwnerRead

			} else if stringAclTmp == "write" {

				stringAcl = PublicReadWrite

			} else if stringAclTmp == "all" {

				stringAcl = BucketOwnerFull

			} else {

				boldYellow.Println("Acl does not exist! Try red | write | all")
				os.Exit(0)

			}

		case "version":

			boldYellow.Println("v.1.0")
			os.Exit(0)

		case "v":

			boldYellow.Println("v.1.0")
			os.Exit(0)

		case "help":

			flag.PrintDefaults()
			os.Exit(0)

		case "h":

			flag.PrintDefaults()
			os.Exit(0)

		default:
			//flag.PrintDefaults()
			//os.Exit(0)
		}
	}

	if stringCmd2 == "" {

		flag.PrintDefaults()
		os.Exit(0)
	}

	//
	//
	//
	var lastValue string

	//
	//
	//
	vetSplit := strings.Split(FileUpload, "/")

	//
	//
	//
	sizeVet := len(vetSplit)

	if sizeVet > 0 {

		lastValue = vetSplit[sizeVet-1]

	} else {

		lastValue = FileUpload
	}

	strcommand := "start upload to [ --put " + lastValue + " --bucket " + Bucket + " --acl " + stringAclTmp + "]"

	boldYellow.Println(strcommand)

	//
	// GetAuth (key, secret)
	//
	// connect := auth.GetAuth("xxxxxx-key", "xxxxx-secret")
	//
	// OR
	//
	// aws/.credentials
	// GetAuth()
	//
	// OR
	//
	// Getenv(Key) and Getenv(Secrete)
	// GetAuth()
	//
	connect := auth.GetAuth()

	//
	// s3.Bucket
	//
	conn := s3.Bucket{
		S3:   connect,
		Name: Bucket,
	}

	//
	// OR
	//
	// bucket := conn.Bucket(Bucket) // change this your bucket name

	if check.Exists(FileUpload) == true {

		var fileCrypt string

		if cryptInt == 1 {

			//
			// File must be encrypted
			//

			cry.Crypt(keyDefault, FileUpload)
			fileCrypt = FileUpload + ".crypt"

			//
			// Will have to reopen etc ...
			//

			fmt.Println("Will encrypt...", fileCrypt)
			fmt.Println("Used key: ", keyDefault)

			FileUpload = fileCrypt

			//os.Exit(0)
		}

		//
		//
		//
		file, errx := os.Open(FileUpload)
		erro.Check(errx)
		defer file.Close()

		//
		//
		//
		fileInfo, _ := file.Stat()

		//
		//
		//
		fileSize := fileInfo.Size()

		//
		//
		//
		bytes := make([]byte, fileSize)

		// read into buffer
		buffer := bufio.NewReader(file)

		//
		//
		//
		_, errx = buffer.Read(bytes)

		//
		//
		//
		erro.Check(errx)

		//
		//
		//
		filetype := http.DetectContentType(bytes)

		//
		//
		//
		multi, err := conn.InitMulti(FileUpload, filetype, stringAcl)

		//
		//
		//
		erro.Check(err)

		//
		//
		//
		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		//
		//
		//
		parts := []s3.Part{}

		//
		//
		//
		HeaderPart := strings.NewReader(string(bytes))

		//
		//
		//
		sizeTotal := (fileSize / (1024 * 1024))

		//
		//
		//
		chunkPart := (fileChunk / (1024 * 1024))

		go func() {
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, os.Interrupt)

			<-sc

			boldRed.Println("\ncanceled!")
			fmt.Print("\033[?25h")
			os.Exit(0)
		}()

		fmt.Print("\033[?25l")

		timer := time.Tick(time.Duration(50) * time.Millisecond)

		s := []rune(`|/~\`)
		//s := []rune(`-=*=`)
		//s := []rune(`◐◓◑◒`)
		i := 0

		go func() {
			for {

				<-timer

				fmt.Print("\r")
				boldWhite.Print(string(s[i]))

				i++

				if i == len(s) {
					i = 0
				}
			}
		}()

		//
		//
		//
		if fileSize < fileChunk {

			//
			//
			//
			boldRed.Println("File Size:", fileSize, "byte")

			//
			// Use another form of putAll shipping
			//
			partsx, erry := multi.PutAll(file, fileChunk)

			//
			//
			//
			if erry != nil {

				fmt.Println("Error using send with putAll: ", erry)
				os.Exit(0)
			}

			//
			//
			//
			erry = multi.Complete(partsx)

			//
			//
			//
			fmt.Print("\033[?25h")

			//
			//
			//
			boldYellow.Println("\n\nUpload completed...")

		} else {

			//
			//
			//
			boldRed.Println("File Size:", fileSize, "byte", sizeTotal, "Mb", "Chunk:", chunkPart, "Mb")

			//
			//
			//
			boldYellow.Println("Uploading by Parts...")

			//
			//
			//
			for i := uint64(0); i < totalPartsNum; i++ {

				<-timer

				//
				//
				//
				partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))

				//
				//
				//
				partBuffer := make([]byte, partSize)

				//
				//
				//
				sizePart, errx2 := io.ReadFull(HeaderPart, partBuffer)

				//
				//
				//
				erro.Check(errx2)

				//
				//
				//
				piece, errx3 := multi.PutPart(int(i)+1, strings.NewReader(string(partBuffer))) // write to S3 bucket part by part

				//
				//
				//
				erro.Check(errx3)

				fmt.Print("\r")
				//
				//
				//
				boldWhite.Printf("Processing %d piece of %d and uploaded %d bytes.\n ", int(i), int(totalPartsNum), int(sizePart))

				//
				//
				//
				parts = append(parts, piece)
			}

			//
			//
			//
			err = multi.Complete(parts)

			//
			//
			//
			erro.Check(err)

			fmt.Print("\033[?25h")
			boldYellow.Println("\n\nUpload completed...")
		}

	} else {

		boldRed.Println("Erro, File [" + FileUpload + "]does not exist!")
		os.Exit(0)
	}
}
