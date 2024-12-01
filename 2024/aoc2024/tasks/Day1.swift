import Foundation
internal import Algorithms

struct Day1: Day {
    typealias R1 = Int
    typealias R2 = Int
    
    static func task1(data: String, param: P1) throws -> R1 {
        let (left, right) = parse(data: data)

        return zip(left.sorted(), right.sorted()).map({ abs($0.0 - $0.1) }).reduce(0, +)
    }
    
    static func task2(data: String, param: P2) throws -> R2 {
        let (left, right) = parse(data: data)
        
        return left.map({ leftValue in leftValue * right.count(where: { rightValue in leftValue == rightValue }) }).reduce(0, +)
    }
    
    private static func parse(data: String) -> ([Int], [Int]) {
        let lines = data.split(whereSeparator: \.isWhitespace).map({ Int($0)! })
        let (leftElements, rightElements) = lines.enumerated().partitioned(by: { $0.offset % 2 == 0 })
        
        return (leftElements.map({ $0.element }), rightElements.map({ $0.element }))
    }
}

extension Day1 {
    static let data = """
58692   56129
45806   95015
61519   31093
77848   16487
98705   66749
40109   81197
26355   75702
68715   92381
56304   23460
37264   12036
43236   27331
43577   28168
77477   20638
53092   44012
37180   48396
20616   32705
97334   49126
73040   28694
21873   15026
84281   40798
42579   91319
58950   66300
41645   92178
94750   65101
32915   60179
44562   71858
93775   59687
19512   81598
99696   83818
21597   11382
56743   16557
78901   99638
41593   96712
32215   53944
44764   53944
18670   40607
96152   17765
25871   16487
51869   24957
18158   54482
58119   77417
56369   81197
18076   27795
39036   31093
10418   62892
45943   78393
97617   54482
58295   64971
85585   93419
52619   23460
90697   71454
74290   30929
75290   86123
20096   74997
95350   26162
38951   98705
75164   16487
91540   63048
15233   92381
85251   96622
68167   89320
50113   11239
71427   10418
24444   98705
36612   65035
45913   72508
61899   97524
17009   90579
36447   80725
29939   31093
98955   78393
86307   90907
79546   17968
20043   68154
89734   10418
22532   98705
38775   82700
51623   71276
23458   36867
20971   78684
28114   21873
98781   31388
55269   46378
15215   31093
25994   40798
18371   21873
57542   68652
49611   69512
86439   21873
28659   11473
40997   92284
69645   63538
54066   10418
77771   40798
68166   82468
60960   27531
70063   90621
23460   89404
15248   83526
71980   93951
93745   81047
69012   40798
32127   49126
37912   78393
44522   66779
46780   91790
17757   52357
22296   12692
21860   37749
13293   91094
65101   16557
80705   49126
89050   25994
88161   26162
74340   20670
31839   21873
25311   32705
94912   29820
19900   43167
38669   44521
42958   29495
43606   36746
82677   77417
98077   78921
30924   81591
25598   95296
28564   74437
11879   15403
27531   77593
39113   56398
53147   47022
12762   16557
89612   83096
83701   56880
73067   57030
24957   50324
77222   67603
17744   99063
62306   38370
14025   96622
59643   75583
67152   79005
58835   89612
61882   53944
28174   81598
84094   81598
84099   51482
42783   49701
45148   59184
55374   91670
53134   62411
38007   25421
26794   54482
56567   24607
10209   92381
13233   16352
18053   54482
69340   31093
32061   21873
58293   78513
46514   85839
74862   83096
66080   54482
95079   17822
26186   40798
16839   86836
55841   77124
92678   26808
49885   65391
56779   10020
35136   67136
35115   81598
70253   13483
28435   67494
34410   96622
39335   72612
60832   98705
61949   77417
16487   89612
43545   85993
39289   94417
64581   79338
13200   17235
58996   89612
21766   12440
33174   53944
96545   96622
13820   89612
75200   63765
63979   91806
54711   14141
34633   79885
49784   77593
34461   81007
65491   17490
48702   16557
93477   40798
50646   43959
98874   44341
83312   31512
87924   81173
24928   32705
33803   77593
34995   27808
31695   16557
96781   25994
52443   10418
31068   16557
45835   65101
39749   23460
52812   49126
84499   61472
76703   74413
15788   77417
46869   65101
35929   47384
31093   25736
13592   99027
47704   54696
17713   64743
71888   87351
35166   80828
32725   80342
92648   37749
86060   37749
18301   75483
57155   96622
76362   69037
99510   98552
27789   93066
84634   82443
57886   55692
65729   16557
62285   48939
32091   56805
75565   21873
90388   64049
47465   92381
98181   75353
85368   23460
47978   54621
14979   67106
61630   92381
56645   23460
23400   43430
23871   92381
33383   32705
57650   70467
79961   76721
30980   98667
65936   42768
98177   37749
74251   50601
19586   65101
82672   89612
97885   69420
16375   40976
13960   42919
48331   68948
71851   61884
39544   81598
26162   37143
96196   37749
56644   24957
44642   49170
75978   34663
23507   86221
10164   38813
80471   52851
23262   47386
50169   22415
82553   20839
37359   21873
66822   94258
77157   95081
89279   83096
70625   95877
96127   85802
57483   49126
84704   21873
95805   99427
10526   31093
41396   79445
27982   92381
67459   23460
47751   22401
96622   51524
75606   34459
48009   26647
71026   84513
85116   78393
51669   19783
41048   16165
25788   53944
55740   44351
84391   55762
62303   32778
61747   17977
54119   57740
95965   25994
37749   77593
34224   49126
63046   53944
55216   40138
78168   10418
57385   22620
36927   10294
64442   40798
81693   22002
28110   19900
88272   92381
14111   31683
44817   51880
81556   23477
44458   21970
12018   77420
96305   68195
37402   92381
18732   44095
42170   62400
46401   27531
87100   64743
39844   40798
57423   83184
26562   50029
54418   84284
58705   61925
82922   11177
87067   32705
47312   27531
38548   81197
53190   80348
29917   93023
42325   81197
95590   27531
54561   23126
97439   64743
11411   54482
58773   53944
70943   85056
60775   53944
28063   54482
47400   74202
10825   64490
72110   27531
24815   83532
78705   16487
99101   75491
12807   78725
54179   64498
68659   67308
33138   77593
49947   98705
80334   66710
34081   89511
52339   59338
76399   16557
93419   33219
41305   12246
70393   37749
97638   44272
83001   16688
72802   10380
52036   81598
35576   24957
11887   37223
35348   37749
45853   49126
53651   81197
63422   84087
39390   32627
50050   32718
86437   54268
89370   65101
24909   31388
98791   31888
87645   98705
83210   77417
96021   19900
53148   37749
30332   31148
94829   22219
31815   77593
19040   93953
85522   32812
71721   65101
55365   64743
90383   27531
77502   54482
91232   78231
82187   21873
89204   89612
20858   37749
69477   93247
81598   61349
64333   10118
14077   21873
29326   93949
24726   61981
54422   24957
30688   54482
20341   21202
50601   65101
86836   49126
44594   42246
86152   45889
63247   93419
82498   40798
79108   54192
91871   53390
99872   77417
41068   49126
96159   11626
62645   31093
45646   31093
26163   31093
90008   13269
87334   66094
87822   85292
70745   31847
21762   81395
41620   89001
70907   89612
26244   27854
79411   58099
55868   87368
11278   77417
37673   95517
25055   29383
54935   12762
43886   21873
37309   61228
55015   84361
63352   92381
58698   85675
81387   40798
29046   66504
58613   10418
82324   41647
13403   41567
10892   29289
77123   24957
33575   89612
48171   43026
90435   26525
89406   53944
27722   31362
80176   45398
93218   25976
41958   25454
56360   52645
90404   50800
39556   89612
84518   12762
25734   19900
34464   67008
67395   90719
99197   81598
62886   64743
66314   31575
64470   81598
91704   92381
49126   31093
44922   16557
61975   24788
60682   60858
64266   24957
76536   16557
16857   80320
52642   30266
54675   23460
45270   97669
25750   32914
81197   50551
61256   16557
18231   30651
36067   37882
26043   21873
29852   80234
96277   16557
91064   50601
63914   58965
69764   75318
24111   83096
57354   34874
91631   84294
97132   84499
74955   32953
12873   81598
75267   99914
53133   65101
20645   30462
98238   93419
12241   91866
67045   26186
73351   78393
10571   70926
18983   77417
61188   53944
53415   12892
11950   46304
19189   90035
49639   50601
50651   70751
79184   51245
93111   31691
48125   32705
32705   48623
91611   54121
56231   74893
61384   16557
50882   19816
68311   24056
18609   75970
31210   12762
97836   81598
30406   48671
21526   77493
29914   58274
82010   68984
25775   83096
29688   18265
90860   19900
95867   24957
31388   50389
65300   66889
25536   40549
40121   38677
77593   13569
67527   24957
12937   32705
49194   42027
27889   11799
48768   15747
39163   88157
64135   76645
29431   78317
20473   98689
11519   81598
83666   50089
71962   53944
12936   47345
67251   89135
25749   67545
16970   76986
79021   12762
24436   25522
40505   89612
48317   29070
99381   37882
62487   52063
19335   45479
87856   76582
85590   44438
39618   27531
61562   77417
45199   17823
80448   12762
17187   10418
79089   81197
32058   26701
82841   92787
40798   53461
86367   63332
13554   89592
68467   40417
35298   10360
35159   93419
85923   27309
72185   24957
56088   49126
63172   59362
58775   21873
43947   67792
27276   16735
76238   35704
93662   42478
86221   31093
56440   96803
47639   64743
85362   31093
26154   92381
13811   57981
35004   92381
61936   83096
77417   77593
51617   83096
71330   11566
72744   31861
42786   23188
62415   23827
33067   31786
28905   93419
50640   50601
14207   77593
30760   50601
33767   42776
21522   77417
37979   31093
98095   80409
78924   50601
23863   10376
11062   81598
14585   42821
59032   29344
65746   23569
37228   77972
96930   26603
91199   77417
23174   66423
74662   13386
84769   52236
95632   86168
98388   53284
19101   29813
74624   84414
41636   40798
50603   43055
73759   16557
62222   67636
36385   81598
94441   53944
66544   37749
29419   54482
37882   83066
76573   49126
41137   59092
72153   88283
40280   10418
89800   21942
20352   53944
97382   26162
48401   38866
95481   83096
34830   94495
52340   74883
72473   67919
23027   77417
84626   12031
55465   69547
90755   54548
36665   97629
28837   45543
88173   18396
48651   54898
29605   10707
61275   49126
21002   64743
40255   40798
34659   64743
30246   20886
70325   45581
83974   67963
61241   54482
11883   45819
12479   76035
17092   93759
39998   26162
71318   13849
49438   58801
34597   16487
64278   40798
80596   55825
91844   59420
74971   49148
25101   86221
22587   77104
31853   92381
23921   50601
61611   21873
23577   50601
62744   84499
85808   10418
79657   16487
62155   12847
15585   26162
99894   80428
41130   96622
19404   49517
25031   21873
68371   92381
46354   93419
68726   71530
84826   69572
76273   21873
64743   65807
51903   16557
36136   23751
40047   90769
95606   66439
53944   21993
94698   25710
25217   63275
86650   37749
92979   70928
10190   92381
86815   24957
96308   65101
40065   89612
92522   23460
70854   69671
80130   95318
80869   41574
63911   12762
78393   99596
37699   89007
54669   91261
88420   40064
97511   43961
93370   81197
87004   78393
88067   60647
91248   54482
26912   77449
62212   16557
35359   40798
43743   96327
23402   40435
52733   72898
38801   31859
69148   42844
23150   31093
11522   40798
87909   23460
74157   13421
87552   89035
97642   87767
58377   81990
36872   60818
28786   83096
10838   19340
51754   54482
49366   24110
62108   63081
46490   97762
92381   81598
14655   95778
94136   64743
52958   15704
35680   75439
99274   49126
95898   32705
80049   24323
78216   64743
34169   27636
37503   37749
91663   89612
44335   40798
18185   28044
81878   37656
43387   81598
90977   21873
72660   64743
39634   63703
47748   48934
79572   19518
61317   94748
99452   80662
59165   26162
28988   24957
10186   73375
97072   43422
23437   72900
25136   23835
70817   16557
81489   19816
16210   84879
79565   88447
16125   68082
18390   64667
22543   23460
90995   58285
88960   33239
34710   90430
20426   84502
30472   16487
81984   25342
71958   40798
70163   63310
33936   65348
96019   83283
15737   95669
47900   20364
45203   37749
73885   64544
41883   27531
19816   77681
23620   89095
93874   89612
37065   45856
13841   64743
20247   53944
23142   38721
31341   47276
83096   54494
26139   27531
32514   89612
88888   19816
83754   10418
66106   88915
38199   45336
57295   61535
46102   89612
79413   46077
13550   71765
67201   71450
11907   28021
70044   37250
33542   24957
11273   16557
85143   49126
90698   98705
66336   63745
52763   23460
21923   85051
12300   76577
64921   38328
49612   15170
42092   62194
81081   29744
85411   26540
88743   24957
67110   57072
66344   23460
13510   83821
10766   40798
59257   35914
39333   79820
41611   89612
89131   16487
79910   74195
79833   10972
70201   43597
45298   40798
93528   39040
67472   83096
28083   73486
99427   83096
57426   44711
70197   26162
41146   77870
32235   39025
83437   52047
19107   80128
62376   65101
37000   41525
38143   39047
97780   98705
54008   14030
62441   89612
95123   29967
71651   49358
43911   83096
77678   24957
40843   92381
81957   88406
97799   65543
34839   88608
79794   40798
41598   40240
75305   68392
47771   46860
38976   17664
36023   78321
35570   49374
48804   16006
23497   39006
21745   21216
35898   78381
68244   50448
49831   56014
70988   54482
19797   77193
50561   54482
58577   64145
94592   86756
57724   55135
54847   14131
72690   23460
47137   76125
23898   89612
87598   81598
13646   31388
53861   50601
83822   31668
21224   87493
16557   32745
28394   43896
15982   92381
20661   35634
98506   42840
16242   91366
32971   35609
33113   67658
87344   31857
14803   75667
34931   47173
38451   17170
67646   27531
80436   53944
81838   22536
95361   20123
65978   90413
76587   26162
70892   85750
23734   27189
89248   92064
17066   57813
52504   38134
98114   50508
30619   10418
91557   75651
88382   77897
95976   50494
72241   41135
58496   21873
93790   68603
30945   50899
54482   66005
96412   33777
10313   84492
32604   65384
31806   31093
18912   21152
32018   60171
38922   21106
82021   81197
90291   12302
62722   32705
37494   21470
33523   64743
23069   98705
84349   92381
26251   81638
42012   82734
77929   32705
39926   81197
98975   97049
33771   82816
22832   96615
75620   86221
35406   47499
70752   83096
88399   70473
51534   98705
85961   89612
14123   27531
85125   14573
34896   21065
82440   83096
99451   77759
44034   52525
11876   53944
47516   92381
75926   18104
35222   26162
99941   56786
34909   65101
24255   54482
11814   98705
98257   83096
95143   21873
50403   31093
71325   44054
82711   16487
11959   89612
88390   21612
52422   83096
39874   53944
72503   29230
13223   54482
53459   85940
51579   81598
31830   67725
51189   62613
10980   27338
94234   13514
34098   20372
83669   96622
22936   49126
26105   35302
51805   81598
35479   61032
10006   64579
82269   85007
77241   81197
86928   83096
18518   19816
86000   12762
41196   25803
81975   64197
24972   10418
45484   91706
20171   53178
59569   19816
85735   35587
30874   14008
15369   52468
25998   79528
37079   53944
"""
}
