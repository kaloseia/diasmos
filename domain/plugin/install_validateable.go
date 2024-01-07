package plugin

type InstallValidateable interface {
	OnInstallValidateStart(attempt Installation) error
	OnInstallValidateSuccess(attempt Installation) error
	OnInstallValidateFailure(attempt Installation, failure error)
}
