package logging

type LogService interface {
	CreateRedirectLog(log RedirectLog) error
	GetLogs() ([]RedirectLog, error)
}

type LogStorage interface {
	AddLog(log RedirectLog) error
	GetLogs() ([]RedirectLog, error)
}
