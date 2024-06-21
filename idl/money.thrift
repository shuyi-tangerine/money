include "base.thrift"

namespace go tangerine.money

struct GenFinanceDetailIDsRequest {
    1: optional i8 amount = 1   // 需要的数量
    255: optional base.RPCRequest Base
}

struct GenFinanceDetailIDsResponse {
    1: optional list<i64> ids
    255: required base.RPCResponse Base
}

struct AddFinanceDetailRequest {
    1: required i64 app_id
    2: required i64 amount
    3: required i64 operated_type
    4: required i64 operated_at
    5: required string operated_by
    6: optional string extra  // 要求 JSON string
    7: required string created_by
    8: optional i64 finance_detail_id = 0  // 填写的时候，幂等操作，不填写就创建然后自动生成
    255: optional base.RPCRequest Base
}

struct AddFinanceDetailResponse {
    1: optional FinanceDetailInfo info
    255: required base.RPCResponse Base
}


struct ListFinanceDetailRequest {
    1: optional i64 app_id = 0
    5: optional TimeRange operated_at
    6: required i64 offset
    7: required i64 limit
    255: optional base.RPCRequest Base
}

struct ListFinanceDetailResponse {
    1: optional i64 offset = 0
    2: optional i32 total = 0
    3: optional list<FinanceDetailInfo> finance_details
    255: required base.RPCResponse Base
}

struct FinanceDetailInfo {
    1: required i64 id
    2: required i64 finance_detail_id
    3: required i64 app_id
    4: required i64 amount
    5: required i64 operated_type
    6: required i64 operated_at
    7: required string operated_by
    8: optional string extra
    9: required i64 created_at
    10: required string created_by
    11: required i64 updated_at
    12: required string updated_by
}

struct TimeRange {
    1: required i64 s  // start
    2: required i64 e  // end
}

service MoneyHandler {
    // 生成 ID
    GenFinanceDetailIDsResponse GenFinanceDetailIDs(1:GenFinanceDetailIDsRequest req)
    // 增加一条资金明细
    AddFinanceDetailResponse AddFinanceDetail(1:AddFinanceDetailRequest req)
    // 获取资金明细列表
    ListFinanceDetailResponse ListFinanceDetail(1:ListFinanceDetailRequest req)
}