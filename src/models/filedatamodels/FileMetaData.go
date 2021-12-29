package filedatamodels

import (
	"file-upload/databaseconnection/postresdb"
)

type FileMetaData struct {
	FileName    string
	ContentType string
}

const ()

func (filemetadata *FileMetaData) Save() error {
	queryInsert := `insert into "file_data"("file_name", "content_type") values($1, $2)`

	_, insErr := postresdb.Client.Exec(queryInsert, filemetadata.FileName, filemetadata.ContentType)
	if insErr != nil {
		return insErr
	}
	return nil
}
