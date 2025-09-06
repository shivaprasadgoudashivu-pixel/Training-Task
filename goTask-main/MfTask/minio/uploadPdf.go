package minio

import (
	"bytes"
	"context"
	"fmt"
	"keycloak-demo/model"
	"log"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadPDF(order *model.ORDER) string {
	// ====== Step 1: Generate PDF ======
	var buf bytes.Buffer
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, " Order  - Report #12345")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Scheme :"+order.Scheme)
	pdf.Ln(10)
	pdf.Cell(40, 10, "Scheme Code : "+order.Scheme_code)
	pdf.Ln(10)
	pdf.Cell(40, 10, "userID : "+strconv.Itoa(order.UserId))
	pdf.Ln(10)
	pdf.Cell(40, 10, "Status : "+(order.Status))
	pdf.Ln(10)
	pdf.Cell(40, 10, "units : "+strconv.Itoa(order.Units))
	pdf.Ln(10)
	pdf.Cell(40, 10, "Amount : "+strconv.Itoa(order.Price))

	err := pdf.Output(&buf) // write PDF into buffer
	if err != nil {
		log.Fatal(err)
	}

	// ====== Step 2: Upload to MinIO ======
	endpoint := "minio:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	bucketName := "report"
	objectName := strconv.Itoa(order.UserId) + strconv.Itoa(order.Id) + "report.pdf"
	fmt.Printf("minio intialize \n", objectName)
	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Ensure bucket exists
	ctx := context.Background()

	fmt.Printf("minio bucket creation try \n")
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists != nil || !exists {
			log.Fatal(err)

		}
	}

	fmt.Printf("PDF uploading to minio \n")
	// Upload PDF buffer
	_, err = minioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader(buf.Bytes()), int64(buf.Len()), minio.PutObjectOptions{ContentType: "application/pdf"})
	if err != nil {
		log.Fatal(err)
		fmt.Print(err)
	}

	fmt.Printf("PDF uploaded to minio \n")

	// ====== Step 3: Get public link (presigned URL) ======
	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, 7*24*time.Hour, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Save `presignedURL.String()` into your DB
	return fmt.Sprintf("%s", presignedURL.String())
}
