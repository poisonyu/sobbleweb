// package bank

// // type Draw struct {
// // 	amount    int
// // 	withdraws chan bool
// // }

// var deposits = make(chan int)
// var balances = make(chan int)

// // var draw = new(Draw)

// var withdraws = make(chan bool)

// func Deposit(amount int) {
// 	deposits <- amount
// }

// func Balance() int {
// 	return <-balances
// }

// // 取款时输入的amount为负
// func Withdraw(amount int) bool {
// 	deposits <- amount
// 	return <-withdraws
// }

// func teller() {
// 	var balance int
// 	for {
// 		select {
// 		case amount := <-deposits:
// 			if amount > 0 {
// 				balance += amount
// 			} else {
// 				balance += amount
// 				if balance < 0 {
// 					balance -= amount
// 					withdraws <- false

// 				}
// 				withdraws <- true

// 			}
// 		case balances <- balance:
// 		}
// 	}
// }

// func init() {
// 	go teller()
// }

package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan *WithdrawChan)

type WithdrawChan struct {
	amount     int
	isWithdraw chan bool
}

func Deposit(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	withdrawChan := new(WithdrawChan)
	withdrawChan.amount = amount
	// withdrawsChan.withdraw := make(chan bool)
	withdraws <- withdrawChan

	return <-withdrawChan.isWithdraw

}
func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdrawChan := <-withdraws:
			amount := withdrawChan.amount
			if balance < amount {
				withdrawChan.isWithdraw <- false
			} else {
				balance -= amount
				withdrawChan.isWithdraw <- true
			}
		case balances <- balance:
		}
	}
}
func init() {
	go teller()
}
