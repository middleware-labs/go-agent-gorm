# GORM instrumentation

## Installation

```shell
go get github.com/middleware-labs/go-agent-gorm
```

## Usage

To instrument GORM, you need to install the plugin provided by otelgorm:

```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
     mw_gorm "github.com/middleware-labs/go-agent-gorm/gorm"
     "github.com/middleware-labs/golang-apm/tracker"
)

db, err := gorm.Open(mysql.Open("username:password@tcp(127.0.0.1:3306)/dbname"), &gorm.Config{})
if err != nil {
   panic(err)
}
if err := db.Use(mw_gorm.NewPlugin(mw_gorm.WithDBName("dbname"), mw_gorm.WithAttributes(tracker.String("db.system", "mysql")))); err != nil {
  panic(err)
}

```

And then use `db.WithContext(ctx)` to propagate the active span

```go
var num int
if err := db.WithContext(ctx).Raw("SELECT 42").Scan(&num).Error; err != nil {
	panic(err)
}
```
