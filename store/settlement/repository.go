package settlement

func Create(settlement *Settlement) (*Settlement, error) {
	result := db.Create(settlement)
	return settlement, result.Error
}

func Save(settlement *Settlement) error {
	result := db.Create(settlement)
	return result.Error
}

func FindOne(id string) (*Settlement, error) {
	settlement := &Settlement{}
	result := db.First(settlement, "id = ?", id)
	return settlement, result.Error
}

func FindOneBy(settlement *Settlement) (*Settlement, error) {
	u := &Settlement{}
	result := db.Where(settlement).First(u)
	return u, result.Error
}

func FindMany(settlement *Settlement) (*[]Settlement, error) {
	settlements := &[]Settlement{}
	result := db.Where(settlement).Find(settlements)
	return settlements, result.Error
}

func Update(id string, settlement *Settlement) (*Settlement, error) {
	settlement.Id = id
	result := db.Updates(settlement)
	return settlement, result.Error
}

func Delete(id string) error {
	settlement := &Settlement{}
	settlement.Id = id
	result := db.Delete(&settlement)
	return result.Error
}

func CountMany(settlement *Settlement) (int64, error) {
	var total int64
	result := db.Where(settlement).Count(&total)
	return total, result.Error
}
