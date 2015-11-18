<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<form method="post" action="{{urlfor "CategoryController.Create"}}">
		{{if .category}}
			<input type="hidden" name="Pid" value="{{.category.Id}}"/>
		{{else}}
			<input type="hidden" name="Pid" value="0"/>
		{{end}}
			分类：<input type="text" name="Name" />
			<input type="submit" name="submit" value="增加" />
		</from>
	</body>
</html>