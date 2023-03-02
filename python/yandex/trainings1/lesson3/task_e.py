digits = set(map(int, input().split()))
digits_to_add = set()

num = int(input())

while num > 0:
    digit = num % 10
    if digit not in digits and digit not in digits_to_add:
        digits_to_add.add(digit)
    num = num // 10

print(len(digits_to_add))

