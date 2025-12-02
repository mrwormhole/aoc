#include <fstream>
#include <print>
#include <string>

#include "../log.h"

constexpr uint START = 50;

unsigned int processFile(std::ifstream &file) {
  unsigned int count = 0;
  unsigned int res = START;

  std::string line;
  while (std::getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    char rotation = line[0];
    unsigned int units;
    try {
      units = std::stoi(line.substr(1));
    } catch (const std::exception &e) {
      log("ERR: Invalid format with line={}", line);
      continue;
    }
    log("Rotation={}, Units={}", rotation, units);

    units = units % 100;
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

unsigned int processFileV2(std::ifstream &file) {
  unsigned int count = 0;
  unsigned int res = START;

  std::string line;
  while (std::getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    char rotation = line[0];
    unsigned int units;
    try {
      units = std::stoi(line.substr(1));
    } catch (const std::exception &e) {
      log("ERR: Invalid format with line={}", line);
      continue;
    }
    log("Rotation={}, Units={}", rotation, units);

    // count how many times we land on 0 during this rotation
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
  std::string filename;
  if (argc > 1) {
    filename = argv[1];
  }

  std::ifstream file(filename);
  if (!file.is_open()) {
    log("ERR: Could not open file={}", filename);
    return 1;
  }

  unsigned int res = processFile(file);
  std::print("Result 1={}\n", res);

  file.clear();  // clear EOF flag
  file.seekg(0); // rewind to beginning

  unsigned int res2 = processFileV2(file);
  std::print("Result 2={}\n", res2);
  return 0;
}
