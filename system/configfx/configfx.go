package configfx

import (
	"io/ioutil"
	"time"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

// ApplicationConfig ...
type ApplicationConfig struct {
	Address      string `yaml:"address"`
	AppName      string `yaml:"name"`
	PasswordSalt string `yaml:"salt"`
	TokenExpire  int    `yaml:"expire"`

	BodyLimit                    int           `yaml:"body_limit"`
	CaseSensitive                bool          `yaml:"case_sensitive"`
	CompressedFileSuffix         string        `yaml:"suffix"`
	Concurrency                  int           `yaml:"concurrency"`
	DisableDefaultContentType    bool          `yaml:"disable_content_type"`
	DisableDefaultDate           bool          `yaml:"disable_date"`
	DisableHeaderNormalizing     bool          `yaml:"disable_header_normalizing"`
	DisableKeepalive             bool          `yaml:"disable_keepalive"`
	DisablePreParseMultipartForm bool          `yaml:"disable_multipart_form"`
	DisableStartupMessage        bool          `yaml:"disable_startup_message"`
	ETag                         bool          `yaml:"etag"`
	EnableIPValidation           bool          `yaml:"enable_ipvalite"`
	EnablePrintRoutes            bool          `yaml:"enable_print_routes"`
	EnableTrustedProxyCheck      bool          `yaml:"enable_trusted_proxy"`
	GETOnly                      bool          `yaml:"get_only"`
	IdleTimeout                  time.Duration `yaml:"idle_timeout"`
	Immutable                    bool          `yaml:"immutable"`
	// JSONDecoder utils.JSONUnmarshal
	// JSONEncoder utils.JSONUnmarshal
	Network           string        `yaml:"network"`
	PassLocalsToViews bool          `yaml:"pass_locals_views"`
	Prefork           bool          `yaml:"prefork"`
	ProxyHeader       string        `yaml:"proxy_header"`
	ReadBufferSize    int           `yaml:"read_buffer_size"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	RequestMethods    []string      `yaml:"request_methods"`
	ServerHeader      string        `yaml:"server_header"`
	StreamRequestBody bool          `yaml:"stream_request_body"`
	StrictRouting     bool          `yaml:"strict_routing"`
	TrustedProxies    []string      `yaml:"trusted_proxies"`
	UnescapePath      bool          `yaml:"unescape_path"`
	WriteBufferSize   int           `yaml:"write_buffer_size"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Port     int    `yaml:"port"`
}

// Config ...
type Config struct {
	ApplicationConfig `yaml:"application"`
	DatabaseConfig    `yaml:"database"`
}

// ProvideConfig provides the standard configuration to fx
func ProvideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
