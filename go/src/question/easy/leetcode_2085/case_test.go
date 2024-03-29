package leetcode_2085

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"plsfzfgawptsrhtppzhlcjqbcmi", "gsccupsxkshslvrixubnxdhcsjshp", "s", "ezrflcmchmhhvojohew", "lmskv", "fxhnharndihttfrnccktwbmxzpqee", "hd", "mcrcjipns", "e", "qrnqz", "csbtrpsmhdnafekmns", "wjb", "xmibfbabpomcgyhb", "hruigjocaexurmfljszxvlxkyolqky", "kiigmajzxjocnixufcy", "ckklctinivzecvhocjbsxymrt", "kxfzrrv", "qjceabulyrwatm", "ojafmsymfkotalkvsufnprdu", "wmkpuftthlvycyqmzozgvb", "umhjoqoenipkhumbikswzheqe", "hgyvfnvjosgxefphhzr", "rqxhyspbloycsxovno", "paicai", "zpnmsfkdljsscaddgwhirtqntv", "fysxtnbcohdlcpibrddtoa", "icwlkeyaiuzasuntg", "wmloodipkkplkbsico", "tbffw", "whlygbzsrewchnbtt", "gip", "psqxgfhpioqyw", "uwbnogpqsu", "ioibrgngksybilgr", "zojhxjnxqfojsfafkuvsiyimewsepo", "vfjvmojjizmvfzbc", "wkjnunyavkdomqcsrowiw", "nbbuavuhhomg", "qsjvrgkypqsxioefxkwh", "boejyfmqyvsbpugbqyrkpl", "efwlbrnpkfc", "huiaolundrnxezqqzyiqauboqm", "jmkjyizpmdumpw", "oykzgyoocwavrjum", "f", "dqurh", "cwguhiozmaeozcqwnffgexmfnku", "jxvlggrhculruzgfkokpmgrbyftdo", "qvrtqvchslzbykfwuxayywnracnlq", "daelyspayiondzyq", "eqghkkijnripfzksjrwulmxbaa", "zmyvzmwwhbgyfcqbbt", "mjucbpastaer", "phyagrxzqrkn", "nnavmxqyfczdd", "irqxbfyarmjeoiomprfjequksm", "ggdht", "hofvrhohgxvevxgjtgfprsmooo", "ntztieup", "ueflxsfystwopaopgldaimuipsaful", "iydkzacinibnpjgljls", "ihsropbisotqwhpndcqq", "lahooipcdsetztoncvjauzcnqh", "sobdrebzxpwxgvqveayc", "wcnddeqeqtdrerbbdwnpquhaxbbpe", "ithgkrvauylxnecjzkbroiaptu", "duceyy", "fhcyatyjbcmbftigjihvpceit", "ygtswpbjneglmhdgwtgrv", "g", "nzogophfgwdotbttenoxy", "lju", "kfb", "sgktjimgogrwvigpfunvfnxc", "rpavqupjhpuj", "vs", "hygvurijfrsrazwvxp", "rablosnrbxhf", "njiswypyccwn", "btvymedpiunurhaqtoxzxaxnorrf", "evxjylrnnbwlocmduito", "zjdduajyynlsuyymugcuwg", "qzesdxpfsj", "a", "iyouogtkdygmmqwmzsjdeppyoskh", "mwsvjmtbrfnm", "kz", "cvvqjhlrhzzmmntmbjyza", "wwaazasyy", "tyzeebwhqmcobmakafgebbwtigbi", "ginky", "dtalssxkpyfqaurgz", "yafm", "qhxehrmsljaufkvhvuwvmnunjno", "y", "jywujvcljmoqghtz", "bwjjtjrtcwlv", "iclszrajbgqvmsuh", "hdwofeucbvxxrsfaiikghvjcahwvy", "hjruyhqev", "drcgjidxzlucrucxdxd", "impxzwjxmiosor", "mkgdfgguom", "mmlbqulasotujisgwpkzoejsc", "lpckqobfpjzvy", "eeabxbb", "gmrjtyfkwgzaiabkirptgnfutztd", "hechoogbsykpoplww", "tyjlifgtizyvpqisgipjk", "mzzudpnujooifzqgatprzfejhmi", "cudoqctbacrxrewlrscddzgqvnyb", "hzhpnymwjhuaoflycsfopmvqcxxv", "pjgs", "ltnsokdyld", "nwzqpqymrxtmbdf", "nxqcezojdgswdx", "iimkiwkpnvuufvfgpmrpydwem", "lzzzxnztaigeuexwjkuve", "vzvmnh", "hqtkmfaguqowsmoqymilhuyqli", "jkj", "xmynmjfuvfjzwbyzllzuepkftsk", "adchcrgqoz", "iekocgvlolololo", "ozgzvbokbauwkipdeyfunptjco", "sdvtnigwgleiral", "ngvwegcsfwlzedmwyzghygj", "acbiywsbpifabqrtasbsiypwtm", "kextotfhtjl", "bgbeuwgfmru", "sprleyhjucfcli", "jchcophgucybaaexrbtfnfmcats", "cqosqjuxlzmseaweieeaposqpaa", "oalrhiwqccgsphcgwjyx", "uhakpqbmflimcknywbzcrwukzmkum", "iwaassehsaojeheeld", "brly", "kq", "wzdvbyfyx", "gzjngsmlxhu", "ofrupfvujzjexteznqu", "tzhwyzpyi", "cnjjpjz", "sfwedbvjwuxkhqzfglqhutznqvx", "xnuopczllpxqga", "gvzapdgeqezouspqbohbgsvodjdid", "aqyfe", "l", "uqiwis", "s", "jxgkgjgzxaunyledp", "bfwivzblomvixmxhrdxomldst", "ejqny", "rjizxyahdjoqnxvalzzx", "narkxounvlpmovyunkl", "uxadyjuoyhvybacvaffsfhspl", "dpjitixb", "mkqiiunsffxpvlysnqpbkrutsoo", "lnrrvbxxkqxdivr", "bhqkobemsduydnsea", "emyejsthfl", "cchkfx", "armbwbclk", "w", "wlzdkkmrmnhedtmipy", "pnymgukydtluubrwkngaohqp", "xuirrfqfyoimwm", "xllhirakkfdajyglr", "gcempkdx", "ahqgqykzitbqphjrorfsdknalmf", "krdbk", "ornevzpgknjx", "vnnvictiqsxwbujcoilmgmywh", "lyohgxylwgrlgsdhagtfgh", "swjetqfxmhbbsplionglesgwtobjig", "yvjyxcpfqzaiwgqw", "uanxy", "chggkz", "chnpnzeizdnliczrmvssguap", "zrrttunymnppvlj", "nrpdmrhwmhzeukjwiiersmlhohimmv", "lzhzr", "ywkymfcgittyxujrjdyvxoy", "agtpspxhilmmlspz", "orocovqkhxrafbjogqfizqe", "apdfsxnjmmqpymeiejrf", "gomvicmc", "et", "ccq", "yxqrmxtrbmrb", "ufniqsmdruabolytejvcgfzw", "vilollcowereynhexxslxggbdqdiv", "ewhszmrpvwkfvhub", "ropevek", "oycxkv", "qsdppfiwxlqahncomkhla", "dtnaflx", "vfseztyahalmrywwvxworxown", "chgbawrqo", "rvfwrqqgmredhitfuieplodgkbqke", "yrj", "tzuplosuzflzjbwl", "volgapwsdopoakj", "ymsllv", "exwegnmdksbt", "vkddjfkssoxfczokjroijjxvjxfz", "fuzpgwk", "byglewgcvgmboxjarhurztb", "hdyijxwbyymq", "bdow", "fijvnajwmxpfvxhfmdvgvzal", "bpxhuqjzfuxijqtpbczpmpo", "mmjt", "tvrwttkamwhwwwe", "eglznvgvwobujenmfhmbvutwydouy", "alvmrqlkvn", "grbhrsyxraezwvzjjlpyqvovtinb", "aijr", "xv", "wiefix", "ifvnxukkdkoal", "nifizflkpvhpdefbqeg", "ksnz", "yemju", "b", "beezbouznuehxuoxrwcxppzzfdii", "bmgmuaplw", "yorkpqtf", "gx", "pnjujio", "nuablp", "ehcwkkzztgavpuczgmbylahctkmahj", "kipvctrf", "omuowicgxorgvjposylpf", "sqpvrxaf", "fxahaigkuljflmymfdvxydncsc", "koz", "yhunffzpa", "lxsdrtfimvmp", "wbboihhzzsyngjw", "pxlbtcjcyeyltgjqmiaplldoih", "tvfgucmqsqawkoulnw", "zgikecpxgkvtyttkblqpgnypemhc", "kpxiescejxnzjbzmhjpadoygpufudg", "pzzakdgffmupv", "jrrlutbebsxzmxwlhykulprzv", "zcerueeudguhodsq", "pddqgygxxd", "fhtgwy", "lihwn", "kvbhujftweg", "mmj", "yuxrpaachruyjki", "nuaowzzownvzofuoovcbl", "jijoiywepebahhqdv", "bhdkwiqchmrqjbsdidxct", "onh", "snelkpzexsfub", "rhjednmlueduletulbfqywwjylyna", "lpbfygesivkzwveukcm", "epyrzjqyafhxddrxwgsowmsmyevkdw", "uktpvfwcxpdsbcab", "vlfhtlqdkfggxusma", "hwpqpwnyhlsokrnmqsjoiwzzlbaylv", "v", "chlkdhpklasay", "xxmzssfoezczvzgpfo", "jpjtklksindyjpvxhejtucset", "fvpnqduebposjklmptdgkaq", "fmcuvpikdrkmizblymnnmwkf", "niqlyrl", "qyvabwvkhxntbacvhh", "vtvvkozwytlhotpop", "lmcxjhudvumnlkqqatflyy", "lgxzmasivwpcnnywnuufsbqesoscr", "yjsfppmproitzv", "lsbkldsem", "tcttcfjhgqnqhvqzyrzjxfn", "udmole", "tgyjtnxdghchkxkuoghexfeydburw", "yxllvjhndb", "qietoq", "uzhhtul", "exzqsigutm", "as", "viamlaohkbtrqmdnunrdowqtgsufkr", "oylmvpsskwlpdtmztdklelxvjux", "gubgujoxhecsq", "efalejqivgmq", "w", "vjxgyddxmq", "ujrpd", "fszpdvkmtqvmilcdv", "xjminawbblb", "ibr", "fewgitzw", "hprffroywugiqixdknxtpjaemi", "tzoyuqczdnfemsgsh", "phumhrexialxyxvm", "wganyhzv", "jzgcvsbkwiuxwxgylygy", "pipxxfxwabozs", "uzbdpprpidfpxcxvlbwqxxc", "dexkfmajwahmpjflbpgkyxr", "hngxdllbufoxleyupznzbczfegpuf", "rbwvfflgygzp", "lerks", "ugquinvxvmunbobyrcijzdoicfmc", "bkhxnvdpugculmh", "melfmzenlvvwpfkjnxofddey", "bogwobnsdavmjsxmxjwyleun", "jq", "ytxacmohwkofts", "jdsmtiozbthrloguxrzjqfhxxme", "bzr", "fyukylveigjhlfiossgg", "kiipzcvlkhohfq", "fybxdscdofqimy", "m", "dzdjmuttofxznkvsod", "rwu", "ndxzwenrdcs", "uskykkpcc", "fkxcdrlkdgtmfaiagufsj", "mcray", "lginz", "cizjbmlxvnzsjrftbtskactlhmz", "daeqjfvdrxiatudqtvxkmxphr", "imrtbxgvhnritxey", "hozwsvgpzkxtqxryw", "aazwqjdjnmkslyxghvne", "eeixrpaitag", "diey", "kfiedkessuqkzgkngjmkrf", "zezwofckwwaomftf", "ueezorhcwakyvxc", "yjxdpjirmlu", "xjytzoijlxbn", "gzaviosoghcbbctkiamrqjgkuu", "tdeqplbibvuk", "jdbxxvn", "p", "zpbwccizb", "edyiiiyytlmcuqpwfmxxsohsoc", "v", "hykbrkqquxzallurqolcmxmlaob", "vgfiaxcyebvbgadcgo", "qzdlklmxudpjxayxiayjwtatagnxn", "mxbgmhobasjyizza", "cqsgdpdcvguopyjedi", "vwpmuqsyjwhwz", "ymzf", "dvj", "qtvyfppmtiwxziyzxkdvohelxmky", "ghomebkvdmegwqmgvpghtmvqx", "eirhdshiwbyiquhyyobnewvtplfgos", "bbjpktlbsckegcveathdoaprpaiiqm", "ujyeuhxiynwwhqblghcphvjkkvrlgh", "qp", "efcvlwujvbokzhnedwmlmceasa", "elnqxwypigkcrwksxzivm", "jviuoywbndunuufqqze", "bgxab", "iirugjnftmpue", "kec", "zzlstxdpqgqbyzsfdfrcbwse", "ulbicasvbqyzpxxlhqilwhqv", "dswdttupegpi", "joervjpln", "fzr", "oyljhnpab", "zfcefvmfvwwqgredyvlqcpftryx", "ozco", "iwihjdwncjpcqayvjaxpqozef", "aimibrckokvzu", "kwokiotuincktuihdh", "xpcervhuhnonfltyoqo", "emtvfafnrxkskzluueftozdc", "nzzcrp", "qgxevxqqpoapwzvomniwwvjecoddqg", "qmqhnhlhpzzsilzszmxgiywwabedx", "oledlvkixxlniku", "cw", "nxrwiqjfpf", "aecjgrlshvckeairimjzmt", "wlivvdjvxhagdrvjgna", "ts", "hxzxuyswcslfjurtgsk", "hcoomavhnegnek", "zxtctesskgziqnmgobjfexjpzbehj", "rekengyardexvhpavvtrgjyk", "upsvbkohtcmhb", "r", "ivqcjggr", "xjmajfjgbztfbuclpkjgroieqaympy", "icymijcffnwjqqye", "yllrgl", "vf", "fnnlgrgcedbdjckyzpdglanozd", "krmkafyuszgiacaxgvauohxd", "jahf", "zpcgg", "lnwsexzg", "ol", "qslmoxaukcqcxmafx", "couixcjyadowkyjrqpjqf", "nrspfiy", "gyzmorecjuqhnsswlfakkucooq", "mtxnwhkukkzbmwinkempus", "oxu", "xpgauwpehnopfj", "tjojwxuptjy", "whyjcgipmzjuz", "ajqfwhhmbmilkuyonqm", "lzvlzatwyfpzjykqyqffdo", "qmzvwfdnqaeeqwnq", "uppfsekumxkqxfacnwerhh", "rwhdsiatdjmefqqpvrigyqx", "nmydtkrjdsxgongzvqhguy", "maicftfg", "lqxbuixfbtn", "s", "alnvwoelsmgudfngj", "qgybscbkzwacbmewgxp", "fhqkckrjo", "sxhcrebvnesoombuojxegnv", "iswgoi", "yr", "qbggupjlgozeeaixzkxjttbwqbzur", "rydgau", "hxumnhwkvnumr", "cczsldgguk", "qdldwaeagjxwrlakmiloqzan", "xxduxsvxbcghbqlowcfosciscmp", "rbtkcvtolrnloqdwmckflkp", "prmp", "whplwmndffhqobx", "d", "rncewagdkchrhtjrnzmadpcwfjb", "eobptzai", "rwk", "jegkfltfhuwrtnizdnwiws", "klrpovcwydaemnohsie", "ayzvgoibazfskub", "iofebmcsyeetbpwotiwh", "fuferesjlbfttbaewyshhrtbxzubwz", "rrtkenfavrcylu"},
			Assist: []string{"kqn", "eb", "gxfghgnbhtovustrozx", "odhlkiozl", "epercrztztzezrtbl", "flxtxebczkieasv", "xmqasqngd", "wnfzotywjgcdxhwpvna", "sxdagxl", "dpzvqzk", "ecjkfoscbegkehjivusighoqqnk", "rqfourzdrcznxk", "oqru", "y", "ung", "qsrqlewfsxvalrqwkwxgvijexcz", "cmsxlyopwcccjcj", "jxymbeionjjkdtqedtgj", "iaozuvwrdf", "cfouijvujkydd", "pgoxrmocffagasup", "dnokakeyuzkfq", "hfulhdmo", "rdxztedplabiuztycfdfgeq", "vhxxgnz", "kekefyp", "enoxfxeokmkqjsia", "wtdnmypkd", "xkeka", "st", "tontipglxozzdbisvivvo", "nokgykgyasqdjawogrzwhyt", "zbvdxczpsvnmdfxhnfvkqbagllqy", "n", "fobiqnitfwnkk", "e", "flirzsqhedcvzqkdolizgkpvd", "maolhnokjhfhfmagsguzyidsw", "meuiinfkihlhvpnpuqzsozsgavmxym", "duifeolcntqb", "lgwfxarv", "bwmxltx", "edbjcvtuckvlrhlzzxzhpzuascoc", "vk", "fqdsbpcettrclxm", "tlsvzglozwlxopzfr", "ecozqjgaqkydphispvthfk", "mfwnnoorheu", "hfcjipwjeydmrmmsaepln", "ddrjgzbnhcjkyjjpobadrfpczsler", "dotnrswglfbgvxzcm", "ojanpanmxlgijinesftli", "sosrajezvi", "zntoxwfbsqmwbiuinxjmwjlvjnglj", "rjgzbwbopuuguvomdrhjd", "hlcjjedjuthesdtkxulbkf", "fqquhhppufiwk", "yllulqmjkbborqdjadtnotyy", "kkoioemifyifbk", "icw", "oyirnuaicvdtdojtac", "jbrgthooif", "hkmbqlyxxqixpip", "sajllhmynbyohpszhlrgsywb", "fwoflyrshsnvenegpftkowsw", "xeosmulijkgkvssmyccxzqrzx", "yunftqycasmvzaodanzvetqudxi", "muusipbyc", "cbatbdbvzkuinz", "afhnfwiuxyohxsr", "skyhv", "adqulaxacdapzwum", "ociilqnruscpgfi", "vedudlxcoybmrkkhsuw", "esnhkkqzyzpat", "kmwdlpfkurztgmvkjxjaaot", "eduwvewmpmxi", "gutstyzyvypmnvbvkpiewqiwgbr", "jyktjhztps", "keocndbcjlqhowjfx", "lavyidmwcttjbjoxyuoloaamjjijp", "oymkjiodvdtgoxvrwtgyuktbecdmo", "whldzabrvr", "ptsyvusbhkwuxctf", "zabbnaqzyangne", "almivgofy", "voopqlzufjwzgtpavyxhreg", "whxcsjcbznwzzxavafgtg", "kxxzrheqylitupzrqnctohtvvqcab", "px", "eccqatnkwbvsrgmd", "ivczfuwypdqtsmbtgmxxcxaus", "bjf", "hbtohelg", "mrwtmatce", "pp", "wpycdwttrapxai", "jdymvpzptouy", "wagbsctivcseqtjaumhbflt", "upuwtlqugbhvxsxgtwrb", "yepspbp", "olcrdqavrthqdbasfnt", "mfkfoczefkaczylkgif", "pd", "rpppwxnupezk", "ivs", "buvpbsxftxdfldspyfxrsbwl", "uhvdkwlwjdvznntftlgii", "lblsoyfgy", "msmfswsnjksztmvgdypjajdczvn", "bqobaouojjm", "ojintbkd", "tc", "anucgoziakuo", "dfxovsgwghydesjme", "s", "vavapblsflgxkvcm", "rmledxxatstcryjtib", "lbdli", "wm", "cvzcdkzz", "dorqwghzzbkangcrttdrdjovp", "gfcrufdagwkz", "hmvbhdsgsrfmi", "vgqcsarsbsjovmggclemfatvoisdu", "sdbonteic", "obomsujilwtpxtrchum", "tipq", "ajbngjobftuhd", "mvrew", "uqcytkmzmjfpsxwlfgqejcxies", "pvbqmajwbrphpz", "affbrxzgfoohlyc", "gotqcpj", "iqzbkpkhxlbkrxuabhd", "sf", "tqguluxwrqghwfwfzqb", "hkudatagdnzqboafneftqjxcrhxgw", "uhdeuebnhmvsevaxqiseqw", "djgajmhofotlmnozgdi", "odwqtgqtbtwrmuxhuepnmlyqzkzpiu", "lnqdzftjftokwgn", "rydjeovuhwgbvmidzxrocdldrmxsmb", "xfhnzkxgedxdupwtptff", "nxuxx", "hghjisddfamudmh", "hlyitt", "dw", "pesyfl", "ccnaipacnankyiupbhjivra", "abajywpwvdxuqsggpe", "msrskqfaaagkwhsvkqdelx", "yx", "ihryvvwttkzjdxyqv", "wodzvtpxclcdfawdkndmzpfvgzbp", "feznmuumsdsrqirfpvj", "ycdffiqqpyay", "obodlrehapwcblfsuuufjvbgbidu", "b", "awgjwfednpqrvnjozfqvmuydzcjed", "nxsry", "budezoplspawxzd", "soetyzvkcugqtodojbfelwjr", "zchiualewvmwosfagzulhkulixwa", "avaududitkvqkagyqqlumjsfbcuoq", "cueajr", "vsxphjelkkesegrout", "yos", "fndafnvxixwoyqjhrlzvxclgwhsfwl", "bqtzvrkl", "nwmwydrcmvcmkrur", "pwryqfsjhdwzoyfyotctb", "tutmcouczivtctrwlcyte", "qmgoohfiamcmqoyvrzzncvzqvuzcrk", "htogy", "bwworpdpnqnxjedrcteiebghqdh", "orkvieoajrxhqgydmjzchwcskksjex", "ic", "t", "pamgcrdezjeaslpmqqjncmt", "uvdmyvz", "jcnhfayreaylnjxpj", "efh", "rxknljpg", "yuwtalfcixee", "ycovnbwpiyejawmmi", "ybprv", "qlrcigxwljoiaxwlssbuozndmcx", "kyhygahrnkieuozpapvyfcc", "pwzfgkhddzcz", "yptmxlydnwyh", "mihukhfshtbp", "dvvcejirjqhpjjbohvtgs", "dxqsvykwhoqdurbnddzclrr", "dpwphfhpgrjvaapyvqnubscaq", "rqmv", "txqolpuiwbjltrnsqwhadglbntxhe", "crqcvylhjslhwhdqs", "izvlucjobdwiayaqvqpvov", "fyvnzliuzbqlwswrhsrzngiyn", "fhqrsqqoow", "lqwdmytnhca", "dljiw", "vttothnvznobyixmrknylgqni", "xnwzsfoscdxwaevehhsgliqlnwwp", "goie", "skobcntqszen", "viydnkccafcrlmfawortdzxui", "ipjaffchpzhxnlnbqow", "trofqukw", "yruowdttgp", "txnbtlyzmbzrfuecbufw", "vixmylkmlrgzjqpyutvqcelk", "miwezqzr", "duuiiknfcrm", "ffdsvnywehqyuzyfbrkhuiozfxorpe", "hufrczlkitvx", "bmeocwxatfxohxfqfa", "lz", "cjlflqcmybtpfcrnrtxthivanswj", "avwyurfesefufzbtendcs", "uxjxofpxncbadyypntfiqqabww", "pxfsqhyyqriawwayomeghlnygl", "hytdwdobtwhcjh", "ceayzx", "k", "udaztje", "yhyymkpidfmasoqvpdminjillmm", "iqfnbhiqsdbrbggrdywzmurx", "qpoag", "lsnmcrlqdrjmzqx", "kphyex", "oidlopayreexiyfwptnsxacxt", "lzuuwzenlwdklpohb", "jlnufqqme", "nlvwzlxhockyoucdjynulddzpylt", "tybiksxwuumpxwnuxfeqgizkf", "clohfofkkkfoffnh", "gqff", "twqyle", "ug", "fclmkugcmtqv", "kxjfiwxpmhmthijjiexbdgcpujkbn", "pdmlmppbrgi", "udvigkpzkssgasnhwhkivxd", "ytuhsbggfcxfzalzacxrrkmibx", "trplynrpejbuvxulbcrbnxueyhnxxd", "zhbimidnledyheunkguehanhomnjnl", "aervcdjl", "ixpjfbaly"},
		},
		TargetData: common.TDefault{
			Value: 3,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"a", "ab"},
			Assist: []string{"a", "a", "a", "ab"},
		}, TargetData: common.TDefault{
			Value: 1,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"b", "bb", "bbb"},
			Assist: []string{"a", "aa", "aaa"},
		}, TargetData: common.TDefault{
			Value: 0,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"leetcode", "is", "amazing", "as", "is"},
			Assist: []string{"amazing", "leetcode", "is"},
		}, TargetData: common.TDefault{
			Value: 2,
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := caseOne(data.SourceData.Value.([]string), data.SourceData.Assist.([]string))
		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)
	}
}

func TestCaseTwo(t *testing.T) {
	for _, data := range dataSet() {
		res := caseTwo(data.SourceData.Value.([]string), data.SourceData.Assist.([]string))
		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)
	}
}
