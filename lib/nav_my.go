package lib

type AutoGenerated struct {
	Expanded bool       `json:Expanded`
	Children []Children `json:Children`
}

type Children struct {
	Text     string     `json:Text`
	Leaf     bool       `json:"leaf,omitempty"`
	Cls      string     `json:"cls,omitempty"`
	Mod      string     `json:"mod,omitempty"`
	ModURL   string     `json:"modUrl,omitempty"`
	Expanded bool       `json:"expanded,omitempty"`
	Children []Children `json:"children,omitempty"`
}

var NavMy = AutoGenerated{
	Expanded: true,
	Children: []Children{
		{
			Text:   "首页",
			Leaf:   true,
			Cls:    "node-link",
			Mod:    "desktop",
			ModURL: "desktop.Desktop",
		},

		{
			Text:     "系统管理",
			Expanded: true,
			Cls:      "node-link",
			Children: []Children{
				{
					Text:   "角色管理",
					Leaf:   true,
					Cls:    "node-link",
					Mod:    "role",
					ModURL: "role.Role",
				},
				{
					Text:   "用户管理",
					Leaf:   true,
					Cls:    "node-link",
					Mod:    "user",
					ModURL: "user.User",
				},
			},
		},

		{
			Text:     "文章管理",
			Expanded: true,
			Children: []Children{
				{
					Text: "文章列表",
					Leaf: true,
				},
				{
					Text: "发布文章",
					Leaf: true,
				},
			},
		},

		{
			Text:     "产品管理",
			Expanded: true,
			Children: []Children{
				{
					Text: "产品列表",
					Leaf: true,
				},
				{
					Text: "新增产品",
					Leaf: true,
				},
				{
					Text: "产品监控",
					Leaf: true,
				},
			},
		},

		{
			Text:     "报表管理",
			Expanded: true,
			Children: []Children{
				{
					Text: "生成报表",
					Leaf: true,
				},
				{
					Text: "报表统计",
					Leaf: true,
				},
				{
					Text: "报表打印",
					Leaf: true,
				},
			},
		},
	},
}