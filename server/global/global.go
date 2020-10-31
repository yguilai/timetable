package global

import (
	"ccsu/config"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Viper         *viper.Viper
	Config        config.Server
	Validator     *v.Validate
	ValidateTrans ut.Translator
	Log           *zap.Logger
	DB            *gorm.DB
	Redis         *redis.Client
)
