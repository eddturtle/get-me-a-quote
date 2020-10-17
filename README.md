
# Quote Generator API

Tiny API to work as a random 'famous quote' generator. Each time it's called you will get a different quote.
Built as a way to keep periodic messages, like emails and slack notifications from cron jobs interesting.

Quotes are sourced from this list: https://github.com/umbrae/reddit-top-2.5-million/blob/master/data/quotes.csv

API Built by Edd Turtle ([designedbyaturtle.com](https://designedbyaturtle.com))

Served at: [getmeaquote.designedbyaturtle.com](http://getmeaquote.designedbyaturtle.com/) (or [https://get-me-a-quote.herokuapp.com](https://get-me-a-quote.herokuapp.com/) if you need https)

---

#### Content Type

Use `/?accept=json` for a json response, `/?accept=xml` for xml. Otherwise plain text will be returned.
