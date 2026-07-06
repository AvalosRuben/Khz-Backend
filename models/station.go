type Station struct{
	Id string `gorm:"type:uuid;primaryKey"`
	Frequency float64 `gorm:"column:frequency;unique"`
	Name string `gorm:"column:name"`
	Genre string `gorm:"column:genre"`
	Stream_Url string `gorm:"column:stream_url"`
	Created_At time.Time `gorm:"column:created_at"`
}