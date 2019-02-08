package models

import "fmt"

type QError struct {

	Where string 
	What string	 

}


func (e *QError) Error() string {
	return fmt.Sprintf("at %s, %s",
		e.Where, e.What)
}