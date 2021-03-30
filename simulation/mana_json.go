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
		"consensus": [
    {
      "shortNodeID": "11111111",
      "nodeID": "11111111111111111111111111111111",
      "mana": 788281304652010.5
    },
    {
      "shortNodeID": "BaNaDA1EumS",
      "nodeID": "5FosDHpsHurmB9rWDBp3VDV9iVz8M7JjxtPYG3EtjF2Q",
      "mana": 266885241.23834494
    },
    {
      "shortNodeID": "KnyxQBcDNtK",
      "nodeID": "8Zaz5xcFu3TPcqAEjYDd5x7nc3s5EWGAmcPskP5tUacV",
      "mana": 206879469.7175765
    },
    {
      "shortNodeID": "7GQ62hsvdAj",
      "nodeID": "3XDykfsEZyBfvRqnxoQ2ui9wVFKyVmyjRZbwsgxLXV3o",
      "mana": 69130677.5535854
    },
    {
      "shortNodeID": "fX4yTtQK2yu",
      "nodeID": "GVwfFDnezWJyfuoKtww1f5LFNs3m9RtdDNqXeR5TEbpN",
      "mana": 29931397.846512526
    },
    {
      "shortNodeID": "YVPUvRzKu1b",
      "nodeID": "DfuD5qCQ6e21DtQfCywSLNzpGVfJG4zFd9dgeMvaC2y5",
      "mana": 19075643.86954299
    },
    {
      "shortNodeID": "brEM773P27x",
      "nodeID": "F2JSWzn7NHQWaM6niPgnGaKTi8bU3vi7gwwaDPpzpbny",
      "mana": 16503969.173613228
    },
    {
      "shortNodeID": "2AatmXGV8qg",
      "nodeID": "UCHQ2agVGCB7hQSrdnGFc7ZwuSsqgJqRS5pmYjATJ5z",
      "mana": 11466922.022501556
    },
    {
      "shortNodeID": "7cmX7HB2JKm",
      "nodeID": "3fRN1KYYBJFN47GDZtvxN9Yxd3aB6AP79KFEsuz4mxob",
      "mana": 10335583.37696328
    },
    {
      "shortNodeID": "2unmZAZkvNP",
      "nodeID": "mabVGTXJpxDRro5U9QXDqW4FQjrCzKBFUnrPZqLxJe2",
      "mana": 7790877.077371529
    },
    {
      "shortNodeID": "EcfmjmpBzgm",
      "nodeID": "6UkGksrpd5sMDfphNpsbWXPJQ2V5VD4e1uDFtAdJ4cay",
      "mana": 7038641.49615559
    },
    {
      "shortNodeID": "VNKR12HbzFT",
      "nodeID": "CR3PyhLvzTueMTGmQXrYSDA5nyGU5WkYcMugE659LtpW",
      "mana": 6820694.565239315
    },
    {
      "shortNodeID": "NYGi8UA4Mwo",
      "nodeID": "9fgTziRHLZHbWXhXHj9T8bdsaZYoUjzvxXNLqYywzsRv",
      "mana": 4922055.349679017
    },
    {
      "shortNodeID": "BNJsNaeBWCQ",
      "nodeID": "5AxKXbnzxjbWhCB1isj8BCBXRZ9ih4KqofsUvRj8mHQS",
      "mana": 4715928.3279479975
    },
    {
      "shortNodeID": "Y16XTksCde2",
      "nodeID": "DUWvUHhRskuaPnNAp8K6vgFA6atEJatpp1VyeuXWPgWU",
      "mana": 4425763.682309439
    },
    {
      "shortNodeID": "i3i5EDo2E7p",
      "nodeID": "HWww8pmDRQqGPKLfwcXVpJfNYBxtZZn8kWdTHYgvVt8P",
      "mana": 4337580.170796012
    },
    {
      "shortNodeID": "YYYAZDPVjc5",
      "nodeID": "DhAiYbMouPGEyp9rLshbR6xrhCYX13wWakGMVMY5Y4nK",
      "mana": 4228827.641216035
    },
    {
      "shortNodeID": "3XN7bsS88CA",
      "nodeID": "21tqpUwDvmEngz9AsqvUqN398q2Qha4whT5tNvW6VshU",
      "mana": 2719598.8485213583
    },
    {
      "shortNodeID": "hm9Y9fE9XS9",
      "nodeID": "HQHSATMzT5UpizCMoPec96uK3DStSHiT9bGmhaFhPbVQ",
      "mana": 2321955.316295727
    },
    {
      "shortNodeID": "dg1fJkmvLKz",
      "nodeID": "FksAPJFGvvm1EZUtFBvNUWwtQVuaBcmsCdFpoG5PpSKc",
      "mana": 2290702.522444421
    },
    {
      "shortNodeID": "81m376rur1o",
      "nodeID": "3pfwcS1A3xtVEMXj6kb2u62UY7XST22Hw9Bovkhd2oZ2",
      "mana": 2273001.0076718773
    },
    {
      "shortNodeID": "9eXQTGkYXPu",
      "nodeID": "4UoZmjwWeT5a2CMogu5s6jqtigVRN5qZcJ9sAByuMUAv",
      "mana": 2214008.475601943
    },
    {
      "shortNodeID": "feUWpUJLgmw",
      "nodeID": "GYvVrkHv4vDLQNzYJpZcoizigh3EJiZiQfhX4nofLJVw",
      "mana": 2068892.2589117277
    },
    {
      "shortNodeID": "gUrmwanvExy",
      "nodeID": "GtPgLDC8ZUHEP1HmZ3tixA6p2g4qTkq3HfZukCR5acji",
      "mana": 1778261.8521135338
    },
    {
      "shortNodeID": "hmpzL2SugVN",
      "nodeID": "HQZJswN6EacbP55E65bBKJ9cEDRfnvbYSX6R2Yfhhk3v",
      "mana": 1595032.0921376052
    },
    {
      "shortNodeID": "aZ6qhPVRccg",
      "nodeID": "EW555akrNkzDDwxVv6WyeL3SRVruVcRcRvfyGugsu2LG",
      "mana": 1585561.0822666993
    },
    {
      "shortNodeID": "U4fbc5CMJBo",
      "nodeID": "BtbqAQpJp8gY1GHP3NDXBV5WhtQJvQf77NNvNZc5PwZT",
      "mana": 1574224.5531766466
    },
    {
      "shortNodeID": "HRbn1KjwzTL",
      "nodeID": "7cJkqDY8UZbRyTz2RedvvgWawVEJXGjKMRMfpfKFnwy2",
      "mana": 1554825.3037581649
    },
    {
      "shortNodeID": "2ZErA3PpnGL",
      "nodeID": "dJzTf8mrZsr4QHh8GKjNPia7iuWUbjFUZ9o4j6yTzjY",
      "mana": 1537010.995643776
    },
    {
      "shortNodeID": "fok53G1956V",
      "nodeID": "GceoKcnbkSChhPNaSCJCREYCVHVReHL54g1CcAKzgdho",
      "mana": 1530701.2701816922
    },
    {
      "shortNodeID": "4vjWxstv2f1",
      "nodeID": "2adqB8ZjxTmKqzE8PtwVhmUbHsRSvFP5vfCq2i6sGYEB",
      "mana": 1523532.8599157108
    },
    {
      "shortNodeID": "uJzrTQQfYT",
      "nodeID": "N3pPcxETtKoAno9CpwT4mkSZvmXL1BNyWsLqwg3i7nG",
      "mana": 1509660.2719238065
    },
    {
      "shortNodeID": "TKVBFGn6Uwm",
      "nodeID": "BbE73ZPQKtM77teNdjGL7Pp7ovvMEQfHiSxm6HowDYAv",
      "mana": 1503430.1602826454
    },
    {
      "shortNodeID": "dW9yPo7mTCG",
      "nodeID": "Fgu8LPXZho3yeT9TWNMjq9jn7w3oXGem2G7kGi627cy5",
      "mana": 1466038.749216622
    },
    {
      "shortNodeID": "AEDxsCtc8Xn",
      "nodeID": "4iN3EomZvb4Up3MCevzx7UE3Zj3XmSKNwuy5GqWxa5Yt",
      "mana": 1431942.3243982203
    },
    {
      "shortNodeID": "9eaAk4JA3zA",
      "nodeID": "4UpgGaBdh2PvoXRjPhjfiwfKZhVD9RsEyUuGybd8WX8F",
      "mana": 1248518.4909675089
    },
    {
      "shortNodeID": "fqsADb5dur7",
      "nodeID": "GdWKyey9Decq5HWasr7cfC14haHqQAtWYnnLx4UNWUrk",
      "mana": 937395.3537074422
    },
    {
      "shortNodeID": "Y9EUdUuuizb",
      "nodeID": "DXnprtbSeQ5PtBgjxVxmSpSSE3fHXD1VNn18B8NSLmbj",
      "mana": 898434.8437444883
    },
    {
      "shortNodeID": "jWScXFguJGf",
      "nodeID": "J73SbCSodtoYhXsgb4c6duD8TnDq1qkbtE3NF3wQemKm",
      "mana": 787615.4743021592
    },
    {
      "shortNodeID": "cRQKkVnjMeP",
      "nodeID": "FFecFtfG4ENSJjYj5ciQJFVVSN7XQnd7PZHnX879cap2",
      "mana": 786558.0055866787
    },
    {
      "shortNodeID": "SNqVehZthk2",
      "nodeID": "BDEh2yiKKvzqLzJBg13KyTnwpDibEgrYTmcuEAyaanwf",
      "mana": 785991.0182786227
    },
    {
      "shortNodeID": "JB7n9yizZH4",
      "nodeID": "7upMy8jVJJQCCAx3phoJ2AFjavafMjKRXWDzswEZhCzY",
      "mana": 785433.8311753809
    },
    {
      "shortNodeID": "G6hekFJHG2K",
      "nodeID": "75N6Vb1D6yEm3ZHL2yGXZYkN4UCufHKxYnDRvSNQtu5G",
      "mana": 784576.2317835987
    },
    {
      "shortNodeID": "6fAEfhc3saS",
      "nodeID": "3H3aPMW67zzyTmudpgmYBdRDvEwydJFnqq2YkpfYmq8v",
      "mana": 783270.0851443092
    },
    {
      "shortNodeID": "RJCWHG6NzMF",
      "nodeID": "An2rPKochukaGpSEy33dkDuUvfS6KXwWnbHw9h8Bsydc",
      "mana": 782180.5785520942
    },
    {
      "shortNodeID": "aYBuckpokB8",
      "nodeID": "EVhmk6sv5kha1pWQyPB2jaK2PYcUS62b7VtGopBvBsuV",
      "mana": 781036.6671511423
    },
    {
      "shortNodeID": "jf6Va1cXkSQ",
      "nodeID": "JAXPTjrjzf5zyzsaXyzyBC8dVBjTECstgW2yzokz4gXN",
      "mana": 777766.6043537228
    },
    {
      "shortNodeID": "D582hi7au2i",
      "nodeID": "5riRHBa2ejzcKhU8gvDAANeL8zD5s1q7ou8MCi8rNswz",
      "mana": 776585.8360609584
    },
    {
      "shortNodeID": "f1pb46WLtwb",
      "nodeID": "GJB51DLqTbRrz7g9h6W7At33DtMwj7t5QXd2MtNHsX5Q",
      "mana": 775755.2185556706
    },
    {
      "shortNodeID": "Mxs9m5SzGn7",
      "nodeID": "9SEqHvep3rmU9cRL5UvShRJj6hM1XrJHFqPtVXbp5xsd",
      "mana": 772507.0279957451
    },
    {
      "shortNodeID": "4mZGE2yNT71",
      "nodeID": "2WwfSGR5wFN6P18AkafAwwpZjymfFnsfzz2cHrr5bdrw",
      "mana": 770406.9484480051
    },
    {
      "shortNodeID": "j6Y2mdU6uUN",
      "nodeID": "HwRWZUjZ9fstzSf9bdyogSiUqDybC1bizj49Zz5f4xmK",
      "mana": 766699.6352724289
    },
    {
      "shortNodeID": "MuWSdFpfLaN",
      "nodeID": "9QtVAuNek6k2K8mCqYuihFzvqu1dXeyfTiBNGcoe7S8w",
      "mana": 757087.3346843715
    },
    {
      "shortNodeID": "2KChVaii1Ar",
      "nodeID": "XfPrVF4MF3R55ZRtDHczKJqrWXnzt2tpxhwTmnBZtfj",
      "mana": 754768.8702482661
    },
    {
      "shortNodeID": "RMoFrpSAXKR",
      "nodeID": "AoUrDpiDusjMRmh6WeRm4ACmstwPTYxw2uctgSXw95Rb",
      "mana": 748979.291969006
    },
    {
      "shortNodeID": "D44Y6bf3GP4",
      "nodeID": "5rHg9oC1qCv157B9LYVm7vae9BFBEU5SAbzfzU4aMezt",
      "mana": 744757.2829094419
    },
    {
      "shortNodeID": "ZPnHjBe4z9d",
      "nodeID": "E2yxifN3hrwR3zNUHaLziyejcyEbU77iuf1LQv3jN1Eo",
      "mana": 739564.4627540021
    },
    {
      "shortNodeID": "RnaddBJCfED",
      "nodeID": "AyTDfYZNfdPdmsVhCdUUW9sPasa5wXq5JJbpwCbH7cs7",
      "mana": 727648.5384411326
    },
    {
      "shortNodeID": "RifYaNfvVHu",
      "nodeID": "AwsquobT57CYWaaHqtiGayfZwPtvM9Xbea5ornhwbro9",
      "mana": 724659.1321022711
    },
    {
      "shortNodeID": "T1f45Lu77By",
      "nodeID": "BU3zgWLNYZQzthMqg58CUnbH7UhFDnz6Nxa8dmuLRBnB",
      "mana": 719117.6542685003
    },
    {
      "shortNodeID": "biEv2Mna8Pn",
      "nodeID": "Ey5xwN4KArsUx9UaCydssr4oqxEay8QzJr4hUyAHn1Lf",
      "mana": 711576.530065114
    },
    {
      "shortNodeID": "5xWx79Btzc6",
      "nodeID": "2zh469DZLVsEDTVzm9rhD9HoogQwRxfbEuzFDBahuNLQ",
      "mana": 671903.968189647
    },
    {
      "shortNodeID": "BBCFLcyF6gm",
      "nodeID": "56UwNCEmCa2C95KKvQmkBomDH4roKVRb6p2bF91HbsML",
      "mana": 651636.8667135197
    },
    {
      "shortNodeID": "A1BEDSspbSn",
      "nodeID": "4d7YZQqcjCkx8isitCnqzR8AET667p1rKjEMGKcHoDNP",
      "mana": 632088.5499540408
    },
    {
      "shortNodeID": "2TG3Amm7b76",
      "nodeID": "auScN9gp9jHGMbojZxKHdCPSFD5aS5shAPCt64G8sqV",
      "mana": 622006.6689298473
    },
    {
      "shortNodeID": "W9jXwRgCcwc",
      "nodeID": "CjKK2Q9JB9tZKj4LmTFKTiN7DfhKojx641qHJJRFLaVF",
      "mana": 471728.84443061525
    },
    {
      "shortNodeID": "fVpi6o3iM6c",
      "nodeID": "GVSajEbRVV3dA2NH9cdQJdiyNLX4RnbertxSJJjLGgLf",
      "mana": 298295.81613648974
    },
    {
      "shortNodeID": "ZE1Yt8e6ePf",
      "nodeID": "Dy3vBKxg3j9EVPri9NSNovCeQVm3Gx9NdAT8KhYwZUgk",
      "mana": 201285.61606969684
    },
    {
      "shortNodeID": "GqmqfHj2uhG",
      "nodeID": "7NhKB4XuqXUcJrT6UnKVrtoW4JrTb2BbWYF9ekKUka8m",
      "mana": 122687.88184691448
    },
    {
      "shortNodeID": "26kWefhuy3h",
      "nodeID": "SeoTvkdJTAhnRynAsCgYJ1V3SvxBBKX1vHRoP6gqXUU",
      "mana": 81548.5744963303
    },
    {
      "shortNodeID": "ZnVW3vjoL9A",
      "nodeID": "EC7yoenZxRLrJLjSkoNw2yR5VFTcdF9tWisN45d7VG3w",
      "mana": 78393.11131112796
    },
    {
      "shortNodeID": "RQGujf6jA5M",
      "nodeID": "ApUeo3K5iKbmpNfL3SuamAjSzozvk2CDWJZqmwWXudAj",
      "mana": 50517.37244615776
    },
    {
      "shortNodeID": "6bJgqiWnun5",
      "nodeID": "3FVdD4Ncv5vmMQLfbMRGWNDd28vEoDRYBN2ebnme79FK",
      "mana": 18413.48130946957
    },
    {
      "shortNodeID": "C3R3tN98yxR",
      "nodeID": "5SgzJz5UZEtFvnPeboCsqv6qMkcTE4WJ3TXs7hd5qJhg",
      "mana": 3628.1442387511725
    },
    {
      "shortNodeID": "CZ3VPmjqeAD",
      "nodeID": "5ecTEY3sdrDCqxt7hDLBHKfjsG9gFx84XaE5BbQJ3V9Y",
      "mana": 3151.806594827808
    },
    {
      "shortNodeID": "CgacgfuYvUx",
      "nodeID": "5heLsHxMRdTewXooaaDFGpAoj5c41ah5wTmpMukjdvi7",
      "mana": 659.5398648264919
    },
    {
      "shortNodeID": "JLNNRCHRD8J",
      "nodeID": "7yYGv48KVwqgV4uWZA1HX28Jm4u3gpB4VpoTTvbMzUX3",
      "mana": 409.72094617595815
    }
  ]
	}`
)

func ParseManaJSON() ConsensusMana {
	var consensusDistribution ConsensusMana

	json.Unmarshal([]byte(manaDistribution), &consensusDistribution)

	return consensusDistribution
}
