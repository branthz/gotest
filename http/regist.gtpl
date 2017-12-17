<!--xxx>
<html>
<head>
<title></title>
</head>
<body>
<form action="http://127.0.0.1:9998/regist" method="post">
    厂家：<input type="text" name="vendor">
    邮箱:<input type="text" name="email">
    ip地址:<input type="text" name="ip">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="注册">
</form>
</body>
</html>
