<a href="{{urlfor "CategoryController.Create" "pid" "0"}}">添加分类</a>

<ul id="data">
{{range .categorys}}
	<li>
		{{.Id}}:{{.Name}}->{{.Pid}}
		&nbsp;&nbsp;
		<a href="{{urlfor "CategoryController.Edit" "id" .Id}}">编辑</a>
		&nbsp;&nbsp;
		<a href="{{urlfor "CategoryController.Del" "id" .Id}}">删除</a>
		&nbsp;&nbsp;
		<a href="{{urlfor "CategoryController.Create" "pid" .Id}}">添加子分类</a>
	</li>
{{end}}
</ul>
{{if gt .paginator.PageNums 1}}
<div class="pagination pagination-sm">
    {{if .paginator.HasPrev}}
        <span><a href="{{.paginator.PageLinkFirst}}">&lt;&lt;</a></span>
        <span><a href="{{.paginator.PageLinkPrev}}">&lt;</a></span>
    {{else}}
        <span class="disabled"><a>&lt;&lt;</a></span>
        <span class="disabled"><a>&lt;</a></span>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <span{{if $.paginator.IsActive .}} class="active"{{end}}>
           <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </span>
    {{end}}
    {{if .paginator.HasNext}}
        <span><a href="{{.paginator.PageLinkNext}}">&gt;</a></span>
        <span><a href="{{.paginator.PageLinkLast}}">&gt;&gt;</a></span>
    {{else}}
        <span class="disabled"><a>&gt;</a></span>
        <span class="disabled"><a>&gt;&gt;</a></span>
    {{end}}
</div>
{{end}}