def main():
    count = 0
    with open("input.txt", "r", encoding="utf-8") as file:
        matrix = []
        for line in file:
            matrix.append(list(line.rstrip("\n")))

    rows = len(matrix)

    for i in range(rows):
        cols = len(matrix[i])
        for j in range(cols):
            if matrix[i][j] == ".":
                continue
            else:
                valid = 0
                # check north
                if (i - 1 >= 0) and (matrix[i - 1][j] == "@"):
                    valid += 1

                # check east
                if (j + 1 <= cols - 1) and (matrix[i][j + 1] == "@"):
                    valid += 1

                # check south
                if (i + 1 <= len(matrix) - 1) and (matrix[i + 1][j] == "@"):
                    valid += 1

                # check west
                if (j - 1 >= 0) and (matrix[i][j - 1] == "@"):
                    valid += 1

                # check northwest
                if (j - 1 >= 0 and i - 1 >= 0) and (matrix[i - 1][j - 1] == "@"):
                    valid += 1

                # check northeast
                if (i - 1 >= 0 and j + 1 <= cols - 1) and (matrix[i - 1][j + 1] == "@"):
                    valid += 1

                # check southeast
                if (i + 1 <= len(matrix) - 1 and j + 1 <= cols - 1) and (matrix[i + 1][j + 1] == "@"):
                    valid += 1

                # check southwest
                if (i + 1 <= len(matrix) - 1 and j - 1 >= 0) and (matrix[i + 1][j - 1] == "@"):
                    valid += 1

            if valid < 4:
                count += 1


    print(count)

main()
                
                


