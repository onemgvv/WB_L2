package marketplace

type Account struct {
	login   string
	code    int
	balance int
}

func NewAccount(login string, code, balance int) *Account {
	return &Account{login, code, balance}
}

func (a *Account) Authorize(login string, code int) bool {
	if a.login == login && code == a.code {
		return true
	}

	return false
}
