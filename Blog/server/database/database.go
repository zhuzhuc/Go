package database

import (
	"log"
	"os"
	"time"

	"github.com/zhuzhuc/blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	// 从环境变量获取数据库配置
	user := os.Getenv("db_user")
	if user == "" {
		user = "root" // 默认用户名
	}
	log.Printf("Database user: %s", user)

	password := os.Getenv("db_password")
	// 密码可以为空，但我们记录是否设置了密码
	if password == "" {
		log.Println("Warning: No database password set")
	} else {
		log.Println("Database password is set")
	}

	dbname := os.Getenv("db_name")
	if dbname == "" {
		dbname = "blog" // 默认数据库名
	}
	log.Printf("Database name: %s", dbname)

	// 获取数据库主机地址，默认为 localhost
	dbhost := os.Getenv("db_host")
	if dbhost == "" {
		dbhost = "127.0.0.1" // 默认主机地址
	}
	log.Printf("Database host: %s", dbhost)

	// 获取数据库端口，默认为 3306
	dbport := os.Getenv("db_port")
	if dbport == "" {
		dbport = "3306" // 默认端口
	}
	log.Printf("Database port: %s", dbport)

	// 打印所有环境变量，帮助调试
	log.Println("Environment variables:")
	for _, env := range os.Environ() {
		log.Println(env)
	}

	// 构建 MySQL 的 DSN
	dsn := user + ":" + password + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	log.Printf("Attempting to connect to MySQL database: %s@tcp(%s:%s)/%s", user, dbhost, dbport, dbname)

	// Configure custom logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level (logger.Info shows all SQL)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			Colorful:                  true,        // Enable color
		},
	)

	// 尝试连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("Error connecting to MySQL: %v", err)
		log.Println("Please check your MySQL configuration:")
		log.Println("1. Make sure MySQL server is running")
		log.Println("2. Check if the username and password are correct")
		log.Println("3. Verify that the 'blog' database exists")
		log.Println("You can create the database with: CREATE DATABASE blog;")

		// 如果无法连接到MySQL，可以考虑使用SQLite作为备选
		log.Println("Attempting to use SQLite as fallback...")

		// 使用SQLite作为备选
		sqliteDB, sqliteErr := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
			Logger: newLogger,
		})

		if sqliteErr != nil {
			// 如果SQLite也失败，则终止程序
			panic("Failed to connect to both MySQL and SQLite: " + err.Error() + ", " + sqliteErr.Error())
		}

		log.Println("Connected to SQLite database successfully")
		db = sqliteDB
	} else {
		log.Println("MySQL database connected successfully")
	}

	// 自动迁移模型
	migErr := db.AutoMigrate(new(model.Blog), new(model.User))
	if migErr != nil {
		log.Printf("Error during auto migration: %v", migErr)
	} else {
		log.Println("Database schema migrated successfully")
	}

	DBConnection = db
}
