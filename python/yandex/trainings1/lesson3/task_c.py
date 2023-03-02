n, k = map(int, input().split(" "))

anya_cubes = set()
borya_cubes = set()

for i in range(0, n):
    anya_cubes.add(int(input()))

for i in range(0, k):
    borya_cubes.add(int(input()))

common_cubes = set()
for cube in anya_cubes:
    if cube in borya_cubes:
        common_cubes.add(cube)

for cube in common_cubes:
    if cube in anya_cubes and cube in borya_cubes:
        anya_cubes.remove(cube)
        borya_cubes.remove(cube)

print(len(common_cubes))
print(*sorted(common_cubes))

print(len(anya_cubes))
print(*sorted(anya_cubes))

print(len(borya_cubes))
print(*sorted(borya_cubes))
