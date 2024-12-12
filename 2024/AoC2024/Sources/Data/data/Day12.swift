import Tasks
import Framework

extension Data {
    public var day12: Day12 { Day12() }
    
    public struct Day12 {
        public var example1: String { return """
AAAA
BBCD
BBCC
EEEC
"""     }
        public var example2: String { return """
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
"""     }
        public var example3: String { return """
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
"""     }
        public var example4: String { return """
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
"""     }
        public var example5: String { return """
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
"""     }
        public var data: String { return """
YYYYYMMMMMMUUUUUUUUNNNNUTTTTTAASAXXXXXHHEEHHEEEEERYYYYYYYYYYHHHHHHHHHHHHHHHHFFFFFFFFFFAAAAAAAAAAAAAOOOOOOOOYYYYYYYYYYYYYYYYVVVVWWWWWWWZZZZZZ
YYYYYMMMMMMUMUUUUUUUUUUUTTAAALAAAAXXXHHFHHHEEEEEERRRRYYYYYYQHHHHHHHHHHHHHHHHHFFFFFFFFFAAAAAAMMMMMMMMMMMOOOOOYYYYYYYYYYYYYYYVVVWWWWWWKWZZZZZZ
YYYYYMMMMMMMMMMUUUUUUUUUTAAAAAAAAXXXXHHHHHHEEEEEERRRYYRYYYYQQHHHHHHHHHHHHHHHHFFFFFFAAAAAAAAAMMMMFMMMMOOOOOOOYYYYYYYYYYYYYYVVVVWWWTWKKWZKKZZZ
YYYYYYMMMMMMMUUUUUUUUUUUUAAAAAAAAHHHHHHHHHHHEEERRRRRYYRQQQQQQHHHHHHHHHHHHHHHHFFFFFAAAAAAAAAAAMMMMMMMMOOOOOOOYYYYYYYYYYYNDDDDDDTTTTTKKKKKKZZZ
YYYYMMMMMMMMMUUUUUUUUUUUUAAAAAAAAHHHHHHHHHHHMEERRRRRRRRRQQQQQQQHHHHHHHHHHHHHFFFFFFAAAAAAAAAAAMMMMMMMMOOOOOOOOYYYYYYYYYYNDDDDDDDDTTKKKKKKKKKK
YYYMMMMMMMMMMMMUUUUUUUUUUAAAAAAAAHHHHHHHHHKHMEERRRRRRRRQQQQQQQQQHHHHHHHHHHHHHFFFFFFFAAAAAAMMMMMMMMMMMOOOOOOOYYKYYYYYYYYNNDDDDDDTTKKKKKKKKKKM
YYYYMMMMMMMMMMMUUUUUUAUUUAAAAAAAHHHHHHHHMMMMMEERRIRRRQQQQQQQQQQQIHHHHHHHHHHHHHFFTTTFAVVAUUUUMMMMMMMMMMOOOOOOOYYYYYCYYDDDDDDDDDDTTTTTKKKKKKKM
YYYMMMMMMMMMMMMMUPPUUAAAAAAAAAAAAHHYYLQQMMMMMEEIIIRRRRQQQQQQQQQQHHHHHHHHHHLLHFFTTTTTRRVVVVUUMMMMMMMMMMMOOOOOOYYYYYFFVVDDDDDDDDDKKKKKKKKKKKKK
YYYMMMMMMMMMMMMMUPPPAAAAAAAAAABBAHHHHLQLLMLLMEIIIIIIRRQQQQQQQQQQQXHHHHHHHHKLHFFTTTTTRRVVVVVUMMMMMMMMPOOOOOOOYYYFYFFFFFFDDDDDDDDDKKKKKKKKKKKK
YYYYMMMMMMMMMMOOUPPPPPNNAAAAABBOHHHHHLLLLLLLLIIIIIINIRRQQQQQQQQQQXHHHHHHHHKKHFFOOTTVVVVVVVVMMMMMMMMMPOOOOOOOOOFFFFFFFFFDDDDDDDDDWWKKKKKKKKKK
YYYYYMMMMMMMMMMUUPPPPNNNAAABBBBOHHHHLLLLLLLLIIIIIIIIIRQQQQQQQQQQQXXHHHHHHHKKHFFKOTTVVVVVVVVVOMMMMMZMPOOOOOOOOOFFFFFFFFFFDDDDDDDDDWWKKKKKKKKK
YEEMMMMMMMZMMMRRRPPNNNNNNAABBBBOOHHLLLLLLLLLLLIIIAIIIRRRQRQQQQQQQQXXHHHKCHKKKKKKOOOVVVVVVVVVVSMMMMMSOOOOOOOOOOFFFFFFFFFDDDDDDDDDWKKKKKKKKKKK
EYEEMMMMMMMMMMRNNPNNNNNNAAAABBBOOLLLLLLLLLLLLIIIIIIIRRRRRRRRQQSQCSXXXXXKKKKKKKKKOOOVVVVVVSVVVSMMMMSSSOOOOOOOOOFFFFFFFFFFDDDDDDDDWKKKKKKKKWKK
EEEEEQMMMMMGMRRTNNNNNNNNAAAAPAJAALLLLLLLLLLLLIIIIIIIRRRRRRRRQSSSSSSXXKKKKKKKKKKKKVVVVVVVRSVVSSMMSSSSSOOOOOOOCCFFFFFFFFFFDDDDDDDDWWKWWWWWKWWW
EEEEEEEMNNMGTTTTTNNNNNNNAAAAAAAAALLLLLLLLLLLLIIIIIIIIRRRRRRRRSSSSSSXXKKKKKKKKKZZVVVVVVVVRSSSSSSSSSSSLOOOOLLOOLZFFFFFFFFFFFFDDDDDWWWWWWWWWWWW
EEEEEEEMMEGGTTTTNNNNNNNWNAAAAAAJALLLLLLLLLLLIIIIIIIIBZZZZRRRSSSSSSSSSKKKKKKKKKZZZVVVVVRRRRRSSSSSSSLLLOLLLLLLLLPFFFFFFPFPFFDDDDDDDWWWWWWWWWWW
EEEEEEEEEEEEVTTTNNNNNNNNNAAAAAAAPLLLWLLWWLLLLIIIIIIIBZZZZRRSSSSSSSSSAKKKKKKKKKZZZZVVVRRRRRRSSSSSSLLLLLLLLLLLLPPFFFFPPPPPFDDDDDDDDVWWWWWWWWWW
EEEEEEEEEEEEETTTTTNNNNNNNAAAAAPPPPPLWWWWWLLLLLLZIIIBBBZZZZRSSSSSSSSSSKKKKKKKZZZZZZZVVRRVRRRRSSTSSLSLLLLLLLLLLLPPPPPPPPPPPDDDDDDDVVWWWWWWWWWW
EEEEEEEEEEEEEEETTTNNNNNNNAAAAPPPPPPPPWWTWLLLLLLZZIIBBBZZZRRRSSSSSSWSKKKKKZZKZZZZZZZZZTRRRRRRTSTTSSSNNLLLLLLLLPPPPPBPPPPVDDDDDDDDDVVVWWWWWWWW
EEEEENEEEEEEEEETTNNNNNNNNAAAAPPPPPPPPWAWWLLWLLZZZZIZZZZZZZRRSSSSSSSHHKKKKKZZZZZZZZZZZTTRRRRRTTTSSSSNNNNLLLLLLPPPPPPPPPBVVVDDDDDVVVVVWWWWWWWW
EEEENNNEEEEEEEENNNNNNNNNNAGGDGGGGPPPPWWWWWWWLLZZZZZZZZZZZZRRSSSSSSSHHKKKKKQZZZZZZZZZZTTTTRRTTTTTSSSNNNNLLLLLLLLPPPPPPPVVVVVDDVDVVVVVWWWWWWWW
EEEENNNNNNEEPPEPNNNNNNNMGGGGDDGGGGPPPWWWWWWWWWZZZZZZZZZZZZRRSSSSSSSHKKKKNNNPPZZZZZZZTTTTRRRTTTNNSNSNNNNLNLLLLLLPPPPPPPPVVVVVVVVVVVVWWWWWWWWW
EEEENNNNNNNNPPPPNNNNNNNMGGGGGGGGGWWPPPWWWWWWWZZZZZZZZZZZZZZRRRSSSSSHHHHHHNNNZZZZZZZZTTTTTTTTNNNNNNNNNNNNNNLLLLPPPPPPEPEVVVVVVVVVVVVVWWWWWWWW
EEENNNNNHNNNPNNNNNNNNNNNNGGWGGGGGGWWWWWWWWWWWZZZZZZZZZZZZZZRSSSSSSSSHHHHHNNNNNZZZZZTTTTTTTTTTNNNNNNNNNNNNNLLLLPPPPPPEPEEEVVEVVVVVVVVWWWWWUUW
ENNNNNNNNNNNNNJJNVVVVVNENGGGGGGGGGWWWWWWWWWSSRZRZZZZZZZZZZZZSSSSSSSHHHHQHHNNNZZZTTTTTTTTTTTTTNNNNNNNNNNNNNNNNLPPPPPPEEEEEEEEEEVVVVVWWVWWWUWW
NDNNNNNNNNNNNNGGGVVVVNNVVGGGGGGGGGWWWWWWWWWWSRRRZZZZZZZZZZZZZSSSSSQHHHQQQQNNNNZTTTTTTTTTTTTTNNNNNNNNNNNNNNNNLLLLPPEEEEEEEEEEEEVVVVVVVVWWWUUU
NNNNNNNNNNNNGGGGGGVVVVVVGGGGSGBBGGGWWWWWWWRRRRRNNNZZZZZZZZZZZQQQSQQQHQQQQNNNNNTTTTTTTTTTTTNTNNNNNNNNNNNNNKLLLLLLPPEKEEEEEEEEEEEEVVVVVWWWWWWU
NNNNNNNNNNNNGGGGGGGGVVVVGGGGBBBBGGGGGWWWWVRRRRRRRNZZZZZZZZZZZQQQSQQQQQQQQNNNNNTTTTTTTTTTTTNNNNNNNNNNNNNNNKLLLLLKKKKKEEEEEEEEEEEEVVVVVWWWWWUU
NNNNNNNNNNNNNGGGGGGGVVVGGBBBBBBBGGGGGGWWWRRRRRRRRRZZZZZZZZZZZQQQQQQQQQQQNNNNNNNTTTTTTTTTTTYNNNNNNNNNNNNYYKKKLLLKKKKKKKEEEEEEMMMMVVVMWWWWWWUW
NNNNNNNNNNNNNZGGGGGVVVGGGGGGGGBGGGGGGOOWXRRRRRRRRRLZWWHZKZRZQQQQQQQQQQQQNNOOONTTTOBBTTTTTTTBBBNNNNNNNNNYYYKKKKKKKKKKKKEKEFEMMMMMZMMMWWWWWWUW
NNNNNNNNNNNNZZZZZJJJJGGGGGGGGGGGGGGGGOOOXRRRRRRRRRRWWWWWWRRZQQQQQQQQQQQQQNNOONMTOOBBBTTCTCTBBBNNNNNNYYNYYYYYKKKFKFFKKKEKEMMMMSMMMMMMWWWWWWWW
NNNNNNNNNNNNNZJJJJJJJJGGGGZGGGGGGGGGGGOOOORRRRRRRRWWWWWWWRWMWQQQQQQQQQPNNNOOOOOOOOOBTTTCCCBBBBNINNNNNYNYYYYYYFFFFFFFAKEKKMMMMMMMMMMWWWWWTTWW
NNNCCNNNNNNNNNNJJJJJJJGGGZZZGGGGGGGGGGGOOLLTTRRRRRWWWWWWWWWWWWQQQQQQQQQQQOOOOOOOOOOBBBTTCCBBBBBNNTNNYYYYYYYYYFFFFFFAAAKKMMMMMMMMMMMWWWWWWTTW
CCCCCNNNNNNNNJJJJJJJJCZZZZZZZGGGGGGGGGGOOLLTTRRRRRWWWWWWWWWQQQQQQQQQQQQQDDOOOOOOBOBBBBBBBBBBBBBYTTNYYYYYYYYYYFFFAALAAAAAMMMMMMMMMMMMWWWWWWTW
CCCCCIINNNNNJJJJJJJJJJZZZZZZZGGGGGGGGGGGGTTTRRRRRWWWWWWWWWWQQQQQQQQQQQQQDDOOOOOOBBBBBBBBBBBBBBBYYYYYYYYYYYYYFFFFFAAAAAAAMMMMMMMMMMMMMWWWWWWW
CCCCCCCCNNNKKKKJJJJJZZZZZZZZGGGGGGGGGGYTTTTTTRRWRWWWWWWKKZQQQQQQQHQZZTQQQOOOOOOOOBBBBBBBBBBBBBBYYYYYYYYYYYYYFFFFAAAAAAAAMMMMMMMMMMMMMWWWWWWW
CCCCCCCCCNNNKJJJJJJZZZZZZZZZZGGGGGGJGTTTTTTTTRWWWWWWWWKKKKQQQQQQQQQQZZQZZOZOOOOOOBBBBBBBBBBBBKKYYYYYYYYYYYYWWFFFAAAAAAMMMMMMMMVVVMMMMWWWWWWW
CCCCCCCCCNNNKMJJJJJZZZZZZWWWNWWKGGOJOTTTTTTTBRWWWWWWWWKKKKQQQQQQQQVZZZZZZZZOOOBBOBBBBBBBBKKBKKKKYYYYYYYYYYWWWFWFAAAAAAAMMMVVVVVVVVMMWWWWWWWW
CCCCCCCCCCNNJJJJJJJZZZZZZZWWWWWWWGOOOOTTTTTTTWWWWWWWWKKKKKQQQQTQQQVVVVZZZZZMGBBBBBBBBBBBBBKKKKKYYYYYYYYYYWWWWWWWAAAAAAAMMMMVVVVVVVWWWWPPPPPW
CCCCCCCCCNNNNJJJJJJZZZZZZZZWWZZZWOOOOTTTTTTTWWWWWWWWWKKKKEEEZZZZZZZVVVZZZZZGGBBBBBBBBBBKKKKKKKLLLYYYYYYYYGWWWWWWWWWWAAMMMMMMVVVVWWWWWWPDPPPW
CCCCCCCCCCNFNFJJJZJJZZZZZZZWWZZBBBBOQQQTTTTTTWWWWWWWWKKKKKEGGZZZZZZZZZZZZZZGGGGGBBBBBKKKKKKKKKKLLYYYYYYYYGWWWWWWWWWWWWPMMMVVVVVVWWWPPPPPPPPP
CCCCCCCCCCNFNFFZZZZZZZZZZZWWWZZBBBBBBBBBBTTRWWWWWWWWWWKKEEEGGGRZHHZZZZZZZZGGGGGBBBBBKKKKKKKKKKKKKYYYYYYYYWWWWWWWWWWWWWPPPPVVVVVVWWWPPPPPPPPP
CCCCCCCCCCCFFFFFFZZZZZZZZZZZZZZBBBBBBBBBBTRRRRRRWWWWWWWWEEGGGGRRRHZZZZZZZZGGGGGGBBBEEEGKKKKKKKKKKYYYYYYYWWWWWWWWWWWWWWWPPPPDDVDWWWPPPPPPPPPP
CCCCCCCCCCCFFFFFFFZZZZZZZZZZZZWBBBBBBBBBBQRRRRRRWRRRWWWWWEEGGGGRHHHZZZZZGGGGGGGGBBEEEEEKKKKKKKKKKKYKMJYYLWWWWWWWWWWWWWWWWWDDDDDWWWPPPPPPPPPP
CCCCCCCCCCCCXPFFFFFZZZZZZZZZZZWBBBBBBBBBBQRRRRRRRRRRWRWWWEWRRRRRHHZZZZZZGGGGGGGGGBEEEEEEEEEKKKKKKKYKMMMLLLLWWWWWWTWWDDDDDDDDDDWWWWPPPPPPPPPP
CCCCCCCCCCCCPPPPFFFZZZZZZZZZZZBBBBBBBBBORRRRRRRRRRRRRRWWWWWWRRRHHHHHHHZZGGGGGGGGGGEEEEEEEEKKKKKKKKKKMMMMMLLLLWWWTTWTTTDDDDDDDDDDDWPPPPRRPPPC
CCCCCCCCCCCPPPPPFFFFZZZZZZZZZZBBBBBBBBBOLLRRRRRRRRRRRRWWWWWRRRHHHHHHHHHGGGGGGGGGGGEEEEEEEEEKKKKKKMMMMMMMMMMLLLWWTTTTTTTTDDDDDDDDDDPPPPPRPPPP
CCCCCCCCCCCCPPPAFPLFZZZZZZZZZZBBBBBLLLLLLLRRRRRRRRRRRRWRWRRRRRHHHHHHHHHHGGGGGGGBBBGEEEEEEEEEKKKKMMMMMMMMMMMLNNNNTTTTTTTDDDDDDDDDCCCPPPPRRRPR
CCCCCCCCCCPPPPPPPPBBZZZZZZZZBBBBBBBLLLLLLLLYYRRRRRRRRRRRRRRRRHHHHHHHHHHHGGGGGGGBBBGGGEEEEKKKKKKKMMMMMMMMMMMLLNNNNTTTTDDDDDDDDDDCCCCPPRRRRRPR
CCCCCCCCCCCPPPPPPPBBBBBBBBDWBBBBBBBLLLLLLLLLRRRRRRRRRRRURRRRHHHHHHHHHHHBBBBBBBBBBBGGGEEEEEKKKKKMMMMMMMMMMMMLLNNNTTTTTTTDDDDDDDCCCRRRRRRRRRRR
CCCCCCCCCCPPPPPPPBBBBBBBDDDWBBBBBBBLLLLLLLLLRBRRARLLLLRUURRRHWHHHHHHHHGBBBBBBBBBBGGQGEEEEEEKKKKMMMMMMMMMMMMMLNNNTTTTTTTTTDDDDDDWRRRRRRRRRRRR
CCCCCCCCCCPPPPPPPBBBBBBDDDDDBBBBBBBBBBOLLLLLBBRRRLLLLLRUUUURRHHHHHHHHHHBBBBBBBBBBGGGGEEEEKKKKKMMMMMMMMMMMNNNNNNNNNNNTTTTWDDDDDWWRRRRRRRRRRRR
CCBBBBCCCCPPPPPPPBBBBBBBBDDDBBBBBBBBBBLLLLLLBBRRLLLLLLRUUUUURHHHHHHHHHHBBBBBBBBBBAAAAEEEEKKKKKKKMMMMMMMMMNNNNNNNNNNNTWWWWWWWDWWWWRRRRRRRRRRR
RRBRRBCBBCPPPPPPPBBBBBMMDDDDBBBBBBBBBBBLLLLLBBBRBLLLJJJUUUUUUEHHHHHHQQHBBBBBBBBBBAAAAEEEKKKKHHHHHHHHMMMMMNNNNNNNNNNNTWXWWWWWWWWWWWHRRRRRRRRR
RRRRRBBBBBBBJPPBBBBBBBBBBDDDBBBBBBBBBBBLLLLLLBBBBBBLJJJBUUUUUEEHHHHQRQQQBBBBQQAAPAAAAEEAKKKKHHHHHHHHMZMMMNNNNNNNNNNNNWXWWWWWWWWWWWHRRRRRRRRR
RRRBBBBBBBBJJPPBBBBBBBBBBBDDIIIIIIIIBBBLLLYLTTBBBBBLJJBBBUUUUEEHHHHQQQQQBBBBQQAAAAAAYEAAAKKKHHHHHHHHMZNNNNNNNNNNNNNXXXXWWZZZWWWJWWRRRRRRRRRR
RRRBBRBBBIJJIBBBBBBBBBBBBBIIIIIIIIIIBBBLLLLLTTBBBBBLBJBBBBUBBHHHHHHHQQQQBBBBAAAAAAAAAAAPYKKKHHHHHHHHMZZZNNNNNNNNNNNNNXXXXZZBZZJJWWRRRRRRRRRR
SSSSSSXBBIJIIIIBBBBBBBBBPIIIIIIIIILLLLLLLLLTTTBBBBBLBBBBBBBBBHHHHHHHQQQQBBBBAAAAAAAAAEEEEKKKYZZHHHHZZZZZZNNNNNNNNNNNXXXXZZZZZZJJJJJARRRRRRRR
SLSSSSSIIIIIIIIBBBBBBBBIIIIIIIIIIIIILLLLLLLLTBBBBBBBBPBBBBBBHHHHHHQQQQQQQQQQAAAAAAAAAEEEEYKKYYYHHHHZZZZZZNNNNNNNNNXXXXXXZZZZZZZJJJJAAARRRRRR
SSSSRSSIIIIIIIBBBBBBBBBBIIIIIIIIIIIELLLLLLTTTTBBBBPPPPPBIIBBBHHHHHHQQQQQQQQQQAAAAAAAAEEEEYYKYRRHHHHZZZZZZNNNNNNNNNNXXXZZZZZZZZJJJAAAAAARRRRR
SSSSSSSPIIIIIIIIBBBBBBBKIIIIIIIIIIILLLLLLLTTTBBIIPPPPPIIIBBBBBHHHHHQQQQQQQQQQAAAAAAAAEEEEYYYYRRHHHHZZZZZZZNNNOONTTNNZXZZZZZZZJJJAAAAAAARRRRR
SSSSSSPPPIIIIIIIIBBEBBBIIIIIIIIIIIILLLLLLLWTIOOIIIIIPIIIIBBBBHHHHHHQQQQQQQQQQQQAAAAAAEEEEYYYYRRRVVVZZZZZZZZNNNNNTTTZZXZZZZZZZZZZAAAAAAARRRRR
SSSSSSIIIIIIIIIIIIEEBBQQIIIIIIIIIIIILLLLLLLLOOOIIIIIPIIIIBHBHHHHHHHHHQQQQQQQFQQAAAAAAEEEEYYYYYRRRVVVZVZZZZQCQQTTTTTTZZZZZZZZZAAAAAAAAAAAKRRR
SSSSSSIIIIIIIIIIIIIIIIQQIIIIIIIIIILLLLLLLLLLROOOIIIIIIIIIHHHHHHHHHHHHQQQQQQQFFFBAAAAAEEEEYYRRRRRRVVVVVVVVQQQQQQTTTTTZZZZZZZZZOAAAAAAAAAAKRRF
SSSSIIIIIIIIIIIIIIIIIIIQQIIIIIIIIILLLLLLLLLLOOOOOIIIIIFIIHHHHHHHDDDDQQQQQQQQQQFFQAEEEEEEEEEYRRRRRVVVVVVCVQQQQQLQTTTZZZZNZZZXAOAAAAAAAAAAAAAF
SSIIIIIIIIIIIIIIIYYIJIIQQIIIIIIIIILLLLLLLLLLIIIIIIIIIIIJIHHHHHHHDDDDDQKQQQQQQQQFQAEEEEEEEEEYRRVVVVVVVVVCCQQQQQQQTTTTZZZNNNTTAAAAAAAAAAAAAAAF
SSISSSIIIIIIIIHIIYIIHQQQQQIAAIIIIILLLLLLLLLLLIIIIIIIIIJJJJHHHHHHDDDDQQKQAQQQQQQQQAEEEEEEEEEYVVVVVVVVVVVVVVQQQQQQTTTTBBNNNTTTTAAAAAAAAAAAAAAF
SSSSSSIIIIIIIHHHHHHHHHQQQAAAAIIAIALLLLLLHLLLIIIIIIIJJJJJHHHHHHHPDDDDDDDQQQSQQQQQQQEEEEEEEEEYYVVVVVVVVVVVFQQQQQQQQQTBBITNTTTTTTTTAAAAAAAAAAAF
SSSSSSIIIHHIHHHHHHHHHHQQQQAAAAAAAAALHHHHHHHIIIIIEEJJJJJJJHHPPPPPDDDDDDDDSSSQQQQQQQEEEEEEEEEYVVVVVVVVVVVVQQQQQQQQQQQBBTTTTTTTTTTTTAAAAAAAABBB
SSSSSSIIIHHHHHHHHHHHHHQAAAAAAAAAAAAHHHHHHHIIEEIEEJJJJJJJJEEPPPDDDDDDDDBYYSSQQQQQQQAYYDYYYYYYVVVRVVVVVVVVQQQQQQQQQQKBBOTTTTTTTTTTAAAAAAABBBBB
RSSSSSSIIHHHHHHHHHHHHHQAAAAAAAHAAAAHHHHHHHHHHEEEEJJJJJJJEEEPDDDDDDDDDYBBYYQQQQQQQYYYYYYYYYYYYYVVVVVVVVVQQQQQQQQQQBBBBBZTTTTTTTTTTTAAAAABOBBB
RSSSSSSHHHHHHHHHHHHHHHQQQDAAAAHHHHHHHHHHEHEEEEEEEEJEJEEEEEEDDDDDDDDDDYYYYYSQQQQQQQYYYYYYYYYYYYVVVVVVVVVVQQQQQQQBHBBBBBBTTTTTTTTTTTHAAAAABBBB
RRRRRRVHHHHHHHHHHHHHHHHDDDAAAHHHHHHHHHHHEEEEEEEEEEJEEEEEDDDDDDDDDDDDDYYYIYSQQQQQQQQYYYYYYYYYYYVVVVVVVVVVVQQQQAAAAAAAAABBTTTTTTTTHHHAHAHBBBBB
RRRRRRVHVVHVHHHHHHHHFFHDDDAAADHHHHHHHHHHECEEEEEEEEEEEEEDDDDDDDDDDDDDDYLYYYSSSSSSQQQYYYYYYYYYYVVVVVVVVVVVVVQQQAAAAAAAAABWTVVTTTTTHHHHHHHHHBBB
RRRRRRVHVVVVVHHHHHHHFFFFDDDDDDHHHHHHHHHHECEEEEEEEEEEEEEDDDDDDDDDDDDDDLLLYQEESSSSSSSYYYYYYYYYYHVVVVVVVVVQQAAAAAAAAAAAAABSSSSTTTAHHHHHHHHHHBBB
RRRRVVVVVVVVVVHHHHHHHFFDDDDDDDDHHDHHHHHHHEEEEEEEEETTEETDDDDDDDDDDDDDLLLLLEEEESSSSSSSSYYYYYYYYHLVVVLVVPIQQAAAAAAAAAAAAASSSSQTSTNNHHHHHHHHHHHH
RRRRRVVVVVVVVVHHHHHHHFFDFDDDDDDDDDDHDHHHHHEEEEEEEETTTTTTTDDDDDDDDDDDLLZZZZEEEEEEESSSSYYYYYYXXLLLLLLIIIIIIAAAAAAAAAAAAASSSSSSSSNNNHHHHHHHHHHH
RRRRVVVVVVVVVVDHHHHHFFFFFIDDDDDDDDDDDNHHHHEEEEEEEETTTTTTTTDDDDDDDDDDDZZZZZZEEEEEESEESEGYNXXXXXLLLLLIIIIIIAAAAAAAAAAAAASSSSSSNNNNNHHHHHHHHHHH
RRRRVVVVVVVVVVVHHHHFFFFFIIDDDDDDDDDDDNNHHNQQEEEEEETTTTTTTZDDDDDDDDDDZZZZZZEEEEEEEEEEEEGXXXXXLLLLLLLLIIIIIAAAAAAAAQSSSSSSSSSSSSVHHHHHHHHHHHHH
RRRRVVVVVVVVVVVJOHJFJJJIIDDDDDDDDDDDDNNHNNNNIEEEEETTTTTTTDDDDDDDDDDDZZZZZZZZZEEEEEEEEEEEXXXXXXXLLLLLLIIAAAAAAAAAAQSSSSSSSSSSGSVHHHHHHHHHHHHH
RRRVVVVVVVVVVVVJOJJJJJJJIDDDDDDDDDDDNNNNNNNIIEYYEKTTTTTTTTDDDDNDNDZZZZZZZZZZEEEEEEEEEEXXXXXXXLLLLLIIIIIAAAAAAAAAAAAAZSSSSSSSSSSSVVVHHHHHHHHH
RKRRVVVVVVVVVVJJJJJJJJJIIIIIDDDDDDDDDNNNNNNIIINYYATNTBTTTTDDDDNNNNNZZZZZZZZZEEEEEEEEEXXXXXXQLLLLLLIIIIIAAAAAAAAAAAAASSSSSSSSSCVVVVVHHHHHHHHH
KKRRRVNVVVVVJJJJJJJJJJIIIIIIDDDDDDDDDNNNNNNNNNNYNAANABTTTTADDDNNNNNNZZZZZZZZEEEEEEEEXXXXXXQQQLLQLQIIIIIAAAAAAAAAAAAAASSSSSSSSVVVVVHHHHHHHHHH
KKKKRNNVVVVVJJJJJJJJJIIIIIIDDDDRDDDDNNNNNNNNNNNNNAAAAATTAAAADDNNNNNNZZZZZZZEEEEEEEEXXXXXXXQQQOJJJJJJQIIAAAAAAAAAAAAAASSSSSSSSVVVVVVVVVVHHHHH
WKKKWNVVVVVFJJJJJJJJJIIIIIIDDDDRRDNNNNNNNNNNNNUUUAAAAAAAAAAAADNNNNNSSZZZZEEEEEEEEEEEXXXJJJJJJJJJJJJJQIIIMEEEEEEEAAAAAEESSSYVVVVVVVVVVWVHHHHH
WWWWWNNVVVVVVJJJJJJJJIIIIIIDDDRRRRRRNNNNNNNNNNNAAAAAAAAAAAAAADNNNNNSSSZZZEEEEEEEEEEFXSSJJJJJJJJJJJJJQQDSMMMEEEEEAAAAAESSSYYYVVVVVVVVVVVHHHHH
WBWWWNNNVVVVJJJJJJJJJIIIIDDDDRRRRRRRNNNNNNNNNNEAAAAAAAAAAAAAANNNNNNZZZZZEEEEEEEEEEEEXSSJJJJJJJJJJJJJQQSSSSSSEEEEAAAAAEESSVVVVVVVVVVVVVVHHHHA
BBWWWNNNVVJJJJJJJJJJJIIDDDDDRRRRRRRNNNNNNNNNNNEEEAAAAAAAAAAAANNNNNNSSSSEEEEEEEQEEENNNSSSJJJJJJJJJJJJQQSSSSSEEEEEAAAAEEEEVVVVVVVVVVVVVVHHHHHA
BBWWWWNNNNNJJJJJJJJJIIIDDLLORRRRRRRRNNNNNNNNNNEEEAAAAAAAAAAAANNNNNNNNNEEEHHHEEEEEEEEESSSJJJJJJJJJJJJQSSSSPPPEPEEAAAAEEEEVVVVVVVVVVVVVVAAAAHA
CBBBGGNGGNNNJJJJJJJJJILLLLLLQLRRRRROOOONNYNEEEEEAAAEEAAAAAAEENNNVVNNNNHHHHHEEEEEEEEESSSJJJJJJJJJJJJJSSSSSPPPPPPPAAAAEEEEEVVVVVVVVVSSVVVAAAAA
CBBBGGGGGNNNNJJXJJJJJXBBBBLLLLRRRRRROOOONNNEOEEEEAAEEAAAAAAEEENNNVNNNNNNNHHEEEEESSSSSSSJJJJJJJJJSSSSSSSSSPPPPPPPPPEEEEEOOVVVVVVVSSSAAAAAAAAA
CCBBGGGGGNNNNJXXJXXXXXBBBKLLLFRURRUUOOOOONOOOOOEEEEEEAEAAAEEENNNNNNNNNNNNEHEEEHEESSSSSSJJJJJSSSSSSSSSSSSPPPPPPPPPEEEEEEEOAAVVVVVVVAAAAAAAAAA
CCBGGGGGGGNNXXXXXXXXXBBKBKKLLKUUUUOOOOOOOOOOOOOEEEEEEAEAEEEENNNNNNNNNNNNREEEEHHEESSSSSSJJJJJSSSSSSSSSSSSPPPPPPPPPEEEEEEEERAAAAAVAAAAAAAAAAAA
CCCCGGGGGGNNXXXXXXXXXKKKKKKKLKKUUUUOOOOOOOOOOEEEEEEEEEEAEEEEEENNNNNNNNNNNNNJJIHHEHSWSSSJJJJJSSSSSSSSSPPMPPPPPPPPPEEEEEEEERAAAAAAAAAAAAAAAAAA
CCCCGGGGGGGGXXXXXXXXXKKKKKKKKKKUUUUUDDOOOOOOOEEEEEEEEEEEEEEENNNNNNNNNNNNNJJJJIHHHHSWWSSJJJJJSKKKKSSSSSPPPPPPPPPPPVVEEEEEEEAAADAAAAAAAAAAAAAA
CCCCGGGGGGGGXXXXXXXXXKKKKKKKKKKKKUUUDDOOOOOOOOEEEEEEEEEEEREEENNNNNNNNNNNNJJJIIIIISSSSVSSSRRSSKKKKSKSSPPZPPPPPPPPPPVEEEEEFFFADDADAAAAAAAAAAAA
CCCGGGGGGGGGXXXXXXXXXXXKKKKKKKKKKKDDDNNOOOOOOEEEEEEEEEEEEREEENNNNNNNNNNXWJJJIIIIVVVVVVVVSKKKKKKKKSKKKKKZZPPPPPPPPPEEOEEEEFFDDDDDDDAAAAAAAAAA
CCCCGGGGGGGGFXXXXXXXXKKKKKKKKKKKDDDDDNNNOOOOEEEEEELEEEEEERRRRRNNNNNNNNNXWJJJIIIIIIVVVVVIIKKKKKKKKKKKKKKKPPPPEEEPPPEEELEEEEDDDDDDDDAAAAAAAJAA
CCCGGGGGGGGGFFXXLLXXKKKKKKKKKKKDDDDDDNNOOOOOEEEELLLLLLEEEERERRRNNNYNNNNWWWWIIIIIIIVIIIVIIYKKKKKKKKKKLLLKKPPPPEEEEPEEZEEEECDDDDDDDDAAAAAAXAAA
CCCCGGGGGGGLLFFLLLLXXKKKKKKKKKKDDDDDNNOOOOOEELLLLLLLLLEEEEEEENNNNNMMMNNWWWIIIIIIIIIIIIIIBKKKKKKKKKKLLLKKEPEEEEEEEEEEEEEAEDDDDDDDDDDAXAXXXAAA
CCCCGGGGGGLLLLLLLLXXDDKKKKKDKDDDDDDDDNNNNNNNNLLLLLLLLLEEEEEZENNRRRMLLLWWWWWWWIIIIIIIIIIIBKKKKKKKKKKKLLLLEEEEEEEEEEEEEEEAEADDDDDDDDDXXXXXXXXC
CCCCGGGGGGWLLLLLLLXDDKKKKKKDKDDDDDDIIINNNNNNLLLLLLLLLLEEEEZZNNNMMMMLMMMWWWWWWIIIIIIIIQIBBBKKKKKKKKKKKKLLEEEEEEEEEEEEEPPAAADDDDDDDXXXXXXXXXXC
CCCCGFGGGGWLLLLLLLLDDKKKDDDDDDDDDIIIIINNNNNNNLLLLLLLLLLEZZZMMMMMMMMMMMMGWWWWIIIIIIIIIQQQBBKKKKKKKKKKKKKLEEEEEEEEEEEEEEPAAADDDDDDDXXXXXXCCCCC
CCCCCLLMGLLLLLLLLLLDDDKDDWDDDDDNNNNIIIIIIINNNNNLLLSSLSLZZZZZMMMMMMMMMMMMMWWWIIIIIIIIIIQQBBBKKKKKKKKKKKKKEEEEEEEEEEEEEPPPPADDDDDHDXXXXXMMMCCC
CCLLLLLLLLLLLLLLLLLDDDKDDWWDWWWWNNNIIIIINNNNNNLLNLSSSSLLZZZMMMMMMMMMMMMMMWWWWIIIIIIIIKQQBBQKKKKKKKKKKKKKEEEEEEEEEEPEEEPPPDDDDDDDDXXXXXXMMMMM
CCCCLLLLLLLLLLLLLLLDDDDDWWWWWWPPNNNIIIIJNNNNNNNNNLSSSULLZZZZMMMMMMMMMMMMMMWWIIIQIQIQQQQQBBQMKKKKKKKKKKKKEEEEEEEEEEPEEPPPDDDDDTDDDXXXXXXMMMMM
CCCCLLLLLLLLLLLLLLLDDDDDDWWWWWWITIIIIIIINNNNNNNNNNSNSUZZZZZMMMMMMMMMMMMMMMWWIIQQQQQQQQQQQBQQXKKKKVVHHKKKKEEEEZEEEEPPPPPPPDDDDDDDXXXXXXXMMMMM
CCCCCLLLLLLLLLLLLLLLLLDDDWWWWWWIIIIIIIIINNNNNNNNNNNNNQQZZZZZMMMMMMMMMMMMMMWWIIZZQQQQQQQQKQQQXXKKVVHHHHHHEEEEEZSEEEPPPPPVPPDVXXXXXXXMXXXMMMMM
CCCCCLTLLLLLLLLLLLPLDDDDDWWWWWWIIIIIIIIINNNNNNNNNNNNQQZZZZZZMMMMMMMMMMMMMMMMZZZQQQQQQQQQQQQXXXXXHHHHHHHHHHZZZZSEPPPPPPPVVVVVVVXXXXXMMMMMMMMM
CCCCCCTTLLLLLLLLTLLDDDDDDDWWWWWWWIIIIIINNNNNNNNNNNNNQQQQQQQQQMMMMMMMOOOMMMMZZZQQQQQQQQQBBQQXXHHHHHHHHHHHHHZZZZZZPPPPPPPVVVVVVVVXVVXMMMMMMMMM
CCCCCCCLLLLLLLLLTLLDDDWWWDWWWWWWWIWTIINNNNNNNNNNNNNNQQQQHQQQKKKKKOMOOXOMMMMZZZZQQQQQQQQQQQQQXHHHHHHHHHHHGZZZZZLZPPPPPPPXNVVVVVVVVVMMMMMMMMMM
RRCCCCRRLYLLLLTTTIDDDDWWWWWOWWWWWWWWIIWNNNNNNNNNNNNNQQQQQQQQQQKKKOOOOOOMMMMZZZZZQQQQQQQQQQQQHHHHHHHHHHHHGZZZZZZZPPPPXXPXVVVVVVVVVVMMMMMMMMMM
RCCCRRRLLYRAALLTRIIIIIBWWWWWWWWWWWWWWWWNNNNNNNXNNNNQQQQQQQQQQKKKKKOOOMMMMMMMZZZZQQQQQQQQQQQQWHHHHHHHHHHHHZZZZZZZYPPPXXXXDVVVVVVVVVVMMMMMMMMM
RRRCRRRRRRRRRZRRRIIIIIBBWWZWWWWWWWWWWWNNNNNNNXXXNNNQQQQQQQQQKKKKKKOOOOMMMMMMZKKKKKQQQQQQQQQQWWHHHHHHHHHHSZZZZZZZMPPPXXXXXVVXMMVVVVMMMMMMMMMM
RRRRRRRRRRRRRRRRIIIIIIBWWWZZWWWWWWWWWNNNNNNTTXXNNNQQQQQQQQQQQQNNNNNOOOMMMMMMZZZZZZQQQQQQQQQWWWHHHHHHHHHSSZZZZZZMMMMMXXXXXXXXMMVVVVTMMMMMMMMM
RRRRRRRRRRRRRRFRIIIIIIBBBWWZZWWWWWWWWNNNNNTTTXNNQQQQQQQQQQQQQQZNNNNNNNMMMMMMMZZZZZQWWQQQQQQWWWWHHHHHHHHHSZZZMZMMMXXXXXXXXXXXXXXXVCMMMMMMMMMM
RRRRRRRRRRRRRRRIIIIIIIIBBWWZZZWWWWWWWWWNCTTTTTQQQQQQQQQQQQQQQSNNNNNNNNNNMMMMMMGGGGQWWWWQQQQWWWWHLHHHHHHSSSZMMMMMMXXXXXXXXXXXXXXCVCMMMMMMMMMM
RRRRRRRRRRRRRRRIIIIIIIISBBBIZZZWWWWWWWTTTTTTTTQQQQQQQBBQQMQPQNNNNNNNNNNNMMMMMMGIGGGGWWWQQVQLLLWLLHHHHHHHSMMMMMMMMMXXXXXXXXXXXXCCCCMMMMMMMMMM
RRRRRRRRRRRRRRRIIIIIIGSSBBIIIIZZWWWWWTTTTTTTTTTQQQGQQBBQQQQQHNNNNNNNNNNNMMMGGGGGGGGGWWWWQQTTLLLLLHUHHHHHSMMMMMMMMMJJTXTTXXXXXXCCLLKKMMMMMMMM
RRRRRRRRRRRRRRCIIIIISSSSBBSIIIIIWWWWWWITTTTTTTTTQQGQQBBBQQQNHNNNNNNNNNNNNMMMMGGGGGGGGGWTTTTTTTLLLLUHHHHSSMMMMMMMMMJJTTTTXXXXXLLLLLKLLMMMMMTM
RZRRRRRRRRRRRRIIIIIIIISSSSSIIIIIWWWIIIITTTTTTTTTTTGQAGBBQBQNNNNNNNNNNNNNNNMMMGGGGGGGGTTTTTTTTTLLLTTHHHHSMMMIMMMJMJJJJTTTXXXXXZZLLLLLLWMWMMMZ
ZZRRRRJRRRRRRRIIIIIIIISSSSSIIIIIWIIIIIIIIITTTTGGGGGGGGBBBBBNNNNNNNNNNNNNNNMMMMMGGGGGGTTTTTTTTTLLTTTHLAAAAMMAJJJJJJJJTTTTXXXXXLLLLLLLLWWWMMMM
ZZZRRRJJRRRRRIIIIIIISSSSSSIIIIIIIIIIIGGIIIIITTGGGGGGGGGGGBBBNNNNNNNNNNNNNMMMMMMMMGGGTTTTTTTTTTTTTTAAAAAAAMAAJJJJJJJJTTTTXXXXXXLLLLLLLLWWMMWW
ZZZRZZJJJJRJIIIIIIIIISPSSSFPIIIIIIIIIGGIIIIGGGGGGGGGGGGGGBNNNNNNNNNNNNWNMMMMMMMMMMMTTTTTTTTTTTTTTTTAAAAAAQAAAAAJJJJJJJTXXXXXXXLLLLLLLLLWWWWW
ZZZZZJJJJJJJJJJIIIIIIIIJSFFFIIIIIIIIIGIIIIIGGGGGGGGGGGGGGGQYNNNNNNNYNNYNMMMMMMMMMMMTDTTTTTTTTTTTTTAAAAAAAAAAAAAJJJJXXXXXXXXXXXLLLLLLLFLWWWWW
CZZZZJJJJJJJJJIIIIIIJIJJFFFFIIIIIIIIIIIIIIGGGGGGGGGGGGGGGYYYYYYNNNNYYYYYYYDMMMMMMMMMDDTTTTTTTTTTTTAAAAAAAAAAAAJJJJJXXXXXXXXXAAAALLLLLLLLWWWW
CZZZZJJZZJJJJJJJJIIIJIJFFFFFIIIIIIIIIIIIIIIGGGGGGGGGGGGGGYYYYYYYYYYYYYYYYYDMMMMMMMMDDDDTTTTTTTTTTTADAAAAAAAAYJJJJJJJJXXXXXXXXXXALLLLLLXBBWWW
ZZZZZZZZAZJJJRJRRRIIJJJOWWYYIIIIIIIIOOIIIIGGGGGGGGGGGGAAAYYYYYYYYYYYYYYYYYYMMXMMMMDDDDDDDDTTDDTTTTDAAAAAAAAAYJJJJJJJWWWXWWWXXAAAAALLYLXBBWBB
ZZZZZZZZAZJRRRRRRJJIJJOOWWYYIIYIOOOOOOOIIGGGGGGGGGGGAGAAYYYYYYYYYYYYYYYYYYYYMMMMEDDDDDDDDDDDDDDDTDDDAAAAAAAAADJJJJJJWWWWWWWWAAAAALLLLLLBBBBB
ZZZZZZZZZZJJRRRJJJJJJOOOWYYYYYYIOOOOOOOOGGGGGGGGGGGGAGAAAYYYYYYYYYYSYYYYYYYEEEMEEEEDDDDDDDDDDDDDDDDDDAAAAAADDDDJDJKWWWWWWWAAAAAALLLSBBBBBBBB
ZZZZZZZZZJJWWWWJJJJOJOOOOOOCYYBOOOOOOOOOGGSGGGGGGGGGAAAAYYYYYYYYYYSSSSYYEEEEEEEEEEDDDDDDDDDDDDDDDDDDDAADDDDDDDDDDKKKWKKWWWWWAAAAALLSBBBBBBBB
ZZZZZZZZZZZWWWWJJJJOOOOOOOOOOFOOOOOOOOOOOGGGGGGDDDQQQQQAAYYYYYYYYSSSSSSYEEEEEEEEEDDDDDDDDDDDDDDDDDDDDDDDDDDDDDKKKKKKKKKWWWXXAAAAASSSSBBBBBBB
ZZZZZZZZZWWWWWWWWJOOOOOOOOOOIOOOOOOOOOOOGBBBGGGDDQQQQQQQYYYYYYYYYYGSSSSSSSEEEEEEEDDDDDDDDDDDDDDDDDDDDDDDDDDDDKKKKKKKKKKWWXXXAAAAAASSSBBBBBBB
ZZZZZZWWWWWWWWWWJJOOOOOOOOOOOOAOOOOOOOOBBBBBGTSSDQQQQQQQYYYYYYYYOOCCSSSSSEEEEEEEEEDDDDDDDDDDDDDDEDDDDDDDSDDDDKEKKKKKKKKXXXXXWWWAAABBBBBBBBBB
ZZZZZZWWWWWWWWWWJJOOOOOOOOOOOOOOOOOOOBBBBBBBBSSSSQQQQQQQYVYYYYYYYOOCCCSSCEEEEEEEEEEDDDDDDDDDDDDDEEEDSSDDSDDDKKKKKKKKKKKXXXXXWWWWBUBBBBBBBBBB
SZZZZZZWWWWWWWWSOOOOOOOOOOOOOOOOOOOOBBBBBBBBBSSSQQQQQQQQYYYYYYYMYOOCCCSSCCCCCCCEEEEDDDDDDDDDDDDEEEEEPSSSSDKKKKIKKKKKKKKBXXXXXWWWBBBBBBBBBBBB
ZZWWWWZWWWWWWWWWOOOOOOOOOOOOOOOOOOOOOOOBBBBBBBBBQQQQQQQQQYYYYOOYYOKKCCCCCCCCCCCCEEEDDDDDDDDZEEEEEEESSSSSSSSSDKKKKKKKKKKXXXXXXWWWWWWBBBBBBBBB
ZWWWWWWWWWWWWWWWWNNNOOOOOOOOOOOOOOOOOOBBBBBBBBBBQQQQQQQQQQQQYOOOOOKKKCCCCCCCCCCEEEEDDDEDDDDDEEEEEEFSSSSSSSSSSKKKKKKKKKXXXXXXWWWWWWWBBBBBBBBB
MWWWWWWWWWWWWWWNNNNNOOOOOOOOOOOOOOOOOOBOBBBBBBBBBSQQQQQQQQQQQOOOOOCCKKCCCCCCCCCEEEEEEEEEEDDEEEEEEEEESSSSSSSSKKKKKKKKKKXXXXXWWWWWWBBBBBBBBBBB
MMWWWWWWWWWWWWWWWNNNOOOOOOOOOOOOOOOOOOOOBBBBBBBBBBQQQQQQQQQQQOOOOOCCCCCCCCCCLLCEEEEEEEEEEEEEEEEEEEEEMSSSSSISRKIIKKKKGKXXXXXXXWWWWBBBBBBBBBBB
"""
        }
    }
}