/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-14 13:51:09
LastEditors:
LastEditTime: 2019-06-14 13:51:09
Description:  对elasticsearch的插入查询等操作
*/

package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"scaffold_go/database"
	"strings"
	"time"
)
type bulkResponse struct {
	Errors bool `json:"errors"`
	Items  []struct {
		Index struct {
			ID     string `json:"_id"`
			Result string `json:"result"`
			Status int    `json:"status"`
			Error  struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
				Cause  struct {
					Type   string `json:"type"`
					Reason string `json:"reason"`
				} `json:"caused_by"`
			} `json:"error"`
		} `json:"index"`
	} `json:"items"`
}
var personIndex string = "cport_person"
var res *esapi.Response
var raw map[string]interface{}
var blk *bulkResponse
var numItems   int
var numErrors  int
var numIndexed int
var numBatches int
var currBatch  int

func ImportJobs(jobs []database.Job){
	// Create the Elasticsearch client
	//
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var buf bytes.Buffer
	for _, job := range jobs {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d", "_type" : "_doc" } }%s`, job.Job_id, "\n"))
		data, _ := json.Marshal(job)
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}
	// 记录导入的开始时间
	start := time.Now().UTC()
	// 进行批量数据导入
	res, err = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(personIndex))
	if err != nil {
		log.Fatalf("Failure indexing : %s", err)
	}
	// If the whole request failed, print error and mark all documents as failed
	//
	if res.IsError() {  // 导入失败
		if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
			log.Fatalf("Failure to to parse response body: %s", err)
		} else {
			log.Printf("  Error: [%d] %s: %s",
				res.StatusCode,
				raw["error"].(map[string]interface{})["type"],
				raw["error"].(map[string]interface{})["reason"],
			)
		}
	} else {	// 导入成功
		if err := json.NewDecoder(res.Body).Decode(&blk); err != nil {
			log.Fatalf("Failure to to parse response body: %s", err)
		} else {
			// 循环返回值 查看每条数据的导入状态
			for _, d := range blk.Items {
				// ... so for any HTTP status above 201 ...
				//
				if d.Index.Status > 201 {
					// ... increment the error counter ...
					// 记录失败的数目
					numErrors++

					// ... and print the response status and error information ...
					log.Printf("  Error: [%d]: %s: %s: %s: %s",
						d.Index.Status,
						d.Index.Error.Type,
						d.Index.Error.Reason,
						d.Index.Error.Cause.Type,
						d.Index.Error.Cause.Reason,
					)
				} else {
					// ... otherwise increase the success counter.
					//
					numIndexed++
				}
			}
		}
	}

	buf.Reset()

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	// 打印导入报告
	log.Println(strings.Repeat("=", 80))

	dur := time.Since(start)

	if numErrors > 0 {
		log.Fatalf(
			"Indexed [%d] documents with [%d] errors in %s (%.0f docs/sec)",
			numIndexed,
			numErrors,
			dur.Truncate(time.Millisecond),
			1000.0/float64(dur/time.Millisecond)*float64(numIndexed),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%d] documents in %s (%.0f docs/sec)",
			numIndexed,
			dur.Truncate(time.Millisecond),
			1000.0/float64(dur/time.Millisecond)*float64(numIndexed),
		)
	}
}

/**
插入一条数据到elasticsearch
 */
func InsertJob(job database.Job){

}

/**
查询数据
 */
func QueryJob(query interface{}){

}