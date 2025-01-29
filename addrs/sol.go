package addrs

import (
	"github.com/dwdwow/goalb/labels"
)

var Sol map[string]labels.AddressLabels = map[string]labels.AddressLabels{

	// ============================== CEX ==============================

	// Binance
	"2ojv9BAiHUrvsm9gxDe7fJSzbNZSJcxZvf8dqmWGHG8S": {
		Address: "2ojv9BAiHUrvsm9gxDe7fJSzbNZSJcxZvf8dqmWGHG8S",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.HOT_WALLET, "solscan: Binance 1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"5tzFkiKscXHK5ZXCGbXZxdw7gTjjD1mBwuoFbhUvuAi9": {
		Address: "5tzFkiKscXHK5ZXCGbXZxdw7gTjjD1mBwuoFbhUvuAi9",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.HOT_WALLET, "solscan: Binance 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM": {
		Address: "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.HOT_WALLET, "solscan: Binance 3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"3yFwqXBfZY4jBVUafQ1YEXw189y2dN3V5KQq9uzBDy1E": {
		Address: "3yFwqXBfZY4jBVUafQ1YEXw189y2dN3V5KQq9uzBDy1E",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.HOT_WALLET, "solscan: Binance 8"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"3gd3dqgtJ4jWfBfLYTX67DALFetjc5iS72sCgRhCkW2u": {
		Address: "3gd3dqgtJ4jWfBfLYTX67DALFetjc5iS72sCgRhCkW2u",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.HOT_WALLET, "solscan: Binance 10"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6QJzieMYfp7yr3EdrePaQoG3Ghxs2wM98xSLRu8Xh56U": {
		Address: "6QJzieMYfp7yr3EdrePaQoG3Ghxs2wM98xSLRu8Xh56U",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.HOT_WALLET, "solscan: Binance 11"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"4f77K3QgVREaoAZ7U6EG1BsQMdPKRjzPbNznzzKT2gEJ": {
		Address: "4f77K3QgVREaoAZ7U6EG1BsQMdPKRjzPbNznzzKT2gEJ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.CEX, labels.CEX_BINANCE, labels.SOL_STAKE, "solscan: Binance Stake Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"3N7s9zXMZ4QqvHQR15t5GNHyqc89KduzMP7423eWiD5g": {
		Address: "3N7s9zXMZ4QqvHQR15t5GNHyqc89KduzMP7423eWiD5g",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.SOL_STAKING_VOTE, "solscan: binance staking Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"DRpbCBMxVnDK7maPM5tGv6MvB3v1sRMC86PZ8okm21hy": {
		Address: "DRpbCBMxVnDK7maPM5tGv6MvB3v1sRMC86PZ8okm21hy",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.SOL_STAKING, "solscan: binance staking"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"Cn4Bs7EBNxcVyEbS2HQaJTrmNHmpqUF2xYQkaxtuCfiA": {
		Address: "Cn4Bs7EBNxcVyEbS2HQaJTrmNHmpqUF2xYQkaxtuCfiA",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.CEX, labels.CEX_BINANCE, labels.SOL_STAKING_STAKE_NET_ADDRESS, "solscan: binance staking StakeNet Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Coinbase
	"H8sMJSCQxfKiFTCfDR3DUMLPwcRbM61LGFJ8N4dK3WjS": {
		Address: "H8sMJSCQxfKiFTCfDR3DUMLPwcRbM61LGFJ8N4dK3WjS",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.HOT_WALLET, "solscan: Coinbase 1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"2AQdpHJ2JpcEgPiATUXjQxA8QmafFegfQwSLWSprPicm": {
		Address: "2AQdpHJ2JpcEgPiATUXjQxA8QmafFegfQwSLWSprPicm",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.HOT_WALLET, "solscan: Coinbase 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9obNtb5GyUegcs3a1CbBkLuc5hEWynWfJC6gjz5uWQkE": {
		Address: "9obNtb5GyUegcs3a1CbBkLuc5hEWynWfJC6gjz5uWQkE",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.HOT_WALLET, "solscan: Coinbase 4"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"beefKGBWeSpHzYBHZXwp5So7wdQGX6mu4ZHCsH3uTar": {
		Address: "beefKGBWeSpHzYBHZXwp5So7wdQGX6mu4ZHCsH3uTar",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.SOL_STAKING_VOTE, "solscan: Coinbase Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6D2jqw9hyVCpppZexquxa74Fn33rJzzBx38T58VucHx9": {
		Address: "6D2jqw9hyVCpppZexquxa74Fn33rJzzBx38T58VucHx9",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.SOL_STAKING_VOTE, "solscan: Coinbase 02 Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"GqcuMuWq4gKeCuCrD8iAjXTozCET2EP6qJXmZDFsSTWK": {
		Address: "GqcuMuWq4gKeCuCrD8iAjXTozCET2EP6qJXmZDFsSTWK",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.SOL_STAKING_VOTE, "solscan: Coinbase 03 Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"XkCriyrNwS3G4rzAXtG5B1nnvb5Ka1JtCku93VqeKAr": {
		Address: "XkCriyrNwS3G4rzAXtG5B1nnvb5Ka1JtCku93VqeKAr",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.SOL_STAKING, "solscan: Coinbase"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"3vxheE5C46XzK4XftziRhwAf8QAfipD7HXXWj25mgkom": {
		Address: "3vxheE5C46XzK4XftziRhwAf8QAfipD7HXXWj25mgkom",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.HOT_WALLET, "solscan: Coinbase Prime"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CW9C7HBwAMgqNdXkNgFg9Ujr3edR2Ab9ymEuQnVacd1A": {
		Address: "CW9C7HBwAMgqNdXkNgFg9Ujr3edR2Ab9ymEuQnVacd1A",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.HOT_WALLET, "solscan: Coinbase 02"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9W3QTgBhkU4Bwg6cwnDJo6eGZ9BtZafSdu1Lo9JmWws7": {
		Address: "9W3QTgBhkU4Bwg6cwnDJo6eGZ9BtZafSdu1Lo9JmWws7",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.HOT_WALLET, "solscan: Coinbase 03"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"GJRs4FwHtemZ5ZE9x3FNvJ8TMwitKTh21yxdRPqn7npE": {
		Address: "GJRs4FwHtemZ5ZE9x3FNvJ8TMwitKTh21yxdRPqn7npE",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_COINBASE, labels.HOT_WALLET, "solscan: Coinbase Hot Wallet"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"4HMkJtMhCJxdDB1cthCDeniHqCkeA5E2NGxw6vDDxoQw": {
		Address: "4HMkJtMhCJxdDB1cthCDeniHqCkeA5E2NGxw6vDDxoQw",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.CEX, labels.CEX_COINBASE, labels.SOL_STAKING_STAKE_NET_ADDRESS, "solscan: Coinbase StakeNet Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6uttVyMD5xcLR6tUd7Lf2NVeR2HH7oQarJxk5fDcPKes": {
		Address: "6uttVyMD5xcLR6tUd7Lf2NVeR2HH7oQarJxk5fDcPKes",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.CEX, labels.CEX_COINBASE, labels.SOL_STAKING_STAKE_NET_ADDRESS, "solscan: Coinbase 02 StakeNet Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Kraken
	"FWznbcNXWQuHTawe9RxvQ2LdCENssh12dsznf4RiouN5": {
		Address: "FWznbcNXWQuHTawe9RxvQ2LdCENssh12dsznf4RiouN5",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_KRAKEN, labels.HOT_WALLET, "solscan: Kraken"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"KRAKEnMdmT4EfM8ykTFH6yLoCd5vNLcQvJwF66Y2dag": {
		Address: "KRAKEnMdmT4EfM8ykTFH6yLoCd5vNLcQvJwF66Y2dag",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_KRAKEN, labels.SOL_STAKING_VOTE, "solscan: Kraken Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"krakeNd6ednDPEXxHAmoBs1qKVM8kLg79PvWF2mhXV1": {
		Address: "krakeNd6ednDPEXxHAmoBs1qKVM8kLg79PvWF2mhXV1",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_KRAKEN, labels.SOL_STAKING, "solscan: Kraken"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"E5THh1zEV5HRLJrmJ1MMmfE9ZpZiqbQmXXpeYZZjM9x6": {
		Address: "E5THh1zEV5HRLJrmJ1MMmfE9ZpZiqbQmXXpeYZZjM9x6",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.CEX, labels.CEX_KRAKEN, labels.SOL_STAKING_STAKE_NET_ADDRESS, "solscan: Kraken StakeNet Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Okx
	"9un5wqE3q4oCjyrDkwsdD48KteCJitQX5978Vh7KKxHo": {
		Address: "9un5wqE3q4oCjyrDkwsdD48KteCJitQX5978Vh7KKxHo",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_OKX, labels.HOT_WALLET, "solscan: OKX 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"2tucttroqFNXsrYeMBQ8LfzKNfgwT2rHBzAF6RzbbHEp": {
		Address: "2tucttroqFNXsrYeMBQ8LfzKNfgwT2rHBzAF6RzbbHEp",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_OKX, labels.SOL_STAKING_VOTE, "solscan: OKX Earn Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"5VCwKtCXgCJ6kit5FybXjvriW3xELsFDhYrPSqtJNmcD": {
		Address: "5VCwKtCXgCJ6kit5FybXjvriW3xELsFDhYrPSqtJNmcD",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_OKX, labels.HOT_WALLET, "solscan: OKX"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"Fc6NNdS2j3EmrWbU6Uqt6wsKB5ef72NjaWfNxKYbULGD": {
		Address: "Fc6NNdS2j3EmrWbU6Uqt6wsKB5ef72NjaWfNxKYbULGD",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_OKX, labels.SOL_STAKING, "solscan: OKX Earn"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"5SEvDSqpp1eAT67apSHPyox2LC38DBqvNSb7PYW8VJ1x": {
		Address: "5SEvDSqpp1eAT67apSHPyox2LC38DBqvNSb7PYW8VJ1x",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.CEX, labels.CEX_OKX, labels.SOL_STAKING_STAKE_NET_ADDRESS, "solscan: OKX Earn StakeNet Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6m2CDdhRgxpH4WjvdzxAYbGxwdGUz5MziiL5jek2kBma": {
		Address: "6m2CDdhRgxpH4WjvdzxAYbGxwdGUz5MziiL5jek2kBma",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.CEX_OKX, labels.DEX, "solscan: OKX DEX: Aggregation Router V2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Bitget
	"A77HErqtfN1hLLpvZ9pCtu66FEtM8BveoaKbbMoZ4RiR": {
		Address: "A77HErqtfN1hLLpvZ9pCtu66FEtM8BveoaKbbMoZ4RiR",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BITGET, labels.HOT_WALLET, "solscan: Bitget Exchange"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Gateio
	"HiRpdAZifEsZGdzQ5Xo5wcnaH3D2Jj9SoNsUzcYNK78J": {
		Address: "HiRpdAZifEsZGdzQ5Xo5wcnaH3D2Jj9SoNsUzcYNK78J",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_GATEIO, labels.HOT_WALLET, "solscan: Gate.io 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"u6PJ8DtQuPFnfmwHbGFULQ4u4EgjDiyYKjVEsynXq2w": {
		Address: "u6PJ8DtQuPFnfmwHbGFULQ4u4EgjDiyYKjVEsynXq2w",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_GATEIO, labels.HOT_WALLET, "solscan: Gate.io"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Bybit
	"8hPk5CbKDoM7dN9LssTdVkFhDykeq7A8CZurA5AQSFJH": {
		Address: "8hPk5CbKDoM7dN9LssTdVkFhDykeq7A8CZurA5AQSFJH",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BYBIT, labels.SOL_STAKING_VOTE, "solscan: Bybit Staking Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"AC5RDfQFmDS1deWZos921JfqscXdByf8BKHs5ACWjtW2": {
		Address: "AC5RDfQFmDS1deWZos921JfqscXdByf8BKHs5ACWjtW2",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BYBIT, labels.HOT_WALLET, "solscan: Bybit"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CG4tRANBKrzUmpv93V5sgftjQznBdiJsc2yPCzZWWuS9": {
		Address: "CG4tRANBKrzUmpv93V5sgftjQznBdiJsc2yPCzZWWuS9",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BYBIT, labels.SOL_STAKING, "solscan: Bybit Staking"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Kucoin
	"57vSaRTqN9iXaemgh4AoDsZ63mcaoshfMK8NP3Z5QNbs": {
		Address: "57vSaRTqN9iXaemgh4AoDsZ63mcaoshfMK8NP3Z5QNbs",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_KUCOIN, labels.HOT_WALLET, "solscan: Kucoin"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"BmFdpraQhkiDQE6SnfG5omcA1VwzqfXrwtNYBwWTymy6": {
		Address: "BmFdpraQhkiDQE6SnfG5omcA1VwzqfXrwtNYBwWTymy6",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_KUCOIN, labels.HOT_WALLET, "solscan: Kucoin"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"HVh6wHNBAsG3pq1Bj5oCzRjoWKVogEDHwUHkRz3ekFgt": {
		Address: "HVh6wHNBAsG3pq1Bj5oCzRjoWKVogEDHwUHkRz3ekFgt",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_KUCOIN, labels.HOT_WALLET, "solscan: Kucoin"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	// hold many tokens and rank top, just transfer with kucoin hot wallet, but not be labeled by solscan
	"GL8T72PKygWaYrLKtSteN9UUvYYnuC8azNDunk4eaqqQ": {
		Address: "GL8T72PKygWaYrLKtSteN9UUvYYnuC8azNDunk4eaqqQ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_KUCOIN, labels.HOT_WALLET},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Huobi
	"88xTWZMeKfiTgbfEmPLdsUCQcZinwUfk25EBQZ21XMAZ": {
		Address: "88xTWZMeKfiTgbfEmPLdsUCQcZinwUfk25EBQZ21XMAZ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_HUOBI, labels.HOT_WALLET, "solscan: Huobi"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Mexc
	"ASTyfSima4LLAdDgoFGkgqoKowG1LZFDr9fAQrg7iaJZ": {
		Address: "ASTyfSima4LLAdDgoFGkgqoKowG1LZFDr9fAQrg7iaJZ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_MEXC, labels.HOT_WALLET, "solscan: MEXC"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"5PAhQiYdLBd6SVdjzBQDxUAEFyDdF5ExNPQfcscnPRj5": {
		Address: "5PAhQiYdLBd6SVdjzBQDxUAEFyDdF5ExNPQfcscnPRj5",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_MEXC, labels.HOT_WALLET, "solscan: MEXC #2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Ftx
	"6b4aypBhH337qSzzkbeoHWzTLt4DjG2aG8GkrrTQJfQA": {
		Address: "6b4aypBhH337qSzzkbeoHWzTLt4DjG2aG8GkrrTQJfQA",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_FTX, labels.HOT_WALLET, "solscan: FTX Cold Storage #1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9uyDy9VDBw4K7xoSkhmCAm8NAFCwu4pkF6JeHUCtVKcX": {
		Address: "9uyDy9VDBw4K7xoSkhmCAm8NAFCwu4pkF6JeHUCtVKcX",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_FTX, labels.HOT_WALLET, "solscan: FTX Cold Storage #2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6wEMcwrcF5AP9jpHWQcPxHXciWA2g217Qq81CTWjbgBw": {
		Address: "6wEMcwrcF5AP9jpHWQcPxHXciWA2g217Qq81CTWjbgBw",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_FTX, labels.HOT_WALLET, "solscan: FTX Cold Storage #3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"H4yiPhdSsmSMJTznXzmZvdqWuhxDRzzkoQMEWXZ6agFZ": {
		Address: "H4yiPhdSsmSMJTznXzmZvdqWuhxDRzzkoQMEWXZ6agFZ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_FTX, labels.SOL_STAKING, "solscan: FTX/Alameda Staking Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6ZRCB7AAqGre6c72PRz3MHLC73VMYvJ8bi9KHf1HFpNk": {
		Address: "6ZRCB7AAqGre6c72PRz3MHLC73VMYvJ8bi9KHf1HFpNk",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_FTX, labels.HOT_WALLET, "solscan: FTX"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"JBpj7yp4Afvb71TmanVwJZXGeX4kqbGFvjCFCRo3EbTM": {
		Address: "JBpj7yp4Afvb71TmanVwJZXGeX4kqbGFvjCFCRo3EbTM",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_FTX, labels.HOT_WALLET, "solscan: FTX.US"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Bitfinex
	"FxteHmLwG9nk1eL4pjNve3Eub2goGkkz6g6TbvdmW46a": {
		Address: "FxteHmLwG9nk1eL4pjNve3Eub2goGkkz6g6TbvdmW46a",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BITFINEX, labels.HOT_WALLET, "solscan: Bitfinex"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Crypto.com
	"6FEVkH17P9y8Q9aCkDdPcMDjvj7SVxrTETaYEm8f51Jy": {
		Address: "6FEVkH17P9y8Q9aCkDdPcMDjvj7SVxrTETaYEm8f51Jy",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_CRYPTOCOM, labels.HOT_WALLET, "solscan: Crypto.com 1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"AobVSwdW9BbpMdJvTqeCN4hPAmh4rHm7vwLnQ5ATSyrS": {
		Address: "AobVSwdW9BbpMdJvTqeCN4hPAmh4rHm7vwLnQ5ATSyrS",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_CRYPTOCOM, labels.HOT_WALLET, "solscan: Crypto.com 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	// hold many tokens and rank top, get tokens from crypto.com hot wallet, but not be labeled by solscan
	"HqqzZC5qsNKtZBrBsYvbXhzxUQaADSzkDaifHWe1GvEH": {
		Address: "HqqzZC5qsNKtZBrBsYvbXhzxUQaADSzkDaifHWe1GvEH",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_CRYPTOCOM, labels.HOT_WALLET},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// ============================== MMK ==============================

	// Wintermute
	"7sMiW38oLg5q9SKNoyPaH2Ee1VhpGdMCDrC4Lo4uLBBE": {
		Address: "7sMiW38oLg5q9SKNoyPaH2Ee1VhpGdMCDrC4Lo4uLBBE",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.MMK_WINTERMUTE, labels.HOT_WALLET, "solscan: Wintermute 1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"4FDKx3S3k9eD7HeAhjQxHeYNLXHtreCD1GTUWktiYUvR": {
		Address: "4FDKx3S3k9eD7HeAhjQxHeYNLXHtreCD1GTUWktiYUvR",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.MMK_WINTERMUTE, labels.HOT_WALLET, "solscan: Wintermute 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"5sTQ5ih7xtctBhMXHr3f1aWdaXazWrWfoehqWdqWnTFP": {
		Address: "5sTQ5ih7xtctBhMXHr3f1aWdaXazWrWfoehqWdqWnTFP",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.MMK_WINTERMUTE, labels.HOT_WALLET, "solscan: Wintermute 3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Jump Crypto
	"AhFjTUE2DgNFnAfRtFLUmLTLYyhxzz7j1cvKbvP18tg9": {
		Address: "AhFjTUE2DgNFnAfRtFLUmLTLYyhxzz7j1cvKbvP18tg9",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.MMK_JUMP_CRYPTO, labels.HOT_WALLET, "solscan: Jump Crypto"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// ============================== DEX ==============================

	// Jupiter
	"8gMBNeKwXaoNi9bhbVUWFt4Uc5aobL9PeYMXfYDMePE2": {
		Address: "8gMBNeKwXaoNi9bhbVUWFt4Uc5aobL9PeYMXfYDMePE2",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, labels.MULTISIG, "solscan: Jupiter Burn Multisig"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"FVhQ3QHvXudWSdGix2sdcG47YmrmUxRhf3KCBmiKfekf": {
		Address: "FVhQ3QHvXudWSdGix2sdcG47YmrmUxRhf3KCBmiKfekf",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, labels.HOT_WALLET, "solscan: Jupiter Community Hot Wallet"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CbU4oSFCk8SVgW23NLvb5BwctvWcZZHfxRD6HudP8gAo": {
		Address: "CbU4oSFCk8SVgW23NLvb5BwctvWcZZHfxRD6HudP8gAo",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, labels.HOT_WALLET, "solscan: Jupiter Team Hot Wallet"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"61aq585V8cR2sZBeawJFt2NPqmN7zDi1sws4KLs5xHXV": {
		Address: "61aq585V8cR2sZBeawJFt2NPqmN7zDi1sws4KLs5xHXV",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, labels.COLD_WALLET, "solscan: Jupiter Team Cold Wallet"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"EXJHiMkj6NRFDfhWBMKccHNwdSpCT7tdvQeRf87yHm6T": {
		Address: "EXJHiMkj6NRFDfhWBMKccHNwdSpCT7tdvQeRf87yHm6T",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, labels.COLD_WALLET, "solscan: Jupiter Community Cold Wallet"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"45ruCyfdRkWpRNGEqWzjCiXRHkZs8WXCLQ67Pnpye7Hp": {
		Address: "45ruCyfdRkWpRNGEqWzjCiXRHkZs8WXCLQ67Pnpye7Hp",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Partner Referral Fee Vault"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CatzoSMUkTRidT5DwBxAC2pEtnwMBTpkCepHkFgZDiqb": {
		Address: "CatzoSMUkTRidT5DwBxAC2pEtnwMBTpkCepHkFgZDiqb",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, labels.SOL_STAKING_VOTE, "solscan: Jupiter Validator Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"H4ND9aYttUVLFmNypZqLjZ52FYiGvdEB45GmwNoKEjTj": {
		Address: "H4ND9aYttUVLFmNypZqLjZ52FYiGvdEB45GmwNoKEjTj",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Perpetuals State"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"Cspp27eGUDMXxPEdhmEXFVRn6Lt1L7xJyALF3nmnWoBj": {
		Address: "Cspp27eGUDMXxPEdhmEXFVRn6Lt1L7xJyALF3nmnWoBj",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter DCA Event Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"DCAKxn5PFNN1mBREPWGdk1RXg5aVH9rPErLfBFEi2Emb": {
		Address: "DCAKxn5PFNN1mBREPWGdk1RXg5aVH9rPErLfBFEi2Emb",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter DCA Keeper 1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"DCAKuApAuZtVNYLk3KTAVW9GLWVvPbnb5CxxRRmVgcTr": {
		Address: "DCAKuApAuZtVNYLk3KTAVW9GLWVvPbnb5CxxRRmVgcTr",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter DCA Keeper 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"DCAK36VfExkPdAkYUQg6ewgxyinvcEyPLyHjRbmveKFw": {
		Address: "DCAK36VfExkPdAkYUQg6ewgxyinvcEyPLyHjRbmveKFw",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter DCA Keeper 3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"D8cy77BBepLMngZx6ZukaTff5hCt1HrWyKk3Hnd9oitf": {
		Address: "D8cy77BBepLMngZx6ZukaTff5hCt1HrWyKk3Hnd9oitf",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Event Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"BQ72nSv9f3PRyRKCBnHLVrerrv37CYTHm5h3s9VSGQDV": {
		Address: "BQ72nSv9f3PRyRKCBnHLVrerrv37CYTHm5h3s9VSGQDV",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Authority 1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"2MFoS3MPtvyQ4Wh4M9pdfPjz6UhVoNbFbGJAskCPCj3h": {
		Address: "2MFoS3MPtvyQ4Wh4M9pdfPjz6UhVoNbFbGJAskCPCj3h",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Authority 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"HU23r7UoZbqTUuh3vA7emAGztFtqwTeVips789vqxxBw": {
		Address: "HU23r7UoZbqTUuh3vA7emAGztFtqwTeVips789vqxxBw",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Authority 3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"3CgvbiM3op4vjrrjH2zcrQUwsqh5veNVRjFCB9N6sRoD": {
		Address: "3CgvbiM3op4vjrrjH2zcrQUwsqh5veNVRjFCB9N6sRoD",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Authority 4"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6LXutJvKUw8Q5ue2gCgKHQdAN4suWW8awzFVC6XCguFx": {
		Address: "6LXutJvKUw8Q5ue2gCgKHQdAN4suWW8awzFVC6XCguFx",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Authority 5"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CapuXNQoDviLvU1PxFiizLgPNQCxrsag1uMeyk6zLVps": {
		Address: "CapuXNQoDviLvU1PxFiizLgPNQCxrsag1uMeyk6zLVps",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Authority 6"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"GGztQqQ6pCPaJQnNpXBgELr5cs3WwDakRbh1iEMzjgSJ": {
		Address: "GGztQqQ6pCPaJQnNpXBgELr5cs3WwDakRbh1iEMzjgSJ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator Authority 7"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"VALaaymxQh2mNy2trH9jUqHT1mTow76wpTcGmSWSwJe": {
		Address: "VALaaymxQh2mNy2trH9jUqHT1mTow76wpTcGmSWSwJe",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter VA"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"JSWX3pKDbj2EdCMu4ei7sPYSpdcwZNyjkDSteoHQ4UH": {
		Address: "JSWX3pKDbj2EdCMu4ei7sPYSpdcwZNyjkDSteoHQ4UH",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Ape Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"JUP4Fb2cqiRUcaTHdrPC8h2gNsA2ETXiPDD33WcGuJB": {
		Address: "JUP4Fb2cqiRUcaTHdrPC8h2gNsA2ETXiPDD33WcGuJB",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator v4"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4": {
		Address: "JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator v6"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"JUP6i4ozu5ydDCnLiMogSckDPpbtr7BJ4FtzYWkb5Rk": {
		Address: "JUP6i4ozu5ydDCnLiMogSckDPpbtr7BJ4FtzYWkb5Rk",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator v1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"JUP2jxvXaqu7NQY1GmNF4m1vodw12LVXYxbFL2uJvfo": {
		Address: "JUP2jxvXaqu7NQY1GmNF4m1vodw12LVXYxbFL2uJvfo",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator v2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"JUP3c2Uh3WA4Ng34tw6kPd2G4C5BB21Xo36Je1s32Ph": {
		Address: "JUP3c2Uh3WA4Ng34tw6kPd2G4C5BB21Xo36Je1s32Ph",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Aggregator v3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"PERPHjGBqRHArX4DySjwM6UJHiR3sWAatqfdBS2qQJu": {
		Address: "PERPHjGBqRHArX4DySjwM6UJHiR3sWAatqfdBS2qQJu",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Labs Perpetuals"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M": {
		Address: "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter DCA program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"jupoNjAxXgZ4rjzxzPMP4oxduvQsQtZzyknqvzYNrNu": {
		Address: "jupoNjAxXgZ4rjzxzPMP4oxduvQsQtZzyknqvzYNrNu",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Limit Order"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"GovaE4iu227srtG2s3tZzB4RmWBzw8sTwrCLZz7kN7rY": {
		Address: "GovaE4iu227srtG2s3tZzB4RmWBzw8sTwrCLZz7kN7rY",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Governance Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"j1o2qRpjcyUwEvwtcfhEQefh773ZgjxcVRry7LDqg5X": {
		Address: "j1o2qRpjcyUwEvwtcfhEQefh773ZgjxcVRry7LDqg5X",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_JUPITER, "solscan: Jupiter Limit Order V2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Raydium
	"8pFhUqCU7Fkxfg2DLytRDf7a9oK4XGtN92PrYwtVQc6G": {
		Address: "8pFhUqCU7Fkxfg2DLytRDf7a9oK4XGtN92PrYwtVQc6G",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Treasury"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"39azUYFWPz3VHgKCf3VChUwbpURdCHRxjWVowf5jUJjg": {
		Address: "39azUYFWPz3VHgKCf3VChUwbpURdCHRxjWVowf5jUJjg",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Pump.fun: Raydium Migration"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"RaydiumJDX8X6om6Fg44xyqz5eukZ9KC3LX61SttLbH": {
		Address: "RaydiumJDX8X6om6Fg44xyqz5eukZ9KC3LX61SttLbH",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_RAYDIUM, labels.SOL_STAKING_VOTE, "solscan: Raydium Validator Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"BihEG2r7hYax6EherbRmuLLrySBuSXx4PYGd9gAsktKY": {
		Address: "BihEG2r7hYax6EherbRmuLLrySBuSXx4PYGd9gAsktKY",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium (RAY) Stake Reward Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"4qD717qKoj3Sm8YfHMSR7tSKjWn5An817nArA6nGdcUR": {
		Address: "4qD717qKoj3Sm8YfHMSR7tSKjWn5An817nArA6nGdcUR",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium (RAY) Stake Pool Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"vvvvbtDs9HsdsE6NskZMnb1RA6muoud1ChQuiF9QhSM": {
		Address: "vvvvbtDs9HsdsE6NskZMnb1RA6muoud1ChQuiF9QhSM",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_RAYDIUM, labels.SOL_STAKING, "solscan: Raydium Validator"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"GThUX1Atko4tqhN2NaiTazWSeFWMuiUvfFnyJyUghFMJ": {
		Address: "GThUX1Atko4tqhN2NaiTazWSeFWMuiUvfFnyJyUghFMJ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Upgrade Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"FPjtrk724AM9GuCCUZg4vqGDjB7L42LNPsYbbVzbZBXz": {
		Address: "FPjtrk724AM9GuCCUZg4vqGDjB7L42LNPsYbbVzbZBXz",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, labels.SOL_STAKING_STAKE_NET_ADDRESS, "solscan: Raydium Validator StakeNet Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"AgJddDJLt17nHyXDCpyGELxwsZZQPqfUsuwzoiqVGJwD": {
		Address: "AgJddDJLt17nHyXDCpyGELxwsZZQPqfUsuwzoiqVGJwD",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium LP Exploiter's Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	// programs
	"EhhTKczWMGQt46ynNeRX1WfeagwwJd7ufHvCDjRxjo5Q": {
		Address: "EhhTKczWMGQt46ynNeRX1WfeagwwJd7ufHvCDjRxjo5Q",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Stake"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9HzJyW1qZsEiSfMUf6L2jo3CcTKAyBmSyKdwQeYisHrC": {
		Address: "9HzJyW1qZsEiSfMUf6L2jo3CcTKAyBmSyKdwQeYisHrC",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium IDO"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK": {
		Address: "CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Concentrated Liquidity"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"routeUGWgWzqBWFcrCfv8tritsqukccJPu3q5GPP3xS": {
		Address: "routeUGWgWzqBWFcrCfv8tritsqukccJPu3q5GPP3xS",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium AMM Routing"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"5Q544fKrFoe6tsEbD7S8EmxGTJYAKtTVhAW5Q5pge4j1": {
		Address: "5Q544fKrFoe6tsEbD7S8EmxGTJYAKtTVhAW5Q5pge4j1",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Authority V4"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"FarmqiPv5eAj3j1GMdMCMUGXqPUvmquZtMy86QH6rzhG": {
		Address: "FarmqiPv5eAj3j1GMdMCMUGXqPUvmquZtMy86QH6rzhG",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium ecosystem farm"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CBuCnLe26faBpcBP2fktp4rp8abpcAnTWft6ZrP5Q4T": {
		Address: "CBuCnLe26faBpcBP2fktp4rp8abpcAnTWft6ZrP5Q4T",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Stake V4"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9KEPoZmtHUrBbhWN1v1KWLMkkvwY6WLtAVUCPRtRjP4z": {
		Address: "9KEPoZmtHUrBbhWN1v1KWLMkkvwY6WLtAVUCPRtRjP4z",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Stake V5"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"5quBtoiQqxF9Jv6KYKctB59NT3gtJD2Y65kdnB1Uev3h": {
		Address: "5quBtoiQqxF9Jv6KYKctB59NT3gtJD2Y65kdnB1Uev3h",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium liquidity pool AMM"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"27haf8L6oxUeXrHrgEgsexjSY5hbVUWEmvv9Nyxg8vQv": {
		Address: "27haf8L6oxUeXrHrgEgsexjSY5hbVUWEmvv9Nyxg8vQv",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Liquidity Pool V3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8": {
		Address: "675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Liquidity Pool V4"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"RVKd61ztZW9GUwhRbbLoYVRE5Xf1B2tVscKqwZqXgEr": {
		Address: "RVKd61ztZW9GUwhRbbLoYVRE5Xf1B2tVscKqwZqXgEr",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium Liquidity Pool V2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CPMMoo8L3F4NbTegBCKVNunggL7H1ZpdTHKxQB5qKP1C": {
		Address: "CPMMoo8L3F4NbTegBCKVNunggL7H1ZpdTHKxQB5qKP1C",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_RAYDIUM, "solscan: Raydium CPMM"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Orca
	"82yxjeMsvaURa4MbZZ7WZZHfobirZYkH1zF8fmeGtyaQ": {
		Address: "82yxjeMsvaURa4MbZZ7WZZHfobirZYkH1zF8fmeGtyaQ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_ORCA, "solscan: Orca Aquafarm"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc": {
		Address: "whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_ORCA, "solscan: Orca Whirlpools"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"DjVE6JNiYqPL2QXyCUUh8rNjHrbz9hXHNYt99MQ59qw1": {
		Address: "DjVE6JNiYqPL2QXyCUUh8rNjHrbz9hXHNYt99MQ59qw1",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_ORCA, "solscan: Orca Token Swap"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP": {
		Address: "9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_ORCA, "solscan: Orca Token Swap V2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Lifinity Protocol
	"2ayMCC4aizr8RGg5ptXYqu8uoxW1whNek1hE1zaAd58z": {
		Address: "2ayMCC4aizr8RGg5ptXYqu8uoxW1whNek1hE1zaAd58z",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_LIFINITY, labels.SOL_STAKING_VOTE, "solscan: Lifinity Protocol Vote Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"RoYFUUD7QD9aQ34UCMcwfye8dC5YvJeXz2J3mmoy5S4": {
		Address: "RoYFUUD7QD9aQ34UCMcwfye8dC5YvJeXz2J3mmoy5S4",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_LIFINITY, labels.SOL_STAKING, "solscan: Lifinity Protocol"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6mjWPU6zWq124pCVxtYg2kagQvRjJ92Z37RnSgRJWPtW": {
		Address: "6mjWPU6zWq124pCVxtYg2kagQvRjJ92Z37RnSgRJWPtW",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_LIFINITY, labels.SOL_STAKING_STAKE_NET_ADDRESS, "solscan: Lifinity Protocol StakeNet Address"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"EewxydAPCCVuNEyrVN68PuSYdQ7wKn27V9Gjeoi8dy3S": {
		Address: "EewxydAPCCVuNEyrVN68PuSYdQ7wKn27V9Gjeoi8dy3S",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_LIFINITY, "solscan: Lifinity Swap"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"2wT8Yq49kHgDzXuPxZSaeLaH1qbmGXtEyPy64bL7aD3c": {
		Address: "2wT8Yq49kHgDzXuPxZSaeLaH1qbmGXtEyPy64bL7aD3c",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_LIFINITY, "solscan: Lifinity Swap V2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Meteora
	"Eo7WjKq67rjJQSZxS6z3YkapzY3eMj6Xy8X5EQVn5UaB": {
		Address: "Eo7WjKq67rjJQSZxS6z3YkapzY3eMj6Xy8X5EQVn5UaB",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_METEORA, "solscan: Meteora Pools Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo": {
		Address: "LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_METEORA, "solscan: Meteora DLMM Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"24Uqj9JCLxUeoC3hGfh5W3s9FM9uCHDS2SG3LYwBpyTi": {
		Address: "24Uqj9JCLxUeoC3hGfh5W3s9FM9uCHDS2SG3LYwBpyTi",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_METEORA, "solscan: Meteora Vault Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"vaU6kP7iNEGkbmPkLmZfGwiGxd4Mob24QQCie5R9kd2": {
		Address: "vaU6kP7iNEGkbmPkLmZfGwiGxd4Mob24QQCie5R9kd2",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_METEORA, "solscan: Meteora DLMM Vault Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Pump.fun
	// "39azUYFWPz3VHgKCf3VChUwbpURdCHRxjWVowf5jUJjg": {
	// 	Address: "39azUYFWPz3VHgKCf3VChUwbpURdCHRxjWVowf5jUJjg",
	// 	Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_PUMPFUN, "solscan: Pump.fun Raydium Migration"},
	// 	Grade:   labels.LabelGradeStrict,
	// 	Remark:  "from solscan",
	// },
	"CebN5WGQ4jvEPvsVU4EoHEpgzq1VV7AbicfhtW4xC9iM": {
		Address: "CebN5WGQ4jvEPvsVU4EoHEpgzq1VV7AbicfhtW4xC9iM",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_PUMPFUN, "solscan: Pump.fun Fee Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"TSLvdd1pWpHVjahSpsvCXUbgwsL3JAcvokwaKt1eokM": {
		Address: "TSLvdd1pWpHVjahSpsvCXUbgwsL3JAcvokwaKt1eokM",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_PUMPFUN, "solscan: Pump.fun Token Mint Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	// Pump.fun Programs
	"6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P": {
		Address: "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_PUMPFUN, "solscan: Pump.fun Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"BANANAesZoESMCSsQDQBwWuEvAUnUbbuEmdSAHggvJVS": {
		Address: "BANANAesZoESMCSsQDQBwWuEvAUnUbbuEmdSAHggvJVS",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_PUMPFUN, "solscan: Pump.fun Router Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Phoenix
	"PhoeNiXZ8ByJGLkxNfZRnkUfjvmuYqLR89jjFHGqdXY": {
		Address: "PhoeNiXZ8ByJGLkxNfZRnkUfjvmuYqLR89jjFHGqdXY",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_PHOENIX, "solscan: Phoenix Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// ApePro
	"JSW99DKmxNyREQM14SQLDykeBvEUG63TeohrvmofEiw": {
		Address: "JSW99DKmxNyREQM14SQLDykeBvEUG63TeohrvmofEiw",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_APEPRO, "solscan: ApePro Smart Wallet Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Fluxbeam
	"beamazjPnFT3JQoe16HjUxkpmHFfsHY6dTqf3VwBXzq": {
		Address: "beamazjPnFT3JQoe16HjUxkpmHFfsHY6dTqf3VwBXzq",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_FLUXBEAM, "solscan: Fluxbeam Fee Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"FLUXubRmkEi2q6K3Y9kBPg9248ggaZVsoSFhtJHSrm1X": {
		Address: "FLUXubRmkEi2q6K3Y9kBPg9248ggaZVsoSFhtJHSrm1X",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_FLUXBEAM, "solscan: Fluxbeam Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Sanctum
	"77pP7iTJKrNDmmHFfUYTA9YxfHP8PHZ4ABxAdwYSTJ8r": {
		Address: "77pP7iTJKrNDmmHFfUYTA9YxfHP8PHZ4ABxAdwYSTJ8r",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Airdrop Vault"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"scoeWYRwSor53KxfQ8EkNCkka1vasF8td3P3nfHQvsv": {
		Address: "scoeWYRwSor53KxfQ8EkNCkka1vasF8td3P3nfHQvsv",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum S Controller"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"6v6B6tUQ9J5D75WejAcq4rYjCAobq8saucATLkw86JnE": {
		Address: "6v6B6tUQ9J5D75WejAcq4rYjCAobq8saucATLkw86JnE",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Airdrop Vault #2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"AYhux5gJzCoeoc1PoJ1VxwPDe22RwcvpHviLDD1oCGvW": {
		Address: "AYhux5gJzCoeoc1PoJ1VxwPDe22RwcvpHviLDD1oCGvW",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Markets"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"FypPtwbY3FUfzJUtXHSyVRokVKG2jKtH29FmK4ebxRSd": {
		Address: "FypPtwbY3FUfzJUtXHSyVRokVKG2jKtH29FmK4ebxRSd",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Unstake Pool"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"ALpzvhALRr35nH8mw9SXk2WvmwEYjfw1dvmpFG9Kosu6": {
		Address: "ALpzvhALRr35nH8mw9SXk2WvmwEYjfw1dvmpFG9Kosu6",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Prefunder Account"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"GRwm4EXMyVwtftQeTft7DZT3HBRxx439PrKq4oM6BwoZ": {
		Address: "GRwm4EXMyVwtftQeTft7DZT3HBRxx439PrKq4oM6BwoZ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Stake Manager"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	// Sanctum Programs
	"5ocnV1qiCgaQR8Jb8xWnVbApfaygJ8tNoZfgPwsgx9kx": {
		Address: "5ocnV1qiCgaQR8Jb8xWnVbApfaygJ8tNoZfgPwsgx9kx",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"stkitrT1Uoy18Dk1fTrgPw8W6MVzoCfYoAFT4MLsmhq": {
		Address: "stkitrT1Uoy18Dk1fTrgPw8W6MVzoCfYoAFT4MLsmhq",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Router Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"f1tUoNEKrDp1oeGn4zxr7bh41eN6VcfHjfrL3ZqQday": {
		Address: "f1tUoNEKrDp1oeGn4zxr7bh41eN6VcfHjfrL3ZqQday",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Flat Fee Pricing"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"unpXTU2Ndrc7WWNyEhQWe4udTzSibLPi25SXv2xbCHQ": {
		Address: "unpXTU2Ndrc7WWNyEhQWe4udTzSibLPi25SXv2xbCHQ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Unstake Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"SPMBzsVUuoHA4Jm6KunbsotaahvVikZs1JyTW6iJvbn": {
		Address: "SPMBzsVUuoHA4Jm6KunbsotaahvVikZs1JyTW6iJvbn",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Multi-Validator SPL Stake Pool Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"SP12tWFxD9oJsVWNavTTBZvMbA6gkAmxtVgxdqvyvhY": {
		Address: "SP12tWFxD9oJsVWNavTTBZvMbA6gkAmxtVgxdqvyvhY",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_SANCTUM, "solscan: Sanctum Single Validator SPL Stake Pool Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// OpenBook Programs
	"srmqPvymJeFKQ4zGQed1GFppgkRHL9kaELCbyksJtPX": {
		Address: "srmqPvymJeFKQ4zGQed1GFppgkRHL9kaELCbyksJtPX",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_OPENBOOK, "solscan: OpenBook Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"opnb2LAfJYbRMAHHvqjCwQxanZn7ReEHp1k81EohpZb": {
		Address: "opnb2LAfJYbRMAHHvqjCwQxanZn7ReEHp1k81EohpZb",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_OPENBOOK, "solscan: OpenBook V2 Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},

	// Kamino Accounts
	"GKnHiWh3RRrE1zsNzWxRkomymHc374TvJPSTv2wPeYdB": {
		Address: "GKnHiWh3RRrE1zsNzWxRkomymHc374TvJPSTv2wPeYdB",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Global Config"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"E35i5qn7872eEmBt15e5VGhziUBzCTm43XCSWvDoQNNv": {
		Address: "E35i5qn7872eEmBt15e5VGhziUBzCTm43XCSWvDoQNNv",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Admin Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"3v6ootgJJZbSWEDfZMA1scfh7wcsVVfeocExRxPqCyWH": {
		Address: "3v6ootgJJZbSWEDfZMA1scfh7wcsVVfeocExRxPqCyWH",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Token Infos"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"3NJYftD5sjVfxSnUdZ1wVML8f3aC6mp1CXCL6L7TnU8C": {
		Address: "3NJYftD5sjVfxSnUdZ1wVML8f3aC6mp1CXCL6L7TnU8C",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Scope Prices"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"B9spsrMK6pJicYtukaZzDyzsUQLgc3jbx5gHVwdDxb6y": {
		Address: "B9spsrMK6pJicYtukaZzDyzsUQLgc3jbx5gHVwdDxb6y",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Reserve 2"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"9DrvZvyWh1HuAoZxvYWMvkf2XCzryCpGgHqrMjyDWpmo": {
		Address: "9DrvZvyWh1HuAoZxvYWMvkf2XCzryCpGgHqrMjyDWpmo",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Reserve 1"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"81BgcfZuZf9bESLvw3zDkh7cZmMtDwTPgkCvYu7zx26o": {
		Address: "81BgcfZuZf9bESLvw3zDkh7cZmMtDwTPgkCvYu7zx26o",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Reserve 3"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"GuWEkEJb5bh8Ai2gaYmZWMTUq8MrFeoaDZ89BrQfB1FZ": {
		Address: "GuWEkEJb5bh8Ai2gaYmZWMTUq8MrFeoaDZ89BrQfB1FZ",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Reserve 4"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"BabJ4KTDUDqaBRWLFza3Ek3zEcjXaPDmeRGRwusQyLPS": {
		Address: "BabJ4KTDUDqaBRWLFza3Ek3zEcjXaPDmeRGRwusQyLPS",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Token Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"2bhkQ6uVn32ddiG4Fe3DVbLsrExdb3ubaY6i1G4szEmq": {
		Address: "2bhkQ6uVn32ddiG4Fe3DVbLsrExdb3ubaY6i1G4szEmq",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Actions Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"CommjpK7mVSaP9HwypJWAHmbxLgVGEaCSXMSt7AFs6Ps": {
		Address: "CommjpK7mVSaP9HwypJWAHmbxLgVGEaCSXMSt7AFs6Ps",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Admin Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"73a2yN8Qd4dN8g9RJBeHCUywMdXaKoh7LRuF7Gj9dL8Z": {
		Address: "73a2yN8Qd4dN8g9RJBeHCUywMdXaKoh7LRuF7Gj9dL8Z",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Treasury Fee Vault Authority"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	// Programs
	"6LtLpnUFNByNXLyCoK9wA2MykKAmQNZKBdY8s47dehDc": {
		Address: "6LtLpnUFNByNXLyCoK9wA2MykKAmQNZKBdY8s47dehDc",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"FarmsPZpWu9i7Kky8tPN37rs2TpmMrAZrC7S7vJa91Hr": {
		Address: "FarmsPZpWu9i7Kky8tPN37rs2TpmMrAZrC7S7vJa91Hr",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Farm"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
	"KLend2g3cP87fffoy8q1mQqGKjrxjC8boSyAYavgmjD": {
		Address: "KLend2g3cP87fffoy8q1mQqGKjrxjC8boSyAYavgmjD",
		Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO, "solscan: Kamino Lending Program"},
		Grade:   labels.LabelGradeStrict,
		Remark:  "from solscan",
	},
}
