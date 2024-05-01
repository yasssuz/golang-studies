package main

import "fmt"

type notifier interface {
	notify(string)
}

type User struct {
	name  string
	email string
}

type Admin struct {
	power int
	User  // // Embed item -> creates Item key based on Item struct, however, it also lifts up all Item keys to the
	// top level of Admin. This means that, you can access Admin.name instead of Admin.User.name.
	// An issue tho, is that if the struct has a conflicting key, the go compiler will not notify you of that and
	// this can create bugs.

	// A general guideline for Embedding is that, you should most of the times do it for behavior rather than
	// concrete data. Is cheaper to duplicate than to have wrong abstractions.

	// User User -> you could also do this way, but you'd lose the features above mentioned.
}

func (user User) notify(m string) {
	fmt.Printf("User %#v notified %#v", user.email, m)
}

func sendNotification(m string, n notifier) {
	n.notify(m)
}

func main() {
	// u := Admin{5, "Karim", "test@gmail.com"}
	u := Admin{
		power: 5,
		// Above field is required during construction
		// Below fields are optional during construction
		User: User{"Karim", "test@gmail.com"},
	}

	// Did you notice that we are sending admin, however, admin can still do admin.notify() ?
	sendNotification("test", u)
}
