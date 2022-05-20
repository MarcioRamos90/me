#!/bin/python3

import requests


def clean_input_arr(message):
	arr = input(message).split(',') 
	if len(arr) == 1 and arr[0] == "":
		arr = []
	return arr

def extract_key_value(arr):
	obj = {}
	for item in arr:
		key, value = item.split("=")
		obj[key] = value
	return obj


url = input("Enter the url to retrieve: ")
cookies_arr = clean_input_arr("set cookies: ")
headers_arr = clean_input_arr("set headers: ")
parameters_arr = clean_input_arr("set parameters: ")

cookies = extract_key_value(cookies_arr)
print('C -->', cookies)

headers = extract_key_value(headers_arr)
print('H -->', headers)

parameters = extract_key_value(parameters_arr)
print("P -->", parameters)

response = requests.post(url, cookies=cookies, headers=headers, data=parameters)

print(response.request.headers)
print(response.text)

