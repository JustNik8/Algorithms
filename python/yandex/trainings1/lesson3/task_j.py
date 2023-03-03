def extend(rect: (), d: int) -> ():
    minPlus, maxPlus, minMinus, maxMinus = rect
    return minPlus - d, maxPlus + d, minMinus - d, maxMinus + d


def intersection(rect1: (), rect2: ()) -> ():
    return max(rect1[0], rect2[0]), min(rect1[1], rect2[1]), max(rect1[2], rect2[2]), min(rect1[3], rect2[3])


def main():
    posRect = (0, 0, 0, 0)
    t, d, n = map(int, input().split())
    for _ in range(0, n):
        posRect = extend(posRect, t)
        navX, navY = map(int, input().split())
        navRect = extend((navX + navY, navX + navY, navX - navY, navX - navY), d)
        posRect = intersection(posRect, navRect)

    points = []
    for xPlusY in range(posRect[0], posRect[1] + 1):
        for xMinusY in range(posRect[2], posRect[3] + 1):
            if (xPlusY + xMinusY) % 2 == 0:
                x = (xPlusY + xMinusY) // 2
                y = xPlusY - x
                points.append((x,y))

    print(len(points))
    for p in points:
        print(*p)


main()
