package types

//Money представляет собой  денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д).
type Money int64

//Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д).
type Category string

//Status представляет собой статус платежей.
type Status string

//Предопределённые статус платежей.
const (
	StatusOk         Status = "Ok"
	StatusFail       Status = "Fail"
	StatusInProgress Status = "INPROGRESS"
)

//Payment представлет информацию о платёже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
