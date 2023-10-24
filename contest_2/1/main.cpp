#include <cmath>

int is_prime(int num) {
  for (int i = 2; i <= std::sqrt(num); i++) {
      if (num % i == 0) {
          return 0;
      }
  }
  return 1;
}