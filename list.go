package main

import (
	"fmt"
)

//Vertex of your chain list.
type list struct {
		X			int
		next	*list
}

//Function to create item who take in arg your struct val, next always = nil
func 	item(X int) *list{
	return &list{X, nil}
}

//Function for add your value to chain list.
func 	add_to_list(lst **list, item *list) {
	var tmp *list = nil

	//check
	if (lst == nil || item == nil) {
		return
	}

	//Add to list.
	if (*lst == nil){
		*lst = item
	} else {
		tmp = *lst
		for (tmp.next != nil) {
			tmp = tmp.next
		}
		tmp.next = item
	}
}

//example of use.
func 	main(){
	var lst *list = nil
	add_to_list(&lst, item(3))
	add_to_list(&lst, item(4))
	add_to_list(&lst, item(5))
	for (lst != nil) {
			fmt.Println(lst.X)
			lst = lst.next
	}
}
