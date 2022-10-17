package lib

import (
	"fmt"
	"strconv"
)

type Node struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	Leaf        bool   `json:"leaf"`
	Cls         string `json:"cls"`
	Class       string `json:"class"`
	AllUserShow bool   `json:"allUserShow"`
	Children    []Node `json:"children"`
	IconCls     string `json:"iconCls"`
	Qtitle      string `json:"qtitle"`
}

var Nav = Node{
	Text: "ERP功能导航",
	Leaf: false,
	Cls:  "folder",
	Children: []Node{
		{
			Text:        "首页",
			Cls:         "i-home",
			Leaf:        true,
			Class:       "HomePage",
			AllUserShow: true,
		},
		{
			Text: "我的",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:        "我的考勤",
					Cls:         "i-shield",
					Leaf:        true,
					Class:       "attendance.my_page",
					AllUserShow: true,
				},
				{
					Text:        "密码设置",
					Cls:         "i-key",
					Leaf:        true,
					Class:       "sys.myPage",
					AllUserShow: true,
				},
				{
					Text:        "场地预约",
					Cls:         "i-hand",
					Leaf:        true,
					Class:       "hr.venue_booking",
					AllUserShow: true,
				},
				{
					Text:        "业务申请",
					Cls:         "i-hand",
					Leaf:        true,
					Class:       "project.business_apply",
					AllUserShow: true,
				},
				{
					Text:        "登录日志",
					Leaf:        true,
					Cls:         "i-log",
					Class:       "tools.user_logs",
					AllUserShow: true,
				},
				{
					Text:  "任务管理",
					Cls:   "i-event",
					Leaf:  true,
					Class: "project.task_manager",
				},
			},
		},
		{
			Text: "HR管理",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "员工管理",
					Cls:   "i-user-manager",
					Leaf:  true,
					Class: "sys.user_manager",
				},
				{
					Text:  "权限报表",
					Cls:   "i-power-report",
					Leaf:  true,
					Class: "sys.power_report",
				},
				{
					Text:  "考勤审批",
					Cls:   "i-event",
					Leaf:  true,
					Class: "attendance.aduit",
				},
				{
					Text:  "假期设置",
					Cls:   "i-set",
					Leaf:  true,
					Class: "attendance.setting",
				},
				{
					Text:  "外网登录设置",
					Cls:   "i-set",
					Leaf:  true,
					Class: "sys.outnet_login",
				},
				{
					Text:  "冻结账号",
					Cls:   "i-set",
					Leaf:  true,
					Class: "sys.freeze_account",
				},
			},
		},
		{
			Text: "财务管理",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "财务预算",
					Cls:   "i-door",
					Leaf:  true,
					Class: "financial.yu_suan_manage",
				}, {
					Text:  "财务预算汇总",
					Cls:   "i-monitor",
					Leaf:  true,
					Class: "financial.yu_suan_statistics",
				}, {
					Text:  "借款流程",
					Cls:   "i-in",
					Leaf:  true,
					Class: "financial.jie_kuan",
				}, {
					Text:  "核销流程",
					Cls:   "i-out",
					Leaf:  true,
					Class: "financial.he_xiao",
				}, {
					Text:  "添加业务账",
					Cls:   "i-money-add",
					Leaf:  true,
					Class: "sell.binding",
				}, {
					Text:  "月消费明细",
					Leaf:  true,
					Class: "financial.card_detail",
				}, {
					Text:  "野狼交易明细",
					Leaf:  true,
					Class: "financial.yl_detail",
				}, {
					Text:  "线上交易明细",
					Leaf:  true,
					Class: "financial.online_bonus",
				}, {
					Text:  "公募收入统计",
					Leaf:  true,
					Class: "financial.fund_fee",
				}, {
					Text: "基金投资",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "基金买卖录入",
							Cls:   "i-form",
							Leaf:  true,
							Class: "financial.fund.trade_manager",
						}, {
							Text:  "仓位探针",
							Cls:   "i-table",
							Leaf:  true,
							Class: "financial.fund.hold",
						}, {
							Text:  "报表分析",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "financial.fund.report",
						},
					},
				},
			},
		},
		{
			Text: "合规管理",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "基金销售报审",
					Cls:   "i-form",
					Leaf:  true,
					Class: "approve.fund_material",
				}, {
					Text:  "基金销售审查",
					Cls:   "i-form",
					Leaf:  true,
					Class: "approve.fund_material_approve",
				}, {
					Text:  "投资咨询报审",
					Cls:   "i-form",
					Leaf:  true,
					Class: "approve.invest_material",
				}, {
					Text:  "投资咨询审查",
					Cls:   "i-form",
					Leaf:  true,
					Class: "approve.invest_material_approve",
				}, {
					Text:  "第五军团审查",
					Cls:   "i-form",
					Leaf:  true,
					Class: "approve.five_legion_approve",
				}, {
					Text:  "电子保存数据资料统计表",
					Cls:   "i-form",
					Leaf:  true,
					Class: "approve.data_file_statistics",
				},
			},
		},
		{
			Text: "问卷管理",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "问卷管理",
					Cls:   "i-info-edit",
					Leaf:  true,
					Class: "fund.sms.risk_eva",
				}, {
					Text:  "风险评估结果查询",
					Cls:   "i-power-report",
					Leaf:  true,
					Class: "fund.sms.sms_risks",
				}, {
					Text:  "投顾风险评估历史查询",
					Cls:   "i-power-report",
					Leaf:  true,
					Class: "fund.sms.adviser_risks",
				}, {
					Text:  "回访问卷查询",
					Cls:   "i-power-report",
					Leaf:  true,
					Class: "fund.sms.sms_questionnaire",
				}, {
					Text:  "公募回访样本抽取",
					Cls:   "i-power-report",
					Leaf:  true,
					Class: "fund.push.gm_questionnaire_push",
				},
			},
		},
		{
			Text: "营销中心",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text: "交易所审核",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "日报表",
							Leaf:  true,
							Class: "exchange.day_report",
						}, {
							Text:  "月报表",
							Leaf:  true,
							Class: "exchange.month_report",
						}, {
							Text:  "月报明细表",
							Leaf:  true,
							Class: "exchange.month_details",
						}, {
							Text:  "实时天狼卡",
							Cls:   "i-eye",
							Leaf:  true,
							Class: "exchange.realTime",
						}, {
							Text:  "天狼卡查询",
							Cls:   "i-zoom",
							Leaf:  true,
							Class: "exchange.card_search",
						}, {
							Text:  "证监会月报",
							Leaf:  true,
							Class: "exchange.month_zjh_report",
						},
					},
				},
				{
					Text: "演示账号(藏)",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "首次下单开通",
							Cls:   "i-cart-add",
							Leaf:  true,
							Class: "sell.special_booking",
						}, {
							Text:  "直接加减有效期",
							Cls:   "i-control-stop-blue",
							Leaf:  true,
							Class: "sell.special_edit",
						},
					},
				},
				{
					Text: "登陆日志",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "明细",
							Cls:   "i-log",
							Leaf:  true,
							Class: "crm.soft_login",
						},
					},
				},
				{
					Text: "野狼军团",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "小金币管理",
							Leaf:  true,
							Class: "wild.gold_manage",
						}, {
							Text:  "客户信息汇总",
							Leaf:  true,
							Cls:   "i-user-manager",
							Class: "wild.cm.wild_wolf_customers",
						}, {
							Text:  "开户申请",
							Cls:   "i-trader",
							Leaf:  true,
							Class: "wild.cm.new_customer_wf",
						}, {
							Text:  "开户信息管理",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "wild.cm.future_company_customers",
						},
						{
							Text:  "野狼白名单",
							Cls:   "i-trader",
							Leaf:  true,
							Class: "wild.cm.ww_cps_orders",
						},
						{
							Text:  "野狼收益分析",
							Cls:   "i-table",
							Leaf:  true,
							Class: "wild.cm.income_analyse",
						},
						{
							Text:  "好Trader成绩",
							Cls:   "i-table",
							Leaf:  true,
							Class: "wild.cm.t_trade_detail",
						},
					},
				},
				{
					Text: "呼叫中心",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "通话记录",
							Leaf:  true,
							Class: "call.sessionRecords",
						}, {
							Text:  "分机管理",
							Leaf:  true,
							Class: "call.phoneSetting",
						},
					},
				},
				{
					Text: "短信平台",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "后台设置",
							Cls:   "i-group",
							Leaf:  true,
							Class: "sms.setting",
						}, {
							Text:  "发短信",
							Cls:   "i-sms",
							Leaf:  true,
							Class: "sms.send",
						},
					},
				},
				{
					Text: "设置",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "线上价格",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "sell.e_shop_setting",
						}, {
							Text:  "软件弹窗",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "sell.popup_setting",
						}, {
							Text:  "优惠券",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.coupons.coupons",
						}, {
							Text:  "软件白名单",
							Leaf:  true,
							Class: "sell.auth_whitelist",
						}, {
							Text:  "短信模板",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "crm.sms_setting",
						}, {
							Text:  "微学伴优惠券",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "sell.wxb",
						}, {
							Text:  "免调查问卷白名单",
							Leaf:  true,
							Class: "sell.risk_whitelist",
						}, {
							Text:  "选股类型配置",
							Leaf:  true,
							Class: "sell.pool_type_list",
						}, {
							Text:  "每日点评",
							Leaf:  true,
							Class: "stock.sell.reviews",
						}, {
							Text:  "产品购买的天数",
							Leaf:  true,
							Class: "sell.product_buy_max_days",
						},
					},
				},
				{
					Text: "售后服务",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "产品退货",
							Cls:   "i-cart-delete",
							Leaf:  true,
							Class: "sell.return",
						}, {
							Text:  "冻结及解冻",
							Cls:   "i-control-stop-blue",
							Leaf:  true,
							Class: "sell.dis_en",
						}, {
							Text:  "客户信息维护",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "crm.custom_edit",
						}, {
							Text:  "野狼平台工具",
							Cls:   "i-sys-admin-tools",
							Leaf:  true,
							Class: "wild.ylChangeTool",
						}, {
							Text:  "交易平台订单",
							Leaf:  true,
							Class: "wild.tp.order",
						}, {
							Text:  "代付交易记录",
							Leaf:  true,
							Class: "wild.tp.payment_record",
						}, {
							Text:  "野狼卡号信息",
							Leaf:  true,
							Class: "wild.tp.num",
						},
					},
				},
				{
					Text: "投顾",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text: "BQ策略",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "预警池",
									Cls:   "i-log",
									Leaf:  true,
									Class: "sell.ia.bq.warn",
								},
							},
						},
					},
				},
				{
					Text: "个人中心业务",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "天狼帐号解绑",
							Cls:   "i-user-manager",
							Leaf:  true,
							Class: "fund.user.unbind_tl_account",
						}, {
							Text:  "手机号失效",
							Cls:   "i-trader",
							Leaf:  true,
							Class: "fund.user.phone_failure",
						}, {
							Text:  "交换手机号",
							Cls:   "i-user-manager",
							Leaf:  true,
							Class: "fund.user.phone_swap",
						}, {
							Text:  "账户注销",
							Cls:   "i-sys-admin-tools",
							Leaf:  true,
							Class: "fund.user.user_cancel",
						},
					},
				},
				{
					Text:  "客户资源(直销)",
					Cls:   "i-user_comment",
					Leaf:  true,
					Class: "crm.res_manage",
				},
				{
					Text:  "下订单",
					Cls:   "i-cart-add",
					Leaf:  true,
					Class: "sell.booking",
				},
				{
					Text:  "待审核订单",
					Cls:   "i-cart-go",
					Leaf:  true,
					Class: "sell.pending",
				},
				{
					Text:  "待成交订单",
					Cls:   "i-cart-go",
					Leaf:  true,
					Class: "sell.unover",
				},
				{
					Text:  "待开通订单",
					Cls:   "i-cart-edit",
					Leaf:  true,
					Class: "sell.opening",
				},
				{
					Text:  "订单列表",
					Cls:   "i-cart",
					Leaf:  true,
					Class: "sell.history",
				},
				{
					Text:  "微学伴订单列表",
					Leaf:  true,
					Class: "sell.yl_college",
				},
				{
					Text:  "野狼客户签约",
					Cls:   "i-group",
					Leaf:  true,
					Class: "sell.ye_lang_send",
				},
				{
					Text:  "投顾客户签约",
					Cls:   "i-group",
					Leaf:  true,
					Class: "sell.tou_gu_send",
				},
				{
					Text:  "业绩分析",
					Cls:   "i-log",
					Leaf:  true,
					Class: "sell.bonus_detail",
				},
				{
					Text:  "发票记录",
					Cls:   "i-log",
					Leaf:  true,
					Class: "sell.receipt",
				},
				{
					Text:  "课程权限管理",
					Cls:   "i-group",
					Leaf:  true,
					Class: "sell.Class_manager",
				},
			},
		},
		{
			Text: "私募营销",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text: "信息维护",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "私募产品信息",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.fund_info",
						}, {
							Text:  "私募发布管理",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.sms_publish_info",
						}, {
							Text:  "私募标签管理",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.sms_fund_index",
						}, {
							Text:  "私募策略管理",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.sms_fund_policy",
						}, {
							Text:  "私募基金经理管理",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.sms_fund_manager",
						}, {
							Text:  "私募基金公司管理",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.sms_fund_company",
						}, {
							Text:  "私募产品信息ext",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.sms_fund_info_ext",
						}, {
							Text:  "报名预约管理",
							Cls:   "i-settings",
							Leaf:  true,
							Class: "fund.sms.app_info",
						}, {
							Text:  "新-报名预约管理",
							Cls:   "i-settings",
							Leaf:  true,
							Class: "fund.sms.sms_yy_info",
						}, {
							Text:  "私募分红",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.sms.fund_dividend",
						}, {
							Text:  "私募基金净值",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.sms.fund_net",
						}, {
							Text:  "私募基金区间收益",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.sms.fund_profit",
						}, {
							Text:  "购买开放日设置",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "fund.sms.fund_open",
						}, {
							Text:  "赎回开放日设置",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "fund.sms.redeem_open",
						}, {
							Text:  "业绩基准比较",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.sms.sms_fund_standard",
						}, {
							Text:  "信批邮件发送",
							Cls:   "i-email-go",
							Leaf:  true,
							Class: "fund.sms.sms_report_email",
						}, {
							Text:  "信批消息发送",
							Cls:   "i-email-go",
							Leaf:  true,
							Class: "fund.sms.sms_report_msg",
						}, {
							Text:  "私募回访",
							Cls:   "i-history",
							Leaf:  true,
							Class: "fund.sms.sms_visit",
						},
					},
				},
			},
		},
		{
			Text: "私募资管",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text: "期权投资管理",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "账号管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.qq.account_manager",
						}, {
							Text:  "历史资金查询",
							Cls:   "i-table",
							Leaf:  true,
							Class: "fund.qq.income_re",
						},
					},
				},
				{
					Text: "期货投资管理",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "账号管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.sm.account_manager",
						}, {
							Text:  "人员管理",
							Cls:   "i-trader",
							Leaf:  true,
							Class: "fund.sm.trader_manager",
						}, {
							Text:  "资金管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.sm.fund_manager",
						}, {
							Text:  "资金调整日志",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.sm.fund_log",
						}, {
							Text:  "成交记录/分析",
							Cls:   "i-table",
							Leaf:  true,
							Class: "fund.sm.trading_re",
						}, {
							Text:  "收益记录",
							Cls:   "i-table",
							Leaf:  true,
							Class: "fund.sm.income_re",
						}, {
							Text:  "历史资金查询",
							Cls:   "i-table",
							Leaf:  true,
							Class: "fund.sm.fund_his",
						}, {
							Text:  "报表分析",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.sm.analyse",
						},
					},
				},
				{
					Text: "股票投资管理",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "账号管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.ss.account_manager",
						}, {
							Text:  "人员管理",
							Cls:   "i-trader",
							Leaf:  true,
							Class: "fund.ss.trader_manager",
						}, {
							Text:  "资金管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.ss.fund_manager",
						}, {
							Text:  "仓位探针",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.shimmer_position",
						}, {
							Text:  "股票池管理",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.stock_pool",
						}, {
							Text:  "资金调整日志",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.fund_log",
						}, {
							Text:  "我的持仓",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.my_position",
						}, {
							Text:  "我的持仓(手机版)",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.my_position_mobile",
						}, {
							Text:  "我的持仓(量化)",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.my_position_lh",
						}, {
							Text:  "盈利汇总(量化)",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.my_collect_lh",
						}, {
							Text:  "指令列表",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.order_list",
						}, {
							Text:  "对账",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.stock_dz",
						}, {
							Text:  "风控参数设置",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.fund_risk",
						}, {
							Text:  "新股股票池",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.stock_new",
						}, {
							Text:  "收益记录",
							Cls:   "i-table",
							Leaf:  true,
							Class: "fund.ss.income_re",
						}, {
							Text:  "主账号关联设置",
							Cls:   "i-table",
							Leaf:  true,
							Class: "fund.ss.stock_relate",
						}, {
							Text:  "仓位委托修改",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.stock_root",
						}, {
							Text:  "历史查询",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.ss.stock_history",
						}, {
							Text:  "下单权重配置",
							Cls:   "i-table",
							Leaf:  true,
							Class: "fund.ss.account_weight",
						}, {
							Text:  "报表分析",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.ss.analyse",
						},
					},
				},
				{
					Text: "基金投资管理",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "账号管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.fd.account_manager",
						}, {
							Text:  "人员管理",
							Cls:   "i-trader",
							Leaf:  true,
							Class: "fund.fd.trader_manager",
						}, {
							Text:  "资金管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.fd.fund_manager",
						}, {
							Text:  "资金调整日志",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.fd.fund_log",
						}, {
							Text:  "我的仓位",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.fd.my_hold",
						}, {
							Text:  "仓位探针",
							Cls:   "i-log",
							Leaf:  true,
							Class: "fund.fd.fund_hold",
						}, {
							Text:  "报表分析",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.fd.analyse",
						}, {
							Text:  "收益率对比",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.fd.profit_compare",
						},
					},
				},
				{
					Text: "投资划拨系统",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "可投基金池",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "fund.zg.fund_pool",
						}, {
							Text:  "投资交易指令",
							Cls:   "i-event",
							Leaf:  true,
							Class: "fund.zg.trade_order",
						},
					},
				},
				{
					Text: "估值系统",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "账号管理",
							Cls:   "i-example",
							Leaf:  true,
							Class: "fund.zg.account_manager",
						}, {
							Text:  "资金/份额配置",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "fund.zg.fund",
						},
						{
							Text:  "资金到账确认",
							Cls:   "i-trader",
							Leaf:  true,
							Class: "fund.zg.fund_transfer_confirm",
						}, {
							Text:  "小F持仓信息",
							Leaf:  true,
							Class: "fund.zg.position",
						},
						{
							Text:  "净值信息",
							Leaf:  true,
							Class: "fund.zg.fund_info",
						},
						{
							Text:  "估值表备份",
							Leaf:  true,
							Class: "fund.zg.valuation_table",
						}, {
							Text:  "修改持仓信息",
							Cls:   "i-info-edit",
							Leaf:  true,
							Class: "fund.zg.modify_position",
						}, {
							Text:  "历史持仓信息",
							Leaf:  true,
							Class: "fund.zg.his_position",
						}, {
							Text:  "报表分析",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.zg.report",
						},
					},
				},
				{
					Text: "投后管理系统",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "基金明细",
							Leaf:  true,
							Class: "fund.zg.pi_info",
						}, {
							Text:  "净值归因报表",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "fund.zg.pi_report",
						}, {
							Text:  "证券持仓明细",
							Leaf:  true,
							Class: "fund.zg.gzb_info",
						}, {
							Text:  "持仓市值统计",
							Leaf:  true,
							Class: "fund.zg.position_porthion",
						}, {
							Text:  "浮动盈亏统计",
							Leaf:  true,
							Class: "fund.zg.profit_porthion",
						}, {
							Text:  "T+0收益统计",
							Leaf:  true,
							Class: "fund.zg.t_plus_zero",
						},
					},
				},
			},
		},
		{
			Text: "公募基金数据统计",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "基金开户统计",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.gm_open_bind",
				}, {
					Text:  "基金推广活动",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.fund_sale_activity",
				}, {
					Text:  "今日未确认数据统计",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.fund_sale_now_list",
				}, {
					Text:  "基金单日销量列表",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.fund_sale_day_list",
				}, {
					Text:  "TA基金销量分类筛查",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.fund_sale_ta_list",
				}, {
					Text:  "单只基金历史销量趋势",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.fund_sale_stat_fund",
				}, {
					Text:  "TA销量历史趋势",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.fund_sale_stat",
				}, {
					Text:  "基金区间销量统计",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.jy.fund_sale_statistics",
				}, {
					Text:  "基金业务统计",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.data_count",
				}, {
					Text:  "基金日报统计",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.fund_com_data",
				}, {
					Text:  "认购数据统计",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.buy_chart",
				}, {
					Text:  "日销售统计",
					Leaf:  true,
					Class: "fund.gm.sales_tatistics",
				}, {
					Text:  "用户收益查询",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.user_income",
				}, {
					Text:  "客户数据统计",
					Cls:   "i-chart",
					Leaf:  true,
					Class: "analyse.user_data_total",
				},
			},
		},
		{
			Text: "公募代销",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "换卡",
					Cls:   "i-hand",
					Leaf:  true,
					Class: "fund.gm.card_apply",
				},
				{
					Text:  "公募基金银行限额",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.bank_limit",
				},
				{
					Text:  "基金特点更新",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.day_operation",
				},
				{
					Text:  "重签约管理",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.rev_bank",
				},
				{
					Text:  "基金管理",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.fund_manager",
				},
				{
					Text:  "意见反馈管理",
					Cls:   "i-user-manager",
					Leaf:  true,
					Class: "fund.bbs.feedback",
				},
				{
					Text:  "行业监管报表",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "fund.gm.day_mon_report",
				},
				{
					Text:  "基金转换折扣配置",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "activity.gm_convert_dis",
				},
			},
		},
		{
			Text: "第五军团",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "小酒桌",
					Cls:   "i-form",
					Leaf:  true,
					Class: "fund.gen.gen_blog",
				}, {
					Text:  "发帖",
					Cls:   "i-form",
					Leaf:  true,
					Class: "fund.gen.fund_blog",
				}, {
					Text:  "评论管理",
					Cls:   "i-form",
					Leaf:  true,
					Class: "fund.gen.five_legion_manage",
				}, {
					Text:  "服务窗",
					Cls:   "i-form",
					Leaf:  true,
					Class: "fund.gen.blog_manage",
				}, {
					Text:  "权限管理",
					Cls:   "i-sys-admin-tools",
					Leaf:  true,
					Class: "fund.gen.admin_manage",
				}, {
					Text:  "APP文章管理",
					Cls:   "i-form",
					Leaf:  true,
					Class: "fund.gen.app_blog",
				}, {
					Text:  "APP评论管理",
					Cls:   "i-form",
					Leaf:  true,
					Class: "fund.gen.app_reply",
				}, {
					Text:  "禁言管理",
					Cls:   "i-form",
					Leaf:  true,
					Class: "fund.gen.banned",
				},
			},
		},
		{
			Text: "活动",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "直播/视频",
					Cls:   "i-form",
					Leaf:  true,
					Class: "activity.app.live_video",
				}, {
					Text:  "教学园地",
					Cls:   "i-form",
					Leaf:  true,
					Class: "activity.app.video_garden",
				}, {
					Text:  "图片链接发布",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "activity.factory.pub_images",
				}, {
					Text:  "活动管理",
					Cls:   "i-form",
					Leaf:  true,
					Class: "activity.manage",
				}, {
					Text:  "文章",
					Cls:   "i-tree-default",
					Leaf:  true,
					Class: "activity.readings",
				},

				{
					Text: "业务工具",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "APP推送",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.push.xg_push",
						}, {
							Text:  "新APP推送",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.push.new_push",
						}, {
							Text:  "微信推送",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.wechat_msg",
						}, {
							Text:  "小工具",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.user.tools",
						}, {
							Text:  "积分",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.points.points",
						},
					},
				},

				{
					Text: "鼎信APP",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "APP版本控制",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.version_control",
						}, {
							Text:  "开机图",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.app_load_ad",
						}, {
							Text:  "首页固定位广告",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.app_head_ad",
						}, {
							Text:  "首页轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.app_rotation_ad",
						}, {
							Text:  "公募固定位广告",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.gm_head_ad",
						}, {
							Text:  "公募轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.gm_rotation_ad",
						}, {
							Text:  "APP首页Banner",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.home_banner_ad",
						}, {
							Text:  "公募首页轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.gm_home_ad",
						}, {
							Text:  "私募顶部广告",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.sm_top_ad",
						}, {
							Text:  "私募首页轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.sm_home_ad",
						}, {
							Text:  "基金经理轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.fund_manager_ad",
						}, {
							Text:  "主题选基轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.fund_theme_ad",
						}, {
							Text:  "定投专区轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.fund_regular_zone",
						}, {
							Text:  "公募详情页Banner广告",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.fund_detail_banner_ad",
						}, {
							Text:  "基金月报",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.jy.fund_month_report",
						}, {
							Text:  "推荐公募",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.jy.fund_recommend",
						}, {
							Text:  "精选定投",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.jy.fund_regular_sift_list",
						}, {
							Text:  "定投学院",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.fund_regular_school",
						}, {
							Text:  "基金经理",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.jy.fund_manager",
						}, {
							Text:  "热销基金",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.jy.fund_hot",
						}, {
							Text:  "主题选基",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.jy.fund_theme",
						}, {
							Text:  "认购专区",
							Cls:   "i-form",
							Leaf:  true,
							Class: "fund.jy.fund_subscribe_zone",
						}, {
							Text:  "教学园地轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.garden_ad",
						}, {
							Text:  "基金公司管理",
							Cls:   "i-tree-default",
							Leaf:  true,
							Class: "activity.app.fund_fc_manager",
						},
					},
				},

				{
					Text: "天狼APP",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "首页轮播",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.tl_app.tl_app_rotation_ad",
						}, {
							Text:  "轮播消息",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.tl_app.tl_app_rotation_msg",
						},
					},
				},
			},
		},

		{
			Text: "财经数据",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text: "天狼软件",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "说明书",
							Cls:   "i-form",
							Leaf:  true,
							Class: "activity.app.instructions",
						}, {
							Text:  "融资融券",
							Cls:   "i-document-invoice",
							Leaf:  true,
							Class: "stock.ziquan",
						}, {
							Text:  "股票质押",
							Cls:   "i-document-invoice",
							Leaf:  true,
							Class: "stock.stock_right",
						},
						{
							Text: "央行数据",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "社会融资规模",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.cbank.society_raising_scale",
								}, {
									Text:  "货币统计概览",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.cbank.money_statistics_summary",
								}, {
									Text:  "金融机构信贷收支统计",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.cbank.credit_statistics",
								}, {
									Text:  "企业商品价格指数",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.cbank.corp_goods_price",
								}, {
									Text:  "宏观经济数据",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.cbank.macroeconomy",
								},
							},
						},
						{
							Text: "期权数据",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "中债国债",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.qq.zzgz_rates",
								},
							},
						},
						{
							Text: "沪深港通",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "成交统计",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "stock.hgt.summary",
								}, {
									Text:  "前十大活跃个股",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "stock.hgt.rank",
								}, {
									Text:  "沪深港通持股",
									Cls:   "i-page-white-link",
									Leaf:  true,
									Class: "stock.hgt.positions",
								},
							},
						},
						{
							Text: "股本结构",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "A股股本结构",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.arch.astock",
								}, {
									Text:  "B股股本结构",
									Cls:   "i-table-relationship",
									Leaf:  true,
									Class: "stock.arch.bstock",
								}, {
									Text:  "股权质押冻结",
									Cls:   "i-slash",
									Leaf:  true,
									Class: "stock.arch.fridge",
								}, {
									Text:  "流通股东持股",
									Cls:   "i-page-white-link",
									Leaf:  true,
									Class: "stock.arch.tardableHolder",
								}, {
									Text:  "股东持股",
									Cls:   "i-page-white-link",
									Leaf:  true,
									Class: "stock.arch.mainSHListNew",
								}, {
									Text:  "A股自由流通盘",
									Cls:   "i-page",
									Leaf:  true,
									Class: "stock.arch.tradableMarket",
								}, {
									Text:  "A股自由流动盘",
									Cls:   "i-market-flow",
									Leaf:  true,
									Class: "stock.arch.capital",
								},
							},
						},

						{
							Text: "板块管理",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "股票",
									Cls:   "i-sitemap",
									Leaf:  true,
									Class: "stock.plate.stock",
								},
							},
						},
					},
				},
				{
					Text: "公募基金",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text: "公募基金",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text: "基金概况",
									Leaf: false,
									Cls:  "folder",
									Children: []Node{
										{
											Text:  "赎回清算日期",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.fund_trade_days",
										},
										{
											Text:  "基本信息",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_FundArchives",
										},
										{
											Text:  "基金公司",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_InvestAdvisorOutline",
										},
										{
											Text:  "净值变动",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_NVChange",
										},
										{
											Text:  "份额变动",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_SharesChange",
										},
										{
											Text:  "基金分红",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_Dividend",
										},
										{
											Text:  "拆分折算",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_SharesSplit",
										},
										{
											Text:  "基金净值",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_NetValue",
										},
										{
											Text:  "净值表现",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_NetValuePerformanceHis",
										},

										{
											Text:  "货基净值表现",
											Cls:   "i-tree-default",
											Leaf:  true,
											Class: "fund.jy.MF_MMYieldPerformance",
										},
									},
								},
							},
						},

						{
							Text: "季报分析",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "十大重仓股",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.MF_KeyStockPortfolio",
								},
								{
									Text:  "全行业股票",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.MF_StockPortfolioDetail",
								},
								{
									Text:  "基金行业统计",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.MF_InvestIndustry",
								},
								{
									Text:  "资产配置",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.MF_AssetAllocation",
								},
								{
									Text:  "十大流通股东",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.LC_StockHoldingSt",
								},
							},
						},

						{
							Text: "基金评价",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "基金打分",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.score.achievement_score",
								},
								{
									Text:  "风险分析",
									Cls:   "i-chart",
									Leaf:  true,
									Class: "fund.jy.risk_analysis",
								},
							},
						},

						{
							Text: "鼎信基金",
							Leaf: false,
							Cls:  "folder",
							Children: []Node{
								{
									Text:  "鼎信严选",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.choice_fund",
								},
								{
									Text:  "鼎信指数",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.dx_fund_index",
								},
								{
									Text:  "风险协议",
									Cls:   "i-tree-default",
									Leaf:  true,
									Class: "fund.jy.dx_risk_deal_link",
								},
							},
						},
					},
				},
			},
		},

		{
			Text: "系统管理员",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text: "监控统计运维",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "登录会话",
							Leaf:  true,
							Cls:   "i-user-manager",
							Class: "tools.session_manager",
						}, {
							Text:  "服务器监控",
							Cls:   "i-sys-monitor",
							Leaf:  true,
							Class: "tools.server_watch",
						}, {
							Text:  "站点访问统计",
							Cls:   "i-chart",
							Leaf:  true,
							Class: "analyse.chart",
						}, {
							Text: "公募服务器监控",
							//Cls: "i-sys-monitor",
							Leaf:  true,
							Class: "tools.gm_day_wc",
						},
					},
				},
				{
					Text: "工作流设置",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "用户组设置",
							Cls:   "i-cmp",
							Leaf:  true,
							Class: "wf.user_group",
						}, {
							Text:  "工作流设置",
							Cls:   "i-pkg",
							Leaf:  true,
							Class: "wf.wf",
						}, {
							Text:  "模板设置",
							Cls:   "i-cmp",
							Leaf:  true,
							Class: "wf.tik_template",
						},
					},
				},

				{
					Text:  "超级数据库",
					Cls:   "i-sys-admin-tools",
					Leaf:  true,
					Class: "tools.admin_database",
				},
				{
					Text:  "交易日历",
					Cls:   "i-equalizer",
					Leaf:  true,
					Class: "tools.calendar",
				},
				{
					Text:  "页面访问统计",
					Leaf:  true,
					Class: "sys.erp_visit_page_log",
				},
			},
		},
		{
			Text: "k50管理",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text: "组合管理",
					Leaf: false,
					Cls:  "folder",
					Children: []Node{
						{
							Text:  "组合调整",
							Leaf:  true,
							Class: "fund.k50.k50",
						},
					},
				},
				{
					Text:  "k50走势",
					Leaf:  true,
					Class: "fund.k50.trend",
				},
			},
		},

		{
			Text: "代理商管理",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "代理商管理",
					Leaf:  true,
					Class: "fund.bpm.agent",
				}, {
					Text:  "代理员工明细",
					Leaf:  true,
					Class: "fund.bpm.staff",
				},
			},
		},

		{
			Text: "客户信息表",
			Leaf: false,
			Cls:  "folder",
			Children: []Node{
				{
					Text:  "客户信息表",
					Leaf:  true,
					Class: "fund.bpm.customer",
				}, {
					Text:  "客户身份证信息表",
					Leaf:  true,
					Class: "fund.user.icard",
				},
			},
		},
	},
}

