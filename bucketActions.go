package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func createRegionClient(region string) BucketBasics {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		panic("couldn't load configuration")
	}

	//Set up S3 client
	s3Client := s3.NewFromConfig(sdkConfig)

	//set up S3 Client Bucket Basics receiver type,
	bb := BucketBasics{s3Client}
	return bb
}

func createDefaultClient() BucketBasics {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		panic("couldn't load configuration")
	}

	//Set up S3 client
	s3Client := s3.NewFromConfig(sdkConfig)

	//set up S3 Client Bucket Basics receiver type,
	bb := BucketBasics{s3Client}
	return bb
}

func getBuckets() {
	//Get all buckets in the account
	var bb = createDefaultClient()
	buckets, _ := BucketBasics.ListBuckets(bb)
	for _, bucket := range buckets {
		fmt.Println(*bucket.Name + ": " + bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}
}

func getBucketsInRegion(region string) {
	//Get all buckets in the account
	var bb = createDefaultClient()
	buckets, _ := BucketBasics.ListBuckets(bb)
	for _, bucket := range buckets {
		//Get bucket region
		reg, _ := BucketBasics.GetBucketRegion(bb, *bucket.Name)
		//Print the bucket info if region is the same as specified in const
		if reg == region {
			fmt.Println(*bucket.Name + ": " + bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
		}
	}
}

func createNewBucket(region string, bucketName string) {
	var bb = createRegionClient(region)
	//Creates a S3 bucket with the given name and specified region
	bb.CreateBucket(bucketName, region)
	fmt.Printf("Bucket %v was created in Region %v \n", bucketName, region)
}

func uploadObject(region string, bucketName string, localFile string, name string) {
	var bb = createRegionClient(region)
	//Uploads an object into a specified bucket
	bb.UploadFile(bucketName, name, localFile)
	fmt.Println("File Uploaded Succcesfully")
}

func deleteObject(region string, bucketName string, name string) {
	var bb = createRegionClient(region)

	//Delete an object from a specified bucket
	objArray := []string{name}
	bb.DeleteObjects(bucketName, objArray)

	fmt.Println("Object Deleted Succcesfully")
}

func deleteObjects(region string, bucketName string, objArray []string) {
	var bb = createRegionClient(region)

	bb.DeleteObjects(bucketName, objArray)

	fmt.Println("Objects Deleted Succcesfully")
}

func downloadObject(region string, bucketName string, name string, localName string) {
	var bb = createRegionClient(region)

	bb.DownloadFile(bucketName, name, localName)

	fmt.Println("Object Downloaded Succcesfully")
}

func getObjects(region string, bucketName string) {
	var bb = createRegionClient(region)

	objects, _ := bb.ListObjects(bucketName)
	for _, object := range objects {
		fmt.Println(*object.Key + ": " + object.LastModified.Format("2006-01-02 15:04:05 Monday") + ", Size:" + fmt.Sprint(*object.Size) + " B")
	}
}
