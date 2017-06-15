# s3goamz

Loads files to s3 aws using the goamz library.

We implemented the multi.PutPart method, it will allow us to send chunks of the file to s3 aws, and in the end it joins all the parts.

The cool thing is that we will send everything cryptographed, but we may choose to encrypt the content before sending or not.

The authentication method is done right in the code we will encapsulate to get the access keys from the default location of aws ~/.aws/credentials, and also allow the keys coming straight from the code, the latter option is very dangerous and can cause a tremendous If the keys fall into the wrong hands, we strongly suggest keeping the keys on disk.


# Example of operation

![image](https://github.com/jeffotoni/s3goamz/blob/master/img/ex1.gif)

# Packages

go get launchpad.net/goamz/aws

go get launchpad.net/goamz/s3

go get github.com/fatih/color

got get github.com/jeffotoni/gocry

go get github.com/jeffotoni/s3goamz

# Install

$ go build s3goamz.go

$ sudo cp s3goamz /usr/bin

# Example 1

```go

$ s3goamz --put ~/Downloads/ex2.pdf --bucket name-bucket

```

# Example 2 

```go

$ s3goamz -put ~/Downloads/ex2.pdf -bucket name-bucket

```

# Example 3 Encrypting file

```go

$ s3goamz -put ~/Downloads/ex2.pdf -bucket name-bucket --crypt

Will encrypt... ~/Downloads/ex2.pdf.crypt
Used key:  DKYPENJXW43SMOJCU6F5TMFVOUANMJNL

```

# Example 4 

```go

$ s3goamz --version

```

# Example 5

```go

$ s3goamz -v

```


# Example 6 

```go

$ s3goamz --help

```

# Example 7

```go
	
$ s3goamz --help

   --acl string
    	Ex: read|write|all (default "read")

  --bucket string
    	Ex: name-bucket

  --crypt string
    	empty value

  --put string
    	Ex: file.pdf


```
