package orm

type IPFS_saved_file struct {
	FileCid    string `orm:"column(cid);pk" json:"cid"`
	FileName   string `orm:"column(name)" json:"name"`
	FileSize   int64  `orm:"column(size)" json:"size"`
	SavedDate  int64  `orm:"column(saved_date)" json:"saved_date"`
	DateMark   string `orm:"column(date_mark)" json:"date_mark"`
	SharedMark bool   `orm:"column(shared_mark)" json:"shared_mark"`
}

func (tb *IPFS_saved_file) TableName() string {
	return "IPFS_saved_file"
}
