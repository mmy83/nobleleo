<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<form method="post" action="/admin/company/edit">
			<input type="hidden" name="Id" value="{{.company.Id}}"/>
			代码：<input type="text" name="Symbol" value="{{.company.Symbol}}"/>
			简称：<input type="text" name="Sname" value="{{.company.Sname}}"/>
			<input type="submit" name="submit" value="编辑" />
		</from>
	</body>
</html>