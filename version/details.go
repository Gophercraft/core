package version

import (
	"encoding/hex"
)

type BuildInfo struct {
	MajorVersion    uint32
	MinorVersion    uint32
	BugfixVersion   uint32
	HotfixVersion   string
	WinAuthSeed     []byte
	Win64AuthSeed   []byte
	Mac64AuthSeed   []byte
	WinChecksumSeed []byte
	MacChecksumSeed []byte
}

// This data will be displayed in the JSON/protobuf realm listing.
var details = map[Build]*BuildInfo{}

// https://github.com/TrinityCore/TrinityCore/blob/master/sql/base/auth_database.sql#L527
func init() {
	addBuild(3368, 0, 5, 3, "", "", "", "", "", "")
	addBuild(3494, 0, 5, 5, "", "", "", "", "", "")
	addBuild(3592, 0, 6, 0, "", "", "", "", "", "")
	addBuild(3694, 0, 7, 0, "", "", "", "", "", "")
	addBuild(3702, 0, 7, 1, "", "", "", "", "", "")
	addBuild(3712, 0, 7, 6, "", "", "", "", "", "")
	addBuild(3734, 0, 8, 0, "", "", "", "", "", "")
	addBuild(3807, 0, 9, 0, "", "", "", "", "", "")
	addBuild(3810, 0, 9, 1, "", "", "", "", "", "")
	addBuild(3892, 0, 10, 0, "", "", "", "", "", "")
	addBuild(3925, 0, 11, 0, "", "", "", "", "", "")
	addBuild(3988, 0, 12, 0, "", "", "", "", "", "")
	addBuild(3980, 1, 0, 0, "", "", "", "", "", "")
	addBuild(3989, 1, 0, 1, "", "", "", "", "", "")
	addBuild(4044, 1, 1, 0, "", "", "", "", "", "")
	addBuild(4062, 1, 1, 1, "", "", "", "", "", "")
	addBuild(4125, 1, 1, 2, "", "", "", "", "", "")
	addBuild(4147, 1, 2, 0, "", "", "", "", "", "")
	addBuild(4150, 1, 2, 1, "", "", "", "", "", "")
	addBuild(4196, 1, 2, 2, "", "", "", "", "", "")
	addBuild(4211, 1, 2, 3, "", "", "", "", "", "")
	addBuild(4222, 1, 2, 4, "", "", "", "", "", "")
	addBuild(4284, 1, 3, 0, "", "", "", "", "", "")
	addBuild(4297, 1, 3, 1, "", "", "", "", "", "")
	addBuild(4341, 1, 4, 0, "", "", "", "", "", "")
	addBuild(4364, 1, 4, 1, "", "", "", "", "", "")
	addBuild(4375, 1, 4, 2, "", "", "", "", "", "")
	addBuild(4442, 1, 5, 0, "", "", "", "", "", "")
	addBuild(4449, 1, 5, 1, "", "", "", "", "", "")
	addBuild(4500, 1, 6, 0, "", "", "", "", "", "")
	addBuild(4544, 1, 6, 1, "", "", "", "", "", "")
	addBuild(4671, 1, 7, 0, "", "", "", "", "", "")
	addBuild(4695, 1, 7, 1, "", "", "", "", "", "")
	addBuild(4735, 1, 8, 0, "", "", "", "", "", "")
	addBuild(4769, 1, 8, 1, "", "", "", "", "", "")
	addBuild(4807, 1, 8, 3, "", "", "", "", "", "")
	addBuild(4878, 1, 8, 4, "", "", "", "", "", "")
	addBuild(4937, 1, 9, 0, "", "", "", "", "", "")
	addBuild(4983, 1, 9, 1, "", "", "", "", "", "")
	addBuild(4996, 1, 9, 2, "", "", "", "", "", "")
	addBuild(5059, 1, 9, 3, "", "", "", "", "", "")
	addBuild(5086, 1, 9, 4, "", "", "", "", "", "")
	addBuild(5195, 1, 10, 0, "", "", "", "", "", "")
	addBuild(5230, 1, 10, 1, "", "", "", "", "", "")
	addBuild(5302, 1, 10, 2, "", "", "", "", "", "")
	addBuild(5428, 1, 11, 0, "", "", "", "", "", "")
	addBuild(5462, 1, 11, 1, "", "", "", "", "", "")
	addBuild(5464, 1, 11, 2, "", "", "", "", "", "")
	addBuild(5595, 1, 12, 0, "", "", "", "", "", "")
	addBuild(5875, 1, 12, 1, "", "", "", "", "95EDB27C7823B363CBDDAB56A392E7CB73FCCA20", "8D173CC381961EEBABF336F5E6675B101BB513E5")
	addBuild(6005, 1, 12, 2, "", "", "", "", "", "")
	addBuild(6141, 1, 12, 3, "", "", "", "", "", "")
	addBuild(8606, 2, 4, 3, "", "", "", "", "319AFAA3F2559682F9FF658BE01456255F456FB1", "D8B0ECFE534BC1131E19BAD1D4C0E813EEE4994F")
	addBuild(9947, 3, 1, 3, "", "", "", "", "", "")
	addBuild(10505, 3, 2, 2, "a", "", "", "", "", "")
	addBuild(11159, 3, 3, 0, "a", "", "", "", "", "")
	addBuild(11403, 3, 3, 2, "", "", "", "", "", "")
	addBuild(11723, 3, 3, 3, "a", "", "", "", "", "")
	addBuild(12340, 3, 3, 5, "a", "", "", "", "CDCBBD5188315E6B4D19449D492DBCFAF156A347", "B706D13FF2F4018839729461E3F8A0E2B5FDC034")
	addBuild(13623, 4, 0, 6, "a", "", "", "", "", "")
	addBuild(13930, 3, 3, 5, "a", "", "", "", "", "")
	addBuild(14545, 4, 2, 2, "", "", "", "", "", "")
	addBuild(15595, 4, 3, 4, "", "", "", "", "", "")
	addBuild(19116, 6, 0, 3, "", "", "", "", "", "")
	addBuild(19243, 6, 0, 3, "", "", "", "", "", "")
	addBuild(19342, 6, 0, 3, "", "", "", "", "", "")
	addBuild(19702, 6, 1, 0, "", "", "", "", "", "")
	addBuild(19802, 6, 1, 2, "", "", "", "", "", "")
	addBuild(19831, 6, 1, 2, "", "", "", "", "", "")
	addBuild(19865, 6, 1, 2, "", "", "", "", "", "")
	addBuild(20182, 6, 2, 0, "a", "", "", "", "", "")
	addBuild(20201, 6, 2, 0, "", "", "", "", "", "")
	addBuild(20216, 6, 2, 0, "", "", "", "", "", "")
	addBuild(20253, 6, 2, 0, "", "", "", "", "", "")
	addBuild(20338, 6, 2, 0, "", "", "", "", "", "")
	addBuild(20444, 6, 2, 2, "", "", "", "", "", "")
	addBuild(20490, 6, 2, 2, "a", "", "", "", "", "")
	addBuild(20574, 6, 2, 2, "a", "", "", "", "", "")
	addBuild(20726, 6, 2, 3, "", "", "", "", "", "")
	addBuild(20779, 6, 2, 3, "", "", "", "", "", "")
	addBuild(20886, 6, 2, 3, "", "", "", "", "", "")
	addBuild(21355, 6, 2, 4, "", "", "", "", "", "")
	addBuild(21463, 6, 2, 4, "", "", "", "", "", "")
	addBuild(21742, 6, 2, 4, "", "", "", "", "", "")
	addBuild(22248, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22293, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22345, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22410, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22423, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22498, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22522, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22566, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22594, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22624, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22747, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22810, 7, 0, 3, "", "", "", "", "", "")
	addBuild(22900, 7, 1, 0, "", "", "", "", "", "")
	addBuild(22908, 7, 1, 0, "", "", "", "", "", "")
	addBuild(22950, 7, 1, 0, "", "", "", "", "", "")
	addBuild(22995, 7, 1, 0, "", "", "", "", "", "")
	addBuild(22996, 7, 1, 0, "", "", "", "", "", "")
	addBuild(23171, 7, 1, 0, "", "", "", "", "", "")
	addBuild(23222, 7, 1, 0, "", "", "", "", "", "")
	addBuild(23360, 7, 1, 5, "", "", "", "", "", "")
	addBuild(23420, 7, 1, 5, "", "", "", "", "", "")
	addBuild(23911, 7, 2, 0, "", "", "", "", "", "")
	addBuild(23937, 7, 2, 0, "", "", "", "", "", "")
	addBuild(24015, 7, 2, 0, "", "", "", "", "", "")
	addBuild(24330, 7, 2, 5, "", "", "", "", "", "")
	addBuild(24367, 7, 2, 5, "", "", "", "", "", "")
	addBuild(24415, 7, 2, 5, "", "", "", "", "", "")
	addBuild(24430, 7, 2, 5, "", "", "", "", "", "")
	addBuild(24461, 7, 2, 5, "", "", "", "", "", "")
	addBuild(24742, 7, 2, 5, "", "", "", "", "", "")
	addBuild(25549, 7, 3, 2, "", "FE594FC35E7F9AFF86D99D8A364AB297", "1252624ED8CBD6FAC7D33F5D67A535F3", "66FC5E09B8706126795F140308C8C1D8", "", "")
	addBuild(25996, 7, 3, 5, "", "23C59C5963CBEF5B728D13A50878DFCB", "C7FF932D6A2174A3D538CA7212136D2B", "210B970149D6F56CAC9BADF2AAC91E8E", "", "")
	addBuild(26124, 7, 3, 5, "", "F8C05AE372DECA1D6C81DA7A8D1C5C39", "46DF06D0147BA67BA49AF553435E093F", "C9CA997AB8EDE1C65465CB2920869C4E", "", "")
	addBuild(26365, 7, 3, 5, "", "2AAC82C80E829E2CA902D70CFA1A833A", "59A53F307288454B419B13E694DF503C", "DBE7F860276D6B400AAA86B35D51A417", "", "")
	addBuild(26654, 7, 3, 5, "", "FAC2D693E702B9EC9F750F17245696D8", "A752640E8B99FE5B57C1320BC492895A", "9234C1BD5E9687ADBD19F764F2E0E811", "", "")
	addBuild(26822, 7, 3, 5, "", "283E8D77ECF7060BE6347BE4EB99C7C7", "2B05F6D746C0C6CC7EF79450B309E595", "91003668C245D14ECD8DF094E065E06B", "", "")
	addBuild(26899, 7, 3, 5, "", "F462CD2FE4EA3EADF875308FDBB18C99", "3551EF0028B51E92170559BD25644B03", "8368EFC2021329110A16339D298200D4", "", "")
	addBuild(26972, 7, 3, 5, "", "797ECC19662DCBD5090A4481173F1D26", "6E212DEF6A0124A3D9AD07F5E322F7AE", "341CFEFE3D72ACA9A4407DC535DED66A", "", "")
	addBuild(28153, 8, 0, 1, "", "", "DD626517CC6D31932B479934CCDC0ABF", "", "", "")
	addBuild(30706, 8, 1, 5, "", "", "BB6D9866FE4A19A568015198783003FC", "", "", "")
	addBuild(30993, 8, 2, 0, "", "", "2BAD61655ABC2FC3D04893B536403A91", "", "", "")
	addBuild(31229, 8, 2, 0, "", "", "8A46F23670309F2AAE85C9A47276382B", "", "", "")
	addBuild(31429, 8, 2, 0, "", "", "7795A507AF9DC3525EFF724FEE17E70C", "", "", "")
	addBuild(31478, 8, 2, 0, "", "", "7973A8D54BDB8B798D9297B096E771EF", "", "", "")
	addBuild(32305, 8, 2, 5, "", "", "21F5A6FC7AD89FBF411FDA8B8738186A", "", "", "")
	addBuild(32494, 8, 2, 5, "", "", "58984ACE04919401835C61309A848F8A", "", "", "")
	addBuild(32580, 8, 2, 5, "", "", "87C2FAA0D7931BF016299025C0DDCA14", "", "", "")
	addBuild(32638, 8, 2, 5, "", "", "5D07ECE7D4A867DDDE615DAD22B76D4E", "", "", "")
	addBuild(32722, 8, 2, 5, "", "", "1A09BE1D38A122586B4931BECCEAD4AA", "", "", "")
	addBuild(32750, 8, 2, 5, "", "", "C5CB669F5A5B237D1355430877173207", "EF1F4E4D099EA2A81FD4C0DEBC1E7086", "", "")
	addBuild(32978, 8, 2, 5, "", "", "76AE2EA03E525D97F5688843F5489000", "1852C1F847E795D6EB45278CD433F339", "", "")
	addBuild(33369, 8, 3, 0, "", "", "5986AC18B04D3C403F56A0CF8C4F0A14", "F5A849C70A1054F07EA3AB833EBF6671", "", "")
	addBuild(33528, 8, 3, 0, "", "", "0ECE033CA9B11D92F7D2792C785B47DF", "", "", "")
	addBuild(33724, 8, 3, 0, "", "", "38F7BBCF284939DD20E8C64CDBF9FE77", "", "", "")
	addBuild(33775, 8, 3, 0, "", "", "B826300A8449ED0F6EF16EA747FA2D2E", "354D2DE619D124EE1398F76B0436FCFC", "", "")
	addBuild(33941, 8, 3, 0, "", "", "88AF1A36D2770D0A6CA086497096A889", "", "", "")
	addBuild(34220, 8, 3, 0, "", "", "B5E35B976C6BAF82505700E7D9666A2C", "", "", "")
	addBuild(34601, 8, 3, 0, "", "", "0D7DF38F725FABA4F009257799A10563", "", "", "")
	addBuild(34769, 8, 3, 0, "", "", "93F9B9AF6397E3E4EED94D36D16907D2", "", "", "")
	addBuild(34963, 8, 3, 0, "", "", "7BA50C879C5D04221423B02AC3603A11", "C5658A17E702163447BAAAE46D130A1B", "", "")
	addBuild(35249, 8, 3, 7, "", "", "C7B11F9AE9FF1409F5582902B3D10D1C", "", "", "")
	addBuild(35284, 8, 3, 7, "", "", "EA3818E7DCFD2009DBFC83EE3C1E4F1B", "A6201B0AC5A73D13AB2FDCC79BB252AF", "", "")
	addBuild(35435, 8, 3, 7, "", "", "BB397A92FE23740EA52FC2B5BA2EC8E0", "8FE657C14A46BCDB2CE6DA37E430450E", "", "")
	addBuild(35662, 8, 3, 7, "", "", "578BC94870C278CB6962F30E6DC203BB", "5966016C368ED9F7AAB603EE6703081C", "", "")
	addBuild(36753, 9, 0, 2, "", "", "386FDE8559B5EAD67B791B490B200B88", "", "", "")
	addBuild(36839, 9, 0, 2, "", "", "356EB4412B8EFCF72E3AF50D5181D529", "", "", "")
	addBuild(36949, 9, 0, 2, "", "", "51C074CD8A09A75384B9B44188C56981", "", "", "")
	addBuild(37142, 9, 0, 2, "", "", "5D9CFB3139F0D1B6C2B304261F9EABC9", "", "", "")
	addBuild(37176, 9, 0, 2, "", "", "3C725EA504EC3DAED143EB6FF3B48CDA", "", "", "")
	addBuild(37474, 9, 0, 2, "", "", "0DE685BBB0551086E7FBDC0B4BB06A5B", "024C9BE7E44237B7E81C6D42E678D433", "", "")
	addBuild(38134, 9, 0, 5, "", "", "32275ED0F13B357C28BDB0E611EF9E31", "", "", "")
	addBuild(38556, 9, 0, 5, "", "", "EC7D5AF64364AC3E7181F3FBA1B3A882", "", "", "")
	addBuild(39653, 9, 1, 0, "", "", "10D015AB1EEB91310428D9C57EE24632", "", "", "")
	addBuild(39804, 9, 1, 0, "", "", "E42D2BBA12ED260A76F9B1E477E19EA5", "", "", "")
	addBuild(40000, 9, 1, 0, "", "", "4CB1433AB637F09F03FBBD1B221B04B0", "", "", "")
	addBuild(40120, 9, 1, 0, "", "", "04F47EAEFD8BDEFE14AA0350EA336678", "853F0F2985CEAED46DF422583CD07A7C", "", "")
	addBuild(40443, 9, 1, 0, "", "", "8597BB43E8AB38C85504E8BFB72ABBF5", "", "", "")
	addBuild(40593, 9, 1, 0, "", "", "BA14570F2D62D5F61953394164A8DAE2", "", "", "")
	addBuild(40725, 9, 1, 0, "", "", "C1EBDBEB9BB2956EBCCEF7C9D27A1B3B", "", "", "")
	addBuild(40906, 9, 1, 5, "", "", "F5FC259C8635488AFE0D0CD023F361D4", "", "", "")
	addBuild(40944, 9, 1, 5, "", "", "368FC7FABAF487A8A049C11970657074", "", "", "")
	addBuild(40966, 9, 1, 5, "", "", "D90F47AF21F381D2D8F3763B994BAC88", "", "", "")
	addBuild(41031, 9, 1, 5, "", "", "019A0FACD6B0D6374B7BA69A5B677449", "", "", "")
	addBuild(41079, 9, 1, 5, "", "", "F8853CF823BC0BBE8A9677A762DFAEE1", "", "", "")
	addBuild(41288, 9, 1, 5, "", "", "871C0C9691DBC536EB24B68EC73FAD5B", "", "", "")
	addBuild(41323, 9, 1, 5, "", "", "E53D0DF1FAC1A59A1C8071B295A04A1D", "", "", "")
	addBuild(41359, 9, 1, 5, "", "", "5F8D7F2A690A4375A1B52A28D6D681FA", "", "", "")
	addBuild(41488, 9, 1, 5, "", "", "1BC91EC368705815F3F532B553DAD981", "", "", "")
	addBuild(41793, 9, 1, 5, "", "", "B3B47DA3B7615570742A55B96614EE1C", "", "", "")
	addBuild(42010, 9, 1, 5, "", "", "302970161D16417B5BE553CC530E011A", "", "", "")
	addBuild(42423, 9, 2, 0, "", "", "0B03614A7E94DD57548596BE420E9DC2", "", "", "")
	addBuild(42488, 9, 2, 0, "", "", "A78755E6928D83A271C5D1EE3CDB6F15", "", "", "")
	addBuild(42521, 9, 2, 0, "", "", "5FE6C12FC407C6B15B4A5D3B5B4A5D3B", "", "", "")
	addBuild(42538, 9, 2, 0, "", "", "71A7504BD53F8DE518F24265D37310AE", "", "", "")
	addBuild(42560, 9, 2, 0, "", "", "115FE8C38A8D67CA4664BB192E0F0DFE", "", "", "")
	addBuild(42614, 9, 2, 0, "", "", "772BE726FEEF42124255D2EA7973CA18", "", "", "")
	addBuild(42698, 9, 2, 0, "", "", "B4497B1CD11FC974C5FB09548AC27269", "", "", "")
	addBuild(42825, 9, 2, 0, "", "", "A14DA228C6A6AFF1DDBA51218939E557", "", "", "")
	addBuild(42852, 9, 2, 0, "", "", "DE9F9F0C3CC8FD54D3AFF99CEFFCE129", "", "", "")
	addBuild(42937, 9, 2, 0, "", "", "F5FC75E70874752C92846B3333920E63", "", "", "")
	addBuild(42979, 9, 2, 0, "", "", "E1DD38AE6450FC4D2AE4609233C59E54", "", "", "")
	addBuild(43114, 9, 2, 0, "", "", "F75C9380CCB24A48A24EEE52C1594A7E", "", "", "")
	addBuild(43206, 9, 2, 0, "", "", "DDE806532C7704FFB75F256DC5F1F3D9", "", "", "")
	addBuild(43340, 9, 2, 0, "", "", "70E46D2D888E84DF9316EA849B068CF4", "", "", "")
	addBuild(43345, 9, 2, 0, "", "", "D911ABFCDA030DEE8CAF4EE3F60DEE13", "", "", "")
	addBuild(43971, 9, 2, 5, "", "", "681CF99E61FB0005A5C7D31D0AAD1ED9", "", "", "")
	addBuild(44015, 9, 2, 5, "", "", "FCF0BDA7C98BFEF92AE6D8C39A217ABD", "", "", "")
	addBuild(44061, 9, 2, 5, "", "", "FD2B5C0B3293FE0E9CAA6EB0B7788119", "", "", "")
	addBuild(44127, 9, 2, 5, "", "", "787887CEC9FCC9B516481F60E4FC34A8", "", "", "")
	addBuild(44232, 9, 2, 5, "", "", "81F0A71DF7E9873BB3750022D64D33CF", "", "", "")
	addBuild(44325, 9, 2, 5, "", "", "138A7D524D268A7F9934C3D148E8F01B", "", "", "")
	addBuild(44730, 9, 2, 5, "", "", "FC0B18C47BB4C79F4300CA0FF3E5CAC7", "", "", "")
	addBuild(44908, 9, 2, 5, "", "", "BFFAEC40C9BCD591C7C959A9D5A8BA8C", "", "", "")
	addBuild(45114, 9, 2, 7, "", "", "D7AFE240BD00F06C30D0C2D16E54A8BE", "", "", "")
	addBuild(45161, 9, 2, 7, "", "", "74BD2E787A98B145B063BDA9A98F6CBD", "", "", "")
	addBuild(45338, 9, 2, 7, "", "", "5CE2094A41B61EDA9DF56378BC3B1DE0", "", "", "")
	addBuild(45745, 9, 2, 7, "", "", "0F6DC90161694D765A595A3AF603166B", "", "", "")
	addBuild(46479, 10, 0, 2, "", "", "CB9AF4D89B60A3ABA288D395D315D932", "", "", "")
	addBuild(46658, 10, 0, 2, "", "", "3F8EFB085428D75360E9EFE25CD8639A", "", "", "")
	addBuild(46689, 10, 0, 2, "", "", "D9A11D188D6AD60906F5467510EFD3AA", "", "", "")
	addBuild(46702, 10, 0, 2, "", "", "01B4D1688FF97DC9AAFCCD0A0B842C0B", "", "", "")
	addBuild(46741, 10, 0, 2, "", "", "4C0F4A7EC2098AF1FBA745848EC79A78", "", "", "")
	addBuild(46801, 10, 0, 2, "", "", "E6AC18D1EA5D36ABFFAE5EDED80630DF", "", "", "")
	addBuild(46879, 10, 0, 2, "", "", "EFEC43936051DD1A210633AF1F6B63DB", "", "", "")
	addBuild(46924, 10, 0, 2, "", "", "E6CE0B1A8119F069ECF1E7DBAA7BB2F8", "", "", "")
	addBuild(47067, 10, 0, 2, "", "", "63862CFCDEA6BD2BD7F740EB36B65657", "", "", "")
	addBuild(47187, 10, 0, 2, "", "", "711F8455C5000C237292E1E6E90631E1", "", "", "")
	addBuild(47213, 10, 0, 2, "", "", "23C50D88CEAC0A8696ADDE8DD244D4A2", "", "", "")
	addBuild(47631, 10, 0, 2, "", "", "F986AB91D0AEB20822EFB72F4256713C", "", "", "")
	addBuild(47777, 10, 0, 5, "", "", "A88C04915AB9E035A104E55C4DCF5F9F", "", "", "")
	addBuild(47799, 10, 0, 5, "", "", "7364EB093C23DB2CDC9513D5A7B4933E", "", "", "")
	addBuild(47825, 10, 0, 5, "", "", "82A3B94E5E727AF3A208B471FF2054C0", "", "", "")
	addBuild(47849, 10, 0, 5, "", "", "DD8BBE2087A28C0AF4984CBE23A1C707", "", "", "")
	addBuild(47871, 10, 0, 5, "", "", "8E4F7D30EE4982B02B3B3F8837C2C4F2", "", "", "")
	addBuild(47884, 10, 0, 5, "", "", "2B7A002BC359F2C31104BC2DE04302BF", "", "", "")
	addBuild(47936, 10, 0, 5, "", "", "833D30D8FBC43B3FAE99CD3898D70849", "", "", "")
	addBuild(47967, 10, 0, 5, "", "", "CFE225D0089E224D9E7541D3B5C26478", "", "", "")
	addBuild(48001, 10, 0, 5, "", "", "4B0260A37BD95B615E71048469E6D5BB", "", "", "")
	addBuild(48069, 10, 0, 5, "", "", "558CDF958FA082E95849779C7C6945E5", "", "", "")
	addBuild(48317, 10, 0, 5, "", "", "C096E37B45B43244E9C79916604DD4AF", "", "", "")
	addBuild(48397, 10, 0, 5, "", "", "64BA8779EAA97E6C57982B6B1A5B32E7", "", "", "")
	addBuild(48526, 10, 0, 5, "", "", "D5B7D3303A2A741E6913EE1AEB0BCB65", "", "", "")
	addBuild(48676, 10, 0, 7, "", "", "E059FB74DFF6438CC20C7F28900F64CA", "", "", "")
	addBuild(48749, 10, 0, 7, "", "", "92DBCCA0E33DFB8A17A2B6A39246B288", "", "", "")
	addBuild(48838, 10, 0, 7, "", "", "9E6F4E1E46EF228D2DE90F7BC48AAA96", "", "", "")
	addBuild(48865, 10, 0, 7, "", "", "4B774ABE7B34D6702571B4279A4B6A13", "", "", "")
	addBuild(48892, 10, 0, 7, "", "", "AA31BF27458321B03A1A346964DD7B9D", "", "", "")
	addBuild(48966, 10, 0, 7, "", "", "823142CA131FBB715FF55D4343E55C6D", "", "", "")
	addBuild(48999, 10, 0, 7, "", "", "79BA6FF0F9672EEF875F64155C8B62D4", "", "", "")
	addBuild(49267, 10, 0, 7, "", "", "EEE77EA5A216E0731ADBB41AEFB1DF31", "", "", "")
	addBuild(49318, 10, 1, 0, "", "", "AF439AEE62EE48B36C1725111E3D9BBF", "", "", "")
	addBuild(49343, 10, 0, 7, "", "", "301A0B4C0942B9B6F605B903AD6C1F60", "", "", "")
	addBuild(49407, 10, 1, 0, "", "", "6413820DC9885BB0693B37090CBB2F30", "", "", "")
	addBuild(49426, 10, 1, 0, "", "", "D85EDFBFE9A94A55E2B4510BE41C19B2", "", "", "")
	addBuild(49444, 10, 1, 0, "", "", "363B2B05285BDD8857419D2866316D3C", "", "", "")
	addBuild(49474, 10, 1, 0, "", "", "44A7D2B352EE3D098A3CB4C2F1065E37", "", "", "")
	addBuild(49570, 10, 1, 0, "", "", "B024DE67ACAEB9E8EE6DB38DC53E8281", "", "", "")
	addBuild(49679, 10, 1, 0, "", "", "9CE59B68D8797EBF00581F41138F4316", "", "", "")
	addBuild(49741, 10, 1, 0, "", "", "0EF181E2BB0E946CF3B7422ADEB6CD1A", "", "", "")
	addBuild(49801, 10, 1, 0, "", "", "0832179567B66CA85DBD5678B604C683", "", "", "")
	addBuild(49890, 10, 1, 0, "", "", "22A5B8A1EB797A64995F705B3DBCB14C", "", "", "")
	addBuild(50000, 10, 1, 0, "", "", "02F06FFA2296FD66384295DBFD5A4C91", "", "", "")
	addBuild(50401, 10, 1, 5, "", "", "3EEF52D902CCE81D16D0E255F0AA4938", "", "", "")
	addBuild(50438, 10, 1, 5, "", "", "0B5F68F06B129CB4C57702F6D30F260B", "", "", "")
	addBuild(50467, 10, 1, 5, "", "", "5E996B1CDCEE68432D6340138E68D1EB", "", "", "")
	addBuild(50469, 10, 1, 5, "", "", "1768CCB6589E16AB3BEFA9D608A393A2", "", "", "")
	addBuild(50504, 10, 1, 5, "", "", "7D5FD20C0B32C9AF5DD65433B391D49C", "", "", "")
	addBuild(50585, 10, 1, 5, "", "", "C4F7CC38A3B84935A485F7EDAD3E764B", "", "", "")
	addBuild(50622, 10, 1, 5, "", "", "D23A26FD75FD9A6073EB7060AA28E6A7", "", "", "")
	addBuild(50747, 10, 1, 5, "", "", "2D3C386A9C45C27304ED3A3C6EB3F7C8", "", "", "")
	addBuild(50791, 10, 1, 5, "", "", "0BE7D0BB07EF37C25CBC682409091EA0", "", "", "")
	addBuild(51130, 10, 1, 5, "", "", "44CD2C91E4F0655DA387483726CE4035", "", "", "")
	addBuild(51187, 10, 1, 7, "", "", "74E2055D3965269447B5CB1B31FC71C6", "", "", "")
	addBuild(51237, 10, 1, 7, "", "", "C8660A21B766646FBD67F481CFCF55C3", "", "", "")
	addBuild(51261, 10, 1, 7, "", "", "1BEBB57AE450331E9F8C301AA7876FAB", "", "", "")
	addBuild(51313, 10, 1, 7, "", "", "35419ED0AB16735CF720858F45DC300C", "", "", "")
	addBuild(51421, 10, 1, 7, "", "", "45E24D6F3335269787DF2B2063939002", "", "", "")
	addBuild(51485, 10, 1, 7, "", "", "EC549E1D0A5DD85C03E7A9D93B7DC6D1", "", "", "")
	addBuild(51536, 10, 1, 7, "", "", "570EEA10A8EC169C3FF9621D1B635BB4", "", "", "")
	addBuild(51754, 10, 1, 7, "", "", "BED5A861C071AB41FEF6087E0C37BB1A", "", "", "")
	addBuild(51886, 10, 1, 7, "", "", "09CF8919FD2EABDEAEBC0C810F53B511", "", "", "")
	addBuild(51972, 10, 1, 7, "", "", "444DC7EF3544B6670C18884DADA00428", "", "", "")
	addBuild(52038, 10, 2, 0, "", "", "A8EF004ADED8A3AFF5A67D2BB8D95795", "", "", "")
	addBuild(52068, 10, 2, 0, "", "", "A44F842BACCC7EE8E2975FAF01F12474", "", "", "")
	addBuild(52095, 10, 2, 0, "", "", "BA36382887D16D274EA9149695F0C9C8", "", "", "")
	addBuild(52106, 10, 2, 0, "", "", "95F43869B7D881212CBC1690B8F393ED", "", "", "")
	addBuild(52129, 10, 2, 0, "", "", "02DD842F2A7162EEB8FD5B9D325606F8", "", "", "")
	addBuild(52148, 10, 2, 0, "", "", "8A969717C8CDC6E7FF4C54D5CB00C224", "", "", "")
	addBuild(52188, 10, 2, 0, "", "", "977DF9993E94855DED5E328BA7A2F21F", "", "", "")
	addBuild(52301, 10, 2, 0, "", "", "821AA3BB237B400B82F44970250539AA", "", "", "")
	addBuild(52393, 10, 2, 0, "", "", "B013ED23B7EF51B29A45594D9BBB0D03", "", "", "")
	addBuild(52485, 10, 2, 0, "", "", "5805CEB4650730AE489258DD30E34441", "", "", "")
	addBuild(52545, 10, 2, 0, "", "", "FB52179A8355A46EDBFBDC8E8E5CDAFD", "", "", "")
	addBuild(52607, 10, 2, 0, "", "", "8F002E4AADCAEABB08ABC2880B31AD60", "", "", "")
	addBuild(52649, 10, 2, 0, "", "", "D0B779FBECEBC1ED5A85D83F03C8A75B", "", "", "")
	addBuild(52808, 10, 2, 0, "", "", "6276712B6C8AEA21CD5D94D52FEE70EE", "", "", "")
	addBuild(52902, 10, 2, 5, "", "", "D4F0A24CDF165628538C1C387A326AF3", "", "", "")
	addBuild(52968, 10, 2, 5, "", "", "2D247FD440C44D4F1BF80B075B8720F2", "", "", "")
	addBuild(52983, 10, 2, 5, "", "", "B1E5ADA5FDD06C9AB5E5D8A6983324AC", "", "", "")
	addBuild(53007, 10, 2, 5, "", "", "A21AFB4D381C56AF471D994258C0EEF5", "", "", "")
	addBuild(53040, 10, 2, 5, "", "", "2F1283BF7B7F307B70DBBD75CC42D7C3", "", "", "")
	addBuild(53104, 10, 2, 5, "", "", "DBD79EC8DF044B53C78931B985CAB406", "", "", "")
	addBuild(53162, 10, 2, 5, "", "", "8A67511FBF8984EEE2B630F7CB23376A", "", "", "")
	addBuild(53212, 10, 2, 5, "", "", "08761EFF2F9B639364B9A9FBFFFFB949", "", "", "")
	addBuild(53262, 10, 2, 5, "", "", "614A72D53126348A4927EC0F53FD2B7A", "", "", "")
	addBuild(53441, 10, 2, 5, "", "", "BFDD7D0FE87D5F75E6DEB4F5C99D7C99", "", "", "")
	addBuild(53584, 10, 2, 5, "", "", "CDD7A93659A03460B5A6CE1C4ACE5554", "", "", "")
	addBuild(53840, 10, 2, 6, "", "", "AC97D745C60DD3DC5F973E55C0E3649E", "", "", "")
	addBuild(53877, 10, 2, 6, "", "", "16320F95B63846A2276E1D2612C34AD4", "", "", "")
	addBuild(53913, 10, 2, 6, "", "", "475680680B2192EBCF6744D14F755199", "", "", "")
	addBuild(53989, 10, 2, 6, "", "", "3AEB90ACB9E18B88BA1021F52D51B857", "", "", "")
	addBuild(54070, 10, 2, 6, "", "", "FDA08264B7587250CF78F9B960218169", "", "", "")
	addBuild(54205, 10, 2, 6, "", "", "589D59EB3F0D6D77C2175D9302F78FDD", "", "", "")
}

