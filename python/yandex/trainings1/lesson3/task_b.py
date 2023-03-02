nums1 = set(map(int, input().split(" ")))
nums2 = set(map(int, input().split(" ")))

for num in nums1:
    if num in nums2:
        print(num, end=" ")
