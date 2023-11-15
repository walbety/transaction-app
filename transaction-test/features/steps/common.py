from behave import then
from time import sleep
from json import JSONEncoder

@then("I wait {period} seconds")
def step_impl(_, period):
    sleep(int(period))


class RequestEncoder(JSONEncoder):
    def default(self, o):
        return o.__dict__
