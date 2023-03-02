n = int(input())

diff_x = set()

for i in range(0, n):
    x, _ = map(int, input().split())
    if x not in diff_x:
        diff_x.add(x)

print(len(diff_x))
