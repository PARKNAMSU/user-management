package logservice_type

/*
로그 서비스 등록한 Application

	appId: 고유 ID
	appName: 어플리케이션 이름
	Url: 어플리케이션 url
	Status: 어플리케이션 사용 상태
*/
type ApplicationSchema struct {
	AppId   uint   `json:"appId" db:"appId"`
	AppName string `json:"appName" db:"appName"`
	Url     string `json:"url" db:"url"`
	Status  string `json:"status" db:"status"`
}
