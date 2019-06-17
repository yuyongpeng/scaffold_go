## 人的信息
```bash
salary : 期望月薪
max_education : 教育程度
working_experience : 工作经验
job_type_id : 职位类别 全职还是兼职
job_area_id : 职位所在地区
job_mode : 职位
modify_time : 职位更新时间
birth_date : 生日

select a.resume_id, a.job_name, b.name, a.salary, a.max_education, a.working_experience, a.job_type_id,a.job_area_id,a.job_mode, a.modify_time,b.birth_date
from resume a left outer join person b on a.person_id = b.person_id


curl -X DELETE http://127.0.0.1:9200/cport_person_x
curl -X PUT "127.0.0.1:9200/cport_person_x?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
        "properties": { 
            "resume_id": {
                "type": "integer"
            },
            "job_name": {
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart"
            },
            "name": {
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart"
            },
            "salary": {
                "type": "integer"
            },
            "max_education": { 
                "type": "integer"
            },
            "working_experience": { 
                "type": "integer"
            },
            "job_type_id": {
                "type": "integer"
            },
            "job_area_id": {
                "type": "integer"
            },
            "job_mode": {
                "type": "integer"
            },
            "modify_time": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
            },
            "birth_date": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
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

```javascript

 var enums = {
  education: {
   '1': '初中及以下',
   '2': '中技',
   '3': '中专',
   '4': '高中',
   '5': '大专',
   '6': '本科',
   '7': '硕士',
   '8': 'MBA/EMBA',
   '9': '博士',
   '10': '博士后',
   '11': '其它',
  },
  experience: {
   '1': '应届毕业',
   '2': '1年',
   '3': '2年',
   '4': '3年',
   '5': '4年',
   '6': '5年',
   '7': '6年',
   '8': '7年',
   '9': '8年',
   '10': '9年',
   '11': '10年',
   '12': '10年以上',
  },
  arrival_time: {
   '1': '已离职，随时到岗',
   '2': '已离职，一个月内到岗',
   '3': '在职，急寻新工作',
   '4': '在职，考虑机会',
   '5': '在职，暂不考虑机会',
  },
  salary: {
   '1': '1k以下',
   '2': '1k-3k',
   '3': '3k-5k',
   '4': '5k-8k',
   '5': '8k-10k',
   '6': '10k-15k',
   '7': '15k-25k',
   '8': '25k-35k',
   '9': '35k-50k',
   '10': '50k及以上',
   '11': '面议',
  },
  sex: {
   '1':'男',
   '2':'女',
  },
  job_mode: {
   '1': '全职',
   '2': '兼职',
   '3': '实习',
   '4': '全/兼职',
  },
  enterprise_size: {
   '1': '1-50人',
   '2': '50-150人',
   '3': '150-500人',
   '4': '500-1000人',
   '5': '1000-2000人',
   '6': '2000-5000人',
   '7': '5000-10000人',
   '8': '10000人以上',
  },
  age: {
   1: '10-20岁',
   2: '25-30岁',
   3: '30-35岁',
  },
}
```

```bash
select 
b.enterprise_name, a.job_name, a.job_description, a.job_area_id, b.field_id as industry_id, a.job_salary, a.job_min_education, a.job_experience, a.job_mode, b.enterprise_size, a.job_status
from job a left outer join enterprise b on a.enterprise_id = b.enterprise_id
curl -X DELETE http://127.0.0.1:9200/cport_job
curl -X PUT "127.0.0.1:9200/cport_job?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
        "properties": { 
            "job_id": {
                "type": "integer"
            },
            "enterprise_name": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "job_name": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "job_description": {
                "type": "text",
                "analyzer": "ik_max_word",
                "search_analyzer": "ik_smart"
            },
            "job_area_id": { 
                "type": "integer"
            },
            "industry_id": {
                "type": "integer"
            },
            "job_salary": {
                "type": "integer"
            },
            "job_min_education": {
                "type": "integer"
            },
            "job_experience": {
                "type": "integer"
            },
            "job_mode": {
                "type": "integer"
            },
            "enterprise_size": {
                "type": "integer"
            },
            "job_status": {
                "type": "integer"
            },
            "modify_time": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
            }
        }
    }
}
'

curl -XPOST http://localhost:9200/cport_person_x/_search?pretty  -H 'Content-Type:application/json' -d'
{
    "query" : { "match" : { "job_name" : "中国" }},
    "highlight" : {
        "pre_tags" : ["<tag1>", "<tag2>"],
        "post_tags" : ["</tag1>", "</tag2>"],
        "fields" : {
            "job_name" : {}
        }
    }
}
'

# 查看所有的索引列表
curl http://127.0.0.1:9200/_cat/indices?v 

# 去掉只读
curl -X PUT "127.0.0.1:9200/cport_person/_settings?pretty" -H 'Content-Type: application/json' -d'
{"index.blocks.read_only_allow_delete": null}
'

# 使用bulk插入数据
curl -X PUT "127.0.0.1:9200/_bulk?pretty" -H 'Content-Type: application/json' -d'
{ "index" : { "_index":"cport_person","_type":"_doc", "_id" : "1" } }
{"company":"公司名称","description_of_job":"职位描述","work_place_id":10,"work_place":"北京","industry_id":11,"industry":"IT","monthly_salary":8000,"education":3,"experience":2,"type":1}
'



```