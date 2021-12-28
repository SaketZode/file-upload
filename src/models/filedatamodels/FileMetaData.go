package filedatamodels

import (
	"file-upload/databaseconnection/postresdb"
	"fmt"
)

type FileMetaData struct {
	FileName    string
	ContentType string
}

const ()

func (filemetadata *FileMetaData) Save() error {
	queryInsert := "INSERT INTO file_data(\"file_name\", \"content_type\") VALUES(\"" + filemetadata.FileName + "\",\"" + filemetadata.ContentType + "\")"
	fmt.Println(queryInsert)

	_, insErr := postresdb.Client.Exec(queryInsert)
	if insErr != nil {
		return insErr
	}
	return nil
}
