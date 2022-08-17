package unfreeze

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/33cn/externaldb/db"
	"github.com/33cn/externaldb/util"
)

func TestConvert_ConvertTx(t *testing.T) {
	convert := NewConvert("", "bty", nil)
	convertCreate(t, convert)
	convertWithdraw(t, convert)
	convertTerminate(t, convert)
}

func convertCreate(t *testing.T, convert db.ExecConvert) {
	blockByte := "0afc0612208d7ded42beb3f761348dbcfee7b2ed4f89996a8bb2c0b9b7c44d55e266c814391a20ac659c9cc3b49d85a4ef42955f8a20f63f681789d9d67dd8a7eec138030312b5222061288501954209741ad992e781ad5c5585607d7591e22b5ed5c568243051c1ae28a30130fcb9c9ea053add030a067469636b657412b102501022ac0208ffffbff9011080cab5ee011a70313271796f6361794e46374c7636433971573461767873324537553431664b5366763a3078386434663130653666313762616533636538663764303239616263623461393839616563333333386238386662333537656165663035613265326465373930343a3030303030303730383822066d6f646966793220e0b8f7b37096a7e42dad587ec8879b0cd4f9af588cb01c5bed6e5628954f34cb3a81011bffed96172ae5c5dcc7d7787c332ca2b94c3ef81ddd2d622e17be88cf8539b8b8072c9dc9a16008e93927d7d3c4782402f1498bb3c308a17fbe8bfeac0b2bc2049bf13d3ee40eaf6489f564713450e04bad5d410cbd39fc81f9e4143fd09f4b4c96eda25f0ba2b33bb9705c628fe74ea910cdd40bef9fca4a60138b173324c0101a6d080112210320bbac09528e19c55b0f89cb37ab265e7e856b1a8c388780322dbbfd194b52ba1a46304402204270ce88a565d040aa1b57451ebfc901ff9c27fc133b401c547ecec22018fd8302202a4051a972299e3f7f3dbc7ff33fd5a76646e40ae4cbad82104a3b6628d3545d20a08d0630efa2abc6be9ed18d1e3a22313668747663424e53454137665a6841644c4a706844775152514a614870794854703aff010a08756e667265657a65124c20010a481205636f696e731a03627479208084af5f2a22313271796f6361794e46374c7636433971573461767873324537553431664b5366763209466978416d6f756e743a06081410c0843d1a6d0801122102504fa1c28caaf1d5a20fefb87c50a49724ff401043420cb3ba271997eb5a43871a46304402202db898956e0bdedc84b6a9025f45dd41ae1a90d4282a634eece78aeb67023e2102202aa8619df941c309eccfa7c3442b0e687ed45ab6c13f2488e3dc0737b76786d020a08d0628f3bac9ea0530bbc698f894f3c7b9293a2231355973714175586545585648676d36525678346f4a6141416e687477716e75334858ffffbff9016220bf932ca9e274b04307b469e5fa647c5fb0a504ed2bbccff3ad1195051e77f53d68a30112db0508021a5e0802125a0a2b10e0c7adfc941a2222313271796f6361794e46374c7636433971573461767873324537553431664b536676122b10c0baa7fc941a2222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a9f010871129a010a70313271796f6361794e46374c7636433971573461767873324537553431664b5366763a3078386434663130653666313762616533636538663764303239616263623461393839616563333333386238386662333537656165663035613265326465373930343a30303030303037303838100218012222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a620805125e0a2d1080a098e0d79ea5352222313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d1080eacdced99ea5352222313668747663424e53454137665a6841644c4a706844775152514a614870794854701a870108081282010a22313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d1880d4ebddd4e0e111222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a2d18809ea1ccd6e0e111222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a620805125e0a2d1080eacdced99ea5352222313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d1080b0d4ddda9ea5352222313668747663424e53454137665a6841644c4a706844775152514a614870794854701a82010808127e0a22313668747663424e53454137665a6841644c4a706844775152514a61487079485470122b1880cca486b5012222314251585336547861595947356d41446157696a344178685a5a55547077393561351a2b188092ab95b6012222314251585336547861595947356d41446157696a344178685a5a555470773935613512bc0308021a5c080212580a2a10a0e0cbbe01222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a74122a1080d3c5be01222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a860108091281010a2231355973714175586545585648676d36525678346f4a6141416e687477716e753348122a1080cab5ee01222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a2f1080c6868f01188084af5f222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741ad00108d10f12ca0112c7010a4e6d61766c2d756e667265657a652d3138613864343833643635356362373864366261643337666363363737326265356638373563333030356531303839313537393233623562613234353363313510fcb9c9ea051a05636f696e732203627479288084af5f322231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a743a22313271796f6361794e46374c7636433971573461767873324537553431664b536676408084af5f4a09466978416d6f756e745206081410c0843d"
	index := int64(1)
	env, err := util.GenEnv(blockByte, index)
	assert.Nil(t, err)

	records, err := convert.ConvertTx(env, db.OpAdd)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(records))
	assert.Equal(t, "unfreeze_tx/unfreeze/0x18a8d483d655cb78d6bad37fcc6772be5f875c3005e1089157923b5ba2453c15", records[0].Key())
	assert.Equal(t, `{"block":{"height":163,"ts":1565678844,"block_hash":"0x78171e8fc1573773a0f66c5a0a596b9e4b356c3ba23fbec2c20761305c20b853","index":1,"send":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","tx_hash":"0x18a8d483d655cb78d6bad37fcc6772be5f875c3005e1089157923b5ba2453c15","height_index":16300001},"creator":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","beneficiary":"12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv","unfreeze_id":"mavl-unfreeze-18a8d483d655cb78d6bad37fcc6772be5f875c3005e1089157923b5ba2453c15","success":true,"action_type":"create","create":{"start_time":1565678844,"asset_exec":"coins","asset_symbol":"bty","total_count":200000000,"means":"FixAmount","fix_amount":{"period":20,"amount":1000000}}}`, string(records[0].Value()))
}

