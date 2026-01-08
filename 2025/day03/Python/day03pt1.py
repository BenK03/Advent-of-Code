def main():
    with open("input.txt", "r", encoding="utf-8") as file:
        ans = []

        for line in file:
            possible = []
            for i in range(len(line) - 1):
                for j in range(i + 1, len(line)):
                    possible.append(int(line[i] + line[j]))

            ans.append(max(possible))
        
    print(sum(ans))

main()
