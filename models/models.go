package models

import (
	"io"
	"ipfs-file-server/models/ipfs"
	"ipfs-file-server/models/orm"
	"time"
)

func SaveFileToIpfs(name string, size int64, r io.Reader) error {
	cid, err := ipfs.AddFile(r)
	if err != nil {
		return err
	}

	cTime := time.Now()

	file := &orm.IPFS_saved_file{
		FileCid:    cid,
		FileName:   name,
		FileSize:   size,
		SavedDate:  cTime.Unix(),
		DateMark:   cTime.Format("2006-01-02"),
		SharedMark: true,
	}

	orm.SavedFiles(file)
	return err
}

type FileList struct {
	List []*orm.IPFS_saved_file `json:"list"`
}

func GetFileList() *FileList {
	list := orm.ReadFileList()
	fileList := &FileList{
		List: list,
	}
	return fileList
}

func ShareFile(cid string, shared bool) error {

	file := &orm.IPFS_saved_file{
		FileCid:    cid,
		SharedMark: shared,
	}

	return orm.ShareFile(file)
}
