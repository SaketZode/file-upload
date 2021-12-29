package services

import (
	"file-upload/constants/filetypes"
	"file-upload/models/filedatamodels"
	"fmt"
	"io/ioutil"
	"net/http"
)

// validates file type
func detectFileType(fileBytes []byte) (string, error) {
	contentType := http.DetectContentType(fileBytes)
	if contentType == filetypes.FILE_TYPE_JPG || contentType == filetypes.FILE_TYPE_PNG {
		return contentType, nil
	}
	return "", fmt.Errorf("file type %s not supported", contentType)
}

//uploads file to specific folder
func UploadFile(w http.ResponseWriter, r *http.Request) {
	//fetching file content from request received
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	//check for file size
	if handler.Size > 8000000 {
		fmt.Fprintf(w, "File size exceeds the limit!")
		return
	}

	//reading file bytes
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	//detecting content-type
	cType, ftErr := detectFileType(fileBytes)
	if ftErr != nil {
		fmt.Fprintf(w, ("File type error : " + ftErr.Error()))
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//creating temporary file for storing the file on local directory
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Something went wrong whole uploading file!")
		return
	}

	defer tempFile.Close()

	//writing the contents to the temporary file created
	tempFile.Write(fileBytes)

	filedata := filedatamodels.FileMetaData{handler.Filename, cType}

	//persisting file meta-data to database
	if err := filedata.Save(); err != nil {
		fmt.Println("Save error:", err)
		fmt.Fprintf(w, "Internal server error!!")
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
