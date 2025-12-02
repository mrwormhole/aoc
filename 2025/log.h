#pragma once

#include <format>
#include <print>
#include <utility>

constexpr bool ENABLE_LOGGING = false;

template<typename... Args>
constexpr void log(std::format_string<Args...> fmt, Args&&... args) {
  if constexpr (ENABLE_LOGGING) {
    std::print(fmt, std::forward<Args>(args)...);
    std::print("\n");
  }
}
