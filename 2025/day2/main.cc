#include <cmath>
#include <fstream>
#include <print>
#include <sstream>
#include <string>
#include <vector>

#include "../log.hpp"

using namespace std;

vector<pair<unsigned long, unsigned long>>
parseRanges(const std::string &line) {
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
      } catch (const std::exception &e) {
        log("ERR: Invalid range format: {}", s);
      }
    }
  }

  return ranges;
}

unsigned long processFile(ifstream &file) {
  unsigned long res = 0;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    auto ranges = parseRanges(line);

    for (const auto &range : ranges) {
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

vector<size_t> computePatternLengths(size_t len) {
  vector<size_t> patternLengths;
  auto sqrtLen = static_cast<size_t>(sqrt(len)) + 1;

  for (size_t i = 1; i < sqrtLen; i++) {
    if (len % i == 0) {
      patternLengths.push_back(i);

      if (i != len / i) {
        patternLengths.push_back(len / i);
      }
    }
  }
  return patternLengths;
}

unsigned long processFileV2(ifstream &file) {
  unsigned long res = 0;

  string line;
  while (getline(file, line)) {
    if (line.empty()) {
      continue;
    }

    auto ranges = parseRanges(line);

    for (const auto &range : ranges) {
      log("Range: {}-{}", range.first, range.second);
      for (auto i = range.first; i < range.second + 1; i++) {
        auto s = to_string(i);
        bool found = false;
        auto len = s.length();

        auto patternLengths = computePatternLengths(len);

        for (auto patternLen : patternLengths) {
          if (patternLen == len) {
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
    log("ERR: Could not open file={}", filename);
    return 1;
  }

  unsigned long res = processFile(file);
  std::print("Result 1={}\n", res);

  file.clear();  // clear EOF flag
  file.seekg(0); // rewind to beginning

  unsigned long res2 = processFileV2(file);
  std::print("Result 2={}\n", res2);

  return 0;
}