func addBuild(build Build, major, minor, bugfix uint32, hotfix string, winAuthSeed, win64AuthSeed, mac64AuthSeed, winChecksumSeed, macChecksumSeed string) {
	bi := &BuildInfo{}
	bi.MajorVersion = major
	bi.MinorVersion = minor
	bi.BugfixVersion = bugfix
	bi.HotfixVersion = hotfix

	var err error
	if winAuthSeed != "" {
		bi.WinAuthSeed, err = hex.DecodeString(winAuthSeed)
		if err != nil {
			panic(err)
		}
	}
	if win64AuthSeed != "" {
		bi.Win64AuthSeed, err = hex.DecodeString(win64AuthSeed)
		if err != nil {
			panic(err)
		}
	}

	if mac64AuthSeed != "" {
		bi.Mac64AuthSeed, err = hex.DecodeString(mac64AuthSeed)
		if err != nil {
			panic(err)
		}
	}

	if winChecksumSeed != "" {
		bi.WinChecksumSeed, err = hex.DecodeString(winChecksumSeed)
		if err != nil {
			panic(err)
		}
	}

	if macChecksumSeed != "" {
		bi.MacChecksumSeed, err = hex.DecodeString(macChecksumSeed)
		if err != nil {
			panic(err)
		}
	}

	details[build] = bi
}

// (legacy auth protocol) Returns a 3-byte field describing the version data. For instance, version 3.3.5 would return []byte{ 3, 3, 5 }
func (b Build) VersionData() []byte {
	d := b.ClientVersion()
	return []byte{
		uint8(d.Major),
		uint8(d.Minor),
		uint8(d.Revision),
	}
}

type ClientVersion struct {
	Build    Build
	Major    uint32
	Minor    uint32
	Revision uint32
}

func (b Build) ClientVersion() *ClientVersion {
	info := b.BuildInfo()
	if info == nil {
		return nil
	}

	ver := new(ClientVersion)
	ver.Build = b
	ver.Major = info.MajorVersion
	ver.Minor = info.MinorVersion
	ver.Revision = info.BugfixVersion

	return ver
}

func (b Build) BuildInfo() *BuildInfo {
	return details[b]
}
