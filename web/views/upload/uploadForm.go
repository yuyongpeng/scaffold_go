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
<form enctype="multipart/form-data" action="http://127.0.0.1:8085/upload" method="POST">
    <input type="file" name="uploadfile" />

    <input type="hidden" name="token" value="{{.}}" />

    <input type="submit" value="upload" />
</form>
</body>
</html>
`
