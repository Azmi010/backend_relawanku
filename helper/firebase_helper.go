package helper

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

func UploadImageToFirebase(bucketName, folderPath, fileName string, file io.Reader) (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY"))))
	if err != nil {
		return "", fmt.Errorf("failed to create Firebase storage client: %w", err)
	}
	defer client.Close()

	objectName := path.Join("relawanku", folderPath, fmt.Sprintf("%d_%s", time.Now().Unix(), fileName))
	object := client.Bucket(bucketName).Object(objectName)

	writer := object.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	downloadToken := uuid.New().String()
	metadata := map[string]string{
		"firebaseStorageDownloadTokens": downloadToken,
	}
	_, err = object.Update(ctx, storage.ObjectAttrsToUpdate{
		Metadata: metadata,
	})
	if err != nil {
		return "", fmt.Errorf("failed to update object metadata: %w", err)
	}

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get object attributes: %w", err)
	}

	encodedName := url.PathEscape(attrs.Name)

	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s",
		attrs.Bucket, encodedName, downloadToken)

	return url, nil
}