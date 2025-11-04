package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds the application configuration
type Config struct {
	App       AppConfig       `yaml:"app"`
	Server    ServerConfig    `yaml:"server"`
	CORS      CORSConfig      `yaml:"cors"`
	JWT       JWTConfig       `yaml:"jwt"`
	OAuth     OAuthConfig     `yaml:"oauth"`
	Jitsi     JitsiConfig     `yaml:"jitsi"`
	AWS       AWSConfig       `yaml:"aws"`
	Logging   LoggingConfig   `yaml:"logging"`
	RateLimit RateLimitConfig `yaml:"rate_limiting"`
	WebSocket WebSocketConfig `yaml:"websocket"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
}

type ServerConfig struct {
	ReadTimeout    time.Duration `yaml:"read_timeout"`
	WriteTimeout   time.Duration `yaml:"write_timeout"`
	IdleTimeout    time.Duration `yaml:"idle_timeout"`
	MaxHeaderBytes int           `yaml:"max_header_bytes"`
}

type CORSConfig struct {
	AllowedOrigins   []string `yaml:"allowed_origins"`
	AllowedMethods   []string `yaml:"allowed_methods"`
	AllowedHeaders   []string `yaml:"allowed_headers"`
	ExposedHeaders   []string `yaml:"expose_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	MaxAge           int      `yaml:"max_age"`
}

type JWTConfig struct {
	Secret             string        `yaml:"secret"`
	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
	Issuer             string        `yaml:"issuer"`
}

type OAuthConfig struct {
	Google GoogleOAuthConfig `yaml:"google"`
}

type GoogleOAuthConfig struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURL  string   `yaml:"redirect_url"`
	Scopes       []string `yaml:"scopes"`
}

type JitsiConfig struct {
	APIURL     string `yaml:"api_url"`
	AppID      string `yaml:"app_id"`
	AppSecret  string `yaml:"app_secret"`
	Domain     string `yaml:"domain"`
	RoomPrefix string `yaml:"room_prefix"`
}

type AWSConfig struct {
	Region string    `yaml:"region"`
	S3     S3Config  `yaml:"s3"`
}

type S3Config struct {
	Bucket           string `yaml:"bucket"`
	RecordingsPrefix string `yaml:"recordings_prefix"`
	MaxFileSize      int64  `yaml:"max_file_size"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Output string `yaml:"output"`
}

type RateLimitConfig struct {
	Enabled            bool `yaml:"enabled"`
	RequestsPerSecond  int  `yaml:"requests_per_second"`
	Burst              int  `yaml:"burst"`
}

type WebSocketConfig struct {
	ReadBufferSize  int           `yaml:"read_buffer_size"`
	WriteBufferSize int           `yaml:"write_buffer_size"`
	MaxMessageSize  int64         `yaml:"max_message_size"`
	PingInterval    time.Duration `yaml:"ping_interval"`
	PongTimeout     time.Duration `yaml:"pong_timeout"`
}

// PostgresConfig holds PostgreSQL configuration
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	Pool     PoolConfig `yaml:"pool"`
}

type PoolConfig struct {
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string            `yaml:"host"`
	Port     int               `yaml:"port"`
	Password string            `yaml:"password"`
	DB       int               `yaml:"db"`
	Pool     RedisPoolConfig   `yaml:"pool"`
	Timeouts TimeoutsConfig    `yaml:"timeouts"`
	Prefixes map[string]string `yaml:"prefixes"`
	TTL      map[string]int    `yaml:"ttl"`
}

type RedisPoolConfig struct {
	MaxActive   int           `yaml:"max_active"`
	MaxIdle     int           `yaml:"max_idle"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
	Wait        bool          `yaml:"wait"`
}

type TimeoutsConfig struct {
	Dial  time.Duration `yaml:"dial"`
	Read  time.Duration `yaml:"read"`
	Write time.Duration `yaml:"write"`
}

// Load loads configuration from YAML files
func Load(appConfigPath, postgresConfigPath, redisConfigPath string) (*Config, *PostgresConfig, *RedisConfig, error) {
	// Load app config
	appConfig, err := loadAppConfig(appConfigPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to load app config: %w", err)
	}

	// Load postgres config
	postgresConfig, err := loadPostgresConfig(postgresConfigPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to load postgres config: %w", err)
	}

	// Load redis config
	redisConfig, err := loadRedisConfig(redisConfigPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to load redis config: %w", err)
	}

	// Expand environment variables
	expandEnvVars(appConfig)
	expandPostgresEnvVars(postgresConfig)
	expandRedisEnvVars(redisConfig)

	return appConfig, postgresConfig, redisConfig, nil
}

func loadAppConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func loadPostgresConfig(path string) (*PostgresConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		Postgres PostgresConfig `yaml:"postgres"`
	}
	if err := yaml.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}

	return &wrapper.Postgres, nil
}

func loadRedisConfig(path string) (*RedisConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		Redis RedisConfig `yaml:"redis"`
	}
	if err := yaml.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}

	return &wrapper.Redis, nil
}

func expandEnvVars(config *Config) {
	config.JWT.Secret = os.ExpandEnv(config.JWT.Secret)
	config.OAuth.Google.ClientID = os.ExpandEnv(config.OAuth.Google.ClientID)
	config.OAuth.Google.ClientSecret = os.ExpandEnv(config.OAuth.Google.ClientSecret)
	config.OAuth.Google.RedirectURL = os.ExpandEnv(config.OAuth.Google.RedirectURL)
	config.Jitsi.APIURL = os.ExpandEnv(config.Jitsi.APIURL)
	config.Jitsi.AppID = os.ExpandEnv(config.Jitsi.AppID)
	config.Jitsi.AppSecret = os.ExpandEnv(config.Jitsi.AppSecret)
	config.Jitsi.Domain = os.ExpandEnv(config.Jitsi.Domain)
	config.AWS.Region = os.ExpandEnv(config.AWS.Region)
	config.AWS.S3.Bucket = os.ExpandEnv(config.AWS.S3.Bucket)
}

func expandPostgresEnvVars(config *PostgresConfig) {
	config.Host = os.ExpandEnv(config.Host)
	config.User = os.ExpandEnv(config.User)
	config.Password = os.ExpandEnv(config.Password)
	config.DBName = os.ExpandEnv(config.DBName)
	config.SSLMode = os.ExpandEnv(config.SSLMode)
}

func expandRedisEnvVars(config *RedisConfig) {
	config.Host = os.ExpandEnv(config.Host)
	config.Password = os.ExpandEnv(config.Password)
}

// GetDSN returns PostgreSQL connection string
func (c *PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

// GetAddr returns Redis address
func (c *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
