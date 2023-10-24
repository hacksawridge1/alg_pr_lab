#include <iostream>
#include <math.h>

int main()
{   
  int a , b , c {};
  std::cin >> a >> b >> c;
  int ab {a-b};
  int ac {a-c};
  if(abs(ab) < abs(ac))
    std::cout << "B " << abs(ab);
  else if(abs(ac) < abs(ab))
    std::cout << "C " << abs(ac);
  else
    std::cout << "Равны";
}