package repository

// // CreateBattleResult recieve entity.BattleResult and save to DB
// func CreateBattleResult(b model.BattleResult) {
// 	db := database.Db()

// 	if !haveBattleResult(db, &b) {
// 		fmt.Println("Saving BattleResult to DB")
// 		if err := db.Create(&b).Error; err != nil {
// 			fmt.Println("Error saving battleresult to DB: ", err)
// 		}
// 	}
// }
