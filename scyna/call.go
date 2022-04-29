package scyna

type ICallRepository interface {
	NewServiceCall()
	NewSignalCall()
	NewEventCall()
}

var CallRepository ICallRepository

func InitCallRepository() {

}

type callRepository struct {
}

func (r *callRepository) NewServiceCall() {

}

func newSignalCall() {

}

func newEventCall() {

}
