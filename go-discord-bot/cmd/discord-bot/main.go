package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cpcchengt/cpchain-discord-bot/internal/airdrop"
	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
// const (
// 	userName = "root"
// 	password = "123456"
// 	ip       = "127.0.0.1"
// 	port     = "3306"
// 	dbName   = "discordbot"
// )

// type discordUserInfo struct {
// 	ID       string `json:"id"`
// 	TOKEN    string `json:"token"`
// 	USERNAME string `json:"username"`
// }

// var discordUserInfos = []discordUserInfo{
// 	{ID: "1", TOKEN: "12345", USERNAME: "John"},
// 	{ID: "2", TOKEN: "67890", USERNAME: "Coltrane"},
// }

// func getInfos(c *gin.Context) {
// 	c.JSON(http.StatusOK, discordUserInfos)
// }

// func sendMsg(c *gin.Context) {
// 	dg, err := discordgo.New("Bot " + Token) // 创建机器人session
// 	if err != nil {
// 		fmt.Println("error creating Discord session,", err)
// 		return
// 	}
// 	dg.ChannelMessageSend("978216307126312960", "test rest api")
// }

// //Db数据库连接池
// var DB *sql.DB

// //注意方法名大写，就是public
// func InitDB() {
// 	fmt.Println("opon database fail")
// 	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
// 	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

// 	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
// 	DB, _ = sql.Open("mysql", path)
// 	//设置数据库最大连接数
// 	DB.SetConnMaxLifetime(100)
// 	//设置上数据库最大闲置连接数
// 	DB.SetMaxIdleConns(10)
// 	//验证连接
// 	if err := DB.Ping(); err != nil {
// 		fmt.Println("opon database fail")
// 		return
// 	}
// 	fmt.Println("connnect success")
// }

// Variables used for command line parameters
var (
	Token string
)

func init() { // 从外部导入token
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// router := gin.Default()
	// router.GET("/infos", getInfos)
	// router.GET("/sendmsg", sendMsg)
	// router.Run("localhost:8080")

	dg, err := discordgo.New("Bot " + Token) // 创建机器人session
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// dg.AddHandler(messageCreate) // 添加处理器
	dg.AddHandler(airdrop.AirdropApply)

	// dg.Identify.Intents = discordgo.IntentsGuildMessages // 文本消息处理

	// websocket
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// 接收退出信号
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// 关闭
	dg.Close()
}
