def main():
    with open("input.txt", "r", encoding="utf-8") as file:
        nums = []
        for line in file:
            split = line.split()
            nums.append(split)

        operations = nums[-1]
        nums = nums[:-1]

        cols = len(operations)
        res = 0

        for col in range(cols):
            col_vals = []
            for row in range(len(nums)):
                col_vals.append(int(nums[row][col]))

            if operations[col] == "+":
                res += sum(col_vals)
            else:
                holder = 1
                for vals in col_vals:
                    holder *= vals
                
                res += holder

    print(res)

main()



        
