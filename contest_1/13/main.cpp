#include <iostream>

int main(){
  int num = 1;
  int inputNum;
  int numInLine = 2;
  int dir = 1;
  int counter = 1;

  std::cin >> inputNum;

  while(num <= inputNum){

    for (int i = 0; i < counter && num <= inputNum; i++)
      std::cout << num++ << ' ';

    std::cout << std::endl;

    counter = counter + dir;

    if (counter < 1){
      dir = 1;
      counter = 2;
      ++numInLine;
    }else if (counter > numInLine){
      dir = -1;
      counter = counter - 2;
    }
  }
}