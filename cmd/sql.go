package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"projection/internal/sql2struct"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql转换和处理",
	Long: "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use: "struct",
	Short: "sql转换",
	Long: "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{DBType: dbType, Host: host, UserName: username, PassWord: password, Charset: charset}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.connect err:%v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.connect err:%v", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns:= template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("dbModel.connect err:%v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "","请输入数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "","请输入数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "49.235.244.9:3306","请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4","请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "dbType", "", "mysql","请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "dbName", "", "","请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "tableName", "", "","请输入数据库表名")

}