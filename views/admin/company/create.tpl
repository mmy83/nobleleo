<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<form method="post" action="{{urlfor "CompanyController.Create"}}">
			代码：<input type="text" name="Symbol" />
			简称：<input type="text" name="Sname" />
			<input type="submit" name="submit" value="增加" />
		</from>
	</body>
</html>