#include <cmath>
#include <expected>
#include <format>
#include <fstream>
#include <print>
#include <sstream>
#include <stdexcept>
#include <string>
#include <vector>

#include "../log.hpp"

using namespace std;

using Pairs = vector<pair<unsigned long, unsigned long>>;

expected<Pairs, runtime_error> parseRanges(const string &line) {
  vector<pair<unsigned long, unsigned long>> ranges;
  stringstream ss(line);
  string s;

  while (getline(ss, s, ',')) {
    size_t dashPos = s.find('-');
    if (dashPos != string::npos) {
      try {
        unsigned long start = stoul(s.substr(0, dashPos));
        unsigned long end = stoul(s.substr(dashPos + 1));
        ranges.push_back({start, end});
      } catch (const exception &e) {
        return unexpected(runtime_error(
            format("failed to parse range({}): ", s).append(e.what())));
      }
    }
  }

  return ranges;
}

expected<unsigned long, runtime_error> processFile(ifstream &file) {
  unsigned long res = 0;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    auto ranges = parseRanges(line);
    if (!ranges.has_value()) {
      return unexpected(ranges.error());
    }

    for (const auto &range : ranges.value()) {
      log("Range: {}-{}", range.first, range.second);
      for (auto i = range.first; i < range.second + 1; i++) {
        auto s = to_string(i);
        if (s.length() % 2 == 1) {
          continue;
        }
        if (s.substr(0, s.length() / 2) ==
            s.substr(s.length() / 2, s.length() / 2)) {
          log("Found special: {}", i);
          res += i;
        }
      }
    }
  }
  return res;
}

expected<unsigned long, runtime_error> processFileV2(ifstream &file) {
  unsigned long res = 0;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    auto ranges = parseRanges(line);
    if (!ranges.has_value()) {
      return unexpected(ranges.error());
    }

    for (const auto &range : ranges.value()) {
      log("Range: {}-{}", range.first, range.second);
      for (auto i = range.first; i < range.second + 1; i++) {
        auto s = to_string(i);
        bool found = false;
        auto len = s.length();

        for (size_t patternLen = 1; patternLen < len / 2 + 1; patternLen++) {
          if (len % patternLen != 0) {
            continue;
          }

          string pattern = s.substr(0, patternLen);
          bool matches = true;

          // check if repeating this pattern creates the full string
          for (auto j = patternLen; j < len; j += patternLen) {
            if (s.substr(j, patternLen) != pattern) {
              matches = false;
              break;
            }
          }

          if (matches) {
            found = true;
            break;
          }
        }

        if (found) {
          log("Found special: {}", i);
          res += i;
        }
      }
    }
  }
  return res;
}

int main(int argc, char *argv[]) {
  string filename;
  if (argc > 1) {
    filename = argv[1];
  }

  ifstream file(filename);
  if (!file.is_open()) {
    log("could not open file=({})", filename);
    return 1;
  }

  if (auto res = processFile(file); res.has_value()) {
    print("Result 1={}\n", res.value());
  } else {
    log("processFile(): {}", res.error().what());
    return 1;
  }

  file.clear();  // clear EOF flag
  file.seekg(0); // rewind to beginning

  if (auto res = processFileV2(file); res.has_value()) {
    print("Result 2={}\n", res.value());
  } else {
    log("processFileV2(): {}", res.error().what());
    return 1;
  }
  return 0;
}