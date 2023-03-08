import time

start_time = time.time()

first = {}
second = {}

n = int(input())

for i in range(n):
    first_word, second_word = input().split()

    first[first_word] = second_word
    second[second_word] = first_word

word = input()

try:
    print(first[word])
except:
    print(second[word])

end_time = time.time()
elapsed_time = end_time - start_time
print(f"Elapsed time: {elapsed_time}")