var pidCount = 0
var existsIds []string

func NavParse(ct Node) {

	// 函数 parse_tree
	Nav.Leaf = Nav.Leaf || false
	for _, node := range ct.Children {
		if node.Class != "" {
			node.ID = node.Class
		} else {
			pidCount++
			node.ID = strconv.Itoa(pidCount)
		}
		check_node(existsIds, node)
		//vct := ct
		//curv := node
		//vct.Children = []Node{} // 将数组变空
		//curv.
	}
}

func check_node(strList []string, node Node) {
	if isExist(node.ID, strList) { // 存在
		panic("导航Error:")
	}
	existsIds = append(existsIds, node.ID)
	if len(node.Children) == 0 {
		if node.Class == "" {
			panic("导航Error:, 缺少class节点")
		}
	}
}

// true 存在; false 不存在
func isExist(str string, strList []string) bool {
	for _, v := range strList {
		if str == v {
			return true
		}
	}
	return false
}

// 获取 _nav_T

type _nav_T struct {
	Name  string `json:"name"`
	ID    string `json:"ID"`
	Class string `json:"class"`
}

var NavT []_nav_T
var count = 0
var depth = 1

func getNavT(node Node) {
	if node.Children == nil {
		NavT = append(NavT, _nav_T{
			Name:  node.Class,
			ID:    node.Class,
			Class: node.Class,
		})
		return
	}
	for _, v := range node.Children {
		getNavT(v)
	}
	count++
	NavT = append(NavT, _nav_T{
		Name: fmt.Sprintf("m%d", count),
	})

	//if node.Leaf {
	//	NavT = append(NavT, _nav_T{
	//		Name: fmt.Sprintf("m%d", count),
	//	})
	//}
	//for _, v := range node.Children {
	//	getNavT(v)
	//}

}
