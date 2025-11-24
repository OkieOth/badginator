package config

// Attention, this file is generated. Manual changes will be
// overwritten with the next run of the codegen
type LoggingLogLevel string

const (
  LoggingLogLevel_debug LoggingLogLevel = "debug"
  LoggingLogLevel_info LoggingLogLevel = "info"
  LoggingLogLevel_error LoggingLogLevel = "error"
)
type Logging struct {
  LogLevel LoggingLogLevel`json:"logLevel,omitempty"`
  Output string`json:"output,omitempty"`
}

func NewLogging() Logging {
  var ret Logging
  ret.LogLevel = LoggingLogLevel_info
  ret.Output = "stdout"
  return ret
}
//  Container object for all configuration of the 'serve' sub command
type Server struct {
  Port int`json:"port,omitempty"`
}

func NewServer() Server {
  var ret Server
  ret.Port = 8080
  return ret
}
//  Container object for all storage configurations
type Storage struct {
  Sqlite Sqlite
}

func NewStorage() Storage {
  var ret Storage
  ret.Sqlite = NewSqlite()
  return ret
}
//  Configurations for the sqlite storage backend
type Sqlite struct {
  File string`json:"file,omitempty"`
}

func NewSqlite() Sqlite {
  var ret Sqlite
  return ret
}
//  The schema defines the model used by the configuration file of the program
type AppConfig struct {
  Server Server
  Storage Storage
  Logging Logging
}

func NewAppConfig() AppConfig {
  var ret AppConfig
  ret.Server = NewServer()
  ret.Storage = NewStorage()
  ret.Logging = NewLogging()
  return ret
}
