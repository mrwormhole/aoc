#include <expected>
#include <fstream>
#include <print>
#include <ranges>
#include <stdexcept>
#include <string>
#include <vector>

#include "../log.hpp"

using namespace std;

int findLargestJoltage(const string &batteries) {
  int max_joltage = 0;

  for (size_t i = 0; i < batteries.size(); ++i) {
    for (size_t j = i + 1; j < batteries.size(); ++j) {
      int first_digit = batteries[i] - '0';
      int second_digit = batteries[j] - '0';
      int joltage = first_digit * 10 + second_digit;
      max_joltage = max(max_joltage, joltage);
    }
  }
  return max_joltage;
}

unsigned long processFile(ifstream &file) {
  unsigned long sum = 0;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    sum += findLargestJoltage(line);
  }
  return sum;
}

expected<unsigned long, runtime_error>
findLargestJoltage12(const string &batteries) {
  const int k = 12;
  int to_remove = batteries.size() - k;
  string result;

  for (char digit : batteries) {
    while (!result.empty() && result.back() < digit && to_remove > 0) {
      result.pop_back();
      to_remove--;
    }
    result.push_back(digit);
  }
  result.resize(k);

  auto to_ulong = [](string_view sv) -> expected<unsigned long, runtime_error> {
    unsigned long result;
    auto [_, ec] = from_chars(sv.data(), sv.data() + sv.size(), result);
    if (ec == errc{0}) { // no err
      return result;
    }

    auto err_msg = make_error_code(ec).message();
    return unexpected(runtime_error(
        format("failed parsing number ({}): {}", sv.data(), err_msg)));
  };
  return to_ulong(result);
}

unsigned long processFileV2(ifstream &file) {
  unsigned long sum = 0;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    auto res = findLargestJoltage12(line);
    if (res.has_value()) {
      sum += res.value();
    } else {
      log("failed to find largest joltage on line({}): {}", line,
          res.error().what());
    }
  }
  return sum;
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

  if (auto res = processFileV2(file)) {
    println("Result 2={}", res);
  }
  return 0;
}
