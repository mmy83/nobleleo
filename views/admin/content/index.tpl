
<ul id="data">
{{range .contents}}
	<li>
		{{.Id}}:{{.Title}}->{{.User.Username}}->{{.Company.Sname}}
		&nbsp;&nbsp;
		<a href="{{urlfor "ContentController.Edit" "id" .Id}}">编辑</a>
		&nbsp;&nbsp;
		<a href="{{urlfor "ContentController.Del" ":id" .Id}}">删除</a>
	</li>
{{end}}
</ul>
