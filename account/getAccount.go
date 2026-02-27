package account

func GetAccount(userId int) (*Account, error) {

	a, err := dbGetAccount(userId)
	if err != nil {
		return nil, err
	}

	return a, nil

}