func convertWithdraw(t *testing.T, convert db.ExecConvert) {
	blockByte := "0ab3071220cdd7e3f86d536a2c01fcd89a0573ea8c189e2aae471dd1ad328c75766c33b0f91a200e831b568af7adf8e1ae58b005faa9baea5d48ab95d23e9ad7a3097a20801c782220622d1824757ddeab8cf745cbd5dbecaf9fd97783cd7f10de0cce26b21a392c1428aa023092bac9ea053a9a040a067469636b657412ed02501022e80208ffffaff8011080cab5ee011a70313271796f6361794e46374c7636433971573461767873324537553431664b5366763a3078386434663130653666313762616533636538663764303239616263623461393839616563333333386238386662333537656165663035613265326465373930343a303030303030323530342242307837373334383238636166386338343435663333396531303436343564623931356462343063656562353430393137313538396638373661303236353630613138322059f36aefe55642cbb928692760749aa248892c55264b03d19a8471a7e5e6983b3a8101ccf8d597e592b664e6910c50765912d67609e62162207d9a52f97aa79ffed3a3f9534d41d52209cd407bcbab92bbb9be1c3057dbe835c3d1991542907388115504f6150e7d6ccf318ccfbaafcd1cbca8710243535ecce77d2756c3b08babd71a62054ac9aedb8d8f0e4316c0a2a3d34732968d273b28514f48e4d09f5699504a421a6e080112210320bbac09528e19c55b0f89cb37ab265e7e856b1a8c388780322dbbfd194b52ba1a4730450221008234ea5aff87096bfc5b9fa52e3e7770b0a2a51dfb2ab97e27a02bfccb28a4be02207ad0733075c22814ae4321020ea1fb3349b8a42e9621b0428d0ff9e91ef042c920a08d0630c7bb9cd1c6d7b6862a3a22313668747663424e53454137665a6841644c4a706844775152514a614870794854703af9010a08756e667265657a651246200212420a40313861386434383364363535636237386436626164333766636336373732626535663837356333303035653130383931353739323362356261323435336331351a6d080112210320bbac09528e19c55b0f89cb37ab265e7e856b1a8c388780322dbbfd194b52ba1a463044022049a1068d628ebdb6be6a3ccd998156acc7310c91bb6f449eff1054109965f7db02201474a7f66a0dcb736464b19c356f59f53d17ac9ec5cd0e44c0a9ecf3e1d5fc7320a08d062889bbc9ea0530ac90eef9be958a92503a2231355973714175586545585648676d36525678346f4a6141416e687477716e75334858ffffaff8016220ab82d5bb840e33fd0f5faefbcc8142deba5e1fe545d2c748997b1298022c3cdb68aa0212db0508021a5e0802125a0a2b1080cbf5f5941a2222313271796f6361794e46374c7636433971573461767873324537553431664b536676122b10e0bdeff5941a2222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a9f010871129a010a70313271796f6361794e46374c7636433971573461767873324537553431664b5366763a3078386434663130653666313762616533636538663764303239616263623461393839616563333333386238386662333537656165663035613265326465373930343a30303030303032353034100218012222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a620805125e0a2d108090cd8aeaa1a5352222313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d1080da82f9eba1a5352222313668747663424e53454137665a6841644c4a706844775152514a614870794854701a870108081282010a22313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d1880daac98d0e2e111222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a2d1880a4e286d2e2e111222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a620805125e0a2d1080da82f9eba1a5352222313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d1080a08988eda1a5352222313668747663424e53454137665a6841644c4a706844775152514a614870794854701a82010808127e0a22313668747663424e53454137665a6841644c4a706844775152514a61487079485470122b1880b698f6cb022222314251585336547861595947356d41446157696a344178685a5a55547077393561351a2b1880fc9e85cd022222314251585336547861595947356d41446157696a344178685a5a555470773935613512870608021a5e0802125a0a2b10e0bdeff5941a2222313271796f6361794e46374c7636433971573461767873324537553431664b536676122b10c0b0e9f5941a2222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a8b0108061286010a2231355973714175586545585648676d36525678346f4a6141416e687477716e753348122f1080c6868f01188084af5f222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a2f1080c6868f011880fbb45e222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a78080612740a2231355973714175586545585648676d36525678346f4a6141416e687477716e75334812242222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a281080897a2222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a9a0308d20f1294030ac7010a4e6d61766c2d756e667265657a652d3138613864343833643635356362373864366261643337666363363737326265356638373563333030356531303839313537393233623562613234353363313510fcb9c9ea051a05636f696e732203627479288084af5f322231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a743a22313271796f6361794e46374c7636433971573461767873324537553431664b536676408084af5f4a09466978416d6f756e745206081410c0843d12c7010a4e6d61766c2d756e667265657a652d3138613864343833643635356362373864366261643337666363363737326265356638373563333030356531303839313537393233623562613234353363313510fcb9c9ea051a05636f696e732203627479288084af5f322231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a743a22313271796f6361794e46374c7636433971573461767873324537553431664b5366764080fbb45e4a09466978416d6f756e745206081410c0843d"
	index := int64(1)
	env, err := util.GenEnv(blockByte, index)
	assert.Nil(t, err)

	records, err := convert.ConvertTx(env, db.OpAdd)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(records))
	assert.Equal(t, "unfreeze_tx/unfreeze/0x53a769f8b2a29b015346d0a6ac8c76224276e74f119f755f45bc52e4394d44f0", records[0].Key())
	assert.Equal(t, `{"block":{"height":298,"ts":1565678866,"block_hash":"0x0cb8d302eb1f21fefd900de04877a4712f1125d2370106480edbbaefbb188412","index":1,"send":"12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv","tx_hash":"0x53a769f8b2a29b015346d0a6ac8c76224276e74f119f755f45bc52e4394d44f0","height_index":29800001},"creator":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","beneficiary":"12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv","unfreeze_id":"mavl-unfreeze-18a8d483d655cb78d6bad37fcc6772be5f875c3005e1089157923b5ba2453c15","success":true,"action_type":"withdraw","create":{"start_time":1565678844,"asset_exec":"coins","asset_symbol":"bty","total_count":200000000,"means":"FixAmount"},"withdraw":{"amount":2000000}}`, string(records[0].Value()))
}

