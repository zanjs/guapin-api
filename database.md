## 电子商务平台中订单模块(Order)数据库设计明细

订单表 (order)
|-- 自动编号（order_id, 自增长主键）
|-- 订单单号（order_no, 唯一值，供客户查询）
|-- 商店编号（shop_id, 商店表自动编号）
|-- 订单状态 (order_status,未付款,已付款,已发货,已签收,退货申请,退货中,已退货,取消交易)
|-- 商品数量 (product_count, 商品项目数量，不是商品)
|-- 商品总价 (product_amount_total)
|-- 订单金额 (order_amount_total，实际付款金额)
|-- 运费金额 (logistics_fee)
|-- 是否开箱验货 (is_unpacking_inspection)
|-- 是否开票（是否开具发票）
|-- 发票编号 (订单发票表自动编号)
|-- 收货地址编号 (address_id, 收货地址表自动编号)
|-- 订单物流编号 (orderlogistics_id, 订单物流表自动编号)
|-- 订单支付渠道 (pay_channel)
|-- 订单支付单号 (out_trade_no/escrow_trade_no,第三方支付流水号)
|-- 创建时间 (下单时间)
|-- 付款时间
|-- 发货时间
|-- 客户编号 (user_id，用户表自动编号)
|-- 客户备注
|-- 订单结算状态 (order_settlement_status，货到付款、分期付款会用到)
|-- 订单结算时间 (order_settlement_time)

订单发票表 (order_invoice)
|-- 自动编号 (invoice_id)
|-- 订单编号 (order_id)
|-- 是否增值税发票 (is_vat, 普通发票,增值发票)
|-- 发票抬头名称 (invoice_title)
|-- 发票抬头内容 (invoice_content)
|-- 发票金额 (invoice_amount)
|-- 发票税号 (invoice_tax_no)
|-- 开票税金 (invoice_tax)
|-- 公司名称[增值税] (vat_company_name)
|-- 公司地址[增值税] (vat_company_address)
|-- 联系电话[增值税] (vat_telphone)
|-- 开户银行[增值税] (vat_bank_name)
|-- 银行帐号[增值税] (vat_bank_account)
|-- 开票时间 (created_time)


订单物流表 (order_logistics)
|-- 自动编号 (orderlogistics_id)
|-- 订单编号 (order_id, 订单表自动编号)
|-- 物流单号 (express_no, 发货快递单号)
|-- 收货人姓名 (consignee_realname, 收货地址表可能更新或删除，因此要在这里记录)
|-- 联系电话 (consignee_telphone, 收货地址表可能更新或删除，因此要在这里记录)
|-- 备用联系电话 (consignee_telphone2, 收货地址表可能更新或删除，因此要在这里记录)
|-- 收货地址 (consignee_address, 收货地址表可能更新或删除，因此要在这里记录)
|-- 邮政编码 (consignee_zip, 收货地址表可能更新或删除，因此要在这里记录)
|-- 物流方式（logistics_type, ems, express）
|-- 物流商家编号 (logistics_id，物流商家表自动编号)
|-- 物流发货运费 (logistics_fee，显示给客户的订单运费)
|-- 快递代收货款费率 (agency_fee, 快递公司代收货款费率，如货值的2%-5%，一般月结)
|-- 物流成本金额 (delivery_amount, 实际支付给物流公司的金额)
|-- 物流状态 (orderlogistics_status)
|-- 物流结算状态 (logistics_settlement_status, 未结算,已结算,部分结算)
|-- 物流最后状态描述 (logistics_result_last)
|-- 物流描述 (logistics_result)
|-- 发货时间 (logistics_create_time)
|-- 物流更新时间 (logistics_update_time)
|-- 物流结算时间 (logistics_settlement_time)
|-- 物流支付渠道
|-- 物流支付单号
|-- 物流公司已对账状态 (reconciliation_status，已对账,未对账)
|-- 物流公司对账日期 (reconciliation_time)
设计说明：收货地址可能被修改、删除等，因此这里要记录发货时用户的收货地址，这样就算后来收货地址被删除了，用户在查看历史订单的时候也依然能看到收货地址的快照信息。


订单退货表 (order_returns)
|-- 自动编号 (order_returns_id)
|-- 退货编号 (returns_no，供客户查询)
|-- 订单编号 (order_id, 订单表自动编号)
|-- 物流单号 (express_no, 退货物流单号)
|-- 收货人姓名 (consignee_realname)
|-- 联系电话 (consignee_telphone)
|-- 备用联系电话 (consignee_telphone2)
|-- 收货地址 (consignee_address)
|-- 邮政编码 (consignee_zip)
|-- 物流方式（logistics_type, ems, express）
|-- 物流商家编号
|-- 物流发货运费 (logistics_fee，退货运费)
|-- 物流状态 (orderlogistics_status)
|-- 物流最后状态描述
|-- 物流描述
|-- 物流更新时间
|-- 物流发货时间
|-- 退货类型 (returns_type, 全部退单,部分退单)
|-- 退货处理方式 (handling_way, PUPAWAY:退货入库;REDELIVERY:重新发货;RECLAIM-REDELIVERY:不要求归还并重新发货; REFUND:退款; COMPENSATION:不退货并赔偿)
|-- 退款金额 (returns_amount)
|-- 退货销售员承担的费用 (seller_punish_fee)
|-- 退货申请时间 (return_submit_time)
|-- 退货处理时间 (handling_time)
|-- 退货原因
设计说明：退货可能被修改、删除等，因此这里要记录退货时商家的退货地址信息，


