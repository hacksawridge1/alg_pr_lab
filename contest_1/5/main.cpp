#include <iostream>
#include <math.h>

double m;
double h;

int main() {
  std::cout << "";
  std::cin >> m >> h;
  double result {
    m / pow(h,2)
  };
  std::cout << result << std::endl;
}