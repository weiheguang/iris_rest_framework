package rftests

// sql 和 model放在一起
const (
	TEST_TABLE_USER_SQL = `create table if not exists user (
		id int(11) not null auto_increment,
		name varchar(255) not null default '',
		age int(11) not null default 0,
		primary key (id)
	) engine=innodb default charset=utf8mb4;`

	INSERT_USER_SQL = `insert into user (name, age) values (?, ?);`
)

// 测试 User 结构体
type User struct {
	Id   int    `json:"id" gorm:"column:id;type:int(11);not null;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null;default:''"`
	Age  int    `json:"age" gorm:"column:age;type:int(11);not null;default:0"`
}
