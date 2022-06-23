package image

import (
	"encoding/base64"
  "bytes"
	"net/http"
  "image"
  "image/png"
  "image/jpeg"
  "github.com/auyer/steganography"
  "log"
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
  var img image.Image
  var err error
  
  reader := bytes.NewReader(imageBytes)

  mimeType := http.DetectContentType(imageBytes)
	switch mimeType {
	  case "image/jpeg":
		  img, err = jpeg.Decode(reader)
	  case "image/png":
		  img, err = png.Decode(reader)
	}
  
  if err != nil {
    log.Fatalf("Error converting file to image %v", err)
  }

  w := new(bytes.Buffer)
  err = steganography.Encode(w, img, []byte(message))
  if err != nil {
    log.Fatalf("Error Encoding file %v", err)
  }

  *outBytes = w.Bytes()
}

func Decode(imageBytes []byte) string {
  var img image.Image
  var err error
  
  reader := bytes.NewReader(imageBytes)

  mimeType := http.DetectContentType(imageBytes)
	switch mimeType {
	  case "image/jpeg":
		  img, err = jpeg.Decode(reader)
	  case "image/png":
		  img, err = png.Decode(reader)
	}

  if err != nil {
    log.Fatalf("Error converting file to image %v", err)
  }

  sizeOfMessage := steganography.GetMessageSizeFromImage(img)
  msg := steganography.Decode(sizeOfMessage, img)

  return string(msg)
}