nums = list(map(int, input().split(" ")))
numsSet = set(nums)

diff = len(nums) - len(numsSet)
print(len(nums) - diff)
