package cephclient

import(
	"fmt"
	"bytes"

	"github.com/ceph/go-ceph/rados"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/awsutil"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3"
    "net/http"
    "os"
)

type CephClient struct {
	host         string
	port         int
	ioContexts   map[string]*rados.IOContext
}

func NewCephClient (host string, port int) (*CephClient, error){
	var cephclient = &CephClient{host: host, port: port}
	cephclient.ioContexts = make(map[string]*rados.IOContext)
	err := cephclient.connect()
	if err != nil {
		return nil, err
	}

	filepath = ""
	err = writeFile(filepath)
	if err != nil {
		return nil, err
	}
	return cephclient, nil

	

}

func (client *CephClient) connect() error {
	connection, err = rados.NewConn()
	if err != nil{
		return error
	}
	connection.ReadDefaultConfigFile()
	if client.host != ""{
		address := fmt.Sprintf("%s:%d", client.host, client.port)
		connection.SetConfigOption("mon_host", address)
	}
	connection.SetConfigOption("client_mount_timeout", "5")
	err = connection.Connect()
	if err != nil {
		connection.Shutdown()
		return err
	}
	client.conn = connection
	return nil
}

func createBucket(bucketname string) (string , error){

	params := &s3.CreateBucketInput{
                 Bucket: aws.String(bucketname),
    }

    result, err := s3client.CreateBucket(params)
    if err != nil{

    	return nil, err
    }

    return result, nil
}

func writeFile(fileToUpload string) error {
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err = creds.Get()
	if err!= nil{
		return err
		os.Exit(1)
	}

	config := &aws.Config{
		Region:		aws.String(""),
		Endpoint:	aws.String("s3.amazonaws.com"),
		S3ForcePathStyle:	aws.Bool(true),
		Credentials:		creds,
		LogLevel:			0,
	}

	s3client := s3.New(config)
	result, _ = createBucket(myawsbucket)
	bucketName = *result.Name
	file, err := os.Open(fileToUpload)

	if err != nil{
		return err
		os.Exit(1)
	}

	defer file.Close()
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)
	file.read(buffer)

	fileBytes = bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	s3Path := aws.String(*result.Location) + file.Name()

	params := &s3.PutObjectInput{
                 Bucket:        aws.String(bucketName), 
                 Key:           aws.String(s3Path),
                 ACL:           aws.String("public-read"),
                 Body:          fileBytes,
                 ContentLength: aws.Long(size),
                 ContentType:   aws.String(fileType),
                 Metadata: map[string]*string{
                         "Key": aws.String("MetadataValue"),
                 },
                 
    }

    res, err := s3client.PutObject(params)

    if err != nil {
    	if awsErr, ok := err.(awserr.Error); ok {
    		return awsErr.OrigErr()
    	}
    	else{
    		return err.Error()
    	}

    }

	return nil
}