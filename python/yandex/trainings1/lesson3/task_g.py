n = int(input())
ans = set()

for i in range(0, n):
    a, b = map(int, input().split())
    if a >= 0 and b >= 0 and a + b == n - 1 and a not in ans:
        ans.add(a)

print(len(ans))
