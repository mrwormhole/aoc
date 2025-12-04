#include <expected>
#include <format>
#include <fstream>
#include <print>
#include <stdexcept>
#include <string>

#include "../log.hpp"

constexpr uint START = 50;

using namespace std;

expected<unsigned int, runtime_error> to_uint(string_view sv) {
  unsigned int result;
  auto [_, ec] = from_chars(sv.data(), sv.data() + sv.size(), result);
  if (ec == errc{0}) { // no err
    return result;
  }

  auto err_msg = make_error_code(ec).message();
  return unexpected(runtime_error(
      format("failed parsing number ({}): {}", sv.data(), err_msg)));
};

expected<unsigned int, runtime_error> processFile(ifstream &file) {
  unsigned int count = 0;
  unsigned int res = START;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    auto units_result = to_uint(line.substr(1));
    if (!units_result.has_value()) {
      return unexpected(
          runtime_error(format("failed to parse units from line({}): ", line)
                            .append(units_result.error().what())));
    }
    auto units = units_result.value() % 100;

    char rotation = line[0];
    if (rotation == 'R') {
      res = (res + units) % 100;
    } else if (rotation == 'L') {
      res = (res + 100 - units) % 100;
    }

    if (res == 0) {
      count += 1;
    }
  }
  return count;
}

expected<unsigned int, runtime_error> processFileV2(ifstream &file) {
  unsigned int count = 0;
  unsigned int res = START;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    auto units_result = to_uint(line.substr(1));
    if (!units_result.has_value()) {
      return unexpected(
          runtime_error(format("failed to parse units from line({}): ", line)
                            .append(units_result.error().what())));
    }
    auto units = units_result.value();

    char rotation = line[0];
    if (rotation == 'R') {
      int zero_count = units / 100;
      if (res + (units % 100) >= 100) {
        zero_count += 1;
      }
      count += zero_count;
      res = (res + units) % 100;
    } else if (rotation == 'L') {
      unsigned int dist_to_zero = (res == 0) ? 100 : res;
      if (units >= dist_to_zero) {
        count += 1;                            // hitting 0 first time
        count += (units - dist_to_zero) / 100; // every 100 clicks after that
      }
      res = (res - (units % 100) + 100) % 100;
    }

    log("Current={}, ZerosPassed={}", res, count);
  }
  return count;
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

  if (auto res = processFile(file); res.has_value()) {
    println("Result 1={}\n", res.value());
  } else {
    log("processFile(): {}", res.error().what());
    return 1;
  }

  file.clear();  // clear EOF flag
  file.seekg(0); // rewind to beginning

  if (auto res = processFileV2(file); res.has_value()) {
    println("Result 2={}", res.value());
  } else {
    log("processFileV2(): {}", res.error().what());
    return 1;
  }
  return 0;
}
