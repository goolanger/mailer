from flask import Flask, jsonify

app = Flask(__name__)

routes = [
  ('Users', '/api/users', [
    {'name':'User 1', 'email': 'user1@email.com', 'password': 'user1password'},
    {'name':'User 2', 'email': 'user2@email.com', 'password': 'user2password'},
    {'name':'User 3', 'email': 'user3@email.com', 'password': 'user3password'},
    {'name':'User 4', 'email': 'user4@email.com', 'password': 'user4password'}
  ])
]

for _, link, resources in routes:
  @app.route(link, methods=['GET', 'POST'])
  def rest_method():
    return jsonify(data=resources)

@app.route('/')
def index():
  return """
    This is an api example to get the results into original app<br>
    Routes are the nexts:<br>
    """ + "\n".join(["<a href=\"{0}\">{1}</a><br>".format(link, name) for name, link, _ in routes]) + """
    <br>
    All this data is fake, it was only created for testing the results.<br>
    Thanks for watching...
  """

if __name__ == '__main__':
  app.run(debug=True)