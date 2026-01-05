def main():
    with open("input.txt", "r", encoding="utf-8") as file:
        lines = []
        for line in file:
            lines.append(line.strip())

        ranges = []
        ids = []

        toggle = True

        for line in lines:
            if line == "":
                toggle = False
                continue
                
            if toggle:
                split = line.split("-")
                front = int(split[0])
                rear = int(split[-1])
                ranges.append([front, rear])
            
            else:
                ids.append(int(line))

        res = 0
        for i in ids:
            for r in ranges:
                if r[0] <= i <= r[-1]:
                    res += 1
                    break # don't over count ids that fall into multiple ranges
            
    print(res)


main()