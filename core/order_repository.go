//this is secondary port
package core

//as i said, this is secondary port
//it defines the methods that the application core expects from the order repository
type OrderRepository interface {
	Save(order Order) error
}
