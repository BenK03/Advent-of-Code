def main():
    ranges = []
    my_set = set() 
    with open("input.txt", "r", encoding="utf-8") as file:
        for line in file:
            line = line.strip()
            if line == "":
                break

            split = line.split("-")
            front = int(split[0])
            rear = int(split[-1])
            ranges.append([front, rear])

    ranges.sort()

    fixed_ranges = []
    for r in ranges:
        front = r[0]
        rear = r[1]

        if len(fixed_ranges) == 0:
            fixed_ranges.append([front, rear])
        else:
            last = fixed_ranges[len(fixed_ranges) - 1]
            last_end = last[1]

            if front <= last_end + 1:
                if rear > last_end:
                    last[1] = rear
            else:
                fixed_ranges.append([front, rear])

    total = 0
    for r in fixed_ranges:
        total += (r[1] - r[0] + 1)

    print(total)


main()