func convertTerminate(t *testing.T, convert db.ExecConvert) {
	blockByte := "0ab30712201b66de4a75f103846ae9231c01ca90758ee76d10b8b389568cec0a7234277dbd1a20e08bffe1927773e65234e8c4dad6d1fdb3f742deed315950e8b5b906d7bef0642220eb4947d2aa13080eac0e475d966d74d35d170d0b195ea5aed006ddaa587d59ff28b0023094bac9ea053a9a040a067469636b657412ed02501022e80208ffffaff8011080cab5ee011a70313271796f6361794e46374c7636433971573461767873324537553431664b5366763a3078386434663130653666313762616533636538663764303239616263623461393839616563333333386238386662333537656165663035613265326465373930343a3030303030303239383022423078373733343832386361663863383434356633333965313034363435646239313564623430636565623534303931373135383966383736613032363536306131383220d7877136e62137ec9bb498e60d1881170a1558fbb76031310e05e682f63b13203a81017e1759e42b3658e52cf2e07e3d11f682ae2d7f2159c62e383d0406f02d76a642915993c19499cd471a34124d0f9a001804288daed58cfda8f347f9bd749851df04f5a23c293881a47a98d4118a3eda1a85c50597f61b808d5d842358c31d278ac349c09b83c800dcbbb17311b7997bab1a067a4b8937179b61a9b3468a2fa70e011a6e080112210320bbac09528e19c55b0f89cb37ab265e7e856b1a8c388780322dbbfd194b52ba1a473045022100c4cdb2e7ddf0f527665ade8db5bb9a634bedf30d53e9c1977e666ca143ea105d022019b9421a997e0318436b1663b42cbb0487c248fa65373bed5b5524fd4d4fe45420a08d0630cdb7d6e287ddb181093a22313668747663424e53454137665a6841644c4a706844775152514a614870794854703af9010a08756e667265657a65124620031a420a40313861386434383364363535636237386436626164333766636336373732626535663837356333303035653130383931353739323362356261323435336331351a6d0801122102504fa1c28caaf1d5a20fefb87c50a49724ff401043420cb3ba271997eb5a43871a46304402207f2201f7dca3ddb8a95b66fb058868e9c55d156e2e0086b689d30f2a52d022e402204c240850ddfcd841aa23ce3b9f6977f78183b1c1e9ffe8664d128d084a83dd9420a08d06288bbbc9ea0530a5d6def3a8a780b4663a2231355973714175586545585648676d36525678346f4a6141416e687477716e75334858ffffaff8016220c0b59fd6a458df2aa9adcfe4a86b75a74d8a903418c7d891c960748b154eab5968b00212db0508021a5e0802125a0a2b10a0eecaf5941a2222313271796f6361794e46374c7636433971573461767873324537553431664b536676122b1080e1c4f5941a2222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a9f010871129a010a70313271796f6361794e46374c7636433971573461767873324537553431664b5366763a3078386434663130653666313762616533636538663764303239616263623461393839616563333333386238386662333537656165663035613265326465373930343a30303030303032393830100218012222313271796f6361794e46374c7636433971573461767873324537553431664b5366761a620805125e0a2d1080f0b5fbfba1a5352222313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d1080baebe9fda1a5352222313668747663424e53454137665a6841644c4a706844775152514a614870794854701a870108081282010a22313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d188096eeaedbe2e111222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a2d1880e0a39ddde2e111222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a620805125e0a2d1080baebe9fda1a5352222313668747663424e53454137665a6841644c4a706844775152514a61487079485470122d108080f2f8fea1a5352222313668747663424e53454137665a6841644c4a706844775152514a614870794854701a82010808127e0a22313668747663424e53454137665a6841644c4a706844775152514a61487079485470122b1880dabfd0d2022222314251585336547861595947356d41446157696a344178685a5a55547077393561351a2b1880a0c6dfd3022222314251585336547861595947356d41446157696a344178685a5a555470773935613512830508021a5c080212580a2a1080d3c5be01222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a74122a10e0c5bfbe01222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a8601080a1281010a2231355973714175586545585648676d36525678346f4a6141416e687477716e753348122f1080c6868f011880fbb45e222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a2a1080c1bbed01222231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a741a970308d30f1291030ac7010a4e6d61766c2d756e667265657a652d3138613864343833643635356362373864366261643337666363363737326265356638373563333030356531303839313537393233623562613234353363313510fcb9c9ea051a05636f696e732203627479288084af5f322231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a743a22313271796f6361794e46374c7636433971573461767873324537553431664b5366764080fbb45e4a09466978416d6f756e745206081410c0843d12c4010a4e6d61766c2d756e667265657a652d3138613864343833643635356362373864366261643337666363363737326265356638373563333030356531303839313537393233623562613234353363313510fcb9c9ea051a05636f696e732203627479288084af5f322231344b454b6259744b4b516d34774d7468534b394a344c61346e41696964476f7a743a22313271796f6361794e46374c7636433971573461767873324537553431664b5366764a09466978416d6f756e7460015206081410c0843d"
	index := int64(1)
	env, err := util.GenEnv(blockByte, index)
	assert.Nil(t, err)

	records, err := convert.ConvertTx(env, db.OpAdd)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(records))
	assert.Equal(t, "unfreeze_tx/unfreeze/0x84dc5ea1393370f70c67d20180e81188b7152767f5a9830f5d513a35130e8346", records[0].Key())
	assert.Equal(t, `{"block":{"height":304,"ts":1565678868,"block_hash":"0xcadc7c6d7eda4968144186c482acc4b397b505e04963ed20411746438f3e9cbc","index":1,"send":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","tx_hash":"0x84dc5ea1393370f70c67d20180e81188b7152767f5a9830f5d513a35130e8346","height_index":30400001},"creator":"14KEKbYtKKQm4wMthSK9J4La4nAiidGozt","beneficiary":"12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv","unfreeze_id":"mavl-unfreeze-18a8d483d655cb78d6bad37fcc6772be5f875c3005e1089157923b5ba2453c15","success":true,"action_type":"terminate","create":{"start_time":1565678844,"asset_exec":"coins","asset_symbol":"bty","total_count":200000000,"means":"FixAmount"},"terminate":{"amount_back":198000000,"amount_left":0}}`, string(records[0].Value()))
}
