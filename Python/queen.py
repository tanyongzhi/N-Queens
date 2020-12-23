# check if location is safe to place a queen
def is_safe(grid, x, y, N):
    # horizontal
    for i in range(N):
        if i != x and grid[y][i] == 1:
            return False
    # vertical
    for i in range(N):
        if i != y and grid[i][x] == 1:
            return False
    # diagonal
    for i, j in [(-1, 1), (1, -1), (1, 1), (-1, -1)]:
        diag_cnt_x, diag_cnt_y = x, y
        while 0 <= diag_cnt_x < N and 0 <= diag_cnt_y < N:
            if x != diag_cnt_x and y != diag_cnt_y and grid[diag_cnt_y][diag_cnt_x] == 1:
                return False
            diag_cnt_x, diag_cnt_y = (diag_cnt_x+i, diag_cnt_y+j)
    return True

# recursive helper function
def helper(grid, y, N):
    if y >= N:
        return True
    for i in range(N):
        if is_safe(grid, i, y, N):
            grid[y][i] = 1
            if helper(grid, y+1, N):
                return True
            else:
                grid[y][i] = 0

    return False

# recursive function wrapper
def solve(N):
    grid = [[0] * N for i in range(N)]
    helper(grid, 0, N)

    return grid

# pretty print
def print_solution(grid, N):
    for i in range(N):
        for j in range(N):
            if grid[i][j]:
                print('Q', end=' ')
            else:
                print('X', end=' ')
        print('\n')

N = int(input('Specify N size: '))
grid = solve(N)
print_solution(grid, N)