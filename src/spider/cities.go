package spider


type Map map[string]string
var List = Map{
	"110000":"北京", //bj
	"310000":"上海", //sh
}
var Abbr = Map{
	"110000":"bj",
	"310000":"sh",
}


var CityList = map[string]string{
	"340800" : "安庆", // aq
	"210300" : "鞍山", // as
	"513200" : "阿坝", // ab
	"520400" : "安顺", // anshun
	"420982" : "安陆", // anlu
	"530181" : "安宁", // an
	"152900" : "阿拉善盟", // alsm
	"610900" : "安康", // ak
	"370784" : "安丘", // anqiu
	"231281" : "安达", // ad
	"652900" : "阿克苏", // aks
	"410500" : "安阳", // ay
	"659002" : "阿拉尔", // alr
	"130683" : "安国", // ag
	"110000" : "北京", // bj
	"450500" : "北海", // bh
	"130600" : "保定", // bd
	"610300" : "宝鸡", // baoji
	"150200" : "包头", // baotou
	"131081" : "霸州", // bazhou
	"511900" : "巴中", // bz
	"340300" : "蚌埠", // bf
	"520500" : "毕节", // bijie
	"530500" : "保山", // baoshan
	"341600" : "亳州", // bozhou
	"210500" : "本溪", // bx
	"220600" : "白山", // bs
	"620400" : "白银", // by
	"220800" : "白城", // bc
	"371600" : "滨州", // binzhou
	"652700" : "博尔塔拉", // brtl
	"451000" : "百色", // baise
	"231181" : "北安", // ba
	"211381" : "北票", // bp
	"130981" : "泊头", // botou
	"652701" : "博乐", // bole
	"450981" : "北流", // bl
	"469029" : "保亭", // bt
	"510100" : "成都", // cd
	"500000" : "重庆", // cq
	"430100" : "长沙", // cs
	"220100" : "长春", // cc
	"430700" : "常德", // changde
	"320400" : "常州", // changzhou
	"130900" : "沧州", // cangzhou
	"411082" : "长葛", // cg
	"140400" : "长治", // changzhi
	"520381" : "赤水", // chishui
	"150400" : "赤峰", // cf
	"341700" : "池州", // chizhou
	"532300" : "楚雄", // chuxiong
	"431000" : "郴州", // chenzhou
	"540300" : "昌都", // changdu
	"211300" : "朝阳", // chaoyang
	"220202" : "昌邑", // cy
	"445100" : "潮州", // chaozhou
	"652300" : "昌吉", // cj
	"320581" : "常熟", // changshu
	"451400" : "崇左", // chongzuo
	"130800" : "承德", // chengde
	"341100" : "滁州", // cz
	"421281" : "赤壁", // cb
	"450481" : "岑溪", // cx
	"510184" : "崇州", // chongzhou
	"330282" : "慈溪", // cixi
	"469023" : "澄迈", // cm
	"210200" : "大连", // dl
	"441900" : "东莞", // dg
	"210600" : "丹东", // dd
	"511700" : "达州", // dazhou
	"510600" : "德阳", // dy
	"140200" : "大同", // dt
	"130682" : "定州", // dingzhou
	"411381" : "邓州", // dengzhou
	"330783" : "东阳", // dongyang
	"420582" : "当阳", // dangyang
	"533100" : "德宏", // dh
	"210882" : "大石桥", // dsq
	"533400" : "迪庆", // diqing
	"220183" : "德惠", // dehui
	"361181" : "德兴", // dx
	"370500" : "东营", // dongying
	"621100" : "定西", // dingxi
	"230600" : "大庆", // dq
	"371400" : "德州", // dezhou
	"232700" : "大兴安岭", // dxal
	"320904" : "大丰", // df
	"320981" : "东台", // dongtai
	"469007" : "东方", // dongfang
	"222403" : "敦化", // dunhua
	"321181" : "丹阳", // danyang
	"211081" : "灯塔", // dengta
	"620982" : "敦煌", // dunhuang
	"632802" : "德令哈", // dlh
	"420281" : "大冶", // daye
	"410185" : "登封", // dengfeng
	"510181" : "都江堰", // djy
	"210681" : "东港", // donggang
	"469003" : "儋州", // dz
	"469021" : "定安", // da
	"532900" : "大理", // dali
	"420700" : "鄂州", // ez
	"150600" : "鄂尔多斯", // eeds
	"422800" : "恩施", // es
	"440785" : "恩平", // ep
	"440600" : "佛山", // fs
	"350100" : "福州", // fz
	"341200" : "阜阳", // fy
	"210400" : "抚顺", // fushun
	"360981" : "丰城", // fengcheng
	"361000" : "抚州", // fuzhou
	"370983" : "肥城", // fc
	"450600" : "防城港", // fcg
	"652302" : "阜康", // fk
	"141182" : "汾阳", // fenyang
	"522702" : "福泉", // fq
	"350981" : "福安", // fa
	"210682" : "凤城", // fch
	"350982" : "福鼎", // fd
	"350181" : "福清", // fuqing
	"440100" : "广州", // gz
	"520100" : "贵阳", // gy
	"450300" : "桂林", // gl
	"360700" : "赣州", // ganzhou
	"321084" : "高邮", // gaoyou
	"511600" : "广安", // ga
	"140181" : "古交", // gujiao
	"513300" : "甘孜", // ganzi
	"421381" : "广水", // gs
	"210881" : "盖州", // gaizhou
	"220381" : "公主岭", // gzl
	"440981" : "高州", // gaozhou
	"623000" : "甘南", // gn
	"370785" : "高密", // gm
	"640400" : "固原", // guyuan
	"450800" : "贵港", // gg
	"450881" : "桂平", // gp
	"632801" : "格尔木", // grm
	"532501" : "个旧", // gejiu
	"360681" : "贵溪", // gx
	"130684" : "高碑店", // gbd
	"510800" : "广元", // guangyuan
	"330100" : "杭州", // hz
	"340100" : "合肥", // hf
	"441300" : "惠州", // hui
	"460100" : "海口", // hk
	"230100" : "哈尔滨", // hrb
	"420200" : "黄石", // huangshi
	"320800" : "淮安", // ha
	"150100" : "呼和浩特", // hhht
	"330500" : "湖州", // huzhou
	"610700" : "汉中", // hanzhong
	"130400" : "邯郸", // hd
	"430400" : "衡阳", // hy
	"131100" : "衡水", // hs
	"371700" : "菏泽", // heze
	"340400" : "淮南", // huainan
	"340600" : "淮北", // huaibei
	"150700" : "呼伦贝尔", // hlbr
	"341000" : "黄山", // huangshan
	"421083" : "洪湖", // honghu
	"341522" : "霍邱", // hq
	"210381" : "海城", // haicheng
	"431200" : "怀化", // hh
	"431281" : "洪江", // hongjiang
	"211400" : "葫芦岛", // hld
	"220282" : "桦甸", // huadian
	"440982" : "化州", // huazhou
	"441600" : "河源", // heyuan
	"230400" : "鹤岗", // hegang
	"630200" : "海东", // haidong
	"632800" : "海西", // haixi
	"231083" : "海林", // hl
	"231100" : "黑河", // heihe
	"650500" : "哈密", // hm
	"451100" : "贺州", // hezhou
	"451200" : "河池", // hc
	"410600" : "鹤壁", // hb
	"653200" : "和田", // ht
	"410602" : "鹤山", // hsh
	"421100" : "黄冈", // hg
	"222404" : "珲春", // hunchun
	"130983" : "黄骅", // huanghua
	"130984" : "河间", // hj
	"610581" : "韩城", // hancheng
	"141081" : "侯马", // houma
	"511681" : "华蓥", // huaying
	"451381" : "合山", // heshan
	"410782" : "辉县", // huixian
	"141082" : "霍州", // huozhou
	"520382" : "仁怀", // hr
	"330481" : "海宁", // haining
	"370687" : "海阳", // haiyang
	"320684" : "海门", // haimen
	"532500" : "红河哈尼族彝族自治州", // honghezhou
	"370101" : "济南", // jn
	"330400" : "嘉兴", // jx
	"440700" : "江门", // jiangmen
	"220200" : "吉林", // jl
	"330700" : "金华", // jh
	"360800" : "吉安", // jian
	"360400" : "九江", // jiujiang
	"370800" : "济宁", // jining
	"420800" : "荆门", // jm
	"410800" : "焦作", // jiaozuo
	"510185" : "简阳", // jianyang
	"321282" : "靖江", // jingjiang
	"419001" : "济源", // jiyuan
	"140500" : "晋城", // jc
	"421000" : "荆州", // jingzhou
	"532801" : "景洪", // jinghong
	"210700" : "锦州", // jinzhou
	"360200" : "景德镇", // jdz
	"620200" : "嘉峪关", // jyg
	"620300" : "金昌", // jinchang
	"620900" : "酒泉", // jq
	"230300" : "鸡西", // jixi
	"445200" : "揭阳", // jieyang
	"230800" : "佳木斯", // jms
	"320413" : "金坛", // jt
	"140700" : "晋中", // jz
	"433101" : "吉首", // jishou
	"350783" : "建瓯", // jo
	"220281" : "蛟河", // jiaohe
	"330182" : "建德", // jd
	"131103" : "冀州", // jizhou
	"321183" : "句容", // jr
	"330421" : "嘉善", // jiashan
	"370281" : "胶州", // jiaozhou
	"360124" : "进贤", // jinxian
	"510781" : "江油", // jiangyou
	"320281" : "江阴", // jy
	"410200" : "开封", // kf
	"530100" : "昆明", // km
	"320583" : "昆山", // ks
	"440783" : "开平", // kp
	"522601" : "凯里", // kaili
	"532502" : "开远", // ky
	"211282" : "开原", // kaiyuan
	"650200" : "克拉玛依", // klmy
	"652801" : "库尔勒", // krl
	"653000" : "克孜勒苏", // kzls
	"653100" : "喀什", // kashi
	"654003" : "奎屯", // kt
	"131000" : "廊坊", // lf
	"410300" : "洛阳", // luoyang
	"450200" : "柳州", // liuzhou
	"620100" : "兰州", // lz
	"513400" : "凉山彝族自治州", // liangshan
	"371300" : "临沂", // linyi
	"320700" : "连云港", // lyg
	"371500" : "聊城", // lc
	"510500" : "泸州", // luzhou
	"411100" : "漯河", // luohe
	"330781" : "兰溪", // lanxi
	"331082" : "临海", // lh
	"331100" : "丽水", // lishui
	"331181" : "龙泉", // lq
	"520200" : "六盘水", // lps
	"140821" : "临猗", // lin
	"141000" : "临汾", // linfen
	"141100" : "吕梁", // lvliang
	"341500" : "六安", // la
	"530700" : "丽江", // lj
	"530900" : "临沧", // lincang
	"430281" : "醴陵", // liling
	"430682" : "临湘", // linxiang
	"350681" : "龙海", // longhai
	"211000" : "辽阳", // liaoyang
	"540100" : "拉萨", // lasa
	"360281" : "乐平", // lp
	"431300" : "娄底", // loudi
	"440881" : "廉江", // lianjiang
	"220400" : "辽源", // liaoyuan
	"440882" : "雷州", // leizhou
	"370682" : "莱阳", // laiyang
	"621200" : "陇南", // ln
	"622900" : "临夏", // lx
	"441882" : "连州", // lianzhou
	"371200" : "莱芜", // lw
	"371481" : "乐陵", // ll
	"371581" : "临清", // linqing
	"320481" : "溧阳", // liyang
	"410581" : "林州", // linzhou
	"451300" : "来宾", // lb
	"511100" : "乐山", // leshan
	"350800" : "龙岩", // ly
	"222405" : "龙井", // longjing
	"220681" : "临江", // linjiang
	"211382" : "凌源", // lingyuan
	"140481" : "潞城", // lucheng
	"422802" : "利川", // lichuan
	"431381" : "冷水江", // lsj
	"431382" : "涟源", // lianyuan
	"511381" : "阆中", // langzhong
	"440281" : "乐昌", // lechang
	"441581" : "陆丰", // lufeng
	"445381" : "罗定", // luoding
	"350122" : "连江", // lianj
	"370285" : "莱西", // laixi
	"420682" : "老河口", // lhk
	"430181" : "浏阳", // liuyang
	"370683" : "莱州", // laizhou
	"469027" : "乐东", // ld
	"469028" : "陵水", // ls
	"469024" : "临高", // lg
	"540400" : "林芝", // linzhi
	"510700" : "绵阳", // mianyang
	"340500" : "马鞍山", // mas
	"511400" : "眉山", // ms
	"341182" : "明光", // mingguang
	"440900" : "茂名", // mm
	"441400" : "梅州", // meizhou
	"231000" : "牡丹江", // mdj
	"410883" : "孟州", // mengzhou
	"230382" : "密山", // mishan
	"150781" : "满洲里", // mzl
	"421181" : "麻城", // mc
	"510683" : "绵竹", // mianzhu
	"320100" : "南京", // nj
	"330200" : "宁波", // nb
	"360100" : "南昌", // nc
	"511300" : "南充", // nanchong
	"320600" : "南通", // nt
	"450100" : "南宁", // nn
	"411300" : "南阳", // ny
	"511000" : "内江", // neijiang
	"350583" : "南安", // na
	"350700" : "南平", // np
	"533300" : "怒江", // nujiang
	"350900" : "宁德", // nd
	"540600" : "那曲", // nq
	"360703" : "南康", // nk
	"230281" : "讷河", // nh
	"231084" : "宁安", // ningan
	"341881" : "宁国", // ng
	"130581" : "南宫", // nangong
	"410900" : "濮阳", // py
	"510400" : "攀枝花", // pzh
	"210214" : "普兰店", // pld
	"350300" : "莆田", // pt
	"211100" : "盘锦", // pj
	"360300" : "萍乡", // pingxiang
	"620800" : "平凉", // pl
	"445281" : "普宁", // pn
	"410400" : "平顶山", // pds
	"370283" : "平度", // pd
	"330482" : "平湖", // ph
	"320382" : "邳州", // pz
	"530800" : "普洱", // pr
	"370200" : "青岛", // qd
	"350500" : "泉州", // quanzhou
	"441800" : "清远", // qy
	"130283" : "迁安", // qa
	"140121" : "清徐", // qx
	"330800" : "衢州", // quzhou
	"520181" : "清镇", // qingzhen
	"522600" : "黔东南", // qdn
	"522700" : "黔南", // qn
	"530300" : "曲靖", // qj
	"429005" : "潜江", // qianjiang
	"370781" : "青州", // qingzhou
	"230200" : "齐齐哈尔", // qqhr
	"621000" : "庆阳", // qingyang
	"370881" : "曲阜", // qf
	"230900" : "七台河", // qth
	"450700" : "钦州", // qinzhou
	"410882" : "沁阳", // qinyang
	"130300" : "秦皇岛", // qhd
	"640381" : "青铜峡", // qtx
	"320113" : "栖霞", // qixia
	"320681" : "启东", // qidong
	"510183" : "邛崃", // ql
	"469002" : "琼海", // qh
	"469030" : "琼中", // qz
	"522300" : "黔西南", // qianxinan
	"330381" : "瑞安", // ra
	"130982" : "任丘", // rq
	"533102" : "瑞丽", // rl
	"540200" : "日喀则", // rkz
	"360781" : "瑞金", // rj
	"371100" : "日照", // rz
	"410482" : "汝州", // ruzhou
	"360481" : "瑞昌", // rc
	"320682" : "如皋", // rg
	"371082" : "荣成", // rongcheng
	"310000" : "上海", // sh
	"440300" : "深圳", // sz
	"130100" : "石家庄", // sjz
	"210100" : "沈阳", // sy
	"320500" : "苏州", // su
	"460200" : "三亚", // san
	"330600" : "绍兴", // sx
	"361100" : "上饶", // sr
	"440500" : "汕头", // st
	"321300" : "宿迁", // sq
	"321322" : "沭阳", // shuyang
	"510900" : "遂宁", // sn
	"411400" : "商丘", // shangqiu
	"131082" : "三河", // sanhe
	"420300" : "十堰", // shiyan
	"421081" : "石首", // shishou
	"421087" : "松滋", // songzi
	"341300" : "宿州", // suzhou
	"350400" : "三明", // sm
	"430500" : "邵阳", // shaoyang
	"350581" : "石狮", // shishi
	"440200" : "韶关", // shaoguan
	"440606" : "顺德", // sd
	"220300" : "四平", // sp
	"611000" : "商洛", // sl
	"220700" : "松原", // songyuan
	"230183" : "尚志", // shangzhi
	"441500" : "汕尾", // sw
	"370783" : "寿光", // shouguang
	"230500" : "双鸭山", // sys
	"640200" : "石嘴山", // szs
	"231200" : "绥化", // suihua
	"330683" : "嵊州", // shengzhou
	"231081" : "绥芬河", // sfh
	"510682" : "什邡", // sf
	"350781" : "邵武", // shaowu
	"430382" : "韶山", // ss
	"131182" : "深州", // shenzhou
	"411200" : "三门峡", // smx
	"120000" : "天津", // tj
	"140100" : "太原", // ty
	"331000" : "台州", // taizhou
	"130200" : "唐山", // ts
	"321200" : "泰州", // tz
	"370900" : "泰安", // ta
	"321283" : "泰兴", // tx
	"520600" : "铜仁", // tr
	"340700" : "铜陵", // tl
	"150500" : "通辽", // tongliao
	"340881" : "桐城", // tc
	"429006" : "天门", // tm
	"211200" : "铁岭", // tieling
	"610200" : "铜川", // tongchuan
	"440781" : "台山", // taishan
	"220500" : "通化", // tonghua
	"370481" : "滕州", // tengzhou
	"620500" : "天水", // tianshui
	"650400" : "吐鲁番", // tlf
	"230781" : "铁力", // tieli
	"341181" : "天长", // tianchang
	"330483" : "桐乡", // tongxiang
	"420100" : "武汉", // wh
	"320200" : "无锡", // wx
	"371000" : "威海", // weihai
	"370700" : "潍坊", // wf
	"340200" : "芜湖", // wuhu
	"330300" : "温州", // wz
	"650100" : "乌鲁木齐", // wlmq
	"130481" : "武安", // wa
	"331081" : "温岭", // wl
	"150300" : "乌海", // wuhai
	"150900" : "乌兰察布", // wlcb
	"210281" : "瓦房店", // wfd
	"532600" : "文山", // ws
	"610500" : "渭南", // weinan
	"440883" : "吴川", // wuchuan
	"620600" : "武威", // ww
	"230184" : "五常", // wuchang
	"640300" : "吴忠", // wuzhong
	"450400" : "梧州", // wuzhou
	"410481" : "舞钢", // wugang
	"410781" : "卫辉", // weihui
	"430581" : "武冈", // wg
	"231182" : "五大连池", // wdlc
	"421182" : "武穴", // wuxue
	"511781" : "万源", // wy
	"130130" : "无极", // wj
	"469006" : "万宁", // wn
	"469001" : "五指山", // wzs
	"469005" : "文昌", // wc
	"350200" : "厦门", // xm
	"610100" : "西安", // xa
	"320300" : "徐州", // xz
	"411000" : "许昌", // xc
	"420600" : "襄阳", // xy
	"410700" : "新乡", // xinxiang
	"610400" : "咸阳", // xianyang
	"130500" : "邢台", // xt
	"630100" : "西宁", // xining
	"420900" : "孝感", // xg
	"130184" : "新乐", // xl
	"411500" : "信阳", // xinyang
	"140900" : "忻州", // xinzhou
	"152502" : "锡林浩特", // xlht
	"210181" : "新民", // xinmin
	"429004" : "仙桃", // xiantao
	"341800" : "宣城", // xuancheng
	"430300" : "湘潭", // xiangtan
	"430381" : "湘乡", // xiangxiang
	"360500" : "新余", // xinyu
	"433100" : "湘西", // xx
	"441481" : "兴宁", // xingning
	"370982" : "新泰", // xintai
	"421200" : "咸宁", // xn
	"152500" : "锡林郭勒盟", // xlglm
	"610481" : "兴平", // xp
	"513401" : "西昌", // xichang
	"522301" : "兴义", // xingyi
	"530381" : "宣威", // xw
	"411681" : "项城", // xiangcheng
	"211481" : "兴城", // xingcheng
	"410182" : "荥阳", // xingyang
	"410183" : "新密", // xinmi
	"320381" : "新沂", // xiny
	"532800" : "西双版纳傣族自治州", // xsbn
	"888886" : "雄安新区", // xan
	"370600" : "烟台", // yt
	"420500" : "宜昌", // yichang
	"640100" : "银川", // yinchuan
	"320900" : "盐城", // yc
	"430600" : "岳阳", // yy
	"321000" : "扬州", // yz
	"210800" : "营口", // yk
	"411024" : "鄢陵", // yanling
	"411081" : "禹州", // yuzhou
	"511500" : "宜宾", // yibin
	"330782" : "义乌", // yw
	"511800" : "雅安", // yaan
	"330784" : "永康", // yongkang
	"140300" : "阳泉", // yq
	"420581" : "宜都", // yd
	"140800" : "运城", // yuncheng
	"420981" : "应城", // yingcheng
	"530400" : "玉溪", // yx
	"350481" : "永安", // yongan
	"430900" : "益阳", // yiyang
	"431100" : "永州", // yongzhou
	"220182" : "榆树", // yushu
	"360600" : "鹰潭", // yingtan
	"610600" : "延安", // ya
	"360900" : "宜春", // ych
	"610800" : "榆林", // yulin
	"222400" : "延边", // yb
	"620981" : "玉门", // ym
	"441700" : "阳江", // yangjiang
	"441781" : "阳春", // yangchun
	"441881" : "英德", // yingde
	"230700" : "伊春", // yichun
	"371482" : "禹城", // yucheng
	"450900" : "玉林", // yl
	"654000" : "伊犁", // yili
	"321081" : "仪征", // yizheng
	"222401" : "延吉", // yj
	"654002" : "伊宁", // yining
	"140881" : "永济", // yongji
	"451281" : "宜州", // yizhou
	"320282" : "宜兴", // yixing
	"430981" : "沅江", // yuanjiang
	"411481" : "永城", // yongcheng
	"410381" : "偃师", // yanshi
	"330281" : "余姚", // yr
	"442000" : "中山", // zs
	"440400" : "珠海", // zh
	"410100" : "郑州", // zz
	"370300" : "淄博", // zb
	"440800" : "湛江", // zhanjiang
	"430200" : "株洲", // zhuzhou
	"350600" : "漳州", // zhangzhou
	"321100" : "镇江", // zj
	"130700" : "张家口", // zjk
	"370400" : "枣庄", // zaozhuang
	"130133" : "赵县", // zhaoxian
	"510300" : "自贡", // zg
	"130281" : "遵化", // zhunhua
	"411600" : "周口", // zk
	"411700" : "驻马店", // zmd
	"512000" : "资阳", // ziyang
	"330900" : "舟山", // zhoushan
	"420583" : "枝江", // zhijiang
	"420881" : "钟祥", // zhongxiang
	"530600" : "昭通", // zt
	"210283" : "庄河", // zhuanghe
	"430800" : "张家界", // zjj
	"370181" : "章丘", // zhangqiu
	"441200" : "肇庆", // zq
	"620700" : "张掖", // zy
	"370685" : "招远", // zhaoyuan
	"370782" : "诸城", // zhucheng
	"370883" : "邹城", // zc
	"640500" : "中卫", // zw
	"231282" : "肇东", // zhaodong
	"320582" : "张家港", // zjg
	"431081" : "资兴", // zx
	"360982" : "樟树", // zhangshu
	"610124" : "周至", // zhouzhi
	"330681" : "诸暨", // zhuji
	"420683" : "枣阳", // zaoyang
	"130681" : "涿州", // zhuozhou
	"520300" : "遵义", // zunyi
}
var CityAbbr = map[string]string{
	"340800" : "aq",
	"210300" : "as",
	"513200" : "ab",
	"520400" : "anshun",
	"420982" : "anlu",
	"530181" : "an",
	"152900" : "alsm",
	"610900" : "ak",
	"370784" : "anqiu",
	"231281" : "ad",
	"652900" : "aks",
	"410500" : "ay",
	"659002" : "alr",
	"130683" : "ag",
	"110000" : "bj",
	"450500" : "bh",
	"130600" : "bd",
	"610300" : "baoji",
	"150200" : "baotou",
	"131081" : "bazhou",
	"511900" : "bz",
	"340300" : "bf",
	"520500" : "bijie",
	"530500" : "baoshan",
	"341600" : "bozhou",
	"210500" : "bx",
	"220600" : "bs",
	"620400" : "by",
	"220800" : "bc",
	"371600" : "binzhou",
	"652700" : "brtl",
	"451000" : "baise",
	"231181" : "ba",
	"211381" : "bp",
	"130981" : "botou",
	"652701" : "bole",
	"450981" : "bl",
	"469029" : "bt",
	"510100" : "cd",
	"500000" : "cq",
	"430100" : "cs",
	"220100" : "cc",
	"430700" : "changde",
	"320400" : "changzhou",
	"130900" : "cangzhou",
	"411082" : "cg",
	"140400" : "changzhi",
	"520381" : "chishui",
	"150400" : "cf",
	"341700" : "chizhou",
	"532300" : "chuxiong",
	"431000" : "chenzhou",
	"540300" : "changdu",
	"211300" : "chaoyang",
	"220202" : "cy",
	"445100" : "chaozhou",
	"652300" : "cj",
	"320581" : "changshu",
	"451400" : "chongzuo",
	"130800" : "chengde",
	"341100" : "cz",
	"421281" : "cb",
	"450481" : "cx",
	"510184" : "chongzhou",
	"330282" : "cixi",
	"469023" : "cm",
	"210200" : "dl",
	"441900" : "dg",
	"210600" : "dd",
	"511700" : "dazhou",
	"510600" : "dy",
	"140200" : "dt",
	"130682" : "dingzhou",
	"411381" : "dengzhou",
	"330783" : "dongyang",
	"420582" : "dangyang",
	"533100" : "dh",
	"210882" : "dsq",
	"533400" : "diqing",
	"220183" : "dehui",
	"361181" : "dx",
	"370500" : "dongying",
	"621100" : "dingxi",
	"230600" : "dq",
	"371400" : "dezhou",
	"232700" : "dxal",
	"320904" : "df",
	"320981" : "dongtai",
	"469007" : "dongfang",
	"222403" : "dunhua",
	"321181" : "danyang",
	"211081" : "dengta",
	"620982" : "dunhuang",
	"632802" : "dlh",
	"420281" : "daye",
	"410185" : "dengfeng",
	"510181" : "djy",
	"210681" : "donggang",
	"469003" : "dz",
	"469021" : "da",
	"532900" : "dali",
	"420700" : "ez",
	"150600" : "eeds",
	"422800" : "es",
	"440785" : "ep",
	"440600" : "fs",
	"350100" : "fz",
	"341200" : "fy",
	"210400" : "fushun",
	"360981" : "fengcheng",
	"361000" : "fuzhou",
	"370983" : "fc",
	"450600" : "fcg",
	"652302" : "fk",
	"141182" : "fenyang",
	"522702" : "fq",
	"350981" : "fa",
	"210682" : "fch",
	"350982" : "fd",
	"350181" : "fuqing",
	"440100" : "gz",
	"520100" : "gy",
	"450300" : "gl",
	"360700" : "ganzhou",
	"321084" : "gaoyou",
	"511600" : "ga",
	"140181" : "gujiao",
	"513300" : "ganzi",
	"421381" : "gs",
	"210881" : "gaizhou",
	"220381" : "gzl",
	"440981" : "gaozhou",
	"623000" : "gn",
	"370785" : "gm",
	"640400" : "guyuan",
	"450800" : "gg",
	"450881" : "gp",
	"632801" : "grm",
	"532501" : "gejiu",
	"360681" : "gx",
	"130684" : "gbd",
	"510800" : "guangyuan",
	"330100" : "hz",
	"340100" : "hf",
	"441300" : "hui",
	"460100" : "hk",
	"230100" : "hrb",
	"420200" : "huangshi",
	"320800" : "ha",
	"150100" : "hhht",
	"330500" : "huzhou",
	"610700" : "hanzhong",
	"130400" : "hd",
	"430400" : "hy",
	"131100" : "hs",
	"371700" : "heze",
	"340400" : "huainan",
	"340600" : "huaibei",
	"150700" : "hlbr",
	"341000" : "huangshan",
	"421083" : "honghu",
	"341522" : "hq",
	"210381" : "haicheng",
	"431200" : "hh",
	"431281" : "hongjiang",
	"211400" : "hld",
	"220282" : "huadian",
	"440982" : "huazhou",
	"441600" : "heyuan",
	"230400" : "hegang",
	"630200" : "haidong",
	"632800" : "haixi",
	"231083" : "hl",
	"231100" : "heihe",
	"650500" : "hm",
	"451100" : "hezhou",
	"451200" : "hc",
	"410600" : "hb",
	"653200" : "ht",
	"410602" : "hsh",
	"421100" : "hg",
	"222404" : "hunchun",
	"130983" : "huanghua",
	"130984" : "hj",
	"610581" : "hancheng",
	"141081" : "houma",
	"511681" : "huaying",
	"451381" : "heshan",
	"410782" : "huixian",
	"141082" : "huozhou",
	"520382" : "hr",
	"330481" : "haining",
	"370687" : "haiyang",
	"320684" : "haimen",
	"532500" : "honghezhou",
	"370101" : "jn",
	"330400" : "jx",
	"440700" : "jiangmen",
	"220200" : "jl",
	"330700" : "jh",
	"360800" : "jian",
	"360400" : "jiujiang",
	"370800" : "jining",
	"420800" : "jm",
	"410800" : "jiaozuo",
	"510185" : "jianyang",
	"321282" : "jingjiang",
	"419001" : "jiyuan",
	"140500" : "jc",
	"421000" : "jingzhou",
	"532801" : "jinghong",
	"210700" : "jinzhou",
	"360200" : "jdz",
	"620200" : "jyg",
	"620300" : "jinchang",
	"620900" : "jq",
	"230300" : "jixi",
	"445200" : "jieyang",
	"230800" : "jms",
	"320413" : "jt",
	"140700" : "jz",
	"433101" : "jishou",
	"350783" : "jo",
	"220281" : "jiaohe",
	"330182" : "jd",
	"131103" : "jizhou",
	"321183" : "jr",
	"330421" : "jiashan",
	"370281" : "jiaozhou",
	"360124" : "jinxian",
	"510781" : "jiangyou",
	"320281" : "jy",
	"410200" : "kf",
	"530100" : "km",
	"320583" : "ks",
	"440783" : "kp",
	"522601" : "kaili",
	"532502" : "ky",
	"211282" : "kaiyuan",
	"650200" : "klmy",
	"652801" : "krl",
	"653000" : "kzls",
	"653100" : "kashi",
	"654003" : "kt",
	"131000" : "lf",
	"410300" : "luoyang",
	"450200" : "liuzhou",
	"620100" : "lz",
	"513400" : "liangshan",
	"371300" : "linyi",
	"320700" : "lyg",
	"371500" : "lc",
	"510500" : "luzhou",
	"411100" : "luohe",
	"330781" : "lanxi",
	"331082" : "lh",
	"331100" : "lishui",
	"331181" : "lq",
	"520200" : "lps",
	"140821" : "lin",
	"141000" : "linfen",
	"141100" : "lvliang",
	"341500" : "la",
	"530700" : "lj",
	"530900" : "lincang",
	"430281" : "liling",
	"430682" : "linxiang",
	"350681" : "longhai",
	"211000" : "liaoyang",
	"540100" : "lasa",
	"360281" : "lp",
	"431300" : "loudi",
	"440881" : "lianjiang",
	"220400" : "liaoyuan",
	"440882" : "leizhou",
	"370682" : "laiyang",
	"621200" : "ln",
	"622900" : "lx",
	"441882" : "lianzhou",
	"371200" : "lw",
	"371481" : "ll",
	"371581" : "linqing",
	"320481" : "liyang",
	"410581" : "linzhou",
	"451300" : "lb",
	"511100" : "leshan",
	"350800" : "ly",
	"222405" : "longjing",
	"220681" : "linjiang",
	"211382" : "lingyuan",
	"140481" : "lucheng",
	"422802" : "lichuan",
	"431381" : "lsj",
	"431382" : "lianyuan",
	"511381" : "langzhong",
	"440281" : "lechang",
	"441581" : "lufeng",
	"445381" : "luoding",
	"350122" : "lianj",
	"370285" : "laixi",
	"420682" : "lhk",
	"430181" : "liuyang",
	"370683" : "laizhou",
	"469027" : "ld",
	"469028" : "ls",
	"469024" : "lg",
	"540400" : "linzhi",
	"510700" : "mianyang",
	"340500" : "mas",
	"511400" : "ms",
	"341182" : "mingguang",
	"440900" : "mm",
	"441400" : "meizhou",
	"231000" : "mdj",
	"410883" : "mengzhou",
	"230382" : "mishan",
	"150781" : "mzl",
	"421181" : "mc",
	"510683" : "mianzhu",
	"320100" : "nj",
	"330200" : "nb",
	"360100" : "nc",
	"511300" : "nanchong",
	"320600" : "nt",
	"450100" : "nn",
	"411300" : "ny",
	"511000" : "neijiang",
	"350583" : "na",
	"350700" : "np",
	"533300" : "nujiang",
	"350900" : "nd",
	"540600" : "nq",
	"360703" : "nk",
	"230281" : "nh",
	"231084" : "ningan",
	"341881" : "ng",
	"130581" : "nangong",
	"410900" : "py",
	"510400" : "pzh",
	"210214" : "pld",
	"350300" : "pt",
	"211100" : "pj",
	"360300" : "pingxiang",
	"620800" : "pl",
	"445281" : "pn",
	"410400" : "pds",
	"370283" : "pd",
	"330482" : "ph",
	"320382" : "pz",
	"530800" : "pr",
	"370200" : "qd",
	"350500" : "quanzhou",
	"441800" : "qy",
	"130283" : "qa",
	"140121" : "qx",
	"330800" : "quzhou",
	"520181" : "qingzhen",
	"522600" : "qdn",
	"522700" : "qn",
	"530300" : "qj",
	"429005" : "qianjiang",
	"370781" : "qingzhou",
	"230200" : "qqhr",
	"621000" : "qingyang",
	"370881" : "qf",
	"230900" : "qth",
	"450700" : "qinzhou",
	"410882" : "qinyang",
	"130300" : "qhd",
	"640381" : "qtx",
	"320113" : "qixia",
	"320681" : "qidong",
	"510183" : "ql",
	"469002" : "qh",
	"469030" : "qz",
	"522300" : "qianxinan",
	"330381" : "ra",
	"130982" : "rq",
	"533102" : "rl",
	"540200" : "rkz",
	"360781" : "rj",
	"371100" : "rz",
	"410482" : "ruzhou",
	"360481" : "rc",
	"320682" : "rg",
	"371082" : "rongcheng",
	"310000" : "sh",
	"440300" : "sz",
	"130100" : "sjz",
	"210100" : "sy",
	"320500" : "su",
	"460200" : "san",
	"330600" : "sx",
	"361100" : "sr",
	"440500" : "st",
	"321300" : "sq",
	"321322" : "shuyang",
	"510900" : "sn",
	"411400" : "shangqiu",
	"131082" : "sanhe",
	"420300" : "shiyan",
	"421081" : "shishou",
	"421087" : "songzi",
	"341300" : "suzhou",
	"350400" : "sm",
	"430500" : "shaoyang",
	"350581" : "shishi",
	"440200" : "shaoguan",
	"440606" : "sd",
	"220300" : "sp",
	"611000" : "sl",
	"220700" : "songyuan",
	"230183" : "shangzhi",
	"441500" : "sw",
	"370783" : "shouguang",
	"230500" : "sys",
	"640200" : "szs",
	"231200" : "suihua",
	"330683" : "shengzhou",
	"231081" : "sfh",
	"510682" : "sf",
	"350781" : "shaowu",
	"430382" : "ss",
	"131182" : "shenzhou",
	"411200" : "smx",
	"120000" : "tj",
	"140100" : "ty",
	"331000" : "taizhou",
	"130200" : "ts",
	"321200" : "tz",
	"370900" : "ta",
	"321283" : "tx",
	"520600" : "tr",
	"340700" : "tl",
	"150500" : "tongliao",
	"340881" : "tc",
	"429006" : "tm",
	"211200" : "tieling",
	"610200" : "tongchuan",
	"440781" : "taishan",
	"220500" : "tonghua",
	"370481" : "tengzhou",
	"620500" : "tianshui",
	"650400" : "tlf",
	"230781" : "tieli",
	"341181" : "tianchang",
	"330483" : "tongxiang",
	"420100" : "wh",
	"320200" : "wx",
	"371000" : "weihai",
	"370700" : "wf",
	"340200" : "wuhu",
	"330300" : "wz",
	"650100" : "wlmq",
	"130481" : "wa",
	"331081" : "wl",
	"150300" : "wuhai",
	"150900" : "wlcb",
	"210281" : "wfd",
	"532600" : "ws",
	"610500" : "weinan",
	"440883" : "wuchuan",
	"620600" : "ww",
	"230184" : "wuchang",
	"640300" : "wuzhong",
	"450400" : "wuzhou",
	"410481" : "wugang",
	"410781" : "weihui",
	"430581" : "wg",
	"231182" : "wdlc",
	"421182" : "wuxue",
	"511781" : "wy",
	"130130" : "wj",
	"469006" : "wn",
	"469001" : "wzs",
	"469005" : "wc",
	"350200" : "xm",
	"610100" : "xa",
	"320300" : "xz",
	"411000" : "xc",
	"420600" : "xy",
	"410700" : "xinxiang",
	"610400" : "xianyang",
	"130500" : "xt",
	"630100" : "xining",
	"420900" : "xg",
	"130184" : "xl",
	"411500" : "xinyang",
	"140900" : "xinzhou",
	"152502" : "xlht",
	"210181" : "xinmin",
	"429004" : "xiantao",
	"341800" : "xuancheng",
	"430300" : "xiangtan",
	"430381" : "xiangxiang",
	"360500" : "xinyu",
	"433100" : "xx",
	"441481" : "xingning",
	"370982" : "xintai",
	"421200" : "xn",
	"152500" : "xlglm",
	"610481" : "xp",
	"513401" : "xichang",
	"522301" : "xingyi",
	"530381" : "xw",
	"411681" : "xiangcheng",
	"211481" : "xingcheng",
	"410182" : "xingyang",
	"410183" : "xinmi",
	"320381" : "xiny",
	"532800" : "xsbn",
	"888886" : "xan",
	"370600" : "yt",
	"420500" : "yichang",
	"640100" : "yinchuan",
	"320900" : "yc",
	"430600" : "yy",
	"321000" : "yz",
	"210800" : "yk",
	"411024" : "yanling",
	"411081" : "yuzhou",
	"511500" : "yibin",
	"330782" : "yw",
	"511800" : "yaan",
	"330784" : "yongkang",
	"140300" : "yq",
	"420581" : "yd",
	"140800" : "yuncheng",
	"420981" : "yingcheng",
	"530400" : "yx",
	"350481" : "yongan",
	"430900" : "yiyang",
	"431100" : "yongzhou",
	"220182" : "yushu",
	"360600" : "yingtan",
	"610600" : "ya",
	"360900" : "ych",
	"610800" : "yulin",
	"222400" : "yb",
	"620981" : "ym",
	"441700" : "yangjiang",
	"441781" : "yangchun",
	"441881" : "yingde",
	"230700" : "yichun",
	"371482" : "yucheng",
	"450900" : "yl",
	"654000" : "yili",
	"321081" : "yizheng",
	"222401" : "yj",
	"654002" : "yining",
	"140881" : "yongji",
	"451281" : "yizhou",
	"320282" : "yixing",
	"430981" : "yuanjiang",
	"411481" : "yongcheng",
	"410381" : "yanshi",
	"330281" : "yr",
	"442000" : "zs",
	"440400" : "zh",
	"410100" : "zz",
	"370300" : "zb",
	"440800" : "zhanjiang",
	"430200" : "zhuzhou",
	"350600" : "zhangzhou",
	"321100" : "zj",
	"130700" : "zjk",
	"370400" : "zaozhuang",
	"130133" : "zhaoxian",
	"510300" : "zg",
	"130281" : "zhunhua",
	"411600" : "zk",
	"411700" : "zmd",
	"512000" : "ziyang",
	"330900" : "zhoushan",
	"420583" : "zhijiang",
	"420881" : "zhongxiang",
	"530600" : "zt",
	"210283" : "zhuanghe",
	"430800" : "zjj",
	"370181" : "zhangqiu",
	"441200" : "zq",
	"620700" : "zy",
	"370685" : "zhaoyuan",
	"370782" : "zhucheng",
	"370883" : "zc",
	"640500" : "zw",
	"231282" : "zhaodong",
	"320582" : "zjg",
	"431081" : "zx",
	"360982" : "zhangshu",
	"610124" : "zhouzhi",
	"330681" : "zhuji",
	"420683" : "zaoyang",
	"130681" : "zhuozhou",
	"520300" : "zunyi",
}
