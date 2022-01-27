package payout

func Create(payout *Payout) (*Payout, error) {
	result := db.Create(payout)
	return payout, result.Error
}

func Save(payout *Payout) error {
	result := db.Create(payout)
	return result.Error
}

func FindOne(id string) (*Payout, error) {
	payout := &Payout{}
	result := db.First(payout, "id = ?", id)
	return payout, result.Error
}

func FindOneBy(payout *Payout) (*Payout, error) {
	u := &Payout{}
	result := db.Where(payout).First(u)
	return u, result.Error
}

func FindMany(payout *Payout) (*[]Payout, error) {
	payouts := &[]Payout{}
	result := db.Where(payout).Find(payouts)
	return payouts, result.Error
}

func Update(id string, payout *Payout) (*Payout, error) {
	payout.Id = id
	result := db.Updates(payout)
	return payout, result.Error
}

func Delete(id string) error {
	payout := &Payout{}
	payout.Id = id
	result := db.Delete(&payout)
	return result.Error
}

func CountMany(payout *Payout) (int64, error) {
	var total int64
	result := db.Where(payout).Count(&total)
	return total, result.Error
}
