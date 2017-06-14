# s3goamz

Loads files to s3 aws using the goamz library.

We implemented the multi.PutPart method, it will allow us to send chunks of the file to s3 aws, and in the end it joins all the parts.

The cool thing is that we will send everything cryptographed, but we may choose to encrypt the content before sending or not.

The authentication method is done right in the code we will encapsulate to get the access keys from the default location of aws ~/.aws/credentials, and also allow the keys coming straight from the code, the latter option is very dangerous and can cause a tremendous If the keys fall into the wrong hands, we strongly suggest keeping the keys on disk.


# Example of operation

![image]()


# Install

go get launchpad.net/goamz/aws

go get launchpad.net/goamz/s3

go get github.com/fatih/color

go get github.com/jeffotoni/s3goamz


# Permissions with ACL

```go

Private           = s3.ACL("private")
PublicRead        = s3.ACL("public-read")
PublicReadWrite   = s3.ACL("public-read-write")
AuthenticatedRead = s3.ACL("authenticated-read")
BucketOwnerRead   = s3.ACL("bucket-owner-read")
BucketOwnerFull   = s3.ACL("bucket-owner-full-control")

```

# Example 1 of submission

```go

go run main.go --put ~/Downloads/ex2.pdf --bucket name-bucket

```

# Example 2 of submission

```go

go run main.go -put ~/Downloads/ex2.pdf -bucket name-bucket

```

# Example 3 of submission

```go

go run main.go --version

```

# Example 4 of submission

```go

go run main.go -v

```


# Example 5 of submission

```go

go run main.go --help

```

# Example 6 of submission

```go

  -bucket string
    	Ex: name-bucket
  -crypt string
    	Exs: des/rsa/md5 (default "des")
  -put string
    	Ex: file.pdf

```
