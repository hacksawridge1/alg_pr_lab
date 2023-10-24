#include <iostream>
#include <string>
#include <algorithm>
#include <cmath>

using namespace std;

string action(string first, string second) {

    sort(first.begin(), first.end());
    sort(second.begin(), second.end());

    if (first == second){
        return "YES";
    } else{
        return "NO";
    }
}


int main() {
    string a, b;
    cin >> a >> b;
    string answer = action(a, b);
    cout << answer << endl;
}