package plugin

type InstallVerifiable interface {
	OnInstallVerifyStart(attempt Installation) error
	OnInstallVerifySuccess(attempt Installation) error
	OnInstallVerifyFailure(attempt Installation, failure error)
}
