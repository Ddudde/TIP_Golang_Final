package db

type Prodct struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Zag   string `json:"zag"`
	Price int32  `json:"price"`
	Gram  int32  `json:"gram"`
	Img   string `json:"img"`
}
