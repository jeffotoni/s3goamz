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

go build s3goamz.go

# Example 1

```go

go run s3goamz.go --put ~/Downloads/ex2.pdf --bucket name-bucket

```

# Example 2 

```go

go run s3goamz.go -put ~/Downloads/ex2.pdf -bucket name-bucket

```

# Example 3 

```go

go run s3goamz.go --version

```

# Example 4 

```go

go run s3goamz.go -v

```


# Example 5 

```go

go run s3goamz.go --help

```

# Example 6 

```go

   -acl string
    	Ex: read|write|all (default "read")
  -bucket string
    	Ex: name-bucket
  -crypt string
    	Exs: des|rsa|md5 (default "des")
  -put string
    	Ex: file.pdf

```
