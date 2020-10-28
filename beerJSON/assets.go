//go:generate statics -i=schema -o=assets.go -pkg=beerJSON -group=Assets -ignore=\.gitignore

package beerJSON

import (
	"os"

	"github.com/go-playground/statics/static"
)

// newStaticAssets initializes a new *static.Files instance for use
func newStaticAssets(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{
		Path:    "/schema",
		Name:    "schema",
		Size:    4096,
		Mode:    os.FileMode(2147484157),
		ModTime: 1603903802,
		IsDir:   true,
		Compressed: `
`,
		Files: []*static.DirFile{{
			Path:    "/schema/beer.json",
			Name:    "beer.json",
			Size:    4128,
			Mode:    os.FileMode(436),
			ModTime: 1597743398,
			IsDir:   false,
			Compressed: `
H4sIAAAAAAAA/8RXwY7bRgy971cQTo4bq7cCvjVFivZQYJEu2kNRLKgZyuJ2NKNwRnaNYv894MiyZa+S
WBsDe7M55CPfG3Jo/38DsHgbTU0NLlawqFNqV0XxGIN/11uXQdaFFazSux9+LHrbm8VtjmM7xMRVUQhu
l2tOdVd2kcQEn8inpQlNURKJQh4/NBgTSXEwLvVTj5p2LSlsKB/JpN6G1nLi4NHdSWhJElNcrKBCFyk7
tGOzkgJYDMkOlmnwbLcUjXCrOfT4YwgJyFFDPkGoAJ2DAQ5sMJ0exOUx/lsFZqeJIrN9QxL5pM6pmj78
1zo2nNwOyJtg6VjRHgC2nGr24DjmqoVMEDsqM8O+FaoUryGMnRCWjh46zynmO3hTWKrYZzKx+LMHvlfV
DhhPR7hF4ob9+mEv59cJ3NcEvT/0/lB1zu2gdyspQhp5VIAeBlkzMwgZKkIVBEqMHCF49adbWAtuOO1u
IQi0vwImQL+DVoKhGCEmar+gQp9ugvp9Pvgy84pEm0Dli+fEhy5DEdyd5j1vtP6KMiWtmP1ayLK2Xaox
gU6RcNklVSdkgfZUVSH9mqfnNAcnas5rGlEeVT7B+5fj6Sl5gKdJIRqOhpxDT6GLD0cC11LFPnbepAjb
mk0NNoAPY12uJovymNDj9zG9yxSpQ/uwQWF6NuovUsFSQnY6FEqs2Tf2/iVRvp3nTx1BHVo4pJ3HvQ7t
1PRnsN1vvgrSYC7rm9xN51In16e9ZUuQg0eMh2TzyO6jJgj/3J/MItxKqPhqb8AWEwkMmNBFssAeSqGt
PlOzeGasCZZ/qf09xgs6OabdVaid3qWpUdAkEo6JTRyPLPQZey3ywD92dq2B644tOfYUocZN9vZAUR8q
jjXZedrkNBPa/KH2C989jPX3aPMTmOAcmbzfQpWXVNSr37BVwsPy4mM3ZllMaJrgQdMf3Oz8KdD4qecO
Y303QF4mxLBPMsCr6TGu4sWijEG+shv120yRhAy31xgl3XzIXhVBaNhzg+5MOR2mUbSaRmv5dqQOoLcQ
Uk0CQp86FrLQomBDiSSCJ1UcZadLVsgIYSJAKDGZWmHnb9lehwltP+aDy8TUUtvm+35j3OXGojHJ887a
v7pwSDeP6jHsOdsPw9llhMvA7tUmS5O/eKI0eIL/+8Bu5gS1aP5FXQWvJsShghercUCYkORuOLtEl5sz
fRbD+C5W8Pfzf5Z7yz83Q3AOfB50/M98o95PN58DAAD//1PBDWcgEAAA
`,
			Files: []*static.DirFile{},
		},
			{
				Path:    "/schema/fermentation.json",
				Name:    "fermentation.json",
				Size:    919,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/5RTTWvjMBC9+1cMyh4Ta/e04B+wsLdAeyulKPbYVog+OhpTQvF/L5LjRE6bftzEe/Pe
07PGrwWA+BXqHo0SFYie2VdS7oOzmwktHXWyIdXy5vdfOWErsU463cyaUElJ6qXsNPfDbghItbOMlsva
GblDpGh5ORgVGCnlyBbJoGXF2tkyIpM7Hz1Ge7fbY80TpppGxzF12JLzSKwxiApadQiYBhpstU0jEY/t
AMS/LGBLrsZmILyf7KeRj+MS3mCoSfuojfRNL0jRGIB7BD8z0DoCj9Q6Mtp2sOh6CfmqVxryOTXfG0BY
ZXCBZG0Ck7adOFPj+qJbNvux3Dq+usk3hfk3eAqM/qaLIlJHsc4po+1/RhMlfxaEPqG5UdxRwjZavQtN
m7aS2cLI/G3vGH1akcxuLK5P51qC8HnQhPGHeLh6ms+rn8jHYvYdi7F4CwAA//+cvAIPlwMAAA==
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/equipment.json",
				Name:    "equipment.json",
				Size:    3412,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/8RXX2/bNhB/96c4qH1oB8dyhq0b8tYUG1osxQIs6MtQGCfxJLHhH+VIOjWGfveBtC1L
imM767Y+BJDvz493v7vjMX9NALLnrmxIY3YBWeN9e5Hnn5w1Z2vpzHKdC8bKn81/yteyZ9k0+Umx9XEX
ec54P6ulb0IRHHFpjSfjZ6XVeUHEEXL3odF54nROTndBtjraxp9raL9qKWLb4hOVfi1DIaSX1qC6ZtsS
e0kuu4AKlaNkIKiSJplEeUwNIPtli36Jjjrx/iOSXJArWbYRJqpvGoJOtCQo0BFEX6gsQ8F0DwJXsEsC
rtkupSAHwci7QNB20YK3UAWlVhvEgsA3BEylbGm2C6Ht57cNGCAzqGkg6aXhPEtTZ53qy3QytnmyX2VZ
H/Gb9nVkQrT/sycDyN5e3QzMALL36Bq4CWYsv8LgifdpLiPRv5H3isaqX4kj8cRjxetamho+kHOkxrpr
LG+xr++pP+4lQ+NnqYNeLK0KD8vwnKmKtGhCF5iwULQIRnqXmvpZ3uvN/ENCuIlE7k6ajE7MOLYUkxgQ
um6BcYkej3LSz2iDvRuJd570zbA3Th2LDiP6xxZft7yhkpxDXoE0MTSM9mlSpBFyKUVAlYYmMr+bmd0x
qNTv1SDlvTQPCR3O+L7qDUAeGa+kU9Y9lP4L9R3FkzBrRmkWWDjLidcFox831j7u45WEbYtMxsO60JBQ
ChJQrCDhTuMlJUtUsEQVyAEywXx2/v2PcOdzVcCL89n8B7jKb+uXqUAIGl3jg5nCfDb/GWpUyWw+e/Wq
Z3b57vXlbDRNT+Tnj5ZKWcnyVJ4KK1XiZtESLxob+ESSNtxEfxJgqwpEiLcWnENEmcImWgEFVZYJ0AjA
Kl5B6IGt1eBJt8ToA9NXZn1qtiJ1RZeuliac3BWbhH2DHhThMq6chuA23ZtTIBeZx7iApG4tezQ+FdVY
cya1JnZxXMtGKkXs1jiltSqh3Fv2gA7kHuz/iZt7knXjTyRjbQy2SnG2kkqKP7pb5wAf9w2Z5LUZCZAO
jPXQMjWEnsRX5vsenTuerdvMySKeeWLSWx+IPgdz/9xy3HwCpIE3qPIX9XdvXn5DTra3wltCf5wbY/2e
6/vAuwZ6S3b4vf36eGT1ptVweKH+w2V6fXx9PtiZ4GiwN488juG/eFd2wSykJz0uRweBzLgavhRHBAyV
+8B6zTT8f+GRl0D3tDlY8qe+tsYJD7thEv++TP4OAAD//xXle3dUDQAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/boil.json",
				Name:    "boil.json",
				Size:    1094,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/5RTy27bQAy86yuITY6Ot7cCvrVfUKBBL0VhrCVKYqB9lKQQpIX/vdjd2pYfReqTF8Ph
cDimfjcA5lHaEb0zGzCjatpY+yIxPFV0HXmwHbtenz58tBV7MKvSR92hRzbWsntdD6TjvJsFuY1BMei6
jd7uEDlLnh7eiSLbCkaa1vlVVfUtYZaNuxdstWKu60gpBjd94ZiQlVDMBno3CRZChz2FQsl43grAfI6U
+S12M+Nzla2l22MK3qG0TClL5fKVBpRRKKAjQjpUoI8MCbmP7CkM4KBsBZ/K74L3SjpCiCCKSYCqjDiP
4AQciLrQOe5AKAwTFlqVOjl8L4xCSsvSYWkAE5zHM2QRhShTGMyxtF+d+s5jubs9RL1w8p+NiXGb998K
/boy/sjYZwGPTmZGt5twOwdSKef0YBdXYb/Fafb1Cm4OKkOUrtO5Z8gzvTui/PH/SsIxuzezWpZI0V/y
F66Omjfc5Nv9qpjOHQHsm8vX0aVh/DkTY/6wv98K5y/2ozm075t98ycAAP//sasZH0YEAAA=
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/style.json",
				Name:    "style.json",
				Size:    2994,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/8RWT2/bOgy/51MIbo9u/N7pAbm99rDLgBVdTysKg7Zom50suZScNhv63QdJ+WNnTpE0
GHYoqpDijz9SpMmfMyGSS1s22EKyEEnjXLfIsidr9FWUzg3XmWSo3NU//2VRdpGkwY7kxsYusozhZV6T
a/qit8il0Q61m5emzQpE9pC7QwvWIQc/mXUrhXN/jLBu1aHHNcUTli7KQEpyZDSoWzYdsiO0yUJUoCyG
CxIr0uGKl/uwhEi+euRrsLgVTcMHuURbMnUewqvvGxRb0RJFARaFtxWVYVEY14hAXNQ9SVSkUTCWhqVN
BWjpf1CH6zsdmyVZT24ubv1ZohW9puc+6NYBCWcESdSOqpUgLWlJsgcVMeyOaTdMwSYuIRINLY4kg2it
Y9J1slW9pTu7EhzWhlfn2Oa6bwvkQxCkHdbI0xghvlyhc4cB1hzSoa6F18+oa9ckC/HvWEP6gKYD7yW8
8MP/V9/E43uUwtN+JCvrOyO7S8bK211kg1rNQonexCT6B93BzfZgE8bnnhh9zz3Et04Hb5eOSadrEo+z
AUhsiPsxuWMbYmsbq1miIF0ZbsFfCE0Rbog1I/oRFDs4UOqLT8DDNsRj0xM6eCrPI4ADXRF0hqkmDSqv
GZbk9it9RKBFsD0jFArzXpOz4dM0pvUpwtyBrmM6R2Bv6dh79fdc+75jDfHTmRcUih+tjfDncrne4h1J
pzTK7Lf4yV5vPMixDoELE+M/2+0O6kjnoErTGJUXq3xpVP/bx/lkCrfIJWp3pHtt3EQvvPMFm4yBTdgM
zgPpOgQGXU5m4BSkSsFyuoJOQWlN75oKUZ0LZJbIoFRObcdo7YEqOwWRdM0oCfV0d54Cha/QduoDJTCb
Om9O44FyF7acc8bKHoII9Y52vTgNR4x1hlEK0gI221XcuP7AjJkMdn9WTwS7t6YkqPs2zGu/+YZ5TTIe
vpu26MsG/LlFkP6/cU1UWiOD4oX0doTP/N/b7FcAAAD//+8SZziyCwAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/mash.json",
				Name:    "mash.json",
				Size:    1028,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/4xTTW/bMAy9+1cQao+ptZ0G5B/0MKDAchuGQLFpm0X0UZLCUAz574PkunGybOuNfuR7
5CPlXw2AuZduQu/MFsykmrbWPksMDzPaRh5tz27Qh09f7IzdmU3lUb9wZGstu5/tSDrlQxbkLgbFoG0X
vT0gcpE8B96JItslntoSzar6mrDIxsMzdjpjru9JKQZ3fOKYkJVQzBYGdxSsBT0OFGpJwYsrAPPVyfTE
scM+M+5m2Tl1u03Fe5SOKRWpkt5NJFDVUUAnhLQIwhAZEvIQ2VMYIQd6yQjFTfmsdSIo7Vn7fzZqUVqn
lnEBTHAeL5CVCVGmMJr31Glz5o3sKOwVfUJ2mvkPkSvHj2WP7giVByseJKbIoLHuQdSxQhzqRz2h2axV
7xmHIufRSWZ0hyPucyCVeuo7u7qY3Z2b1DPd9BGiXq3kgwsow+1FMf2V7Zjd6+X4nsKjoi+UzxcJekPX
Qmu7S7MbNst7/KaYLj0CnJrr6H18w/iSibH8aN+v3sI/b3zb/xv4o1manZpT8zsAAP//QPzwTQQEAAA=
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/mash_step.json",
				Name:    "mash_step.json",
				Size:    2309,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/7xVTW/bOBC9+1cMlBw2gCNtdjcf8C2XYA8tUCBBL0VhjKWRxNgimeEwRlDkvxck41hW
XLuB215saT7eaB7fDL+NALJjV7bUYTaBrBWxk6K4d0afJmtuuCkqxlpO/74sku0oG8c8Va1y3KQoGJd5
o6T1M++IS6OFtOSl6YoZEQfI9UOHToiL1XM7dUI2D68JWp4sBWwzu6dSkg2rSokyGhef2FhiUeSyCdS4
cBQDKqqVjiHBHloDyD6ia2+F7F1CTNbtFaK9IleysgEluPvpcAoIlhjCxwKTZXKkBUMsmLL0zEo3UPn4
Jy1BaC1fY+/rIAbZvmv1uQCZxo42LL0mnISS2avreTwaxuzKG/d9pH2XTeBLzwaQKV17p16Op49EnSVG
8UxDV0WlKWVLjjOJocDOj3xLw/IGkFHpmAXi36Ja5IaynvHrVj6wM17LkJFjpjow0hE6z4SzBU29VuKi
KI+KnraKz2bhO4qC2lohyGPaJ+aAWndrmH0F1Vt9vKuS2tUTY7e9xGBe7lqCxDCYGkICgLQoYNE5cjCj
2jCBtMqlMZpRo7TLgRrozGM4+ppNB5jOOYb8FX/PTgLi2X9XN2MQAwialtBjOcWGkPOLm5ecf06gwycQ
nBNcQae0F3IhuyWU9YTCbQyGpfGLClp8JEAIHacGTL1OzjeH5dcRTLr6o5rZPLd3b5UlCvG0CRM55bAB
9+jieuEMzLVZakD3ynzQQTnX5JIAzvLLc3iQYjEDw/BvfnEOH4p5cwDnt5ZKVaty38jG9Ua7DmCo857w
gtBbgkgJ1IYBNaz2ZVTlIaL5uflHlqn9/xDRXJeqUvK0W6C/o8RoUCpjevCKqdq4hNL1N7zWdi3drfvx
xZZuhlD6efQ8+h4AAP//WHvslwUJAAA=
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/recipe.json",
				Name:    "recipe.json",
				Size:    7258,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
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
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/timing.json",
				Name:    "timing.json",
				Size:    2476,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/7xWYWvcRhD97l8xXPLxIjltaIu/uTTYhRYCdWihFDPSjqRJpF2xM7I5gv97mV3pTmdf
YkpovvjOs6O3+57evL1PZwCbl1J3NODmAjad6nhRlh8k+Fe5WoTYli5io6/Ofyxz7cVmm55jtzwjF2UZ
8b5oWbupmoRiHbyS16IOQ1kRRYM8fBlQlGLap1Qe2LeFfc+4uhvJgEP1gWrNNXSOlYPH/l0MI0Vlks0F
NNgLpQZHDfvUYnXjBbB5L3STwXJhBS4a2bcJPNUdSR15NABb/oWbhiJ5ZVQS0I5ARqq54RrGGGoSAYOC
+448aMcC7NtIjskrLKcFFpiEXHHYh/w0bC7g7/n/TO1Ww+2A0u3bDuUqcH+i3FAcyCum8z5dHrH+iC1t
5oV/0udD7tvcJME/J8xK9VPC3HQE+Y1B7oRm6vsd5LZq1mrpaAD9QY171g5CghJoQoQKhQWCt37aQhvx
jnW3hRBhvAZUQL/byy1K40rI5xyRmsb10qeDTLbfUeUp0T871HQucJNZBfDoJPZq0TlyWyBbu8N+IiP8
HTjcZXoILu6gC+NKgjD1DirKzy7N7DUk3dav9RHhdMaXkRo73EAoUySserqdPKuk8XlRrqagvOEhu38P
8LAyipsizlS/pMJ1uIc+GHt/0uCRBmQv2zwC9ygQqaEYyYEGQMkKsk/sfiaKf/3+G4iidxhdAW+LqwIu
4c05DOwnpSSWWX6/g8BA6MXQUt103bfLFtA7wKwjLJweib8HiDSEOwJWwEYpzur/bxJbBLKfwiTPiHwJ
h9aj7JAxEjoIkwLdke93mW6yzuxK09ViKtKRPZMrf9jL+uu7S6h24ELbsHTQGariRxLAvofQGMrap7Js
lMZ5AO1imNrODrLa0F7II/WWFKlC6An9aWGWJL2dB/4Zed5LdhN7xzXqnLnrYFlm0QKFnAUKWiCx2XAf
2/Nmi+mcO0ynZsg/rgzpdXH++qevMMVV3ufzvhivvzFhahPdKpLOTMdrg/i+ePMVPC9rdl/kaT7870xR
c74+vlPnK2UtQYoa8rX9fIC3V3C5TuGUAU8vyzmBLRNSQrmjSRKqg3fPxvDicvZKLcXT9Cd5csUs8h4L
ufxKOaCcrT/t78PZw9m/AQAA//8w1oIIrAkAAA==
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/packaging_vessel.json",
				Name:    "packaging_vessel.json",
				Size:    1904,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/7yUwW7bMAyG73kKQu2xiXcbkNuAAbv2UPQyDAZt044aW1IpKkMx5N0HSc3iJE6CIlhv
yU/y/0SJ5p8ZgLr39YoGVEtQKxG3LIoXb808qwvLXdEwtjL/8rXI2p16SHW62dX4ZVEw/l50WlahCp64
tkbIyKK2Q1ERcbTc/xjQC3HiFA7rNXbadOWGvKd+EdVMkDdHEWGrF6ola9g0WrQ12D+ydcSiyasltNh7
SgkNtdqklKjHDgHU4w7ynBhP2TgHp0FJb8jXrF00i+EJF5gDgiOGfHhgckyejGAsAtvG8K4MHNuavF/s
CdfaSUluHNodGkAZHOhAGbXihbXp1L/Q9mF2nHOp7mEcIxMGtYSfIw1ArekwDUBVVqSnY7VGvz7WBM2J
1mpea3OsWlkRq5H2a7Knw6f68JXkN6KyQTm5mnumNtYPhD4wYdVTGYwWn0b1rhhNXPEdhdJ0TVK8IEsp
NDhilMA3oZ72NueJZJpP5XkhV4o+ncsPkfRw9RI7xo2Wt1swP7LF5cv7DE7uyK1ugXyrdXO1mf+MqJEr
a/DSN2jCUI0/53F5XqHlxvbhtvl5Tg7nz/kOeg1oZOJtz551dmSlmF6DZmoOlmNey+9/87KKldvZdvY3
AAD//xE4R0BwBwAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/fermentable.json",
				Name:    "fermentable.json",
				Size:    9588,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
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
HE5BwwTzxzE7nMi/v07+GwAA//9IhiABdCUAAA==
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/packaging.json",
				Name:    "packaging.json",
				Size:    938,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/5STTe6bMBDF95xi5GSZ4O4qse4BsqiyqarI4AGcgk3H41RRxd0rm0IITfVXdvZ8/ObN
M/zOAMTeVy32ShQgWuahkPLqnT1O0dxRIzWpmo+fPsspthOH1Gf03OMLKUn9yhvDbSiDR6qcZbScV66X
JSJF5OPQK89IaY4cVPVDNcY2ebxOaL4PGNmuvGLFU0xpbdg4q7oTuQGJDXpRQK06j6lAY21sKonxuBqA
OM30E7kKdSD8OrGn/OtZKa7RV2SGyIvpL+laogduEYaZBrUjWFaAuwsEJd6QVIP5A/aR+FQ0rFOzPgBh
VY9PkZVqz2RsI5bUeHj0TbJQX26uC/8i9oR1RPSofCBUZYeXYA379BA7ufJTnhMhWfdy1LNZbyu1jjdL
v7Wisc3lht5j91+IIlJ3cVinDGO/rV/ZskW/cGX5uM6p4tkegDHbnhbpgvBnMITxD/q2eem/1+/Z3Dlm
Y/YnAAD//zPpHMCqAwAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/measureable_units.json",
				Name:    "measureable_units.json",
				Size:    11643,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
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
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/culture.json",
				Name:    "culture.json",
				Size:    6286,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
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
drBz1eln9zNtlv7vZv8EAAD//9tQ1T6OGAAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/hop.json",
				Name:    "hop.json",
				Size:    6181,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/8xYT2/bxhK/+1MMlBzeAxQp7/IK+Fa3QeND4KCwcykCZUQOxQ2WO5vdoR228HcvhqRs
klpJdhIVPRigZufPzuzvNzvrv84AZi9jVlKFs3OYlSL+fLn8HNm96qQLDptlHrCQV69/WnayF7N5a2fy
rU08Xy4D3i02Rsp6XUcKGTshJ4uMq+WaKKjLx48Ko1Bo4yxL9gv96JxK40m98vozZdLJMM+NGHZo3wf2
FMRQnJ1DgTZSq5BTYVyronJNCmD2lv0HDIakucBID/J0jFaeU8yC8epHl8cOwAe+NTlFqJ35Ure/+72A
MJicnJiiAeNyc2vyGi0EyjjkEbgAhJI93HbuFo8h/TCh7QYBZg4rGkkG244SjNvMHpbu5492PnBeZxS+
3VZW7bE+25qD2Rj3LZYN4Tftt+BQHbGbD9fI1ar/x0Cm0q8ScICCXmwJi5QM/nNH8t/piidraceH57uc
wo7U1oOEAD4mk0PrS1xhtnsYLwMVmmJFGOtAuLa0qp2R2LLoxXJAhuV7Chk5udaqJMOsSU4V5WwSbRbo
S20C5aMz6JCeTvxsWKDez6yn5KXT48eeq8+l9q4TyNhayiSClAQoEsy6Ftolr5I9CgcCjD3DwTz6GXAb
rb0qRskmKzyu5aRrpQ5t5GVPAxkWYyw9TBLYS5QupcAV7hgojowIJb0pUS3e8pQHj/6WB40fFpfH3Dw5
zFZzovhx9Pt+QlvHkqjxgVaVcuI7pqwsR0n5msAUflVkUK5QU1j25rghhaWCsmULqDswDv4PFTspW9Aq
RnFDi93j/W5qpzKL9TqKkfoHFImNXfXjwxNqdGUs/NJpD3kIXQc4kP440Stjey/HczXulpxwaFLb20vr
y61VIsBZ6vs+2f/esv+5n4eux/x+xlwz9HCo8xFmZYsz4zaBcqNFLjhAHUnxhtoAjadWZW3sv6r5mUqR
dhxA1yVBp9xjBora2gY6tTV1ddlqFIAOtgMp3BkpgVtXsS3MGqOJwE71aQ6bgLdGmjlwAP8WUABdo5Nj
RjFCFPIHANqFTJDyul04jlOsuE5ziB1NTidV350dPaVjfGBbV7S7ucQGf0i8dxjjnmiH+/thzh2ZWvrS
DuaWHm9Jzl5e3LyJYioUmnJ2AsabSLlOGLkpCgr6pEAhuCtNVsLlxY1irKotgomwJgVkrQYDKi7gshhq
VZybwvQ6roE7bOZAm3aKqb3nIOo9WM9s9V6siGt5AHgEkkz/5uAt6StImf/pSkoKn1rAS0AXPQZy2fBd
s68V7XvvVCQl7x1Cxyd+eXHzrlXfM3BOSz9QTjTLyRC0O/zMfsfNaIqfXRsXScqh6DcMJH8OJW2N0nCY
3DXP7+CDG/Kxe6O14wtwOrhq4dE4RY0wsLHQ+5iDZ9v4khzbOAd0OUhp2MYFvNELYPCkzbjyXLtcoUVf
faAYu/EEt8OJxm0bJgvaNkrP4oqcPOHl29qtNMPKrjyF1f9ev5628WQLb+N1vNQ9sLFzMC6zda4Zl00e
OMOwZhfnQFEobFOl4MkRoM24ZBuVKW2zUS+V7X1peqBb2U5eFcY4bt4Pp+fqak0h/eAq66q25HYe9z/0
VZdhaNiXjT19JG4T4hOHqRpVOG2MAoOjeOooGwroDNvTvupX3rhTZ2KNQ8snzsSaik+eiaNw4jT+geMY
tPCTxvmKTkpl/SlKlrzCd95OT7gwj/3LGA5dQcm5ec/MvPf99z2z8mRO/q4Y6fl4OP1+3Ff/M/27P/s7
AAD//6A3XtwlGAAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/misc.json",
				Name:    "misc.json",
				Size:    3283,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/8yWwY7jNgyG73kKwrvHbNxbgdzaU3tYtIfZXorFQLHpmAtZ9JJUgqDIuxeyk4ntOJnZ
mUHRQwCHsn5SPz9J/mcBkH3UosbGZWvIarN2nefflMOnPrpi2ealuMo+/fRz3sc+ZMtuHpXnObrOc3H7
1ZasjpuoKAUHw2Crgpt8gyhJ8vLQODWUvH8mLVbpqVe1Q4tJljffsLA+VmJFgYw4aLaGVDRA9pm0QO9d
QI76q1N8GppX6eIlaiHUJqk0fKUBrfCOSlSIgb7H7n+LYoQKxkAlBqPqABRK2lEZnQfBgqVU4AoobAVL
wmAKVjsDJwgByWoUqLnVJQSWcw5wkFwS2sRUT9K3GmErbkd2SHp7Fltdir/UMlgqQBZcg6PIwAA1obDN
noaOy8u8VriMBcrr59pjx8APzz69c2/ecjiGITbZGv4exAAybanA0YsAWSJlMh0g2ztDAbfFYNOhGmVz
JeLdjuVKhLmcxjj1NhvEvl4WvJgsPBP8HkmwHC2lb9/UnMVQ7aQwxvVh7OGrkE8aULD3WHTIIjjrgcQO
aBcGTCdA1VgQnJ6oBwoVS+OS+ABU5/0f1WiRo05/FEyj2Yd8sLHz6/08h85I6MaG6Mai4mPFU7bnHPmi
WKa19eENdj60UVpWTCakv82wuIEnS8DVdgX7msRXngsghZj0KhYovBOyw2rCzJ09MllsjwfbzPp+UITC
DoOxHOaEXtCP38/zHwZwnlIt5p6PzwP8S1l2qd4D5KHWPaDRFfXNZnZNi4pAAVwinFr8n1Ft1KRePw/1
Q43Qvwy9j1BF7w9PkPfmnN/odro7eQh7shq4k9LOk41TUkhXFDW4PN9QS0h32W+Q7rlwSLdagaqghu0M
8md7+pTdhT926qEbmOFrgrJrOAabs4ADTho05+9VRQ06jYJu4/ExBjKdKe4v9rHB6+JmCnyXfJ+d6n+X
7Usgu5FtEvn6kq1/g+vZ2w8GLb17i945SMaH0wtOkjPpzv853GuV84rPfm3N8neDvZtH7VuYmxDwphzz
nL1nhnm2jm/6YhrzciJjkX7Hxb8BAAD//465i/LTDAAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/water.json",
				Name:    "water.json",
				Size:    2835,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/8xWMW/cPAzd71cQTsbk/G0fcFvbpVs7BO1QFAUt0WcGtuhQVIKgyH8vZF9yduIEbYoY
HQzIlMgn8lF6+rkBKE6ja6jDYgdFY9bvyvIySjgfrVvRfekVazv/7/9ytJ0UZ4Mf+3ufuCtLxZvtnq1J
VYqkToJRsK2TrqyINIc8DjqMRjrglDdopNs8HMPabU85rlSX5Gy0ofdsLAHbzyo9qTHFYgc1tpGGBZ5q
DsOSbM9pARRfc+T3GOnBtBx+sHuKTrnPIfL0gy/0KtfsKUIKfJWG/8MOwATYUzCub4GD52v2CVtQcqI+
gtQAldINhz2MWR7h+mke95sDKAJ2NLNMthxNOeyLh6m7s6Nfr+KTI32Nr8PWceoeu54q1dm1I4xJCauW
fqTAFgeyTspJzcsPEhwFU8y/FxlzEalih1pJQHuS5BugrYnVi2GM61SRdejRt4YJnJesUbyMxKsg1a0k
Zb8GVExtvVKfN62slZT4dXq8w32gtzxPm0eYhdJVYqWsad8e3ccLd+Uzl9oC/0tEPa3oYuoH2/fNZKej
Ml2M1/qrVC37gpO2JWcRrCFAM+UqGQ2ahXPRyioXTZQA40HagEMt2g1lPWJh236qZ9VbJG5O0VGjl9pg
FuAZyRznPj6xTaoSUleRFrPpCcpItNhC4BcUFCZdNB/fPc/bu8ND5m/4m8Z4iUdC1xwovH8/QS0KKRJw
AMxkck//BoHYSQq2VP8/OfBfpE0dzU/6b/O0yd/drwAAAP//Zm8EQxMLAAA=
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/boil_step.json",
				Name:    "boil_step.json",
				Size:    2307,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/7xVwY4bNwy9+yuISQ4N4PW0BYoCvrU99FogeysKQ57hzDDRSCpJ2dgW++8FJa899i4S
BGn2YngeyUdRfKT+XQE0b6WbcHbNFppJNW3b9oPEcFfRTeSx7dkNevf9z23F3jTrEkf9U4xs25bdcTOS
TnmfBbmLQTHopotzu0dko7z8mZ0oclvBSH4nimljn5VaHxIad9x/wE4r5vqelGJw/g+OCVkJpdnC4Lxg
cehxoFBcDLfSAJpfI/n3ium+Mlb05QwF71E6pmQsZl6Gwx04SMhghwXGxCgY1JkvxAEcWCmQOHYosobO
BdgjZMEeNILklCIrJMbiZyyyhhDDnX1TGCHZtWSmfyrnyeNkXcNxIvYpxnOsCz10E3kzby41fO6milNa
mp6uBaAJbsYrZHFZokxhbM6mx/Ul7vrivjhc1LHuFOeE7DTzszO8ZRyMZEYnmdHtPe5yIJUimzftovvt
/YWm9P3FjBj6V83Hbk47pee3eyO5+wnBzTEHNVFZAOjk1MQhKLDHIbJBJFWIexwpyAZwhDkeTEYDx/mk
RvssXt+V3x/emRLdjZJOxh/fweweQN1HhJ9gppAVZQPviw2OMfseJndAcGC11KPF4eK7hikmIIkznjVs
Et2TKnJAEeic77IvJoEjeQ8B63i4ris1D5FrcQZwT2H0Dxdtf3lraP5ET6zuF3vyf6YwYY/sDqQPX5Pm
90rxaUG/Rp5aUZq+JskvHfWfLeYbp3janDu9fhteGsrfTr5gviBYxt2m0Wk3nXfw2oaQ5hlZTPsFRhbb
3FhmFgGDEiMcos91eo72JpDAnmMeJ4U+HgNQgLqabKWAkzKy0eMaDgIUPAU8J12QP3HVxL2NuUegyo/m
27OjgH15S+zgLoDHOtNCY6CBOhd0sX+eT3Tsumx7vBxywvrofURVjzdzerP3r2wY8txs4c8FBtCU67xy
BGhqvc0C/OvSz9VNXxvGvzMx9lfk9VVbLaMt8nH1uPovAAD//6FPtmUDCQAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "/schema/fermentation_step.json",
				Name:    "fermentation_step.json",
				Size:    1676,
				Mode:    os.FileMode(436),
				ModTime: 1597743398,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/7xVzW7bTAy8+ykIJcfY+m4f4Fsv7bVAfSsKg5ZGFgNpVyWpuEHhdy9Wqmv5B06DoDnZ
JoczO8vB+ueMKLu3okbL2ZKy2r1b5vmjxTAfq4uo27xUrnz+3//5WLvLHoY5KQ8ztsxz5d1iK173m96g
RQyO4IsitvkG0ER5/NKyOXTQyStoi+DsEsPaHN0ilUcJf+6QNOLmEYWPNS5LSVhuPmvsoC6wbEkVN4YB
UKKSMEBSPVkkyj5OVL44utXIPHavKw31ElaodGksta/R0JyYOiilw5OiU9gBQ7EipqlF4iJ9LI4SLxka
QN20dTg1URa4xUll4sVcJWyzP639w3Hu1Nerx81Zfe1oOyh7rxdnuFdUiaQFW6/gTYN1H8Rt2O5dPllS
vjrSDGu5qohQvqteWuba5fJ2X6Uk7Q2JSoG1il1InIdOAUo4EqPeUJJHklBKwY7zeA0h3NVQkNcg/Ihe
Q1spTmFixE0TdyOZcmJP+MkV0068jr2TIkVhiC2tarE0DEm0tNIeFHUM6zHTJznaxNiAw60gbZWfxJ/f
ctWfRorbAXoPndFRV79F5EMh5Ytm/rHEE8zQ/PXjMDtjyBTfe1Gk/4ivZ+/V75/fZofJ/Ww/+xUAAP//
fDVAVYwGAAA=
`,
				Files: []*static.DirFile{},
			},
		},
	})
}
