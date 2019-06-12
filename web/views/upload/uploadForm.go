/**
      ___           ___           ___
     /\__\         /\__\         /\  \
    /:/ _/_       /:/  /        /::\  \
   /:/ /\  \     /:/  /        /:/\:\  \
  /:/ /::\  \   /:/  /  ___   /:/ /::\  \
 /:/_/:/\:\__\ /:/__/  /\__\ /:/_/:/\:\__\
 \:\/:/ /:/  / \:\  \ /:/  / \:\/:/  \/__/
  \::/ /:/  /   \:\  /:/  /   \::/__/
   \/_/:/  /     \:\/:/  /     \:\  \
     /:/  /       \::/  /       \:\__\
     \/__/         \/__/         \/__/
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-12 07:57:29
LastEditors:
LastEditTime: 2019-06-12 07:57:29
Description:
*/
package upload

var Body = `
<html>
<head>
    <title>Upload file</title>
</head>
<body>
<h1>请将退役士兵的数据导入数据库中：</h1>
<form enctype="multipart/form-data" action="http://127.0.0.1:8085/upload" method="POST">
    <input type="file" name="uploadfile" />

    <input type="hidden" name="token" value="{{.}}" />

    <input type="submit" value="upload" />
</form>
<h2>规则说明</h2>
只支持CSV格式的文件</br>
CSV 文件的格式："姓名","身份证号","出生年月","兵种","专业","退役号"</br>
正确示例："光头强","362101198721010087","1982-01-08","步兵","列兵","110982091"</br>
错误示例："光头强,"362101198721010087","1982-01-08","步兵","列兵","110982091"</br>
错误示例："光头强,362101198721010087","1982-01-08","步兵","列兵","110982091"</br>
</body>
</html>
`
