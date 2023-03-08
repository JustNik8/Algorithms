n = int(input())

# key - width, value - set of heights
blocks = dict()
for i in range(n):
    width, height = map(int, input().split())
    try:
        heights = blocks[width]
        heights.add(height)
    except:
        blocks[width] = {height}

sortedWidths = sorted(blocks.keys(), reverse=True)
maxWidth = 0
for width in sortedWidths:
    maxWidth += max(blocks[width])

print(maxWidth)
