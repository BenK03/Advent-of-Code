def main():
    ans = 0

    with open("input.txt", "r", encoding="utf-8") as file:
        s = file.read()

    for part in s.split(","):
        start, end = map(int, part.split("-"))

        for i in range(start, end + 1):
            s = str(i)
            n = len(s)

            for j in range(1, len(s) // 2 + 1):

                pattern = s[:j]
                repeats =  n // j

                # check if pattern repeated is the same as s
                if pattern * repeats == s: # if yes invalid
                    ans += int(s)
                    break

    print(ans)

main()