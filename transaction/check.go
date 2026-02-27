package transaction

func (t *Transaction) Check() (bool, error) {

	f, err := t.dbCheck()
	if err != nil {
		return false, err
	}

	return f, nil
}
