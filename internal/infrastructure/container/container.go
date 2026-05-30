package container

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/infrastructure/mysql"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var v *viper.Viper

type (
	Container struct {
		Mysqldb        *gorm.DB
		Apps           *Apps
		AuthUsc        usecase.AuthUseCase
		UserUsc        usecase.UserUseCase
		StoreUsc       usecase.StoreUseCase
		AddressUsc     usecase.AddressUseCase
		CategoryUsc    usecase.CategoryUseCase
		ProductUsc     usecase.ProductUseCase
		TransactionUsc usecase.TransactionUseCase
	}

	Apps struct {
		Name      string `mapstructure:"name"`
		Host      string `mapstructure:"host"`
		Version   string `mapstructure:"version"`
		Address   string `mapstructure:"address"`
		HttpPort  int    `mapstructure:"httpport"`
		SecretJwt string `mapstructure:"secretJwt"`
	}
)

func loadEnv() {
	projectDirName := "go-example-cruid"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	v.SetConfigFile(string(rootPath) + string(os.PathSeparator) + ".env")
	v.SetConfigType("env")
}

func init() {
	v = viper.New()

	v.AutomaticEnv()
	loadEnv()

	path, err := os.Executable()
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("os.Executable panic : %s", err.Error()), err)
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprintf("failed read config : %s", err.Error()), err)
	}

	helper.Logger(helper.LoggerLevelInfo, "Succeed read configuration file", nil)
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, fmt.Sprint("Error when unmarshal configuration file : ", err.Error()), err)
	}
	helper.Logger(helper.LoggerLevelInfo, "Succeed when unmarshal configuration file", err)
	return
}

func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	utils.InitJWT(apps.SecretJwt)
	mysqldb := mysql.DatabaseInit(v)

	// Initialize repositories
	userRepo := repository.NewUserRepository(mysqldb)
	storeRepo := repository.NewStoreRepository(mysqldb)
	addressRepo := repository.NewAddressRepository(mysqldb)
	categoryRepo := repository.NewCategoryRepository(mysqldb)
	productRepo := repository.NewProductRepository(mysqldb)
	transactionRepo := repository.NewTransactionRepository(mysqldb)
	productLogRepo := repository.NewProductLogRepository(mysqldb)

	// Initialize usecases
	authUsc := usecase.NewAuthUseCase(userRepo, storeRepo)
	userUsc := usecase.NewUserUseCase(userRepo)
	storeUsc := usecase.NewStoreUseCase(storeRepo, userRepo)
	addressUsc := usecase.NewAddressUseCase(addressRepo, userRepo)
	categoryUsc := usecase.NewCategoryUseCase(categoryRepo)
	productUsc := usecase.NewProductUseCase(productRepo, storeRepo, categoryRepo)
	transactionUsc := usecase.NewTransactionUseCase(transactionRepo, productLogRepo, productRepo, addressRepo, userRepo)

	return &Container{
		Apps:           &apps,
		Mysqldb:        mysqldb,
		AuthUsc:        authUsc,
		UserUsc:        userUsc,
		StoreUsc:       storeUsc,
		AddressUsc:     addressUsc,
		CategoryUsc:    categoryUsc,
		ProductUsc:     productUsc,
		TransactionUsc: transactionUsc,
	}
}
