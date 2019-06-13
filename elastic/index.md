## 人的信息
```bash
curl -X DELETE http://127.0.0.1:9200/cport_person
curl -X PUT "127.0.0.1:9200/cport_person?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
        "properties": { 
            "content": {
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart"
            },
            "dt": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
            },
            "keyword_arr": {
                "type": "object"
            },
            "stas": {
                "type": "integer"
            },
            "is_published": {
                "type": "boolean"
            },
            "create_at": {
                "type": "date_range",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
            },
            "count": {
                "type": "integer_range"
            },
            "json_obj": {
                "type": "object"
            }
        }
    }
}
'
```

## 职位信息

* company : 公司名称
* description_of_job : 职位描述
* work_place_id : 工作地点表的id
* work_place : 工作地点
* industry_id : 行业表的id
* industry : 行业名称
* monthly_salary : 月薪
* education : 职位匹配的教育程度
* experience : 职位需要的工作经验 年
* type : 全职还是兼职

```bash
curl -X DELETE http://127.0.0.1:9200/cport_person
curl -X PUT "127.0.0.1:9200/cport_person?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
        "properties": { 
            "company": {
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart"
            },
            "description_of_job": {
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart"
            },
            "work_place_id": { 
                "type": "integer"
            },
            "work_place": {
                "type": "keyword"
            },
            "industry_id": {
                "type": "integer"
            },
            "industry": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "monthly_salary": {
                "type": "integer_range"
            },
            "education": {
                "type": "integer"
            },
            "experience": {
                "type": "integer"
            },
            "type": {
                "type": "integer"
            }
        }
    }
}
'
```