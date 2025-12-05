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

// 8 directions
constexpr auto directions = to_array<pair<int, int>>({
    {-1, 0},  // up
    {1, 0},   // down
    {0, -1},  // left
    {0, 1},   // right
    {-1, -1}, // up-left
    {-1, 1},  // up-right
    {1, -1},  // down-left
    {1, 1}    // down-right
});

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

  const auto rows = static_cast<int>(grid.size());
  const auto cols = static_cast<int>(grid[0].size());

  auto positions =
      views::cartesian_product(views::iota(0, rows), views::iota(0, cols));

  auto at_symbols = positions | views::filter([&](const auto &pos) {
                      const auto &[row, col] = pos;
                      return grid[row][col] == '@';
                    });

  return ranges::count_if(at_symbols, [&](const auto &pos) {
    const auto &[row, col] = pos;

    // count adjacent '@' symbols
    auto adjacent_count = directions | views::transform([&](const auto &dir) {
                            const auto &[dr, dc] = dir;
                            const int new_row = row + dr;
                            const int new_col = col + dc;
                            return make_pair(new_row, new_col);
                          }) |
                          views::filter([&](const auto &neighbor) {
                            const auto &[r, c] = neighbor;
                            return r >= 0 && r < rows && c >= 0 && c < cols;
                          }) |
                          views::filter([&](const auto &neighbor) {
                            const auto &[r, c] = neighbor;
                            return grid[r][c] == '@';
                          }) |
                          ranges::to<vector>();

    return adjacent_count.size() < 4;
  });
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

  return 0;
}
