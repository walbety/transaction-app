
from time import sleep
import logging
import requests
from behave import given, then, when
from features.steps.config import Config
from features.steps.common import RequestEncoder

@given("I save a single purchase with given values")
def step_impl(context):

    for item in context.table:
        amount = item.get("amount")
        date = item.get("date")
        description = item.get("description")

    purchase = {"amount" : amount,"description" : description,"date" : date,}

    logging.debug(f' purchase request: {purchase}')

    
    response = requests.post(url=f'http://{Config.get_transaction_url()}/v1/transaction',
                              json=purchase,
                              headers={"Content-Type": "application/json"})
    context.response = response
    logging.debug(f' Transaction Response({context.response.status_code}): {context.response.text}')
    
    context.response_json = response.json()
    logging.info(f' Transaction Response({context.response}): {context.response_json}')
    
@then('I expect the status response to be {status}')
def step_impl(context,status):
    assert (int(context.response.status_code) == int(status)
            ), f"Expected {status} on transaction request and got {context.response.status_code}"

@then('I expect the id of the purchase to be returned')
def step_impl(context):
    context.purchase_id = context.response_json['id']
    assert (context.purchase_id != ''), f"Expected id on transaction request response and got {context.purchase_id}"

@when('I send a request to get the given purchase using the id returned and the the currency as {currency}')
def step_impl(context,currency):
    context.response = requests.get(url=f'http://{Config.get_transaction_url()}/v1/transaction?id=' + context.purchase_id +'&currency='+currency)
    logging.debug(f' Transaction Response({context.response.status_code}): {context.response.text}')

@then('The returned response should have the given values')
def step_impl(context):
    for item in context.table:
        originalAmount = item.get("originalAmount")
        description = item.get("description")
        transactionDate = item.get("transactionDate")
        exchangeRate = item.get("exchangeRate")
        currency = item.get("currency")
        convertedAmount = item.get("convertedAmount")

    response = context.response.json()
    logging.debug(response)

    assert (response['originalAmount'] == originalAmount), f"expected amount: {originalAmount} got {response['originalAmount']}"
    assert (response['description'] == description), f"expected description: {description} got {response['description']}"
    assert (response['transactionDate'] == transactionDate), f"expected transactionDate: {transactionDate} got {response['transactionDate']}"
    assert (response['exchangeRate'] == exchangeRate), f"expected exchangeRate: {exchangeRate} got {response['exchangeRate']}"
    assert (response['currency'] == currency), f"expected currency: {currency} got {response['currency']}"
    assert (response['convertedAmount'] == convertedAmount), f"expected convertedAmount: {convertedAmount} got {response['convertedAmount']}"