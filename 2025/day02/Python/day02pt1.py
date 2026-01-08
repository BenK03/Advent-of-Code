def main():
    ans = 0

    with open("input.txt", "r", encoding="utf-8") as file:
        s = file.read()

    for part in s.split(","):
        start, end = map(int, part.split("-"))

        for i in range(start, end + 1):
            s = str(i)

            # odd number of digits can not be invalid
            if len(s) % 2 != 0:
                continue
            
            mid = len(s) // 2

            # check if first half == last half, if so add to ans
            if s[:mid] == s[mid:]:
                ans += i

    print(ans)

main()