订单商品详情表 (order_detail)
|-- 自动编号
|-- 订单编号
|-- 商品编号
|-- 商品名称 (product_name, 商品可能删除,所以这里要记录，不能直接读商品表)
|-- 商品价格 (product_price, 商品可能删除,所以这里要记录)
|-- 商品型号 (product_marque，前台展示给客户)
|-- 商品条码 (product_store_barcode, 商品仓库条码)
|-- 商品型号信息 (product_mode_desc，记录详细商品型号，如颜色、规格、包装等)
|-- 商品型号参数 (product_mode_params, JSON格式，记录单位编号、颜色编号、规格编号等)
|-- 折扣比例 (discount_rate 打几折)
|-- 折扣金额 (discount_amount)
|-- 购买数量 (number)
|-- 小计金额 (subtotal)
|-- 商品是否有效 (is_product_exists)
|-- 客户商品备注 (remark)
设计说明：商品可能被修改、删除等，因此这里要记录下单时用户关注的商品交易摘要信息，如价格、数量、型号、型号参数等。这样就算后来商品被删除了，用户在查看历史订单的时候也依然能看到商品的快照信息。


收货地址表 (delivery_address)
|-- 自动编号 (address_id)
|-- 用户编号 (user_id, 用户表自动编号)
|-- 收件人姓名 (realname)
|-- 联系电话 (telphone)
|-- 备用联系电话 (telphone2)
|-- 国家 (country)
|-- 省份 (province)
|-- 城市 (city)
|-- 地区 (area)
|-- 街道/详细收货地址 (street)
|-- 邮政编码 (zip)
|-- 是否默认收货地址 (is_default_address)
|-- 创建时间 (created_time)


购物车表 (shoppingcart)
|-- 自动编号 (id)
|-- 用户编号 (user_id)
|-- 商店编号 (shop_id)
|-- 商品编号 (product_id)
|-- 是否有效 (is_product_exists)
|-- 购买数量 (number)
|-- 创建时间 (created_time)
设计说明：商品价格和小计金额是要通过实时关联商品表来读取和计算，因为商户可能会更改商品价格，或者商品已售罄，或者商品已下架等，因此这里只需要记录商品id就可以，商品价格等要实时从商品表读取。

 

===============================用于电话营销的订单模块的扩展设计======================================

订单业务审核流程表 (order_auditbiz)
|-- 自动编号 (order_auditbiz_id)
|-- 订单编号 (order_id)
|-- 订单状态 (0:未审核或发起交易;1:交易完成;20:核单通过;24:核单失败;30:已发货;未签收;34:仓库退回;40:座席取消;41:买家取消;42:逾期取消;43:订单无效取消;50:客户签收;54:客户拒签;55:客户退货)
|-- 销售员直接确认订单(不需要订单审核员确认，直接强制审核通过，如客户退货则销售员必须承担退货运费) (is_seller_risk_confirm)
|-- 订单退货,销售员是否承担运费 (is_seller_punish
_logistics_fee)
|-- 销售员是否提成 (is_seller_commission)
|-- 销售员提成比例 (seller_commission_rate, 无提成则填0)
|-- 销售员提成金额 (seller_commission_amount)
|-- 销售员订单备注（seller_remark，给订单审核员看的备注）
|-- 订单审核员订单备注 (confirmer_remark，给仓管看的备注)
|-- 仓管备注（storekeeper_returnback_remark，仓管退给订单审核员看的备注）
|-- 财务备注 (cashier_remark, 财务给销售员看的备注)
|-- 销售员用户编号 (seller_uid)
|-- 订单审核员用户编号 (auditor_uid)
|-- 收款人用户编号 (cashier_uid，收款人不一定是财务)
|-- 财务用户编号 (accountant_uid, 财务人员用户编号)
|-- 订单来源 (order_source, 销售下单,内部购买)
|-- 订单审核员审核时间 (auditor_audited_time)
|-- 仓管员审核时间 (storekeeper_audited_time)
|-- 财务审核时间 (accountant_audited_time)


订单提成表 (order_commission)
|-- 自动编号 (order_commission_id)
|-- 订单编号 (order_id)
|-- 销售员用户编号 (seller_uid)
|-- 提成金额 (commission_amount)
|-- 结算状态 (settlement_status)
|-- 结算时间 (settlement_time)
|-- 财务人员用户编号 (cashier_uid)


订单调度表 (order_dispatch)
|-- 自动编号
|-- 订单编号
|-- 被调度的营销人员用户编号 (from_seller_uid)
|-- 营销人员用户编号 (to_seller_uid)
|-- 调度原因 (dispatch_reason)
|-- 调度管理员 (diapatch_admin_uid)
|-- 调度日期 (created_time)


数据库设计原则是：

1. 为提高读的性能，尽可能把写的操作拆分到另一张表，因为对表的更新操作会导致锁表，会降低数据表的读取的性能。
2. 交易时一些关联信息可能在后来会被修改或删除，如商品、收货地址等，因此要在订单中记录交易时的商品信息和收货地址，一边后来商品或收货地址被删除的时候，依然能在历史订单中看到快照信息。
3. 不要怕拆分成很多表，读的时候多张表关联读取，会比读取一张字段非常多的数据量庞大的表效率高很多。