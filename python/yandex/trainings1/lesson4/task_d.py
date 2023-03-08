n = int(input())
keys = dict()

inputKeys = list(map(int, input().split()))
for i in range(0, n):
    keys[i + 1] = inputKeys[i]

k = int(input())

clicks = list(map(int, input().split()))
for i in range(0, k):
    keys[clicks[i]] -= 1

sortedKeys = sorted(keys.keys())
for key in sortedKeys:
    if keys[key] >= 0:
        print("NO")
    else:
        print("YES")

