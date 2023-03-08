with open('task2_input.txt', 'r') as file:
    frequency = dict()
    for line in file:
        for word in line.split():
            if word not in frequency:
                print(0, end=' ')
                frequency[word] = 1
            else:
                print(frequency[word], end=' ')
                frequency[word] += 1
