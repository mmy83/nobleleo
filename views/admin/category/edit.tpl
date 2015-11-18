<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<form method="post" action="{{urlfor "CategoryController.Edit"}}">
			<input type="hidden" name="Id" value="{{.category.Id}}">
			分类：<input type="text" name="Name" value="{{.category.Name}}"/>
			<input type="submit" name="submit" value="保存" />
		</from>
	</body>
</html>