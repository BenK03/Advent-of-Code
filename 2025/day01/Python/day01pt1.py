def main():

    land_zero = 0
    tracker = 50

    with open('input.txt', 'r', encoding='utf-8') as file:
        for line in file:
            dist = int(line[1:])
            if line[0] == 'L':
                tracker = (tracker - dist) % 100
                if tracker < 0:
                    tracker += 100
                if tracker == 0:
                    land_zero += 1

            # R operation
            else:
                tracker = (tracker + dist) % 100
                if tracker == 0:
                    land_zero += 1
    
    print(land_zero)



main()