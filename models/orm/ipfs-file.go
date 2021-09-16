package orm

func SavedFiles(file *IPFS_saved_file) {
	saveSingleData(file)
}

func ReadFileList() []*IPFS_saved_file {
	var list []*IPFS_saved_file
	if readAllOrderBy(new(IPFS_saved_file), "-saved_date", &list) {
		return list
	}

	return nil
}

func ShareFile(file *IPFS_saved_file) error {

	feilds := []string{"shared_mark"}

	return updateSingleData(file, feilds)
}
