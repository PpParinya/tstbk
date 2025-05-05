package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/database"
	"github.com/tst/backend/model/entity"
)

//*******************************************************//

// type Account struct {
// 	UserID       int        `gorm:"column:UserID"   json:"userid"`
// 	FullName     string     `gorm:"column:FullName"    json:"fullname"`
// 	ParentUserID *int       `gorm:"column:ParentUserID" json:"-"`
// 	Children     []*Account `json:"children,omitempty" gorm:"-"`
// }

func GetParentUserIDByUserID(UserID string) string {

	var users entity.Users

	err := database.DB.Debug().Table("users").Find(&users, "UserID = ?", UserID).Error
	if err != nil || users.UserID == 0 {
		log.Fatal("❌ query error:", err)
	}
	ParentUserID := strconv.FormatInt(*users.ParentUserID, 10)
	return ParentUserID
}

func GetAccountsByUserID(ctx *fiber.Ctx) error {

	userIDStr := ctx.Locals("UserID").(string)
	rootByParentUserID := GetParentUserIDByUserID(userIDStr)

	UserID := rootByParentUserID

	// 🧠 Query เฉพาะ ID userID และลูกทุกระดับ
	rawSQL := `
	  WITH AccountTree AS (
		  SELECT *
		  FROM Users
		  WHERE UserID = ` + UserID + `
  
		  UNION ALL
  
		  SELECT c.*
		  FROM Users c
		  INNER JOIN AccountTree a ON c.ParentUserID = a.UserID
	  )
	  SELECT * FROM AccountTree;
	  `

	var accounts []*entity.Users
	err := database.DB.Debug().Raw(rawSQL).Scan(&accounts).Error
	if err != nil {
		log.Fatal("❌ query error:", err)
	}

	if len(accounts) == 0 {
		log.Fatal("⚠️ No accounts found")
	}

	return ctx.JSON(accounts)

}

func GetAccountsTreeByUserID(ctx *fiber.Ctx) error {

	userIDStr := ctx.Locals("UserID").(string)
	rootByParentUserID := GetParentUserIDByUserID(userIDStr)

	UserID := rootByParentUserID

	// 🧠 Query เฉพาะ ID userID และลูกทุกระดับ
	rawSQL := `
	  WITH AccountTree AS (
		  SELECT *
		  FROM Users
		  WHERE UserID = ` + UserID + `
  
		  UNION ALL
  
		  SELECT c.*
		  FROM Users c
		  INNER JOIN AccountTree a ON c.ParentUserID = a.UserID
	  )
	  SELECT * FROM AccountTree;
	  `

	var accounts []*entity.Users
	err := database.DB.Debug().Raw(rawSQL).Scan(&accounts).Error
	if err != nil {
		log.Fatal("❌ query error:", err)
	}

	if len(accounts) == 0 {
		log.Fatal("⚠️ No accounts found")
	}

	userIDStrs := fmt.Sprintf("%v", ctx.Locals("UserID"))
	userIDs, err := strconv.ParseInt(userIDStrs, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID format",
		})
	}

	// 🌳 สร้างโครงสร้าง Tree จากผลลัพธ์
	root := buildTree(accounts, userIDs)
	if root == nil {
		log.Fatal("❌ root node not found")
	}

	return ctx.JSON(root)

}

// buildTree สร้างต้นไม้จาก flat list ของบัญชี
func buildTree(accounts []*entity.Users, rootID int64) *entity.Users {
	accountMap := make(map[int64]*entity.Users)

	for _, acc := range accounts {
		accountMap[acc.UserID] = acc
	}

	var root *entity.Users
	for _, acc := range accounts {
		if acc.UserID == rootID {
			root = acc
		}

		if acc.ParentUserID != nil {
			parent, ok := accountMap[*acc.ParentUserID]
			if ok && parent != nil {
				parent.Children = append(parent.Children, acc)
			} else {
				log.Printf("⚠️ Missing parent %d for account %d", *acc.ParentUserID, acc.UserID)
			}
		}
	}

	return root
}

type UserInfos struct {
	UserID int64 `gorm:"column:UserID"   json:"UserID"`
}
