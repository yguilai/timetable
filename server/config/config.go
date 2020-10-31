package config

type Server struct {
	System  System  `json:"system" yaml:"system"`
	Mysql   Mysql   `json:"mysql" yaml:"mysql"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	Log     Log     `json:"log" yaml:"log"`
	Jwt     Jwt     `json:"jwt" yaml:"jwt"`
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	Remote  Remote  `json:"remote" yaml:"remote"`
}

type System struct {
	Port string `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
	Name string `json:"name" yaml:"name"`
}

type Mysql struct {
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	Path         string `json:"path" yaml:"path"`
	Dbname       string `json:"dbname" yaml:"dbName"`
	Config       string `json:"config" yaml:"config"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
	LogMode      bool   `json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
	Enable   bool   `json:"enable" yaml:"enable"`
}

type Log struct {
	Encoder    string `json:"encoder" yaml:"encoder"`
	Level      string `json:"level" yaml:"level"`
	Path       string `json:"path" yaml:"path"`
	MaxSize    int    `json:"maxSize" yaml:"maxSize"`
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"`
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

type Jwt struct {
	SigningKey string `json:"signingKey" yaml:"signingKey"`
}

type Captcha struct {
	Length int `json:"length" yaml:"length"`
	Width  int `json:"width" yaml:"width"`
	Height int `json:"height" yaml:"height"`
}

type Remote struct {
	TimetableApi string `json:"timetableApi" yaml:"timetableApi"`
}
