# s3goamz

Loads files to s3 aws using the goamz library.

We implemented the multi.PutPart method, it will allow us to send chunks of the file to s3 aws, and in the end it joins all the parts.

The cool thing is that we will send everything cryptographed, but we may choose to encrypt the content before sending or not.

The authentication method is done right in the code we will encapsulate to get the access keys from the default location of aws ~/.aws/credentials, and also allow the keys coming straight from the code, the latter option is very dangerous and can cause a tremendous If the keys fall into the wrong hands, we strongly suggest keeping the keys on disk.


# Example of operation

![image](https://github.com/jeffotoni/s3goamz/blob/master/img/s3goamz.gif)

# Packages

go get launchpad.net/goamz/aws

go get launchpad.net/goamz/s3

go get github.com/fatih/color

got get github.com/jeffotoni/gocry

go get github.com/jeffotoni/s3goamz

# Install

$ go build s3goamz.go

$ sudo cp s3goamz /usr/bin

# help

```go
	
$ s3goamz --help

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

```

# Example 1

```go

$ s3goamz --put ~/Downloads/ex2.pdf --bucket name-bucket

```

# Example 2 Encrypting file

```go

$ s3goamz -put ~/Downloads/ex2.pdf -bucket name-bucket --crypt

Will encrypt... ~/Downloads/ex2.pdf.crypt
Used key:  DKYPENJXW43SMOJCU6F5TMFVOUANMJNL

```

# Example 3

```go

$ s3goamz --version

```