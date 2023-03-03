n = int(input())

langs = dict()

for i in range(0, n):
    m = int(input())
    for j in range(0, m):
        lang = input()
        if lang not in langs:
            langs[lang] = 1
        else:
            langs[lang] += 1

allKnowsCnt = 0
allKnows = []
for k, v in langs.items():
    if v == n:
        allKnows.append(k)
        allKnowsCnt += 1

print(allKnowsCnt)
for i in range(0, allKnowsCnt):
    print(allKnows[i])

print(len(langs))
for k in langs:
    print(k)
