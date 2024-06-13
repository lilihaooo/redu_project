package craw

import (
	"encoding/json"
	"fmt"
	"log"
	"redu/flag/craw/receive"
	"redu/global"
	"redu/model"
)

func CategoryTask() {
	categories := []receive.Category{}
	jsonData := `[
    {
        "id": "12",
        "level": "1",
        "parent_id": "0",
        "name": "食品饮料",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/jyXIeRHeeD4FnuspuDZMepGgMi7SHR1S.png",
        "num": 0,
        "childer": [
            {
                "id": "28",
                "level": "2",
                "parent_id": "12",
                "name": "营养保健",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/wtf4RLFv1m0M2jNBjH4msotMmcskrYgz.png",
                "num": 0,
                "childer": [
                    {
                        "id": "45",
                        "level": "3",
                        "parent_id": "28",
                        "name": "膳食保健",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/08UgaufzStZBayNBRb3TWU4DYSd00UjG.png",
                        "num": 0
                    },
                    {
                        "id": "44",
                        "level": "3",
                        "parent_id": "28",
                        "name": "营养膳补",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/tXDfFSjSxonmANbCMhljHzHc1fvHVNgp.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "27",
                "level": "2",
                "parent_id": "12",
                "name": "乳品冲饮",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221109/iJOer33-ikTG08INT05GWRNnGwdAOrm-.png",
                "num": 0,
                "childer": [
                    {
                        "id": "43",
                        "level": "3",
                        "parent_id": "27",
                        "name": "咖啡冲饮",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/1CNPAPw8T0kQ6HLcZorVeeoDqyf7xJX1.png",
                        "num": 0
                    },
                    {
                        "id": "42",
                        "level": "3",
                        "parent_id": "27",
                        "name": "乳制品/奶制品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/AmF50pAypFMXn-5IomnYxt0su-p2vDiA.png",
                        "num": 0
                    },
                    {
                        "id": "41",
                        "level": "3",
                        "parent_id": "27",
                        "name": "饮料",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yJFIPXnem3eAPaa1Y_1CQeqIJsqcYTDb.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "26",
                "level": "2",
                "parent_id": "12",
                "name": "休闲零食",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221109/UPdZHnQmo7fibUEe8SiF2A-xwPMQzfNM.png",
                "num": 0,
                "childer": [
                    {
                        "id": "40",
                        "level": "3",
                        "parent_id": "26",
                        "name": "海苔/即食海产",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2_HgdmMqpbAj1pg46piXHNqqDK4vzoXW.png",
                        "num": 0
                    },
                    {
                        "id": "39",
                        "level": "3",
                        "parent_id": "26",
                        "name": "豆干",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/3lY62-rKsS23Lzi6HRl-JXgxUVGEEpb_.png",
                        "num": 0
                    },
                    {
                        "id": "37",
                        "level": "3",
                        "parent_id": "26",
                        "name": "糖果",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/TjUH46B8r9rNA14SBdbJcl1umUpjN2ZA.png",
                        "num": 0
                    },
                    {
                        "id": "36",
                        "level": "3",
                        "parent_id": "26",
                        "name": "蜜饯/果干",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5L8ig7WVAq-jTuG6QjsVjWPIuogl9c4T.png",
                        "num": 0
                    },
                    {
                        "id": "35",
                        "level": "3",
                        "parent_id": "26",
                        "name": "膨化食品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/vPQWILbRyHxv2SHdGkVGu8pvMAvqoUCY.png",
                        "num": 0
                    },
                    {
                        "id": "34",
                        "level": "3",
                        "parent_id": "26",
                        "name": "巧克力",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/PGLuHP5BQYwSVXtOnLzdnm9moKl_AUI_.png",
                        "num": 0
                    },
                    {
                        "id": "33",
                        "level": "3",
                        "parent_id": "26",
                        "name": "枣制品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/RrPMHO-fOntHttxhx0HnBEAKmj-EbJHi.png",
                        "num": 0
                    },
                    {
                        "id": "32",
                        "level": "3",
                        "parent_id": "26",
                        "name": "糕点",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/z1242JfEKwMT0T09eoQrsNk-aHGaNtVq.png",
                        "num": 0
                    },
                    {
                        "id": "31",
                        "level": "3",
                        "parent_id": "26",
                        "name": "饼干",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/GQbDXymzvkJaecKUhgz1LzIarkvW_8xA.png",
                        "num": 0
                    },
                    {
                        "id": "30",
                        "level": "3",
                        "parent_id": "26",
                        "name": "坚果",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/-8236XWpI5UkC2WZURMOaZQhQztz4tXt.png",
                        "num": 0
                    },
                    {
                        "id": "29",
                        "level": "3",
                        "parent_id": "26",
                        "name": "小零食",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/D67wj4w4oVoYRHAP-6n3dEXl9In7Oin8.png",
                        "num": 0
                    },
                    {
                        "id": "226",
                        "level": "3",
                        "parent_id": "26",
                        "name": "肉干肉脯/卤味",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/X-KdktKvuLsfFXGlFZbXL6hz0qtwkMRl.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "132",
                "level": "2",
                "parent_id": "12",
                "name": "粮油速食",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/3S4IwVVaA2Gm-ykFPBQJwEum9CnYsAYS.png",
                "num": 0,
                "childer": [
                    {
                        "id": "133",
                        "level": "3",
                        "parent_id": "132",
                        "name": "粮油米面",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/p-qt5s0C_x_hS6UXMBPbJuvVIy2puSav.png",
                        "num": 0
                    },
                    {
                        "id": "134",
                        "level": "3",
                        "parent_id": "132",
                        "name": "南北干货",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/AglSGagX-TYEBKd8sTUac2xnWkS0yqDi.png",
                        "num": 0
                    },
                    {
                        "id": "135",
                        "level": "3",
                        "parent_id": "132",
                        "name": "调味品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/YRvzW6bPihX6LUqn4lTQOI9038_bApsA.png",
                        "num": 0
                    },
                    {
                        "id": "228",
                        "level": "3",
                        "parent_id": "132",
                        "name": "食品添加剂/辅料",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/oFad2myPIcH_i2O9CZ52as0ziPAQyfyE.png",
                        "num": 0
                    },
                    {
                        "id": "229",
                        "level": "3",
                        "parent_id": "132",
                        "name": "方便速食",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/0awmTMajlmvdwBTfiCIxRVREvgmbhP9J.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "172",
                "level": "2",
                "parent_id": "12",
                "name": "传统滋补",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/K4KU3SeOqiaIwFO4-XC0pns8azqfylQO.png",
                "num": 0,
                "childer": [
                    {
                        "id": "230",
                        "level": "3",
                        "parent_id": "172",
                        "name": "枸杞",
                        "status": "1",
                        "image": "",
                        "num": 0
                    },
                    {
                        "id": "231",
                        "level": "3",
                        "parent_id": "172",
                        "name": "阿胶",
                        "status": "1",
                        "image": "",
                        "num": 0
                    },
                    {
                        "id": "232",
                        "level": "3",
                        "parent_id": "172",
                        "name": "蜂蜜",
                        "status": "1",
                        "image": "",
                        "num": 0
                    },
                    {
                        "id": "233",
                        "level": "3",
                        "parent_id": "172",
                        "name": "燕窝",
                        "status": "1",
                        "image": "",
                        "num": 0
                    },
                    {
                        "id": "234",
                        "level": "3",
                        "parent_id": "172",
                        "name": "参茸贵细",
                        "status": "1",
                        "image": "",
                        "num": 0
                    },
                    {
                        "id": "235",
                        "level": "3",
                        "parent_id": "172",
                        "name": "养生原料",
                        "status": "1",
                        "image": "",
                        "num": 0
                    },
                    {
                        "id": "236",
                        "level": "3",
                        "parent_id": "172",
                        "name": "其他滋补品",
                        "status": "1",
                        "image": "",
                        "num": 0
                    },
                    {
                        "id": "238",
                        "level": "3",
                        "parent_id": "172",
                        "name": "食疗滋补",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/YS0wh9KQH1UDIBO6B6WHQ1enYeHeTJhT.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "14",
        "level": "1",
        "parent_id": "0",
        "name": "日用家居",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/xzWPAclrCMZacX4MdYsywDHwaoir7FFU.png",
        "num": 0,
        "childer": [
            {
                "id": "53",
                "level": "2",
                "parent_id": "14",
                "name": "家具家装",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/Dikq_8nIJALYwzFbM2MhwmF4X771rlFT.png",
                "num": 0,
                "childer": [
                    {
                        "id": "163",
                        "level": "3",
                        "parent_id": "53",
                        "name": "电子电工",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/wm0OVorAhLW9NQfu4SCjr_MTFCU3gXOa.png",
                        "num": 0
                    },
                    {
                        "id": "164",
                        "level": "3",
                        "parent_id": "53",
                        "name": "基础建材",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/4iFgdHz5FdxLEEmlqIw9JmxtIIHfPnYa.png",
                        "num": 0
                    },
                    {
                        "id": "165",
                        "level": "3",
                        "parent_id": "53",
                        "name": "五金工具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/R9xut2gHosv3AEnV2Q_qSgMUxw1xIaag.png",
                        "num": 0
                    },
                    {
                        "id": "166",
                        "level": "3",
                        "parent_id": "53",
                        "name": "家装主材",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/vbI2rWfGxotNEnDzOYpmRpEihBZINMrk.png",
                        "num": 0
                    },
                    {
                        "id": "167",
                        "level": "3",
                        "parent_id": "53",
                        "name": "灯饰照明",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5Swq0T7dBL14c4hUSbZyzkgy6A1KlCOk.png",
                        "num": 0
                    },
                    {
                        "id": "168",
                        "level": "3",
                        "parent_id": "53",
                        "name": "家具装饰",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/XXREC3ntFtqp0LNPRF5x5RlOPYzuopJz.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "50",
                "level": "2",
                "parent_id": "14",
                "name": "日用百货",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/LgUiFuSn0fqra9hsA8AVAcqxSVrD1X0g.png",
                "num": 0,
                "childer": [
                    {
                        "id": "308",
                        "level": "3",
                        "parent_id": "50",
                        "name": "生活日用",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/tItnhtIb6JxxM3_fJPkFGMsQ9wolL7hx.png",
                        "num": 0
                    },
                    {
                        "id": "310",
                        "level": "3",
                        "parent_id": "50",
                        "name": "办公文具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/N-nm6Gn7wavgVFR7pqjJD9spbJUzdL7k.png",
                        "num": 0
                    },
                    {
                        "id": "311",
                        "level": "3",
                        "parent_id": "50",
                        "name": "创意家居",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/S27trIX_XCzYM374lt3Csfl-mzpyPxvo.png",
                        "num": 0
                    },
                    {
                        "id": "374",
                        "level": "3",
                        "parent_id": "50",
                        "name": "收纳整理",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/W6XpTUGuFqMk9NtbBWye7FFwQ_H4ARhi.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "222",
                "level": "2",
                "parent_id": "14",
                "name": "生活家电",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/BA1xO5OJOnyYpCbYFQihlEzY4uNdaCcM.png",
                "num": 0,
                "childer": [
                    {
                        "id": "223",
                        "level": "3",
                        "parent_id": "222",
                        "name": "个护家电",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/vWm9SDbnGJwl-nLv7G5ld0IrjXalext-.png",
                        "num": 0
                    },
                    {
                        "id": "224",
                        "level": "3",
                        "parent_id": "222",
                        "name": "生活电器",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/WNI8DlUXb1VXTpitTSvjPgrU9nPppAH1.png",
                        "num": 0
                    },
                    {
                        "id": "225",
                        "level": "3",
                        "parent_id": "222",
                        "name": "厨房小电",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/XBDs6brFlmbaAA3BvR9pMhnO14iUUyf9.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "162",
                "level": "2",
                "parent_id": "14",
                "name": "厨具烹饪",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/zFeWw06Lkmo72cpkDKfGC9S2BD4H__J0.png",
                "num": 0,
                "childer": [
                    {
                        "id": "312",
                        "level": "3",
                        "parent_id": "162",
                        "name": "餐饮具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Og4U-zhh_vnQB9QHV1zkkDSf9bkiVClb.png",
                        "num": 0
                    },
                    {
                        "id": "313",
                        "level": "3",
                        "parent_id": "162",
                        "name": "烹饪用具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5ZiFQW12zKqQ6TZcZKfUcmrXT1MaRb_d.png",
                        "num": 0
                    },
                    {
                        "id": "314",
                        "level": "3",
                        "parent_id": "162",
                        "name": "厨房配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/lXkUOomB8C2VyLyp0GvJ4_o_daQyTI_r.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "488",
        "level": "1",
        "parent_id": "0",
        "name": "家庭清洁",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20220928/x3QZW5i-1ey4BW3r7WhbTaWq_pyQTGks.jpg",
        "num": 0,
        "childer": [
            {
                "id": "454",
                "level": "2",
                "parent_id": "488",
                "name": "纸品/湿巾",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "456",
                        "level": "3",
                        "parent_id": "454",
                        "name": "抽纸/纸巾/手帕纸",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/FQIpikJKQgAy_TT0G_mlVQnHkKz9VqJT.png",
                        "num": 0
                    },
                    {
                        "id": "457",
                        "level": "3",
                        "parent_id": "454",
                        "name": "棉柔巾/洗脸巾",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/dGNKBg3ltZBQ4HSOuPmesM5VdcljEn7y.png",
                        "num": 0
                    },
                    {
                        "id": "458",
                        "level": "3",
                        "parent_id": "454",
                        "name": "湿巾/湿厕纸",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/0a7_ktsxYQb4aTRd6nvkKzpBthEam94L.png",
                        "num": 0
                    },
                    {
                        "id": "459",
                        "level": "3",
                        "parent_id": "454",
                        "name": "卷纸/卫生纸",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/jSpxEX-jc4ew2LOHeGbtXmWdkCigk866.png",
                        "num": 0
                    },
                    {
                        "id": "460",
                        "level": "3",
                        "parent_id": "454",
                        "name": "厨房用纸",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/zUCsLANaj4cdKu_SGS1fc9SRbmeSwYkf.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "495",
                "level": "2",
                "parent_id": "488",
                "name": "家庭/个人清洁工具",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "497",
                        "level": "3",
                        "parent_id": "495",
                        "name": "家务/地板清洁用具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/4Sh8kwQwnshTAvYA0zCZG2RCx0RpvgdI.png",
                        "num": 0
                    },
                    {
                        "id": "498",
                        "level": "3",
                        "parent_id": "495",
                        "name": "卫浴/置物用具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/N4a0s14Tb5jAxluyRYqVI891nPLf7L0X.png",
                        "num": 0
                    },
                    {
                        "id": "496",
                        "level": "3",
                        "parent_id": "495",
                        "name": "个人洗护清洁用具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/g1uendqoT0dTogW9Vn7sOpOSUXarrsG5.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "450",
                "level": "2",
                "parent_id": "488",
                "name": "室内除臭/芳香用品",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "461",
                        "level": "3",
                        "parent_id": "450",
                        "name": "空气清新/净化/芳香剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/swEMhQNO5C-quej7g8RRR3ySKSUlx3rr.png",
                        "num": 0
                    },
                    {
                        "id": "462",
                        "level": "3",
                        "parent_id": "450",
                        "name": "甲醛清除剂/甲醛测试",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/psXTD6Adocv8Gk_Hc6BoxWAoEsH8cz5L.png",
                        "num": 0
                    },
                    {
                        "id": "463",
                        "level": "3",
                        "parent_id": "450",
                        "name": "冰箱除味剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/dPboe2iJQZyBGYbZkt1Qkubs3ZoBtfh4.png",
                        "num": 0
                    },
                    {
                        "id": "464",
                        "level": "3",
                        "parent_id": "450",
                        "name": "除臭喷雾/剂/除臭用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/XyhbpZwp4RYHOfFEmAe9op_F58sRoBdu.png",
                        "num": 0
                    },
                    {
                        "id": "465",
                        "level": "3",
                        "parent_id": "450",
                        "name": "干燥剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/cx33QchaIe7I01AaiNsTHbIbraQ1JU79.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "451",
                "level": "2",
                "parent_id": "488",
                "name": "家庭清洁剂",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/kZf0gIlsqhqCkDo50OCYVR5XC2xkqInn.png",
                "num": 0,
                "childer": [
                    {
                        "id": "466",
                        "level": "3",
                        "parent_id": "451",
                        "name": "厨房清洁",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/4s-S0TRHq4GxPfA7AgN-_FFALhANscrY.png",
                        "num": 0
                    },
                    {
                        "id": "467",
                        "level": "3",
                        "parent_id": "451",
                        "name": "居家清洁",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/KfgIFiKHyoBKhpkE8Mdj40-MVaQGGeIF.png",
                        "num": 0
                    },
                    {
                        "id": "468",
                        "level": "3",
                        "parent_id": "451",
                        "name": "消毒用品",
                        "status": "1",
                        "image": "",
                        "num": 0
                    }
                ]
            },
            {
                "id": "452",
                "level": "2",
                "parent_id": "488",
                "name": "衣物清洁/护理剂",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "469",
                        "level": "3",
                        "parent_id": "452",
                        "name": "洗衣液/洗衣凝珠/留香珠",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/8vnIWnCIwi4WHyqHYq06Dk971dSOatJ8.png",
                        "num": 0
                    },
                    {
                        "id": "470",
                        "level": "3",
                        "parent_id": "452",
                        "name": "洗衣粉/爆炸盐/漂白剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/TJRXEW80UAH0SS7TnJdM2azyJBJFphSY.png",
                        "num": 0
                    },
                    {
                        "id": "471",
                        "level": "3",
                        "parent_id": "452",
                        "name": "洗衣皂/内衣皂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/UvZOTXH2mMzr-R61KHSW7iZaulkyjAN6.png",
                        "num": 0
                    },
                    {
                        "id": "472",
                        "level": "3",
                        "parent_id": "452",
                        "name": "柔顺剂/吸色片/织物喷雾",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/qCEd1X-2U_IYWGoRoMLPhjIlq9KLIVK2.png",
                        "num": 0
                    },
                    {
                        "id": "473",
                        "level": "3",
                        "parent_id": "452",
                        "name": "其他衣物洗涤剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/VykOkKZG-w5gtN48xjxSQOZSaeIzZ5Tk.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "453",
                "level": "2",
                "parent_id": "488",
                "name": "驱虫用品",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "474",
                        "level": "3",
                        "parent_id": "453",
                        "name": "家用杀虫剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/lklS1MsrqWMVyBDi8MxcPq2_0ivd7z5h.png",
                        "num": 0
                    },
                    {
                        "id": "475",
                        "level": "3",
                        "parent_id": "453",
                        "name": "蚊香液/电蚊香",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/BiYLTs-ZxOOpriMNv_EuFC15rGO5BuBz.png",
                        "num": 0
                    },
                    {
                        "id": "476",
                        "level": "3",
                        "parent_id": "453",
                        "name": "除螨喷雾",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/lsgkcoKgptzmeAKOuZUpsNKaYHW2W3Q9.png",
                        "num": 0
                    },
                    {
                        "id": "477",
                        "level": "3",
                        "parent_id": "453",
                        "name": "蚊香片/蚊香盘",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/HucAJ7dgflhh6v63P9HpsKJi71rFpmbZ.png",
                        "num": 0
                    },
                    {
                        "id": "478",
                        "level": "3",
                        "parent_id": "453",
                        "name": "驱蚊液/驱蚊凝露/驱蚊喷雾",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/TgdE4joj8JY-hdCxDdPwbuMRFBOOt-iH.png",
                        "num": 0
                    },
                    {
                        "id": "479",
                        "level": "3",
                        "parent_id": "453",
                        "name": "其他驱虫用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/-_lLf0gx48gLcDfH3BEnvTMjuT6DZK66.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "455",
                "level": "2",
                "parent_id": "488",
                "name": "其他家清用品",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "480",
                        "level": "3",
                        "parent_id": "455",
                        "name": "家私清洁/皮具护理剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/JQu5EmV-E-HEBa9KI5LTo73iI8B005yX.png",
                        "num": 0
                    },
                    {
                        "id": "482",
                        "level": "3",
                        "parent_id": "455",
                        "name": "清洁慕斯",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/sZHWq_35bPOSN0Kd3jvrItlvE_cYF5vX.png",
                        "num": 0
                    },
                    {
                        "id": "483",
                        "level": "3",
                        "parent_id": "455",
                        "name": "香薰用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/jjnedJKl21dA7QivaVJ8t6eatvkE75cB.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "16",
        "level": "1",
        "parent_id": "0",
        "name": "美妆护肤",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/ulmJ2dSHDTNYIQF68QCYyYhU8MmLB2Wb.png",
        "num": 0,
        "childer": [
            {
                "id": "72",
                "level": "2",
                "parent_id": "16",
                "name": "美容护肤",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/AJL6MVCls18oBLXNqfV2Fo24qTvb-qsf.png",
                "num": 0,
                "childer": [
                    {
                        "id": "84",
                        "level": "3",
                        "parent_id": "72",
                        "name": "卸妆",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/wwoNQQEeF8rCRDl7Uj5Hsker7za82snX.png",
                        "num": 0
                    },
                    {
                        "id": "83",
                        "level": "3",
                        "parent_id": "72",
                        "name": "乳液面霜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/YXPec7ksq7PZPtL3E0PxLJcrTRSo7W9Y.png",
                        "num": 0
                    },
                    {
                        "id": "82",
                        "level": "3",
                        "parent_id": "72",
                        "name": "防晒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/av1sri3Iz6IKyl2_wnXZIbeBQlTrOZxn.png",
                        "num": 0
                    },
                    {
                        "id": "81",
                        "level": "3",
                        "parent_id": "72",
                        "name": "护肤套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/a995WF_Qvw4Nf49xmeCRtFU1lYpFDBGL.png",
                        "num": 0
                    },
                    {
                        "id": "79",
                        "level": "3",
                        "parent_id": "72",
                        "name": "洁面",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/PfEhyVDAApL7W7ElCFAOB5ZTI4ny7lEH.png",
                        "num": 0
                    },
                    {
                        "id": "78",
                        "level": "3",
                        "parent_id": "72",
                        "name": "眼部护理",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/oSoRhvqDWr8pD5bgBrfX_5yPp9Q-LkA3.png",
                        "num": 0
                    },
                    {
                        "id": "77",
                        "level": "3",
                        "parent_id": "72",
                        "name": "精华",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/JRU7FU1Z0tF5fffRWJn0XEwYysJa8fZX.png",
                        "num": 0
                    },
                    {
                        "id": "76",
                        "level": "3",
                        "parent_id": "72",
                        "name": "面膜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/_0JTrBTNK3sdfg1LLjsynRIQ5kYrUTQQ.png",
                        "num": 0
                    },
                    {
                        "id": "319",
                        "level": "3",
                        "parent_id": "72",
                        "name": "精油/纯露",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/EWS_DHNkqJZvAX0tpOR1mOnyr_cOZ_KP.png",
                        "num": 0
                    },
                    {
                        "id": "320",
                        "level": "3",
                        "parent_id": "72",
                        "name": "基础护肤",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/uSSDyHq8N4Z2vaSA9__Lc2z9S1Rh_MDg.png",
                        "num": 0
                    },
                    {
                        "id": "206",
                        "level": "3",
                        "parent_id": "72",
                        "name": "爽肤水",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/zNigSZmvH9yuHKUNPUrUZqW4Ii1bWWET.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "73",
                "level": "2",
                "parent_id": "16",
                "name": "彩妆香水",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/WS70RvvhURaSjPozPpFFgvoXrckLQ5qv.png",
                "num": 0,
                "childer": [
                    {
                        "id": "90",
                        "level": "3",
                        "parent_id": "73",
                        "name": "睫毛膏",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/YcpW_qqXQkl2EIm-V-8QQ5T399FhbMHp.png",
                        "num": 0
                    },
                    {
                        "id": "89",
                        "level": "3",
                        "parent_id": "73",
                        "name": "眉笔",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/eTlDwjg-_tf0IRmgmnpChvTdrVcdiurE.png",
                        "num": 0
                    },
                    {
                        "id": "88",
                        "level": "3",
                        "parent_id": "73",
                        "name": "隔离",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Z8JpJ4EY6pMhm0BIh5_CK1tkEg51-DfS.png",
                        "num": 0
                    },
                    {
                        "id": "86",
                        "level": "3",
                        "parent_id": "73",
                        "name": "口红",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/T0Ojt5B4qoqRv2BIu8kt08vYL3tDi24k.png",
                        "num": 0
                    },
                    {
                        "id": "85",
                        "level": "3",
                        "parent_id": "73",
                        "name": "香水香膏",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Fq81SNGzgsl1WfSHYZCqGnh0sGxqqgHi.png",
                        "num": 0
                    },
                    {
                        "id": "270",
                        "level": "3",
                        "parent_id": "73",
                        "name": "气垫",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5GHjGeAoDnEzzBk8lskTOkQXHRSbp_sG.png",
                        "num": 0
                    },
                    {
                        "id": "271",
                        "level": "3",
                        "parent_id": "73",
                        "name": "粉底遮瑕",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/dCjWOaRqiIVNZHqd6xrJXwKnVkDyuhVA.png",
                        "num": 0
                    },
                    {
                        "id": "272",
                        "level": "3",
                        "parent_id": "73",
                        "name": "散粉",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yVskkkgx-mAnjWMzFzdUhZ8XdiOPjRfz.png",
                        "num": 0
                    },
                    {
                        "id": "273",
                        "level": "3",
                        "parent_id": "73",
                        "name": "彩妆套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ArWq1j6CV8XpKkpycLsbv0PZ3YDBvsd4.png",
                        "num": 0
                    },
                    {
                        "id": "274",
                        "level": "3",
                        "parent_id": "73",
                        "name": "眼影眼线",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/neXW2h-9Dd6GEsb6Y4H8XZOzMS2IptwJ.png",
                        "num": 0
                    },
                    {
                        "id": "275",
                        "level": "3",
                        "parent_id": "73",
                        "name": "腮红/修容",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/-v4Z30byjzWuTjMF1qCIoPDmMAGxNIzY.png",
                        "num": 0
                    },
                    {
                        "id": "321",
                        "level": "3",
                        "parent_id": "73",
                        "name": "美妆工具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/pKL9CLz5aknuAlMRbHYYF5QRcIpxOWYQ.png",
                        "num": 0
                    },
                    {
                        "id": "322",
                        "level": "3",
                        "parent_id": "73",
                        "name": "其他",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/MAbrIlCEJ-XjzHTtCEHWud9z78VRBx00.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "74",
                "level": "2",
                "parent_id": "16",
                "name": "美容美体",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/cc8wNI0XTgJnkHBocfVCje87IPTpAY26.png",
                "num": 0,
                "childer": [
                    {
                        "id": "276",
                        "level": "3",
                        "parent_id": "74",
                        "name": "美容美体医疗器械",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/9CB2UsswG5FSbcSvo35pPK50yCJT7cKs.png",
                        "num": 0
                    },
                    {
                        "id": "277",
                        "level": "3",
                        "parent_id": "74",
                        "name": "美容/个护仪器",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/GgoCAAIQNCDJN2igM5LrjW9flQ6lAORU.png",
                        "num": 0
                    },
                    {
                        "id": "278",
                        "level": "3",
                        "parent_id": "74",
                        "name": "女士脱毛",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Cp44V-Vl_LhMneetBhukEGVA1jSC6776.png",
                        "num": 0
                    },
                    {
                        "id": "279",
                        "level": "3",
                        "parent_id": "74",
                        "name": "洁面仪",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/o0cJmTnZ3VNlVw9hDQrfyssj67A1NlX8.png",
                        "num": 0
                    },
                    {
                        "id": "280",
                        "level": "3",
                        "parent_id": "74",
                        "name": "蒸脸器",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/I8GJbH_ZCK3mcwcrt85hudJgiXOM-6kS.png",
                        "num": 0
                    },
                    {
                        "id": "334",
                        "level": "3",
                        "parent_id": "74",
                        "name": "美容服务",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/wMEfAnEETUTnbqcPUWCm7AzVJM9QBg67.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "487",
        "level": "1",
        "parent_id": "0",
        "name": "个人护理",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20220928/okSWOoXmmNlbsJ9DAHnEQO-6V5PMZ15d.png",
        "num": 0,
        "childer": [
            {
                "id": "395",
                "level": "2",
                "parent_id": "487",
                "name": "洗发护发",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "404",
                        "level": "3",
                        "parent_id": "395",
                        "name": "洗发水",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/sMUihQn6n2zKbTTeepwI_Zn7IKX3-Du5.png",
                        "num": 0
                    },
                    {
                        "id": "405",
                        "level": "3",
                        "parent_id": "395",
                        "name": "洗护套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/DTdARi_SwHIxao1rdabtHo7XF2dq-Fxx.png",
                        "num": 0
                    },
                    {
                        "id": "406",
                        "level": "3",
                        "parent_id": "395",
                        "name": "护发素/发膜/精油",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/F78Yx5VdFCCTDHynN98pa_qjkbRVWMd2.png",
                        "num": 0
                    },
                    {
                        "id": "407",
                        "level": "3",
                        "parent_id": "395",
                        "name": "其它护发",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/WLgU8sLLjaScMTBQpQZ-xWAhA9LLIHZ3.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "396",
                "level": "2",
                "parent_id": "487",
                "name": "口腔护理",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "408",
                        "level": "3",
                        "parent_id": "396",
                        "name": "牙膏",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/jQ4ZvAgR--TrYjuwF1Lh29DOavkSauBx.png",
                        "num": 0
                    },
                    {
                        "id": "409",
                        "level": "3",
                        "parent_id": "396",
                        "name": "牙刷/口腔清洁工具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/EN-UtBiuwDw_1cZfmBn3Zs0LzTo03Nd9.png",
                        "num": 0
                    },
                    {
                        "id": "410",
                        "level": "3",
                        "parent_id": "396",
                        "name": "牙齿美白/护理剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/x-KS40P-H0KQE6pJ-1GKixfknoJ8KNOI.png",
                        "num": 0
                    },
                    {
                        "id": "411",
                        "level": "3",
                        "parent_id": "396",
                        "name": "口腔护理套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/vrB08Hl-Tc97ZiJQ-YZp31uYHx3-IYzN.png",
                        "num": 0
                    },
                    {
                        "id": "412",
                        "level": "3",
                        "parent_id": "396",
                        "name": "漱口水/口腔喷剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/xBFzhDICxh7pcmeqs5b8Gtq17_ocytva.png",
                        "num": 0
                    },
                    {
                        "id": "413",
                        "level": "3",
                        "parent_id": "396",
                        "name": "牙线/牙签/牙线棒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5oE4XJLTDGEYybr0MhPu3Ac_OHIICF7_.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "397",
                "level": "2",
                "parent_id": "487",
                "name": "身体清洁",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "414",
                        "level": "3",
                        "parent_id": "397",
                        "name": "沐浴露/油",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/K5nUuBwbzT3J8K10YTkbvbTHOMi95zPk.png",
                        "num": 0
                    },
                    {
                        "id": "415",
                        "level": "3",
                        "parent_id": "397",
                        "name": "香皂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/mfNB4-oN4pQ0QPB_aK8JUCEdLK1vSc10.png",
                        "num": 0
                    },
                    {
                        "id": "416",
                        "level": "3",
                        "parent_id": "397",
                        "name": "身体磨砂膏/去角质膏",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/bM9j_wI_lGPOdS1EYJVg7yFHUPU4SHEe.png",
                        "num": 0
                    },
                    {
                        "id": "417",
                        "level": "3",
                        "parent_id": "397",
                        "name": "洗手液",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/TMkBtUY-6DX-WMFXCp1Rj6sRCDzNsKKD.png",
                        "num": 0
                    },
                    {
                        "id": "418",
                        "level": "3",
                        "parent_id": "397",
                        "name": "其他",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Pwa_FgwnEbVjA8xeIg_UBQN75pWYlRNi.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "398",
                "level": "2",
                "parent_id": "487",
                "name": "身体护理",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "419",
                        "level": "3",
                        "parent_id": "398",
                        "name": "身体乳/霜/精油/膏",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yMDwP7pclAxdIdmu_O6zFJ_GvQXawxK-.png",
                        "num": 0
                    },
                    {
                        "id": "420",
                        "level": "3",
                        "parent_id": "398",
                        "name": "身体护理套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/TpK_4m0UJ2ZZS1k4MPoLdXNnuvO-tPY2.png",
                        "num": 0
                    },
                    {
                        "id": "422",
                        "level": "3",
                        "parent_id": "398",
                        "name": "颈霜/颈膜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/aUjkUCOPdWlYaozr_IbNcMPmRQvZQFiR.png",
                        "num": 0
                    },
                    {
                        "id": "423",
                        "level": "3",
                        "parent_id": "398",
                        "name": "止汗露/消臭剂",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/BTnxtVUJbSJmAhxVf7KW3Aq9CUQPEan5.png",
                        "num": 0
                    },
                    {
                        "id": "424",
                        "level": "3",
                        "parent_id": "398",
                        "name": "肩颈/膝肘贴",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/GrK6pzD6y33vjyqBIyE0Z565UcZ3QIaK.png",
                        "num": 0
                    },
                    {
                        "id": "425",
                        "level": "3",
                        "parent_id": "398",
                        "name": "足部护理",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/cYuwpzZn8w8eCazgOiPFz12Z003arM0s.png",
                        "num": 0
                    },
                    {
                        "id": "426",
                        "level": "3",
                        "parent_id": "398",
                        "name": "胸部护理",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/4r8lTIM4Cx-e9KmXvQ366icRYzNphE3S.png",
                        "num": 0
                    },
                    {
                        "id": "427",
                        "level": "3",
                        "parent_id": "398",
                        "name": "脱毛膏/凝胶",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/buylVlJnGGLx5gvJAMw3RwqNbcuQBS24.png",
                        "num": 0
                    },
                    {
                        "id": "428",
                        "level": "3",
                        "parent_id": "398",
                        "name": "清凉油/防暑/醒神药油",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/bTLPjMF44zTrSm_spoqWew15dSiUpxTS.png",
                        "num": 0
                    },
                    {
                        "id": "429",
                        "level": "3",
                        "parent_id": "398",
                        "name": "身体皂/身体喷雾",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2qnYitkZeKckVX-8U4RfP3vNUXOauVab.png",
                        "num": 0
                    },
                    {
                        "id": "430",
                        "level": "3",
                        "parent_id": "398",
                        "name": "其他",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/xp_rUpKuspaKTVkw0oYCQF6JdhY3gue3.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "399",
                "level": "2",
                "parent_id": "487",
                "name": "卫生巾/私处护理",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "431",
                        "level": "3",
                        "parent_id": "399",
                        "name": "卫生巾/护垫/卫生裤",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/bxmgEjpukm1P17LiFIZq2UuwrAqcgqWb.png",
                        "num": 0
                    },
                    {
                        "id": "432",
                        "level": "3",
                        "parent_id": "399",
                        "name": "卫生棉条",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/CyJ6jH_7a792PZkrA4OpcyF_41SLR-eE.png",
                        "num": 0
                    },
                    {
                        "id": "433",
                        "level": "3",
                        "parent_id": "399",
                        "name": "私处洗液/湿巾",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/P1LwosFPEHWesATcYnbhbuU3MEJV_EZk.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "400",
                "level": "2",
                "parent_id": "487",
                "name": "染发烫发",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "434",
                        "level": "3",
                        "parent_id": "400",
                        "name": "染发产品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/DhR2ity2sYChVNfBES4-S4bWFeM3_eki.png",
                        "num": 0
                    },
                    {
                        "id": "435",
                        "level": "3",
                        "parent_id": "400",
                        "name": "摩丝/啫喱/头发造型",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/zVllUZEbEmx1b6tCMisqagBnvWiCcMA6.png",
                        "num": 0
                    },
                    {
                        "id": "436",
                        "level": "3",
                        "parent_id": "400",
                        "name": "烫发产品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/m3V_kkX1T9611UNJHvfzrPu_4d7EPTMw.png",
                        "num": 0
                    },
                    {
                        "id": "437",
                        "level": "3",
                        "parent_id": "400",
                        "name": "美发工具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/TyGCvKoV3N8fKkAsSfpNe43_lJO7eYmh.png",
                        "num": 0
                    },
                    {
                        "id": "493",
                        "level": "3",
                        "parent_id": "400",
                        "name": "其它",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/uKNB3fG2XKk6Oshjc-oM0iP0kafMAm0d.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "401",
                "level": "2",
                "parent_id": "487",
                "name": "手部护理",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "438",
                        "level": "3",
                        "parent_id": "401",
                        "name": "护手霜/手膜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/sM_lI0UyPBZH-nfoH7uUWT4CiVTXGhiR.png",
                        "num": 0
                    },
                    {
                        "id": "439",
                        "level": "3",
                        "parent_id": "401",
                        "name": "手部保养套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/4c1WdYopWUAWu1bDbmZ6LZOhjBdVO9QD.png",
                        "num": 0
                    },
                    {
                        "id": "440",
                        "level": "3",
                        "parent_id": "401",
                        "name": "手部修复/清洁",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/UXDxYrqiu7jcf3Ea6KEQG9fQwVRZZfDk.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "402",
                "level": "2",
                "parent_id": "487",
                "name": "个护工具",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "441",
                        "level": "3",
                        "parent_id": "402",
                        "name": "剃须刀/指甲钳",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/xxqwdiRNR5NnyJvdsswfbXft2Ejqx1mf.png",
                        "num": 0
                    },
                    {
                        "id": "442",
                        "level": "3",
                        "parent_id": "402",
                        "name": "塑身贴/塑形带",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/YCg9_FX6xESIZ1A0iF7GJKcisRxv_4Rl.png",
                        "num": 0
                    },
                    {
                        "id": "443",
                        "level": "3",
                        "parent_id": "402",
                        "name": "发圈/发网/发带",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/EqW_p7F_HzfRWWxuW0iK8vazLd1XBUkC.png",
                        "num": 0
                    },
                    {
                        "id": "444",
                        "level": "3",
                        "parent_id": "402",
                        "name": "梳子/卷发棒/化妆镜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/i6fG9LxDMna7Yvf_YqmAoA1jZihqaN_h.png",
                        "num": 0
                    },
                    {
                        "id": "445",
                        "level": "3",
                        "parent_id": "402",
                        "name": "蒸汽眼罩",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/7bmEbe7QW5FQLBymIQTYNYIpSc0vOsIr.png",
                        "num": 0
                    },
                    {
                        "id": "446",
                        "level": "3",
                        "parent_id": "402",
                        "name": "按摩用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/R_1y9ZUfQWXkxoh4481_I3EPaBmlnGwu.png",
                        "num": 0
                    },
                    {
                        "id": "447",
                        "level": "3",
                        "parent_id": "402",
                        "name": "其他个护工具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/uq3zJkOdQazM2ibaJBOgq6LP1Q0fHw7Q.png",
                        "num": 0
                    },
                    {
                        "id": "494",
                        "level": "3",
                        "parent_id": "402",
                        "name": "个护电器",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Yr3_vegMf4HbQ8H1ICch0Zqz-_82TCAG.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "403",
                "level": "2",
                "parent_id": "487",
                "name": "其他个护用品",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "448",
                        "level": "3",
                        "parent_id": "403",
                        "name": "假发及配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/6HynqqcAWB05-SIwMyaDsuy19j8eilar.png",
                        "num": 0
                    },
                    {
                        "id": "449",
                        "level": "3",
                        "parent_id": "403",
                        "name": "身体精油芳疗",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/kCRUtDF4ZoGBtqCDW-jSehx3ijZAnkay.png",
                        "num": 0
                    },
                    {
                        "id": "489",
                        "level": "3",
                        "parent_id": "403",
                        "name": "男士个护",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/9eOe1o2E0RiHrUKy-gbwx32feYnlCXo6.png",
                        "num": 0
                    },
                    {
                        "id": "490",
                        "level": "3",
                        "parent_id": "403",
                        "name": "中老年护理用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yfKMdWiXgoHjJpRCApUS-uhUxGkrcRuI.png",
                        "num": 0
                    },
                    {
                        "id": "491",
                        "level": "3",
                        "parent_id": "403",
                        "name": "隐形眼镜/护理液",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/w7iL-YXMpF9_OoBmCElpkc8HWlSg6FJI.png",
                        "num": 0
                    },
                    {
                        "id": "492",
                        "level": "3",
                        "parent_id": "403",
                        "name": "个护服务",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/WzM4u0fB4rh2tcc9oxj-ujNWgdpC6JmJ.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "375",
        "level": "1",
        "parent_id": "0",
        "name": "家纺布艺",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20220915/x_adQHyHW3apnt2weFMX_4Py1kjdQlw-.jpg",
        "num": 0,
        "childer": [
            {
                "id": "376",
                "level": "2",
                "parent_id": "375",
                "name": "家居布艺",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "378",
                        "level": "3",
                        "parent_id": "376",
                        "name": "窗帘窗纱",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Dx_HyaLO6cPor1j5IpCSoriOWzxIAG6t.png",
                        "num": 0
                    },
                    {
                        "id": "379",
                        "level": "3",
                        "parent_id": "376",
                        "name": "地毯/地垫",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/6JLP5P7ThKmtWJS66MabYKcZK9njUSMp.png",
                        "num": 0
                    },
                    {
                        "id": "380",
                        "level": "3",
                        "parent_id": "376",
                        "name": "抱枕/靠垫/座垫",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Mw3Y44MQ5zZsd1EXfv1FR7I0fI1RYkLa.png",
                        "num": 0
                    },
                    {
                        "id": "381",
                        "level": "3",
                        "parent_id": "376",
                        "name": "桌布/罩件/布料",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/7N0n88FxpSsEtSXSnw18s4xjeye3MqyU.png",
                        "num": 0
                    },
                    {
                        "id": "382",
                        "level": "3",
                        "parent_id": "376",
                        "name": "毛巾/浴巾/浴袍",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/B7TpqY5ZEBAjc-64IZW86n_O5lnkxhXX.png",
                        "num": 0
                    },
                    {
                        "id": "383",
                        "level": "3",
                        "parent_id": "376",
                        "name": "其他配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/j-vPsNZ_xmIArk_yCZRwXWrUJLbF3Fsf.png",
                        "num": 0
                    },
                    {
                        "id": "389",
                        "level": "3",
                        "parent_id": "376",
                        "name": "拖鞋/居家鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/fdIYK4ZIz76084K3f0asPKd63xt2f7Q_.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "377",
                "level": "2",
                "parent_id": "375",
                "name": "床上用品",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "384",
                        "level": "3",
                        "parent_id": "377",
                        "name": "被子/枕头",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/qn30GDi1SvmiooB0C-aRrKqC8Gw0n7j0.png",
                        "num": 0
                    },
                    {
                        "id": "385",
                        "level": "3",
                        "parent_id": "377",
                        "name": "四件套/床笠床套",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/aUoRlZiGZxq7ust4wITaN9Ge9gKlCusE.png",
                        "num": 0
                    },
                    {
                        "id": "386",
                        "level": "3",
                        "parent_id": "377",
                        "name": "电热毯/毛毯",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Bij5632c5Icl-WQC04ywLUVeDz7IEWG6.png",
                        "num": 0
                    },
                    {
                        "id": "387",
                        "level": "3",
                        "parent_id": "377",
                        "name": "凉席/蚊帐",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/-6iI9SH4Zifj8rS_mOBQmcPL36bmkZoz.png",
                        "num": 0
                    },
                    {
                        "id": "388",
                        "level": "3",
                        "parent_id": "377",
                        "name": "儿童床品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5I83-Tk2t0qkuJYU0SVkAaNeGaaukYVg.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "17",
        "level": "1",
        "parent_id": "0",
        "name": "母婴童玩",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/cRgc82GaZTvwHH2ruBfAl-JJgpxNjEfT.png",
        "num": 0,
        "childer": [
            {
                "id": "107",
                "level": "2",
                "parent_id": "17",
                "name": "童装",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/PTIJKKcZkxuOMtLKnPI4uRm7k2FoXibP.png",
                "num": 0,
                "childer": [
                    {
                        "id": "188",
                        "level": "3",
                        "parent_id": "107",
                        "name": "儿童上衣",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/0_M9Xs3d7kmA70cawArHm53oHjoybfXE.png",
                        "num": 0
                    },
                    {
                        "id": "189",
                        "level": "3",
                        "parent_id": "107",
                        "name": "亲子装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/1uButCFq7TFPbjVyMg7GlZvTBwAJSJ1W.png",
                        "num": 0
                    },
                    {
                        "id": "190",
                        "level": "3",
                        "parent_id": "107",
                        "name": "儿童下装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/7vGTKmkZqB0sB4EsCcwbMyr_LPcxCktZ.png",
                        "num": 0
                    },
                    {
                        "id": "191",
                        "level": "3",
                        "parent_id": "107",
                        "name": "内衣裤袜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/oLnjArmkOCtlKw1t6QIq2S3YxclRGIS-.png",
                        "num": 0
                    },
                    {
                        "id": "192",
                        "level": "3",
                        "parent_id": "107",
                        "name": "儿童套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yoKNv0r0GSmil-zpzThqPxi9uqMSdhrR.png",
                        "num": 0
                    },
                    {
                        "id": "193",
                        "level": "3",
                        "parent_id": "107",
                        "name": "儿童配饰",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/d01_ffBA2JZwcFheUw9UhB6X4oz1qMb-.png",
                        "num": 0
                    },
                    {
                        "id": "194",
                        "level": "3",
                        "parent_id": "107",
                        "name": "儿童外套",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/xdi9j6DNAn5W6-rtRC-LDiQNHv34kHw1.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "105",
                "level": "2",
                "parent_id": "17",
                "name": "孕妇用品",
                "status": "1",
                "image": "https://static2.redu.com/uploads/images/2022/0523/20220523163402_55080.png",
                "num": 0,
                "childer": [
                    {
                        "id": "330",
                        "level": "3",
                        "parent_id": "105",
                        "name": "孕妇装/妈咪装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/sxejBSTRG1Lha-DFlJtCyEb1_uUiLb3T.png",
                        "num": 0
                    },
                    {
                        "id": "182",
                        "level": "3",
                        "parent_id": "105",
                        "name": "孕妇个护",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Le_6H5GEY2GZQyOEom9aK-dxJVyZNewC.png",
                        "num": 0
                    },
                    {
                        "id": "184",
                        "level": "3",
                        "parent_id": "105",
                        "name": "孕妇家居",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Vog3b-v1fFIB7JkpWfOzCT8KZNQ1nm-r.png",
                        "num": 0
                    },
                    {
                        "id": "185",
                        "level": "3",
                        "parent_id": "105",
                        "name": "孕妇营养",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Wlw0ob-wun1tEAUzVSx3QoiZ_uroJds3.png",
                        "num": 0
                    },
                    {
                        "id": "186",
                        "level": "3",
                        "parent_id": "105",
                        "name": "孕妇配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/lcNYlu5WLAhvskvIuzM03d7lpdNP2aN-.png",
                        "num": 0
                    },
                    {
                        "id": "187",
                        "level": "3",
                        "parent_id": "105",
                        "name": "骨盆带/塑身衣",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/aEEfjnpLtawgU3bKyi_6AwGT4KNzdtpi.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "104",
                "level": "2",
                "parent_id": "17",
                "name": "婴童用品",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/4SRX3oP1AjZWe5o-VDB98L4H2c1A8VBX.png",
                "num": 0,
                "childer": [
                    {
                        "id": "331",
                        "level": "3",
                        "parent_id": "104",
                        "name": "日用起居",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/qZ-lSFs39IkF847tt6c7VNpaXSVnDeS8.png",
                        "num": 0
                    },
                    {
                        "id": "332",
                        "level": "3",
                        "parent_id": "104",
                        "name": "安全防护",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/kzUZuQTAv1A5ao4_UxP5HaG6XeX7Gpa5.png",
                        "num": 0
                    },
                    {
                        "id": "175",
                        "level": "3",
                        "parent_id": "104",
                        "name": "婴童尿裤",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/O303x_UFRVUNkf-mqSXsJUuMcznjUSuN.png",
                        "num": 0
                    },
                    {
                        "id": "176",
                        "level": "3",
                        "parent_id": "104",
                        "name": "儿童餐具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/gfoaZqb-NbwZygJmQPysn1HupWBfFMfu.png",
                        "num": 0
                    },
                    {
                        "id": "177",
                        "level": "3",
                        "parent_id": "104",
                        "name": "童车座椅",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/f_xL2dmK-qzhPXghwvhpqQFPOH_35Tl1.png",
                        "num": 0
                    },
                    {
                        "id": "178",
                        "level": "3",
                        "parent_id": "104",
                        "name": "婴幼个护",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/M4Xi1zwT4ahzptTA5v6JpuuOvpCZt6Zg.png",
                        "num": 0
                    },
                    {
                        "id": "179",
                        "level": "3",
                        "parent_id": "104",
                        "name": "喂养用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/RKbM9v3stCl2dLLJXUFzw-a70jCx6ihD.png",
                        "num": 0
                    },
                    {
                        "id": "181",
                        "level": "3",
                        "parent_id": "104",
                        "name": "宝宝护肤",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/DD1Z2pi06PA8Dyp_3hiwk-pXf9TgdvQe.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "103",
                "level": "2",
                "parent_id": "17",
                "name": "童鞋",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/1agUM21CuTwxCgC8kozyBdQelz66MA-Q.png",
                "num": 0,
                "childer": [
                    {
                        "id": "329",
                        "level": "3",
                        "parent_id": "103",
                        "name": "童鞋配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/mQS55lj-YVbFVaAyKHofqqrVtK4CQaoC.png",
                        "num": 0
                    },
                    {
                        "id": "199",
                        "level": "3",
                        "parent_id": "103",
                        "name": "儿童拖鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/NiRkPii8uAWFwNUbIVZtWHmcrIhrM-mZ.png",
                        "num": 0
                    },
                    {
                        "id": "200",
                        "level": "3",
                        "parent_id": "103",
                        "name": "儿童凉鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/3uNzuzPYGDDBOcsG_gXJE_elQHVGA2oO.png",
                        "num": 0
                    },
                    {
                        "id": "201",
                        "level": "3",
                        "parent_id": "103",
                        "name": "休闲鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/v_URiHUrbn-22lnLh8Q_lgFcoqtNhhQX.png",
                        "num": 0
                    },
                    {
                        "id": "202",
                        "level": "3",
                        "parent_id": "103",
                        "name": "学步鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/bOCgxhDyOvZfs9kOekwYgcVMaJb0vPLH.png",
                        "num": 0
                    },
                    {
                        "id": "203",
                        "level": "3",
                        "parent_id": "103",
                        "name": "运动鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/cZbnfbVzQLOLt3OhlcU3NL1TrAR5JkEF.png",
                        "num": 0
                    },
                    {
                        "id": "204",
                        "level": "3",
                        "parent_id": "103",
                        "name": "儿童皮鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/gNb2gXpu13D_75S-P7HJxvyZdQfepdId.png",
                        "num": 0
                    },
                    {
                        "id": "508",
                        "level": "3",
                        "parent_id": "103",
                        "name": "其他童鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Bqct9jhx-aMm-OCIjycasJViykM4qSQH.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "361",
                "level": "2",
                "parent_id": "17",
                "name": "宠物/宠物食品及用品",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "363",
                        "level": "3",
                        "parent_id": "361",
                        "name": "宠物食品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/hZNw5ChJWYFwfrnHIYAnrgm6lMr9pf1x.png",
                        "num": 0
                    },
                    {
                        "id": "364",
                        "level": "3",
                        "parent_id": "361",
                        "name": "宠物用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/FeqYTXQWpLEmNZQB8VPhteVPg91D0Nha.png",
                        "num": 0
                    },
                    {
                        "id": "486",
                        "level": "3",
                        "parent_id": "361",
                        "name": "宠物",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Lu49N4uv7atvBLOLXdQO1fKOR5pNIYDU.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "365",
                "level": "2",
                "parent_id": "17",
                "name": "图书杂志",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "366",
                        "level": "3",
                        "parent_id": "365",
                        "name": "文学小说类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5JYacJhjnrUMnB5kNyoW3RDVPOLRPq0T.png",
                        "num": 0
                    },
                    {
                        "id": "367",
                        "level": "3",
                        "parent_id": "365",
                        "name": "儿童读物类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/WAg0MHXgb9MEERt_iWb9EqJBT7b0uEzV.png",
                        "num": 0
                    },
                    {
                        "id": "368",
                        "level": "3",
                        "parent_id": "365",
                        "name": "教材辅导类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ke34PAEFVLBxgr8Jpf3-dHwm7mAnBKda.png",
                        "num": 0
                    },
                    {
                        "id": "369",
                        "level": "3",
                        "parent_id": "365",
                        "name": "休闲杂志类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/dicCSMuYec0m-oIjMnfT6IYmB3pkE76b.png",
                        "num": 0
                    },
                    {
                        "id": "370",
                        "level": "3",
                        "parent_id": "365",
                        "name": "社会科学类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yZJPCRJAKWDhr_1d1D3dNDoR1J6X6hgM.png",
                        "num": 0
                    },
                    {
                        "id": "390",
                        "level": "3",
                        "parent_id": "365",
                        "name": "进口书籍",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/dgaEYt4ykNqH0qyW4dcv7bZ9F-BF7gc4.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "372",
                "level": "2",
                "parent_id": "17",
                "name": "学习用品",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "373",
                        "level": "3",
                        "parent_id": "372",
                        "name": "文具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/3BMrFAgiESsAo_I7fA9jj-C59eEf-b95.png",
                        "num": 0
                    },
                    {
                        "id": "500",
                        "level": "3",
                        "parent_id": "372",
                        "name": "画具/画材/书法用品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/sNWHQL0ZoHXJTjL_V8dzG453Gxf4d2mW.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "130",
                "level": "2",
                "parent_id": "17",
                "name": "奶粉辅食",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/Xv1yV-_6qFnUPkSlYNBMadSE5ryCRL2z.png",
                "num": 0,
                "childer": [
                    {
                        "id": "195",
                        "level": "3",
                        "parent_id": "130",
                        "name": "婴儿调味",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/crQ1Opu36e3Qx3-t5Lr3voFEt8x1cb17.png",
                        "num": 0
                    },
                    {
                        "id": "196",
                        "level": "3",
                        "parent_id": "130",
                        "name": "辅食零食",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/unkx-bEbmoS-ZspildVZXpi-7_jlguAf.png",
                        "num": 0
                    },
                    {
                        "id": "197",
                        "level": "3",
                        "parent_id": "130",
                        "name": "儿童营养",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/xaTerWF3eptU4cC7mHG9sriAL71983yF.png",
                        "num": 0
                    },
                    {
                        "id": "198",
                        "level": "3",
                        "parent_id": "130",
                        "name": "奶粉",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2axHzvC6cQ_sKvrGRw6UUDK1WZ3WsV49.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "207",
                "level": "2",
                "parent_id": "17",
                "name": "婴幼玩具",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/Au0pH8OtQAelYDIMLIgW691038OqOrcg.png",
                "num": 0,
                "childer": [
                    {
                        "id": "323",
                        "level": "3",
                        "parent_id": "207",
                        "name": "益智玩具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/_6IbN559GEBsTeT-uCz1CstijOPJo-4y.png",
                        "num": 0
                    },
                    {
                        "id": "324",
                        "level": "3",
                        "parent_id": "207",
                        "name": "手工DIY",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Fni30vcjxAmIaNPnZgYrZTVq9c3cqyrI.png",
                        "num": 0
                    },
                    {
                        "id": "325",
                        "level": "3",
                        "parent_id": "207",
                        "name": "积木/模型",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/MQKU3l5_VPpKbwr2XOOZbl-AtFP9WFC3.png",
                        "num": 0
                    },
                    {
                        "id": "326",
                        "level": "3",
                        "parent_id": "207",
                        "name": "电动玩具",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/q92YNcEXgIdmouXn3c9Qn9a1oQzVzWk9.png",
                        "num": 0
                    },
                    {
                        "id": "327",
                        "level": "3",
                        "parent_id": "207",
                        "name": "童车运动",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/DQA7nSWgoGNJWv1aC9zSRGEuhr4y-mlL.png",
                        "num": 0
                    },
                    {
                        "id": "328",
                        "level": "3",
                        "parent_id": "207",
                        "name": "宝宝纪念品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/hs54fF14VRrIKBkXFAigHfP1EqKQGuOp.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "15",
        "level": "1",
        "parent_id": "0",
        "name": "服饰鞋包",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/ryG9XGu9ET19xgdNI5huoIZubcqa7AXz.png",
        "num": 0,
        "childer": [
            {
                "id": "58",
                "level": "2",
                "parent_id": "15",
                "name": "内衣裤/袜子",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/MQUKP9O6WxrtdFAEG4Y0lfyRYH_Z9XAk.png",
                "num": 0,
                "childer": [
                    {
                        "id": "71",
                        "level": "3",
                        "parent_id": "58",
                        "name": "家居服",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/hYE-6FlpMAYRwTea30iKsoAJ-Mh-BfID.png",
                        "num": 0
                    },
                    {
                        "id": "70",
                        "level": "3",
                        "parent_id": "58",
                        "name": "内裤",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/okOfLpgP_Mt99KHHsckSBEMM3_bIupng.png",
                        "num": 0
                    },
                    {
                        "id": "69",
                        "level": "3",
                        "parent_id": "58",
                        "name": "文胸塑身",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/U0V0h6QCEEykiGuMEm6JuevwOf0adSyZ.png",
                        "num": 0
                    },
                    {
                        "id": "360",
                        "level": "3",
                        "parent_id": "58",
                        "name": "保暖套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/58e7-0LfPUO22Z2QGSyy7q4llG8Qkc7y.png",
                        "num": 0
                    },
                    {
                        "id": "160",
                        "level": "3",
                        "parent_id": "58",
                        "name": "袜子",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/SbHLoXxwoGh0RzbpfdUGM_-yjHlIWJ9m.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "57",
                "level": "2",
                "parent_id": "15",
                "name": "女装",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/I7rax5tUX7nrH0Gd5DSoMSW8tlrfDSog.png",
                "num": 0,
                "childer": [
                    {
                        "id": "300",
                        "level": "3",
                        "parent_id": "57",
                        "name": "连衣裙",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5qyXmcLECZLuiak_5wTpabRD_XreiwXM.png",
                        "num": 0
                    },
                    {
                        "id": "301",
                        "level": "3",
                        "parent_id": "57",
                        "name": "女士上衣",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/25RMbG9GDwTOu-54e6KBosaNhvm4GWKG.png",
                        "num": 0
                    },
                    {
                        "id": "302",
                        "level": "3",
                        "parent_id": "57",
                        "name": "裤子",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Gldvkx3CVtXvuypzdC1fp4z2J_O18pFe.png",
                        "num": 0
                    },
                    {
                        "id": "303",
                        "level": "3",
                        "parent_id": "57",
                        "name": "套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/KVr2NVv_YF9UY8xGlYk7NeVBiKnQjOae.png",
                        "num": 0
                    },
                    {
                        "id": "304",
                        "level": "3",
                        "parent_id": "57",
                        "name": "衬衫",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/zAiJBOCFh_wOXjVY-9itHNS26kC07OzS.png",
                        "num": 0
                    },
                    {
                        "id": "305",
                        "level": "3",
                        "parent_id": "57",
                        "name": "大码女装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/RH5n-X5L_CySnYJkDeSTxSzyH7jPH_PG.png",
                        "num": 0
                    },
                    {
                        "id": "356",
                        "level": "3",
                        "parent_id": "57",
                        "name": "民族服装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/oohMSEm77SJq8rPCMngEPlCr8AcM2lVU.png",
                        "num": 0
                    },
                    {
                        "id": "357",
                        "level": "3",
                        "parent_id": "57",
                        "name": "短裙",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/InldbAcghpW_TZmzn0GpDrwnPdLe9QQI.png",
                        "num": 0
                    },
                    {
                        "id": "358",
                        "level": "3",
                        "parent_id": "57",
                        "name": "内搭/背心",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/KRvu-MJhVRB2-YtjnQRWzFLbMKZEm8PN.png",
                        "num": 0
                    },
                    {
                        "id": "359",
                        "level": "3",
                        "parent_id": "57",
                        "name": "礼服",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/gVLwfkkZ78oz-vuy-HAmAgy255X2uo21.png",
                        "num": 0
                    },
                    {
                        "id": "503",
                        "level": "3",
                        "parent_id": "57",
                        "name": "毛衣/针织衫",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/7hVZTSa07OAr1vWECQf3Zfnn_rUF__No.png",
                        "num": 0
                    },
                    {
                        "id": "505",
                        "level": "3",
                        "parent_id": "57",
                        "name": "女式外套",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/XJyj5qUqGGyMw7d6wDuQiO52a0azNlpj.png",
                        "num": 0
                    },
                    {
                        "id": "506",
                        "level": "3",
                        "parent_id": "57",
                        "name": "女式短裤",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/T81AHi9ONTQ4jQ58IAkzMSC_mhVgM14s.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "56",
                "level": "2",
                "parent_id": "15",
                "name": "男装",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/nzOG94dQ6215BBcdbS3Pmkf0GJOOySY4.png",
                "num": 0,
                "childer": [
                    {
                        "id": "65",
                        "level": "3",
                        "parent_id": "56",
                        "name": "男士内搭",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/FgXCKeFEZFrhVofFJ9Hv2R4YFp6ZVR7L.png",
                        "num": 0
                    },
                    {
                        "id": "64",
                        "level": "3",
                        "parent_id": "56",
                        "name": "男士外套",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/pH8u4xmyCKMypRFGcIJNFvB20tq1TdfZ.png",
                        "num": 0
                    },
                    {
                        "id": "63",
                        "level": "3",
                        "parent_id": "56",
                        "name": "男士裤装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/JRqv2P4KZzPpDoeyzyJb6M-EoKkkwWKk.png",
                        "num": 0
                    },
                    {
                        "id": "352",
                        "level": "3",
                        "parent_id": "56",
                        "name": "套装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/FDje5ZR8cjQS2-UF46-95z2QtvWSs00K.png",
                        "num": 0
                    },
                    {
                        "id": "353",
                        "level": "3",
                        "parent_id": "56",
                        "name": "男士上衣",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2fsuxZA0A7M584kAV-1tD3JTLJvgILA9.png",
                        "num": 0
                    },
                    {
                        "id": "354",
                        "level": "3",
                        "parent_id": "56",
                        "name": "短裤",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/GEPCqS3eKzYJ05_ZcM8ARwPEKBC7BSJg.png",
                        "num": 0
                    },
                    {
                        "id": "355",
                        "level": "3",
                        "parent_id": "56",
                        "name": "民族服装",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2oq_1JcrZK-7SGHzoC4eU--93yglA2ei.png",
                        "num": 0
                    },
                    {
                        "id": "504",
                        "level": "3",
                        "parent_id": "56",
                        "name": "毛衣/针织衫",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/_bYeykPXX1gKEhtqnIKyfvNSCxK8Oh3n.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "55",
                "level": "2",
                "parent_id": "15",
                "name": "珠宝配饰",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/WAIKVsIM5e4LTxu0AvXJhPJZEK5fNyuB.png",
                "num": 0,
                "childer": [
                    {
                        "id": "62",
                        "level": "3",
                        "parent_id": "55",
                        "name": "皮带/手表",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/KyKOVAfJBL2ZMueab4aIHYKboX-gJLEH.png",
                        "num": 0
                    },
                    {
                        "id": "336",
                        "level": "3",
                        "parent_id": "55",
                        "name": "黄金珠宝",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/LQ4dN_f4Ormm6FMKVPql2C7iRh32vWAw.png",
                        "num": 0
                    },
                    {
                        "id": "337",
                        "level": "3",
                        "parent_id": "55",
                        "name": "时尚配饰",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/70SHLjQiaH-rgUaEV4VvZ3DuBSgdVOTs.png",
                        "num": 0
                    },
                    {
                        "id": "141",
                        "level": "3",
                        "parent_id": "55",
                        "name": "打火机/眼镜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/La2vAhDSOBIv2FqmDHlpB9VSIjG-WQf4.png",
                        "num": 0
                    },
                    {
                        "id": "142",
                        "level": "3",
                        "parent_id": "55",
                        "name": "帽子/围巾/手套",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/GS5W41ZuFceAfv6xluaLOk_mxJ5BgCaN.png",
                        "num": 0
                    },
                    {
                        "id": "143",
                        "level": "3",
                        "parent_id": "55",
                        "name": "其他配饰",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Kw-i0C4-FFTedYyxJKVF_01W7ntX479d.png",
                        "num": 0
                    },
                    {
                        "id": "484",
                        "level": "3",
                        "parent_id": "55",
                        "name": "木作/书画/钱币/收藏",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yu0DghXN9_xcISRaCEI6Ubk1bkKEdlSc.png",
                        "num": 0
                    },
                    {
                        "id": "485",
                        "level": "3",
                        "parent_id": "55",
                        "name": "珠宝/文玩/收藏品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ypvLmROL1T_HYGQFhLrp_NQ_krrdxY8M.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "338",
                "level": "2",
                "parent_id": "15",
                "name": "时尚箱包",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/HSj7fqFpPL-jPhnqi8XR_cI3gyBN2n5B.png",
                "num": 0,
                "childer": [
                    {
                        "id": "339",
                        "level": "3",
                        "parent_id": "338",
                        "name": "功能箱包",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/QpJUfj1duvlH7GkT64nfRFgU-o4HB_MC.png",
                        "num": 0
                    },
                    {
                        "id": "340",
                        "level": "3",
                        "parent_id": "338",
                        "name": "女士包袋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/u6P34umLMuofKTLv723VDQmeNTn_x7pD.png",
                        "num": 0
                    },
                    {
                        "id": "341",
                        "level": "3",
                        "parent_id": "338",
                        "name": "男士包袋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/EMp3qIdlK1Xb8oR1VPcwK_Gbn8WvAm6c.png",
                        "num": 0
                    },
                    {
                        "id": "342",
                        "level": "3",
                        "parent_id": "338",
                        "name": "卡套",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ifo3Q92IRaPu1hmCTS88Ju6mfIVMQT-h.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "343",
                "level": "2",
                "parent_id": "15",
                "name": "运动鞋靴",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/rbILdkVZS_S_zQR6xVxQ2SHIYTn7olEi.png",
                "num": 0,
                "childer": [
                    {
                        "id": "344",
                        "level": "3",
                        "parent_id": "343",
                        "name": "运动鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/-fLZfWE0wsOXColwN1KzXW5ER9XSj_xR.png",
                        "num": 0
                    },
                    {
                        "id": "345",
                        "level": "3",
                        "parent_id": "343",
                        "name": "休闲鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/CGVW9H9hPg39JUT32FwprXfiJgwVncKE.png",
                        "num": 0
                    },
                    {
                        "id": "346",
                        "level": "3",
                        "parent_id": "343",
                        "name": "皮鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/tvFrLnVoadvWNr-I_BfuHLDWAo31hrvo.png",
                        "num": 0
                    },
                    {
                        "id": "347",
                        "level": "3",
                        "parent_id": "343",
                        "name": "靴子",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ELLZvwc78jmppgyPrS_ufKwYYX46ZY5o.png",
                        "num": 0
                    },
                    {
                        "id": "348",
                        "level": "3",
                        "parent_id": "343",
                        "name": "高跟鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/mBM0Mz_md5g8HQBVWbs7WgxD9Ls13GLa.png",
                        "num": 0
                    },
                    {
                        "id": "349",
                        "level": "3",
                        "parent_id": "343",
                        "name": "拖鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2w7p4TnksHwdFONo-dogI3D5VWs3GuoT.png",
                        "num": 0
                    },
                    {
                        "id": "350",
                        "level": "3",
                        "parent_id": "343",
                        "name": "凉鞋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/j4CNUU8BvoxVgIGgLFYcTUOczxHL0nZu.png",
                        "num": 0
                    },
                    {
                        "id": "351",
                        "level": "3",
                        "parent_id": "343",
                        "name": "其他",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/T5qnWRAJPSGJY2CTHxbJboQ4MBRG41V4.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "13",
        "level": "1",
        "parent_id": "0",
        "name": "生鲜水果",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/_fxqgmnimtGykAtogeOXpWSsDS_wgZH8.png",
        "num": 0,
        "childer": [
            {
                "id": "49",
                "level": "2",
                "parent_id": "13",
                "name": "海鲜水产",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/PF_NY1LGReuLJRV5H1bY34QpKKnHxNHK.png",
                "num": 0,
                "childer": [
                    {
                        "id": "248",
                        "level": "3",
                        "parent_id": "49",
                        "name": "蟹类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/olJS2FgfJfvCyfyQ_lkL2z8Q9nwSM8gU.png",
                        "num": 0
                    },
                    {
                        "id": "249",
                        "level": "3",
                        "parent_id": "49",
                        "name": "海鲜制品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/DI5uMvhvNSKxz1911U84y7hW35FRBZH_.png",
                        "num": 0
                    },
                    {
                        "id": "250",
                        "level": "3",
                        "parent_id": "49",
                        "name": "海参",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/aeSmbhyRXqOvfVCEJmsa9auzegf6SYsC.png",
                        "num": 0
                    },
                    {
                        "id": "251",
                        "level": "3",
                        "parent_id": "49",
                        "name": "贝类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/E2TYfyFpW5eEaDmlSRrA02PPV_NlB4Dr.png",
                        "num": 0
                    },
                    {
                        "id": "252",
                        "level": "3",
                        "parent_id": "49",
                        "name": "虾类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/k1J1qFb26QrfMhVqZtJKLPx_e6UUF8cX.png",
                        "num": 0
                    },
                    {
                        "id": "253",
                        "level": "3",
                        "parent_id": "49",
                        "name": "鱼类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/iRv8GOGejlN6PgaWzT4OpaXdzF4aT_E9.png",
                        "num": 0
                    },
                    {
                        "id": "254",
                        "level": "3",
                        "parent_id": "49",
                        "name": "其他海产",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/_ewU2YypWCTjDx8MN7LalFKI5lLpcZLT.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "48",
                "level": "2",
                "parent_id": "13",
                "name": "新鲜水果",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/sEUrfGFV7plCKZ6k5GvlgHZip0HhuhyX.png",
                "num": 0,
                "childer": [
                    {
                        "id": "262",
                        "level": "3",
                        "parent_id": "48",
                        "name": "柑橘橙柚",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Wbw6cR30BXBSaE0nGltAQfqyUkNI5StI.png",
                        "num": 0
                    },
                    {
                        "id": "263",
                        "level": "3",
                        "parent_id": "48",
                        "name": "葡提莓果",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/exfbH0K1gjOyDVMfvZtSKQVxEGEVsWFm.png",
                        "num": 0
                    },
                    {
                        "id": "264",
                        "level": "3",
                        "parent_id": "48",
                        "name": "苹果蕉梨",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/zaRPIadmuOI7BgibFNuILA2QzwUTWr9A.png",
                        "num": 0
                    },
                    {
                        "id": "265",
                        "level": "3",
                        "parent_id": "48",
                        "name": "桃李杏枣",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/V9viGw6HSNsmt_v977bHA4Gp0y84lueB.png",
                        "num": 0
                    },
                    {
                        "id": "266",
                        "level": "3",
                        "parent_id": "48",
                        "name": "热带水果",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/wTlFdR2RRThrlhOruRRQ67Gwx9eIcZXn.png",
                        "num": 0
                    },
                    {
                        "id": "267",
                        "level": "3",
                        "parent_id": "48",
                        "name": "西瓜/蜜瓜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/PJANF7ISJvM-LSWrE2OjJYGnI1an44XS.png",
                        "num": 0
                    },
                    {
                        "id": "268",
                        "level": "3",
                        "parent_id": "48",
                        "name": "猕猴桃",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/IhTvqlM8IrFjCk76A2YFqsGXI0npV_Vb.png",
                        "num": 0
                    },
                    {
                        "id": "269",
                        "level": "3",
                        "parent_id": "48",
                        "name": "水果礼盒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/k9EAokD5neBz6vevIkKxc4nHUCEfqEFb.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "47",
                "level": "2",
                "parent_id": "13",
                "name": "蔬菜蛋品",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/7s4lyMZXlev2MsjCt0mytS_koOPJHk6i.png",
                "num": 0,
                "childer": [
                    {
                        "id": "256",
                        "level": "3",
                        "parent_id": "47",
                        "name": "根茎类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/FsMzb-aAorz24YdD7Fn6z8zg9ouP57xF.png",
                        "num": 0
                    },
                    {
                        "id": "257",
                        "level": "3",
                        "parent_id": "47",
                        "name": "鲜菌菇",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/PHhfrgbGQWQXaR23AdMGAza-mpb6Lpv8.png",
                        "num": 0
                    },
                    {
                        "id": "258",
                        "level": "3",
                        "parent_id": "47",
                        "name": "茄果瓜类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/iq-VxgjBlGBrzPzAcXx95u5dyQ37Konp.png",
                        "num": 0
                    },
                    {
                        "id": "259",
                        "level": "3",
                        "parent_id": "47",
                        "name": "快手菜",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/JdoC40kKNVtq4JsuC6BXl4GKm7tKYqVx.png",
                        "num": 0
                    },
                    {
                        "id": "260",
                        "level": "3",
                        "parent_id": "47",
                        "name": "葱姜蒜椒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Od6WKcNj7tEI18WxKg-KWBHA8yM3CKZj.png",
                        "num": 0
                    },
                    {
                        "id": "261",
                        "level": "3",
                        "parent_id": "47",
                        "name": "蛋品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/5C0G1EdqpBvGLM1vzHDbx8V38_YaTvnM.png",
                        "num": 0
                    },
                    {
                        "id": "307",
                        "level": "3",
                        "parent_id": "47",
                        "name": "豆制品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/1FwcyvMHx1Lgu5FnO6ZHm3y7xiMoxqlH.png",
                        "num": 0
                    },
                    {
                        "id": "507",
                        "level": "3",
                        "parent_id": "47",
                        "name": "其他农产品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/QWGKdCvsOw_I6n0uN58YCqrhXYIyxkij.png",
                        "num": 0
                    },
                    {
                        "id": "255",
                        "level": "3",
                        "parent_id": "47",
                        "name": "叶菜类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/t1OqT0NwNKHbdNxJ-lgkkxFcbCd0DCQJ.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "46",
                "level": "2",
                "parent_id": "13",
                "name": "精选肉类",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/53hlppRZnLV9jEBo4eyOfZutYEbHNMc6.png",
                "num": 0,
                "childer": [
                    {
                        "id": "306",
                        "level": "3",
                        "parent_id": "46",
                        "name": "其他肉类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2GeI3pem2jj7uiHz4Z9DMp_LrDv5Bvmj.png",
                        "num": 0
                    },
                    {
                        "id": "239",
                        "level": "3",
                        "parent_id": "46",
                        "name": "猪肉类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/gTMhYdCbe-DpNUKIUtnrorRoJVpInnbs.png",
                        "num": 0
                    },
                    {
                        "id": "240",
                        "level": "3",
                        "parent_id": "46",
                        "name": "牛羊肉类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/HhTGjzUHJMWNIG0sMUgndJMupa_WXHMj.png",
                        "num": 0
                    },
                    {
                        "id": "241",
                        "level": "3",
                        "parent_id": "46",
                        "name": "鸡鸭肉类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/4UM0wR7E4qUUdeoeERsaiFf6EsQmk4D_.png",
                        "num": 0
                    },
                    {
                        "id": "242",
                        "level": "3",
                        "parent_id": "46",
                        "name": "肉制品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/tbzNNr5v1Tvz0oxOg8ACtjr2PN_aKnye.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "391",
                "level": "2",
                "parent_id": "13",
                "name": "生鲜卡券",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "392",
                        "level": "3",
                        "parent_id": "391",
                        "name": "水果/肉蛋/蔬菜卡券",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/4hqSZRs91Q5n-leBd8bW0zI6upV8toIH.png",
                        "num": 0
                    },
                    {
                        "id": "393",
                        "level": "3",
                        "parent_id": "391",
                        "name": "海鲜/水产品/制品卡券",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/57OUI77B9WkvVV3aItS59Oviu6PHwVWh.png",
                        "num": 0
                    },
                    {
                        "id": "394",
                        "level": "3",
                        "parent_id": "391",
                        "name": "烘焙蛋糕卡券",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/NpY5rcKR4a9kalSaOElowoOFj1-MPsfd.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "161",
                "level": "2",
                "parent_id": "13",
                "name": "冷饮冻食",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/RQdaTWSFtz1KZi-gpMdwP_nuiAzKTKL5.png",
                "num": 0,
                "childer": [
                    {
                        "id": "243",
                        "level": "3",
                        "parent_id": "161",
                        "name": "包点/面点",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/W5XxBuukwQzFhmOs-wdvCReyNqMwjcPe.png",
                        "num": 0
                    },
                    {
                        "id": "244",
                        "level": "3",
                        "parent_id": "161",
                        "name": "煎饼类",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/_wOs_wXbTaKE7zDyjKGmmpQhI8Gpokb8.png",
                        "num": 0
                    },
                    {
                        "id": "245",
                        "level": "3",
                        "parent_id": "161",
                        "name": "水饺馄饨",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/gGU7lXEiber_Xt8slqsgSBYGB5JOggVX.png",
                        "num": 0
                    },
                    {
                        "id": "246",
                        "level": "3",
                        "parent_id": "161",
                        "name": "冰淇淋",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ncz2-GxnTLQoQGcv_EdStaIs_3cu1Jzd.png",
                        "num": 0
                    },
                    {
                        "id": "247",
                        "level": "3",
                        "parent_id": "161",
                        "name": "冻半成品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/DqTc8mDFIxTdLseeDq70QcHEsCS2FiL0.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "18",
        "level": "1",
        "parent_id": "0",
        "name": "3C数码家电",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/hcBGgSeOuLlxXWW-9YxHDPFZqgVpZ53Z.png",
        "num": 0,
        "childer": [
            {
                "id": "109",
                "level": "2",
                "parent_id": "18",
                "name": "大家电",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221109/Mph1TlwZcmIl9WqTK9oHk6027qodEEYb.png",
                "num": 0,
                "childer": [
                    {
                        "id": "122",
                        "level": "3",
                        "parent_id": "109",
                        "name": "热水器",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/74n0APtHbuhDdXlkgPvyE99e_IkSC8Dx.png",
                        "num": 0
                    },
                    {
                        "id": "121",
                        "level": "3",
                        "parent_id": "109",
                        "name": "厨房大家电",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/oXuorE3b-RY9QvNbBYlFWSc-3qN5Gbv9.png",
                        "num": 0
                    },
                    {
                        "id": "120",
                        "level": "3",
                        "parent_id": "109",
                        "name": "洗衣机",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/dlh9A2KumK1XRXOMvmg30VqxdQF32tmJ.png",
                        "num": 0
                    },
                    {
                        "id": "119",
                        "level": "3",
                        "parent_id": "109",
                        "name": "冰箱",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/NPNe_IEgHFhmlqv-Sv-FAyEL_LudzH5e.png",
                        "num": 0
                    },
                    {
                        "id": "118",
                        "level": "3",
                        "parent_id": "109",
                        "name": "空调",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/h6efkEP9d8vWq5iFsbqwxbP7MWCUZxJF.png",
                        "num": 0
                    },
                    {
                        "id": "117",
                        "level": "3",
                        "parent_id": "109",
                        "name": "电视",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/8VN3IuIitg-92SdIQypn52XF8RASu9do.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "108",
                "level": "2",
                "parent_id": "18",
                "name": "手机",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/nJgTNUbK6dHNjvlrmWZcl91A6uP4kvQB.png",
                "num": 0,
                "childer": [
                    {
                        "id": "110",
                        "level": "3",
                        "parent_id": "108",
                        "name": "智能手机",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/0-zhGSeTpBtghwk5e9mN7jIpUlV0CocR.png",
                        "num": 0
                    },
                    {
                        "id": "299",
                        "level": "3",
                        "parent_id": "108",
                        "name": "老年机",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Jw7z_K4yk_1J6qznuMiz7ksSEEqjYLKv.png",
                        "num": 0
                    },
                    {
                        "id": "213",
                        "level": "3",
                        "parent_id": "108",
                        "name": "老人机",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/H9Lf2eL0YbO6dBs3f0-3J68hlKR79m64.png",
                        "num": 0
                    },
                    {
                        "id": "214",
                        "level": "3",
                        "parent_id": "108",
                        "name": "游戏手机",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/sFQUaLvFbRTlEZ7ajJFZv2mXlGNOh52g.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "137",
                "level": "2",
                "parent_id": "18",
                "name": "3C数码配件",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/n9WFYT2PC64FUEdpgbj3H2N3MVnmCccC.png",
                "num": 0,
                "childer": [
                    {
                        "id": "138",
                        "level": "3",
                        "parent_id": "137",
                        "name": "手机配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Qh74WlOZmPgePHFKlnanLww_7ybpuGzB.png",
                        "num": 0
                    },
                    {
                        "id": "139",
                        "level": "3",
                        "parent_id": "137",
                        "name": "电脑配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/bu_P3zGXyMZLVsGrr1sT3Z0oB7Y2tqlN.png",
                        "num": 0
                    },
                    {
                        "id": "140",
                        "level": "3",
                        "parent_id": "137",
                        "name": "其他数码配件",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/aK_rTO0XI4janMJUtTEDIyU9W-iX14ak.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "211",
                "level": "2",
                "parent_id": "18",
                "name": "电脑办公",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/UFsM852A0JGQV8xsCgGtg6LlgbsyQK49.png",
                "num": 0,
                "childer": [
                    {
                        "id": "217",
                        "level": "3",
                        "parent_id": "211",
                        "name": "台式机",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ReJb-ruXS05AGjxRSaXzztjJIftm4vQa.png",
                        "num": 0
                    },
                    {
                        "id": "218",
                        "level": "3",
                        "parent_id": "211",
                        "name": "笔记本",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/-xCWbeLgZ-0FMdYf3VLDV45ccdjgEXZi.png",
                        "num": 0
                    },
                    {
                        "id": "219",
                        "level": "3",
                        "parent_id": "211",
                        "name": "平板",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/N7Pm4F3HeXOmRWz9YukIa_6c3Eod5iJg.png",
                        "num": 0
                    },
                    {
                        "id": "220",
                        "level": "3",
                        "parent_id": "211",
                        "name": "键盘鼠标",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/yrPI-ymgAIeUjOxgrKmdgLawk3mK7Rrn.png",
                        "num": 0
                    },
                    {
                        "id": "221",
                        "level": "3",
                        "parent_id": "211",
                        "name": "办公设备",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/rZeqsOi7xvRVfH1W8lh782NaLXhaEqBY.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "212",
                "level": "2",
                "parent_id": "18",
                "name": "影音智能",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/kOZVA1UL7sDx_JFSeFaN7kXf7WOF_9kq.png",
                "num": 0,
                "childer": [
                    {
                        "id": "215",
                        "level": "3",
                        "parent_id": "212",
                        "name": "影音娱乐",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/E9jkKk6NgjuUeJmCoDjYtPBalJeRUIfk.png",
                        "num": 0
                    },
                    {
                        "id": "216",
                        "level": "3",
                        "parent_id": "212",
                        "name": "智能设备",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/isqV7Q2NmAi9C-Kc8_wSD1_Mlz7g4ezL.png",
                        "num": 0
                    },
                    {
                        "id": "502",
                        "level": "3",
                        "parent_id": "212",
                        "name": "摄影摄像",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/ykEVrm0PMs_4_kG5iav2hzMs3xQncNFw.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "20",
        "level": "1",
        "parent_id": "0",
        "name": "酒类/茶叶",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221109/743-xKRL1UWrYW5aJjYQYewIbCSJL70X.png",
        "num": 0,
        "childer": [
            {
                "id": "54",
                "level": "2",
                "parent_id": "20",
                "name": "酒类",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/5fwPi6hOfZVUPPVdOW9q23GyOiBTR3Ve.png",
                "num": 0,
                "childer": [
                    {
                        "id": "371",
                        "level": "3",
                        "parent_id": "54",
                        "name": "啤酒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230323/nZspz_etOjeb9hucD7296SeYZbpHHFvk.png",
                        "num": 0
                    },
                    {
                        "id": "149",
                        "level": "3",
                        "parent_id": "54",
                        "name": "白酒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/BrprLzLOR6qe_F_XnhfNgzxHZ2PAF1jG.png",
                        "num": 0
                    },
                    {
                        "id": "150",
                        "level": "3",
                        "parent_id": "54",
                        "name": "葡萄酒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/pCx62S5xdnI6DYoc0ESEMiEeBaoWPk1N.png",
                        "num": 0
                    },
                    {
                        "id": "151",
                        "level": "3",
                        "parent_id": "54",
                        "name": "洋酒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/A8P8z4f6-Exd9MrLgSNVIJ5pbrlhJ3fe.png",
                        "num": 0
                    },
                    {
                        "id": "152",
                        "level": "3",
                        "parent_id": "54",
                        "name": "黄酒/米酒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/_hN6DjGEASgLYe4sPMGa0dRG9kt33Yc-.png",
                        "num": 0
                    },
                    {
                        "id": "153",
                        "level": "3",
                        "parent_id": "54",
                        "name": "预调酒/果酒",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/HfvE3iuiygE6y6eLMqgiFT9dIu9_9tEu.png",
                        "num": 0
                    }
                ]
            },
            {
                "id": "148",
                "level": "2",
                "parent_id": "20",
                "name": "茗茶",
                "status": "1",
                "image": "https://static0.reduxingxuan.com/20221110/iDdxEBUVA96TicMkMYj28-NmoUzGemp2.png",
                "num": 0,
                "childer": [
                    {
                        "id": "333",
                        "level": "3",
                        "parent_id": "148",
                        "name": "乌龙茶",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/6b87XmJhLMciiH2FyrRXdD8YRoGgNydu.png",
                        "num": 0
                    },
                    {
                        "id": "154",
                        "level": "3",
                        "parent_id": "148",
                        "name": "绿茶",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/o95s4JalZ5ylnOpkF4pBHQfTMu2-wc8y.png",
                        "num": 0
                    },
                    {
                        "id": "155",
                        "level": "3",
                        "parent_id": "148",
                        "name": "红茶",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/f10s5EQY9GSIfIttmhrFayZNoBwVKXEm.png",
                        "num": 0
                    },
                    {
                        "id": "156",
                        "level": "3",
                        "parent_id": "148",
                        "name": "普洱",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/JVnb5uMGuIarvbTGCzhN1A5m1uq44MCO.png",
                        "num": 0
                    },
                    {
                        "id": "157",
                        "level": "3",
                        "parent_id": "148",
                        "name": "黑茶/白茶",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/2uXQpHGGsa0hYeeEb9P1JaC7-smzC2Eb.png",
                        "num": 0
                    },
                    {
                        "id": "158",
                        "level": "3",
                        "parent_id": "148",
                        "name": "花果茶",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/Jlr-hhX_-VS0hpdF5o7x11RAjkJOAUa7.png",
                        "num": 0
                    },
                    {
                        "id": "159",
                        "level": "3",
                        "parent_id": "148",
                        "name": "茶制品",
                        "status": "1",
                        "image": "https://static0.reduxingxuan.com/20230324/NN_99-iRWP2ms-TwT-Wh0TZMoBi2hedX.png",
                        "num": 0
                    }
                ]
            }
        ]
    },
    {
        "id": "125",
        "level": "1",
        "parent_id": "0",
        "name": "其他",
        "status": "1",
        "image": "https://static0.reduxingxuan.com/20221108/Co2SRKrO5MB1Qz5wNNd0DzW6bum7IrvN.png",
        "num": 0,
        "childer": [
            {
                "id": "509",
                "level": "2",
                "parent_id": "125",
                "name": "其他",
                "status": "1",
                "image": "",
                "num": 0,
                "childer": [
                    {
                        "id": "510",
                        "level": "3",
                        "parent_id": "509",
                        "name": "其他",
                        "status": "1",
                        "image": "",
                        "num": 0
                    }
                ]
            }
        ]
    }
]`
	err := json.Unmarshal([]byte(jsonData), &categories)
	if err != nil {
		log.Fatal(err)
	}
	// 将嵌套的切片转换为JSON字符串
	c := getCategory(categories)
	err = global.DB.Create(c).Error
	if err != nil {
		fmt.Println("数据入库失败: ", err)
	}
	fmt.Println("数据已成功插入数据库")
}

// 递归
func getCategory(categories []receive.Category) *[]model.Category {
	c := []model.Category{}
	for _, one := range categories {
		category := model.Category{
			ID:       one.ID,
			Level:    one.Level,
			ParentID: one.ParentID,
			Name:     one.Name,
			Image:    one.Image,
			Status:   one.Status,
			Num:      one.Num,
		}
		c = append(c, category)
		if len(one.Childer) > 0 {
			c = append(c, *getCategory(one.Childer)...)
		}
	}
	return &c
}
