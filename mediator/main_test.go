package mediator

import "testing"


func TestMediator(t *testing.T) {
    mediator := NewMediator()
    user := &User{Name: "John", Age: 30}
    mediator.AddUser(user)

    if mediator.GetUser("John").Name != "John" {
        t.Error("Expected user name to be John")
    }

    mediator.UpdateUser("John", "Jane")
    if mediator.GetUser("Jane").Name != "Jane" {
        t.Error("Expected user name to be Jane")

    }
}
