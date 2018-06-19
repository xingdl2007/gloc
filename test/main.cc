/*
 * summary: a simple test file
 */

#include <map>
#include <vector>
#include <iostream>

using namespace std;

// macro with line-continuation character
#define TESTMACRO "multiple \
                  line test \
                  macro\n"

// hash table
class Solution {
public:
  vector<int> twoSum(vector<int> &nums, int target) {
    map<int, int> m;
    vector<int> indexs;

    for (int i = 0; i < nums.size(); ++i) {
      if (m.find(nums[i]) == m.end()) {
        m[target - nums[i]] = i;
      } else {
        indexs.push_back(m[nums[i]] + 1);
        indexs.push_back(i + 1);
        return indexs;
      }
    }
  }
};

/*
*  // embedded single-line comment in multiple comment
*
   empty line in multi-line comment

   line-continuation character in multi-line comment
   #define TESTMACRO "multiple \
                     line test \
                     macro\n"

    recursive
    "/* some string "
    "//----"
*/

void vector_dump(vector<int> &input) {
  for (vector<int>::iterator item = /* inline comment */ input.begin(); item != input.end(); item++)
    cout << *item << " ";
  cout << endl;  // useless comment
}

int main(int argc, char *argv[]) { /* single line comment */
  int numbers[] = {2, 7, 11, 15};  // ditto
  int target = 9;  /*
    multiple-line comment
    hello world
  */   /* some // recursive
  other
  comment
  */

  vector<int> nums(numbers, numbers + sizeof(numbers)/sizeof(int)); // /*dsfsdlfklsj df*/
  vector<int> indexs = (new Solution())->twoSum(nums, target);

  cout << "/* this line is /* */ not comment */\n\\";
  cout << "// ditto\n";

  /* prefix comment, same line with code */ vector_dump(nums);
  vector_dump(indexs);
  cout << TESTMACRO;

  cout << " multiple \
  line               \
  string is          \
  also ok            \
  \n";
  // 你好，世界
}
