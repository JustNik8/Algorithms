# key - customer name, value - dict of products
customers = dict()

with open("task_f_input.txt") as file:
    for line in file:
        name, product, countStr = line.split()
        count = int(countStr)
        if name not in customers:
            customers[name] = {product: count}
        else:
            if product not in customers[name]:
                customers[name][product] = count
            else:
                customers[name][product] += count

sortedCustomers = sorted(customers.keys())

for customer in sortedCustomers:
    print(f'{customer}:')
    sortedProducts = sorted(customers[customer].keys())
    for product in sortedProducts:
        print(f'{product} {customers[customer][product]}')
