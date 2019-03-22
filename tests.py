from colorama import Fore, Back, Style
import requests

TOKEN="qwlehfowiuq4hnir2qc342y34o8tu2n34ihmto35utc8924y5g9ot4u7t24y5"
NEWTOKEN="qwlehfowiuq4hnir2qc342y34o8tu2n34ihmto35utc8924y5g9ot4u7t2"

MAILCONFIG = {
  'SMTPOptions': {
    'SMTPAddress': 'smtp.gmail.com',
    'SMTPPort': '587',
    'From': '...@gmail.com',
    'Pass': '******'
  },
  'Mail': 'This is a testing purpose only email message',
  'Emails': [
    'amauryuh@gmail.com',
    'a.caballero@estudiantes.matcom.uh.cu',
    'inexistent@testing.com'
  ],
  'Threads': 2,
  'Retry': 3,
  'Wait': 1
}

def mail_url(path):
  return "http://localhost:5550{}".format(path)

def test_method(message):
  def func(function):
    def wrapper(*args, **kwargs):
      print(Fore.WHITE, "\n\t\t", message)
      return function(*args, **kwargs)
    return wrapper
  return func

def assert_equals(response, expected, success_text="test passed!"):
  if response == expected:
    print(Fore.GREEN, success_text)
  else:
    print(Fore.RED, success_text, "FAIL")

@test_method("Token Validation")
def test_token_validation():
  response = requests.get(mail_url('/token'), headers={ 'User-Token': TOKEN })
  assert_equals(response.status_code, 200)
  assert_equals(response.json(), {'Token': TOKEN})

  response = requests.get(mail_url('/token'), headers={ 'User-Token': NEWTOKEN })
  assert_equals(response.status_code, 401)

@test_method("Token Reset")
def test_token_reset():
  response = requests.put(mail_url('/set-token'), headers={ 'User-Token': TOKEN }, json={ 'Token': NEWTOKEN })
  assert_equals(response.status_code, 200)
  assert_equals(response.json()['Token'], NEWTOKEN)
  
  response = requests.get(mail_url('/token'), headers={ 'User-Token': NEWTOKEN })
  assert_equals(response.status_code, 200)
  assert_equals(response.json(), {'Token': NEWTOKEN})

  response = requests.get(mail_url('/token'), headers={ 'User-Token': TOKEN })
  assert_equals(response.status_code, 401)

@test_method("Send Mail")
def test_send_mail():
  response = requests.post(mail_url('/send'), headers={ 'User-Token': NEWTOKEN }, json=MAILCONFIG)
  assert_equals(response.status_code, 200)
  assert_equals(response.json()['Failed'][0], 'inexistent@testing.com')

def main():
  test_token_validation()
  test_token_reset()
  test_send_mail()
  print(Fore.WHITE, "\n\t\t", "End Tests")

if __name__ == '__main__':
  main()