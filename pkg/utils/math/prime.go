package math

import (
	"math/big"
)

const (
	defaultTestTime = 10
)

var (
	fastPrimes []*big.Int
)

// 递增地初始化一些不同空间占用的已知素数，用作快速加解密需求
func init() {
	fastPrimes = make([]*big.Int, 0, 33)
	prime, _ := new(big.Int).SetString("AwGeptL1TMEBFSqZfp4BXWGY80D", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("bp4BIde7QUfxFaGMulAG6mXcLgIclN9M8j", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("3yLcGnEyq06hbqngIyXNqhUeKJFzKcsuGSD1CnGan", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("16VojPatgWtwXcxnxocPse5M8CZAWEv6Z3VKi4xumNW7FQdf", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("lwegNercKn81SX8boBDos0ksm0cQpaHm7Nu6qJPmJMk2Kgz25o4ggX", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("6IznUBFJvXlEgv1jFCGH4OsE4zPmhPVcHvcyrzVXPOMjCuu1gZxZgbq1CAWmd", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("263M6rAkafjX0vtPBzMJ1DjDEa7Nou7PPzedP6hGWawQ74IGy36ByU1uUBN8cYnahSrr", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("EBEGXHDtAaZMcVEs1ZYx0WVnBxM6ulVZPPQqmkfYjPVSZBZsdTRB5cKmRhR2zJlXMUSMkub4vx", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("cG3zX2IxSBqUNIojtY7wTHmb8NIie4tRODQe47WX4sDkE60lSa5J5bCsZY7rwE7JFiSb8POjIZecbiwzb", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("1eCKtmevq8JCaHsG4M7LaoItPW98tY9sFXMP51SHU744zuZkVBjR83UBUsauNGxZVLWlUY0RQJeXd3xzrCVshJB41rJmcHV", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("7t6q9Yg7hpdczJqwuZLjbtcPtt9zRiCGjx7jbfMRDTIVYKOu0GCEarjarHcoC6HQh4tzMu07S2qfvvZQyJ1EFLF2V08DX2E9zNGf7C7aIUOn", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("J8Tb5MGpxCyIY2pFBqUb7x9SEYG5ID3ehjhGt4VgN0hSL1zpytWcCGuyyJrnHRFyU5z5sDzN4Cxj8mQxECdu25FaA9kzpSZCF9q8F6rkeAeV5I3RW50NVEEWN", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("4oPPXhCZsHLTW96ahHGDfXBZxOZYixGBnignbTzQ6XF06L1fAHysU0YsBiUppygzjGKQndYGRRlKRw9vYmWTUqml2eeFw79iiICumWME6gHKwov6FLS4iHnuIUQCte02Pfupr2V", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("qAXnr1tPpxzrBRhsOWeEwlcK19OJlNhYkIBX9FjYL74dN1slhyYr3I5tOlvQB4eNukoew8hGpSqXGcCcYE5qyU0HC5XcV1D7aYAJt2Fd4SkEB4G7D3JhvNKIQSHkF1kHG2pLF69hfTGdwePpffjf", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("2AJULLRuiEdzBbA9MTKVqSYqqkyrBr8wNFNBj6y2lOZn2TO8yhr7Kqe0ArLgUZKqIYKXXIiRCeSfYUw06pz3Gne83uBki1cCqm1ywRk7QwnHk0aUA0SL2uC2RdX0dNznz3tbnBS5B63euATqYsmskP5ExjU75iq5G9", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("ObxoLSR38Vjstir2mhFGqeiwcFPVoQsKkkXT8x13t2WW4ufNcbpd4lMmc4FhjrKU5SSBXeQmAvSE8zrdFO7BCHQFLX3DX3FiNQ81B5Un6H4ByVDjBXIoWDCNmb2grMgIQNrNHfcss1Jz9DB1jwj6xoP1TBsjLcq3D6oVsTrj", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("UYShNRSeJrPWRFR7lSLuwaAjiGhTGHKsUp5bHtQbPHMOWlYw54dyiDJ9Yhun3cZU4jlmj2Fx01vKUVDLkBHJPbPh5oA1hAPzBH6XFQHJFR7KIRIZipsf9921RoTQ3TWUrDAuZOTFgAJhswqa0WQzFGDsbcUN9dW21aielkO7dgKX", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("fFtFthL5piR6QfHLJEfW8YEiTdjF8er4cFcwjSXinsnB1LHj41Tco18PrpoByJLldIkUr54P0KI3IGuaI8EACpBo76BvpGG6zDAS41iplpHjjS2zBGJEJYWD2FoTCQkuZoU0lQHV9VeYU005SMrCUlVJdwEMD8YczTgyTSMkeeQbXwL", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("1wHo9vkBI3D4BSUOgirtGVhQNOFDaED4YV3kQKReZf0mQe6pNx9nh38lJVk5S5n9PQEZtaj2uV4NILJGdftFeiaD5mxT6M1LJ1wkCWVohY7jroAez8Qqh6lg36ruCtHyks3WAGRzbxVAjPETUEYURbXhfAbALBdpX05HhXWzdjx7hPYslWcclODHMORzH", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("9elj8jsXBRaN3o5Hnq5tBJHTBCUsEdTo6SOx4VNZw6DHLJDd3pNEaTG2CRMdCuTfHu7hsRIGKFB9jraCRDryilQb9DLcTCiEuZeS0XAqkMpaetHhNfrGvqjDcNrEyOHeLNJceOZVjS2WLDu7wi7MFmKjFc2KMqFymfBSmgCd2qDg4mRyIneF1whQljW4wUMg07TaBWyO4X", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("TN6VoaYM0li2RYq8Ki7pTrRl7JnOjTwlBfs6J0B3rSLM3SWfSzgjNq0tWUYiTLWaa43k9P7uNcUnXe8ucqqtteY0mbiwELYcdJnoKOa0Jth7hdWxgkOy9BrDe5OGnOFKRRlC4M4pyc4KvBSak9yA6R26r0qS5k1zAuHrt1ho2sOoTWekLPxWBQ5mMuqb3R6lOcbzJNG7UsUHUa2x2jbFMhN", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("5rc3YaA7sKTgqKz98b610KSf3t14pAkAlveTN3B3qrwYGE138jDOoQ0hKrHuKnFDwPPaNlBfzhRLtRaxdARRJTc6uiPoNptnNAWvIFrJAo7VK8oKNRfDrvG1crEhtQJsAUmzH6oNVT3lhLIdIv7HQaVh3UQ0ItI9x4otAsFO8RNIOW9aY9Ld8e34WpGQvfOozYzoRn7aqQiCm1zFIY6x2nb4gXWm6EATspLYt", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("wRVn8mSnssu8Q29Kham6yH7XnHNEb4w77K4ZwRH3TVm0ENQSlI6LXWe2IVAafon4pVwXnpVBtoSb0XgiRxKKLcppuSfZLzlMCoq42qgBlq2cU3JtFXqfS0mw4kVsd3dFZAgAgTPxMUvOPsgh7jSGjyXoqCmnYEfiB7VvjQNPUyfDfz5EHDr3SDYhAneeclaVlsbUranPHIq2NHqIczzsqQcT7xZmgTnEUxTx4MR22S1dIL2OVR", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("3cEKdxbNL0hXVOxRKenbD64qqvRoVit9OioS4sIZFIibmkyXNZYd7zkAgpzIgqLcz8d8WHGd8I7X0pc02fUoMJ0dXVod3aqx2Ci7o3LOmvKssKchi9kuwvVvNSx7WvRvHLKVP8jRA02W316FzviDLOeubeon51jYKjNRnaBLXjBkzpFyFYAeyYeTLz3rQ73HtMNXXJokGitT98OhnOXNenpNwLurtclWdeXbX5iCPEAc5jO0mT7bBZjPHjV9OSfx", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	// 若是没有开素数fast获取开关，建议秘密字符串len < 200
	prime, _ = new(big.Int).SetString("jmDxdlKz7LDiuoOOQtCanRKR0lg2FIcC4Mv4U07BoHxm5SqJB8Yw3ni81kYFOeglqS7Jmjcgw5YkOXT1BdjiFebJDU31QZbylnrUrHeYYdZUXwzhgkJll7zjfE9clKdnXVLNKxdWy93SnxaKiny7U5YFzsCwSng3v84jUQrqWQKYvt00duKCaWk7xslaEOOYsSsklW29haWMoYx5jG1jcy8i5XoE7OWjNZO5VxfpsEOtKD2O17sCGyG3sUwVIY4GUiBZS4sUMZYyX", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("ECE160ZTXqYCEQk8DX30xllVEsgM4LctUk1zu9HpBIWMEHzGBQ1Ue8Q95ia2T5JKl5nxXedq2xvq3RHfHnh3LtcqIuFymL5n4lCS4XlHOE7IPbEeHb5yd0ZS4Qg9J4XBzcNPGFLsO86zt43LZ5ZmSWFxl0G0Lp8BfFvJU42LNplaf5Et597T6FkHI5wmJFLZ8WpjIKrtrlR60aRL0hO6KjrGAXv7VnoZnT80gtMl58wxIAkyc08IcTzmRIdljNtR4aIbZPPdE6e55YpNQmEJqplHj0Rqas5ivgADMc3nPZcapAq9CafJ8qWJUNLOJx9EGHuZK7BLOSYBYXc5", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("1ndvinsJGzMUGS3Mi1mY1rQ6QuCumbvLANwjsq71lHf8iLHMVnrle4fnf2xjP41zLSM7TU2QzCRADjV2GdoywGYOpLOW686M61trJjFzWfCrlTwjjBbVRCI14GXmK1GzZZwehspd44Q3f5oghSvt3D6ZX9tBqHHELyxUeC4EC2iMvLdye80fy47NmrVoQqLHrDCQ193nDfphuyt2jJAyazPanJxeIOmiPUkqF1l3QPKRajl8BmUTZVkoE3gYcFHaqncdO53ZnGvXqeH88H0RHKpxOw5BlR1sXZtEKQlnHtrLINa8bRYfShcVGkBnbFCghBNinyfpuzMmBbB7OFTkioWdLwnxizWoP5v7pmVz69WuwhUlsxZtZLG9k7Ta3lF3dr6w1n9DRALxdySjLJGD", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("2SLvomxsjCttZG5N1EBUFVxqrlkLKrYnZjPMDK7lVD6BHjopeOF1im2DXak2EddiGkuKzSg25cUR0z94DoJznFPvHN9tHCeUXe9HgoKyplsyPZm3RpFScQaW3yHSsJcjOZw0TH8fIe0YOBqdLyE98lUBBztq0orLFwOMhSbPKfaimluTGzJvtMr98qPfye2gfNadrsfm2xlgS6rHCTFC6U56zeDHumj7KBXr6Mhjv8oSXnw2AXKpLUGiL1EU82EEQYfHi2I44jMPzbaEh9zAHpVm2Ccb3pTroyq8R1y404Aw3kNfcJwzR5F76sh39DqeP9dgNbqT4y7v1BshbMuX5FI4zv6daJjxIBhqkbbatzu5NetqYKggvbhbuzuIqcf828XYi5qHI3VR8lBUevSCJqstkMaJpZ2stfaDE2g3xhjoGaF4X85SpHyzX1Od1hN8UCAqIqBtnvO1B9ijWR8Cl03", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("630uOXSzciIdfVyREWC8Deq9PiX9YHmfrtfBvQscs86wp4OPg5oQiETRQJoPrPO8H0CFG7YwYRKSVCXC4enhiHklcT5bgdVbz8377g8ypG6dtAVoyLFG3v2jy1WnUP547WbZ9RlcT5knkDHOSgn3au3i3QxOFpLgal4n8da5R8zqu7RbXph3DguuivGCy9X1iPWZg8hVXHqjLRwJCUQCyVSay9W0PQVGKcLC91xPizSATUsrJ4iBX4Hfj3qOdcjfWbslceXBqCj4ZotFV8nKJ3rn0lywaS0AQfPDmibmmunvvdmsPbsD1xlpsob5xnVi6S0SkqvWHqa0fpJCrb0hj89xGMYg79Qf1Ejv6xp7Yz9GhHUdivU8ZkzS3T5Y8Ei5mq6ZuXV9M8COyThq9pk9d9KknBT5h79sqDK1JR9aK5inMUWXbLUanoMi8gaOUwzNrHZzUXBmPRNgyiPg2ACNz2OiB5w0amQu60qKu8i18tUkfwe1nWfnsD2VBPc77M0kkfbvikwHYPbOHQbVByBkeCMJ6D", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("cGFSBIMGpLJhGsj0MdckwvfeuiqM7z6jc2V7QZjL5YSEdlkp0esEqzwdQBSOwjamJZrcnsyNM2H67N2YBOzOs39KLKhv8BOC6QfQotcYFkzklcP3n6JHxNEfn0chLN5SJ7ELncjBJDxbs6PNZuVNQKtQ7NfBS4KeQ6lY7pPBgQKqOp23dH6L0UXPr5CYMeFF7uUbC63BHQYhGMNXfp99O3rddKyYAnr8G0xNsSrAhFWYPhY0MwqzTcD1ytASKjAavBna0FwpqFgHY4ZhvZU2SN04PtmfupH4KjySY885CYCEBM4RYlIks32rYxzuts4BzmeYa62Aoq1rDKTJgpRhgO87TqcD3jmoVtGFRGTcnGy3UYT4sD1R1EFWLCvzQnMV1IT7Wjb9id8VU2UjxN8WAVhLhf6BGgAa44jevHHMQ3PCtifj5rhDvM2NMgxF8eOTFyuiQFHMhrwbkqn7EFBFbrH7y2YqUCzm2wZlcRoebTUOLjYl43w6UUk8FefOfHVrb11o2H9ymDG3PLaFw7CG5faEkAj0WGU0KuseuJaCTOxJey8UczzweOPgdp606OFpePTtsRA4ftprv1FZIfR0NSqiEwXPV", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("qCfISW4SzM3MKuzgBRg4l6rqcjeuzdhyxSnHkuda0JWTYdCVhZlCmFeyEfLnLLVlEUFfpGEPVOAX9H1tNlRAy9ApsgMFvy2BrTuwMKLcXQgaPiSZRcWCE25GyjnkUX8RffUI0Xv3P9KCXcfO5QybJWYl0175B1SAYLSZzFynpQ6SZOcq7gb5u0248suaGiDAPNtnY35wKNMacJwk3FreWhfTqFTbUjlf4pqT1tvSN3X04jtszHhtyTD5TIVyTtIbG9LBrx5QPAuhgzmvWxRsujcrtlBUd7aENeXxdpmrCyCuZDD4RdfQBTaU6lIwd804egVzvMJTik7tJLCpL8Sitv8zocKL3fNiP45s3zfkTdnbV3cNWS7dHFR2A0nanp5dUpDP5agpqf0GRW9r4rHoHvtJBdBmtDSBGfYOduNXXFytMdeMz4JKYa2rmfeiUcswidmRXekCI3cnu2VaJCfOC1NWv4ND9OL2YXU0dzEmnoRW9mtjcxwb5WJrt5YcwSmJzI0cMSZ6lU6nCNVRylTqHWXM8PGiEJLrlVKEyvE9MRx51H7ARa0tVGcsNOWi2LOYJerJcv2WXXDb3jzlKPUlItmYz5G4fYxdYXYDglYUkgeB2PAeZt31PWr1SroKKXZFYhnSVzzvk0PEn9AMNu0VKB85Tbdq6Wdx", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("TPPtg4XDmRfUza5cLAvfE2UQJenIvUBietpy4QNftZGaoQY2S1UAUKcO02EwpTmRK06ES82cTuok7046TKORxyztNBaCA0558u5ErFKZ9PKv1EGaeYcAlISTfHvevHFkURIzvfRQZ0SzkO7YJ4xjzxMQedR0EyPk6XPduHkUtQaP0n0n0mPzeUXNYoZWzQiTBdYffGpEbEcOeBhoEJqPc82HHlKsWIqXpN1NiuyH4QwhaGIzkWfUls0FeO3u8mLQZ95pEsz3wehziixypvXfkUkXasGM0B6OeHcIWjnc5sQFLLBC72UWIZK1ZokDKsPX5kMlF4PsczJbxTwTA5PScU6fL3nw3bZ8qUmMRGQTo8G1lYZhSu9bWyPhlv6Xs0nEJzSj5miOhvVCKpLXws0yyCkuJJIzYc8nHXZUzxosdMur619pIQCkO9c6O5dDUw13YqtVRTkprQBXjufd0uTk7UjE6avSijlaJI5C6R2SNV58HpVmvQa680QAPByzkxVfKN1Rp0zHKg30cZ5aadq9w0Scmttv6Q2DRV1IB8HXK3kGcPhI6GRJi23Q2RISetEdGcuI4YKGXUeYQZNS1YQoiQa54S2N2Xszt1ypQQ37rbkHoDBO5nJlXH2BoKl0JDW8EXhYZ59RNUzgNjFnzJ46pognVWo7JO97iClK9XRvAyejzERe9Krw0YNmm2hg6Ov5YflBaXYRClvYcUi2quiNtjEsskvGCVrMMSL", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
	prime, _ = new(big.Int).SetString("1T84KLQ6epeLGpYInFR8rE4CFweHzVdrSXOweLuW1BiXJiTUIgKzisvLXUwriXutM550HlRywsXapQ1VfLTICppq4o0jsLNzBPkTftm8Q5rZeSR92H9S78RPUOO8P7xa38x0QxhsIfVwkTN8GkRAGAOkX7XM9QC35tltZ2xIui6avWFxNDB8r4LgpdoeOms7iyIkYvwG9hsPlKu8yX9U7bgJKBJh6Bhzn7QkeCzm6tUL5KnwHELbwXTl9oH7TDQ6AwjKPnnzlabWmU6Vj2hFI0wTohsWYFRMRzPDUiCC4ru4fx7w68PQMxkKR1EhLQYiSXOoLalqe544jym3hukfJoiVTM1klWb5qEFKwKJhN6gvS5k6kwuPpNxVplCYKv8sIo8K8KTcKSI3aZVIACdUhH2VFDymhZOcDoAUR4lWRlj3LEjlS0eKZyAXzSV8CTWnZV6L9H4ZnXEee2Tt9d91jKO1OXcM0m5hPdL1gZOzU0K9BkMkYCCwKL77hC0I66ItGbDW9omZrGK860YTReb9Adv2KECmO31r3NN9Xn9ggyv51TYuQYLGTXwFbMYG6hkoARzqPda8zjHL9dnZNImXki3JZkdKfkzMGuv1xIYat4FAyaDSKCZrVj1EsZOYbD79UTt8dUrWxvUZo0qzHGnq88PxAhrXILYFrvSl6HwWhLqzsHGp9GIohp2E6b764gTrhK97L33UkbRErwWDoakoO3bmtPsNzOdnIOVBhm39iTEIeQ5wms6F4PZtdyhlGfVPY0gIWqgy45Z45e4Gfv3MtWA9kUU3pxHYU3zOpqJ", big.MaxBase)
	fastPrimes = append(fastPrimes, prime)
}

var incrementalSize = big.NewInt(1)

// NextPrime 返回 > num 的第一个素数
func NextPrime(num *big.Int) *big.Int {
	if num.Cmp(big.NewInt(3)) < 0 {
		return big.NewInt(2)
	}

	result := new(big.Int).Set(num)
	result.Add(result, incrementalSize)

	for !result.ProbablyPrime(defaultTestTime) {
		result.Add(result, incrementalSize)
	}

	return result
}

// FastPrime 快速返回 >= num 的一个素数
// 这个函数将会从预先配置的素数列表里找，若找不到最终会调用 NextPrime
func FastPrime(num *big.Int) *big.Int {
	if num.Cmp(big.NewInt(3)) < 0 {
		return big.NewInt(2)
	}

	for _, prime := range fastPrimes {
		if prime.Cmp(num) > 0 {
			return new(big.Int).Set(prime)
		}
	}

	return NextPrime(num)
}
