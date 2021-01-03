package logging

type loggingService struct {
	storage LogStorage
}

func InitService(st LogStorage) LogService {
	return &loggingService{st}
}

func (logSrv *loggingService) CreateRedirectLog(log RedirectLog) error {
	err := logSrv.storage.AddLog(log)
	if err != nil {
		return err
	}
	return nil
}

func (logSrv *loggingService) GetLogs() ([]RedirectLog, error) {
	result, err := logSrv.storage.GetLogs()
	if err != nil {
		return nil, err
	}
	return result, nil
}
