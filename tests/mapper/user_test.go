package mapper_test

import (
	"encoding/json"
	"fmt"
	"testing"

	gormAgent "github.com/OJ-lab/oj-lab-services/src/core/agent/gorm"
	"github.com/OJ-lab/oj-lab-services/src/service/mapper"
	"github.com/OJ-lab/oj-lab-services/src/service/model"
)

func TestUserMapper(t *testing.T) {
	db := gormAgent.GetDefaultDB()
	user := model.User{
		Account:  "test",
		Password: func() *string { s := "test"; return &s }(),
		Roles:    []*model.Role{{Name: "test"}},
	}
	err := mapper.CreateUser(db, user)
	if err != nil {
		t.Error(err)
	}

	dbUser, err := mapper.GetUser(db, user.Account)
	if err != nil {
		t.Error(err)
	}
	userJson, err := json.MarshalIndent(dbUser, "", "\t")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", string(userJson))

	dbPublicUser, err := mapper.GetPublicUser(db, user.Account)
	if err != nil {
		t.Error(err)
	}
	publicUserJson, err := json.MarshalIndent(dbPublicUser, "", "\t")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", string(publicUserJson))

	err = mapper.DeleteUser(db, user)
	if err != nil {
		t.Error(err)
	}

}
