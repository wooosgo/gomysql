package main

import (
	"encoding/json"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

//-------------------------------------------------------------------------------------------------------------------------------------
//데이터 구조 및 소스 정의
//-------------------------------------------------------------------------------------------------------------------------------------
  type ArpltnInforInqireSvcVo struct {
        ReturnType string `json:"_returnType"`
        CoGrade string `json:"coGrade"`
      }

  type JResponse struct {
        List[] ArpltnInforInqireSvcVo `json:"list"`
        Parm ArpltnInforInqireSvcVo `json:"parm"`
        ArpltnInforInqireSvcVo ArpltnInforInqireSvcVo `json:"ArpltnInforInqireSvcVo"`
        TotalCount int `json:"totalCount"`
      }

  jsrc := `{"list": [{"_returnType": "json","coGrade": "1"}],"parm": {"_returnType": "json","coGrade": ""},"ArpltnInforInqireSvcVo": {"_returnType": "json","coGrade": ""},"totalCount": 98}`

	//-------------------------------------------------------------------------------------------------------------------------------------
	//데이터 출력
	//-------------------------------------------------------------------------------------------------------------------------------------
  // 결과 출력(Json)
  var jsonResponse JResponse
  err := json.Unmarshal([]byte(jsrc), &jsonResponse)
  if err != nil {
    panic(err)
  }
  fmt.Printf("%+v\n", err)

  // fmt.Println(jsonResponse)
  fmt.Printf("%+v\n", jsonResponse)

	//-------------------------------------------------------------------------------------------------------------------------------------
	//mysql connection
	//-------------------------------------------------------------------------------------------------------------------------------------
	db, err := sql.Open("mysql", "jonas:ahffk18!@tcp(127.0.0.1:3306)/test")
  if err != nil {
    panic(err)
  }
  defer db.Close()


	//-------------------------------------------------------------------------------------------------------------------------------------
	//SELECT
	//-------------------------------------------------------------------------------------------------------------------------------------

	type Tag struct {
	    ID   int    `json:"id"`
	    Name string `json:"name"`
	}

	results, err := db.Query("SELECT id, name FROM users")
    if err != nil {
	  	panic(err.Error()) // proper error handling instead of panic in your app
  }

  for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
		}
						// and then print out the tag's Name attribute
		fmt.Println(tag.ID, tag.Name)

  }

}
