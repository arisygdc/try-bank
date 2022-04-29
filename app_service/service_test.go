package appservice

import (
	"context"
	"testing"
	"time"
	"try-bank/app_service/account"
	"try-bank/app_service/company"
	virtualaccount "try-bank/app_service/virtual_account"
	"try-bank/config"
	"try-bank/database"
	"try-bank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// param Seeder for unit testing
func getRepository(ctx context.Context) (database.IRepository, error) {
	env := config.Environment{
		DBDriver: "postgres",
		DBSource: "postgresql://postgresTest:secret@localhost:5432/bank?sslmode=disable",
	}

	return database.NewRepository(env, ctx)

}

func getRegisterClientParam(accountType_id uuid.UUID) []account.CreateUserParam {
	return []account.CreateUserParam{
		{
			Firstname:   "arisy",
			Lastname:    "musyafa'",
			Email:       "arisy@gdc.com",
			Phone:       "081217827013",
			Pin:         "025361",
			Birth:       time.Now(),
			AccountType: accountType_id,
		},
	}
}

func getRegisterCompanyParam() []company.RegisterCompanyParam {
	var outputParam = []company.RegisterCompanyParam{
		{
			PublicInfo_comp: company.PublicInfo_comp{
				Name:  "company companyan",
				Email: "mail@company.com",
				Phone: "081638293643",
			},
			Pin:     "083682",
			Deposit: 30000000,
		},
	}
	return outputParam
}

func TestUserAccount(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repos, err := getRepository(ctx)
	assert.Nil(t, err)

	svc := account.New(repos)
	accountType, err := svc.GetAccountType(ctx, account.LevelSilver)
	assert.Nil(t, err, "cannot get account type")

	register := getRegisterClientParam(accountType.ID)

	for _, v := range register {
		_, err = svc.CreateCustomerAccount(ctx, v)
		assert.Nil(t, err, "error create account")
	}
}

func TestCreateCompany(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repos, err := getRepository(ctx)
	assert.Nil(t, err)

	register := getRegisterCompanyParam()

	svc := company.New(repos)
	for _, v := range register {
		_, err = svc.CreateCompanyAccount(ctx, v)
		assert.Nil(t, err)
	}
}

func TestVirtualaccount(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo, err := getRepository(ctx)

	assert.Nil(t, err)

	svc := virtualaccount.New(repo)

	ca, err := repo.Query().TestGetAllCompaniesAccount(ctx)
	assert.Nil(t, err, err)

	_, err = svc.Register(ctx, uuid.New(), "http://kon.trol")
	assert.Error(t, err)
	for _, v := range ca {
		_, err := svc.Register(ctx, v.CompanyID, "http://"+util.RandString(8))
		assert.Nil(t, err, err)
	}
}
