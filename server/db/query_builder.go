package db

import (
	"database/sql"
	"fmt"

	"strings"
)

type Config struct {
	Db *sql.DB
}

func SelectBuilder(tableName string, where map[string]interface{}, orderBy map[string]bool, groupBy []string) string {
	mainQuery := fmt.Sprintf("SELECT * FROM %s ", tableName)
	if len(where) > 0 {
		sqlQuery := Where
		var conditions []string
		var limitOrCount []string
		for key, value := range where {
			if key != Count && key != Skip {
				switch value.(type) {
				case string:
					conditions = append(conditions, fmt.Sprintf(" %s='%v' ", key, value))
				case []string:
					temp := fmt.Sprintf(" '%v'=ANY(ARRAY[", key)
					var vals []string
					for _, t := range value.([]string) {
						vals = append(vals, fmt.Sprintf("'%s'", t))
					}
					temp += strings.Join(vals, ",") + "]) "
					conditions = append(conditions, temp)
				default:
					conditions = append(conditions, fmt.Sprintf(" %s=%v ", key, value))
				}
			} else {
				if key == Count {
					switch value.(type) {
					case int, int32, int64:
						limitOrCount = append(limitOrCount, fmt.Sprintf(" LIMIT %v ", value))
					}
				}
				if key == Skip {
					switch value.(type) {
					case int, int32, int64:
						limitOrCount = append(limitOrCount, fmt.Sprintf(" OFFSET %v ", value))
					}
				}
			}
		}
		mainQuery += sqlQuery + strings.Join(conditions, AND) + strings.Join(limitOrCount, " ")
	}
	if len(orderBy) > 0 {
		sqlQuery := OrderBy
		var columnsWithOrder []string
		for key, order := range orderBy {
			if order {
				columnsWithOrder = append(columnsWithOrder, fmt.Sprintf(" %s ASC ", key))
			} else {
				columnsWithOrder = append(columnsWithOrder, fmt.Sprintf(" %s DESC ", key))
			}
		}
		mainQuery += sqlQuery + strings.Join(columnsWithOrder, ",")
	}
	if len(groupBy) > 0 {
		sqlQuery := GroupBy
		mainQuery += sqlQuery + strings.Join(groupBy, " , ")
	}
	mainQuery += ";"
	fmt.Println("Query: ", mainQuery)
	return mainQuery
}

func FullTextSearchBuilder(tableName string, query string, sortRankBy bool, where map[string]interface{}, orderBy map[string]bool, groupBy []string) string {
	mainQuery := fmt.Sprintf("SELECT * FROM %s, plainto_tsquery('%s') AS q ", tableName, query)
	mainQuery += Where + " (tsv @@ q) "

	if len(where) > 0 {
		var sqlQuery string
		var conditions []string
		var limitOrCount []string
		for key, value := range where {
			if key != Count && key != Skip {
				switch value.(type) {
				case string:
					conditions = append(conditions, fmt.Sprintf(" %s='%v' ", key, value))
				case []string:
					temp := fmt.Sprintf(" '%v'=ANY(ARRAY[", key)
					var vals []string
					for _, t := range value.([]string) {
						vals = append(vals, fmt.Sprintf("'%s'", t))
					}
					temp += strings.Join(vals, ",") + "]) "
					conditions = append(conditions, temp)
				default:
					conditions = append(conditions, fmt.Sprintf(" %s=%v ", key, value))
				}
			} else {
				if key == Count {
					switch value.(type) {
					case int, int32, int64:
						limitOrCount = append(limitOrCount, fmt.Sprintf(" LIMIT %v ", value))
					}
				}
				if key == Skip {
					switch value.(type) {
					case int, int32, int64:
						limitOrCount = append(limitOrCount, fmt.Sprintf(" OFFSET %v ", value))
					}
				}
			}
		}
		mainQuery += sqlQuery + AND + strings.Join(conditions, AND) + strings.Join(limitOrCount, " ")
	}
	mainQuery += OrderBy + fmt.Sprintf(" ts_rank_cd(tsv, plainto_tsquery('%s')) ", query)
	if sortRankBy {
		mainQuery += " ASC "
	} else {
		mainQuery += " DESC "
	}
	if len(orderBy) > 0 {
		var sqlQuery string
		var columnsWithOrder []string
		for key, order := range orderBy {
			if order {
				columnsWithOrder = append(columnsWithOrder, fmt.Sprintf(" %s ASC ", key))
			} else {
				columnsWithOrder = append(columnsWithOrder, fmt.Sprintf(" %s DESC ", key))
			}
		}
		mainQuery += sqlQuery + " , " + strings.Join(columnsWithOrder, ",")
	}
	if len(groupBy) > 0 {
		sqlQuery := GroupBy
		mainQuery += sqlQuery + strings.Join(groupBy, " , ")
	}
	mainQuery += ";"
	fmt.Println("Query: ", mainQuery)
	return mainQuery
}

func UpdateBuilder(tableName string, set map[string]interface{}, condition []string) string {
	var updates []string
	query := fmt.Sprintf("UPDATE %s SET ", tableName)
	for key, value := range set {
		switch value.(type) {
		case string:
			updates = append(updates, fmt.Sprintf(" %s='%v' ", key, value))
		default:
			updates = append(updates, fmt.Sprintf(" %s=%v ", key, value))
		}
	}
	updates = append(updates, fmt.Sprintf(" %s=%v ", UpdatedAt, CurrentTimestamp))
	query += strings.Join(updates, " , ")
	if len(condition) > 0 {
		query += Where + strings.Join(condition, AND)
	}
	query += ";"
	return query
}

func DeleteBuilder(tableName string, condition []string) string {
	query := fmt.Sprintf("DELETE FROM %s ", tableName)

	if len(condition) > 0 {
		query += Where + strings.Join(condition, AND)
	}
	query += ";"
	return query
}

const (
	// SQL

	Where            = " WHERE "
	Count            = "count"
	Skip             = "skip"
	AND              = " AND "
	OrderBy          = " ORDER BY "
	GroupBy          = " GROUP BY "
	CurrentTimestamp = "CURRENT_TIMESTAMP"
	UpdatedAt        = "updated_at"
)
