package image

import (
	"encoding/base64"
  "bytes"
	"net/http"
  "image"
  "github.com/auyer/steganography"
  "fmt"
)

func ToBase64(imageBytes []byte) string {
	var base64Encoding string

	mimeType := http.DetectContentType(imageBytes)
	switch mimeType {
	  case "image/jpeg":
		  base64Encoding += "image/jpeg;base64,"
	  case "image/png":
		  base64Encoding += "image/png;base64,"
	}

  base64Encoding += base64.StdEncoding.EncodeToString(imageBytes)
	return base64Encoding
}

func Encode(outBytes *[]byte, imageBytes []byte, message string) {
  reader := bytes.NewReader(imageBytes)
  img, _, err := image.Decode(reader)
  if err != nil {
    fmt.Printf("Error converting file to image %v", err)
  }

  w := new(bytes.Buffer)
  err = steganography.Encode(w, img, []byte(message))
  if err != nil {
    fmt.Printf("Error Encoding file %v", err)
  }

  *outBytes = w.Bytes()
}

func Decode(imageBytes []byte) string {
  reader := bytes.NewReader(imageBytes)
  img, _, _ := image.Decode(reader)

  sizeOfMessage := steganography.GetMessageSizeFromImage(img)
  msg := steganography.Decode(sizeOfMessage, img)

  return string(msg)
}