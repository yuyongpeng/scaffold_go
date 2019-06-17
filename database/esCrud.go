/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-14 13:31:39
LastEditors:
LastEditTime: 2019-06-14 13:31:39
Description:  处理elasticsearch的数据库操作
*/
package database

import (
	"time"
)

type Escrud struct {
}

/**
每次查询1000条数据
*/

type Job struct {
	Job_id            int    `json:"job_id"`            // `职位表的id`
	Enterprise_name   string `json:"enterprise_name"`   // `公司名称`
	Job_name          string `json:"job_name"`          // `职位名称`
	Job_description   string `json:"job_description"`   // `职位描述`
	Job_area_id       int    `json:"job_area_id"`       // `职位地区`
	Industry_id       int    `json:"industry_id"`       // `公司所属行业`
	Job_salary        int    `json:"job_salary"`        // `职位月薪`
	Job_min_education int    `json:"job_min_education"` // `最低的教育程度`
	Job_experience    int    `json:"job_experience"`    // `毕业时间`
	Job_mode          int    `json:"job_mode"`          // `全职、兼职`
	Enterprise_size   int    `json:"enterprise_size"`   // `公司规模`
	Job_status        int    `json:"job_status"`        // `职位状态`
	Modify_time       string `json:"modify_time"`       // `职位修改时间`
}

/**
获得所有可以导入的职位数量
*/
func (crud *Escrud) GetJobsCount() (count int) {
	sql := "select count(*) cu from job a left outer join enterprise b on a.enterprise_id = b.enterprise_id "
	//var logger *logrus.Logger = log.Log
	db := getDB()
	//db.LogMode(true)
	defer db.Close()
	type Result struct {
		Cu int
	}
	var result Result
	//db.Exec(sql).Scan(&result)
	db.Raw(sql).Scan(&result)
	count = result.Cu
	return
}

func (crud *Escrud) GetJobs(start, end int) (jobs []Job) {
	sql := "select a.job_id, b.enterprise_name, a.job_name, a.job_description, a.job_area_id, " +
		"b.field_id as industry_id, a.job_salary, a.job_min_education, a.job_experience, " +
		"a.job_mode, b.enterprise_size, a.job_status, a.modify_time " +
		"from job a left outer join enterprise b on a.enterprise_id = b.enterprise_id "
	//var logger *logrus.Logger = log.Log
	db := getDB()
	//db.LogMode(true)
	defer db.Close()
	rows, _ := db.Raw(sql).Limit(end).Offset(start).Rows()
	defer rows.Close()

	for rows.Next() {
		//var job Job = new(Job)
		var job_id int
		var enterprise_name string
		var job_name string
		var job_description string
		var job_area_id int
		var industry_id int
		var job_salary int
		var job_min_education int
		var job_experience int
		var job_mode int
		var enterprise_size int
		var job_status int
		var modify_time time.Time
		rows.Scan(&job_id, &enterprise_name, &job_name, &job_description,
			&job_area_id, &industry_id, &job_salary, &job_min_education,
			&job_experience, &job_mode, &enterprise_size, &job_status, &modify_time)
		//fmt.Printf(modify_time.String()+ "    ")
		//fmt.Println(modify_time.Format("2006-01-02 15:04:05"))
		var job Job = Job{job_id, enterprise_name,
			job_name, job_description,
			job_area_id, industry_id,
			job_salary, job_min_education,
			job_experience, job_mode,
			enterprise_size, job_status, modify_time.Format("2006-01-02 15:04:05"),}
		jobs = append(jobs, job)
	}
	return
}

/**
每次查询1000条数据
*/

type Person struct {
	Resume_id          int    `json:"resume_id"`          // `求职表的id`
	Job_name           string `json:"job_name"`           // `职位名称`
	Person_name        string `json:"name"`               // `人名`
	Salary             int    `json:"salary"`             // `职位月薪`
	Max_education      int    `json:"max_education"`      // `教育程度`
	Working_experience int    `json:"working_experience"` // `工作经验`
	Job_type_id        int    `json:"job_type_id"`        // `职位类型`
	Job_area_id        int    `json:"job_area_id"`        // `职位所在地区`
	Job_mode           int    `json:"job_mode"`           // `全职、兼职`
	Modify_time        string `json:"modify_time"`        // `职位更新时间`
	Birth_date         string `json:"birth_date"`         // `求职人生日`
}

/*
获得所有可以导入的求职人的数量
*/
func (crud *Escrud) GetPersonsCount() (count int) {
	sql := "select count(*) cu from resume a left outer join person b on a.person_id = b.person_id "
	//var logger *logrus.Logger = log.Log
	db := getDB()
	//db.LogMode(true)
	defer db.Close()
	type Result struct {
		Cu int
	}
	var result Result
	//db.Exec(sql).Scan(&result)
	db.Raw(sql).Scan(&result)
	count = result.Cu
	return
}

/*
查询求职人的数据
*/
func (crud *Escrud) GetPersons(start, end int) (persons []Person) {
	sql := "select a.resume_id, a.job_name, b.name, a.salary, a.max_education," +
		" a.working_experience, a.job_type_id,a.job_area_id,a.job_mode, a.modify_time,b.birth_date " +
		"from resume a left outer join person b on a.person_id = b.person_id "
	//var logger *logrus.Logger = log.Log
	db := getDB()
	//db.LogMode(true)
	defer db.Close()
	rows, _ := db.Raw(sql).Limit(end).Offset(start).Rows()
	defer rows.Close()

	for rows.Next() {
		var resume_id int
		var job_name string
		var name string
		var salary int
		var max_education int
		var working_experience int
		var job_type_id int
		var job_area_id int
		var job_mode int
		var modify_time time.Time
		var birth_date time.Time
		rows.Scan(&resume_id, &job_name, &name, &salary,
			&max_education, &working_experience, &job_type_id, &job_area_id,
			&job_mode, &modify_time, &birth_date)
		//fmt.Printf(modify_time.String()+ "    ")
		//fmt.Println(modify_time.Format("2006-01-02 15:04:05"))
		var person Person = Person{resume_id, job_name,
			name, salary, max_education,
			working_experience, job_type_id,
			job_area_id, job_mode,
			modify_time.Format("2006-01-02 15:04:05"),
			birth_date.Format("2006-01-02 15:04:05"),}
		persons = append(persons, person)
	}
	return
}
