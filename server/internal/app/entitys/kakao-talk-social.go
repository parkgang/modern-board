package entitys

type KakaoTalkSocial struct {
	Id             uint   `gorm:"primary_key;"`
	Email          string `gorm:"unique;not null;size:50;"`
	NickName       string `gorm:"not null;size:20;"`
	ThumbnailImage string `gorm:"size:2083;"` // 이미지의 url을 저장하기 위한 최대 크기로 잡아놓았습니다.
}
