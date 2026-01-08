def main():
    res = 0
    with open("input.txt", "r", encoding="utf-8") as file:
        for line in file:
            line = line.strip()
            
            k = len(line) - 12
            stack = []

            for c in line:
                while k > 0 and stack and stack[-1] < c:
                    stack.pop()
                    k -= 1
                stack.append(c)

            if k > 0:
                stack = stack[:-k]

            res += int("".join(stack[:12]))

    print(res)

main()