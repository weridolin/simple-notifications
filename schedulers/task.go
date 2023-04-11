package schedulers

type Task interface {
	Run()
	GetUpInfo()
}
