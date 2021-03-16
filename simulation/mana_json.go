package simulation

import (
	"encoding/json"
)

type manaJSON struct {
	ShortNodeID string
	NodeID      string
	Mana        float64
}

type ConsensusMana struct {
	Consensus []manaJSON
}

var (
	manaDistribution = `{
		"consensus":[
			{"shortNodeID":"RcXN8ST2yM3","nodeID":"AuQXPFmRu9nKNtUq3g1RLqVgSmxNrYeogt6uRwqYLGvK","mana":999670278720337.9},
			{"shortNodeID":"EyJFsP3WaVX","nodeID":"6d3iJoug2kcpbtS93hTwA9uFP31DQyE4s9iDQE8xuJvX","mana":292333010271.5598},
			{"shortNodeID":"11111111","nodeID":"11111111111111111111111111111111","mana":9621323336.976624},
			{"shortNodeID":"391FRvqq5C5","nodeID":"ru1brgnQsLVtiwb4SBoeg6gMRH4RUAzg7h26x155MwW","mana":2973202821.7559648},
			{"shortNodeID":"X9F548nmNbm","nodeID":"D8TTjCA5TF2Tuy9yV8qU212UFC5Ae2nbbXYTvr8ddFud","mana":2042306238.6674123},
			{"shortNodeID":"brEM773P27x","nodeID":"F2JSWzn7NHQWaM6niPgnGaKTi8bU3vi7gwwaDPpzpbny","mana":1565082135.1100168},
			{"shortNodeID":"A1BEDSspbSn","nodeID":"4d7YZQqcjCkx8isitCnqzR8AET667p1rKjEMGKcHoDNP","mana":728811890.5439975},
			{"shortNodeID":"dzoxo5HpD6i","nodeID":"FtRspufNmx5M5kjpHuCAVN495Ao37MTEGFFEhK3Gs7rx","mana":115912629.13527815},
			{"shortNodeID":"VzsXqXuxW9","nodeID":"CfkVFzXRjJdshjgPpQAZ4fccZs2SyVPGkTc8LmtnbsT","mana":103102972.1197696},
			{"shortNodeID":"SnTjUBuapRs","nodeID":"BNk3dqncprQ6DiPxjrgvFEofLBdYru3wYZkz2cohmJJj","mana":88137321.22943768},
			{"shortNodeID":"2ZErA3PpnGL","nodeID":"dJzTf8mrZsr4QHh8GKjNPia7iuWUbjFUZ9o4j6yTzjY","mana":53826849.79699167},
			{"shortNodeID":"KnnMze58hRd","nodeID":"8ZWKXMojDnmenenkw4LvbmH7Pj9PLPBfobSXswsdqP3J","mana":23518964.30215719},
			{"shortNodeID":"EMoz6R6u48A","nodeID":"6NmWPJ7KvoziqgKeskxCxTuxtBAjmWFuesQCVYdUtFym","mana":19489186.191191025},
			{"shortNodeID":"7cmX7HB2JKm","nodeID":"3fRN1KYYBJFN47GDZtvxN9Yxd3aB6AP79KFEsuz4mxob","mana":11727831.62982483},
			{"shortNodeID":"hmpzL2SugVN","nodeID":"HQZJswN6EacbP55E65bBKJ9cEDRfnvbYSX6R2Yfhhk3v","mana":7098568.113006732},
			{"shortNodeID":"FMW9RZYKpGJ","nodeID":"6myw3HvvCP3Tq1n86Yz5YZLCsUnTJ9XZSZZLHgjfcHeY","mana":6999352.317111199},
			{"shortNodeID":"TVpoxUk7QfK","nodeID":"BfPPGcZt5sCNugwUBuoavtNyEALuZvcP4SMG8epHqrve","mana":6178522.632130066},
			{"shortNodeID":"ed1pRUk6LuH","nodeID":"G8zpECZcM55iXAHCVwk2hmMYkuEMf6Uj5PsJcReo4mW2","mana":5999627.345726015},
			{"shortNodeID":"Z2gy9PAufkd","nodeID":"DtVitHT96YEezxMv2upEBTHknzJ9BuWHW2MkHuMaes95","mana":5589067.77302075},
			{"shortNodeID":"Bg7PUsJaU5S","nodeID":"5J7nSkm5UFJVe6oeFdFv5Umk5FFbu6r7WVuSRSMXSnLx","mana":5090027.188078557},
			{"shortNodeID":"UfubZjkKZHo","nodeID":"C8nHzNuA9oxGvXkFo9HgVZj6RrK9KChMDnFJgL89skTD","mana":4995322.514972735},
			{"shortNodeID":"GPoUTpxGwuH","nodeID":"7CFBdxXGb4hhX9QeRLjcyht7zsmanFZquCKqTDU8yCrp","mana":4986318.039502273},
			{"shortNodeID":"Ta4jFHTZbV9","nodeID":"Bh6LX4yYm9rsiD3KvWn7QM3yDWr5EiZ1ufY4CNGwsssC","mana":4796613.904995576},
			{"shortNodeID":"QChVuCLJH8a","nodeID":"ALUtRvTfg1DxfCu8guF7889E5S5Qt132Z6joRQurB6dk","mana":3730224.470097488},
			{"shortNodeID":"44MfjVURih4","nodeID":"2EN97HVBQ4gTDD3m6XWvEw2rmXbE7GLrhTwsy7Ns4NaJ","mana":2999930.654804991},
			{"shortNodeID":"NhvvaBgvLvj","nodeID":"9jZtGuDGA9BXEFxGcvN193p1Evpuzy4wgWefe1VsNgyF","mana":2999918.121554702},
			{"shortNodeID":"ZSMrGhr5tj","nodeID":"E429B31QddGUwrTa2d6ecMoHJbja4u91GYpUMnD9VsV","mana":2805816.2366705523},
			{"shortNodeID":"9p5cWhXtUfY","nodeID":"4YeZsLeekQ2vPCtSo4hJHun68UkgABjMFZphKuP3gWKq","mana":2011249.2621057085},
			{"shortNodeID":"jT2c8brAQFL","nodeID":"J5fmWepv9wwZQNuJiDJFrrPkexAQUX4X25ZKY6aXj2wD","mana":1507156.8842214583},
			{"shortNodeID":"RTSK169ZJFp","nodeID":"AqkT1QXEesBKj1rngXjAYPjYqQw6qBk7ugdyWjjnKDvi","mana":1009479.2314520499},
			{"shortNodeID":"aQwCSQLDseN","nodeID":"ESnVDjyeSNnMjvJNqZJC33eT9ijep5rf5U5tRZQECb6e","mana":1005316.2171513195},
			{"shortNodeID":"7Fu74WKw8Ym","nodeID":"3X2KNEwgkpSp2K89AUafA1DqDL2qiTufS9WnG7gLLp1q","mana":1001077.6909874098},
			{"shortNodeID":"hf6bArVfYtH","nodeID":"HMrDkESPvZLpSKy9cFFViE27126zt7MKtrsLq4ExT8Cm","mana":999974.9112577096},
			{"shortNodeID":"bR34o9MfSNV","nodeID":"ErATA7RtKdfRdNs8XyBiY82d7H6VbUm5Q3Yxua1fhXkp","mana":999974.4364983559},
			{"shortNodeID":"2rnYKpjjtTS","nodeID":"kNVMYvF8eDjwqenkvjdZn4wU754P37zYnkyPMTME7s5","mana":999973.3540146283},
			{"shortNodeID":"Jn56ZVMMC7D","nodeID":"89t7UL8d3zqJR5csQ8Y1dgLMPUYrBMJYTgZ27sAWGHDu","mana":999971.1868171847},
			{"shortNodeID":"YZURQA2U2Gs","nodeID":"DhYYmbBYAr3kEWBK2kQzUoJeTfAJHTqdLNDt3L7bza35","mana":999969.3562922529},
			{"shortNodeID":"QMT7Zp6tWud","nodeID":"AQ19971smfm6tUZE7SWSvE3Vw9r4VWokesauBc4JdjJM","mana":999953.6209828657},
			{"shortNodeID":"182cARyUjDQ","nodeID":"13q1i7wvYBxgfw3Uf4J4KEYD2s53ahyG8iPwxQxHRyor","mana":999937.990505039},
			{"shortNodeID":"jjzEvpsZEm","nodeID":"JCVZZMDojEq6Qyss9BQx5XFqWGEevtqEFdg6xVWeby2","mana":999826.7900317154},
			{"shortNodeID":"4bNcLWrSQNz","nodeID":"2Sqzn9otyCHB2eu4qZAH4nDEhn7sW7EwJ16iYXKHSFMU","mana":999801.7019598096},
			{"shortNodeID":"hctQeVUE83v","nodeID":"HLxeFRmsuiitsY1R1ErAAfJSWaHWU8jsmfu6F5XedJES","mana":951390.874799818},
			{"shortNodeID":"XSqqpEAuZg4","nodeID":"DFYCV44JLMWGcjFGsU477M9QNTCYtvoikS5AMcSDzeRA","mana":925237.0485005786},
			{"shortNodeID":"fX4yTtQK2yu","nodeID":"GVwfFDnezWJyfuoKtww1f5LFNs3m9RtdDNqXeR5TEbpN","mana":603577.9644285233},
			{"shortNodeID":"8Fz9QhaZ8cM","nodeID":"3vPwzAf9ig8gieaA6F3wGENruzXYCt5MDhYc2wCpeDBt","mana":355127.79413727345},
			{"shortNodeID":"Na8YM2ruWTz","nodeID":"9gRrZXwercvv2iCio6RDamPniihXhgPkhMy1wawGZSjX","mana":235255.44190906134},
			{"shortNodeID":"Suy14QHZh5U","nodeID":"BRmBw4d2p8VUXYkytWgzRWamW229PMfcCjKwuF16sErC","mana":200920.52645307814},
			{"shortNodeID":"ZSMrGhr5tj","nodeID":"E429B31Qddbr3AjraG7vCbd9KiLqkNrvfyp7wTkvLfS","mana":194029.5364318932},
			{"shortNodeID":"DXHGTs1jDZz","nodeID":"63EvYNYnJrVipSygRKZ1pTFoUdPt93b6yc3qEknmr4dE","mana":189976.05495598097},
			{"shortNodeID":"7YzQ89gjWAo","nodeID":"3duCBD1igTuVrFgrgtAi33Ptn7jx2U1DhEnFM7nynhiu","mana":168374.09538069944},
			{"shortNodeID":"V1WRaM931FV","nodeID":"CGfjG4dkCSQQtxKUwojE6wDR55yCeoEXTCSRAdSPJCnx","mana":161633.97903636468},
			{"shortNodeID":"NGHE9CQST8v","nodeID":"9ZFGRPn2wTpfR5LNM5jpDx5gniDg4fJFfuNArdQicQCk","mana":153019.5347600506},
			{"shortNodeID":"VtHYihzvmmH","nodeID":"Cd6nWAh67FpFF3KkQSaMb1qHAZXKJ8jQQqyEXmFpis3C","mana":140951.15537180533},
			{"shortNodeID":"hxCJqHhMU3P","nodeID":"HUjGTJvT3EkJEn4JybgAV9qH8QpxR6VvYgbzaf4Axpzj","mana":100109.10226571765},
			{"shortNodeID":"RT4umhDZLSk","nodeID":"AqbqZaMJbCHXd6FzpbNRGszNov3X7miXXq6CFQDnQLd1","mana":99023.49119892197},
			{"shortNodeID":"FMt2UayAXVu","nodeID":"6n8jgQxZGEZj2SHwJqkkzNXR4gVy8Q9jouQCuu5s58cP","mana":52478.42922540393},
			{"shortNodeID":"fTWmNH8t7MR","nodeID":"GUWgkeqR9qTij3SLZYo7qXD53fPQTCbhjZYx4cBsXpw2","mana":49606.43040442365},
			{"shortNodeID":"BCvbJ4epVcw","nodeID":"57BKCF5svY9VQwWouCUjj2tGuJ5bfsZM4vgBJS45fZkr","mana":49577.950398410554},
			{"shortNodeID":"MMKqsdYHyPJ","nodeID":"9BwQWKQrnyNgrZjxXUCe2NLaFxjnUmjKA87dLQmVvtWx","mana":48581.98171131777},
			{"shortNodeID":"XgoLarDPuUS","nodeID":"DM9vVStRddvwEi6C1NMGyKr89XA1Y6qCgtsok6n8aSuQ","mana":41933.292590775556},
			{"shortNodeID":"UvyfxqDzsdp","nodeID":"CEr1ACrhi5JGvsRNXqNvKVMrHSuEGGiP9qmjME4XzPVY","mana":22432.834095985647},
			{"shortNodeID":"eF2kimrLapj","nodeID":"Fz9kokNjd5qNhcqf6cF2Ep7s9fck6QgTDzMZUxWXjoXk","mana":10743.6143014908},
			{"shortNodeID":"AsfVZkteKXY","nodeID":"4yRpMmtorxMUwPvbLgq8dZNPFW2aYUDXLHN6PmWtP31q","mana":9741.42615482147},
			{"shortNodeID":"8GdLWgUE4jV","nodeID":"3veuxcJjYgQ63TbJ1keMockJXZdbdirPAXmpnPpaEahX","mana":9677.224824960898},
			{"shortNodeID":"BN1LHW2as3","nodeID":"5AqGGpowRZWXsFUngca6LiK4nmi4ogaxoQjvgaQ1LXt","mana":9329.966647611183},
			{"shortNodeID":"7gu9QnQZqK4","nodeID":"3h5nPh6Se5DyAMLMYLX7Wuj2Teumyy9u3qwy8ddrtr3G","mana":9217.994702158883},
			{"shortNodeID":"gcfGdNK1XT4","nodeID":"GwXkeKdd913yPXumS2RXxNasK8pESX12fc6q2KEM69zx","mana":8416.743528625935},
			{"shortNodeID":"fScFud34GkK","nodeID":"GU9ZKm2M3WgvYLMCQozAc5Evu4icLZSYe7gSSVhUwWEU","mana":4790.139754486423},
			{"shortNodeID":"335ajqscnEy","nodeID":"pWj6zX2NnYDu5SC7pK8Mk7SZGATtzc5T9gqk1z6Hw17","mana":2490.1229471848233},
			{"shortNodeID":"DYUCd5mQLZh","nodeID":"63ifKMFPpGSBymgECNhFZoXB7QKcuW6NxhJDMwqbL5X4","mana":2073.874060177712},
			{"shortNodeID":"3Tcpzzk2RSx","nodeID":"zPQo9jN4LgyM6XwL6od7wRJYJzFphLEqCudRp2zCYKJ","mana":1446.440246868049},
			{"shortNodeID":"NzsEyHTVvi6","nodeID":"9rP9ac57s8dEv1DkSdJ5Cc9FewtjDrqGGoV9KVkdKjpo","mana":1386.0610848991862},
			{"shortNodeID":"4GeazEE7AVC","nodeID":"2KK1BbStuN5zGJY8DMGfGX8DXEKMvtvmnGW4F3tvVwng","mana":1280.8167300829773},
			{"shortNodeID":"5X4caCJd74y","nodeID":"2pT1MyDbabb9cZwqjmxUQRrXNjoR7SFnCbEJmuXPepnv","mana":1007.033585161805},
			{"shortNodeID":"aoXHDFV73nr","nodeID":"EbsduJuWX9usWtQQMA9EKfgZ2HoJTxzPmKSnBG5se3kL","mana":969.5310985198116},
			{"shortNodeID":"ereMcqJyYgC","nodeID":"GEUunm7YW5Ri3mn118pitBNyM3DFgUzsztBtkGUceid6","mana":922.2630079160228},
			{"shortNodeID":"WWvXNgE7YMz","nodeID":"CsqqBXxRECn7oSvv7NH8PbC3botfwEdwqqLj1G3MzPAJ","mana":911.6565133397212},
			{"shortNodeID":"MfcbS4d65qN","nodeID":"9KJEJGxyuppqfRRu2tuwEkvJF4nE9JX9sGE9aYp4hD6q","mana":724.2822543427093},
			{"shortNodeID":"ccHe8UYKxQx","nodeID":"FL2dvRke1EBMp6HxS2wAUQM7nb9TBetfiLTmKwaDbfA1","mana":724.1884092579478},
			{"shortNodeID":"8hpVpEBTwms","nodeID":"46nrVxnZtK9o4WLkoiVXbgQLa8tGVFvo4953bkDE7LLf","mana":602.8124108112142},
			{"shortNodeID":"HZv7BmXcVwr","nodeID":"7feqTaM2qeVSc6LQMyYGE3o6ZvspVLy11tWbjjK2GG2i","mana":409.4482247787799},
			{"shortNodeID":"gk4m4qyP97H","nodeID":"GzWa5oWRVcPMgFMP5TAZHK8z3n64Rzez6hJEGXDWABNx","mana":341.59795332503626},
			{"shortNodeID":"QctgKTXCHSV","nodeID":"AWDVpL9nwJSqFWBtuhSsiT7V6ddYas6RDAS4pTJ3M5SA","mana":299.98964727821823},
			{"shortNodeID":"4V3peLMv2rZ","nodeID":"2QJQh1YrtTSQz5qz7jpWdFGNHAq96hbjv6fQR3h6ntcf","mana":200.26873865442838},
			{"shortNodeID":"MN99cRjkxGL","nodeID":"9CGSXdjCKk8ZU3PFjiNasp8HXnmKZnuzkpLM1WmFqp9M","mana":182.36736953925586},
			{"shortNodeID":"8JfPR3fL1sP","nodeID":"3wUR1XoSv79vNEBsGGztZAjNV3HjFwT4or9ox5SzYz6T","mana":85.76033298570454},
			{"shortNodeID":"A21e6DFtfn9","nodeID":"4dT2PZjPP1N4oMVEQfYwDPxu9KGPwMDmQt1fzd5mvUiD","mana":46.99062825702913},
			{"shortNodeID":"DzKwTzbnFC4","nodeID":"6E88CQ2WrHvRV2gh29aPsso8sMQr89uUZpKqscUg5sov","mana":38.686880000808905},
			{"shortNodeID":"NFVCvxbgjri","nodeID":"9YvkPPGcKvH7GrFGLtZRGHWn54nzT5rpH8yvBPjnRQXe","mana":32.44075999946826},
			{"shortNodeID":"dwcw1mQXcdg","nodeID":"Fs9RakVn3gwcCrGEjHwM2ingnLrsvoNAteuCD8fsCfUA","mana":18.95291097352776},
			{"shortNodeID":"HCRCKneubhd","nodeID":"7X174KgeR1VjBszbn4k5GVvPZkVcGveS8S2yJ233MF6v","mana":15.010301038202229},
			{"shortNodeID":"XMVFXBKhzxX","nodeID":"DDPDuiRS5eubdwoSJsjaPwAEFSkAPG3vio391dqTaQ3s","mana":0.3159145503006809}
		]
	}`
)

func ParseManaJSON() ConsensusMana {
	var consensusDistribution ConsensusMana

	json.Unmarshal([]byte(manaDistribution), &consensusDistribution)

	return consensusDistribution
}
