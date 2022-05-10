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

// TODO
// - Implement test table

// param Seeder for unit testing
func getRepository(ctx context.Context) (database.IRepository, error) {
	env := config.Environment{
		DBDriver: "postgres",
		DBSource: "postgresql://postgresTest:secret@localhost:5432/bank?sslmode=disable",
	}

	return database.NewRepository(env)

}

func getRegisterClientParam(accountType_id uuid.UUID) []account.CreateCostumerParam {
	return []account.CreateCostumerParam{
		{
			Firstname:   "si",
			Lastname:    "pitung'",
			Email:       "si@pitung.com",
			Phone:       "081217843623",
			Pin:         "025361",
			TopUp:       1000000,
			Birth:       time.Now(),
			AccountType: accountType_id,
		}, {
			Firstname:   "si",
			Lastname:    "gatel'",
			Email:       "si@gatel.com",
			Phone:       "0812132435346",
			Pin:         "025361",
			TopUp:       1000000,
			Birth:       time.Now(),
			AccountType: accountType_id,
		}, {
			Firstname:   "mbak",
			Lastname:    "yeyen'",
			Email:       "mbakyeyen@gatel.com",
			Phone:       "0812132433436",
			Pin:         "025361",
			TopUp:       1000000,
			Birth:       time.Now(),
			AccountType: accountType_id,
		},
	}
}

func getRegisterCompanyParam() []company.RegisterCompanyParam {
	var outputParam = []company.RegisterCompanyParam{
		{
			PublicInfo_company: company.PublicInfo_company{
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

func TestTransfer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo, err := getRepository(ctx)

	assert.Nil(t, err)

	svc := account.New(repo)
	customers, err := repo.Query().TestGetAllAccount(ctx)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(customers), 2)
	for i := 0; i < len(customers)-1; i++ {
		err := svc.Transfer(ctx, customers[i].WalletID, customers[i+1].WalletID, float64(10000+i))
		assert.Nil(t, err)
	}
}
