package transaction

func (t *Transaction) Get() error {
	if err := t.dbGet(); err != nil {
		return err
	}

	return nil
}
