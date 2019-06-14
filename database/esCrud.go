/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-14 13:31:39
LastEditors:
LastEditTime: 2019-06-14 13:31:39
Description:  处理elasticsearch的数据库操作
*/
package database

type Escrud struct {
}

/**
每次查询1000条数据
*/

type Job struct {
	Job_id            int    `职位表的id`
	Enterprise_name   string `公司名称`
	Job_name          string `职位名称`
	Job_description   string `职位描述`
	Job_area_id       int    `职位地区`
	Industry_id       int    `公司所属行业`
	Job_salary        int    `职位月薪`
	Job_min_education int    `最低的教育程度`
	Job_experience    int    `毕业时间`
	Job_mode          int    `全职、兼职`
	Enterprise_size   int    `公司规模`
	Job_status        int    `职位状态`
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
		"a.job_mode, b.enterprise_size, a.job_status " +
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
		rows.Scan(&job_id, &enterprise_name, &job_name, &job_description,
			&job_area_id, &industry_id, &job_salary, &job_min_education,
			&job_experience, &job_mode, &enterprise_size, &job_status)
		var job Job = Job{job_id,enterprise_name,
			job_name, job_description,
			job_area_id, industry_id,
			job_salary, job_min_education,
			job_experience, job_mode,
			enterprise_size, job_status,}
		jobs = append(jobs, job)
	}
	return
}
