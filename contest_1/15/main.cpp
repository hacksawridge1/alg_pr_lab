#include <iostream>
#include <string>

std::string encode(std::string str) {
  std::string encoding = "";
  int counter;

  for( int i = 0; str[i]; i++) {
    counter = 1;
    while (str[i] == str[i + 1]){
      counter++;
      i++;
    };
    if(counter > 1) {
    encoding += str[i] + std::to_string(counter);
    }else {
      encoding += str[i];
    };
  };
  return encoding;
}

int main(){
  std::string str;
  std::cin >> str;
  std::cout << encode(str);
}