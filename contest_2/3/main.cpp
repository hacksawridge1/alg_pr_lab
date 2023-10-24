bool compare(const std::string& a, const std::string& b) {
  int num1 = std::count(a.begin(), a.end(), '1');
  int num2 = std::count(b.begin(), b.end(), '1');

  if (num1 == num2) {
    return stoi(a) < stoi(b);
  }
  else {
    return num1 > num2;
  }
}