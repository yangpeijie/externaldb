package transaction

// mapping
const (
	TxRecordMapping = `{
    "mappings":{
        "properties":{
            "height":{
                "type":"long"
            },
            "block_time":{
                "type":"long"
            },
            "block_hash":{
                "type":"keyword"
            },
            "success":{
                "type":"boolean"
            },
            "is_para":{
                "type":"boolean"
            },
            "index":{
                "type":"long"
            },
            "hash":{
                "type":"keyword"
            },
            "from":{
                "type":"keyword"
            },
            "to":{
                "type":"keyword"
            },
            "next":{
                "type":"keyword"
            },
            "execer":{
                "type":"keyword"
            },
            "amount":{
                "type":"long"
            },
            "fee":{
                "type":"long"
            },
            "action_name":{
                "type":"keyword"
            },
            "group_count":{
                "type":"long"
            },
            "is_withdraw":{
                "type":"boolean"
            },
            "options":{
                "properties":{
                    "symbol":{
                        "type":"keyword"
                    },
                    "to":{
                        "type":"keyword"
                    },
                    "exec_name":{
                        "type":"keyword"
                    },
                    "amount":{
                        "type":"long"
                    },
                    "name":{
                        "type":"keyword"
                    },
                    "introduction":{
                        "type":"keyword"
                    },
                    "total":{
                        "type":"long"
                    },
                    "price":{
                        "type":"long"
                    },
                    "owner":{
                        "type":"keyword"
                    },
                    "category":{
                        "type":"long"
                    },
                    "note":{
                        "type":"keyword"
                    }
                }
            }
        }
    }
}`
)
