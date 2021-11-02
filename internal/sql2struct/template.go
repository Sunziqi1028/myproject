package sql2struct

import (
	"fmt"
	"html/template"
	"os"
	"projection/internal/word"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelcase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{
		structTpl: structTpl,
	}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn  {
	tpColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tpColumns = append(tpColumns, &StructColumn{
			Name: column.ColumnName,
			Type: column.ColumnType,
			Tag: fmt.Sprintf("`json:"+"%s"+"`",column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tpColumns
}

func (t *StructTemplate) Generate(tableName string, tpColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{"ToTemplate": word.UnderscoreToUpperCamelCase,}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns: tpColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}