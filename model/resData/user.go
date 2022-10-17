package resData

import "erp_center/model"

type ResUser struct {
	Account   string               `json:"account"`
	Dep       *model.SysGroupUsers `json:"dep"`
	DeptName  string               `json:"dept_name"`
	Name      string               `json:"name"`
	Title     *model.SysGroups     `json:"title"`
	TitleName string               `json:"title_name"`
	User      *model.SysUsers      `json:"user"`
}

type Records struct {
	Pg      string `json:"pg"`
	PgSize  string `json:"pgSize"`
	Recodes []any  `json:"recodes"`
}
