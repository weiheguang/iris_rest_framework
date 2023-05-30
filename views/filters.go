package views

const (
	// 精确查询
	EXACT = "exact"
	// 不区分大小写精确查询
	IEXACT = "iexact"
	// 包含查询
	CONTAINS = "contains"
	// 不区分大小写包含查询
	ICONTAINS = "icontains"
	// 小于查询
	LT = "lt"
	// 小于等于查询
	LTE = "lte"
	// 大于查询
	GT = "gt"
	// 大于等于查询
	GTE = "gte"
	// 包含查询
	IN = "in"
	// 开头查询
	STARTSWITH = "startswith"
	// 结尾查询
	ENDSWITH = "endswith"
	// 年查询
	YEAR = "year"
	// 月查询
	MONTH = "month"
	// 日查询
	DAY = "day"
	// 时查询
	HOUR = "hour"
	// 分查询
	MINUTE = "minute"
	// 秒查询
	SECOND = "second"
	// 是否为空查询
	ISNULL = "isnull"
	// 搜索查询
	SEARCH = "search"
	// 正则查询
	REGEX = "regex"
	// 不区分大小写正则查询
	IREGEX = "iregex"
)

// var QUERY_CONDITIO = []string{EXACT, IEXACT, CONTAINS, ICONTAINS, GT, GTE, LT, LTE, IN}

// 过滤字段
// type FilterField struct {
// 	// 	# get user with id 5
// 	// 	?id=5
// 	// 	# get user with id 5 and 6
// 	// 	?id=5,6
// 	// 	# get user with id greater than 5
// 	// 	?id__gt=5
// 	// 	# get user with id greater than or equal to 5
// 	// 	?id__gte=5
// 	// 	# get user with id less than 5
// 	// 	?id__lt=5
// 	// 	# get user with id less than or equal to 5
// 	// 	?id__lte=5
// 	lookup string // 查询条件
// 	expr   string // 查询表达式
// 	value  string // 查询值
// }

// // 过滤器
// type Filter struct {
// 	db     *gorm.DB
// 	model  interface{}
// 	fields []FilterField
// }

// // 新建过滤器
// func NewFilter(db *gorm.DB) *Filter {
// 	return &Filter{
// 		db: db,
// 	}
// }

// // 设置过滤字段, 默认为 exact
// /*
// 	参数类型:
// 		id=5 (id_exact=5)
// 		id__gt=5
// 		id__gte=5
// 		id__lt=5
// 		id__lte=5
// 		id__in=5,6,7
// 		id__contains=5
// 		id__icontains=5
// 		id__startswith=5
// 		id__istartswith=5
// 		id__endswith=5
// 		id__iendswith=5
// 		id__year=2020
// 		id__month=2020-01
// 		id__day=2020-01-01
// 		id__hour=2020-01-01 12
// 		id__minute=2020-01-01 12:00
// 		id__second=2020-01-01 12:00:00
// 		id__isnull=true
// 		id__search=5
// 		id__regex=5
// 		id__iregex=5
// */

// func (f *Filter) AddField(key string, value string) {
// 	field := FilterField{}
// 	// 如果 key 中包含 __, 则按照 __ 分割, 否则默认为 exact
// 	if strings.Contains(key, "__") {
// 		field.lookup = strings.Split(key, "__")[0]
// 		field.expr = strings.Split(key, "__")[1]
// 		field.value = value
// 	} else {
// 		field.lookup = key
// 		field.expr = EXACT
// 		field.value = value
// 	}
// 	f.fields = append(f.fields, field)
// }

// // 过滤
// func (f *Filter) Filter() *gorm.DB {
// 	for _, field := range f.fields {
// 		switch field.expr {
// 		case "exact":
// 			f.db = f.db.Where(field.lookup+" = ?", field.value)
// 		case "iexact":
// 			f.db = f.db.Where(field.lookup+" = ?", field.value)
// 		case "contains":
// 			f.db = f.db.Where(field.lookup+" like ?", "%"+field.value+"%")
// 		case "icontains":
// 			f.db = f.db.Where(field.lookup+" ilike ?", "%"+field.value+"%")
// 		case "gt":
// 			f.db = f.db.Where(field.lookup+" > ?", field.value)
// 		case "gte":
// 			f.db = f.db.Where(field.lookup+" >= ?", field.value)
// 		case "lt":
// 			f.db = f.db.Where(field.lookup+" < ?", field.value)
// 		case "lte":
// 			f.db = f.db.Where(field.lookup+" <= ?", field.value)
// 		case "in":
// 			f.db = f.db.Where(field.lookup+" in (?)", field.value)
// 		case "startswith":
// 			f.db = f.db.Where(field.lookup+" like ?", field.value+"%")
// 		case "istartswith":
// 			f.db = f.db.Where(field.lookup+" ilike ?", field.value+"%")
// 		case "endswith":
// 			f.db = f.db.Where(field.lookup+" like ?", "%"+field.value)
// 		case "iendswith":
// 			f.db = f.db.Where(field.lookup+" ilike ?", "%"+field.value)
// 		case "range":
// 			f.db = f.db.Where(field.lookup+" between ? and ?", field.value)
// 		case "year":
// 			f.db = f.db.Where(field.lookup+" = ?", field.value)
// 		case "month":
// 			f.db = f.db.Where(field.lookup+" = ?", field.value)
// 		case "day":
// 			f.db = f.db.Where(field.lookup+" = ?", field.value)
// 		case "isnull":
// 			f.db = f.db.Where(field.lookup + " is null")
// 		}
// 	}
// 	return f.db
// }
