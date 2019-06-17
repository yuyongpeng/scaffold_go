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
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"log"
	"net"
	"net/http"
	"scaffold_go/config"
	"scaffold_go/database"
	"strconv"
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

var personIndex string = "cport_person_x"
var res *esapi.Response
var raw map[string]interface{}
var blk *bulkResponse
var numErrors int
var numIndexed int

var cfg = elasticsearch.Config{
	Addresses: config.Cfg.Elastic.Addresses,
	Transport: &http.Transport{
		MaxIdleConnsPerHost:   config.Cfg.Elastic.MaxIdleConnsPerHost,
		ResponseHeaderTimeout: time.Second,
		DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS11,
			// ...
		},
	},
}

/**
批量导入数据
 */
func ImportJobs(jobs []database.Job) {
	// Create the Elasticsearch client
	//
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var buf bytes.Buffer
	for _, job := range jobs {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_index": "%s" , "_id" : "%d", "_type" : "_doc" } }%s`, personIndex, job.Job_id, "\n"))
		data, _ := json.Marshal(job)
		data = append(data, "\n"...)
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
		fmt.Print(string(meta))
		fmt.Print(string(data))
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
	if res.IsError() { // 导入失败
		if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
			log.Fatalf("Failure to to parse response body: %s", err)
		} else {
			log.Printf("  Error: [%d] %s: %s",
				res.StatusCode,
				raw["error"].(map[string]interface{})["type"],
				raw["error"].(map[string]interface{})["reason"],
			)
		}
	} else { // 导入成功
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
func InsertElastic(job *database.Job, indexName string) (e error) {
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		e = err
		log.Fatalf("Error creating the client: %s", err)
	}

	jobByte, err := json.Marshal(job)
	if err != nil {
		e = err
		fmt.Println(err)
	}

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: strconv.Itoa(job.Job_id),
		Body:       strings.NewReader(string(jobByte)),
		Refresh:    "true",
	}
	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		e = err
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		//打印body信息
		red := bufio.NewReader(res.Body)
		var retErrStr string = ""
		for {
			line, err := red.ReadString('\n')
			if err == io.EOF{
				retErrStr = retErrStr + line
				fmt.Println(line)
				break
			}
			if err != nil {
				retErrStr = retErrStr + line
				fmt.Println(line)
			}
		}
		e = fmt.Errorf(retErrStr)
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), job.Job_id)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			e = err
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return
}

/**
查询数据 elastic search
query = map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"job_name" : "中国",
					},
				},
				"filter": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []interface{}{
							map[string]interface{}{
								"term": map[string]interface{}{
									"job_mode" : "1",
							}},
							map[string]interface{}{
								"term": map[string]interface{}{
									"job_salary" : "3",
							}},
						},
					},
				},
			},
		},
		"size": 10,		// 显示应该返回的结果数量，默认是 10
		"from": 0,		// 显示应该跳过的初始结果数量，默认是 0
		"sort": []interface{}{
			map[string]interface{}{
				"modify_time": map[string]string{
					"order": "desc",
				},
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags": []string{
				"<tag1>", "<tag2>",
			},
			"post_tags": []string{
				"<tag1>", "<tag2>",
			},
			"fields": map[string]interface{}{
				"job_name": map[string]interface{}{
					"number_of_fragments": 0,
				},
			},
		},
	}
*/
func QueryElastic(query map[string]interface{}, indexName string) (retJson map[string]interface{}, e error){
	// 连接 elasticSearch
	es, err := elasticsearch.NewClient(cfg)
	// Build the request body.
	var buf bytes.Buffer
	//var a interface{} = interface{}{nil}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		e = err
		log.Fatalf("Error encoding query: %s", err)
	}
	fmt.Println(strings.Repeat("*", 30))
	fmt.Print(buf.String())
	fmt.Println(strings.Repeat("*", 30))
	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),  //  跟踪点击总数，如果加了这条语句，会导致解析json时，不会解析hits的数据出来，看不到命中的数据
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	//fmt.Println(res.String())
	defer res.Body.Close()

	if res.IsError() {
		// 处理错误信息的body
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	// 将body的内容转换为对象操作
	//var r  map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&retJson); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(retJson["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(retJson["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	//}

	return
}
