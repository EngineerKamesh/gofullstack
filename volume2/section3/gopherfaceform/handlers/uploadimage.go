package handlers

import (
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

type UploadImageForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

func DisplayUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	RenderTemplate(w, "./templates/uploadimageform.html", u)
}

func ProcessUploadImage(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {

	file, fileheader, err := r.FormFile("imagefile")

	if err != nil {
		log.Println("Encountered error when attempting to read uploaded file: ", err)
	}

	randomFileName := GenerateUUID()

	if fileheader != nil {

		extension := filepath.Ext(fileheader.Filename)
		r.ParseMultipartForm(32 << 20)

		defer file.Close()

		imageFilePathWithoutExtension := "./static/uploads/images/" + randomFileName
		f, err := os.OpenFile(imageFilePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)

		thumbImageFilePath := imageFilePathWithoutExtension + "_thumb.png"
		originalimagefile, err := os.Open(imageFilePathWithoutExtension + extension)

		if err != nil {
			log.Println(err)
			return
		}

		img, err := png.Decode(originalimagefile)

		if err != nil {
			log.Println("Encountered Error while decoding image file: ", err)
		}

		thumbImage := resize.Resize(270, 0, img, resize.Lanczos3)
		thumbImageFile, err := os.Create(thumbImageFilePath)

		if err != nil {
			log.Println("Encountered error while resizing image:", err)
		}

		defer thumbImageFile.Close()

		png.Encode(thumbImageFile, thumbImage)

		m := make(map[string]string)
		m["thumbnailPath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + "_thumb.png"
		m["imagePath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + ".png"

		RenderTemplate(w, "./templates/imagepreview.html", m)

	} else {
		w.Write([]byte("Failed to process uploaded file!"))
	}
}

func ValidateUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {

	ProcessUploadImage(w, r, u)

}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	u := UploadImageForm{}
	u.Fields = make(map[string]string)
	u.Errors = make(map[string]string)

	switch r.Method {

	case "GET":
		DisplayUploadImageForm(w, r, &u)
	case "POST":
		ValidateUploadImageForm(w, r, &u)
	default:
		DisplayUploadImageForm(w, r, &u)
	}

}
