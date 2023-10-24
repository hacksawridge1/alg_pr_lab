#include <iostream>

double a;
double b;

int main() {
  std::cout << "";
  std::cin >> a >> b;
  a > b ? std::cout << a << std::endl : std::cout << b << std::endl;
}