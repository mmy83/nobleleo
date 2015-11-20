<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<form method="post" enctype="multipart/form-data" action="{{urlfor "ContentController.Create"}}">
			标题：<input type="text" name="Title" />
			公司：
			<select name="company_id">
			{{range .companys}}
				<option value="{{.Id}}">{{.Symbol}}->{{.Sname}}</option>
			{{end}}
			</select>
			分类：
			<select name="category_id">
			{{range .categorys}}
				<option value="{{.Id}}">{{.Name}}</option>
			{{end}}
			</select>
			文件：<input type="file" name="Filepath"/>
			内容：
			<textarea name="Content"></textarea>
			<input type="submit" name="submit" value="增加" />
		</from>
	</body>
</html>