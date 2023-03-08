frequency = dict()

with open('task3_input.txt', 'r') as file:
    for line in file:
        for word in line.split():
            if word not in frequency:
                frequency[word] = 0
            else:
                frequency[word] += 1


sortedKeys = sorted(frequency.keys())
mostCommonWord = sortedKeys[0]
maxCnt = frequency[mostCommonWord]

for key in sortedKeys:
    if frequency[key] > maxCnt:
        maxCnt = frequency[key]
        mostCommonWord = key

print(mostCommonWord)
