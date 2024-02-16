package models

import "gorm.io/gorm"

type Startup struct {
	gorm.Model
	Name             string  `json:"name" bson:"name"`
	CEO              string  `json:"ceo" bson:"ceo"`
	ValueProposition string  `json:"value_proposition" bson:"value_proposition"`
	Industry         string  `json:"industry" bson:"industry"`
	Founder          Founder `json:"founder" bson:"founder" gorm:"embedded"`
}

type Founder struct {
	gorm.Model
	Name       string `json:"name" bson:"name"`
	Profession string `json:"profession" bson:"profession"`
	Age        int    `json:"age" bson:"age"`
}

func (startup *Startup) Save() (*Startup, error){
	err := Database.Model(&startup).Create(&startup).Error
	if err != nil {
		return &Startup{}, err
	   }
	   return startup, nil
}

func FetchAllStartups () (*[]Startup, error){
	var startups [] Startup
	err := Database.Find(&startups).Error
	if err != nil{
		return &[]Startup{}, err
	}
	return &startups, nil
}

func FetchStartup(id string) (*Startup, error)  {
	var startup Startup
	err := Database.Where("id = ?", id).First(&startup).Error
	if err != nil{
		return &Startup{},err
	}
	return &startup, nil
}

func (startup *Startup) UpdateStartup(id string) (*Startup, error) {
	err := Database.Model(&Startup{}).Where("id = ?", id).Updates(startup).Error
	if err != nil {
	 return &Startup{}, err
	}
	return startup, nil
   }

func DeleteStartup(id string) error {
	err := Database.Model(&Startup{}).Where("id = ?", id).Delete(&Startup{}).Error
	if err != nil {
	 return err
	}
	return nil
   }
