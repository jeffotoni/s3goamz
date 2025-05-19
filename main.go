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
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	cry "github.com/jeffotoni/gocry/pkg"
	auth "github.com/jeffotoni/s3goamz/pkg/auth"
	check "github.com/jeffotoni/s3goamz/pkg/check"
	erro "github.com/jeffotoni/s3goamz/pkg/erro"
	runer "github.com/jeffotoni/s3goamz/pkg/runer"
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
	var FileUpGet string

	//
	//
	//
	var Bucket string

	//
	//
	//
	var cryptInt int

	//
	//
	//
	var putGet string

	//
	//
	//
	var strcommand string

	//
	//
	//
	var lastValue string

	//
	// FileUpGet, Bucket, stringAcl
	//
	FileUpGet, Bucket, stringAclTmp, cryptInt, putGet = check.GenArgs()

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

	if putGet == "get" {

		strcommand = "start download to [ --get " + lastValue + " --bucket " + Bucket + "]"

		boldYellow.Println(strcommand)

		//
		//
		//
		timer := runer.RunerTimer()

		dbyte, erroget := conn.Get(FileUpGet)
		<-timer

		fmt.Print("\r")
		fmt.Print("\033[?25h")

		erro.Check(erroget)

		err3 := ioutil.WriteFile("./"+FileUpGet, dbyte, 0644)

		erro.Check(err3)

		// Writing bytes to disk

		boldYellow.Println("Download done successfully: ", FileUpGet)

		//
		//
		//

		if cryptInt > 0 {

			cry.Decrypt(keyDefault, FileUpGet)

			//
			// Will have to reopen etc ...
			//

			boldWhite.Println("Decrypted file... ", FileUpGet+".descr")
			boldYellow.Println("Used key: ", keyDefault)
		}

	} else {

		if stringAclTmp == "read" {

			stringAcl = BucketOwnerRead

		} else if stringAclTmp == "write" {

			stringAcl = PublicReadWrite

		} else if stringAclTmp == "all" {

			stringAcl = BucketOwnerFull
		}

		//
		//
		//
		if cryptInt > 0 {

			strcommand = "start upload to [ --put " + lastValue + " --bucket " + Bucket + " --acl " + stringAclTmp + " --crypt ]"

		} else {

			strcommand = "start upload to [ --put " + lastValue + " --bucket " + Bucket + " --acl " + stringAclTmp + " ]"
		}

		boldYellow.Println(strcommand)

		//
		//
		//
		vetSplit := strings.Split(FileUpGet, "/")

		//
		//
		//
		sizeVet := len(vetSplit)

		if sizeVet > 0 {

			lastValue = vetSplit[sizeVet-1]

		} else {

			lastValue = FileUpGet
		}

		//
		// OR
		//
		// bucket := conn.Bucket(Bucket) // change this your bucket name

		if check.Exists(FileUpGet) == true {

			var fileCrypt string

			if cryptInt == 1 {

				//
				// File must be encrypted
				//

				cry.Crypt(keyDefault, FileUpGet)
				fileCrypt = FileUpGet + ".crypt"

				//
				// Will have to reopen etc ...
				//

				fmt.Println("Will encrypt...", fileCrypt)
				fmt.Println("Used key: ", keyDefault)

				FileUpGet = fileCrypt

				//os.Exit(0)
			}

			//
			//
			//
			file, errx := os.Open(FileUpGet)
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
			multi, err := conn.InitMulti(FileUpGet, filetype, stringAcl)

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
			timer := runer.RunerTimer()

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

			boldRed.Println("Erro, File [" + FileUpGet + "]does not exist!")
			os.Exit(0)
		}
	}
}
