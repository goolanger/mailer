from colorama import Fore, Back, Style
import requests

TOKEN="qwlehfowiuq4hnir2qc342y34o8tu2n34ihmto35utc8924y5g9ot4u7t24y5"

def mail_url(path):
  return "http://localhost:5550{}".format(path)

def test_method(message):
  def func(function):
    def wrapper(*args, **kwargs):
      print(Fore.WHITE + "\n\t\t", message)
      return function(*args, **kwargs)
    return wrapper
  return func

def assert_equals(response, expected, success_text="test passed!"):
  if response == expected:
    print(Fore.GREEN + success_text)
  else:
    print(Fore.RED + success_text + " FAIL")

@test_method("Token Validation")
def test_token_validation():
  response = requests.get(mail_url('/token'), headers={ 'User-Token': TOKEN })
  assert_equals(response.status_code, 200)
  assert_equals(response.json(), {'Token': TOKEN})

@test_method("Token Reset")
def test_token_reset():
  pass

def main():
  test_token_validation()
  test_token_reset()

if __name__ == '__main__':
  main()