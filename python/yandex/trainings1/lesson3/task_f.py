genom1 = input()
genom2 = input()

pairs2 = set()
genom2_len = len(genom2)

for i in range(1, genom2_len):
    pair = f'{genom2[i - 1]}{genom2[i]}'
    pairs2.add(pair)

genom1_len = len(genom1)
ans = 0
for i in range(1, genom1_len):
    pair = f'{genom1[i - 1]}{genom1[i]}'
    if pair in pairs2:
        ans += 1

print(ans)
