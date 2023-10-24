#include <numeric>

std::tuple<int, int, int> find_lcm(int a, int b) {
    int c = std::lcm(a, b);
    return std::make_tuple(c, c/a, c/b);
}

std::tuple<int, int> reduce(int a, int b) {
    int c = std::gcd(a, b);
    return std::make_tuple(a/c, b/c);
}