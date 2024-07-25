package mediator

type User struct {
    Name string
    Age  int
}

func (u *User) UpdateName(name string) {
    u.Name = name
}

func (u *User) IncrementAge() {
    u.Age++
}

type Mediator struct {
    Users map[string]*User
}

func NewMediator() *Mediator {
    return &Mediator{Users: make(map[string]*User)}
}

func (m *Mediator) AddUser(user *User) {
    m.Users[user.Name] = user
}

func (m *Mediator) GetUser(name string) *User {
    return m.Users[name]
}

func (m *Mediator) UpdateUser(name, newName string) {
    if user, ok := m.Users[name]; ok {
        user.UpdateName(newName)
        m.Users[newName] = user
        delete(m.Users, name)
    }
}
