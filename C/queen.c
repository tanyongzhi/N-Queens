#include <stdio.h>
#include <stdlib.h>

int N;

// check if spot is safe to place queen
int is_safe(int* grid, int x, int y) {
    // horizontal
    for (int i = 0; i < N; i++) {
        if (i != x && grid[y*N + i])
            return 0;
    }
    // vertical
    for (int i = 0; i < N; i++) {
        if (i != y && grid[i*N + x])
            return 0;
    }
    // diagonal
    int directions[4][2] = {{-1, 1}, {1, -1}, {1, 1,}, {-1, -1}};
    int diag_cnt_x, diag_cnt_y;

    for (int i = 0; i < 4; i++) {
        diag_cnt_x = x;
        diag_cnt_y = y;

        while (diag_cnt_x >= 0 && diag_cnt_x < N && diag_cnt_y >= 0 && diag_cnt_y < N) {
            if (x != diag_cnt_x && y != diag_cnt_y && grid[diag_cnt_y*N + diag_cnt_x]) 
                return 0;
            diag_cnt_x += directions[i][0];
            diag_cnt_y += directions[i][1];
        }
    }
    return 1;
}

// pretty print
void pretty_print(int* grid) {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            if (grid[i*N + j])
                printf("Q ");
            else
                printf("X ");
        }
        printf("\n");
    }
}

// recursive helper function to solve problem
int helper(int* grid, int y) {
    if (y >= N) {
        return 1;
    }
    for (int i = 0; i < N; i++) {
        if (is_safe(grid, i, y)) {
            grid[y*N + i] = 1;
            if (helper(grid, y+1))
                return 1;
            else
                grid[y*N + i] = 0;
        }
    }
    return 0;
}

int main() {
    printf("Specify N size: ");
    scanf("%d", &N);

    int* grid = (int*)calloc(N * N, sizeof(int));
    helper(grid, 0);

    pretty_print(grid);

    free(grid);
}