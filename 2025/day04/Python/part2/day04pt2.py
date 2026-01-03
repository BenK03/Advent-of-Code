def main():
    total_removed = 0

    with open("input.txt", "r", encoding="utf-8") as file:
        matrix = []
        for line in file:
            matrix.append(list(line.rstrip("\n")))

    rows = len(matrix)

    # repeat until nothing can be removed
    while True:
        to_remove = []  # store (i, j) positions to remove this round

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
                        to_remove.append((i, j))

        # if nothing removable this round, stop
        if len(to_remove) == 0:
            break

        # remove everything found this round
        for (r, c) in to_remove:
            matrix[r][c] = "."

        total_removed += len(to_remove)

    print(total_removed)


main()