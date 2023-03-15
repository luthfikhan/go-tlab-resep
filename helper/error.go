package helper

import "gorm.io/gorm"

func CheckErrorToCommitOrRollback(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	} else {
		tx.Commit()
	}
}

func PanifIfError(err error) {
	if err != nil {
		panic(err)
	}
}
