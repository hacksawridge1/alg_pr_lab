#include <iostream>

int main() {
  long double a;
  long double b;
  std::cout << "" << std::endl;
  std::cin >> a >> b;
  long double result {
    a / b
  };
  std::cout << result << std::endl;
}