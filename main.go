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

func main() {

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
	putFlag := flag.String("put", "", "Ex: file.pdf")

	//
	//
	//
	bucketFlag := flag.String("bucket", "", "Ex: name-bucket")

	if len(os.Args) < 5 || len(os.Args) > 5 {

		boldRed.Println("You must enter the name of the file and bucket you want to send")
		boldYellow.Println("-put [file.pdf] -bucket [s3-bucket]")
		os.Exit(0)

	} else if len(os.Args) == 2 {

		boldRed.Println("ok.")
	}

	flag.Parse()

	//fmt.Printf("putFlag: %s %t\n", *putFlag)

	if *putFlag == "" || *bucketFlag == "" {

		flag.PrintDefaults()
		os.Exit(0)
	}

	var FileUpload string
	var Bucket string

	//
	//
	//
	flag.Visit(func(f *flag.Flag) {

		//fmt.Println(f)

		switch f.Name {

		case "put":

			FileUpload = fmt.Sprintf("%s", f.Value)

			//
			// if /dir/dir/file
			//

		case "bucket":

			Bucket = fmt.Sprintf("%s", f.Value)

		default:
			flag.PrintDefaults()

		}
	})

	//fmt.Println("Flags: ", len(os.Args))

	//
	// if "/" remove / , only the file name
	// FileUpload
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

	strcommand := "start upload to [ -put " + lastValue + " -bucket " + Bucket + " ]"

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

		if fileSize < fileChunk {

			fmt.Println("Error, The file has to be larger than 5mb to send in parts to amazon s3.")
			os.Exit(0)

		}
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
		multi, err := conn.InitMulti("/"+FileUpload, filetype, BucketOwnerRead)

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

		//
		//
		//
		boldRed.Println("File Size:", fileSize, "byte", sizeTotal, "Mb", "Chunk:", chunkPart, "Mb")

		//
		//
		//
		boldYellow.Println("Uploading...")

		go func() {
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, os.Interrupt)

			<-sc

			fmt.Println("\ncanceled!")
			fmt.Print("\033[?25h")
			os.Exit(0)
		}()

		fmt.Print("\033[?25l")

		timer := time.Tick(time.Duration(150) * time.Millisecond)

		s := []rune(`|/~\`)
		//s := []rune(`-=*=`)
		//s := []rune(`◐◓◑◒`)
		i := 0

		go func() {
			for {

				<-timer

				fmt.Print("\r")
				fmt.Print(string(s[i]))

				i++

				if i == len(s) {
					i = 0
				}
			}
		}()

		//go func() {
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

			//
			//
			//
			boldWhite.Printf("Processing %d piece of %d and uploaded %d bytes.\n ", int(i), int(totalPartsNum), int(sizePart))

			//
			//
			//
			parts = append(parts, piece)
		}
		//}()

		//
		//
		//
		err = multi.Complete(parts)

		//
		//
		//
		erro.Check(err)

		boldYellow.Println("\n\nUpload completed...")

	} else {

		boldRed.Println("Erro, File [" + FileUpload + "]does not exist!")
		os.Exit(0)
	}

	// bucket := connection.Bucket("name-your-bucket") // change this your bucket name

}
