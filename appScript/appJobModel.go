package appScript

import (
	"encoding/json"
	"fmt"
)

type JobConfigData struct {
	Job struct {
		Setting struct {
			Speed struct {
				Channel int `json:"channel"`
			} `json:"speed"`
			ErrorLimit struct {
				Record     int     `json:"record"`
				Percentage float64 `json:"percentage"`
			} `json:"errorLimit"`
		} `json:"setting"`
		Content []struct {
			Reader map[string]interface{} `json:"reader"`
			Writer map[string]interface{} `json:"writer"`
		} `json:"content"`
	} `json:"job"`
}

type OracleReader struct {
	Name      string `json:"name"`
	Parameter struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Connection []struct {
			QuerySql []string `json:"querySql"`
			JdbcUrl  []string `json:"jdbcUrl"`
		} `json:"connection"`
	} `json:"parameter"`
}

type OracleWriter struct {
	Name      string `json:"name"`
	Parameter struct {
		Username   string   `json:"username"`
		Password   string   `json:"password"`
		Column     []string `json:"column"`
		PreSql     []string `json:"preSql"`
		Connection []struct {
			JdbcUrl string   `json:"jdbcUrl"`
			Table   []string `json:"table"`
		} `json:"connection"`
	} `json:"parameter"`
}
type MysqlReader struct {
	Name      string `json:"name"`
	Parameter struct {
		Username   string       `json:"username"`
		Password   string       `json:"password"`
		Column     []string     `json:"column"`
		SplitPk    string       `json:"splitPk"`
		Where      string       `json:"where"`
		Connection []Connection `json:"connection"`
	} `json:"parameter"`
}

type Connection struct {
	QuerySql []string `json:"querySql"`
	Table    []string `json:"table"`
	JdbcUrl  []string `json:"jdbcUrl"`
}

type MysqlWriter struct {
	Name      string `json:"name"`
	Parameter struct {
		WriteMode  string        `json:"writeMode"`
		Session    []interface{} `json:"session"`
		PreSql     []string      `json:"preSql"`
		Username   string        `json:"username"`
		Password   string        `json:"password"`
		Column     []string      `json:"column"`
		Connection []struct {
			Table   []string `json:"table"`
			JdbcUrl string   `json:"jdbcUrl"`
		} `json:"connection"`
	} `json:"parameter"`
}

type TxtReader struct {
	Name      string `json:"name"`
	Parameter struct {
		Path           []string    `json:"path"`
		Encoding       string      `json:"encoding"`
		Column         []TxtColumn `json:"column"`
		FieldDelimiter string      `json:"fieldDelimiter"`
		SkipHeader     string      `json:"skipHeader"`
	} `json:"parameter"`
}

type TxtWriter struct {
	Name      string `json:"name"`
	Parameter struct {
		Path           string `json:"path"`
		FileName       string `json:"fileName"`
		WriteMode      string `json:"writeMode"`
		FileFormat     string `json:"fileFormat"`
		Encoding       string `json:"encoding"`
		FieldDelimiter string `json:"fieldDelimiter"`
	} `json:"parameter"`
}
type TxtColumn struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
}

// ToStruct 将 json 数组转为 Struct
func ToStruct(conn any) (data map[string]interface{}) {
	// 将 TxtReader 转换为 JSON 字节数组
	jsonBytes, err := json.Marshal(conn)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	// 将 JSON 字节数组解析为 struct {}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		fmt.Println("Error converting to struct {}:", err)
		return
	}
	return
}
