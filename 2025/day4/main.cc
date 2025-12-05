#include <algorithm>
#include <expected>
#include <fstream>
#include <print>
#include <ranges>
#include <stdexcept>
#include <string>
#include <vector>

#include "../log.hpp"

using namespace std;

constexpr auto DIRECTIONS = to_array<pair<int, int>>({
    {-1, 0},  // up
    {1, 0},   // down
    {0, -1},  // left
    {0, 1},   // right
    {-1, -1}, // up-left
    {-1, 1},  // up-right
    {1, -1},  // down-left
    {1, 1}    // down-right
});

// count adjacent '@' symbols for a position
int countAdjacent(const vector<string> &grid, int rows_count, int cols_count,
                  int row, int col) {
  return ranges::count_if(
      DIRECTIONS,
      [&grid, rows_count, cols_count, row, col](const auto &dir) -> bool {
        const auto &[dx, dy] = dir;
        const int neighbor_row = row + dx;
        const int neighbor_col = col + dy;
        return neighbor_row >= 0 && neighbor_row < rows_count &&
               neighbor_col >= 0 && neighbor_col < cols_count &&
               grid[neighbor_row][neighbor_col] == '@';
      });
}

unsigned long processFile(ifstream &file) {
  vector<string> grid;

  string line;
  while (getline(file, line)) {
    if (!line.empty()) {
      grid.push_back(line);
    }
  }

  if (grid.empty()) {
    return 0;
  }

  const auto rows_count = static_cast<int>(grid.size());
  const auto cols_count = static_cast<int>(grid[0].size());

  auto positions = views::cartesian_product(views::iota(0, rows_count),
                                            views::iota(0, cols_count));

  auto at_symbols = positions | views::filter([&grid](const auto &pos) -> bool {
                      const auto &[row, col] = pos;
                      return grid[row][col] == '@';
                    });

  return ranges::count_if(
      at_symbols, [&grid, rows_count, cols_count](const auto &pos) -> bool {
        const auto &[row, col] = pos;
        return countAdjacent(grid, rows_count, cols_count, row, col) < 4;
      });
}

unsigned long processFile2(ifstream &file) {
  vector<string> grid;

  string line;
  while (getline(file, line)) {
    if (!line.empty()) {
      grid.push_back(line);
    }
  }

  if (grid.empty()) {
    return 0;
  }

  const auto rows_count = static_cast<int>(grid.size());
  const auto cols_count = static_cast<int>(grid[0].size());

  auto positions = views::cartesian_product(views::iota(0, rows_count),
                                            views::iota(0, cols_count));

  unsigned long total_removed = 0;
  while (true) {
    auto to_remove =
        positions | views::filter([&](const auto &pos) -> bool {
          const auto &[row, col] = pos;
          return grid[row][col] == '@' &&
                 countAdjacent(grid, rows_count, cols_count, row, col) < 4;
        }) |
        ranges::to<vector>();

    if (to_remove.empty()) {
      break;
    }

    ranges::for_each(to_remove, [&](const auto &pos) {
      const auto &[row, col] = pos;
      grid[row][col] = '.';
    });
    total_removed += to_remove.size();
  }

  return total_removed;
}

int main(int argc, char *argv[]) {
  string filename;
  if (argc > 1) {
    filename = argv[1];
  }

  ifstream file(filename);
  if (!file.is_open()) {
    log("could not open file({})", filename);
    return 1;
  }

  if (auto res = processFile(file)) {
    println("Result 1={}", res);
  }

  file.clear();  // clear EOF flag
  file.seekg(0); // rewind to beginning

  if (auto res = processFile2(file)) {
    println("Result 2={}", res);
  }

  return 0;
}
