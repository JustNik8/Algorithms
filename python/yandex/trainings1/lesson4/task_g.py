import random
import string

def generate_random_string(length):
    # define the character set to choose from
    characters = string.ascii_letters + string.digits

    # generate a random string of the specified length
    random_string = ''.join(random.choice(characters) for i in range(length))

    return random_string


n = 1_000_000

with open('text.txt', 'a') as file:
    # append to file
    file.write(f'{n}\n')
    d = dict()

    for i in range(0, n):
        first = generate_random_string(100)
        second = generate_random_string(100)
        file.write(f'{first} {second}\n')
        d[first] = second
        print(i)

    random_key = random.choice(list(d.keys()))
    file.write(random_key)

    file.close()



