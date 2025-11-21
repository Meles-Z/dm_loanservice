package goconf

type DefaultConfig struct {
	Debug            bool             `yaml:"debug"`
	App              *App             `yaml:"app"`
	MigrationsFolder string           `yaml:"migrations_folder"`
	Rest             *Rest            `yaml:"rest"`
	Grpc             *Grpc            `yaml:"grpc"`
	Cors             *Cors            `yaml:"cors"`
	Postgres         *Postgres        `yaml:"postgres"`
	Redis            *Redis           `yaml:"redis"`
	Jwt              *Jwt             `yaml:"jwt"`
	MailTrap         *MailTrap        `yaml:"mailtrap"`
	DocService       *Docservice      `yaml:"docservice"`
	CustomerService  *CustomerService `yaml:"customerservice"`
	SmsConfig        *SmsConfig       `yaml:"smsconfig"`
	Gcs              *Gcs             `yaml:"gcs"`
	Salt             string           `yaml:"salt"`
	Provider         *ProviderDetail  `yaml:"provider"`
	Auth             *AuthTemplate    `yaml:"auth"`
	FeatureFlag      *FeatureFlag     `yaml:"featureFlag"`
	MFASecret        *MFASecret       `yaml:"mfasecret"`
}

type App struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}

type Rest struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
	IdleTimeout  int `yaml:"idle_timeout"`
}

type Grpc struct {
	Port             int    `yaml:"port"`
	TlsEnabled       bool   `yaml:"tls_enabled"`
	UnaryInterceptor string `yaml:"unary_interceptor"`
	ServerCert       string `yaml:"server_cert"`
	ServerKey        string `yaml:"server_key"`
}

type Cors struct {
	Hosts []string `yaml:"hosts"`
}

type Postgres struct {
	Master *PostgresDB `yaml:"master"`
	Slave  *PostgresDB `yaml:"slave"`
}

type PostgresDB struct {
	Addr     string `yaml:"addr"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Net      string `yaml:"net"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Driver   string `yaml:"driver"`
	Moc      int    `yaml:"moc"`
	Mic      int    `yaml:"mic"`
	Timeout  int    `yaml:"timeout"`
}

type Redis struct {
	Address  string     `yaml:"address"`
	Password string     `yaml:"password"`
	Otp      *RedisTime `yaml:"otp"`
}

type RedisTime struct {
	LockRequestTime  int `yaml:"lockRequestTime"`
	LockNotMatchTime int `yaml:"lockNotMatchTime"`
	RegisterTTL      int `yaml:"registerTTL"`
}

type Jwt struct {
	KeyAccess  string `yaml:"key_access"`
	KeyRefresh string `yaml:"key_refresh"`
	Algorithm  string `yaml:"algorithm"`
	Expire     string `yaml:"expire"`
}

type MailTrap struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Addr     string `yaml:"addr"`
	Host     string `yaml:"host"`
}

type Docservice struct {
	Address        string `yaml:"address"`
	PathCredential string `yaml:"path_credential"`
	ServerName     string `yaml:"servername"`
	Timeout        int    `yaml:"timeout"`
}

type CustomerService struct {
	Address        string `yaml:"address"`
	PathCredential string `yaml:"path_credential"`
	ServerName     string `yaml:"servername"`
	Timeout        int    `yaml:"timeout"`
}

type SmsConfig struct {
	UserId   string `yaml:"userid"`
	Template string `yaml:"template"`
	Version  string `yaml:"version"`
}

type Gcs struct {
	CredentialPath     string `yaml:"credential_path"`
	BucketNameProfiles string `yaml:"bucket_name_profiles"`
}

type ProviderDetail struct {
	Google *ProviderConf `yaml:"google"`
}
type MFASecret struct {
	MFASecret string `yaml:"mfasecret"`
}
type ProviderConf struct {
	Key         string   `yaml:"key"`
	Secret      string   `yaml:"secret"`
	CallbackUrl string   `yaml:"callbackURL"`
	Scope       []string `yaml:"scope"`
}

type AuthTemplate struct {
	DocsApiKey           string              `yaml:"DOCS_API_KEY"`
	ForgotPassword       *AuthTemplateDetail `yaml:"forgotPassword"`
	OtpPassword          *AuthTemplateDetail `yaml:"otpPassword"`
	SuccessResetPassword *AuthTemplateDetail `yaml:"successResetPassword"`
}

type AuthTemplateDetail struct {
	Route string                   `yaml:"route"`
	Email *AuthTemplateDetailEmail `yaml:"email"`
	Sms   *AuthTemplateDetailEmail `yaml:"sms"`
}

type AuthTemplateDetailEmail struct {
	TemplateId      string `yaml:"templateId"`
	TemplateVersion string `yaml:"templateVersion"`
	UserId          string `yaml:"userId"`
}

type FeatureFlag struct {
	ForgotPasswordValidateMobile *Flag `yaml:"forgotPasswordValidateMobile"`
}

func (f *FeatureFlag) Exist() (exist bool) {
	return f != nil
}

type Flag bool

func (f *Flag) Valid() (valid bool) {
	return f != nil && *f == true
}
