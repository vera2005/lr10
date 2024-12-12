package usecase

func (u *Usecase) FetchQuery() (string, error) {
	msg, err := u.p.SelectCount()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetQuery(v float32) error {
	err := u.p.InsertCount(v)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) ChangeQuery(v float32) error {
	err := u.p.UpdateCount(v)
	if err != nil {
		return err
	}

	return nil
}
