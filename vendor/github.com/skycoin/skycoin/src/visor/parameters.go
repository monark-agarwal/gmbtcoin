package visor

/*
* CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
* AVOID EDITING THIS MANUALLY
 */

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 5000000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 0
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
	// MaxDropletPrecision represents the decimal precision of droplets
	MaxDropletPrecision uint64 = 3
	//DefaultMaxBlockSize is max block size
	DefaultMaxBlockSize int = 32768 // in bytes
)

var distributionAddresses = [DistributionAddressesTotal]string{
"XbEBHKtuxv1XHv5DPBSxhQQ2BxJRzeypM5",
"2MAmmP52yMo7c1mxdRjBwHXWrHWVxouaVbA",
"5mgFrAbV4JRA7eSYiRFKeadB4sFEEPd5v4",
"9ataqJLSthHyb7P56d6fefbQB4P9L9hQ2T",
"h2NG6jua3MUjJhr4rQ7tL6Xq9ZEHK7jtLU",
"2FEDwZKQ4x3X2bK79TaP3A5FXuPUuXtciSJ",
"6bk6gQQrVp4mLq6HvC3XgtL6zgZdGXAnMV",
"2UQj6E1epx6eyY4boJXiQcnwNDBgFwFrgT1",
"JxZxtfwSYjo2rPgJuZytGPhBkHja6VJbNe",
"2gU5oF4bJSmqYVsNzjuyB9DUKPYKFMvHAWq",
"Uxaffq9iFdEUUgPSZCyEiqzAVxVE2KQYaN",
"qN5t36Y7BRqQfQgkEmAX2nNnMfLTFsxPnv",
"N1rK3xeKxokNVSP97d57hG3wSWo6KDj19w",
"H4goZv3VKKVUoz97h6oL2ZSCvp5G53mBCK",
"b8gunN5ftZjxeTY5yZk74vy3VWUFxLmGst",
"2ZPUoQQUsynTPaU9Hj1LeNSf1ed7H9f8aqG",
"mpDGasroDRxYVR88bqShHE4VmsVH24V629",
"2etvp1LN28SHTTzLVjryCU62SjExv7N6YF8",
"2Lbnm5rTybB4AVXK2CNXxHnCr1qgHEvSq9U",
"tQnLy3YgRaLT1PqrEA5P8siaumXbxX33gu",
"2YHNWN1qpKX679saaozi1FQm7aGMiNwNLTR",
"2bZQVVGk6Ryky6iWerqMTk1KT1uDUQ2Xdb1",
"2GWrULfjuBDiBtB6SCNW9JJqXKY6qneZAH9",
"2YvMQrsxASwbagD1ynGfirR6iPcLaMtFDki",
"2ETbG5dCG7ogAQa4QR32SDCrSpqfcLij8Nw",
"TAaVJT3UxXXVaGuNdStAPtD1CvwUCdgzWp",
"2dP5JKokZw5CZ68gcFyraKgN8zHsKbVFDTX",
"Ymr5eeSEQC8j8zk8Tvp8dZ59S4W2d6uYwr",
"22T5AZwNY54jNwXYt8Rf5wB489r3pVkAQfb",
"tcGN7x2o7fzPWa1MvfLzheXBb6kLc1Y5jS",
"5RaZD7Qw6zq1RKvsPywnyjMUNjhkvy8uLQ",
"28KyUbdxxy71v1mJztDn2Nu8v8TKZnw5D5j",
"dH7LdVLhQqaqnydv6rePRybQecPbET84Qc",
"kY25jVwgaifvtY6U9ogD7VAph4RiHCAn7N",
"sBEQ3BfcwiiBo3MSz6Loj8nsbbAuSR47fk",
"bCBBFFqj5RDirLrb3pEMqgu8txMERACuuE",
"b9ZfgqJezcskTCPMWL92Q3ZwhFKoGbav1Z",
"sHYQDVDh1hS2wK28SqGDLoFzr9DPm9jXoV",
"3QPFV847FapWqAUaMnLPkF96YZysn3kjQs",
"2QJkC8RErDmVchJzxTQVXe6bqMBNE41DZa",
"YVmirWqH2G6ufNy86ZkEQMtpfdJg48BRKM",
"2iBerkhcVspJLDcE2ZMkhzDBzRgmz6jzs4i",
"24Bq8FFZMfwe88Exvt1bMseAoatSNxu93WU",
"P41JjNa6QT1sawRNbWJ8AB9A3a3iy5EEFL",
"Sp6HvvQk5qUWyyo1Y2Ts7Msx9Di7xiMBpV",
"cPhetHrYWihhkj2262Z9pmQcPPsRYemAzi",
"2kU42JnpVXsjL1DHcE5oMwiPY1vF43E9Zzq",
"2KaACFcjgKLV4PBf1tqrVXS2TVaNNe57ZPY",
"R5319mrVXw4fAT9tQ1zFwwcMDn1Z7DqCfj",
"2Y3vwFdyF4K5eMs11KzsBPHppM3rnQ2Tq2p",
"2eJMEVFUy3gFDVRTDYD65DjvfnbBmdhZZRR",
"fTjZBrnSyhQGXUf2cYA992sthAQqFm6xdi",
"2Q1jtVZ3VVZdR1CQzN2pbxjQ3cWHAtFiajX",
"28HxrWTsim4dvEepeA2qeR9hKz3YSNUwNUY",
"PdPegPHoi1WnxQMMCQS2fu7wwfwQFsdEwp",
"2U2AQD4rqA6avpM3HQwWEvYy4NRbtEsADBu",
"22YBeQXkfod1iNLrtoqeoDnEJv2otf1Z4Hc",
"tkJKCWxGQ2EWXXHxbaCizGyMXLuyoiKpMY",
"PrHfJ8vW8HTXBbXrP4ymfQdqQJpdTr6nkH",
"vuG63nm5U7pegavENGxdWW9hXoQm9kjJJp",
"xBa9wUEQZdJjb75jUjPige5xysZLWHigiz",
"2W1Dawb39wkYvwm3jG6M1Jms9DxwJuyyDu3",
"2CLHXdQo23CV5vz1K4dPxV2dz6a2ojZsQX8",
"2CK2ZW8KUzTugo6hV5pbrgkFLd6K74X8wyN",
"e252FCKg6Cug6XzRZ6AsuZdPyKfp7wEoLz",
"RDgB9YUxHkLYfnYuTfAZwh75eVG84cXe9v",
"x5zaPqaCpQWtsDtWFdep5QQtbW6NuqEv48",
"PJsnEL16rkrjHKjRRYydvQ7GBceYU2HqHs",
"akno9R4AvCVwuysTfHRCb8SCbNrZiYuHXk",
"grmBTpvkLEm3mVMJXnirp1sNNXbixMKqXN",
"ZmZNmfufUNT4VVqd5gsqGpMzSH9mTJzY8i",
"HSrbRL46u94sCVv8m7c7iUNfVYbf4GBsAi",
"ijj9m3xVNZLVfzBTegiy3KCKJfR5k4m751",
"47ZvSFWytpAjngncWafZqgxiQxb4wETzoc",
"oAqynx5Cx5osapLuFXmWg5B94YkKMWCwk9",
"2PnnmD1aPNevNGh9Rq2TSNCoMS8JDsb99QF",
"2XnTTa6ZG9NjM9BcH314y29e2J3RTu8kR5M",
"zGSTxvM2XpoThVrhiQYMKPsmwwByM8PH2B",
"2TWfwXcdzCkJRdMKV5YQTaKk3rUy2R7P49f",
"26kbTMKZFpyuzXdWScNaTufFdqdr5XUHqpr",
"gqyRtWUa6nGCsnF9cXmPuJqx3Mj86KxV8",
"26EiansLCXxLpMnVZpNKrpkrFhsD2LYA6f5",
"29rjCxFFQgN7UakstEHVMMRMDpAd5fNfptX",
"2fsS6MLHCqqY6huAUcuCaQomLMSs2Yvaz17",
"2S5tfUgmtyUAbQnrjycD5XFTcZWnKCnGj5D",
"YcSvwBiMYByXYxduUn3HEB2FHUTF6h61z8",
"2XjopbmArYmwcFQNZp19M87LhmWVnTbE6cw",
"zFnmRASJWybBArHMuWMXGB2m3vgnmCNdTU",
"2KL2xMfYYJCq6Mk5zjJht8S2ySEMBAwnaXB",
"3grZnxFMChhdNqnQPCaPA7g7UiV3jJdb7Z",
"edChe5Bmr9WDaKKBkLQPpNFEmG7455QM3B",
"2Lj7StwHefMV5VRZoHAxZ2VP3c9PrVRsDw2",
"2cfoLQkDpepZ8ef3JWNKU7hv27tvHWpDsfQ",
"3Ei2qYt9kbC2iRZMteZ5f8P1q1PUr1tXCz",
"2kCeEpPEhyZTaU8z5VSB5FWxtnvA95q7hGA",
"2dQWbYH1gL2qWXBwf1TVvB46dFyHLm1ao4Y",
"AC3eF3fzBUf2RDqjngWSBhsEQiwSZmZ8TV",
"ZbrBVTbsURa5v11Ed14eMJmucJt73grS9",
"HBPqooyJDhZSsgxc4eyxPceSW1k8WmeMBP",
"2UnUj8AEXcM9HF93Xssww5Kt9ozM2eATHVg",}
