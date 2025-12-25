def main():
    pass_zero = 0
    tracker = 50

    with open('input.txt', 'r', encoding='utf-8') as file:
        for line in file:
            operation = line[0]
            dist = int(line[1:])

            start = tracker

            if operation == 'L':
                first_hit = start % 100
                if first_hit == 0:
                    first_hit = 100

                if dist >= first_hit:
                    hits = 1
                    remaining = dist - first_hit
                    hits += remaining // 100
                    pass_zero += hits

                tracker = (tracker - dist) % 100

            else:  # R
                first_hit = (100 - start) % 100
                if first_hit == 0:
                    first_hit = 100

                if dist >= first_hit:
                    hits = 1
                    remaining = dist - first_hit
                    hits += remaining // 100
                    pass_zero += hits

                tracker = (tracker + dist) % 100

    print(pass_zero)

main()