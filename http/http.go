package http

import (
	"html/template"
	"io/ioutil"
	"main/image"
	"net/http"
)

func getUploadedImage(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseMultipartForm(10 << 20)
	imgForm, _, _ := r.FormFile("image")
	imgBytes, _ := ioutil.ReadAll(imgForm)
  return imgBytes
}

func createImageData(imgB64 string, message string) map[string]string {
	return map[string]string{"image": imgB64, "message": message}
}

func handleEncode(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
	// if GET, show UI
	case http.MethodGet:
		imageSelectTemplate := template.Must(template.ParseFiles("templates/uploadImage.html"))
		imageSelectTemplate.Execute(w, nil)

	// if POST, upload and encode image
	case http.MethodPost:
    imgBytes := getUploadedImage(w, r)
    message := r.Form.Get("message")

    var o []byte
    image.Encode(&o, imgBytes, message)
    imgB64 := image.ToBase64(o)
    imgData := createImageData(imgB64, message)

    imgTemplate := template.Must(template.ParseFiles("templates/showImage.html"))
		imgTemplate.Execute(w, imgData)
	}
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
	// if GET, show UI
	case http.MethodGet:
		imageSelectTemplate := template.Must(template.ParseFiles("templates/uploadImage.html"))
		imageSelectTemplate.Execute(w, nil)

	// if POST, upload and encode image
	case http.MethodPost:
    imgBytes := getUploadedImage(w, r)
    message := image.Decode(imgBytes)
    
    imgB64 := image.ToBase64(imgBytes)
    imgData := createImageData(imgB64, message)

    imgTemplate := template.Must(template.ParseFiles("templates/showImage.html"))
		imgTemplate.Execute(w, imgData)
	}
}

func InitEndpoints() {
	// Intital web entry UI
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeTemplate := template.Must(template.ParseFiles("templates/uploadImage.html"))
		homeTemplate.Execute(w, nil)
	})

  // encode
  http.HandleFunc("/encode", handleEncode)

  // decode
  http.HandleFunc("/decode", handleDecode)

  // css
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

}

func StartServer() {
	http.ListenAndServe(":80", nil)
}
