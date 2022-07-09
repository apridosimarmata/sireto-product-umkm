package Utils

import (
	"bytes"
	"fmt"
	"strings"

	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// TODO fill these in!
const (
	S3_REGION = "us-west-2"
	S3_BUCKET = "sireto"
)

/*
func main() {

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	// Upload
	err = AddFileToS3(s, "thumbnail.png")
	if err != nil {
		log.Fatal(err)
	}
}
*/
// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, base64Image string, name string) error {
	// Open the file for use
	/*file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()*/

	// Get file size and read the file content into a buffer
	/*fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)*/

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	i := strings.Index(base64Image, ",")
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Image[i+1:]))
	buf := new(bytes.Buffer)
	buf.ReadFrom(dec)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String("images/" + name + ".png"),
		ACL:    aws.String("private"),
		Body:   bytes.NewReader(buf.Bytes()),
		/*ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),*/
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	fmt.Println(err)
	return err
}
