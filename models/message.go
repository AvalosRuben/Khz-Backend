type Message struct {
	Id string `gorm:"type:uuid;primaryKey"`
	Created_At time.Time `gorm: "column:created_at"`
	Station_Id string `gorm:"type:uuid;column:station_id"`
	Station Station `gorm:"foreignKey:StationId"`
	User_Id string `gorm:"type:uuid;column:user_id"`
	User Profile `gorm:"foreignKey:ProfileId`
	Content string `gorm:"column:content"`
}