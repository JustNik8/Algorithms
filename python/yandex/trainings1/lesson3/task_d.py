words = set()

with open('input.txt', 'r') as file:
    for line in file:
        for word in line.split():
            words.add(word)

print(len(words))
