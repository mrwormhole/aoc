import math

current_path = []
folders_sizes = dict()
folders_tree = dict()

def parse(full_command):
    lines = full_command.split("\n")
    command = lines[0]
    command_outputs = lines[1:]

    parts = command.split(" ")
    if parts[0] == "cd":
        if parts[1] == "..":
            current_path.pop()
        else:
            current_path.append(parts[1])
        return

    abspath = "/".join(current_path)

    sizes = []
    for line in command_outputs:
        if not line.startswith("dir"):
            sizes.append(int(line.split(" ")[0]))
        else:
            dir_name = line.split(" ")[1]
            try:
                folders_tree[abspath].append(f"{abspath}/{dir_name}")
            except KeyError:
                temp = []
                temp.append(f"{abspath}/{dir_name}")
                folders_tree[abspath] = temp
    folders_sizes[abspath] = sum(sizes)

def search(abspath):
    size = folders_sizes[abspath]
    try:
        for folder in folders_tree[abspath]:
            size += search(folder)
    except:
        return size
    return size

def solve():
    with open("input.txt") as f:
        commands = ("\n" + f.read().strip()).split("\n$ ")[1:]
    for command in commands:
        parse(command)

    result_f = 0
    for f in folders_sizes:
        print("abspath:", f)
        acc_value = search(f)
        if acc_value <= 100000:
            result_f += acc_value
    print("part 1: ", result_f)

    result_s = math.inf
    required = 30000000 - 70000000 + search("/")
    for f in folders_sizes:
        acc_value = search(f)
        if acc_value >= required:
            result_s = min(result_s, acc_value)
    print("part 2: ", result_s)

solve()